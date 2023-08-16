package main

import (
	"fmt"
	"github.com/VeronicaAlexia/requests"
)

func main() {
	params := map[string]any{"key": "", "page": 1, "size": 5}
	headers := map[string]interface{}{"Content-Type": requests.ContentTypeForm}
	response := requests.Get("http://localhost:8080/api/v1/Demo/GetDemo", params, headers)
	fmt.Println(response.String())
	fmt.Println("=====================================")
	fmt.Println(response.Json())
	fmt.Println("=====================================")

}
