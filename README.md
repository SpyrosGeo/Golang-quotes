# Quotes

## A simple cli tool to fetch inspirational quotes from [ProgrammingQuotesApi](https://programming-quotes-api.herokuapp.com/quotes/random)

The imports that we use are :
```go
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
```
This is the struct represents the json data and was created using [json-to-go](https://mholt.github.io/json-to-go/)
```go

type Quote struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	En     string `json:"en"`
}

```
GET request :
```go
	res,err :=http.Get("https://programming-quotes-api.herokuapp.com/quotes/random")
	if err != nil{
		fmt.Println("No repsonse from request")	
	}
	defer res.Body.Close()
	//assign response Body to variable body
	body,err := ioutil.ReadAll(res.Body)//response Body
	//fmt.Println(string(body))
```
and after creating the struct to represent the json data I created a variable with the type of the Struct
```go

	var result Quote
	//unmarshall aka convert json string to object(type)
	if err := json.Unmarshal(body,&result); err!=nil{
		fmt.Println("Cannot unmarshal JSON")
	}
```
and finally print the quote 
```
	fmt.Println(result.En)
```
