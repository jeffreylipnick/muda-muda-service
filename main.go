package main // import "github.com/jeffreylipnick/muda-muda-service"

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jeffreylipnick/muda-muda-service/api"
	conf "github.com/jeffreylipnick/muda-muda-service/config"
)

func main() {
    config := conf.NewConfig()

	http.HandleFunc("/", api.HelloHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":" + config.Port, nil))
}
