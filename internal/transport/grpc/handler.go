package transportgrpc

import (
	"context"

	userpb "github.com/IbadT/project-protos/proto/user"
	"github.com/IbadT/user_service-golang_microservice-.git/internal/user"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc user.Service
	// taskClient taskpb
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	email := req.Email
	user, err := h.svc.CreateUser(email)
	if err != nil {
		return &userpb.CreateUserResponse{}, err
	}

	pbUser := userpb.User{
		Id:    user.Id.String(),
		Email: user.Email,
	}
	return &userpb.CreateUserResponse{User: &pbUser}, nil
}

// func (h *Handler) GetUserWithTasks(ctx context.Context, req *userpb.GetUserWithTasksRequest) (*userpb.GetUserWithTasksResponse, error) {

// }

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	user, err := h.svc.GetUser(userId)
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	pbUser := &userpb.User{
		Id:    user.Id.String(),
		Email: user.Email,
	}

	return &userpb.GetUserResponse{
		User: pbUser,
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *emptypb.Empty) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.ListUsers()
	if err != nil {
		return &userpb.ListUsersResponse{}, err
	}

	pbUsers := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		usr := userpb.User{
			Id:    u.Id.String(),
			Email: u.Email,
		}
		pbUsers = append(pbUsers, &usr)
	}

	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	userId, err := uuid.Parse(req.User.Id)
	if err != nil {
		return &userpb.UpdateUserResponse{}, err
	}
	email := req.User.Email
	user, err := h.svc.UpdateUser(userId, email)
	if err != nil {
		return &userpb.UpdateUserResponse{}, err
	}

	return &userpb.UpdateUserResponse{User: &userpb.User{
		Id:    user.Id.String(),
		Email: user.Email,
	}}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return &userpb.DeleteUserResponse{}, err
	}

	err = h.svc.DeleteUser(userId)
	if err != nil {
		return &userpb.DeleteUserResponse{}, err
	}
	return &userpb.DeleteUserResponse{Id: req.Id}, nil
}

// go get github.com/IbadT/project-protos@latest
// go mod tidy
