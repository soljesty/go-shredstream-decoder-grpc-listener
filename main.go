package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-listener/pb"
)

func main() {
    addr := "" // IP to connect to listen for transactions
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Connection could not be established: %v", err)
    }
    defer conn.Close()
    
    client := pb.NewTransactionServiceClient(conn)
    req := &pb.StreamTransactionsRequest{}

    stream, err := client.StreamTransactions(context.Background(), req)
    if err != nil {
        log.Fatalf("Error calling StreamTransactions: %v", err)
    }

    for {
        resp, err := stream.Recv()
        if err != nil {
            log.Fatalf("Error receiving message: %v", err)
        }
        currentTimestamp := time.Now().UnixMicro() 
        latencyMicro := currentTimestamp - int64(resp.GetTimestamp())
        latencyMilli := latencyMicro / 1000
        
        fmt.Printf("Transaction received: %s, Latency: %d ms\n", resp.GetTransactionJson(), latencyMilli)
    }
}
