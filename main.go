package main
import (
    "net/http"
)

func main(){
    router := http.NewServeMux()
    
    router.HandleFunc("/getshit", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Geniggas"))
    })

    http.ListenAndServe(":6969", router)
}
