package handlers

import (
	"fmt"

	w "videoChat/pkg/webrtc"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c,room.Peers)
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}

func RoomViewerWebsocket(c *websocket.Conn){

}

func roomViewerConn(c *websocket.Conn,p *w.Peers)
{

}

type websocketMessage struct{
	Event string `json:"event"`
	Data string `json:"data"`
}