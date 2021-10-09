package main

import (
	//"fmt"
	"net/http" 	//Package http provides HTTP client and server implementations.
	"log"
	"encoding/json"		//Package json implements encoding and decoding of JSON
	"time"

)


type UData struct{
	ID string `json:"ID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

type PData struct{
	
	ID string `json:"ID"`
	Caption string `json:"caption"`
	Image_url string `json:"url"`
	Stamp time.Time `json:"time"`

}

type Daata []UData 		//array for user data
type postDaata []PData  //array for Post Data

func userData(w http.ResponseWriter, r *http.Request){
	data:=Daata{
		UData{ID:"001", Name:"anant", Email:"anantExample@gmail.com", Password:"1234@"},
	}
	
	json.NewEncoder(w).Encode(data)
}

func postsData(w http.ResponseWriter, r *http.Request){
	now:=time.Now()
	p_data:=postDaata{
		PData{ID:"001", Caption:"Hey, its my First Post", Image_url:"abcd",Stamp:now},
	}
	json.NewEncoder(w).Encode(p_data)
}


func homepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Instagram</h1>"))
}


func HandlereRequest(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users", userData)
	http.HandleFunc("/posts", postsData)
	
	log.Fatal(http.ListenAndServe(":3377", nil))
}


func main(){
	HandlereRequest()
}
