package clients

import (
	"github.com/lempiy/dora/shared/pb/prs"
	"github.com/lempiy/dora/shared/help"
	"google.golang.org/grpc/credentials"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"golang.org/x/net/context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

type ParserClient struct {
	domain string
	port int
	prs.ParseServiceClient
	conn *grpc.ClientConn
}

func NewParserClient()*ParserClient {
	domain := "dora-parser"
	port := 7000
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

	conn, err := help.BlockingDial(ctx, "tcp", fmt.Sprintf("%s:%d", domain, port), creds, opts...)
	if err != nil {
		log.Printf("NewParserClient.BlockingDial %s", err)
		return nil
	}

	return &ParserClient{
		domain:domain,
		port: port,
		conn: conn,
		ParseServiceClient: prs.NewParseServiceClient(conn),
	}
}

func (b *ParserClient) Close() {
	b.conn.Close()
}

