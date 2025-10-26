package main

import (
	"reflect"
	"testing"
)

func TestPrintType(t *testing.T) {
	if got := printType(42); got != "int" {
		t.Errorf("printType(42) = %v, want int", got)
	}
	if got := printType(3.14); got != "float64" {
		t.Errorf("printType(3.14) = %v, want float64", got)
	}
	if got := printType("Golang"); got != "string" {
		t.Errorf("printType(\"Golang\") = %v, want string", got)
	}
	if got := printType(true); got != "bool" {
		t.Errorf("printType(true) = %v, want bool", got)
	}
	if got := printType(complex64(1 + 2i)); got != "complex64" {
		t.Errorf("printType(complex64(1+2i)) = %v, want complex64", got)
	}
}

func TestToStringSlice(t *testing.T) {
	got := toStringSlice(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	want := []string{"42", "42", "42", "3.14", "Golang", "true", "(1+2i)"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("toStringSlice(...) = %v, want %v", got, want)
	}
}

func TestConcatStrings(t *testing.T) {
	got := concatStrings([]string{"42", "42", "42", "3.14", "Golang", "true", "1+2i"})
	want := "[42 42 42 3.14 Golang true 1+2i]"
	if got != want {
		t.Errorf("concatStrings(...) = %v, want %v", got, want)
	}
}

func TestStringToRunes(t *testing.T) {
	got := stringToRunes("Go!")
	want := []rune{'G', 'o', '!'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("stringToRunes(\"Go!\") = %v, want %v", got, want)
	}
}

func TestInsertSalt(t *testing.T) {
	runes := []rune("abcdef")
	salt := "go-2024"
	got := insertSalt(runes, salt)
	want := []rune("abcgo-2024def")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("insertSalt(...) = %v, want %v", string(got), string(want))
	}
}

func TestHashRunes(t *testing.T) {
	runes := []rune("testgo-2024string")
	got := hashRunes(runes)
	if len(got) != 64 {
		t.Errorf("hashRunes(...) length = %d, want 64", len(got))
	}
}
