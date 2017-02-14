package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetDocumentContent(filename string) (documentContent string) {

	if !strings.HasSuffix(filename,"docx") {
		fmt.Println("Not a valid docx file")
		return
	}
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var f *zip.File
	var found bool
	// Iterate through the files in the archive,
	// Searching for document.xml
	for _, f = range r.File {
		if strings.EqualFold(f.Name, "word/document.xml") {
			found = true
			break
		}
	}

	if !found {
		log.Fatal("Not a valid docx file")
		return
	}

	// open word/document.xml
	rc, err := f.Open()

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	documentContent = strings.ToLower(string(data)) //right here do a "to lower" for search matching ease
	return
}

func SearchKeywords(data string, keywords map[string]bool) {
	for keyword, _ := range keywords {
		if strings.Contains(data, keyword) {
			keywords[keyword] = true
		}
	}
}

func main() {
	fmt.Printf("\n")
	if len(os.Args) < 2 {
		fmt.Println("\n Usage: GoFinder <filePath> <keyword1> [<keyword2>...n]" +
			"\n e.g. GoFinder C:\\Users\\joe.smith\\Desktop\\sample_contract.docx\n")
		return
	}

	keywords := make(map[string]bool)

	// keywords to check for, left in this fashion for readability
	// all keywords should be in lowercase to match text from docx file
	var text_check [10]string
	text_check[0] = "expiration"
	text_check[1] = "confidentiality"
	text_check[2] = "background"
	text_check[3] = "pricing"
	text_check[4] = "executive summary"
	text_check[5] = "assumptions"
	text_check[6] = "scope of services"
	text_check[7] = "pricing"
	text_check[8] = "travel"
	text_check[9] = "schedule"


	for _, keyword := range text_check {
		keywords[keyword] = false
	}

	filePath := os.Args[1]

	documentContent := GetDocumentContent(filePath)
	if len(documentContent) == 0 {
		fmt.Println("Empty document")
		return
	}

	SearchKeywords(documentContent, keywords)

	for keyword, exists := range keywords {
		fmt.Println(exists, "\t", keyword)
	}
}
