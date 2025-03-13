package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "grpc-live-chat/grpc-live-chat/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Masukkan nama user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan username: ")
	username, _ := reader.ReadString('\n')
	username = username[:len(username)-1]

	// Receive goroutine
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Gagal menerima: %v", err)
			}
			fmt.Printf("[%s]: %s\n", in.Sender, in.Message)
		}
	}()

	// Kirim pesan
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		if err := stream.Send(&pb.ChatMessage{
			Sender:    username,
			Message:   text,
			Timestamp: time.Now().Unix(),
		}); err != nil {
			log.Fatalf("Gagal kirim: %v", err)
		}
	}
}
