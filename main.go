package main

import (
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pssaravanan/gotodoapp/routers"
)

func elbCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func main(){
	r := gin.Default()
	r.GET("/todos", routers.GetAllTodo)
	r.Run("localhost:8080")
}