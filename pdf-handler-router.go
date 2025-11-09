package main

import (
	"net/http"
)

func PDFHandlerRouter () *http.ServeMux {
	local_multiplexer := http.NewServeMux()

	local_multiplexer.HandleFunc("POST /merge", PDFHandlerFunctionMerge)


	return local_multiplexer
}