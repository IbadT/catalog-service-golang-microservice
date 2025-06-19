package transportgrpc

import (
	"log"
	"net"

	"github.com/IbadT/catalog-service-golang-microservice.git/internal/catalog"
	productpb "github.com/IbadT/project-protos/proto/product"
	"google.golang.org/grpc"
)

func RunGRPC(svc catalog.Service) error {
	listen, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервиса %v", err)
		return err
	}

	grpcSvc := grpc.NewServer()

	productpb.RegisterProductServiceServer(grpcSvc, NewHandler(svc))

	log.Printf("gRPC сервер запущен на порту 50054")
	if err := grpcSvc.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске grpc сервера %v", err)
		return err
	}
	return nil
}
