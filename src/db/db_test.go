package db

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestAdd(t *testing.T) {
	dbh, err := Init("user=test_task_testdb host=localhost password=postgres dbname=test_task_testdb sslmode=disable")
	if err != nil {
		t.Fatalf("can't connect to test db: %v", err)
	}
	defer dbh.db.Close()
	ctx := context.Background()
	randUUID, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("can't generate uuid: %v", err)
	}
	expectedT := Transaction{
		UUID:      randUUID.String(),
		Amount:    "55.5",
		State:     "win",
		Canceled:  false,
		BalanceID: 1,
	}
	var resultB1 balance
	var resultB2 balance
	if err = dbh.db.Get(&resultB1, `SELECT value FROM balances WHERE id=1`); err != nil {
		t.Fatalf("can't select from test db balances: %v", err)
	}
	var resultT []Transaction
	if err := dbh.Add(ctx, 1, randUUID.String(), "55.5"); err != nil {
		t.Fatalf("adding failed: %v", err)
	}
	if err := dbh.Add(ctx, 1, randUUID.String(), "66.6"); !errors.Is(err, ErrUUIDExists) {
		t.Errorf("expected error ErrUUIDExists is not thrown, got: %v", err)
	}
	if err := dbh.db.Select(&resultT, `SELECT * FROM transactions WHERE uuid=$1`, randUUID); err != nil {
		t.Fatalf("can't select from test db transactions: %v", err)
	}
	if err = dbh.db.Get(&resultB2, `SELECT value FROM balances WHERE id=1`); err != nil {
		t.Fatalf("can't select from test db balances: %v", err)
	}
	if (resultB1.Value + 55.5) != resultB2.Value {
		t.Fatalf("expected balance %v, got %v", (resultB1.Value + 55.5), resultB2.Value)
	}
	if len(resultT) != 1 {
		t.Fatalf("expect to have 1 result, got %d", len(resultT))
	}
	expectedT.ID = resultT[0].ID
	if expectedT != resultT[0] {
		t.Fatalf("expected does not match with result: %+v != %+v", expectedT, resultT[0])
	}

}

func TestSubtract(t *testing.T) {
	dbh, err := Init("user=test_task_testdb host=localhost password=postgres dbname=test_task_testdb sslmode=disable")
	if err != nil {
		t.Fatalf("can't connect to test db: %v", err)
	}
	defer dbh.db.Close()
	ctx := context.Background()
	randUUID, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("can't generate uuid: %v", err)
	}
	expectedT := Transaction{
		UUID:      randUUID.String(),
		Amount:    "55.5",
		State:     "lost",
		Canceled:  false,
		BalanceID: 1,
	}
	var resultB1 balance
	var resultB2 balance
	if err = dbh.db.Get(&resultB1, `SELECT value FROM balances WHERE id=1`); err != nil {
		t.Fatalf("can't select from test db balances: %v", err)
	}
	var resultT []Transaction
	if err := dbh.Subtract(ctx, 1, randUUID.String(), "99999999"); !errors.Is(err, ErrNegativeBalance) {
		t.Fatalf("expected error ErrNegativeBalance is not thrown, got: %v", err)
	}
	if err := dbh.Subtract(ctx, 1, randUUID.String(), "55.5"); err != nil {
		t.Fatalf("subtracting failed: %v", err)
	}
	if err := dbh.Subtract(ctx, 1, randUUID.String(), "66.6"); !errors.Is(err, ErrUUIDExists) {
		t.Errorf("expected error ErrUUIDExists is not thrown, got: %v", err)
	}
	if err := dbh.db.Select(&resultT, `SELECT * FROM transactions WHERE uuid=$1`, randUUID); err != nil {
		t.Fatalf("can't select from test db transactions: %v", err)
	}
	if err = dbh.db.Get(&resultB2, `SELECT value FROM balances WHERE id=1`); err != nil {
		t.Fatalf("can't select from test db balances: %v", err)
	}
	if (resultB1.Value - 55.5) != resultB2.Value {
		t.Fatalf("expected balance %v, got %v", (resultB1.Value - 55.5), resultB2.Value)
	}
	if len(resultT) != 1 {
		t.Fatalf("expect to have 1 result, got %d", len(resultT))
	}

	expectedT.ID = resultT[0].ID
	if expectedT != resultT[0] {
		t.Fatalf("expected does not match with result: %+v != %+v", expectedT, resultT[0])
	}
}
