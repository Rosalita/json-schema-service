package main

import (
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/validation", validationHandler)

    http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "upload handler")
}

func downloadHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "download handler")
}

func validationHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "validation handler")
}
