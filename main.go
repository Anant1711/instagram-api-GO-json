package main

import (
	"fmt"
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
	n1:="Anant"
	n2:="Test"
	data:=Daata{
		UData{ID:"001", Name:n1, Email:"anantExample@gmail.com", Password:"1234@"},UData{ID:"002", Name:n2, Email:"TestExample@gmail.com", Password:"Test@"},
	}
	json.NewEncoder(w).Encode(data)

	keys, ok := r.URL.Query()["id"]
        if !ok || len(keys[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }
    	key := keys[0]
		fmt.Println(key)
		switch k:=key; k{
		case "1":
			fmt.Println("Password: 1234@")
			w.Write([]byte(n1))
		case "2":
			fmt.Println("Password: Test@")
			w.Write([]byte(n2))

		}
		
}

func postsData(w http.ResponseWriter, r *http.Request){
	now:=time.Now()
	id1,id2:=001,002
	c1:="Hey myslef Anant, this is my First Post"
	c2:="Hey myslef Test, this is my First Post"
	p_data:=postDaata{
		PData{ID:"001", Caption:c1, Image_url:"abcd.com",Stamp:now},PData{ID:"002", Caption:c2, Image_url:"efgh.com",Stamp:now},
	}
	json.NewEncoder(w).Encode(p_data)

	keys, ok := r.URL.Query()["id"]
        if !ok || len(keys[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }
    	key := keys[0]
		fmt.Println(key)
		switch k:=key; k{
		case "1":
			fmt.Println(id1)
			w.Write([]byte(c1))
		case "2":
			fmt.Println(id2)
			w.Write([]byte(c2))

		}
}

func homepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Instagram</h1>"))
}

func HandlereRequest(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users/", userData)
	http.HandleFunc("/posts/", postsData)
	log.Fatal(http.ListenAndServe(":3377", nil))
}
func main(){
	HandlereRequest()
	
}
