package workers

import (
	"time"

	"github.com/darwinnn/TestTask/src/db"
)

type logFunc func(string, ...interface{})

type Canceller struct {
	ticker        *time.Ticker
	dbh           db.DB
	nTransactions int
	logger        logFunc
}

func Init(dbh db.DB, logger logFunc, d time.Duration, nTransactions int) *Canceller {
	return &Canceller{
		ticker:        time.NewTicker(d),
		dbh:           dbh,
		nTransactions: nTransactions,
		logger:        logger,
	}
}

func (c *Canceller) Work() {
	for _ = range c.ticker.C {
		transactions, err := c.dbh.GetNLastOddTransactions(c.nTransactions)
		if err != nil {
			c.logger("ERR: can't get last %d transactions from database: %v", c.nTransactions, err)
			continue
		}
		if len(transactions) < c.nTransactions {
			// don't cancel transactions if there's not exactly nTransactions left in the database
			c.logger("won't cancel %d transactions, since the value is less than nTransactions %d", len(transactions), c.nTransactions)
			continue
		}
		c.logger("Canceling transactions: %+v", transactions)
		if err := c.dbh.CancelTransactions(transactions); err != nil {
			c.logger("ERR: can't cancel transactions: %v", err)
			continue
		}
		c.logger("Finished canceling transactions")
	}
}
