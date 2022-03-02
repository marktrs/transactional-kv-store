package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionStack(t *testing.T) {
	ts := NewTransactionStack()
	globalStore := make(map[string]string)

	t.Run("commit empty stack", func(t *testing.T) {
		err := ts.Commit(globalStore)
		assert.Error(t, err)
		assert.Equal(t, err, ErrNoTransaction)
	})

	t.Run("pop empty stack", func(t *testing.T) {
		err := ts.PopTransaction()
		assert.Error(t, err)
		assert.Equal(t, err, ErrNoTransaction)
	})

	t.Run("get stack peek", func(t *testing.T) {
		ts.Peek()
	})

	t.Run("push trasaction", func(t *testing.T) {
		ts.PushTransaction()
	})

	t.Run("push trasaction", func(t *testing.T) {
		ts.PushTransaction()
	})

	t.Run("commit transaction", func(t *testing.T) {
		assert.NoError(t, ts.Commit(globalStore))
	})

	t.Run("commit transaction", func(t *testing.T) {
		assert.NoError(t, ts.Commit(globalStore))
	})

	t.Run("pop trasaction", func(t *testing.T) {
		assert.NoError(t, ts.PopTransaction())
	})

}
