package model

import (
	"errors"
)

type TransactionStack interface {
	PushTransaction()
	PopTransaction() error
	Peek() *Transaction
	Commit(globalStore map[string]string) error
}

type transactionStack struct {
	top  *Transaction
	size int
}

var (
	ErrNoTransaction error = errors.New("no transaction")
)

func NewTransactionStack() TransactionStack {
	return &transactionStack{}
}

func (ts *transactionStack) PushTransaction() {
	temp := Transaction{Store: make(map[string]string)}
	temp.Next = ts.top
	ts.top = &temp
	ts.size++
}

func (ts *transactionStack) PopTransaction() error {
	if ts.top == nil {
		return ErrNoTransaction
	} else {
		ts.top = ts.top.Next
		ts.size--
	}

	return nil
}

func (ts *transactionStack) Peek() *Transaction {
	return ts.top
}

func (ts *transactionStack) Commit(globalStore map[string]string) error {
	ActiveTransaction := ts.Peek()
	if ActiveTransaction != nil {
		for key, value := range ActiveTransaction.Store {
			globalStore[key] = value
			if ActiveTransaction.Next != nil {
				ActiveTransaction.Next.Store[key] = value
			}
		}
	} else {
		return ErrNoTransaction
	}

	return nil
}
