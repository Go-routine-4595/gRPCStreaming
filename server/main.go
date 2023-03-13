package main

import (
	"crypto/tls"
	"crypto/x509"
	pb "gRPCStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
)

const (
	//serverAddr = "www.opentv.com:8080"
	serverAddr = "0.0.0.0:8080"
	certFile   = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/server-cert.pem"
	keyFile    = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/server-key.pem"
	caCertFile = "/home/chris/Documents/Dev/go/src/gRPCStreaming/certAlt/ca-cert.pem"
)

func (e *EventService) mustEmbedUnimplementedEventServer() {

}

func main() {

	//servInSecured()
	servSecured()
}

func servSecured() {

	var jwtManager *JWTManager

	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to setup tcp listner: %v", err)
	}

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
	serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load server certificate key pary: %v", err)
	}

	// Create a tls object/config
	tlsConf := &tls.Config{
		Certificates:       []tls.Certificate{serverCert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		RootCAs:            caCertPool,
		ClientCAs:          caCertPool,
		InsecureSkipVerify: true,
	}

	// load server certificate and key file
	// creds, err := credentials.NewClientTLSFromFile(certFile, keyFile)
	// if err != nil {
	//	log.Fatalf("Failed to generate credentials: %v", err)
	//}

	creds := credentials.NewTLS(tlsConf)

	// create an interceptor object
	interceptor := NewAuthInterceptor(jwtManager, accessibleRoles())
	// create gRPC server with TLS configuration
	serverOpts := []grpc.ServerOption{
		grpc.Creds(creds),
		//grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		//grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	s := grpc.NewServer(serverOpts...)

	pb.RegisterEventServer(s, new(EventService))

	log.Print("Starting RPC server on port 8080...")

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func servInSecured() {

	list, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to setup tcp listner: %v", err)
	}

	serverOpts := []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
	}

	s := grpc.NewServer(serverOpts...)
	pb.RegisterEventServer(s, new(EventService))

	log.Print("Starting RPC server on port 8080...")

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
