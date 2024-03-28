package analysisLog

import (
	"errors"
	"fmt"
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

func FilterDatas(inputStructDatas []AnalysisInDataStruct) []FilteredAnalysisDataStruct {
	var filteredDatas []FilteredAnalysisDataStruct
	counter := make(map[string]map[string]int)
	// construct the map[string]map[string]int to count the nb de occurence
	// for each pair of filename/msg per day and hour
	for _, inputData := range inputStructDatas {
		formattedDate := inputData.Date.Format("02012006")
		formattedHour := inputData.Date.Hour()
		dateKey := fmt.Sprintf("%s-%d", formattedDate, formattedHour)
		if len(counter[dateKey]) == 0 {
			counter[dateKey] = make(map[string]int)
		}
		formattedMsgKey := fmt.Sprintf("%s-%s", inputData.FileName, inputData.ErrMsg)
		counter[dateKey][formattedMsgKey]++
	}

	// find the most occured pair of filename/errMsg and put it in the returned array
	for dateKey, msgMap := range counter {
		var mostOccuredMsg string
		max := 0
		for msgKey, count := range msgMap {
			if count > max {
				max = count
				mostOccuredMsg = msgKey
			}
		}

		splitedDateStr := strings.Split(dateKey, "-")
		date := splitedDateStr[0]
		hour := splitedDateStr[1]
		splitedMsgStr := strings.Split(mostOccuredMsg, "-")
		filename := splitedMsgStr[0]
		errMsg := splitedMsgStr[1]
		tempFilteredData := FilteredAnalysisDataStruct{
			DateFormatted: date,
			HourFormatted: hour,
			FileName:      filename,
			ErrMsg:        errMsg,
		}
		filteredDatas = append(filteredDatas, tempFilteredData)
	}

	return filteredDatas
}
