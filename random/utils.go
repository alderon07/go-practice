package main

import (
	"errors"
)

type number interface {
    uint64 | uint32 | uint16 | uint8 | int64 | int32 | int16 | int8 | float32 | float64
}


func SumArray[T number](array []T) (T, error) {
    var sum T
    
    if len(array) == 0{
        return 0, errors.New("the array is empty")
    }

    if len(array) == 1 {
        sum = array[0]
        return sum, nil
    }
    
    for _, elem := range array {
        sum += elem
    }

    return sum, nil
}


func Map[T any](array []T, f func(T) T) ([]T, error) {
	var result []T

	if len(array) == 0 {
		return nil, errors.New("the array is empty")
	}

	for _, elem := range array {
		result = append(result, f(elem))
	}

	return result, nil
}

func Filter[T any](array []T, f func(T) bool) ([]T, error) {
	var result []T

	if len(array) == 0 {
		return nil, errors.New("the array is empty")
	}

	for _, elem := range array {
		if f(elem) {
			result = append(result, elem)
		}
	}

	return result, nil
}
