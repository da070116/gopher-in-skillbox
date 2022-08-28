package storage

import (
	"fmt"
	s "gopher-in-skillbox/module_app/student"
)

type StudentsStorage map[int]*s.Person

func NewStudentStorage() StudentsStorage {
	return make(map[int]*s.Person)
}

func (stor StudentsStorage) Put(idx int, s *s.Person) {
	stor[idx] = s
}

func (stor StudentsStorage) Get(idx int) (*s.Person, error) {
	st, exists := stor[idx]
	if !exists {
		return nil, fmt.Errorf("Нет студента с идентификатором %v\n", idx)
	} else {
		return st, nil
	}
}
