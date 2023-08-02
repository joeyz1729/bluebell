package router

import (
	"context"
	"fmt"

	"github.com/YiZou89/bluebell/logic"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	ID uint64 `json:"id"`
}
type UserResponse struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"`
}

func getUser(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(UserRequest)

	name := fmt.Sprintf("User %d", req.ID)
	return UserResponse{Name: name}, nil
}

// 定义端点（Endpoint）
func makeGetUserEndpoint(us logic.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		username, err := us.Info(ctx, req.ID)
		if err != nil {
			return UserResponse{Name: username, Err: err.Error()}, nil
		}
		//name := fmt.Sprintf("User %d", req.ID)
		return UserResponse{Name: username}, nil
	}
}
