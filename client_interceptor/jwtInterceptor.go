package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationKey = "Authorization"
	invalidToken = "eyJhbGciOiJFUzI1NiIsImtpZCI6IjAxRU5TNUM5OUFNUFFaQVZUSkJGRjc4N0ZIXzE2MDM5NDEwNDkiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJpZ2RhdGEiLCJleHAiOjE2Mjc2MzY3MTQsImp0aSI6IjAxRkJSUlhWTllaM1I0SlhBMkI3WUFIWDdZIiwiaWF0IjoxNjI3NTUwMzE0LCJpc3MiOiJodHRwczovL2F1dGguZ2h0a2xhYi5jb20iLCJzdWIiOiIwMUVOUzVDOTlBTVBRWkFWVEpCRkY3ODdGSCIsInNjcCI6WyJpZ2RhdGE6cXVlcnkucmVhZCJdLCJjbGllbnRfaWQiOiIwMUVOUzVDOTlBTVBRWkFWVEpCRkY3ODdGSCJ9.tt8juVzcnazYRY07N_T4hyLXvQkYLI5dLywMcIKNx2UPfRS0wqQn5UzqN53vchy7nKElIlCYkFJAQaTFwcMv7A"
	validToken = "eyJhbGciOiJFUzI1NiIsImtpZCI6IjAxRU5TNUM5OUFNUFFaQVZUSkJGRjc4N0ZIXzE2MDM5NDEwNDkiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJpZ2RhdGEiLCJleHAiOjE2MjgyMjY2ODAsImp0aSI6IjAxRkNBQko2RjBLQUZKN0tFMUJTMzE3QkhYIiwiaWF0IjoxNjI4MTQwMjgwLCJpc3MiOiJodHRwczovL2F1dGguZ2h0a2xhYi5jb20iLCJzdWIiOiIwMUVOUzVDOTlBTVBRWkFWVEpCRkY3ODdGSCIsInNjcCI6WyJpZ2RhdGE6cXVlcnkucmVhZCJdLCJjbGllbnRfaWQiOiIwMUVOUzVDOTlBTVBRWkFWVEpCRkY3ODdGSCJ9.x-gNogbAQFgxpl29PqRqkGM2YOY2bYE_UE9GASVuMfM5U3ObeJKn5jttSM8zAGwA1wsjTg3OIhb_FCJx3osyIQ"
	invalidScope = "eyJhbGciOiJFUzI1NiIsImtpZCI6IjAxRU5TNUM5OUFNUFFaQVZUSkJGRjc4N0ZIXzE2MDM5NDEwNDkiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhdXRoIiwiZXhwIjoxNjI4MTM4MDE3LCJqdGkiOiIwMUZDN1EwRDdKOVY1WTBQM01WQTkyNzZDWSIsImlhdCI6MTYyODA1MTYxNywiaXNzIjoiaHR0cHM6Ly9hdXRoLmdodGtsYWIuY29tIiwic3ViIjoiMDFFTlM1Qzk5QU1QUVpBVlRKQkZGNzg3RkgiLCJjbGllbnRfaWQiOiIwMUVOUzVDOTlBTVBRWkFWVEpCRkY3ODdGSCJ9.WPzK3dECKaCS1B4p48-cQRQyczEcKeO2IdSk26hrMrRkkjmiwjZs64F6nwcohEq-6dXgej_CAmmS0r5H1iJwZg"
	ClientIP = "1234567"
	ClientKey = "client_ip"
)

type middleware struct {
}

type middlewareOpt func(middleware2 *middleware)

func NewMiddleware(opts... middlewareOpt) *middleware{
	m := &middleware{}
	for _, opt := range opts{
		opt(m)
	}
	return m
}

func (m *middleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor{
	return 	func(ctx context.Context, method string, req interface{},
		reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, AuthorizationKey, "Bearer " + validToken)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func (m *middleware)  StreamClientInterceptor() grpc.StreamClientInterceptor{
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
		method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error){
		newCtx := metadata.AppendToOutgoingContext(ctx, AuthorizationKey, "Bearer " + validToken)
		return streamer(newCtx, desc, cc, method, opts...)
	}
}
