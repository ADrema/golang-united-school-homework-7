package coverage

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	stringErrorFormat  = "Expected: %s, got %s"
	intErrorFormat     = "Expected: %d, got %d"
	complexErrorFormat = "Expected: %v, got %v"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// TEST DATA
func prepareTestData() People {
	var peopleCollection People
	var timeStamp = time.UTC
	peopleCollection = append(peopleCollection, Person{
		"Vasia",
		"Vasiliev",
		time.Date(2000, time.April, 1, 10, 10, 10, 0, timeStamp),
	})

	peopleCollection = append(peopleCollection, Person{
		"Ivan",
		"Ivanov",
		time.Date(2001, time.April, 11, 11, 11, 11, 0, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Valia",
		"Valieva",
		time.Date(2001, time.March, 11, 11, 11, 11, 0, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Marina",
		"Marinova",
		time.Date(2007, time.January, 5, 5, 5, 5, 0, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Semen",
		"Semenov",
		time.Date(2007, time.January, 5, 5, 5, 5, 0, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Semen",
		"Filipov",
		time.Date(2007, time.January, 5, 5, 5, 5, 0, timeStamp),
	})

	return peopleCollection
}

// TESTS

// People Len() tests
func TestLenFunctionIsNotEmpty(t *testing.T) {
	var peoples = prepareTestData()
	var expected = 6
	var actual = peoples.Len()
	if expected != actual {
		t.Errorf(intErrorFormat, expected, actual)
	}
}

func TestLenFunctionIsEmpty(t *testing.T) {
	var peoples People
	var expected = 0
	var actual = peoples.Len()
	if expected != actual {
		t.Errorf(intErrorFormat, expected, actual)
	}
}

// People less tests
func TestLessFunctionSortsIsRequired(t *testing.T) {
	var peoples = prepareTestData()
	var actual = peoples.Less(1, 2)
	if !actual {
		t.Errorf(complexErrorFormat, true, actual)
	}
}

func TestLessFunctionSortsIsNotRequired(t *testing.T) {
	var peoples = prepareTestData()
	var actual = peoples.Less(0, 3)
	if actual {
		t.Errorf(complexErrorFormat, false, actual)
	}
}

func TestLessFunctionEqualDatesDiffFirstNames(t *testing.T) {
	var peoples = prepareTestData()
	var actual = peoples.Less(3, 4)
	if !actual {
		t.Errorf(complexErrorFormat, false, actual)
	}
}

func TestLessFunctionEqualDatesSameFirstNames(t *testing.T) {
	var peoples = prepareTestData()
	var actual = peoples.Less(4, 5)
	if actual {
		t.Errorf(complexErrorFormat, false, actual)
	}
}

// People swap tests
func TestSwapFunctionDifferentElements(t *testing.T) {
	var peoples = prepareTestData()
	var originalCollection = prepareTestData()
	var firstIndex = 0
	var secondIndex = 3

	peoples.Swap(firstIndex, secondIndex)

	var expected = originalCollection[secondIndex].firstName
	var actual = peoples[firstIndex].firstName
	if actual != expected {
		t.Errorf(stringErrorFormat, expected, actual)
	}
	expected = originalCollection[firstIndex].firstName
	actual = peoples[secondIndex].firstName
	if actual != expected {
		t.Errorf(stringErrorFormat, expected, actual)
	}
}

func TestSwapFunctionSameElement(t *testing.T) {
	var peoples = prepareTestData()
	var originalCollection = prepareTestData()
	peoples.Swap(3, 3)

	var expected = originalCollection[3].firstName
	var actual = peoples[3].firstName
	if actual != expected {
		t.Errorf(stringErrorFormat, expected, actual)
	}
}

func TestSwapFunctionLastWithFirst(t *testing.T) {
	var peoples = prepareTestData()
	var originalCollection = prepareTestData()
	var firstIndex = 5
	var secondIndex = 0

	peoples.Swap(firstIndex, secondIndex)

	var expected = originalCollection[secondIndex].firstName
	var actual = peoples[firstIndex].firstName
	if actual != expected {
		t.Errorf(stringErrorFormat, expected, actual)
	}
	expected = originalCollection[firstIndex].firstName
	actual = peoples[secondIndex].firstName
	if actual != expected {
		t.Errorf(stringErrorFormat, expected, actual)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Matrix tests
func TestEmptyMatrixCreation(t *testing.T) {
	_, err := New("")
	var expectedError = "strconv.Atoi: parsing \"\": invalid syntax"
	if err.Error() != expectedError {
		t.Errorf(stringErrorFormat, expectedError, err)
	}
}

func TestSingleSpaceMatrixCreation(t *testing.T) {
	_, err := New(" ")
	var expectedError = "strconv.Atoi: parsing \"\": invalid syntax"
	if err.Error() != expectedError {
		t.Errorf(stringErrorFormat, expectedError, err)
	}
}

func TestSpaceLineMatrixCreation(t *testing.T) {
	_, err := New("1\n ")
	var expectedError = "strconv.Atoi: parsing \"\": invalid syntax"
	if err.Error() != expectedError {
		t.Errorf(stringErrorFormat, expectedError, err)
	}
}

func TestMultiLineMatrixCreation(t *testing.T) {
	matrix, err := New("1 1 1\n1 1 1\n1 1 1")
	var expectedMatrix = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	if err != nil {
		t.Errorf(complexErrorFormat, nil, err)
	}
	if !reflect.DeepEqual(matrix.data, expectedMatrix) {
		t.Errorf(complexErrorFormat, expectedMatrix, matrix.data)
	}
}

func TestMultiLineMatrixWithDiffLineLengthCreation(t *testing.T) {
	_, err := New("1 1 1\n1 1\n1 1 1")
	var expectedErrorMessage = "Rows need to be the same length"
	if err.Error() != expectedErrorMessage {
		t.Errorf(stringErrorFormat, expectedErrorMessage, err.Error())
	}
}

func TestMultiLineMatrixWithCharSymbolsLengthCreation(t *testing.T) {
	_, err := New("1 1 x\n1 1 y\n1 1 z")
	var expectedErrorMessage = "strconv.Atoi: parsing \"x\": invalid syntax"
	if err.Error() != expectedErrorMessage {
		t.Errorf(stringErrorFormat, expectedErrorMessage, err.Error())
	}
}

// Matrix rows tests
func TestMultiLevelsMatrixRowsFunction(t *testing.T) {
	matrix, err := New("1 2 3\n4 5 6\n7 8 9")
	var expectedRows = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	if !reflect.DeepEqual(expectedRows, matrix.Rows()) {
		t.Errorf(complexErrorFormat, expectedRows, matrix.Rows())
	}
	if err != nil {
		t.Errorf(complexErrorFormat, "No errors expected", err)
	}
}

func TestSingleMatrixRowsFunction(t *testing.T) {
	matrix, err := New("1")
	var expectedRows = [][]int{
		{1},
	}
	if !reflect.DeepEqual(expectedRows, matrix.Rows()) {
		t.Errorf(complexErrorFormat, expectedRows, matrix.Rows())
	}
	if err != nil {
		t.Errorf(complexErrorFormat, "No errors expected", err)
	}
}

// Matrix cols tests
func TestMultiLevelsMatrixColFunction(t *testing.T) {
	matrix, err := New("1 2 3\n4 5 6\n7 8 9")
	var expectedCols = [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	if !reflect.DeepEqual(expectedCols, matrix.Cols()) {
		t.Errorf(complexErrorFormat, expectedCols, matrix.Cols())
	}
	if err != nil {
		t.Errorf(complexErrorFormat, "No errors expected", err)
	}
}

func TestSingleValueMatrixColFunction(t *testing.T) {
	matrix, err := New("8")
	var expectedCols = [][]int{
		{8},
	}
	if !reflect.DeepEqual(expectedCols, matrix.Cols()) {
		t.Errorf(complexErrorFormat, expectedCols, matrix.Cols())
	}
	if err != nil {
		t.Errorf(complexErrorFormat, "No errors expected", err)
	}
}

// Matrix set tests
func TestNegativeRowInSetFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(-1, 1, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestRowIndexMoreThenMatrixRowLengthFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(4, 1, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestRowIndexEqualToMatrixRowLengthFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(3, 1, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestNegativeColInSetFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(1, -1, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestRowIndexMoreThenMatrixColLengthFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(1, 4, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestColIndexEqualToMatrixRowLengthFunction(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var isSet = matrix.Set(1, 3, 1)

	if isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
}

func TestSetFunctionColAndRowHaveCorrectValues(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var expectedValue = 100
	var isSet = matrix.Set(0, 0, expectedValue)

	if !isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
	var actualValue = matrix.data[0]
	if actualValue != expectedValue {
		t.Errorf(complexErrorFormat, expectedValue, actualValue)
	}
}

func TestSetValueInTheMiddleOfMatrix(t *testing.T) {
	matrix, _ := New("1 2 3\n4 5 6\n7 8 9")
	var expectedValue = 100
	var isSet = matrix.Set(1, 1, expectedValue)
	fmt.Println(matrix.data)
	if !isSet {
		t.Errorf(complexErrorFormat, false, isSet)
	}
	var actualValue = matrix.data[4]
	if actualValue != expectedValue {
		t.Errorf(complexErrorFormat, expectedValue, actualValue)
	}
}
