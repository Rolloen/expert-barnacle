package analysisLog

import "time"

type AnalysisInDataStruct struct {
	Date     time.Time `json:"date"`
	FileName string    `json:"fileName"`
	ErrMsg   string    `json:"errMsg"`
}
