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
		"capitalize": strings.Title,
		"dict":       dict,
		"div":        div,
		"divides":    divides,
		"iterate":    iterate,
		"last":       last,
		"list":       printList,
		"mul":        mul,
		"mul2":       mul2,
		"notNil":     notNil,
		"reverse":    reverse,
		"sub":        sub,
		"toLower":    strings.ToLower,
		"toUpper":    strings.ToUpper,
	}
}

func notNil(input *interface{}) bool {
	return input != nil
}

func last(input []interface{}) interface{} {
	return input[len(input)-1]
}

var stringBuilderPool = sync.Pool{New: func() interface{} { return strings.Builder{} }}

func printList(input []interface{}) string {

	if len(input) == 0 {
		return ""
	}

	builder := stringBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	defer stringBuilderPool.Put(builder)

	builder.WriteString(fmt.Sprint(input[0]))

	for i := 1; i < len(input); i++ {
		builder.WriteString(", ")
		builder.WriteString(fmt.Sprint(input[i]))
	}

	return builder.String()
}

func iterate(maxBound int) (r []int) {
	for i := 0; i < maxBound; i++ {
		r = append(r, i)
	}
	return
}

func reverse(input []int) []int {
	toReturn := make([]int, len(input))
	j := 0
	for i := len(input) - 1; i >= 0; i-- {
		toReturn[j] = input[i]
		j++
	}
	return toReturn
}
func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func sub(a, b int) int {
	return a - b
}
func mul2(a int) int {
	return a * 2
}
func div(a, b int) int {
	return a / b
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
func divides(c1, c2 interface{}) bool {
	switch cc1 := c1.(type) {
	case int:
		switch cc2 := c2.(type) {
		case int:
			return cc2%cc1 == 0
		case string:
			c2Int, err := strconv.Atoi(cc2)
			if err != nil {
				panic(err)
			}
			return c2Int%cc1 == 0
		}
	case string:
		c1Int, err := strconv.Atoi(cc1)
		if err != nil {
			panic(err)
		}
		switch cc2 := c2.(type) {
		case int:
			return cc2%c1Int == 0
		case string:
			c2Int, err := strconv.Atoi(cc2)
			if err != nil {
				panic(err)
			}
			return c2Int%c1Int == 0
		}
	}
	panic("unexpected type")
}
