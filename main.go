/*
 * Go automate for BrowserStack
 * taken from sourcegraph go-selenium examples
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"encoding/csv"
	"time"
	"strings"
	utils "github.com/FriendlyUser/google-definition-fetch/utils" 
)


func main() {
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}
	// Open the file
	csvfile, err := os.Open("stock_names.csv")
	if err != nil {
		log.Fatalln("Couldn't open the input csv file", err)
	}
	// new output file
	outputFile, err := os.Create("result.tsv")
	outputFile.WriteString("abbr\tfull_name\n")
	if err != nil {
		log.Fatalln("Couldn't open the output csv file", err)
	}
	defer outputFile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		searchTerm := record[1]
		newExtract, err := utils.FindTerm(searchTerm, &spaceClient)
		if err != nil {
			continue
		}
		newExtract = strings.TrimSuffix(newExtract, "\n")
		outputLine := fmt.Sprintf("%s\t%s\t%s", record[0], record[1], newExtract)
		outputFile.WriteString(outputLine)
		outputFile.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = utils.FindTerm("Ease of movement", &spaceClient)
}