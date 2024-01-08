package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// LoggingUnaryInterceptor logs the unary request
func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	// logging
	log.Printf(`
================== gRPC Unary Call ===================
Method: %v
Duration: %v
Error: %v
Response: %v
======================================================
`,
		info.FullMethod,
		time.Since(start),
		err,
		h)

	return h, err
}

// LoggingStreamInterceptor logs the stream request
func LoggingStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()

	err := handler(srv, ss)

	// logging
	log.Printf(`
================== gRPC Streaming Call ===================
Method: %v
Duration: %v
Error: %v
Response: %v
======================================================
`,
		info.FullMethod,
		time.Since(start),
		err,
		ss)

	return err
}
