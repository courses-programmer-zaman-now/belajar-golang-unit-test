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

// sub benchmark
func BenchmarkSubPerhitungan(b *testing.B) {
	b.Run("tambah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Penjumlahan(2, 6)
		}
	})

	b.Run("kurang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Pengurangan(3, 1)
		}
	})
}

// membuat Benchmark Tabel
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "danil",
			request: "danil syah arihardjo",
		},
		{
			name:    "haykal",
			request: "haykal dafiansyah",
		},
		{
			name:    "fika",
			request: "nufika",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkTablePerhitungan(b *testing.B) {
	values := []struct {
		name string
		bil1 int
		bil2 int
	}{
		{
			name: "hitungan1",
			bil1: 5,
			bil2: 3,
		},
		{
			name: "hitungan2",
			bil1: 4,
			bil2: 1,
		},
		{
			name: "hitungan3",
			bil1: 6,
			bil2: 7,
		},
	}

	for _, benchmark := range values {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Penjumlahan(benchmark.bil1, benchmark.bil2)
			}
		})
	}
	for _, benchmark := range values {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Pengurangan(benchmark.bil1, benchmark.bil2)
			}
		})
	}

}
