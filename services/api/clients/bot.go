package clients

import (
	"github.com/lempiy/dora/shared/pb/bot"
	"github.com/lempiy/dora/shared/help"
	"google.golang.org/grpc/credentials"
	"log"
	"google.golang.org/grpc"
	"fmt"
)

type BotClient struct {
	domain string
	port int
	bot.BotServiceClient
	conn *grpc.ClientConn
}

func NewBotClient()*BotClient {
	domain := "dora-bot"
	port := 6000
	dir := help.GetCurrentDir()
	creds, err := credentials.NewClientTLSFromFile(dir+"/cert/client.crt", "")
	if err != nil {
		log.Fatal("Cannot read credentials from file: ", err)
	}
	log.Printf("Dialing %s:%d...", domain, port)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithBlock()}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", domain, port), opts...)
	if err != nil {
		log.Fatal("Cannot dial TCP dora-bot", err)
	}
	client := bot.NewBotServiceClient(conn)
	bc := BotClient{
		domain:domain,
		port: port,
		BotServiceClient: client,
		conn: conn,
	}
	return &bc
}

func (b *BotClient) Close() {
	b.conn.Close()
}

