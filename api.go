package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetApi() {
	const url = "http://localhost:8000/get"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code :", res.StatusCode)
	fmt.Println("Content Length :", res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Content Length :", byteCount)
	fmt.Println("Get Api Content :", responseString.String())
}

// Post Request with json Payload
func PostApi() {
	const url = "http://localhost:8000/post"
	// Payload data Creation
	requestBody := strings.NewReader(`{
		"name": "Prakash Anand",
		"age": "25",
		"Occupation" : "Software Engineer"
	}`)
	res, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content), res.StatusCode, res.Status)
}

// Post Request with Form Data
func PostApiWithFormPayload() {
	const myurl = "http://localhost:8000/postform"
	// FormData Creation
	data := url.Values{}
	data.Add("firstName", "Prakash")
	data.Add("lastName", "Anand")
	data.Add("email", "test@gmail.com")
	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content), res.StatusCode, res.Status)
}
