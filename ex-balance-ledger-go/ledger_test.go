package ledgerinterview

import (
	"testing"
)

// TestSimpleDeposit goes through depositing some money and expecting
// it to be present and correct
func TestSimpleDeposit(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	d1 := ledger.Deposit(100)
	assertNotEmptyEntryID(t, d1)
	assertBalanceEqual(t, 100, ledger.Balance())
}

// TestSimpleWithdrawal goes through withdrawing some money and expecting
// the balance to be reflected appropriately
func TestSimpleWithdrawal(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	w1 := ledger.Withdraw(100)
	assertNotEmptyEntryID(t, w1)
	assertBalanceEqual(t, int(-100), ledger.Balance())
}

// TestSimpleBalance1 goes through a chain of making a deposit and withdrawal
// to assert that the flow works correctly
func TestSimpleBalance1(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	d1 := ledger.Deposit(100)
	assertNotEmptyEntryID(t, d1)
	assertBalanceEqual(t, 100, ledger.Balance())

	w1 := ledger.Withdraw(10)
	assertEntryIDsNotEqual(t, d1, w1)
	assertBalanceEqual(t, 90, ledger.Balance())
}

// TestSimpleBalance2 goes through a chain of making a withdrawal and deposit
// to assert that the flow works correctly
func TestSimpleBalance2(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	w1 := ledger.Withdraw(10)
	assertBalanceEqual(t, int(-10), ledger.Balance())

	d1 := ledger.Deposit(100)
	assertEntryIDsNotEqual(t, d1, w1)
	assertBalanceEqual(t, 90, ledger.Balance())
}

// TestBalanceAt asserts that we are correctly tracking balances between
// operations correctly, so we can retrieve balances at a point in time
func TestBalanceAt(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	d1 := ledger.Deposit(100)
	assertBalanceEqual(t, 100, ledger.Balance())

	balance, err := ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	w1 := ledger.Withdraw(10)
	assertEntryIDsNotEqual(t, d1, w1)
	assertBalanceEqual(t, 90, ledger.Balance())

	balance, err = ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	balance, err = ledger.BalanceAt(w1)
	assertNoError(t, err)
	assertBalanceEqual(t, 90, balance)
}

// TestBalanceAtInvalidID ensures that a bad transaction ID results in an error
func TestBalanceAtInvalidID(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())
	balance, err := ledger.BalanceAt("BAD_ID")
	assertError(t, err)
	assertBalanceEqual(t, 0, balance)
}

// TestTransactionFlow goes through a simple transaction flow
func TestTransactionFlow(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Start our transaction
	ledger.Begin()

	// Make a deposit and a withdrawak
	d1 := ledger.Deposit(100)
	w1 := ledger.Withdraw(10)
	assertBalanceEqual(t, 90, ledger.Balance())
	assertEntryIDsNotEqual(t, d1, w1)

	// Commit our transaction
	err := ledger.Commit()
	assertNoError(t, err)

	// Expect all of it to have been written appropriately
	assertBalanceEqual(t, 90, ledger.Balance())

	balance, err := ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	balance, err = ledger.BalanceAt(w1)
	assertNoError(t, err)
	assertBalanceEqual(t, 90, balance)
}

// TestTransactionRollback tests writing a transaction and then calling
// rollback, making sure that nothing is committed
func TestTransactionRollback(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Start our transaction
	ledger.Begin()

	// Make a deposit and a withdrawal
	d1 := ledger.Deposit(100)
	w1 := ledger.Withdraw(10)
	assertBalanceEqual(t, 90, ledger.Balance())
	assertEntryIDsNotEqual(t, d1, w1)

	// Rollback our transaction
	err := ledger.Rollback()
	assertNoError(t, err)

	// Expect none of it to have been written
	assertBalanceEqual(t, 0, ledger.Balance())

	_, err = ledger.BalanceAt(d1)
	assertError(t, err)

	_, err = ledger.BalanceAt(w1)
	assertError(t, err)
}

// TestTransactionNested tests writing nested transactions with commit
func TestTransactionNested(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Start our transaction, create a deposit
	ledger.Begin()
	d1 := ledger.Deposit(100)

	// Create a new transaction
	ledger.Begin()
	w1 := ledger.Withdraw(10)
	assertBalanceEqual(t, 90, ledger.Balance())

	// Commit
	err := ledger.Commit()
	assertNoError(t, err)

	// Expect all our data to be written
	assertBalanceEqual(t, 90, ledger.Balance())

	balance, err := ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	balance, err = ledger.BalanceAt(w1)
	assertNoError(t, err)
	assertBalanceEqual(t, 90, balance)
}

// TestTransactionNestedRollback tests writing nested transactions with a
// rollback before committing
func TestTransactionNestedRollback(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Start our transaction, create a deposit
	ledger.Begin()
	d1 := ledger.Deposit(100)

	// Create a new transaction but roll it back
	ledger.Begin()
	w1 := ledger.Withdraw(10)
	assertBalanceEqual(t, 90, ledger.Balance())
	err := ledger.Rollback()
	assertNoError(t, err)

	// Expect only d1 to be written
	assertBalanceEqual(t, 100, ledger.Balance())
	balance, err := ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	balance, err = ledger.BalanceAt(w1)
	assertError(t, err)
	assertBalanceEqual(t, 0, balance)

	// Do a commit, expect only d1 again
	err = ledger.Commit()
	assertNoError(t, err)
	assertBalanceEqual(t, 100, ledger.Balance())
	balance, err = ledger.BalanceAt(d1)
	assertNoError(t, err)
	assertBalanceEqual(t, 100, balance)

	balance, err = ledger.BalanceAt(w1)
	assertError(t, err)
	assertBalanceEqual(t, 0, balance)
}

// TestTransactionNestedFullRollback tests writing nested transactions but
// with all transactions roll backed, we expect nothing to be written
func TestTransactionNestedFullRollback(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Start our transaction, create a deposit
	ledger.Begin()
	d1 := ledger.Deposit(100)

	// Create a new transaction
	ledger.Begin()
	w1 := ledger.Withdraw(10)

	// Now rollback both our transactions
	err := ledger.Rollback()
	assertNoError(t, err)
	err = ledger.Rollback()
	assertNoError(t, err)

	// We shouldn't be in a transaction, so expect an error
	assertBalanceEqual(t, 0, ledger.Balance())
	err = ledger.Commit()
	assertError(t, err)

	// Make sure we did not write d1 or w1
	assertBalanceEqual(t, 0, ledger.Balance())
	balance, err := ledger.BalanceAt(d1)
	assertError(t, err)
	assertBalanceEqual(t, 0, balance)
	balance, err = ledger.BalanceAt(w1)
	assertError(t, err)
	assertBalanceEqual(t, 0, balance)
}

// TestTransactionBlockInactiveError tests for appropriate errors if COMMIT
// or ROLLBACK are called outside of a transaction
func TestTransactionBlockInactiveError(t *testing.T) {
	ledger := NewLedger()
	assertBalanceEqual(t, 0, ledger.Balance())

	// Commit outside an active transaction should cause an error
	err := ledger.Commit()
	assertError(t, err)

	// Rollback outside an active transaction should cause an error
	err = ledger.Rollback()
	assertError(t, err)
}

// assertBalanceEqual takes two int values and expects them to be equal
func assertBalanceEqual(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("expected balance to be %d, got %d", expected, actual)
	}
}

// assertNotEmptyEntryID takes in an entry ID and expects it to be non-empty
func assertNotEmptyEntryID(t *testing.T, entryID string) {
	t.Helper()
	if len(entryID) == 0 {
		t.Fatalf("expected non-empty entry ID string")
	}
}

// assertEntryIDsNotEqual takes in two entry IDs and ensures they are different
// This will also check to make sure they are not empty
func assertEntryIDsNotEqual(t *testing.T, a, b string) {
	t.Helper()

	assertNotEmptyEntryID(t, a)
	assertNotEmptyEntryID(t, b)

	if a == b {
		t.Fatalf("expected both strings (%v) to be different", a)
	}
}

// assertError takes in an error and makes sure it is not nil
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// assertNoError takes in an error and makes sure it is nil / empty
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
