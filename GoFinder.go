package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"regexp"
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

	// regular expression to remove ms formatting values for more accurate string matches
	reg, err := regexp.Compile("\\<(.*?)\\>")
	if err != nil {
		log.Fatal(err)
	}
	documentContent = reg.ReplaceAllString(documentContent, "")
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
		fmt.Println("\n Usage: GoFinder <filePath>" +
			"\n e.g. GoFinder C:\\Users\\joe.smith\\Desktop\\sample_contract.docx\n")
		return
	}

	keywords := make(map[string]bool)

	// keywords to check for, left in this fashion for readability
	// all keywords should be in lowercase to match text from docx file
	// items will be printed to the screen in this order
	text_check := []string {
	"expiration",
	"confidentiality",
	"background",
	"pricing",
	"executive summary",
	"assumptions",
	"scope of services",
	"travel",
	"schedule"}

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
	
	// Request submitted to print items in specific order every time for comparison ease
	// There may be a more efficient way than a double loop
	for _, o := range text_check {
		for keyword, exists := range keywords{
			if keyword == o {
				fmt.Println(exists, "\t", keyword)
			}
		}
	}
}
