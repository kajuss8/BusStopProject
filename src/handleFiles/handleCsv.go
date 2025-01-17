package handleFiles

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func ReadFromCsvAllWithSameId(filePath, columnName, id string) [][]string {
	
	var result [][]string
	
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	headers[0] = strings.TrimPrefix(headers[0], "\ufeff")

	columnIndex := -1
	for i, header := range headers {
		if header == columnName {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		log.Fatalf("Column %s not found", columnName)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		if record[columnIndex] == id {
			result = append(result, record)
		}
	}

	return result
}

