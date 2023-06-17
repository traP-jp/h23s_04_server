package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// WebSocket接続をアップグレードします
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// 接続が確立されたことをクライアントに通知します
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Connected to server")); err != nil {
		log.Println("Write error:", err)
		return
	}

	// メッセージを読み取り、処理するループ
	for {
		// クライアントからメッセージを受信します
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// 受信したメッセージを処理します（ここでは単純にログに出力します）
		log.Println("Received message:", string(message))

		// クライアントにメッセージを送信します
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Message received")); err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	// 接続を閉じます
	conn.Close()
}

func main() {
	http.HandleFunc("/websocket", handleWebSocket)
	fmt.Print("Hello")

	// サーバーを起動します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
