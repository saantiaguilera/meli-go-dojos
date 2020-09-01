package pkg_test

import (
	"github.com/mercadolibre/fury_shipping-dx-dojo/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCalculator_GivenTwoInts_WhenSumming_ThenCorrectSumIsReturned(t *testing.T) {
	// Given
	a := 1
	b := 2
	expectedResult := 3
	calculator := pkg.NewBasicCalculator()

	// When
	result := calculator.Sum(a, b)

	// Then
	assert.Equal(t, expectedResult, result)
}

func TestResultAwareCalculator_GivenTwoSetOfInts_WhenSummingTwice_ThenSecondSumIsReturnedWithPrevious(t *testing.T) {
	// Given
	firstA := 1
	firstB := 2
	secondA := 2
	secondB := 4
	expectedResult := 9
	calculator := pkg.NewResultAwareCalculator()

	// When
	calculator.Sum(firstA, firstB)
	result := calculator.Sum(secondA, secondB)

	// Then
	assert.Equal(t, expectedResult, result)
}
