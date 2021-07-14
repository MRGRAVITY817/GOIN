package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	// struct field tag shows the name that we want to show
	// rather than actual struct field name
	// because struct field names are forced to be upper case
	// when it needed to be exported
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	// we can selectively show or not with "omit empty"
	Payload string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add blocks",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
	// These 3 lines are same as above
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

func main() {
	// handle route and callback function
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
