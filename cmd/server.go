package cmd

import (
	"context"
	pb "example/go-grpc/pkg/todo"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":9000"
)

// server implements the TodoServer from the proto definition (i.e expose a single GetTodo method)
type Server struct {
	pb.UnimplementedTodoServer
}

// Define a local slice of dummy Todos to search
type Todo struct {
	Id      int64  `json:"id"`
	Content string `json:"name"`
}

var todos = []Todo{
	{Id: 1, Content: "todo one"},
	{Id: 2, Content: "todo two"},
	{Id: 3, Content: "todo three"},
	{Id: 4, Content: "todo four"},
	{Id: 5, Content: "todo five"},
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the GRPC server",
	Long:  `Starts the GRPC server`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		// init GRPC server
		grpcServer := grpc.NewServer()

		// register services
		pb.RegisterTodoServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

// Search the Todos slice and return the struct fields if found
// Otherwise, return an error
// This is attached as a method to the Server struct (i/e it is in parentheses before the func name)
func (s *Server) GetTodo(ctx context.Context, req *pb.TodoRequest) (*pb.TodoReply, error) {
	res := &pb.TodoReply{}

	// Check request exists
	if req == nil {
		return res, status.Error(codes.InvalidArgument, "request must not be nil")
	}

	// pull the search ID from the incoming message (TodoRequest in proto)
	searchId := req.GetId()

	log.Printf("Searching for id: [%v]", searchId)

	// find a todo using the incoming name from the GRPC request object
	var foundTodo Todo
	for _, v := range todos {
		if v.Id == searchId {
			foundTodo = v
			break
		}
	}

	// error if no todo found
	if (foundTodo == Todo{}) {
		return res, status.Error(codes.NotFound, "Not found")
	}

	// assign the return value to the response to send back to GRPC caller
	// response is typed from the TodoReply in the proto file
	res.Id = foundTodo.Id
	res.Content = foundTodo.Content

	return res, nil
}
