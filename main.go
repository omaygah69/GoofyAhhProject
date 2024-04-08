package main

import (
	"fmt"
	"net/http"
)

func main(){
    router := http.NewServeMux()
    
    router.HandleFunc("/getshit", asd)
    router.HandleFunc("/test", getTitle)
    
    server := http.Server{
        Addr: ":6969",
        Handler: router,
    }
    fmt.Printf("Started, Listening To: %v", server.Addr)
    server.ListenAndServe()
}
