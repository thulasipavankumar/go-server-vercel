
package handler

import (
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h3>Hello World! <p>He's alive! Alive!</p></h3>")
  }
  
  
  func Main() {
	  r := mux.NewRouter()
	  r.HandleFunc("/", Greet).Methods("GET")
  
	  http.ListenAndServe(":8080", r)
  }