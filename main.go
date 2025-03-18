package main

import (
	"Eventos/src/Eventos/infraestructure/routes"
	"context"
	"fmt"
	"log"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)


func initFirebaseApp() (*firebase.App, error) {
	// Cambia "ruta/al/archivo/serviceAccountKey.json" a la ruta real donde está tu archivo JSON
	opt := option.WithCredentialsFile("/clave.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error inicializando la app de Firebase: %v", err)
	}
	return app, nil
}


func main() {
	app, err := initFirebaseApp()
	if err != nil {
		log.Fatalf("Error inicializando Firebase: %v", err)
	}
	fmt.Println("Firebase app inicializada correctamente:", app)
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080") // Servidor en el puerto 8080
}
