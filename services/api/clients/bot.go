package clients

import (
	"github.com/lempiy/dora/shared/pb/bot"
	"github.com/lempiy/dora/shared/help"
	"google.golang.org/grpc/credentials"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"net"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
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

	conn, err := BlockingDial(ctx, "tcp", fmt.Sprintf("%s:%d", domain, port), creds, opts...)
	if err != nil {
		log.Printf("NewBotClient.BlockingDial %s", err)
		return nil
	}

	return &BotClient{
		domain:domain,
		port: port,
		conn: conn,
		BotServiceClient: bot.NewBotServiceClient(conn),
	}
}

// BlockingDial is a helper method to dial the given address, using optional TLS credentials,
// and blocking until the returned connection is ready. If the given credentials are nil, the
// connection will be insecure (plain-text).
func BlockingDial(ctx context.Context, network, address string, creds credentials.TransportCredentials, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// grpc.Dial doesn't provide any information on permanent connection errors (like
	// TLS handshake failures). So in order to provide good error messages, we need a
	// custom dialer that can provide that info. That means we manage the TLS handshake.
	result := make(chan interface{}, 1)

	writeResult := func(res interface{}) {
		// non-blocking write: we only need the first result
		select {
		case result <- res:
		default:
		}
	}

	dialer := func(address string, timeout time.Duration) (net.Conn, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		conn, err := (&net.Dialer{Cancel: ctx.Done()}).Dial(network, address)
		if err != nil {
			writeResult(err)
			return nil, err
		}
		if creds != nil {
			conn, _, err = creds.ClientHandshake(ctx, address, conn)
			if err != nil {
				writeResult(err)
				return nil, err
			}
		}
		return conn, nil
	}

	// Even with grpc.FailOnNonTempDialError, this call will usually timeout in
	// the face of TLS handshake errors. So we can't rely on grpc.WithBlock() to
	// know when we're done. So we run it in a goroutine and then use result
	// channel to either get the channel or fail-fast.
	go func() {
		opts = append(opts,
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithDialer(dialer),
			grpc.WithInsecure(), // we are handling TLS, so tell grpc not to
		)
		conn, err := grpc.DialContext(ctx, address, opts...)
		var res interface{}
		if err != nil {
			res = err
		} else {
			res = conn
		}
		writeResult(res)
	}()

	select {
	case res := <-result:
		if conn, ok := res.(*grpc.ClientConn); ok {
			return conn, nil
		}
		return nil, res.(error)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (b *BotClient) Close() {
	b.conn.Close()
}

