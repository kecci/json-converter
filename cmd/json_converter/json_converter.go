package json_converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/nqd/flat"
	"github.com/xuri/excelize/v2"
)

// JsonToExcel converter
func JsonToExcel(jsonLocation, excelLocation string) {

	// Location file path
	// > /Users/kecci/Downloads/example.json
	// > C:/Program Files/Documents/example.json
	jsonFile, err := os.Open(jsonLocation)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonMaps []map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &jsonMaps)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonNormalized := BulkNormalize(jsonMaps)
	if len(jsonNormalized) == 0 {
		return
	}

	f := excelize.NewFile()
	sheetName := "Sheet1"
	sheetIndex := f.GetSheetIndex(sheetName)

	// Header
	jsonFirst := jsonNormalized[0]
	headerKeys := make([]string, 0, len(jsonFirst))
	for k := range jsonFirst {
		headerKeys = append(headerKeys, k)
	}
	sort.Strings(headerKeys)

	columnHeaderCount := 1
	for _, k := range headerKeys {
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", GetColumnName(columnHeaderCount)), k)
		columnHeaderCount += 1
	}

	rowCount := 2
	for _, jsonData := range jsonNormalized {
		columnCount := 1
		for _, k := range headerKeys {
			f.SetCellValue(sheetName, fmt.Sprintf("%s%v", GetColumnName(columnCount), rowCount), jsonData[k])
			columnCount += 1
		}
		rowCount += 1
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(sheetIndex)

	// Save spreadsheet by the given path.
	// Location file path
	// > /Users/kecci/Downloads/example.xlsx
	// > C:/Program Files/Documents/example.xlsx
	if err := f.SaveAs(excelLocation); err != nil {
		log.Fatal(err)
	}

	println("Success saved on", excelLocation)
}

// BulkNormalize returns a BulkNormalize
func BulkNormalize(jsonMaps []map[string]interface{}) (res []map[string]interface{}) {
	for _, jsonMap := range jsonMaps {
		out, err := flat.Flatten(jsonMap, nil)
		if err != nil {
			return res
		}

		res = append(res, out)
	}
	return res
}

// GetColumnName excel
func GetColumnName(col int) string {
	name := make([]byte, 0, 3) // max 16,384 columns (2022)
	const aLen = 'Z' - 'A' + 1 // alphabet length
	for ; col > 0; col /= aLen + 1 {
		name = append(name, byte('A'+(col-1)%aLen))
	}
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}
