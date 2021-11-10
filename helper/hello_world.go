package helper

import "runtime"

func HelloWorld(name string) string {
	return "Hello " + name
}

func KelilingSegiTiga(sisi int) int {
	keliling := sisi + sisi + sisi
	return keliling
}

func CheckOsVersion() string {
	version := runtime.GOOS
	return version
}

func Penjumlahan(bil1, bil2 int) int {
	result := bil1 + bil2
	return result
}

func Pengurangan(bil1, bil2 int) int {
	result := bil1 - bil2
	return result
}
