package main

import (
	"net/http"

	"github.com/cctechwiz/go-core-language-path/go_getting_started/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
