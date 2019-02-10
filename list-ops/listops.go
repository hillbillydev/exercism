// Package listops gives you a list structure with some methods.
package listops

// IntList list that contains ints.
type IntList []int

type predFunc func(n int) bool
type unaryFunc func(x int) int
type binFunc func(x, y int) int

// Foldr folds the list from right to left.
func (list IntList) Foldr(fn binFunc, init int) int {
	var result = init
	for _, n := range list.Reverse() {
		result = fn(n, result)
	}
	return result
}

// Foldl folds the list from left to right.
func (list IntList) Foldl(f binFunc, init int) int {
	var result = init
	for _, n := range list {
		result = f(result, n)
	}
	return result
}

// Length gives back the length.
func (list IntList) Length() int {
	return len(list)
}

// Reverse reverses the int list.
func (list IntList) Reverse() IntList {
	var result IntList
	if len(list) == 0 {
		return list
	}
	for i := len(list); i != 0; i-- {
		result = append(result, list[i-1])
	}
	return result
}

// Map takes the int list and execute the passed in function A -> B
func (list IntList) Map(u unaryFunc) IntList {
	var result IntList
	if len(list) == 0 {
		return list
	}
	for _, n := range list {
		result = append(result, u(n))
	}
	return result
}

// Append takes another list and appends it to the tail of the target.
func (list IntList) Append(s IntList) IntList {
	return append(list, s...)
}

// Concat takes multiple IntLists and appends it to the tail of the target.
func (list IntList) Concat(s []IntList) IntList {
	if len(list) == 0 {
		return list
	}
	for _, numbers := range s {

		list = append(list, numbers...)
	}
	return list
}

// Filter takes a predicate and make sure all the items fullfill that predicate.
func (list IntList) Filter(predicate predFunc) IntList {
	var result IntList
	if len(list) == 0 {
		return list
	}
	for _, n := range list {
		if !predicate(n) {
			continue
		}
		result = append(result, n)
	}
	return result
}
