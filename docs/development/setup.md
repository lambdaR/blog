# Development Setup

This guide will help you set up your development environment for working with the Micro Blog project.

## Prerequisites

Before you begin, make sure you have the following installed:

1. **Go 1.24 or higher**

   Check your Go version:
   ```bash
   go version
   ```

   If you need to install or update Go, visit the [official Go download page](https://golang.org/dl/).

2. **Micro CLI**

   The Micro Blog project uses go-micro v5. Install the Micro CLI:
   ```bash
   go install github.com/micro/micro/v5@master
   ```

   Make sure that `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH` so you can use the `micro` command:
   ```bash
   export PATH=$PATH:$GOPATH/bin
   # or
   export PATH=$PATH:$HOME/go/bin
   ```

   Verify the installation:
   ```bash
   micro --version
   ```

3. **Protocol Buffers Compiler**

   Install the Protocol Buffers compiler (protoc):

   - **Linux**:
     ```bash
     apt install -y protobuf-compiler
     ```

   - **macOS**:
     ```bash
     brew install protobuf
     ```

   - **Windows**: Download from [GitHub releases](https://github.com/protocolbuffers/protobuf/releases)

   Verify the installation:
   ```bash
   protoc --version
   ```

4. **Protocol Buffers Go Plugins**

   Install the Go plugins for Protocol Buffers:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install go-micro.dev/v5/cmd/protoc-gen-micro@latest
   ```

## Project Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/micro/blog.git
   cd blog
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

3. **Generate Protocol Buffer Code**

   The project includes a Makefile with commands to generate Protocol Buffer code:
   ```bash
   make gen-proto
   ```

   This will generate the necessary Go code from the Protocol Buffer definitions.

## IDE Setup

For the best development experience, we recommend using an IDE with good Go support:

### Visual Studio Code

1. Install Visual Studio Code
2. Install the Go extension
3. Open the project folder
4. The Go extension will prompt you to install various Go tools - accept these installations

### GoLand

1. Install GoLand
2. Open the project folder
3. GoLand should automatically detect the Go module and set up the project

## Environment Configuration

The Micro Blog project uses default configurations for development. If you need to customize the environment:

1. **Service Registry**

   By default, go-micro v5 uses an in-memory registry. For development, this is sufficient.

2. **Data Storage**

   By default, the services use go-micro's in-memory store. For development, this is sufficient.

3. **Network Configuration**

   The Web Service runs on port 42096 by default. Make sure this port is available on your system.

## Next Steps

Once your development environment is set up, you can:

1. [Run the services](running.md)
2. [Build the services](building.md)
3. Start exploring and modifying the code