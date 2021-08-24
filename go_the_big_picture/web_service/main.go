package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"] //Can specify more than one name query parameter
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string {"name": name}
		enc := json.NewEncoder(w) //Anything encoded with this encoder is writen to the writer
		enc.Encode(m)
	})

	err := http.ListenAndServe(":3000", nil) //nil specifies to use the built-in handler
	if err != nil {
		log.Fatal(err)
	}
}
