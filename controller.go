package main
import (
    "net/http"
    "encoding/json"
)

type Data struct{
    Title string
}

func asd(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("bruh"))
}

func getTitle(w http.ResponseWriter, r *http.Request){
    sampledata := Data{
        Title: "Nigga",
    }

    jsonData, err := json.MarshalIndent(sampledata, "", " ")
    if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")

    w.Write(jsonData)
}

