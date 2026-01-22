package main

import (
	"fmt"
	"myTodo/internal/router"
	"net/http"
)

func main(){
	r := router.SetUpRouter()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
