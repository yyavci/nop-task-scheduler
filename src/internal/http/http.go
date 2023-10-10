package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type HttpResponse struct {
	Status     string
	StatusCode int
	Data       any
}

func PostJsonRequest(url string, json string) (HttpResponse, error) {
	fmt.Printf("sending http post request with json data...\n")
	fmt.Printf("url:%s\n", url)

	body := []byte(json)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("Error creating post request! Err:%+v\n", err)
		return HttpResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending post request! Err:%+v\n", err)
		return HttpResponse{}, err
	}

	var httpResponse HttpResponse
	httpResponse.Status = res.Status
	httpResponse.StatusCode = res.StatusCode

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error sending post request! Err:%+v\n", err)
		return httpResponse, err
	}

	data := string(resBody)

	httpResponse.Data = data

	return httpResponse, nil
}
