// package controllers

// import (
// 	"Eventos/src/Eventos/application"
// 	"Eventos/src/Eventos/infraestructure"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"log"
// )

// func GetEventosHandler(c *gin.Context) {
// 	// Crear repositorio y caso de uso
// 	repo := infraestructure.NewMySQLEventoRepository()
// 	useCase := application.NewGetEventos(repo)

// 	// Obtener eventos
// 	eventos, err := useCase.Execute()
// 	if err != nil {
// 		log.Printf("Error al obtener eventos: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los eventos", "detalle": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, eventos)
// }


package controllers

import (
	"Eventos/src/Eventos/application"
	"Eventos/src/Eventos/infraestructure"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Estructura para el mensaje de FCM
type FCMMessage struct {
	To    string            `json:"to"`
	Data  map[string]string `json:"data"`
	Title string            `json:"title"`
	Body  string            `json:"body"`
}

const fcmURL = "https://fcm.googleapis.com/fcm/send"

func GetEventosHandler(c *gin.Context) {
	
	repo := infraestructure.NewMySQLEventoRepository()
	useCase := application.NewGetEventos(repo)

	eventos, err := useCase.Execute()
	if err != nil {
		log.Printf("Error al obtener eventos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los eventos", "detalle": err.Error()})
		return
	}

	// Verificar temperatura en los eventos
	for _, evento := range eventos {
		if evento.Valor < 15 {
			log.Println("🔴 Temperatura baja detectada:", evento.Valor)
			err := sendFCMNotification(evento.Valor)
			if err != nil {
				log.Printf("Error enviando notificación FCM: %v", err)
			}
		}
	}

	c.JSON(http.StatusOK, eventos)
}

// sendFCMNotification envía la notificación si la temperatura es baja
func sendFCMNotification(temp float64) error {
	
	fcmKey := os.Getenv("FCM_SERVER_KEY")
	if fcmKey == "" {
		return fmt.Errorf("FCM_SERVER_KEY no está configurada")
	}

	// Crear mensaje
	message := FCMMessage{
		To: "/topics/alertas", // Enviar a todos suscritos al topic 'alertas'
		Data: map[string]string{
			"alerta": "Temperatura baja",
			"valor":  fmt.Sprintf("%.2f°C", temp),
		},
		Title: " Alerta de Temperatura",
		Body:  fmt.Sprintf("La temperatura ha bajado a %.2f°C", temp),
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error serializando mensaje FCM: %v", err)
	}

	// Hacer petición a FCM
	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error creando request FCM: %v", err)
	}


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+fcmKey)

	// Enviar petición
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error enviando petición a FCM: %v", err)
	}
	defer resp.Body.Close()

	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("FCM respondió con código %d", resp.StatusCode)
	}

	log.Println(" Notificación enviada a FCM con éxito")
	return nil
}
