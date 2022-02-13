package main // import "github.com/jeffreylipnick/muda-muda-service"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jeffreylipnick/muda-muda-service/api"
)

func main() {



	http.HandleFunc("/", api.HelloHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
