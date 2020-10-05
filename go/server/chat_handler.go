package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     canConnectWS,
	}
)

func chatWebSocket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("uid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro no paramêtro UID: %s", err)
		return
	}

	fmt.Println(id)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro na conexão com o WebSocket: %s", err)
		return
	}

	handleReadMessage(conn)
}

func handleReadMessage(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.CloseMessage,
				[]byte("Não foi possível ler sua mensagem, efetue o login novamente."))

			return
		}
		fmt.Println("Message: ", string(msg))
	}
}

func canConnectWS(r *http.Request) bool {
	return true
}
