package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json: "nombre"`
	Apellido string `json: "nombre"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	router.POST("/saludo", func(ctx *gin.Context) {

		var persona Persona
		if err := ctx.BindJSON(&persona); err != nil {
			saludo := fmt.Sprintf("Hola %v %v", persona.Nombre, persona.Apellido)
			ctx.String(200, saludo)
		}

	})
	router.Run()
}
