package binaryTree

import (
	"bytes"
	"fmt"
	"testing"
)

type TestCase struct {
	id             int
	description    string
	input          Input
	expectedResult ExpectedResult
}

type Input struct {
	values []int
}

type ExpectedResult struct {
	stringTree string
}

type TestSuite []TestCase

var tsInsert = TestSuite{
	TestCase{
		id:          1,
		description: "Insert multiply values",
		input: Input{
			values: []int{5, 4, 6, 8, 9, 7, 1, 3, 2, 26, 88, 29, 0, 11, 5, 6},
		},
		expectedResult: ExpectedResult{
			stringTree: `------------------------------------------------
                                   ┌──────[ 88
                                          └──────[ 29
                            ┌──────[ 26
                                   └──────[ 11
                     ┌──────[ 9
              ┌──────[ 8
                     └──────[ 7
       ┌──────[ 6
──────[ 5
       └──────[ 4
                     ┌──────[ 3
                            └──────[ 2
              └──────[ 1
                     └──────[ 0
------------------------------------------------
`,
		},
	},
	TestCase{
		id:          2,
		description: "Empty init input",
		input: Input{
			values: nil,
		},
		expectedResult: ExpectedResult{
			stringTree: `------------------------------------------------
------------------------------------------------
`,
		},
	},
	TestCase{
		id:          3,
		description: "One value repeated",
		input: Input{
			values: []int{99, 99, 99, 99, 99, 99},
		},
		expectedResult: ExpectedResult{
			stringTree: `------------------------------------------------
──────[ 99
------------------------------------------------
`,
		},
	},
}

func TestTreeNode_Init(t *testing.T) {

	for _, tc := range tsInsert {
		var outWriter = new(bytes.Buffer)

		t.Run(fmt.Sprintf("Test%d", tc.id), func(t *testing.T) {

			tree := InitTree(tc.input.values...)
			tree.Print(outWriter)
			result := outWriter.String()
			if result != tc.expectedResult.stringTree {
				fmt.Printf(" - Got: [%s]\n - Expected: [%s]\n", result, tc.expectedResult.stringTree)
				t.Fail()
			}
		})
	}

}
