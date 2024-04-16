package main

import (
	"fmt"
	"net/http"
    "log"
)

func main(){

    // File Server :
    File_Server := http.FileServer(http.Dir("./Pages/"))

    //Router:
    router := http.NewServeMux()
    
    router.Handle("/", File_Server)
    router.HandleFunc("/bruh", asd)
    router.HandleFunc("/test/getTodo", getTodo)
    router.HandleFunc("/getNote", getNotes)
    
    server := http.Server{
        Addr: ":6969",
        Handler: router,
    }

    fmt.Printf("Started, Listening To%v", server.Addr)
    log.Fatal(server.ListenAndServe())
}
