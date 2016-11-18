
package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Account struct {
    NAME string `json:"NAME"`
    ID string `json:"ID"`
    PHONE_NUMBER string `json:"PHONE_NUMBER"`
    TEAM_NAME string `json:"TEAM_NAME"`
}

var Info []Account

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w,r)
            return 
        } 

        if r.Method == "GET" {
            showData(w,r)
        } else if r.Method == "POST" {
            modifyData(w,r)
        } else if r.Method == "PUT" {
            insertData(w,r)   
        } else if r.Method == "DELETE" {
            deleteData(w,r) 
        } else {
            http.Error(w, "Invalid request method.", 405)
        }
        return 
    })

    http.ListenAndServe("localhost:8000", nil)
}

func showData(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    fmt.Fprintf(w, "Show Data\n")
    json.NewEncoder(w).Encode(Info) 
}

func insertData(w http.ResponseWriter, r *http.Request) {
    var req Account
    _ = json.NewDecoder(r.Body).Decode(&req)
    
    Info = append(Info, req)
    fmt.Fprintf(w, "Insert Success.")
    json.NewEncoder(w).Encode(r)
}

func deleteData(w http.ResponseWriter, r *http.Request) {    
    var req Account
    _ = json.NewDecoder(r.Body).Decode(&req)

    for index, item := range Info {

        if (item.ID == req.ID) {
            Info = append(Info[:index], Info[index+1:]...)
        }
    }

    fmt.Fprintf(w, "Delete Success.\nResult=")
    json.NewEncoder(w).Encode(Info)    
}

func modifyData(w http.ResponseWriter, r *http.Request) {
    var req Account
    _ = json.NewDecoder(r.Body).Decode(&req)
  
    for index, item := range Info {

        if (item.ID == req.ID) {
            Info[index] = req
        }
    }
    
    fmt.Fprintf(w, "Modify Success.\nResult=")
    json.NewEncoder(w).Encode(Info)

}