package cmd

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "example/go-grpc/pkg/todo"
)

const (
	address = "localhost:9000"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Search Todos",
	Long:  `Client for calling the GRPC server to search Todo objects`,
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn

		// Connect to the GRPC server
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		// create a new TdoClient using the grpc connection
		client := pb.NewTodoClient(conn)

		// name will be set to defaultName, and overwritten if a name param is passed into the cli command
		var searchId int64

		// cancel if no search arg
		if len(os.Args) < 3 {
			log.Fatal("missing search argument")
		}

		// try to convert incoming id into int64 to pass to GRPC
		parsedArg, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			log.Fatalf("invalid search argument: [%v]", os.Args[2])
		}
		searchId = parsedArg

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Call the GRPC endpoint on the grpc server
		res, err := client.GetTodo(ctx, &pb.TodoRequest{Id: searchId})
		if err != nil {
			log.Fatalf("could not fetch todo: %v", err.Error())
		}
		log.Printf("Found todo: : %v | %s", res.GetId(), res.GetContent())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
