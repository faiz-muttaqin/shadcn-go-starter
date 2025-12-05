package ws

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	Clients = make(map[string]*websocket.Conn)
	Mutex   = sync.Mutex{}
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request, clientID string, db *gorm.DB) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Println(err)
		return
	}

	Mutex.Lock()
	Clients[clientID] = conn
	Mutex.Unlock()

	go HandleMessages(clientID, conn, db)
}

func HandleMessages(clientID string, conn *websocket.Conn, db *gorm.DB) {
	defer func() {
		Mutex.Lock()
		delete(Clients, clientID)
		Mutex.Unlock()

		conn.Close()

		// Start the async check for reconnection
		go checkForReconnection(clientID, db)
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logrus.Println(err)
			return
		}

		logrus.Printf("Received message from %s", p)

		// Handle the message (you can implement your own message format)
		HandleMessage(messageType, p)
	}
}

func HandleMessage(messageType int, message []byte) {
	// Example: Assume messages have the format "recipientID:message"
	parts := strings.SplitN(string(message), ":", 2)
	if len(parts) != 2 {
		logrus.Println("Invalid message format:", string(message))
		return
	}

	recipientID := parts[0]
	actualMessage := parts[1]

	// Broadcast the message to the intended recipient
	SendMessageToRecipient(messageType, actualMessage, recipientID)
}

// SendMessageToRecipient(1, "the message", "email") //1 is text message, 2 is binary
func SendMessageToRecipient(messageType int, message, recipientID string) {
	Mutex.Lock()
	defer Mutex.Unlock()

	if clientConn, ok := Clients[recipientID]; ok {
		if clientConn != nil {
			err := clientConn.WriteMessage(messageType, []byte(message))
			if err != nil {
				logrus.Println(err)
			}
		}
	}
}

func CloseWebsocketConnection(clientID string) {
	Mutex.Lock()
	if clientConn, ok := Clients[clientID]; ok {
		if clientConn != nil {
			clientConn.Close()
		}
		delete(Clients, clientID)
	}
	Mutex.Unlock()
}

func checkForReconnection(clientID string, db *gorm.DB) {
	disconectionTimeStr := os.Getenv("MAX_DISCONECTION_TIME_S") // Parse the disconection time as an integer
	disconectionExpiredSeconds, err := strconv.Atoi(disconectionTimeStr)
	if err != nil {
		disconectionExpiredSeconds = 30
	}
	// Set a timeout of 20 seconds
	timer := time.After(time.Duration(disconectionExpiredSeconds) * time.Second)

	select {
	case <-timer:
		Mutex.Lock()
		_, connected := Clients[clientID]
		Mutex.Unlock()

		if !connected {
			logrus.Printf("Client %s did not reconnect within %s seconds", clientID, disconectionTimeStr)
			updates := map[string]interface{}{
				"Session":        "",
				"SessionExpired": 0,
			}

			if err := db.Model(&model.Admin{}).Where("email = ?", clientID).Updates(updates).Error; err != nil {
				return
			}
		}
	}
}

func IsClientConnected(clientID string) bool {
	Mutex.Lock()
	defer Mutex.Unlock()

	_, connected := Clients[clientID]
	return connected
}

// BroadcastMessage sends a message to all connected clients.
// Text Message: 1 (or websocket.TextMessage)
// Binary Message: 2 (or websocket.BinaryMessage)
func BroadcastMessage(messageType int, message string) {
	Mutex.Lock()
	defer Mutex.Unlock()

	for clientID, clientConn := range Clients {
		if clientConn != nil && !strings.HasPrefix(clientID, "broker_mqtt") {
			err := clientConn.WriteMessage(messageType, []byte(message))
			if err != nil {
				logrus.Error(err)
				log.Printf("Error sending message to %s: %v", clientID, err)
			}
		}
	}
}
