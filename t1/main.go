package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func printType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func toStringSlice(vars ...interface{}) []string {
	var result []string
	for _, v := range vars {
		switch val := v.(type) {
		case int:
			result = append(result, strconv.Itoa(val))
		case float64:
			result = append(result, strconv.FormatFloat(val, 'f', -1, 64))
		case string:
			result = append(result, val)
		case bool:
			result = append(result, strconv.FormatBool(val))
		case complex64:
			result = append(result, fmt.Sprintf("%g", val))
		default:
			result = append(result, fmt.Sprintf("%v", val))
		}
	}
	return result
}

func concatStrings(strs []string) string {
	return fmt.Sprint(strs)
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func insertSalt(runes []rune, salt string) []rune {
	mid := len(runes) / 2
	saltRunes := []rune(salt)
	return append(runes[:mid], append(saltRunes, runes[mid:]...)...)
}

func hashRunes(runes []rune) string {
	b := []byte(string(runes))
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	// 1
	fmt.Println("Types:")
	fmt.Println("numDecimal:", printType(numDecimal))
	fmt.Println("numOctal:", printType(numOctal))
	fmt.Println("numHexadecimal:", printType(numHexadecimal))
	fmt.Println("pi:", printType(pi))
	fmt.Println("name:", printType(name))
	fmt.Println("isActive:", printType(isActive))
	fmt.Println("complexNum:", printType(complexNum))

	// 2
	strs := toStringSlice(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	concat := concatStrings(strs)
	fmt.Println("\nConcatenated string:", concat)

	// 3
	runes := stringToRunes(concat)
	fmt.Println("\nRune slice:", runes)

	// 4
	runesSalted := insertSalt(runes, "go-2024")
	fmt.Println("\nRune slice with salt:", runesSalted)

	// 5
	hash := hashRunes(runesSalted)
	fmt.Println("\nSHA256 hash:", hash)
}
