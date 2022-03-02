package store

import (
	"errors"
	"fmt"

	"github.com/marktrs/transactional-kv-store/pkg/store/model"
)

type MemoryStore interface {
	Set(key, value string)
	Get(key string) error
	Delete(key string)
	Count(value string)
	Begin()
	Commit() error
	RollBack() error
}

var (
	ErrKeyNotSet error = errors.New("key not set")
)

type memoryStore struct {
	globalStore      map[string]string
	transactionStack model.TransactionStack
}

func NewMemoryStore() MemoryStore {
	return &memoryStore{
		globalStore:      make(map[string]string),
		transactionStack: model.NewTransactionStack(),
	}
}

func (m *memoryStore) Set(key string, value string) {
	ActiveTransaction := m.transactionStack.Peek()
	if ActiveTransaction == nil {
		m.globalStore[key] = value
	} else {
		ActiveTransaction.Store[key] = value
	}
}

func (m *memoryStore) Get(key string) error {
	ActiveTransaction := m.transactionStack.Peek()
	if ActiveTransaction == nil {
		if val, ok := m.globalStore[key]; ok {
			fmt.Println(val)
		} else {
			return ErrKeyNotSet
		}
	} else {
		if val, ok := ActiveTransaction.Store[key]; ok {
			fmt.Println(val)
		} else if val, ok := m.globalStore[key]; ok {
			fmt.Println(val)
		} else {
			return ErrKeyNotSet
		}
	}

	return nil
}

func (m *memoryStore) Delete(key string) {
	ActiveTransaction := m.transactionStack.Peek()
	if ActiveTransaction == nil {
		delete(m.globalStore, key)
	} else {
		delete(ActiveTransaction.Store, key)
	}
}

func (m *memoryStore) Count(value string) {
	var count int = 0
	ActiveTransaction := m.transactionStack.Peek()
	if ActiveTransaction == nil {
		for _, v := range m.globalStore {
			if v == value {
				count++
			}
		}
	} else {
		for _, v := range ActiveTransaction.Store {
			if v == value {
				count++
			}
		}
	}
	fmt.Println(count)
}

func (m *memoryStore) Begin() {
	m.transactionStack.PushTransaction()
}

func (m *memoryStore) Commit() error {
	if err := m.transactionStack.Commit(m.globalStore); err != nil {
		return err
	}

	if err := m.transactionStack.PopTransaction(); err != nil {
		return err
	}

	return nil
}

func (m *memoryStore) RollBack() error {
	return m.transactionStack.PopTransaction()
}
