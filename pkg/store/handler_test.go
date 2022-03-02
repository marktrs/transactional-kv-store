package store

import (
	"testing"

	"github.com/marktrs/transactional-kv-store/pkg/store/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	s := &mocks.MemoryStore{}
	h := NewStoreHandler(s)

	k := "foo"
	v := "bar"

	t.Run("handle operation", func(t *testing.T) {
		h.handleOperation("", "", "")
	})

	t.Run("handle operation SET", func(t *testing.T) {
		s.On("Set", k, v).Return()
		h.handleOperation("SET", k, v)
	})

	t.Run("handle operation GET", func(t *testing.T) {
		s.On("Get", k).Return(nil)
		h.handleOperation("GET", k, v)
	})

	t.Run("handle operation DELETE", func(t *testing.T) {
		s.On("Delete", k).Return()
		h.handleOperation("DELETE", k, v)
	})

	t.Run("handle operation COUNT", func(t *testing.T) {
		s.On("Count", k).Return()
		h.handleOperation("COUNT", k, v)
	})

	t.Run("handle operation BEGIN", func(t *testing.T) {
		s.On("Begin").Return()
		h.handleOperation("BEGIN", k, v)
	})

	t.Run("handle operation COMMIT", func(t *testing.T) {
		s.On("Commit").Return(nil)
		h.handleOperation("COMMIT", k, v)
	})

	t.Run("handle operation ROLLBACK", func(t *testing.T) {
		s.On("RollBack").Return(nil)
		h.handleOperation("ROLLBACK", k, v)
	})

	t.Run("validate operation", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{""})
		assert.Error(t, err)
		assert.Equal(t, err, ErrInvalidOperation)
	})

	t.Run("validate operation set invalid", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"SET"})
		assert.Error(t, err)
		assert.Equal(t, err, ErrInvalidParamenter)
	})

	t.Run("validate operation set", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"SET", k, v})
		assert.NoError(t, err)
	})

	t.Run("validate operation get invalid", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"GET"})
		assert.Error(t, err)
		assert.Equal(t, err, ErrInvalidParamenter)
	})

	t.Run("validate operation get", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"GET", "k"})
		assert.NoError(t, err)
	})

	t.Run("validate operation commit invalid", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"COMMIT", "1"})
		assert.Error(t, err)
		assert.Equal(t, err, ErrInvalidParamenter)
	})

	t.Run("validate operation commit", func(t *testing.T) {
		_, _, _, err := h.validateOperation([]string{"COMMIT"})
		assert.NoError(t, err)
	})
}

func ExampleHandler() {
	h := NewStoreHandler(&mocks.MemoryStore{})
	h.handleOperation("foo", "bar", "baz")
	// Output:ERROR: 'foo' operation unknown
}

func ExampleHandler_get() {
	s := &mocks.MemoryStore{}
	h := NewStoreHandler(s)
	s.On("Get", "baz").Return(ErrKeyNotSet)
	h.handleOperation("GET", "baz", "")
	// Output:key not set
}

func ExampleHandler_set() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("GET", "foo", "")
	// Output:bar
}
func ExampleHandler_delete() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("DELETE", "foo", "")
	h.handleOperation("GET", "foo", "")
	// Output:key not set
}

func ExampleHandler_commit() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "baz")
	h.handleOperation("COMMIT", "", "")
	h.handleOperation("GET", "foo", "")
	// Output:baz
}

func ExampleHandler_commit_nested() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "baz")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("COMMIT", "", "")
	h.handleOperation("GET", "foo", "")
	// Output:bar
}

func ExampleHandler_rollback() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "bar")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "baz")
	h.handleOperation("ROLLBACK", "", "")
	h.handleOperation("GET", "foo", "")
	// Output:bar
}

func ExampleHandler_rollback_nested() {
	s := NewMemoryStore()
	h := NewStoreHandler(s)
	h.handleOperation("SET", "foo", "123")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "456")
	h.handleOperation("BEGIN", "", "")
	h.handleOperation("SET", "foo", "789")
	h.handleOperation("ROLLBACK", "", "")
	h.handleOperation("ROLLBACK", "", "")
	h.handleOperation("GET", "foo", "")
	// Output:123
}
