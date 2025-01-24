package handleFiles

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func ReadFile(filePath string) ([][]string ,error){

	file, err := os.Open(filePath)  
    if err != nil { 
        log.Fatal("Error while reading the file", err) 
    } 
  
    defer file.Close() 
  
    reader := csv.NewReader(file) 

    headers, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	headers[0] = strings.TrimPrefix(headers[0], "\ufeff")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	return records, nil
}