package requestsystem

import (
	"encoding/json"
	"fmt"
	Rsys "information/ResponseSys"
	"net/http"
	"strconv"
)

type Request struct {
}
type Person struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

var Database = make(map[string]*Person)

func (req *Request) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0 {
		req.ShowAllUsers(w, r)
	}
	if r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0 {
		req.SearchUser(w, r)
	}
	if r.Method == http.MethodPost {
		req.CrateNewUser(w, r)
	}
}
func (req *Request) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	Cv := new(Rsys.ResponseHandler)
	for index, item := range Database {
		Cv.MakeResponseHander(http.StatusOK, fmt.Sprintf("The index is: %s The Item is: %v", index, item), nil, w)
	}
}
func (req *Request) CrateNewUser(w http.ResponseWriter, r *http.Request) {
	Cv := new(Rsys.ResponseHandler)
	p := new(Person)
	apikey := r.Header.Get("api-key")

	if apikey == "alireza" {
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			Cv.MakeResponseHander(http.StatusBadRequest, "Who are you?", nil, w)
			return
		} else {
			Cv.MakeResponseHander(http.StatusOK, fmt.Sprintf("User %v \n", p), nil, w)
			Database[strconv.Itoa(p.Id)]=p
			return
		}
	} else {
		Cv.MakeResponseHander(http.StatusBadRequest, "Who are you?", nil, w)
		return
	}
}
func (req *Request) SearchUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Cv := new(Rsys.ResponseHandler)
	Result, ok := Database[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		Cv.MakeResponseHander(http.StatusNotFound, "Your User NotFound", nil, w)
		return
	} else {
		Cv.MakeResponseHander(http.StatusOK, fmt.Sprintf("User %v \n", Result), nil, w)
		return
	}

}
