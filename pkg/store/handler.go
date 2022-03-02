package store

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidOperation  error = errors.New("invalid operation")
	ErrInvalidParamenter error = errors.New("not enough parameter")
)

type Handler interface {
	Handle()
	handleOperation(ops, key, value string)
	validateOperation(args []string) (ops, key, value string, err error)
}

type storeHandler struct {
	repo Repository
}

func NewStoreHandler(repo Repository) Handler {
	return &storeHandler{repo}
}

func (h *storeHandler) Handle() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		args := strings.Fields(text)
		if len(args) == 0 {
			continue
		}

		ops, key, value, err := h.validateOperation(args)
		if err != nil {
			fmt.Println(err)
			continue
		}

		h.handleOperation(ops, key, value)
	}
}

func (h *storeHandler) handleOperation(ops, key, value string) {
	var err error
	switch ops {
	case "SET":
		h.repo.Set(key, value)
	case "GET":
		err = h.repo.Get(key)
	case "DELETE":
		h.repo.Delete(key)
	case "COUNT":
		h.repo.Count(key)
	case "BEGIN":
		h.repo.Begin()
	case "COMMIT":
		err = h.repo.Commit()
	case "ROLLBACK":
		err = h.repo.RollBack()
	default:
		fmt.Printf("ERROR: '%s' operation unknown\n", ops)
	}

	if err != nil {
		fmt.Println(err)
	}
}

func (h *storeHandler) validateOperation(args []string) (ops, key, value string, err error) {
	ops = args[0]

	switch ops {
	case "SET":
		err = validateParam(args, 2)
		if err != nil {
			return ops, key, value, err
		}
		key = args[1]
		value = args[2]
	case "GET", "DELETE", "COUNT":
		err = validateParam(args, 1)
		if err != nil {
			return ops, key, value, err
		}
		key = args[1]
	case "BEGIN", "ROLLBACK", "COMMIT":
		err = validateParam(args, 0)
		if err != nil {
			return ops, key, value, err
		}
	default:
		return ops, key, value, ErrInvalidOperation
	}

	return ops, key, value, nil
}

func validateParam(args []string, requiredCount int) error {
	if len(args)-1 != requiredCount {
		return ErrInvalidParamenter
	}

	return nil
}
