# Shredstream Decoder Go Listener

Go listener that gets for transactions from the **Shredstream Decoder's** gRPC server. The listener connects to a specified gRPC endpoint, receives a stream of messages (all of which contain a time-stamped JSON transaction), and calculates the latency between the server timestamp and the time the message arrives.

# Prerequisites

- Go 1.17+ (use Go 1.21 or later)
- Protocol Buffers (protoc)
- Go plugins for Protocol Buffers:
  
`protoc-gen-go`

`protoc-gen-go-grpc`

# Installation

### 1. Install Go

Follow the official install directions or your package manager. 
```
sudo apt update
sudo apt install golang-go
```
Or install from: `https://go.dev/dl/`

### 2. Install the Protocol Buffer plugins
```
ngo install google.golang.org/protobuf/cmd/protoc-gen-go@latest
ngo install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
- Ensure your `$HOME/go/bin` is in your PATH:

```
export PATH=$PATH:$HOME/go/bin
source ~/.bashrc
```

# Configuration

- #### Update the Server Address

In `main.go` file, replace the `addr` variable with your Shredstream Decoder's gRPC server public IP and port.

- #### Verify Import Paths

Make sure the import path of the generated package in `main.go` is identical to the module name and directory structure in your go.mod.

# Run the listener

To execute the client, go to the project directory in a terminal and execute:
`go run main.go`
