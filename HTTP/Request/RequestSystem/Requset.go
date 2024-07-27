package requestsystem

import (
	"fmt"
	"net/http"
)

type Request struct {
}

func (req *Request) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && len(r.URL.Query().Get("id")) == 0 {
		req.ShowAllUsers(w,r)
	}
	if r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0 {
        req.SearchUser(w,r)
	}
	if r.Method == http.MethodPost {
        req.CrateNewUser(w,r)
	}
}
func (req *Request) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	apikey := r.Header.Get("api-key")
	if apikey == "alireza" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hi alireza_rubik welcome"))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("who are you?"))
	}
}
func (req *Request)CrateNewUser(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user has been created"))
}
func (req *Request)SearchUser(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	id:=r.URL.Query().Get("id")
	fmt.Fprintf(w, "Search: %s", id)

}