package main

import (
	"flag"
	"golang-study/init/cmd"
)

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", helloWorld)

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Error Occured")
// 		panic(err)
// 	}
// }

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("hello World")
// }

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
	// cmd.NewCmd("./config.toml")
}
