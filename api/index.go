
package handler

import (
	"fmt"
    "net/http"
	"strconv"
	"github.com/gorilla/mux"
)


func ArthematicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a,_ := strconv.Atoi(vars["a"])
	b,_ := strconv.Atoi(vars["b"])
	fmt.Fprintf(w,"Addition: "+strconv.Itoa(Add(a,b))+"\nSubtraction: "+strconv.Itoa(Subract(a,b))+"\nMutiply: "+strconv.Itoa(Multiply(a,b)))
}
func UniversalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`Hello World! \nUse "<URL>/<number>/<number>" to get the add,subtract and mutiplication of the numbers `)
}


func Add(a int,b int)int {
	return a+b
}
func Subract(a int,b int)int{
	return a-b
}
func Multiply(a int,b int)int{
	return a*b
}
func main() {
	  r := mux.NewRouter()
	  r.HandleFunc("/",UniversalHandler).Methods("GET")
	  r.HandleFunc("/{a}/{b}",ArthematicHandler).Methods("GET") 
	  http.ListenAndServe(":8080",r)
}