package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, GitOps!")
    })

    http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, User!")
    })

    http.HandleFunc("/friend", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, Friend!")
    })

    http.HandleFunc("/sauce", func(w http.ResponseWriter, r *http.Request) {
        fileContent, err := ioutil.ReadFile("./application-secret/sauce")
        var secretSauce = "unknown"
        if err == nil {
            secretSauce = string(fileContent)
        }
        fmt.Fprint(w, "The secret sauce is: ", secretSauce)
    })

    http.ListenAndServe(":8080", nil)
}
