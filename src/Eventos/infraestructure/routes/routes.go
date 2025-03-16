package routes

import (
	"Eventos/src/Eventos/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	eventos := router.Group("/eventos") // Ruta principal para los eventos
	{
		eventos.POST("/", controllers.CreateEventoHandler) // Crear un nuevo evento
		//eventos.GET("/", controllers.ObtenerEventos) // Obtener todos los eventos
		// Si necesitas más rutas para actualizar, eliminar, etc. puedes agregarlas aquí
	}
}
