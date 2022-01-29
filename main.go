package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Variable Decleration
	// var implictVariable = "Hello Go!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
	// fmt.Println(implictVariable)
	// var explictVariable string = "Hello World"
	// fmt.Println(explictVariable)
	// var flag = false
	// fmt.Printf("Data type of flag is %T\n", flag)
	// var a int
	// fmt.Println(a)
	// withoutVar := "This is new type variable decleration without var keyword"
	// fmt.Println(withoutVar)
	// withoutVar1 := "These type of variables only decleared inside functions"
	// fmt.Println(withoutVar1)

	// Input From Console
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Print("You Entered: ", input)

	// Type Casting String To Integer
	// reader1 := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter integer number: ")
	// numInput, _ := reader1.ReadString('\n')
	// aInt, err := strconv.ParseInt(strings.TrimSpace(numInput), 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	//  panic("Entered Value is not an Integer")
	// } else {
	// 	fmt.Println(aInt)
	// 	fmt.Printf("%T", aInt)
	// }

	// // Type Casting String To Float
	// reader2 := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter float Number: ")
	// floatInput, _ := reader2.ReadString('\n')
	// aFloat, err := strconv.ParseFloat(strings.TrimSpace(floatInput), 64)
	// if err != nil {
	// 	fmt.Println(err)
	//  panic("Entered Value is not a float")
	// } else {
	// 	fmt.Println(aFloat)
	// 	fmt.Printf("%T\n", aFloat)
	// }

	// Type Casting String To Boolean
	// reader3 := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter boolean: ")
	// boolInput, _ := reader3.ReadString('\n')
	// aBool, err := strconv.ParseBool(strings.TrimSpace(boolInput))
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Entered Value is not a boolean")
	// } else {
	// 	fmt.Println(aBool)
	// 	fmt.Printf("%T\n", aBool)
	// }

	type UserPermissions struct {
		Groups []struct {
			Name  string `json:"name"`
			Roles []struct {
				Name            string `json:"name"`
				ApplicationName string `json:"applicationName"`
				ApplicationID   int    `json:"applicationId"`
			} `json:"roles"`
		} `json:"groups"`
		Companies []struct {
			ID               int           `json:"id"`
			BeCode           string        `json:"beCode"`
			CountryCode      string        `json:"countryCode"`
			FunctionCode     string        `json:"functionCode"`
			Type             string        `json:"type"`
			RelatedCompanies []interface{} `json:"relatedCompanies"`
		} `json:"companies"`
		UserInfo struct {
			ID        int    `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			UserName  string `json:"userName"`
			UserType  string `json:"userType"`
		} `json:"userInfo"`
	}

	type UserRole struct {
		Name         string
		Priority     int
		CheckShipper bool
	}

	userPermissions := `{
		"groups": [
			{
				"name": "SCMFLOWConsignee",
				"roles": [
					{
						"name": "SCMFLOWConsignee",
						"applicationName": "SCM AEC",
						"applicationId": 0
					}
				]
			}
		],
		"companies": [
			{
				"id": 0,
				"beCode": "EECOJLA",
				"countryCode": "US",
				"functionCode": "HQ",
				"type": "Consignee"
			},
			{
				"id": 0,
				"beCode": "LEGACYFU",
				"countryCode": "US",
				"functionCode": "HEADQ",
				"type": "Consignee"
			},
			{
				"id": 0,
				"beCode": "TIMBER",
				"countryCode": "VN",
				"functionCode": "VDR",
				"type": "Shipper"
			},
			{
				"id": 0,
				"beCode": "TOCCLTD",
				"countryCode": "VN",
				"functionCode": "VDR",
				"type": "Shipper"
			},
			{
				"id": 0,
				"beCode": "JZHFCL23",
				"countryCode": "CN",
				"functionCode": "VDR",
				"type": "Shipper"
			}
		],
		"userInfo": {
			"id": 0,
			"firstName": "Nscp",
			"lastName": "FlowConsignee",
			"email": "nscptest@hotmail.com",
			"userName": "nscpflowconsignee",
			"userType": "external"
		}
	}`

	fmt.Println(userPermissions)
	fmt.Printf("%T", userPermissions)
	up := UserPermissions{}
	fmt.Println("-----------------------------------------------")
	fmt.Println(up)
	fmt.Println("-----------------------------------------------")
	err := json.Unmarshal([]byte(userPermissions), &up)
	fmt.Println(up.Companies)
	fmt.Printf("%T\n", up.Companies)
	for _, i := range up.Companies {
		fmt.Println(i.BeCode)
	}
	fmt.Println(err)

	//Get User Permission FunctionCode
	r, err := http.NewRequest("GET", "", nil)
	fmt.Println(r)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Set(
		"User-Permissions",
		"eyJncm91cHMiOlt7ImlkIjpudWxsLCJuYW1lIjoiU0NNRkxPV0NvbnNpZ25lZSIsImRlc2NyaXB0aW9uIjpudWxsLCJyb2xlcyI6W3siaWQiOm51bGwsIm5hbWUiOiJTQ01GTE9XQ29uc2lnbmVlIiwiYXBwbGljYXRpb25OYW1lIjoiU0NNIEFFQyIsImFwcGxpY2F0aW9uSWQiOjB9XX1dLCJjb21wYW5pZXMiOlt7ImlkIjowLCJiZUNvZGUiOiJFRUNPSkxBIiwiY291bnRyeUNvZGUiOiJVUyIsImZ1bmN0aW9uQ29kZSI6IkhRIiwidHlwZSI6IkNvbnNpZ25lZSIsIm5hbWUiOm51bGwsInJlbGF0ZWRDb21wYW5pZXMiOltdfSx7ImlkIjowLCJiZUNvZGUiOiJMRUdBQ1lGVSIsImNvdW50cnlDb2RlIjoiVVMiLCJmdW5jdGlvbkNvZGUiOiJIRUFEUSIsInR5cGUiOm51bGwsIm5hbWUiOm51bGwsInJlbGF0ZWRDb21wYW5pZXMiOltdfSx7ImlkIjowLCJiZUNvZGUiOiJUSU1CRVIiLCJjb3VudHJ5Q29kZSI6IlZOIiwiZnVuY3Rpb25Db2RlIjoiVkRSIiwidHlwZSI6IlNoaXBwZXIiLCJuYW1lIjpudWxsLCJyZWxhdGVkQ29tcGFuaWVzIjpbXX0seyJpZCI6MCwiYmVDb2RlIjoiVE9DQ0xURCIsImNvdW50cnlDb2RlIjoiVk4iLCJmdW5jdGlvbkNvZGUiOiJWRFIiLCJ0eXBlIjoiU2hpcHBlciIsIm5hbWUiOm51bGwsInJlbGF0ZWRDb21wYW5pZXMiOltdfSx7ImlkIjowLCJiZUNvZGUiOiJKWkhGQ0wyMyIsImNvdW50cnlDb2RlIjoiQ04iLCJmdW5jdGlvbkNvZGUiOiJWRFIiLCJ0eXBlIjoiU2hpcHBlciIsIm5hbWUiOm51bGwsInJlbGF0ZWRDb21wYW5pZXMiOltdfV0sInVzZXJJbmZvIjp7ImlkIjowLCJmaXJzdE5hbWUiOiJOc2NwIiwibGFzdE5hbWUiOiJGbG93Q29uc2lnbmVlIiwiZW1haWwiOiJhYmhpc2hlay5zb25pMDhAaG90bWFpbC5jb20iLCJ1c2VyTmFtZSI6Im5zY3BmbG93Y29uc2lnbmVlIiwidXNlclR5cGUiOiJleHRlcm5hbCJ9fQ==",
	)
	// b64EncodedUserPermissions := r.Header.Get("User-Permissions")
	// fmt.Println(b64EncodedUserPermissions)
	// userPermissions, err := base64.StdEncoding.DecodeString(b64EncodedUserPermissions)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// fmt.Println(userPermissions)
	// fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// var up UserPermissions
	// err = json.Unmarshal(userPermissions, &up)
	// fmt.Println(up) // final return statement from GetuserPermissions function
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Get User role function
	roles := []UserRole{
		{"SCM System Admin", 1, false},
		{"SCM Origin Customer Service Admin", 2, false},
		{"SCMFlowCustomerService", 3, false},
		{"SCM System Viewer", 4, false},
		{"SCMShipper", 5, true},
		{"SCMConsignee", 5, false},
		{"SCMFLOWShipper", 5, true},
		{"SCMFLOWConsignee", 5, false},
	}
	fmt.Println(roles) // Array of Object
	ur := UserRole{Priority: 10}
	fmt.Println(ur)
	for _, r := range roles {
		fmt.Println(r.Name)
	}
	//Get Parties function

}
