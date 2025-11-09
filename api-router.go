package main

import (
	"net/http"
)

func ApiRouter () *http.ServeMux {
	local_multiplexer := http.NewServeMux()

	local_multiplexer.HandleFunc("/first", ApiMainHandlerFirst)
	local_multiplexer.HandleFunc("/second", ApiMainHandlerSecond)


	return local_multiplexer
}