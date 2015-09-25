package golobals

import (
	"fmt"
	"testing"
)

type IdentitySource struct{}

func (i IdentitySource) Get(varName string) string {
	return varName
}

func (i IdentitySource) IsLive() bool {
	return false
}

func Test_IdentitySource(t *testing.T) {
	source := IdentitySource{}
	golobals := CreateGolobals([]Source{source}...)

	x := TestConfig{}

	initedTest := golobals.Init(x).(TestConfig)

	if initedTest.X.Get() != "x.y.z" {
		t.Error("Struct variable set incorrectly")
	}

	if initedTest.Y.Get() != "1.2.3" {
		t.Error("Struct variable set incorrectly")
	}
}
