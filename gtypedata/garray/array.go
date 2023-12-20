package garray

import (
	"fmt"
)

// Contains checks if an element is present in the slice.
// It returns true if the element is found, otherwise false.
func Contains[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// AppendUniqueVal appends unique elements to the destination slice.
// It checks for duplicates using the Contains function.
func AppendUniqueVal[T comparable](dst []T, elem ...T) []T {
	for _, e := range elem {
		if !Contains(dst, e) {
			dst = append(dst, e)
		}
	}

	return dst
}

// AppendUniqueValWithErr appends unique elements to the destination slice.
// It returns an error if any of the elements already exist in the slice.
func AppendUniqueValWithErr[T comparable](dst []T, elem ...T) ([]T, error) {
	var newDst []T
	for _, e := range elem {
		if !Contains(dst, e) {
			newDst = append(newDst, e)
		} else {
			return dst, fmt.Errorf("failed append array because elements is already exist. value: %v", e)
		}
	}

	dst = append(dst, newDst...)
	return dst, nil
}

func FilterDifferentElem[T comparable](src []T, ref []T) []T {
	var res []T
	for _, v := range src {
		if Contains(ref, v) {
			res = append(res, v)
		}
	}

	return res
}

func SlicesMatch[T comparable](src []T, ref []T) bool {
	if len(src) != len(ref) {
		return false
	}

	for i, v := range src {
		if v != ref[i] {
			return false
		}
	}

	return true
}
