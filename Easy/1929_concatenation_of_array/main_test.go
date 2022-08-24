package main

import (
	"reflect"
	"testing"
)

func TestCorrectGetConcatenation(t *testing.T) {
	res := getConcatenation([]int{1, 2, 1})
	if !reflect.DeepEqual(res, []int{1, 2, 1, 1, 2, 1}) {
		t.Error("Method is not correct working")
	}
}
