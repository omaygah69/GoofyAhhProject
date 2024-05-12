package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Todo struct{
    Title string
    Is_Completed bool
}

func init(){
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil{
        log.Fatal(err) 
    }
}

func getTodo(w http.ResponseWriter, r *http.Request){
    collection := client.Database("Lobotomy").Collection("ToDo")
    //Find all documents
    result, err := collection.Find(context.Background(), bson.D{})
    if err != nil{
        log.Fatal(err)
    }
    defer result.Close(context.Background())
	var Items []Todo
	for result.Next(context.Background()){
		var td Todo
		if err := result.Decode(&td);err != nil{
			log.Fatal(err)
		}
		Items = append(Items, td)
	}
	json.NewEncoder(w).Encode(Items)
}

func sampletest(w http.ResponseWriter){
    sampledata := Todo {
        Title: "Nigga",
    }
    jsonData, err := json.MarshalIndent(sampledata, "", " ")
	if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func postToDo(w http.ResponseWriter, r *http.Request){
    var td Todo
    err := json.NewDecoder(r.Body).Decode(&td)
    if err != nil{
        http.Error(w, "Error parsing json", http.StatusBadRequest)
        return
    }
    collection := client.Database("Lobotomy").Collection("ToDo")
    _, err = collection.InsertOne(context.Background(), td)
    if err != nil{
        http.Error(w, "Error Inserting Item", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Item Creation Completed"))
}

func deleteToDo(w http.ResponseWriter, r *http.Request){
    
}
