package main

import (
	"fmt"
	"net/http"
)

func main ()  {
	const PORT = 8085
	multiplekser := http.NewServeMux();

	multiplekser.Handle("/api/", http.StripPrefix("/api", ApiRouter()))

	fmt.Printf("✅ Server started successfully on port %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), multiplekser); err != nil {
		fmt.Printf("Server failed\n")
	}
}