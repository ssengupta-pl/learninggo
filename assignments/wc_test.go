package assignments

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

type WcTestError string

func (e WcTestError) Error() string {
	return string(e)
}

func testInternal(testString string, expected map[string]WordPositions) error {
	result := WordCount(testString)

	// Test if the map is returned at all.
	if result == nil {
		return WcTestError("No map returned")
	}

	// Test if each word in the test string is accounted for.
	arrayOfStrings := strings.Fields(testString)
	for _, val := range arrayOfStrings {
		if _, exists := result[val]; !exists {
			return WcTestError(fmt.Sprintf("Key %s could not be found", val))
		}
	}

	// Test if each word has the right information attached.
	for key, val := range expected {
		returnedResult := result[key]

		// First test for occurences.
		expectedOccurences := val.occurences
		returnedOccurences := returnedResult.occurences

		if expectedOccurences != returnedOccurences {
			return WcTestError(fmt.Sprintf("Mismatch in occurences for %s, expected %v, returned %v", key, expectedOccurences, returnedOccurences))
		}

		// Next test for positions.
		expectedPositions := val.positions
		returnedPositions := returnedResult.positions

		if slices.Compare(expectedPositions, returnedPositions) != 0 {
			return WcTestError(fmt.Sprintf("Mismatch in positions for %s, expected %v, returned %v", key, expectedPositions, returnedPositions))
		}
	}

	return nil
}

func TestWordCountForFreq1(t *testing.T) {
	testString := "I am learning Go!!"

	expected := make(map[string]WordPositions)
	expected["I"] = WordPositions{1, []int{0}}
	expected["am"] = WordPositions{1, []int{1}}
	expected["learning"] = WordPositions{1, []int{2}}
	expected["Go!!"] = WordPositions{1, []int{3}}

	testResults := testInternal(testString, expected)

	if testResults != nil {
		t.Fatal(testResults)
	}
}
func TestWordCountForFreqN(t *testing.T) {
	testString := "Hip Hooray! Hip Hip Hooray!"

	expected := make(map[string]WordPositions)
	expected["Hip"] = WordPositions{3, []int{0, 2, 3}}
	expected["Hooray!"] = WordPositions{2, []int{1, 4}}

	testResults := testInternal(testString, expected)

	if testResults != nil {
		t.Fatal(testResults)
	}
}

func TestWordPositionsStringForOnePosition(t *testing.T) {
	testString := "This"

	expected := "1 occurences in positions [0]:"

	result := WordCount(testString)
	wordPositionObj := result["This"]
	obtained := wordPositionObj.String()

	// Test if the toString method worked
	if strings.Compare(expected, obtained) != 0 {
		t.Fatalf("Mismatch in WordPositions tostring, expected %v, obtained %v", expected, obtained)
	}
}

func TestWordPositionsStringForManyPositions(t *testing.T) {
	testString := "Rat cat bat Rat that mat Rat"

	expected := "3 occurences in positions [0 3 6]:"

	result := WordCount(testString)
	wordPositionObj := result["Rat"]
	obtained := wordPositionObj.String()

	// Test if the toString method worked
	if strings.Compare(expected, obtained) != 0 {
		t.Fatalf("Mismatch in WordPositions tostring, expected %v, obtained %v", expected, obtained)
	}
}
