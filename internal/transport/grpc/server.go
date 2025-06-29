package transportgrpc

import (
	"log"
	"net"

	userpb "github.com/IbadT/project-protos/proto/user"
	"github.com/IbadT/user_service-golang_microservice-.git/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервиса %v", err)
		return err
	}

	grpcSrv := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	log.Printf("GRPC запущен на порту 50051")
	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске grpc сервера %v", err)
		return err
	}
	return nil
}
