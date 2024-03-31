package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"techTest/internal/analysisLog"
	csvService "techTest/pkg/csv"
)

const FORMATTED_LOG_NAME = "formattedLog"

func GetDataLogHandler(w http.ResponseWriter, r *http.Request) {

	// Read the input CSV file
	data, err := csvService.ReadCSVFromSources(csvService.CSV_SOURCE_FILE)
	if err != nil {
		fmt.Println("error : ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Parse and format the input data to wanted output data struct
	structuredData, err := analysisLog.FormatDataAnalysisToStruct(data)
	if err != nil {
		fmt.Println("error : ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the structured data to find the most logged error grouped by day and hour
	filteredStructuredData := analysisLog.FilterDatas(structuredData)
	// NOTES: not the most optimal way but I couldn't find a optimized way of filtering data AND keeping the order
	sortedFilteredStrucData := analysisLog.SortStrucDataByDateAndHour(filteredStructuredData)
	// convert struct data to CSV readable data
	parsedCsvData := analysisLog.ConvertStructDataToCSVData(sortedFilteredStrucData)
	// write the ouput CSV data
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+FORMATTED_LOG_NAME+".csv")
	writer := csv.NewWriter(w)
	if err := writer.WriteAll(parsedCsvData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
