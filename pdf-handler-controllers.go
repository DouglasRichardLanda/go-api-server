package main

import (
	"fmt"
	"net/http"
)

// this function will merge some pdfs for me into 1.
func PDFHandlerFunctionMerge (res http.ResponseWriter, req *http.Request) {
	fmt.Println("PDFHandlerFunctionMerge")
}
