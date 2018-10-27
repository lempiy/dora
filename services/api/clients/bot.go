package clients

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/lempiy/dora/shared/help"
	"github.com/lempiy/dora/shared/pb/bot"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
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
	var opts []grpc.DialOption
	ctx := context.Background()
	dir := help.GetCurrentDir()
	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(dir+"/cert/client.crt", dir+"/cert/client.key")
	if err != nil {
		log.Println(err)
		return nil
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(dir+"/cert/ca.crt")
	if err != nil {
		log.Println(err)
		return nil
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Println(err)
		return nil
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   domain, // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
	if err != nil {
		log.Fatal("Cannot read credentials from file: ", err)
	}

	conn, err := help.Connect(ctx, fmt.Sprintf("%s:%d", domain, port), creds, opts...)
	if err != nil {
		log.Printf("NewBotClient. Cannot connect to server: `%s`", err)
		return nil
	}

	return &BotClient{
		domain:domain,
		port: port,
		conn: conn,
		BotServiceClient: bot.NewBotServiceClient(conn),
	}
}

func (b *BotClient) Close() {
	b.conn.Close()
}

