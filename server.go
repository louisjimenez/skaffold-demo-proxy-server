package main

import (
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/louisjimenez/skaffold-demo-config"
)

const (
	address = "localhost:50051"
)

type server struct {
	pb.UnimplementedTodoServer
	todoItems []*pb.TodoItem
}

var mockData = []byte(`[{
    "description": "Attend meeting",
	"category": {
        "name": "work"
    },
	"id": 1
}, {
	"description": "Write code",
    "category": {
        "name": "work"
    },
	"id": 2
}]`)

func (s *server) ListTodos(cat *pb.Category, stream pb.Todo_ListTodosServer) error {
	for _, todo := range s.todoItems {
		if todo.Category.Name == cat.Name {
			if err := stream.Send(todo); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *server) loadTodoItems() {
	if err := json.Unmarshal(mockData, &s.todoItems); err != nil {
		log.Fatalf("Unable to load mock todo items: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := &server{}
	s.loadTodoItems()
	pb.RegisterTodoServer(grpcServer, s)
	grpcServer.Serve(lis)
}
