package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"

	"github.com/cryptomkt/cryptomkt-go/conn"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func sendJoin(c *gosocketio.Client, method string, data interface{}) {
	log.Printf("Emiting %s", method)
	resp, err := c.Ack(method, data, 120*time.Second)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s", resp)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("worker.cryptomkt.com", 443, true),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("/message", func(h *gosocketio.Channel, args Message) {
		log.Println("--- Got chat message: ", args)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	var keysfile = "../keys.txt"

	file, err := os.Open(keysfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	apiKey := scanner.Text()
	scanner.Scan()
	apiSecret := scanner.Text()
	client := conn.NewClient(apiKey, apiSecret)

	respAsReference, _ := client.SocketAuthInfo()

	respMap := make(map[string]string)

	respMap["uid"] = strconv.Itoa(respAsReference.Data.Uid)
	respMap["socid"] = respAsReference.Data.Socid

	sendJoin(c, "user-auth", respMap)

	log.Println(" [x] Complete")

}
