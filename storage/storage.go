package storage

type Storage interface {
	PutObject(key, obj any)
	GetObject(key any) (any, bool)
	GetAll() []any
}
