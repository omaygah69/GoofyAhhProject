package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
    // File Server :
    File_Server := http.FileServer(http.Dir("./Pages/"))

    //Router:
    router := http.NewServeMux()
    
    router.Handle("/", File_Server)
    router.HandleFunc("/getTodo", getTodo)
    router.HandleFunc("POST /post", postToDo)
    server := http.Server{
        Addr: ":6969",
        Handler: router,
    }
    fmt.Printf("Started, Listening To%v", server.Addr)
    log.Fatal(server.ListenAndServe())
}
