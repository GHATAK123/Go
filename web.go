package main

import (
	"encoding/json"
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

type course struct {
	Name     string `json:"coursename"`
	Price    float64
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// Json Encoding
func EncodeJson() {
	myCourse := []course{
		{"ReactJS BootCamp", 456.65, "Youtube", "abc123", []string{"webdev", "frontend", "js"}},
		{"VueJS BootCamp", 856.65, "Youtube", "abc123412", nil},
		{"Angular BootCamp", 1456.65, "Youtube", "abczxc123", []string{"webdev", "fullstack", "js"}},
	}

	// finalJson, err := json.Marshal(myCourse)
	finalJson, err := json.MarshalIndent(myCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", finalJson)
	fmt.Println(string(finalJson))
}

// JSON Decoding(Json Consume)
func JSONDecode() {
	jsonData := []byte(`
	{
		"coursename": "ReactJS BootCamp",
		"Price": 456.65,
		"website": "Youtube",
		"tags": ["webdev","frontend","js"]
    }
	`)

	var myCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON Was Valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%T \t %T\n", jsonData, myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Invalid Json Data")
	}

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)
	fmt.Println()
	for key, value := range myData {
		fmt.Printf("%v : %v\n", key, value)
	}
}
