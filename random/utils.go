package main

import (
    "errors"
)

// Number is a type constraint that allows any of the specified numeric types.
type Number interface {
    uint64 | uint32 | uint16 | uint8 | int64 | int32 | int16 | int8 | float32 | float64
}

// SumArray calculates the sum of the elements in the provided array.
// It returns an error if the array is empty.
//
// Parameters:
// - array: A slice of elements of a numeric type that satisfies the Number constraint.
//
// Returns:
// - The sum of the elements in the array.
// - An error if the array is empty.
func SumArray[T Number](array []T) (T, error) {
    var sum T

    if len(array) == 0 {
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

// Map applies a given function to each element of the provided array and returns a new array with the results.
// It returns an error if the array is empty.
//
// Parameters:
// - array: A slice of elements of any type.
// - f: A function that takes an element of type T and returns an element of type T.
//
// Returns:
// - A new slice with the results of applying the function to each element of the original array.
// - An error if the array is empty.
func Map[T any, P any](values []T, f func(T) P) []P {
    res := make([]P, 0, len(values))

    for _, v := range values {
        res = append(res, f(v))
    }

    return res
}

// Filter returns a new array containing only the elements of the provided array that satisfy the given predicate function.
// It returns an error if the array is empty.
//
// Parameters:
// - array: A slice of elements of any type.
// - f: A predicate function that takes an element of type T and returns a boolean indicating whether the element should be included in the result.
//
// Returns:
// - A new slice with the elements that satisfy the predicate function.
// - An error if the array is empty.
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
