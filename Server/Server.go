package main

import (
	proto "ChitChatty/grpc"
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Server struct {
	proto.UnimplementedRecieveMessageServer // Necessary
	name string
	port int
}

var port = flag.Int("port", 0, "server port number")

var clients map[int64](flag.Value);

func main() {
	
	clients = make(map[int64]flag.Value)
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name: "serverName",
		port: *port,
	}

	// Start the server
	go startServer(server)

	// Keep the server running until it is manually quit
	for {

	}
}

func startServer(server *Server) {

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterRecieveMessageServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (c *Server) RecieveMessage(ctx context.Context, in *proto.ClientSendMessage) (*proto.ServerBroadcastMessage, error) {
	log.Printf("Client with ID %d posted: %s", in.ClientId, in.Msg);
	
	
	
	// Change out the above four lines with this:
	return nil;	
}

func (c *Server) Broadcast(CLients []int64, msg string){
	//The underscore omits index
	for _, curClient :=	range CLients {
		//function clientside for waiting for msg here?
		//Change waitForTimeRequest 
		log.Printf("Client: %d msg: %s", curClient, msg)
	}
}

func (c *Server) Listener(){
	//msg string = "";
	
	/* switch msg {
		case "new connection":
			RecieveConnection()
		case "End connection":
			endConnection()
		default:
			c.Broadcast(Clients, msg);
	} */
} 

func connectToClient() (client proto.RecieveMessageClient, err error, port string, portID int, portHelper string) {
	// Dial the client at the specified port.
	var clientPort = flag.Int(port, portID, portHelper);
	flag.Parse()
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*clientPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *clientPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *clientPort)
	}
	return proto.NewRecieveMessageClient(conn), nil, "", 0, "";
}

/*


func (c *Server) listener(.....) {
	msg = recieve from client
	switch(msg) {
		case "new connection":
			RecieveConnection()
		case "end connection":
			endConnection()
		default:
			Broadcast(msg)
	}
}	

func (c *Server) RecieveConnection(...) {
	add connection to array
	Broadcast (client <id> joined the chat)
}

func (c *Server) endConnection(...) {
	remove connection from array
	Broadcast (client <id> left the chat)
}

*/