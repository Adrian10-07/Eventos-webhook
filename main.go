package main

import (
	"Eventos/src/Eventos/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080") // Servidor en el puerto 8080
}
