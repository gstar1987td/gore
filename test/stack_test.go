package test

import (
	"gore/pkg/models"
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	stk := models.NewArrayStack()
	if !stk.Empty() {
		t.Fatal("not empty")
	}
	stk.Push(1)
	p, ok := stk.Pop()
	if !ok {
		t.Fatal("no value")
	}
	if !reflect.DeepEqual(p, 1) {
		t.Fatal("not equal")
	}
}
