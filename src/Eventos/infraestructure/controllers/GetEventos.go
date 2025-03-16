package controllers

import (
	"Eventos/src/Eventos/application"
	"Eventos/src/Eventos/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func GetEventosHandler(c *gin.Context) {
	// Crear repositorio y caso de uso
	repo := infraestructure.NewMySQLEventoRepository()
	useCase := application.NewGetEventos(repo)

	// Obtener eventos
	eventos, err := useCase.Execute()
	if err != nil {
		log.Printf("Error al obtener eventos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los eventos", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusOK, eventos)
}
