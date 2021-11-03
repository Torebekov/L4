package test

import (
	"github.com/Torebekov/L4/internal/topwords"
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T){
	type nWords struct{
		words string
		num int
	}
	testTable := []struct{
		numWords nWords
		expected []string
	}{
		{
			numWords: nWords{"one,one  two. two!one", 1},
			expected: []string{"one"},
		},
		{
			numWords: nWords{"one,one three1three  two. three!", 2},
			expected: []string{"three","one"},
		},
	}
	for _, testCase := range testTable{
		result := topwords.TopWords(testCase.numWords.words,testCase.numWords.num)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect %v, got %v",
				testCase.expected, result)
		}
	}
}
