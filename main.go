package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CreatedAt time.Time
}

type goHandler struct{}

func (f *goHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler (w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "World"
	}
	fmt.Fprintf(w, "USERs %s", name)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/user", barHandler)
	mux.Handle("/go", &goHandler{})
	fmt.Print("http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}