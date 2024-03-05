package cmd

import (
	"connectrpc.com/validate"
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todov1 "connect-todo-tutorial-backend/gen/proto/todo/v1"
	"connect-todo-tutorial-backend/gen/proto/todo/v1/todov1connect"
)

var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverRun,
}

func serverRun(_ *cobra.Command, _ []string) error {
	log.Println("start server")
	interceptor, err := validate.NewInterceptor()
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	// mux.Handle(todov1connect.NewTodoServiceHandler(&TodoServer{}))
	mux.Handle(todov1connect.NewTodoServiceHandler(&TodoServer{}, connect.WithInterceptors(interceptor)))
	reflector := grpcreflect.NewStaticReflector(
		todov1connect.TodoServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	if err := http.ListenAndServe(
		"0.0.0.0:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
	return nil
}

type TodoServer struct{}

func (s *TodoServer) ListTodos(ctx context.Context, req *connect.Request[todov1.ListTodosRequest]) (*connect.Response[todov1.ListTodosResponse], error) {
	log.Println("call ListTodos")
	res := connect.NewResponse(&todov1.ListTodosResponse{})
	// res.Header().Set("TodoService-Version", "v1")
	return res, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, req *connect.Request[todov1.GetTodoRequest]) (*connect.Response[todov1.GetTodoResponse], error) {
	log.Println("call GetTodo")
	//v, err := protovalidate.New()
	//if err != nil {
	//	return nil, err
	//}
	//if err = v.Validate(req.Msg); err != nil {
	//	return nil, err
	//}
	res := connect.NewResponse(&todov1.GetTodoResponse{
		Todo: &todov1.Todo{
			Id:     req.Msg.Id,
			Name:   "John",
			IsDone: false,
		},
	})
	// res.Header().Set("TodoService-Version", "v1")
	return res, nil
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *connect.Request[todov1.CreateTodoRequest]) (*connect.Response[todov1.CreateTodoResponse], error) {
	log.Println("call CreateTodo")
	res := connect.NewResponse(&todov1.CreateTodoResponse{})
	// res.Header().Set("TodoService-Version", "v1")
	return res, nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *connect.Request[todov1.UpdateTodoRequest]) (*connect.Response[todov1.UpdateTodoResponse], error) {
	log.Println("call UpdateTodo")
	res := connect.NewResponse(&todov1.UpdateTodoResponse{})
	// res.Header().Set("TodoService-Version", "v1")
	return res, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *connect.Request[todov1.DeleteTodoRequest]) (*connect.Response[todov1.DeleteTodoResponse], error) {
	log.Println("call DeleteTodo")
	res := connect.NewResponse(&todov1.DeleteTodoResponse{})
	// res.Header().Set("TodoService-Version", "v1")
	return res, nil
}
