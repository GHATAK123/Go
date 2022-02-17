package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// {"status":"success","data":{"id":1,"employee_name":"Tiger Nixon","employee_salary":320800,"employee_age":61,"profile_image":""},"message":"Successfully! Record has been fetched."}
func makeGet(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", data)
	fmt.Println()
	content := string(data)
	fmt.Println(content)
}

func urlHandler(myurl string) {
	info, _ := url.Parse(myurl)
	fmt.Println(info.Scheme)
	fmt.Println(info.Host)
	fmt.Println(info.Path)
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "scm-preprod.maersk.com",
		Path:   "/cshub/dashboard",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
