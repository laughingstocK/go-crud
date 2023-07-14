package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/laughingstocK/go-crud/author"
	"github.com/laughingstocK/go-crud/models"
	"google.golang.org/grpc"

	pb "github.com/laughingstocK/go-crud/proto/author"
)

type grpcAuthorRepo struct {
	GrpcConn *grpc.ClientConn
}

func NewGrpcAuthorRepo(GrpcConn *grpc.ClientConn) author.Repository {
	return &grpcAuthorRepo{GrpcConn}
}

func (m *grpcAuthorRepo) Create(ctx context.Context, name string) (*models.Author, error) {
	// Contact the server and print out its response.
	c := pb.NewAuthorClient(m.GrpcConn)
	fmt.Printf("22222 Author: %+v\n", name)
	r, err := c.Create(ctx, &pb.CreateRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	return nil, nil
}

func (m *grpcAuthorRepo) GetByID(ctx context.Context, id int64) (*models.Author, error) {

	return nil, nil
}
