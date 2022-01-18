package main

import (
	"fmt"
//	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	En     string `json:"en"`
}

func main(){
	//TODO
	//get quote from api
	res,err :=http.Get("https://programming-quotes-api.herokuapp.com/quotes/random")
	if err != nil{
		fmt.Println("No repsonse from request")	
	}
	defer res.Body.Close()
	//assign response Body to variable body
	body,err := ioutil.ReadAll(res.Body)//response Body
	//fmt.Println(string(body))
// after creating the struct to represent the json data
// create a varablie of type of the struct in this case Quote
	var result Quote
	//unmarshall aka convert json string to object(type)
	if err := json.Unmarshal(body,&result); err!=nil{
		fmt.Println("Cannot unmarshal JSON")
	}

//	fmt.Println(PrettyPrint(result))
//	print the quote part from the struct
	fmt.Println(result.En)


}
//this is used to make the print pretty for the full json
func PrettyPrint(i interface {})string {
	s,_:=json.MarshalIndent(i,"","\t")
	return  string(s)
}
