package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	RouterRegister()

	log.Fatal(http.ListenAndServe(":80", nil))
}
