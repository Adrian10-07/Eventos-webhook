package controllers

import (
	"Eventos/src/Eventos/application"
	"Eventos/src/Eventos/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"database/sql"
)

func CrearEvento(c *gin.Context) {
	var eventoRequest struct {
		TipoSensor string  `json:"tipo_sensor"`
		Valor      float64 `json:"valor"`
	}

	if err := c.ShouldBindJSON(&eventoRequest); err != nil {
		c.JSON(400, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	db, err := sql.Open("mysql", "root:password@/dbname")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al conectar con la base de datos"})
		log.Println("Error al conectar con la base de datos:", err)
		return
	}
	defer db.Close()

	// Crear el repositorio para manejar los eventos
	repo := mysql.NuevoRepositorioEventos(db)

	// Crear el servicio de eventos e inyectar el repositorio
	servicio := application.NuevoServicioEvento(repo)

	// Llamar a la capa de aplicación para crear el evento
	evento, err := servicio.CrearEvento(eventoRequest.TipoSensor, eventoRequest.Valor)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al crear evento"})
		return
	}

	c.JSON(201, evento)
}
