package main
import (
    "net/http"
    "encoding/json"
)

type Todo struct{
    Title string
    Is_Completed bool
}

type Note struct{
    Title string 
    Text string 
}

func asd(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("bruh"))
}

func getTodo(w http.ResponseWriter, r *http.Request){
    sampledata := Todo{
        Title: "Nigga",
    }

    jsonData, err := json.MarshalIndent(sampledata, "", " ")
    if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")

    w.Write(jsonData)
}

func getNotes(w http.ResponseWriter, r *http.Request){
    sampleData := Note{
        Title: "Note_Number 1",
    }
    jsonData, err := json.MarshalIndent(sampleData, "", " ")
    if err != nil{
        http.Error(w, "Error", http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")

    w.Write(jsonData)
}

func createNote(w http.ResponseWriter, r *http.Request){
    
}

