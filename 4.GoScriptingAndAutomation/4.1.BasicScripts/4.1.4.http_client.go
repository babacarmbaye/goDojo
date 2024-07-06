package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("HTTP Client:")
	baseURL := "https://httpbin.org"

	// http get request
	getResp, err := http.Get(baseURL + "/get")
	if err != nil {
		fmt.Println("Error on get request: ", err)
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)
	fmt.Println("GET response :", string(getBody))

	// http post request
	postBody := []byte(`{ "key" : "value" }`)
	postResp, err := http.Post(baseURL+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error on post request: ", err)
	}
	defer postResp.Body.Close()
	postBodyResp, _ := io.ReadAll(postResp.Body)
	fmt.Println("POST response :", string(postBodyResp))

	// http put request
	client := &http.Client{}
	putBody := []byte(`{ "key" : "updated body" }`)
	putReq, err := http.NewRequest(http.MethodPut, baseURL+"/put", bytes.NewBuffer(putBody))
	putReq.Header.Set("Content-Type", "application/json")
	putResp, err := client.Do(putReq)
	if err != nil {
		fmt.Println("Error on put request: ", err)
		return
	}
	defer putResp.Body.Close()
	putBodyResp, _ := io.ReadAll(putResp.Body)
	fmt.Println("Put response: ", string(putBodyResp))

	// http delete request
	deleteReq, err := http.NewRequest(http.MethodDelete, baseURL+"/delete", nil)
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		fmt.Println("Error on DELETE request:", err)
		return
	}
	defer deleteResp.Body.Close()
	deleteBodyResp, _ := io.ReadAll(deleteResp.Body)
	fmt.Println("DELETE Response:", string(deleteBodyResp))
}
