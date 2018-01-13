package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"bufio"
	"strings"
)

// Server Struct (Model)
type userPC struct {
	IPaddress string `json: "ipaddress"`
	KeyValue string `json: "keyvalue"`
	PropertyTag string `json: "propertytag"`
}

//init serverPC
var pcServer []userPC

func initPCServer() {
	file, err := os.Open("serverList.txt")
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
       eachLine := scanner.Text()
       userInfo := strings.Split(eachLine, " ")
			 pcServer = append(pcServer, userPC{IPaddress : userInfo[0], KeyValue : userInfo[1], PropertyTag : userInfo[2]})
       fmt.Println(eachLine)
       fmt.Println(userInfo[0])
       fmt.Println(userInfo[1])
       fmt.Println(userInfo[2])
   }

   if err := scanner.Err(); err != nil {
       log.Fatal(err)
   }
}

func extractPacket(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var Server userPC
	_ = json.NewDecoder(r.Body).Decode(&Server)
	fmt.Println(Server.IPaddress)
	fmt.Println(Server.KeyValue)
	fmt.Println(Server.PropertyTag)

	for index, item := range pcServer {
		if Server.KeyValue == item.KeyValue {
			pcServer = append(pcServer[:index], pcServer[index+1:]...)
			break
		}
	}
	pcServer = append(pcServer, Server)

	///Save to File
	f, err := os.OpenFile("serverList.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
			log.Fatal(err)
	}
	if _, err := f.Write([]byte(Server.IPaddress + " ")); err != nil {
			log.Fatal(err)
	}
	if _, err := f.Write([]byte(Server.KeyValue + " ")); err != nil {
			log.Fatal(err)
	}
	if _, err := f.Write([]byte(Server.PropertyTag + "\n")); err != nil {
			log.Fatal(err)
	}
	if err := f.Close(); err != nil {
			log.Fatal(err)
	}

	fmt.Println("Success")
}

//Get All ServerList
func getServerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pcServer)
}

func main() {
	r := mux.NewRouter()

	//Mocking the data
	initPCServer()
	r.HandleFunc("/serverlist", getServerList).Methods("GET")
	r.HandleFunc("/serverlist", extractPacket).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
