package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func hogeHandler(w http.ResponseWriter, r *http.Request)  {
	 io.WriteString(w, "Hello from a HandleFunc (>_<) #1!\n", )
}

type UserInfo struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Password string `json:"password`
}

func userInfoHandler(w http.ResponseWriter, r *http.Request)  {
	user := UserInfo{ Name: "Take", Email: "test@test.com", Age:24, Password:"password123" }

	res, err := json.Marshal(user)
	// w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Printf("%s\n", res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/hoge", hogeHandler)
  http.HandleFunc("/user", userInfoHandler)
	http.ListenAndServe(":8000", nil)
}
