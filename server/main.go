package main

import (
	"context"
	"gRPC/pb"

	"log"
	"net"

	"google.golang.org/grpc"
)

// 1. Создаем структуру, которая будет реализовывать наш интерфейс.
type server struct {
	greetpb.UnimplementedGreeterServer // Важно: встраиваем сгенерированную структуру
}

// 2. Реализуем метод SayHello, точно как объявлено в интерфейсе.
func (s *server) SayHello(ctx context.Context, req *greetpb.HelloRequest) (*greetpb.HelloResponse, error) {
	// req — это уже готовый, автоматически разобранный из сети HelloRequest.
	// Вам не нужно самим парсить байты!
	log.Printf("Received: %v", req.GetName())

	// Формируем ответ. greetpb.HelloResponse — тоже сгенерированная структура.
	return &greetpb.HelloResponse{
		Greeting: "Hello, " + req.GetName(),
	}, nil // Возвращаем ответ и ошибку (nil — значит, ошибки нет)
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // Создаем экземпляр gRPC-сервера

	// 3. Регистрируем нашу реализацию сервера в gRPC-сервере.
	// Функция RegisterGreeterServer была сгенерирована автоматически.
	greetpb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
