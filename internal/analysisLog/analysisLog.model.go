package analysisLog

import "time"

type AnalysisInDataStruct struct {
	Date     time.Time `json:"date"`
	FileName string    `json:"fileName"`
	ErrMsg   string    `json:"errMsg"`
}

type FilteredAnalysisDataStruct struct {
	DateFormatted string `json:"dateFormatted"`
	HourFormatted string `json:"hourFormatted"`
	FileName      string `json:"fileName"`
	ErrMsg        string `json:"errMsg"`
}
