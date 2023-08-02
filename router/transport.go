package router

import (
	"context"
	"encoding/json"
	"net/http"
)

// 定义编解码器
func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func encodeUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
