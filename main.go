package main
import (
    "net/http"
)

func main(){
    router := http.NewServeMux()
    
    router.HandleFunc("/getshit", asd)
    
    server := http.Server{
        Addr: ":6969",
        Handler: router,
    }

    server.ListenAndServe()
}

func asd(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("bruh"))
}
