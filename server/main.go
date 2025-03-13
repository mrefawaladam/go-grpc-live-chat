package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	pb "grpc-live-chat/grpc-live-chat/proto"

	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients map[string]pb.ChatService_ChatStreamServer
}

func NewChatServer() *chatServer {
	return &chatServer{
		clients: make(map[string]pb.ChatService_ChatStreamServer),
	}
}

func (s *chatServer) ChatStream(stream pb.ChatService_ChatStreamServer) error {
	var username string

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			s.mu.Lock()
			delete(s.clients, username)
			s.mu.Unlock()
			log.Printf("User %s keluar.", username)
			return nil
		}
		if err != nil {
			return err
		}

		username = msg.Sender

		// Tambah ke daftar client
		s.mu.Lock()
		s.clients[username] = stream
		s.mu.Unlock()

		log.Printf("Pesan dari %s: %s", msg.Sender, msg.Message)

		// Broadcast ke semua user
		s.broadcast(msg)
	}
}

func (s *chatServer) broadcast(msg *pb.ChatMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for user, client := range s.clients {
		if user == msg.Sender {
			continue // Skip sender
		}

		if err := client.Send(&pb.ChatMessage{
			Sender:    msg.Sender,
			Message:   msg.Message,
			Timestamp: time.Now().Unix(),
		}); err != nil {
			log.Printf("Gagal kirim ke %s: %v", user, err)
		}
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, NewChatServer())
	fmt.Println("Chat server running on :50051")
	s.Serve(lis)
}
