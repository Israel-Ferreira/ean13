package ean13

import (
	"fmt"
	"testing"
)


func assertResult(t *testing.T, expected, result string) {
	t.Helper()

	if expected != result {
		t.Errorf("Result: %s, expected: %s", result, expected)
	}
}

func TestGetEan13VerificationDigit(t *testing.T) {
	t.Run("Should works on Valid code", func(t *testing.T) {
		expected := "789901476726-8"
		result,err := GetEan13VerficationDigit("789901476726")

		if err != nil {
			t.Errorf("Error %e", err)
		}
	
		assertResult(t, result, expected)
	})



}


func ExampleGetEan13VerificationDigit(){
	result, _ := GetEan13VerficationDigit("789901476726")
	fmt.Println(result)
	// Output: 789901476726-8
}