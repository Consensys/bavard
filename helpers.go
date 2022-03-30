// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bavard

import (
	"errors"
	"fmt"
	"math/big"
	"math/bits"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

// Template helpers (txt/template)
func helpers() template.FuncMap {
	// functions used in template
	return template.FuncMap{
		"add":        add,
		"bits":       getBits,
		"bytes":      intBytes, //TODO: Do this directly
		"capitalize": strings.Title,
		"dict":       dict,
		"div":        div,
		"divides":    divides,
		"first":      first,
		"interval":   interval,
		"iterate":    iterate,
		"last":       last,
		"list":       makeSlice,
		"log": fmt.Println,
		"mod":        mod,
		"mul":        mul,
		"mul2":       mul2,
		"noFirst":    noFirst,
		"noLast":     noLast,
		"notNil":     notNil,
		"printList":  printList,
		"reverse":    reverse,
		"sub":        sub,
		"toInt64":    toInt64,
		"toLower":    strings.ToLower,
		"toTitle":	strings.Title,
		"toUpper":    strings.ToUpper,
		"words64":    printBigIntAsUint64Slice,
	}
}

func getBits(a interface{}) ([]bool, error) {

	var res []bool
	aI, err := toInt64(a)

	if err != nil {
		return res, err
	}

	for aI != 0 {
		res = append(res, aI%2 != 0)
		aI /= 2
	}

	return res, nil
}

func toInt64(a interface{}) (int64, error) {
	switch i := a.(type) {
	case uint8:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case uint16:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case uint64:
		if i >> 63 != 0 {
			return 0, fmt.Errorf("uint64 value too large, won't fit in an int64")
		}
		return int64(i), nil
	case int64:
		return i, nil
	case int:
		return int64(i), nil
	default:
		return 0, fmt.Errorf("cannot convert to int64 from type %T", i)
	}
}

func mod(a, b interface{}) (int64, error) {

	var err error
	A, err := toInt64(a)

	if err != nil {
		return 0, err
	}

	B, err := toInt64(b)

	if err != nil {
		return 0, err
	}
	return A % B, nil
}

func intBytes(i big.Int) []byte {
	return i.Bytes()
}

func interval(begin, end interface{}) ([]int64, error) {
	beginInt, err := toInt64(begin)
	if err != nil {
		return nil, err
	}
	endInt, err := toInt64(end)
	if err != nil {
		return nil, err
	}

	l := endInt - beginInt
	r := make([]int64, l)
	for i := int64(0); i < l; i++ {
		r[i] = i + beginInt
	}
	return r, nil
}

// Adopted from https://stackoverflow.com/a/50487104/5116581
func notNil(input interface{}) bool {
	isNil := input == nil || (reflect.ValueOf(input).Kind() == reflect.Ptr && reflect.ValueOf(input).IsNil())
	return !isNil
}

func assertSlice(input interface{}) (reflect.Value, error) {
	s := reflect.ValueOf(input)
	if s.Kind() != reflect.Slice {
		return s, fmt.Errorf("value %s is not a slice", fmt.Sprint(s))
	}
	return s, nil
}

func first(input interface{}) (interface{}, error) {
	s, err := assertSlice(input)
	if err != nil {
		return nil, err
	}
	if s.Len() == 0 {
		return nil, fmt.Errorf("empty slice")
	}
	return s.Index(0).Interface(), nil
}

func last(input interface{}) (interface{}, error) {
	s, err := assertSlice(input)
	if err != nil {
		return nil, err
	}
	if s.Len() == 0 {
		return nil, fmt.Errorf("empty slice")
	}
	return s.Index(s.Len() - 1).Interface(), nil
}

var stringBuilderPool = sync.Pool{New: func() interface{} { return &strings.Builder{} }}

func printBigIntAsUint64Slice(in interface{}) (string, error) {

	var input *big.Int

	switch i := in.(type) {
	case big.Int:
		input = &i
	case *big.Int:
		input = i
	default:
		return "", fmt.Errorf("unsupported type %T", in)
	}

	words := input.Bits()

	if len(words) == 0 {
		return "0", nil
	}

	builder := stringBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	defer stringBuilderPool.Put(builder)

	for i := 0; i < len(words); i++ {
		w := uint64(words[i])

		if bits.UintSize == 32 && i < len(words)-1 {
			i++
			w = (w << 32) | uint64(words[i])
		}

		builder.WriteString(strconv.FormatUint(w, 10))

		if i < len(words)-1 {
			builder.WriteString(", ")
		}
	}

	return builder.String(), nil
}

func printList(input interface{}) (string, error) {

	s, err := assertSlice(input)

	if err != nil || s.Len() == 0 {
		return "", err
	}

	builder := stringBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	defer stringBuilderPool.Put(builder)

	builder.WriteString(fmt.Sprint(s.Index(0).Interface()))

	for i := 1; i < s.Len(); i++ {
		builder.WriteString(", ")
		builder.WriteString(fmt.Sprint(s.Index(i).Interface()))
	}

	return builder.String(), nil
}

func iterate(maxBound interface{}) (r []int64, err error) {

	var maxBoundI int64

	if maxBoundI, err = toInt64(maxBound); err != nil {
		return
	}
	for i := int64(0); i < maxBoundI; i++ {
		r = append(r, i)
	}
	return
}

func reverse(input interface{}) interface{} {

	s, err := assertSlice(input)
	if err != nil {
		return err
	}
	l := s.Len()
	toReturn := reflect.MakeSlice(s.Type(), l, l)

	l--
	for i := 0; i <= l; i++ {
		toReturn.Index(l - i).Set(s.Index(i))
	}
	return toReturn.Interface()
}

func noFirst(input interface{}) interface{} {
	s, err := assertSlice(input)
	if s.Len() == 0 {
		return input
	}
	if err != nil {
		return err
	}
	l := s.Len() - 1
	toReturn := reflect.MakeSlice(s.Type(), l, l)
	for i := 0; i < l; i++ {
		toReturn.Index(i).Set(s.Index(i + 1))
	}
	return toReturn.Interface()
}

func noLast(input interface{}) interface{} {
	s, err := assertSlice(input)
	if s.Len() == 0 {
		return input
	}
	if err != nil {
		return err
	}
	l := s.Len() - 1
	toReturn := reflect.MakeSlice(s.Type(), l, l)
	for i := 0; i < l; i++ {
		toReturn.Index(i).Set(s.Index(i))
	}
	return toReturn.Interface()
}

func add(a, b interface{}) (int64, error) {
	aI, err := toInt64(a)
	if err != nil {
		return 0, err
	}
	var bI int64
	if bI, err = toInt64(b); err != nil {
		return 0, err
	}
	return aI + bI, nil
}
func mul(a, b interface{}) (int64, error) {
	aI, err := toInt64(a)
	if err != nil {
		return 0, err
	}
	var bI int64
	if bI, err = toInt64(b); err != nil {
		return 0, err
	}
	return aI * bI, nil
}
func sub(a, b interface{}) (int64, error) {
	aI, err := toInt64(a)
	if err != nil {
		return 0, err
	}
	var bI int64
	if bI, err = toInt64(b); err != nil {
		return 0, err
	}
	return aI - bI, nil
}
func mul2(a interface{}) (int64, error) {
	aI, err := toInt64(a)
	if err != nil {
		return 0, err
	}

	return aI * 2, nil
}
func div(a, b interface{}) (int64, error) {
	aI, err := toInt64(a)
	if err != nil {
		return 0, err
	}
	var bI int64
	if bI, err = toInt64(b); err != nil {
		return 0, err
	}
	return aI / bI, nil
}

func makeSlice(values ...interface{}) []interface{} {
	return values
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// return true if c1 divides c2, that is, c2 % c1 == 0
func divides(c1, c2 interface{}) (bool, error) {
	/*switch cc1 := c1.(type) {
	case int:
		switch cc2 := c2.(type) {
		case int:
			return cc2%cc1 == 0, nil
		case string:
			c2Int, err := strconv.Atoi(cc2)
			if err != nil {
				return false, err
			}
			return c2Int%cc1 == 0, nil
		}
	case string:
		c1Int, err := strconv.Atoi(cc1)
		if err != nil {
			panic(err)
		}
		switch cc2 := c2.(type) {
		case int:
			return cc2%c1Int == 0, nil
		case string:
			c2Int, err := strconv.Atoi(cc2)
			if err != nil {
				return false, err
			}
			return c2Int%c1Int == 0, nil
		}
	}*/

	//try to convert to int64
	c1Int, err := toInt64(c1)
	if err != nil {
		return false, err
	}
	var c2Int int64
	c2Int, err = toInt64(c2)
	if err != nil {
		return false, err
	}

	return c2Int%c1Int == 0, nil
}
