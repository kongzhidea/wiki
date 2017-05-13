// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package listdict

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Simple list
type List []interface{}

// Return new List with specified length
func NewList(length int) List {
	return make(List, length)
}

var (
	// ErrRemoveFromEmptyList is returned when user want to remove element
	// from empty list
	ErrRemoveFromEmptyList = errors.
		New("Trying to remove element from empty list")
)

//=============================================================================

// Append adds an element to the end of the list.
func (list *List) Append(values ...interface{}) {
	*list = append(*list, values...)
}

// AppendIfMissing adds an element to the end of the list if it's not already
// in the list.
func (list *List) AppendIfMissing(value interface{}) {
	for _, ele := range *list {
		if ele == value {
			// Element exists, exit
			goto exit
		}
	}
	*list = append(*list, value)
exit:
}

// Count returns the number of times value appears in the list.
func (list List) Count(value interface{}) int {
	counter := 0
	for _, listValue := range list {
		if listValue == value {
			counter++
		}
	}
	return counter
}

// Delete removes element with given index from the list.
func (list *List) Delete(index int) error {
	if len(*list) <= 0 {
		return ErrRemoveFromEmptyList
	}

	listLen := len(*list)

	copy((*list)[index:], (*list)[index+1:])
	(*list)[listLen-1] = nil
	*list = (*list)[:listLen-1]
	return nil
}

// Extend one list with the contents of the other list.
func (list *List) Extend(otherList List) {
	for _, value := range otherList {
		*list = append(*list, value)
	}
}

// Index returns the index of the first item in the list whose value is val.
// It is -1 if there is no such item.
func (list List) Index(val interface{}) (int, error) {
	for index, listValue := range list {
		if listValue == val {
			return index, nil
		}
	}
	errorString := fmt.Sprintf("%v is not in list", val)
	return -1, errors.New(errorString)
}

// Insert an element at a given position. If the position is past the end
// of the list, append to the end.
func (list *List) Insert(index int, values ...interface{}) {

	if len(*list) > index {
		for i := 0; i < len(values); i++ {
			*list = append(*list, 0)
		}

		copy((*list)[index+len(values):], (*list)[index:])
		copy((*list)[index:], values)
	} else {
		*list = append(*list, values...)
	}
}

// IsEqual returns true if lists are equal.
func (list List) IsEqual(otherList List) bool {
	return reflect.DeepEqual(list, otherList)
}

// Remove and returns the last element in the list.
func (list *List) Pop() (interface{}, error) {
	if len(*list) <= 0 {
		return nil, ErrRemoveFromEmptyList
	}

	listLen := len(*list)
	val := (*list)[listLen-1]
	(*list).Delete(listLen - 1)

	return val, nil
}

// Remove and returns the element at the given position in the list.
func (list *List) PopItem(index int) (interface{}, error) {
	if len(*list) <= 0 {
		return nil, ErrRemoveFromEmptyList
	}

	val := (*list)[index]
	(*list).Delete(index)

	return val, nil
}

// Remove the first element from the list whose value matches the given value.
// Error if no match is found.
func (list *List) Remove(val interface{}) error {
	errorString := fmt.Sprintf("%v is not in list", val)
	if len(*list) > 0 {
		for index, listValue := range *list {
			if listValue == val {
				(*list).Delete(index)
				return nil
			}
		}
	}
	return errors.New(errorString)
}

// Reverse the elements of the list in place.
func (list *List) Reverse() {
	if len(*list) > 0 {
		maxIndex := len(*list) - 1
		for index := 0; index < (maxIndex/2)+1; index++ {
			(*list)[index], (*list)[maxIndex-index] =
				(*list)[maxIndex-index], (*list)[index]
		}
	}
}

// Sort the list in place ordering elements from smallest to largest.
//func (list *List) Sort() {
//}

// String returns list values as string
//		l := listdict.List{"one", 2, "three"}
// 		l.String() => "one, 2, three"
func (list List) String() string {
	var out []string
	for _, val := range list {
		out = append(out, fmt.Sprintf("%v", val))
	}
	return strings.Join(out, ", ")
}