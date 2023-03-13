package main

// to check:
// https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "gRPCStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
)

const (
	serverAddr = "www.opentv.com:8080"
	certFile   = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/client-cert.pem"
	//certFile = "/home/chris/Documents/Dev/go/src/gRPCStreaming/cert/certs/cli.pem"
	keyFile = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/client-key.pem"
	//keyFile    = "/home/chris/Documents/Dev/go/src/gRPCStreaming/cert/private/cli.key"
	caCertFile = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/server-cert.pem"
)

type customCredential struct {
	token string
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.token,
	}, nil
}

//func (c customCredential) RequireTransportSecurity() bool {
//	return false
//}

func main() {

	conSecured("worker 1")
	//conInSecured("worker 1")
	//go con("worker 2")

	fmt.Println("2 listener launched....")
	for {

	}
}

func authMethods() map[string]bool {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"

	return map[string]bool{
		"/event.grpc.Event/GetEvent": true,
	}
}

func conSecured(name string) {

	// load trusted CA certificate file
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Failed to read CA certificate file: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("Failed to append CA certificate to pool")
	}

	// use TLS package and load the certificate and key file
	clientCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load server certificate key pary: %v", err)
	}

	c := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{clientCert},
	}
	creds := credentials.NewTLS(c)

	// creds := insecure.NewCredentials()

	// load client certificate and key files
	// creds, err := credentials.NewClientTLSFromFile(certFile, keyFile)
	//if err != nil {
	//	log.Fatalf("Failed to generate credentials: %v", err)
	//}
	// create gRPC client with TLS configuraiton
	//customCred := customCredential{token: "your-auth-token"}

	interceptor, err := NewAuthInterceptor(serviceToken.authMethods, serviceToken.token())
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
		//grpc.WithBlock(),
		//grpc.WithPerRPCCredentials(customCred),
		//grpc.WithAuthority("www.opentv.com"),
	}

	conn, err := grpc.Dial(serverAddr, dialOpts...)
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	defer conn.Close()

	client := pb.NewEventClient(conn)

	e := &pb.Request{
		Term:      "test",
		MaxResult: 2,
	}
	stream, err := client.GetEvent(context.Background(), e)

	if err != nil {
		log.Fatal("unable to connect to the GetEvent service", err)
	}

	done := make(chan bool)

	go func() {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("%v.ListEvent(_) = _, %v", client, err)
			}
			log.Println("worker: ", name, "event: ", event)
		}
	}()

	<-done
	log.Printf("finished")

}

func conInSecured(name string) {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:8080", opts...)
	//conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	defer conn.Close()

	client := pb.NewEventClient(conn)

	e := &pb.Request{
		Term:      "test",
		MaxResult: 2,
	}

	stream, err := client.GetEvent(context.Background(), e)
	if err != nil {
		log.Fatal("unable to connect to the GetEvent service", err)
	}

	done := make(chan bool)

	go func() {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("%v.ListEvent(_) = _, %v", client, err)
			}
			log.Println("worker: ", name, "event: ", event)
		}
	}()

	<-done
	log.Printf("finished")

}
