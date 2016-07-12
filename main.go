package main 

import (
	"fmt"
	"net/http"
	"log"
)

func main () {
	http.HandleFunc("/sendmail", handleSendMail)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":80", nil)
}


func handleSendMail(w http.ResponseWriter, r *http.Request) {
	log.Println("request to sendmail")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "mailer")
}




