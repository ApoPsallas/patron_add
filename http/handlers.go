package http

import (
	"context"
	"strconv"

	"github.com/thebeatapp/patron/sync"
)

func jsonAPIResponse(data interface{}) *sync.Response {
	return sync.NewResponse(map[string]interface{}{"data": data})
}

//HandleMarco returns Polo
func HandleMarco(ctx context.Context, req *sync.Request) (*sync.Response, error) {
	data := map[string]string{"response": "Polo"}
	return jsonAPIResponse(data), nil
}

//HandleAddition returns the sum of two integers
func HandleAddition(ctx context.Context, req *sync.Request) (*sync.Response, error) {

	x, _ := strconv.Atoi(req.Fields["x"])
	y, _ := strconv.Atoi(req.Fields["y"])
	sum := x + y

	data := map[string]int{"sum": sum}
	return jsonAPIResponse(data), nil
}
