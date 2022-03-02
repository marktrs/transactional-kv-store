package store

import (
	"testing"

	"github.com/marktrs/transactional-kv-store/pkg/store/model"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore(t *testing.T) {
	m := NewMemoryStore()

	k := "foo"
	v := "bar"

	t.Run("set", func(t *testing.T) {
		m.Set(k, v)
	})

	t.Run("get", func(t *testing.T) {
		assert.NoError(t, m.Get(k))
	})

	t.Run("get non-exist key", func(t *testing.T) {
		assert.Error(t, m.Get("baz"))
	})

	t.Run("count", func(t *testing.T) {
		m.Count(v)
	})

	t.Run("commit empty", func(t *testing.T) {
		err := m.Commit()
		assert.Error(t, err)
		assert.Equal(t, err, model.ErrNoTransaction)
	})

	t.Run("rollback empty", func(t *testing.T) {
		err := m.RollBack()
		assert.Error(t, err)
		assert.Equal(t, err, model.ErrNoTransaction)
	})

	t.Run("begin", func(t *testing.T) {
		m.Begin()
	})

	t.Run("set", func(t *testing.T) {
		m.Set(k, "baz")
	})

	t.Run("begin", func(t *testing.T) {
		m.Begin()
	})

	t.Run("set", func(t *testing.T) {
		m.Set(v, "baz")
	})

	t.Run("get", func(t *testing.T) {
		assert.NoError(t, m.Get(k))
	})

	t.Run("count", func(t *testing.T) {
		m.Count(v)
	})

	t.Run("commit", func(t *testing.T) {
		assert.NoError(t, m.Commit())
	})

	t.Run("delete", func(t *testing.T) {
		m.Delete(k)
	})

	t.Run("begin", func(t *testing.T) {
		m.Begin()
	})

	t.Run("set", func(t *testing.T) {
		m.Set("baz", v)
	})

	t.Run("commit", func(t *testing.T) {
		assert.NoError(t, m.Commit())
	})

	t.Run("count", func(t *testing.T) {
		m.Count(v)
	})

	t.Run("rollback", func(t *testing.T) {
		assert.NoError(t, m.RollBack())
	})

	t.Run("delete", func(t *testing.T) {
		m.Delete(v)
	})
}

func ExampleNewMemoryStore() {
	m := NewMemoryStore()

	k := "foo"
	v := "bar"

	m.Set(k, v)
	err := m.Get(k)
	if err != nil {
		panic(err)
	}
	// Output: bar
}

func ExampleNewMemoryStore_forth() {
	m := NewMemoryStore()

	k := "foo"
	v := "bar"

	m.Set(k, v)
	m.Count(v)
	// Output: 1
}
