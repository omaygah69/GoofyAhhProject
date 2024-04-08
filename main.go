package main

import (
	"encoding/json"
	"net/http"
)
type Data struct{
    Title string
}

func main(){
    router := http.NewServeMux()
    
    router.HandleFunc("/getshit", asd)
    router.HandleFunc("/test", getTitle)
    
    server := http.Server{
        Addr: ":6969",
        Handler: router,
    }
    server.ListenAndServe()
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
