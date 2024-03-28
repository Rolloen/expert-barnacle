package analysisLog

import (
	"errors"
	"testing"
	"time"
)

type testFormatDataToStrucTableStruct struct {
	name  string
	input [][]string
	want  []AnalysisInDataStruct
}

func assertFormatedAnalysisDataEqualsExpectedData(t *testing.T, expectedOutput []AnalysisInDataStruct, actualOutput []AnalysisInDataStruct) {
	for index, val := range expectedOutput {
		actOut := actualOutput[index]
		if actOut.Date.Compare(val.Date) != 0 {
			t.Errorf("Not same date, expected: %s but got: %s", val.Date, actOut.Date)
		}
		if actOut.ErrMsg != val.ErrMsg {
			t.Errorf("Not same error messsage , expected: %s but got: %s", val.ErrMsg, actOut.ErrMsg)
		}
		if actOut.FileName != val.FileName {
			t.Errorf("Not same file name , expected: %s but got: %s", val.FileName, actOut.FileName)
		}
	}
}
func TestFormatDataAnalysisToStructNoErrors(t *testing.T) {
	utc2 := time.FixedZone("CEST", 60*60*2)
	utc1 := time.FixedZone("CEST", 60*60)
	testTables := []testFormatDataToStrucTableStruct{
		{
			name: "Test FormatDataToStruct() : no special case , with 3 correct fields per subarray",
			input: [][]string{
				{"2019-04-30T12:01:39+02:00", "network.go", "Network connection established"},
				{"2019-04-30T12:01:42+02:00", "db.go", "Transaction failed"},
			},
			want: []AnalysisInDataStruct{
				{
					Date:     time.Date(2019, 04, 30, 12, 01, 39, 00, utc2),
					FileName: "network.go",
					ErrMsg:   "Network connection established",
				},
				{
					Date:     time.Date(2019, 04, 30, 12, 01, 42, 00, utc2),
					FileName: "db.go",
					ErrMsg:   "Transaction failed",
				},
			},
		},
		{
			name: `Test FormatDataToStruct() : wrong input format , have more than 3 fields
				 (the additionnal fields are caused by the use of "," in the Error message)`,
			input: [][]string{
				{"2023-02-11T21:09:51+01:00", "theMatrix.go", "Red pill taken", "welcome to the real world"},
				{"2023-02-12T00:32:23+01:00", "theHitchhiker.go",
					"Error: Failed to find the answer to the ultimate question of life", "the universe", "and everything"},
			},
			want: []AnalysisInDataStruct{
				{
					Date:     time.Date(2023, 02, 11, 21, 9, 51, 00, utc1),
					FileName: "theMatrix.go",
					ErrMsg:   "Red pill taken, welcome to the real world",
				},
				{
					Date:     time.Date(2023, 02, 12, 00, 32, 23, 00, utc1),
					FileName: "theHitchhiker.go",
					ErrMsg:   "Error: Failed to find the answer to the ultimate question of life, the universe, and everything",
				},
			},
		},
	}

	for _, tt := range testTables {
		ttValue := tt
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actualOutput, _ := FormatDataAnalysisToStruct(ttValue.input)
			// Assert
			assertFormatedAnalysisDataEqualsExpectedData(t, ttValue.want, actualOutput)
		})
	}
}

// Test FormatDataToStruct() : wrong input format , missing time field (not present or not at the first index)
// input : [[network.go,2019-04-30T12:01:39+00:00,Network connection established]]
func TestFormatDataAnalysisToStructWithNoTime(t *testing.T) {
	// arrange
	inputData := [][]string{
		{"network.go", "2019-04-30T12:01:39+02:00", "Network connection established"},
	}
	expectedErrOutput := errors.New("First field must be the time of the log")

	// act
	_, actualErr := FormatDataAnalysisToStruct(inputData)
	// assert
	if actualErr.Error() != expectedErrOutput.Error() {
		t.Errorf("No/wrong error thrown, expected : %s but got : %s", expectedErrOutput.Error(), actualErr.Error())
	}
}

// Test FormatDataToStruct() : empty input
// input : []
func TestFormatDataAnalysisToStructEmptyInput(t *testing.T) {
	// arrange
	inputData := [][]string{}
	expectedErrOutput := errors.New("No data input")

	// act
	_, actualErr := FormatDataAnalysisToStruct(inputData)
	// assert
	if actualErr.Error() != expectedErrOutput.Error() {
		t.Errorf("No/wrong error thrown, expected : %s but got : %s", expectedErrOutput.Error(), actualErr.Error())
	}
}
