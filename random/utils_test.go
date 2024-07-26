package main

import (
	"testing"
)

func TestSumArray(t *testing.T) {
    var intArray []int64 = []int64{1, 2, -3, 4, 5}
    var floatArray []float64 = []float64{1.5, 2.5, -4.0, 5.5}
    var uIntArray []uint64 = []uint64{4, 4, 4}

    uIntSumResult, err := SumArray(uIntArray)
    var uIntSumExpected uint64 = 12
    if uIntSumResult != uIntSumExpected || err != nil {
        t.Errorf("Sum of unsinged int array %v = %d; want %d", uIntArray, uIntSumResult, uIntSumExpected)
    }

    intSumResult, err := SumArray(intArray)
    var intSumExpected int64 = 9
    if intSumResult != intSumExpected || err != nil {
        t.Errorf("Sum of int array %v = %d; want %d", uIntArray, uIntSumResult, uIntSumExpected)
    }

    floatSumResult, err := SumArray(floatArray)
    var floatSumExpected float64 = 5.5
    if floatSumResult != floatSumExpected || err != nil {
        t.Errorf("Sum of float array %v = %f; want %f", floatArray, floatSumResult, floatSumExpected)
    }
}

func TestMap(t *testing.T) {
    intArray := []int64{1, 2, -3, 4, 5}
    floatArray := []float64{1.5, 2.5, -4.0, 5.5}
    uIntArray := []uint64{4, 4, 4}
    stringArray := []string{"I", "am", "a", "bug"}

    double := func(x int64) int64 {
        return x * 2
    }

    doubleFloat := func(x float64) float64 {
        return x * 2
    }

    doubleUInt := func(x uint64) uint64 {
        return x * 2
    }

    addBug := func(x string) string {
        return x + "bug"
    }

    intArrayResult := Map(intArray, double)
    intArrayExpected := []int64{2, 4, -6, 8, 10}
    for i := 0; i < len(intArrayResult); i++ {
        if intArrayResult[i] != intArrayExpected[i] {
            t.Errorf("Map of int array %v = %v; want %v", intArray, intArrayResult, intArrayExpected)
            break
        }
    }

    floatArrayResult := Map(floatArray, doubleFloat)
    floatArrayExpected := []float64{3.0, 5.0, -8.0, 11.0}
    for i := 0; i < len(floatArrayResult); i++ {
        if floatArrayResult[i] != floatArrayExpected[i] {
            t.Errorf("Map of float array %v = %v; want %v", floatArray, floatArrayResult, floatArrayExpected)
            break
        }
    }

    uIntArrayResult := Map(uIntArray, doubleUInt)
    uIntArrayExpected := []uint64{8, 8, 8}
    for i := 0; i < len(uIntArrayResult); i++ {
        if uIntArrayResult[i] != uIntArrayExpected[i] {
            t.Errorf("Map of uint array %v = %v; want %v", uIntArray, uIntArrayResult, uIntArrayExpected)
            break
        }
    }

    stringArrayResult := Map(stringArray, addBug)
    stringArrayExpected := []string{"Ibug", "ambug", "abug", "bugbug"}
    for i := 0; i < len(stringArrayResult); i++ {
        if stringArrayResult[i] != stringArrayExpected[i] {
            t.Errorf("Map of string array %v = %v; want %v", stringArray, stringArrayResult, stringArrayExpected)
            break
        }
    }
}

func TestFilter(t *testing.T) {
    intArray := []int64{1, 2, -3, 4, 5}
    floatArray := []float64{1.5, 2.5, -4.0, 5.5}
    uIntArray := []uint64{4, 1, 4}
    stringArray := []string{"I", "am", "a", "bug"}

    isPositive := func(x int64) bool {
        return x > 0
    }

    isPositiveFloat := func(x float64) bool {
        return x > 0
    }

    isGreaterThanOneUInt := func(x uint64) bool {
        return x > 1
    }

    isBug := func(x string) bool {
        return x != "bug"
    }

    intArrayResult, err := Filter(intArray, isPositive)
    intArrayExpected := []int64{1, 2, 4, 5}
    for i := 0; i < len(intArrayResult); i++ {
        if intArrayResult[i] != intArrayExpected[i] || err != nil {
            t.Errorf("Filter of int array %v = %v; want %v", intArray, intArrayResult, intArrayExpected)
            break
        }
    }

    floatArrayResult, err := Filter(floatArray, isPositiveFloat)
    floatArrayExpected := []float64{1.5, 2.5, 5.5}
    for i := 0; i < len(floatArrayResult); i++ {
        if floatArrayResult[i] != floatArrayExpected[i] || err != nil {
            t.Errorf("Filter of float array %v = %v; want %v", floatArray, floatArrayResult, floatArrayExpected)
            break
        }
    }

    uIntArrayResult, err := Filter(uIntArray, isGreaterThanOneUInt)
    uIntArrayExpected := []uint64{4, 4}
    for i := 0; i < len(uIntArrayResult); i++ {
        if uIntArrayResult[i] != uIntArrayExpected[i] || err != nil {
            t.Errorf("Filter of uint array %v = %v; want %v", uIntArray, uIntArrayResult, uIntArrayExpected)
            break
        }
    }

    stringArrayResult, err := Filter(stringArray, isBug)
    stringArrayExpected := []string{"I", "am", "a"}
    for i := 0; i < len(stringArrayResult); i++ {
        if stringArrayResult[i] != stringArrayExpected[i] || err != nil {
            t.Errorf("Filter of string array %v = %v; want %v", stringArray, stringArrayResult, stringArrayExpected)
            break
        }
    }
}
