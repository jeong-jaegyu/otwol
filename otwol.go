package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
	EOF         = "\x04"
	TOKEN       = "AAAA-1234"
	TOKEN_AUTH  = "%OTWOL_password"
)

var (
	Warn *log.Logger
	Info *log.Logger
	Err  *log.Logger
)

func init() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Err = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

// TODO: cleaner buffers, currently these are going to be shit
func processClient(connection net.Conn) {
	buffer := make([]byte, 2048)
	mLen, err := connection.Read(buffer)

	if err != nil {
		Err.Fatal(err.Error())
	}

	// check for auth_req. if it's not an auth req we close the connection.

	if !bytes.Equal(buffer[:mLen], []byte("Auth_Req")) {
		Info.Printf(
			"Client at [%s] tried to connect without an Auth_Req",
			connection.RemoteAddr(),
		)
		connection.Close()
		return
	}
	// cleanup this
	buffer = make([]byte, 2048)

	// send the Auth_Req_ACK packet
	Info.Printf(
		"Sending Auth_Req_ACK to client at [%s]",
		connection.RemoteAddr(),
	)
	connection.Write(
		[]byte("Auth_Req_ACK"),
	)

	// resonse await
	mLen, err = connection.Read(buffer)

	if err != nil {
		Err.Fatal(err.Error())
	}

	// Token Identification + Auth here
	// This is hardcoded dogshit, replace before using

	split := bytes.Index(buffer, []byte("|"))
	Token := buffer[:split]
	Token_Auth := buffer[split:]
	fmt.Printf(
		"1st: %s, 2nd: %s",
		Token, Token_Auth,
	)
	// if !bytes.Equal(buffer, []byte(TOKEN)) {
	// }

	// End of this section of hardcoded dogshit, replace before using.
	connection.Close()
}

// fmt.Println("Received: ", string(buffer[:mLen]))
// _, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
