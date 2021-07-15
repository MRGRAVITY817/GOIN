package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MRGRAVITY817/goin/blockchain"
	"github.com/MRGRAVITY817/goin/utils"
	"github.com/gorilla/mux"
)

var port string

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	// struct field tag shows the name that we want to show
	// rather than actual struct field name
	// because struct field names are forced to be upper case
	// when it needed to be exported
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	// we can selectively show or not with "omit empty"
	Payload string `json:"payload,omitempty"`
}

func (u urlDescription) String() string {
	return "Hello I'm the url Description"
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add blocks",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See a block",
		},
	}
	json.NewEncoder(rw).Encode(data)
	// These 3 lines are same as above
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var a addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&a))
		blockchain.GetBlockchain().AddBlock(a.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

// Middleware is a function before all api endpoint.
// Which does the dirty stuff prehandedly.
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	// we need individual mux for each explorer and rest package
	// or else it will tie them into same router, and it will
	// eventually cause routing crash.
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Api Server: http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
