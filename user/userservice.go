package user

import (
	api "api/protobuf"
	"api/user/persister"
	"context"
	"log"
	"time"
)

type UserServiceServer struct {
	api.UnimplementedUserServiceServer
}

func (s *UserServiceServer) CreateUser(ctx context.Context, request *api.NewUser) (*api.User, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := persister.CreateUser(cancellableCtx, request.Email)
	if err != nil {
		log.Printf("Failed to create new user %v, error %v", request, err)
	}

	return user, err
}

func (s *UserServiceServer) UpdateUserEmail(ctx context.Context, request *api.UpdateUserEmailRequest) (*api.User, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := persister.UpdateUserEmail(cancellableCtx, request.UserId, request.Email)

	if err != nil {
		log.Printf("Failed to update user email %v, error %v", request, err)
	}

	return user, err
}

func (s *UserServiceServer) GetUserByEmail(ctx context.Context, request *api.GetUserByEmailRequest) (*api.User, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	user, err := persister.GetUser(cancellableCtx, nil, &request.Email)

	if err != nil {
		log.Printf("Couldn't get user %v, error %v", request, err)
	}

	return user, err
}

func (s *UserServiceServer) RemoveUser(ctx context.Context, request *api.RemoveUserRequest) (*api.Empty, error) {
	cancellableCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := persister.RemoveUser(cancellableCtx, request.UserId)

	if err != nil {
		log.Printf("Failed to remove user %v, error %v", request, err)
	}

	return &api.Empty{}, err
}
