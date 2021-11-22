package storage

type Storage interface {
	Save(v interface{})
	Load(v interface{}) error
}

type noOpStorage struct {
}

func (n *noOpStorage) Save(interface{}) {
}

func (n *noOpStorage) Load(interface{}) error {
	return nil
}

func NewNoOpStorage() Storage {
	return &noOpStorage{}
}

var _ = NewNoOpStorage
