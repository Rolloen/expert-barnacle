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

type testFilterDataTableStruct struct {
	name  string
	input []AnalysisInDataStruct
	want  []FilteredAnalysisDataStruct
}

var (
	utc2 = time.FixedZone("CEST", 60*60*2)
	utc1 = time.FixedZone("CEST", 60*60)
)

/* FormatDataAnalysisToStruct() TESTS */

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
			name: `Test FormatDataToStruct() : wrong input format, have more than 3 fields
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

/* FilterDatas() TESTS */
func assertFilteredDataEqualsExpectedData(t *testing.T, expectedOutput []FilteredAnalysisDataStruct, actualOutput []FilteredAnalysisDataStruct) {
	for index, val := range expectedOutput {
		actOut := actualOutput[index]
		if actOut.DateFormatted != val.DateFormatted {
			t.Errorf("Not same formatted date , expected: %s but got: %s", val.DateFormatted, actOut.DateFormatted)
		}
		if actOut.HourFormatted != val.HourFormatted {
			t.Errorf("Not same formatted hour , expected: %s but got: %s", val.HourFormatted, actOut.HourFormatted)
		}
		if actOut.ErrMsg != val.ErrMsg {
			t.Errorf("Not same error messsage , expected: %s but got: %s", val.ErrMsg, actOut.ErrMsg)
		}
		if actOut.FileName != val.FileName {
			t.Errorf("Not same file name , expected: %s but got: %s", val.FileName, actOut.FileName)
		}
	}
}
func TestFilterDatasWithValidInput(t *testing.T) {
	testTables := []testFilterDataTableStruct{
		{
			name: "Test FilterDatas() : valid input with 3 struct",
			input: []AnalysisInDataStruct{
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
				{
					Date:     time.Date(2019, 04, 30, 12, 10, 42, 00, utc2),
					FileName: "network.go",
					ErrMsg:   "Network connection established",
				},
			},
			want: []FilteredAnalysisDataStruct{
				{
					DateFormatted: "30042019",
					HourFormatted: "12",
					FileName:      "network.go",
					ErrMsg:        "Network connection established",
				},
			},
		},
		{
			name: `Test FilterDatas() :valid input with 4 input, 2 diffrent pair of date/hour`,
			input: []AnalysisInDataStruct{
				{
					Date:     time.Date(2022, 03, 12, 15, 01, 39, 00, utc2),
					FileName: "memeGenerator.go",
					ErrMsg:   "Error: Failed to generate a meme",
				},
				{
					Date:     time.Date(2022, 03, 11, 11, 01, 42, 00, utc2),
					FileName: "theHitchhiker.go",
					ErrMsg:   "Error: So long, and thanks for all the fish",
				},
				{
					Date:     time.Date(2022, 03, 12, 15, 10, 42, 00, utc2),
					FileName: "memeGenerator.go",
					ErrMsg:   "Error: Failed to generate a meme",
				},
				{
					Date:     time.Date(2022, 03, 11, 11, 10, 42, 00, utc2),
					FileName: "theHitchhiker.go",
					ErrMsg:   "Error: So long, and thanks for all the fish",
				},
			},
			want: []FilteredAnalysisDataStruct{
				{
					DateFormatted: "12032022",
					HourFormatted: "15",
					FileName:      "memeGenerator.go",
					ErrMsg:        "Error: Failed to generate a meme",
				},
				{
					DateFormatted: "11032022",
					HourFormatted: "11",
					FileName:      "theHitchhiker.go",
					ErrMsg:        "Error: So long, and thanks for all the fish",
				},
			},
		},
	}
	for _, tt := range testTables {
		ttValue := tt
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actualOutput := FilterDatas(ttValue.input)
			// Assert
			assertFilteredDataEqualsExpectedData(t, ttValue.want, actualOutput)
		})
	}
}

// special case where the hour is >=23 then it count for the next day, hour 00:00
func TestFilterDatasWithSpecialCase(t *testing.T) {

	// Arrange
	input := []AnalysisInDataStruct{
		{
			Date:     time.Date(2019, 04, 13, 23, 01, 39, 00, utc2),
			FileName: "network.go",
			ErrMsg:   "Network connection established",
		},
		{
			Date:     time.Date(2019, 04, 13, 12, 01, 42, 00, utc2),
			FileName: "db.go",
			ErrMsg:   "Transaction failed",
		},
		{
			Date:     time.Date(2019, 04, 14, 00, 10, 42, 00, utc2),
			FileName: "network.go",
			ErrMsg:   "Network connection established",
		},
		{
			Date:     time.Date(2019, 04, 14, 00, 10, 42, 00, utc2),
			FileName: "db.go",
			ErrMsg:   "Transaction failed",
		},
	}

	expectedOutput := []FilteredAnalysisDataStruct{
		{
			DateFormatted: "14042019",
			HourFormatted: "00",
			FileName:      "network.go",
			ErrMsg:        "Network connection established",
		},
		{
			DateFormatted: "13042019",
			HourFormatted: "12",
			FileName:      "db.go",
			ErrMsg:        "Transaction failed",
		},
	}

	// Act
	actualOuput := FilterDatas(input)
	// Assert
	assertFilteredDataEqualsExpectedData(t, expectedOutput, actualOuput)

}

// // special case where there are no most
// func TestFilterDatasWithSpecialCaseError(t *testing.T) {

// 	// Arrange
// 	input := []AnalysisInDataStruct{
// 		{
// 			Date:     time.Date(2019, 04, 13, 23, 01, 39, 00, utc2),
// 			FileName: "network.go",
// 			ErrMsg:   "Network connection established",
// 		},
// 		{
// 			Date:     time.Date(2019, 04, 13, 12, 01, 42, 00, utc2),
// 			FileName: "db.go",
// 			ErrMsg:   "Transaction failed",
// 		},
// 		{
// 			Date:     time.Date(2019, 04, 14, 00, 10, 42, 00, utc2),
// 			FileName: "network.go",
// 			ErrMsg:   "Network connection established",
// 		},
// 		{
// 			Date:     time.Date(2019, 04, 14, 00, 10, 42, 00, utc2),
// 			FileName: "db.go",
// 			ErrMsg:   "Transaction failed",
// 		},
// 	}

// 	expectedOutput := []FilteredAnalysisDataStruct{
// 		{
// 			DateFormatted: "14042019",
// 			HourFormatted: "00",
// 			FileName:      "network.go",
// 			ErrMsg:        "Network connection established",
// 		},
// 		{
// 			DateFormatted: "13042019",
// 			HourFormatted: "12",
// 			FileName:      "db.go",
// 			ErrMsg:        "Transaction failed",
// 		},
// 	}

// 	// Act
// 	actualOuput := FilterDatas(input)
// 	// Assert
// 	assertFilteredDataEqualsExpectedData(t, expectedOutput, actualOuput)

// }
