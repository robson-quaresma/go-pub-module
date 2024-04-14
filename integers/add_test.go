package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	num1, num2 := 2, 1
	expected := 3

	// Act
	result := Add(num1, num2)

	// Assert
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", num1, num2, result, expected)
	}
}

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}
