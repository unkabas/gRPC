package main

import (
	"context"
	greetpb "gRPC/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Устанавливаем insecure-соединение с сервером.
	// В продакшене используйте grpc.WithTransportCredentials(credentials.NewTLS(...))
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. Создаем новый клиент.
	// Функция NewGreeterClient и структура GreeterClient были сгенерированы автоматически.
	c := greetpb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 3. Вызываем метод SayHello!
	// Мы просто используем готовую функцию. Вся работа по сети, сериализации
	// запроса и десериализации ответа происходит под капотом.
	r, err := c.SayHello(ctx, &greetpb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// r — это уже готовый, автоматически разобранный HelloResponse.
	log.Printf("Greeting: %s", r.GetGreeting())
}
