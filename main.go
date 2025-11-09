package main

import (
	"fmt"
	"net/http"
)

func main ()  {
	const PORT = 8085
	multiplekser := http.NewServeMux();


	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), multiplekser); err != nil {
		fmt.Printf("Server successfully started on port:: %d", PORT)
	}
}