package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"techTest/internal/analysisLog"
)

const FORMATTED_LOG_NAME = "formattedLog"

func GetDataLogHandler(w http.ResponseWriter, r *http.Request) {

	// Read the input CSV file
	file, err := os.Open("journaux.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error : ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Parse and format the input data to wanted output data struct
	structuredData, _ := analysisLog.FormatDataAnalysisToStruct(data)

	// Use the structured data to find the most logged error grouped by day and hour
	filteredStructuredData := analysisLog.FilterDatas(structuredData)
	log.Println(filteredStructuredData)
	// write the ouput CSV data
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+FORMATTED_LOG_NAME+".csv")
	writer := csv.NewWriter(w)
	if err := writer.WriteAll(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
