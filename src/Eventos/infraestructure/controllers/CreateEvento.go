package controllers

import (
	"Eventos/src/Eventos/application"
	"Eventos/src/Eventos/domain"
	"Eventos/src/Eventos/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Crear evento
func CreateEventoHandler(c *gin.Context) {
	var evento domain.Evento

	// Decodificar JSON
	if err := c.ShouldBindJSON(&evento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Crear repositorio y caso de uso
	repo := infraestructure.NewMySQLEventoRepository()
	useCase := application.NewCreateEvento(repo)

	// Guardar evento
	if err := useCase.Execute(evento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el evento"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Evento creado con éxito"})
}
