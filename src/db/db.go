package db

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DBH struct {
	db *sqlx.DB
}

type balance struct {
	Value float64 `db:"value"`
}
type Transaction struct {
	ID        int    `db:"id"`
	UUID      string `db:"uuid"`
	Amount    string `db:"amount"`
	State     string `db:"state"`
	BalanceID int    `db:"balance_id"`
}

type DB interface {
	Add(ctx context.Context, balanceID int, transactionID, amount string) error
	Subtract(ctx context.Context, balanceID int, transactionID, amount string) error
	GetNLastOddTransactions(n int) ([]Transaction, error)
	CancelTransactions(transactions []Transaction) error
}

func Init(connStr string) (*DBH, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}
	return &DBH{db}, err
}

func (d *DBH) Add(ctx context.Context, balanceID int, transactionID, amount string) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("can't begin database transaction: %w", err)
	}
	if err := d.addTransaction(tx, "win", balanceID, transactionID, amount); err != nil {
		tx.Rollback()
		return fmt.Errorf("can't add transaction: %w", err)
	}
	if err := d.addBalance(tx, balanceID, amount); err != nil {
		tx.Rollback()
		return fmt.Errorf("can't update balance: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("can't commit transaction: %w", err)
	}
	return nil
}

func (d *DBH) Subtract(ctx context.Context, balanceID int, transactionID, amount string) error {
	bal := new(balance)
	err := d.db.GetContext(ctx, bal, `SELECT value FROM balances WHERE id=$1`, balanceID)
	if err != nil {
		return fmt.Errorf("can't select balance from db: %w", err)
	}
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("can't parse float %s: %w", amount, err)
	}
	if (bal.Value - amountFloat) < 0 {
		return fmt.Errorf("can't subtract %v from %v: %w", amountFloat, bal.Value, ErrNegativeBalance)
	}
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}
	if err := d.addTransaction(tx, "lost", balanceID, transactionID, amount); err != nil {
		tx.Rollback()
		return fmt.Errorf("can't add transaction: %w", err)
	}
	if err := d.subBalance(tx, balanceID, amount); err != nil {
		tx.Rollback()
		return fmt.Errorf("can't update balance: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("can't commit transaction: %w", err)
	}
	return nil
}

func (d *DBH) GetNLastOddTransactions(n int) ([]Transaction, error) {
	var transactions []Transaction
	err := d.db.Select(&transactions, `SELECT id, uuid, amount, state, balance_id FROM transactions WHERE (NOT canceled) AND id%2!=0 ORDER BY id DESC LIMIT $1`, n)
	return transactions, err
}

func (d *DBH) CancelTransactions(transactions []Transaction) error {
	transactionIDs := make([]int, len(transactions))
	for i, t := range transactions {
		transactionIDs[i] = t.ID
	}
	tx, err := d.db.Beginx()
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	query, args, _ := sqlx.In(`UPDATE transactions SET canceled=TRUE WHERE id IN (?)`, transactionIDs)
	query = d.db.Rebind(query)
	if _, err := tx.Exec(query, args...); err != nil {
		tx.Rollback()
		return fmt.Errorf("can't update transactions: %w", err)
	}

	// correcting balances in a for loop
	// start with lost transactions so we won't violate sql constrain balance check
	for _, transaction := range transactions {
		if transaction.State == "lost" {
			err = d.addBalance(tx, transaction.BalanceID, transaction.Amount)
		}
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("can't update balance: %w", err)
		}
	}
	for _, transaction := range transactions {
		if transaction.State == "win" {
			err = d.subBalance(tx, transaction.BalanceID, transaction.Amount)
		}
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("can't update balance: %w", err)
		}
	}
	tx.Commit()
	return nil
}

func (DBH) subBalance(tx *sqlx.Tx, balanceID int, amount string) error {
	fmt.Println("del balance, sum:", amount)
	_, err := tx.Exec(`UPDATE balances SET value=value-$1 WHERE id=$2`, amount, balanceID)
	return err
}
func (DBH) addBalance(tx *sqlx.Tx, balanceID int, amount string) error {
	fmt.Println("add balance, sum:", amount)
	_, err := tx.Exec(`UPDATE balances SET value=value+$1 WHERE id=$2`, amount, balanceID)
	return err
}

func (DBH) addTransaction(tx *sqlx.Tx, state string, balanceID int, transactionID, amount string) error {
	_, err := tx.Exec(`INSERT INTO transactions (uuid, state, balance_id, amount) VALUES ($1, $2, $3, $4)`, transactionID, state, balanceID, amount)
	if err, ok := err.(*pq.Error); ok {
		if err.Constraint == "transactions_uuid_key" {
			return ErrUUIDExists
		}
	}
	return err
}
