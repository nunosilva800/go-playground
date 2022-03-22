package ledgerinterview

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Ledger interface {
	// Top Level Operations --------------------------------------------

	// Balance returns an int of the most recent total balance at this
	// point in time
	Balance() (amount int)

	// Deposit records a deposit in to the ledger for amount. Returns
	// an unique entry ID string to identify the deposit
	Deposit(amount uint) (entryID string)

	// Withdraw records a withdrawal in to the ledger for amount. Returns
	// an unique entry ID string to identify the withdrawal
	Withdraw(amount uint) (entryID string)

	// Additional Operations -------------------------------------------

	// BalanceAt returns an int of the total balance at the point (and
	// including) of a particular entry ID. Returns an error if that
	// entry ID was not in the ledger
	BalanceAt(entryID string) (amount int, err error)

	// Transaction Operations ------------------------------------------

	// Begin starts a transaction block. Transactions can be nested
	Begin()

	// Commit finishes and writes (commits) all open transactions. If Commit
	// is called without a transaction being started, it returns an error
	Commit() error

	// Rollback finishes the current active transaction block but discards all
	// the changes. If Rollback is called without a transaction being started,
	// it returns an error
	Rollback() error
}

type historyElem struct {
	id        string
	operation string
	value     uint
}

type transaction struct {
	balance int
	history []historyElem
}

type inMemoryLedger struct {
	balance int
	history []historyElem

	transactions []transaction
}

// NewLedger returns an implementation of the Ledger interface
func NewLedger() Ledger {
	return &inMemoryLedger{
		balance: 0,
		history: []historyElem{},
	}
}

func (m *inMemoryLedger) Balance() (amount int) {
	if len(m.transactions) != 0 {
		return m.transactions[len(m.transactions)-1].balance
	}
	return m.balance
}

func (m *inMemoryLedger) Deposit(amount uint) (entryID string) {
	id := generateEntryID()
	newHist := historyElem{
		id:        id,
		operation: "deposit",
		value:     amount,
	}

	if len(m.transactions) != 0 {
		currTransaction := m.transactions[len(m.transactions)-1]

		m.transactions[len(m.transactions)-1] = append(
			currTransaction,
			newHist,
		)
		return id
	}

	m.history = append(m.history, newHist)
	m.balance += int(amount)

	return id
}

func (m *inMemoryLedger) Withdraw(amount uint) (entryID string) {
	id := generateEntryID()
	newHist := historyElem{
		id:        id,
		operation: "withdraw",
		value:     amount,
	}

	if len(m.transactions) != 0 {
		m.transactions[len(m.transactions)-1] = append(
			m.transactions[len(m.transactions)-1],
			newHist,
		)
		return id
	}

	m.history = append(m.history, newHist)
	m.balance -= int(amount)

	return id
}

func (m *inMemoryLedger) BalanceAt(entryID string) (amount int, err error) {
	found := false

	for _, elem := range m.history {
		switch elem.operation {
		case "deposit":
			amount += int(elem.value)
		case "withdraw":
			amount -= int(elem.value)
		default:
			return 0, errors.New("unsupported operation: " + elem.operation)
		}

		if elem.id == entryID {
			found = true
			break
		}
	}

	if !found {
		return 0, errors.New("bad id")
	}

	return amount, nil
}

func (m *inMemoryLedger) Begin() {
	m.transactions = append(m.transactions, make(transaction, 0))
	return
}

func (m *inMemoryLedger) Commit() error {
	txAmount := 0
	for _, t := range m.transactions {
		for _, elem := range t {
			switch elem.operation {
			case "deposit":
				// txAmount += int(elem.value)

			case "withdraw":
				// txAmount -= int(elem.value)

			default:
				return errors.New("unsupported operation: " + elem.operation)
			}

			m.history = append(m.history, elem)
		}
	}

	m.balance += txAmount

	return nil
}

func (m *inMemoryLedger) Rollback() error {
	// TODO: implement me
	return nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

// generateEntryID is a helper function provided to generate a random ID
// to use for an entry. You can assume that these are completely unique
func generateEntryID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("err generating entry id: %v", err)
	}
	return fmt.Sprintf("entry-%x", b[0:8])
}
