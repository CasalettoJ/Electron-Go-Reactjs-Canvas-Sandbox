package hello

import (
	"strings"

	pb "github.com/casalettoj/exploration/App/proto"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SendChat repeat request as response to prove connection
func (s *Server) SendChat(ctx context.Context, in *pb.Chat) (*pb.Chat, error) {
	response := [2]string{"Recieved Message: ", in.Message}
	responseStr := strings.Trim(strings.Join(response[:], ""), " ")
	return &pb.Chat{Message: responseStr}, nil
}
