package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Payload struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    http.HandleFunc("/post", postHandler)

    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    var payload Payload
    err = json.Unmarshal(body, &payload)
    if err != nil {
        http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
        return
    }

    fmt.Printf("Received payload: %+v\n", payload)
}