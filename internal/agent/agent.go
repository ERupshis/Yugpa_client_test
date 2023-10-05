package agent

import (
	"context"
	"fmt"
	"net"
	"yugpa_test/internal/messages"
)

type Agent struct {
	serverAddr string
}

func Create(serverAddr string) Agent {
	return Agent{
		serverAddr: serverAddr,
	}
}

func (a *Agent) Serve(ctx context.Context, connCount int) {

}

func (a *Agent) makeRequestToServer() {
	// Connect to the server
	conn, err := net.Dial("tcp", a.serverAddr)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	request := messages.Request{
		Path: "C:\\kafka\\",
	}

	// Serialize the request
	requestBytes, err := request.Serialize()
	if err != nil {
		fmt.Printf("Error serializing request: %v\n", err)
		return
	}

	// Send the request to the server
	_, err = conn.Write(requestBytes)
	if err != nil {
		fmt.Printf("Error sending request to the server: %v\n", err)
		return
	}

	// Receive and decode the response
	var response messages.Response
	if err = response.Deserialize(conn); err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("%v\n", response)
}
