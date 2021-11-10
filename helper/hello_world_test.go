package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMain
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

// unit test menggunakan assertions library testify
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Udin")
	assert.Equal(t, "Hello Udin", result, "Result must be 'Hello Udin'")
	fmt.Println("TestHelloWorld with assert Done")
}

// unit test menggunakan require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Udin")
	require.Equal(t, "Hello Udin", result, "Result must be 'Hello Udin'")
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}

	result := HelloWorld("Danil")
	require.Equal(t, "Hello Danil", result, "Result must be 'Hello Danil'")
}

func TestCheckOsVersion(t *testing.T) {
	result := CheckOsVersion()
	require.Equal(t, "linux", result, "Result must be 'windows'")
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
