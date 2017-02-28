package tanks

import (
	"log"
	"net/http"
)

type command struct {
	Command string `json:"command"`
}

func handlePlayer(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", handlePlayer)

	log.Fatal(http.ListenAndServe("", nil))
}
