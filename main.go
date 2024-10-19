package go_digest

import (
	"math"
	"math/cmplx"
	"math/rand/v2"
	"unsafe"
)

// GetCharByIndex returns the i-th character from the given string.
func GetCharByIndex(str string, idx int) rune {
	pos := 0
	for _, runeValue := range str {
		if pos == idx {
			return runeValue
		}
		pos++
	}
	panic("none")
}

// GetStringBySliceOfIndexes returns a string formed by concatenating specific characters from the input string based
// on the provided indexes.
func GetStringBySliceOfIndexes(str string, indexes []int) string {
	result := make([]rune, len(indexes))
	for i := range indexes {
		result[i] = GetCharByIndex(str, indexes[i])
	}

	return string(result)
}

// ShiftPointer shifts the given pointer by the specified number of bytes using unsafe.Add.
func ShiftPointer(pointer **int, shift int) {
	*pointer = (*int)(unsafe.Add(unsafe.Pointer(*pointer), shift))
}

const float64EqualityThreshold = 1e-4

func almostEqual(a, b float64) bool {
	return a == b || math.Abs(a-b) <= float64EqualityThreshold
}

// IsComplexEqual compares two complex numbers and determines if they are equal.
func IsComplexEqual(a, b complex128) bool {
	return almostEqual(real(a), real(b)) && almostEqual(imag(a), imag(b))
}

// GetRootsOfQuadraticEquation returns two roots of a quadratic equation ax^2 + bx + c = 0.
func GetRootsOfQuadraticEquation(a, b, c float64) (complex128, complex128) {
	d := complex(math.Pow(b, 2)-4*a*c, 0)
	dRoot := cmplx.Sqrt(d)

	bComplex := complex(b, 0)
	aComplex := complex(a, 0)

	x1 := (-bComplex + dRoot) / (2 * aComplex)
	x2 := (-bComplex - dRoot) / (2 * aComplex)

	return x1, x2
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// Sort sorts in-place the given slice of integers in ascending order.
func Sort(source []int) {
	rand.Shuffle(len(source), func(i, j int) {
		source[i], source[j] = source[j], source[i]
	})
	quickSort(source, 0, len(source)-1)
}

// ReverseSliceOne in-place reverses the order of elements in the given slice.
func ReverseSliceOne(s []int) {
	for i := 0; i < len(s)/2; i++ {
		pair := len(s) - i - 1
		s[i], s[pair] = s[pair], s[i]
	}
}

// ReverseSliceTwo returns a new slice of integers with elements in reverse order compared to the input slice.
// The original slice remains unmodified.
func ReverseSliceTwo(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)

	ReverseSliceOne(result)

	return result
}

// SwapPointers swaps the values of two pointers.
func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

// IsSliceEqual compares two slices of integers and returns true if they contain the same elements in the same order.
func IsSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// DeleteByIndex deletes the element at the specified index from the slice and returns a new slice.
// The original slice remains unmodified.
func DeleteByIndex(s []int, idx int) []int {
	if idx >= len(s) || idx < 0 || len(s) <= 0 {
		panic("lol")
	}
	result := make([]int, len(s)-1)
	for i, v := range s {
		if i > idx {
			result[i-1] = v
		} else if i < idx {
			result[i] = v
		}
	}

	return result
}
