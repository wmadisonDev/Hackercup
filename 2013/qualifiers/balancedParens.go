package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "flag"
  "regexp"
)

var inputFileName *string = flag.String("file", "", "The input file for this script")

var balanceFinder *regexp.Regexp

func main() {
	balanceFinder, _ = regexp.Compile("^([a-z\\s:]+|(:\\))+|(:\\()+|\\(([a-z\\s:]|(:\\))|(:\\())+\\))+$")

  flag.Parse()

  if *inputFileName == "" {
    panic("No Input File Provided!")
  }

  inputFile, openError := os.Open(*inputFileName)

  if openError != nil {
    errorDesc := fmt.Sprint("Error opening %s for reading!", inputFileName)
    panic(errorDesc)
  }

  bufferedInput := bufio.NewReader(inputFile)

  var currentLine []byte

  currentLine, isTruncated, readError := bufferedInput.ReadLine()

  if readError != nil {
    fmt.Fprintf(os.Stdout, "Error reading from %s!: ", inputFileName, readError)
  }

  //First line of file is always the number of test cases...
  numCases,_ := strconv.Atoi(string(currentLine))

  //For each line in the file
  for caseNumber := 1; caseNumber <= numCases && readError == nil; caseNumber++ {
    currentLine, isTruncated, readError = bufferedInput.ReadLine()

    if isTruncated {
      fmt.Fprintf(os.Stdout, "Buffer was too small! Line was truncated to %s ", string(currentLine))
    } else {
      evaluateMessage(caseNumber, string(currentLine))
    }

  }

  defer inputFile.Close()
}

func evaluateMessage(caseNumber int, currentLine string) {
  balanced := "NO"

  if isBalanced(currentLine) {
    balanced = "YES"
  }

  fmt.Fprintf(os.Stdout, "Case #%d: %s\n", caseNumber, balanced)
}

func isBalanced(text string) bool {
	balanced := false

	if text == "" {
		balanced = true
	} else {
		matches := balanceFinder.MatchString(text)
		balanced = matches
	}

	return balanced
}
