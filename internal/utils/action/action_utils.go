package utils

import (
	t_ "coolsim/internal/types"
	"reflect"
)

// Equals interface issue
// @see https://github.com/golang/go/issues/30557
// @see https://go.dev/doc/faq#t_and_equal_interface
func Equals(action1, action2 t_.Action) bool {
	return reflect.TypeOf(action1) == reflect.TypeOf(action2)
}

func Intersection(allActions, capabilities []t_.Action) []t_.Action {
	dict := make(map[reflect.Type]t_.Action)
	for _, action := range capabilities {
		key := reflect.TypeOf(action)
		dict[key] = action
	}

	var sequence []t_.Action
	for _, action := range allActions {
		key := reflect.TypeOf(action)
		if _, exists := dict[key]; exists {
			sequence = append(sequence, action)
		}

	}

	return sequence
}
