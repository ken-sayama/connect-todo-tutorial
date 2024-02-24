// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/todo/v1/todo.proto

package todov1connect

import (
	v1 "connect-todo-tutorial-backend/gen/proto/todo/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "proto.todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceListTodosProcedure is the fully-qualified name of the TodoService's ListTodos RPC.
	TodoServiceListTodosProcedure = "/proto.todo.v1.TodoService/ListTodos"
	// TodoServiceGetTodoProcedure is the fully-qualified name of the TodoService's GetTodo RPC.
	TodoServiceGetTodoProcedure = "/proto.todo.v1.TodoService/GetTodo"
	// TodoServiceCreateTodoProcedure is the fully-qualified name of the TodoService's CreateTodo RPC.
	TodoServiceCreateTodoProcedure = "/proto.todo.v1.TodoService/CreateTodo"
	// TodoServiceUpdateTodoProcedure is the fully-qualified name of the TodoService's UpdateTodo RPC.
	TodoServiceUpdateTodoProcedure = "/proto.todo.v1.TodoService/UpdateTodo"
	// TodoServiceDeleteTodoProcedure is the fully-qualified name of the TodoService's DeleteTodo RPC.
	TodoServiceDeleteTodoProcedure = "/proto.todo.v1.TodoService/DeleteTodo"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	todoServiceServiceDescriptor          = v1.File_proto_todo_v1_todo_proto.Services().ByName("TodoService")
	todoServiceListTodosMethodDescriptor  = todoServiceServiceDescriptor.Methods().ByName("ListTodos")
	todoServiceGetTodoMethodDescriptor    = todoServiceServiceDescriptor.Methods().ByName("GetTodo")
	todoServiceCreateTodoMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("CreateTodo")
	todoServiceUpdateTodoMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("UpdateTodo")
	todoServiceDeleteTodoMethodDescriptor = todoServiceServiceDescriptor.Methods().ByName("DeleteTodo")
)

// TodoServiceClient is a client for the proto.todo.v1.TodoService service.
type TodoServiceClient interface {
	ListTodos(context.Context, *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error)
	GetTodo(context.Context, *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.GetTodoResponse], error)
	CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error)
	UpdateTodo(context.Context, *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.UpdateTodoResponse], error)
	DeleteTodo(context.Context, *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error)
}

// NewTodoServiceClient constructs a client for the proto.todo.v1.TodoService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		listTodos: connect.NewClient[v1.ListTodosRequest, v1.ListTodosResponse](
			httpClient,
			baseURL+TodoServiceListTodosProcedure,
			connect.WithSchema(todoServiceListTodosMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTodo: connect.NewClient[v1.GetTodoRequest, v1.GetTodoResponse](
			httpClient,
			baseURL+TodoServiceGetTodoProcedure,
			connect.WithSchema(todoServiceGetTodoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createTodo: connect.NewClient[v1.CreateTodoRequest, v1.CreateTodoResponse](
			httpClient,
			baseURL+TodoServiceCreateTodoProcedure,
			connect.WithSchema(todoServiceCreateTodoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateTodo: connect.NewClient[v1.UpdateTodoRequest, v1.UpdateTodoResponse](
			httpClient,
			baseURL+TodoServiceUpdateTodoProcedure,
			connect.WithSchema(todoServiceUpdateTodoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTodo: connect.NewClient[v1.DeleteTodoRequest, v1.DeleteTodoResponse](
			httpClient,
			baseURL+TodoServiceDeleteTodoProcedure,
			connect.WithSchema(todoServiceDeleteTodoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	listTodos  *connect.Client[v1.ListTodosRequest, v1.ListTodosResponse]
	getTodo    *connect.Client[v1.GetTodoRequest, v1.GetTodoResponse]
	createTodo *connect.Client[v1.CreateTodoRequest, v1.CreateTodoResponse]
	updateTodo *connect.Client[v1.UpdateTodoRequest, v1.UpdateTodoResponse]
	deleteTodo *connect.Client[v1.DeleteTodoRequest, v1.DeleteTodoResponse]
}

// ListTodos calls proto.todo.v1.TodoService.ListTodos.
func (c *todoServiceClient) ListTodos(ctx context.Context, req *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error) {
	return c.listTodos.CallUnary(ctx, req)
}

// GetTodo calls proto.todo.v1.TodoService.GetTodo.
func (c *todoServiceClient) GetTodo(ctx context.Context, req *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.GetTodoResponse], error) {
	return c.getTodo.CallUnary(ctx, req)
}

// CreateTodo calls proto.todo.v1.TodoService.CreateTodo.
func (c *todoServiceClient) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	return c.createTodo.CallUnary(ctx, req)
}

// UpdateTodo calls proto.todo.v1.TodoService.UpdateTodo.
func (c *todoServiceClient) UpdateTodo(ctx context.Context, req *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.UpdateTodoResponse], error) {
	return c.updateTodo.CallUnary(ctx, req)
}

// DeleteTodo calls proto.todo.v1.TodoService.DeleteTodo.
func (c *todoServiceClient) DeleteTodo(ctx context.Context, req *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error) {
	return c.deleteTodo.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the proto.todo.v1.TodoService service.
type TodoServiceHandler interface {
	ListTodos(context.Context, *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error)
	GetTodo(context.Context, *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.GetTodoResponse], error)
	CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error)
	UpdateTodo(context.Context, *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.UpdateTodoResponse], error)
	DeleteTodo(context.Context, *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	todoServiceListTodosHandler := connect.NewUnaryHandler(
		TodoServiceListTodosProcedure,
		svc.ListTodos,
		connect.WithSchema(todoServiceListTodosMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceGetTodoHandler := connect.NewUnaryHandler(
		TodoServiceGetTodoProcedure,
		svc.GetTodo,
		connect.WithSchema(todoServiceGetTodoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceCreateTodoHandler := connect.NewUnaryHandler(
		TodoServiceCreateTodoProcedure,
		svc.CreateTodo,
		connect.WithSchema(todoServiceCreateTodoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceUpdateTodoHandler := connect.NewUnaryHandler(
		TodoServiceUpdateTodoProcedure,
		svc.UpdateTodo,
		connect.WithSchema(todoServiceUpdateTodoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	todoServiceDeleteTodoHandler := connect.NewUnaryHandler(
		TodoServiceDeleteTodoProcedure,
		svc.DeleteTodo,
		connect.WithSchema(todoServiceDeleteTodoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceListTodosProcedure:
			todoServiceListTodosHandler.ServeHTTP(w, r)
		case TodoServiceGetTodoProcedure:
			todoServiceGetTodoHandler.ServeHTTP(w, r)
		case TodoServiceCreateTodoProcedure:
			todoServiceCreateTodoHandler.ServeHTTP(w, r)
		case TodoServiceUpdateTodoProcedure:
			todoServiceUpdateTodoHandler.ServeHTTP(w, r)
		case TodoServiceDeleteTodoProcedure:
			todoServiceDeleteTodoHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) ListTodos(context.Context, *connect.Request[v1.ListTodosRequest]) (*connect.Response[v1.ListTodosResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.ListTodos is not implemented"))
}

func (UnimplementedTodoServiceHandler) GetTodo(context.Context, *connect.Request[v1.GetTodoRequest]) (*connect.Response[v1.GetTodoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.GetTodo is not implemented"))
}

func (UnimplementedTodoServiceHandler) CreateTodo(context.Context, *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.CreateTodo is not implemented"))
}

func (UnimplementedTodoServiceHandler) UpdateTodo(context.Context, *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.UpdateTodoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.UpdateTodo is not implemented"))
}

func (UnimplementedTodoServiceHandler) DeleteTodo(context.Context, *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.todo.v1.TodoService.DeleteTodo is not implemented"))
}
