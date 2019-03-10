package main

import (
  "github.com/southernvg/go-protobuf-lambda-pb"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "log"
)

func main() {
  var conn *grpc.ClientConn

  conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  c := ping_service.NewPingClient(conn)

  response, err := c.Ping(context.Background(), &ping_service.PingRequest{ Text: "Hello"})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Text)
}