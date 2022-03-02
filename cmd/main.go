package main

import (
	"github.com/marktrs/transactional-kv-store/pkg/store"
)

func main() {
	memoryStore := store.NewMemoryStore()
	storeHandler := store.NewStoreHandler(memoryStore)
	storeHandler.Handle()
}
