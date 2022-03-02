package model

type Transaction struct {
	Store map[string]string
	Next  *Transaction
}
