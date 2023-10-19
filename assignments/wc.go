package assignments

import (
	"fmt"
	"strings"
)

type WordPositions struct {
	occurences int
	positions  []int
}

func (wp WordPositions) String() string {
	return fmt.Sprintf("%d occurences in positions %v:", wp.occurences, wp.positions)
}

func WordCount(s string) map[string]WordPositions {
	arrayOfStrings := strings.Fields(s)

	wordPositionsMap := make(map[string]WordPositions)

	for idx, val := range arrayOfStrings {
		// Does the current word in the array version of the input sentence
		// exist in the map?
		wordPositionInstance, exists := wordPositionsMap[val]

		if !exists {
			// If word does not exist, then initialize the word position struct
			// with an empty positions array and a count of 0 for occurences.
			// This instance is then updated as a new entry in the map.
			wordPositionInstance = WordPositions{0, []int{}}
		}

		// If the word does exist or for the newly created struct for a word,
		// increment the count of the occurences and the latest found position
		// to the array.
		wordPositionInstance.occurences += 1
		wordPositionInstance.positions = append(wordPositionInstance.positions, idx)

		// Attach the created/updated word position struct for that word.
		wordPositionsMap[val] = wordPositionInstance
	}

	return wordPositionsMap
}
