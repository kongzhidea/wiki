// Copyright 2012 Dobrosław Żybort
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package listdict

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Simple dict.
type Dict map[string]interface{}

// Return new Dict.
func NewDict() Dict {
	return make(Dict)
}

var (
	// ErrRemoveFromEmptyDict is returned when user want to remove element
	// from empty dictionary
	ErrRemoveFromEmptyDict = errors.
		New("Trying to remove element from empty dict")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//=============================================================================

// DictFromKeys creates a new dictionary with keys from list and values set
// to defaultVal.
func DictFromKeys(list List, defaultVal interface{}) Dict {
	newDict := NewDict()
	for _, value := range list {
		newDict[fmt.Sprintf("%v", value)] = defaultVal
	}
	return newDict
}

//=============================================================================

// Clear removes all elements from the dictionary.
func (dict Dict) Clear() {
	for key, _ := range dict {
		delete(dict, key)
	}
}

// Get returns value for the given key or defaultVal if key is NOT in
// the dictionary. defaultVal should be same type as you expect to get.
//		d := listdict.Dict{"one": 1, "two": 2}
// 		d["one"]          => 1
// 		d.Get("one", 4)   => 1
//		d["three"]        => error
// 		d.Get("three", 3) => 3
// 		// d = {'one': 1, 'two': 2}
func (dict Dict) Get(key string, defaultVal interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	return defaultVal
}

// HasKey returns true if key is in the dictionary, false otherwise.
func (dict Dict) HasKey(key string) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}

// IsEqual returns true if dicts are equal.
func (dict Dict) IsEqual(otherDict Dict) bool {
	return reflect.DeepEqual(dict, otherDict)
}

// Items returns an unordered list of the dictionary's [key, value] pairs.
func (dict Dict) Items() []List {
	list := []List{}
	for key, value := range dict {
		list = append(list, List{key, value})
	}
	return list
}

// Keys returns a list of the dictionary's keys, unordered.
func (dict Dict) Keys() List {
	list := NewList(len(dict))
	i := 0
	for key, _ := range dict {
		list[i] = key
		i++
	}
	return list
}

// Pop returns value and remove the given key from the dictionary.
// If the given key is NOT in the dictionary return defaultVal.
// defaultVal should be same type as you expect to get.
func (dict Dict) Pop(key string, defaultVal interface{}) (interface{}, error) {
	if len(dict) <= 0 {
		return defaultVal, ErrRemoveFromEmptyDict
	}
	if dict.HasKey(key) {
		val := dict[key]
		delete(dict, key)
		return val, nil
	}
	return defaultVal, nil
}

// PopItem return and remove a random key-value pair as List from
// the dictionary.
func (dict Dict) PopItem() (List, error) {
	if len(dict) <= 0 {
		return List{}, ErrRemoveFromEmptyDict
	}

	// Get dict keys
	dictKeys := dict.Keys()
	// Return random key as string
	randKey := fmt.Sprintf("%v", dictKeys[rand.Intn(len(dictKeys))])

	list := NewList(2)
	list = List{randKey, dict[randKey]}

	delete(dict, randKey)

	return list, nil

}

// SetDefault is like Get but will set dict[key] to defaultVal if key is not
// already in dict. defaultVal should be same type as you expect to get.
// 		d := listdict.Dict{"one": 1, "two": 2}
// 		d["one"]                 => 1
// 		d.SetDefault("one", 4)   => 1
//		d["three"]               => error
// 		d.SetDefault("three", 3) => 3
// 		// d = {'one': 1, 'two': 2, 'three': 3}
func (dict Dict) PutIfAbsent(key string, value interface{}) interface{} {
	if dict.HasKey(key) {
		return dict[key]
	}
	dict[key] = value
	return value
}

func (dict Dict) Put(key string, value interface{}) interface{} {

	dict[key] = value
	return value
}


// Update updates the dictionary with the key-value pairs in the dict2
// dictionary replacing current values and adding new if found.
func (dict Dict) Update(dict2 Dict) {
	for key, value := range dict2 {
		dict[key] = value
	}
}

// Values returns a list of the dictionary's values, unordered.
func (dict Dict) Values() List {
	list := NewList(len(dict))
	i := 0
	for _, value := range dict {
		list[i] = value
		i++
	}
	return list
}