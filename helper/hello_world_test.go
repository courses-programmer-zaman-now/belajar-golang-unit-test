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

// Testing Sub Test
func TestPenghitungan(t *testing.T) {
	t.Run("tambah", func(t *testing.T) {
		result := Penjumlahan(5, 5)
		require.Equal(t, 10, result, "Result must be '10'")
	})
	t.Run("kurang", func(t *testing.T) {
		result := Pengurangan(5, 2)
		require.Equal(t, 3, result, "Result must be '3'")
	})
}

// Membuat Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "test1",
			request:  "danil",
			expected: "Hello danil",
		},
		{
			name:     "test2",
			request:  "haykal",
			expected: "Hello haykal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

// belajar membuat benchmark untuk mengetest kecepatan proses function
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("danil")
	}
}

func BenchmarkPenghitunganJumlah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Penjumlahan(2, 5)
	}
}
