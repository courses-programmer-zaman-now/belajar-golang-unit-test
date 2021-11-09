package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldFika(t *testing.T) {
	result := HelloWorld("fika")
	if result != "Hello fika" {
		// error
		// panic("Result is not 'Hello danil'")
		// t.Fail()
		t.Error("Result must be 'Hello fika'")
	}
	fmt.Println("TestHelloWorldFika done")

}

func TestHelloWorldDanil(t *testing.T) {
	result := HelloWorld("danil")
	if result != "Hello danil" {
		// error
		// panic("Result is not 'Hello danil'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello danil'")
	}
	fmt.Println("TestHelloWorldDanil Done")
}

func TestKelilingSegiTiga(t *testing.T) {
	result := KelilingSegiTiga(20)
	if result != 60 {
		// error
		// panic("Result is not '60'")
		t.Fatal("Result Must Be = 60")
	}
}
