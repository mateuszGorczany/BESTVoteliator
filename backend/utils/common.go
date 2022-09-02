package common

import (
	"errors"
)

var (
	ErrorNotImplementedYet = errors.New("Not implemented yet.")
)

type ID_t int64

type Stream [T any]struct {
	value []T
}

func (s Stream[T]) Collect() []T {
	return s.value
}


func (s *Stream[T]) Map(function func(element T) T) *Stream[T] {
	newList := make([]T, len(s.value))
	for i, element := range s.value {
		newList[i] = function(element)
	}
	return &Stream[T]{newList}
}


func (s *Stream[T]) Filter(filterFunction func(element T) bool) *Stream[T] {
	var newList []T
	for _, element := range s.value {
		if filterFunction(element) {
			newList = append(newList, element)
		}
	}
	return &Stream[T]{newList}
}