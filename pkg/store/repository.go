package store

type Repository interface {
	Set(key, value string)
	Get(key string) error
	Delete(key string)
	Count(value string)
	Begin()
	Commit() error
	RollBack() error
}
