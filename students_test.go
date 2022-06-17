package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	stringErrorFormat  = "Expected: %s, got %s"
	intErrorFormat     = "Expected: %d, got %d"
	complexErrorFormat = "Expected: %v, got %v"
	invalidSyntaxError = "strconv.Atoi: parsing \"\": invalid syntax"

	simpleMatrixString               = "1 2 3\n4 5 6\n7 8 9"
	matrixWithNegativeElementsString = "-1 2 3\n4 -5 6\n7 8 -9"
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

// WRITE YOUR CODE BELOW

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

//People LEN tests

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

func TestLessFunction(t *testing.T) {
	tData := []struct {
		I        int
		J        int
		Expected bool
	}{
		{I: 1, J: 2, Expected: true},
		{I: 0, J: 3, Expected: false},
		{I: 2, J: 3, Expected: false},
		{I: 4, J: 5, Expected: false},
	}
	var peoples = prepareTestData()
	for _, v := range tData {
		var actual = peoples.Less(v.I, v.J)
		if v.Expected != actual {
			t.Errorf(complexErrorFormat, v.Expected, actual)
		}
	}
}

// People swap tests

func TestSwapFunction(t *testing.T) {
	tData := []struct {
		FirstIndex  int
		SecondIndex int
	}{
		{FirstIndex: 0, SecondIndex: 3},
		{FirstIndex: 3, SecondIndex: 3},
		{FirstIndex: 5, SecondIndex: 0},
	}
	for _, v := range tData {
		var peoples = prepareTestData()
		var originalCollection = prepareTestData()

		peoples.Swap(v.FirstIndex, v.SecondIndex)

		var expected = originalCollection[v.SecondIndex].firstName
		var actual = peoples[v.FirstIndex].firstName

		if actual != expected {
			t.Errorf(stringErrorFormat, expected, actual)
		}
		expected = originalCollection[v.FirstIndex].firstName
		actual = peoples[v.SecondIndex].firstName
		if actual != expected {
			t.Errorf(stringErrorFormat, expected, actual)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Matrix tests
func TestIncorrectMatrixCreationFunction(t *testing.T) {
	tData := []struct {
		InitialString string
		ErrMsg        string
	}{
		{InitialString: "", ErrMsg: invalidSyntaxError},
		{InitialString: " ", ErrMsg: invalidSyntaxError},
		{InitialString: "1\n ", ErrMsg: invalidSyntaxError},
		{InitialString: "1 1 1\n1 1\n1 1 1", ErrMsg: "Rows need to be the same length"},
		{InitialString: "1 1 x\n1 1 y\n1 1 z", ErrMsg: "strconv.Atoi: parsing \"x\": invalid syntax"},
	}
	for _, v := range tData {
		_, err := New(v.InitialString)
		if err.Error() != v.ErrMsg {
			t.Errorf(stringErrorFormat, v.ErrMsg, err)
		}
	}
}

func TestValidMatrixCreationFunction(t *testing.T) {
	tData := []struct {
		InitialString  string
		ExpectedMatrix []int
		ErrMsg         error
	}{
		{InitialString: "1 1 1\n1 1 1\n1 1 1", ExpectedMatrix: []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, ErrMsg: nil},
		{InitialString: "1", ExpectedMatrix: []int{1}, ErrMsg: nil},
		{InitialString: "-1 -2\n-3 -4\n-5 -6", ExpectedMatrix: []int{-1, -2, -3, -4, -5, -6}, ErrMsg: nil},
	}
	for _, v := range tData {
		matrix, err := New(v.InitialString)
		if err != nil {
			t.Errorf(complexErrorFormat, nil, err)
		}
		if !reflect.DeepEqual(matrix.data, v.ExpectedMatrix) {
			t.Errorf(complexErrorFormat, v.ExpectedMatrix, matrix.data)
		}
	}
}

// Matrix rows tests
func TestMatrixRowsFunction(t *testing.T) {
	tData := []struct {
		InitialString string
		ExpectedRows  [][]int
		ErrMsg        error
	}{
		{
			InitialString: simpleMatrixString,
			ExpectedRows: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			ErrMsg: nil,
		},
		{
			InitialString: "1",
			ExpectedRows:  [][]int{{1}},
			ErrMsg:        nil,
		},
		{
			InitialString: matrixWithNegativeElementsString,
			ExpectedRows: [][]int{
				{-1, 2, 3},
				{4, -5, 6},
				{7, 8, -9},
			},
			ErrMsg: nil,
		},
	}
	for _, v := range tData {
		matrix, err := New(v.InitialString)
		if !reflect.DeepEqual(v.ExpectedRows, matrix.Rows()) {
			t.Errorf(complexErrorFormat, v.ExpectedRows, matrix.Rows())
		}
		if err != nil {
			t.Errorf(complexErrorFormat, "No errors expected", err)
		}
	}
}

// Matrix cols tests

func TestMatrixColsFunction(t *testing.T) {
	tData := []struct {
		InitialString string
		ExpectedCols  [][]int
		ErrMsg        error
	}{
		{
			InitialString: simpleMatrixString,
			ExpectedCols: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			ErrMsg: nil,
		},
		{
			InitialString: "8",
			ExpectedCols:  [][]int{{8}},
			ErrMsg:        nil,
		},
		{
			InitialString: matrixWithNegativeElementsString,
			ExpectedCols: [][]int{
				{-1, 4, 7},
				{2, -5, 8},
				{3, 6, -9},
			},
			ErrMsg: nil,
		},
	}
	for _, v := range tData {
		matrix, err := New(v.InitialString)
		if !reflect.DeepEqual(v.ExpectedCols, matrix.Cols()) {
			t.Errorf(complexErrorFormat, v.ExpectedCols, matrix.Cols())
		}
		if err != nil {
			t.Errorf(complexErrorFormat, "No errors expected", err)
		}
	}
}

// Matrix set tests
func TestNegativeValuesMatrixSetFunction(t *testing.T) {
	tData := []struct {
		InitialString string
		Row           int
		Col           int
		Value         int
		ExpInSet      bool
	}{
		{InitialString: simpleMatrixString, Row: -1, Col: 1, Value: 1, ExpInSet: false},
		{InitialString: simpleMatrixString, Row: 4, Col: 1, Value: 1, ExpInSet: false},
		{InitialString: simpleMatrixString, Row: 3, Col: 1, Value: 1, ExpInSet: false},
		{InitialString: simpleMatrixString, Row: 1, Col: -1, Value: 1, ExpInSet: false},
		{InitialString: simpleMatrixString, Row: 1, Col: 4, Value: 1, ExpInSet: false},
		{InitialString: simpleMatrixString, Row: 1, Col: 3, Value: 1, ExpInSet: false},
	}
	for _, v := range tData {
		matrix, _ := New(v.InitialString)
		var inSet = matrix.Set(v.Row, v.Col, v.Value)

		if inSet != v.ExpInSet {
			t.Errorf(complexErrorFormat, v.ExpInSet, inSet)
		}
	}
}

func TestValidValuesMatrixSetFunction(t *testing.T) {
	tData := []struct {
		InitialString string
		Row           int
		Col           int
		Value         int
		ExpInSet      bool
		ValueIndex    int
	}{
		{InitialString: simpleMatrixString, Row: 0, Col: 0, Value: 100, ExpInSet: true, ValueIndex: 0},
		{InitialString: simpleMatrixString, Row: 1, Col: 1, Value: 75, ExpInSet: true, ValueIndex: 4},
		{InitialString: matrixWithNegativeElementsString, Row: 1, Col: 1, Value: -100, ExpInSet: true, ValueIndex: 4},
	}
	for _, v := range tData {
		matrix, _ := New(v.InitialString)
		var inSet = matrix.Set(v.Row, v.Col, v.Value)

		if inSet != v.ExpInSet {
			t.Errorf(complexErrorFormat, v.ExpInSet, inSet)
		}
		var actualValue = matrix.data[v.ValueIndex]
		if actualValue != v.Value {
			t.Errorf(complexErrorFormat, v.Value, actualValue)
		}
	}
}
