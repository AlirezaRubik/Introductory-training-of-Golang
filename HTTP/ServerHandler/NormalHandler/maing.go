package main

import(
"net/http"
"time"
)
func main(){
	Mux:=http.NewServeMux()
	Mux.Handle("/google",http.RedirectHandler("https://google.com",http.StatusOK))
	Mux.Handle("/rubikcomputer",http.RedirectHandler("https://www.rubikcomputer.ir",http.StatusOK))
	Server:=http.Server{
        Addr: ":5005",
		WriteTimeout: time.Second*10,
		ReadTimeout:time.Second *10,
		Handler:Mux,
	}
	Server.ListenAndServe()
}