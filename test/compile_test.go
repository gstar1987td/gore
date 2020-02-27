package test

import (
	"context"
	"fmt"
	gore "gore/pkg/rule"
	"testing"
)

var examples = "../examples/"

func TestCompileFile(t *testing.T) {
	var file = examples + "demo.dsl"
	r, err := gore.NewGoreRuleFromFile(context.Background(), file)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(r.ID, ".", r.Version)
}
