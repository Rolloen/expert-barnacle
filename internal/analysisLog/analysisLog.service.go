package analysisLog

import (
	"errors"
	"strings"
	"time"
)

// Parse a given CSV datas to an array of custom struct
//
// Take into account if fields of the CSV has errors (like more than 3 fields)
//
// The rule is that the fields in the CSV datas are always in this order : time -> filename -> error message
func FormatDataAnalysisToStruct(inpDatas [][]string) ([]AnalysisInDataStruct, error) {
	if len(inpDatas) == 0 {
		return []AnalysisInDataStruct{}, errors.New("No data input")
	}
	structuredDatas := make([]AnalysisInDataStruct, len(inpDatas))
	for i, data := range inpDatas {
		dataTime, err := time.Parse(time.RFC3339, data[0])
		if err != nil {
			return []AnalysisInDataStruct{}, errors.New("First field must be the time of the log")
		}
		structuredDatas[i] = AnalysisInDataStruct{
			Date:     dataTime,
			FileName: data[1],
		}
		if len(data) > 3 {
			structuredDatas[i].ErrMsg = strings.Join(data[2:], ", ")
		} else {
			structuredDatas[i].ErrMsg = data[2]
		}
	}
	return structuredDatas, nil
}
