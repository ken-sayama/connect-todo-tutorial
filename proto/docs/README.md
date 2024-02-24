# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/todo/v1/todo.proto](#proto_todo_v1_todo-proto)
    - [CreateTodoRequest](#proto-todo-v1-CreateTodoRequest)
    - [CreateTodoResponse](#proto-todo-v1-CreateTodoResponse)
    - [DeleteTodoRequest](#proto-todo-v1-DeleteTodoRequest)
    - [DeleteTodoResponse](#proto-todo-v1-DeleteTodoResponse)
    - [GetTodoRequest](#proto-todo-v1-GetTodoRequest)
    - [GetTodoResponse](#proto-todo-v1-GetTodoResponse)
    - [ListTodosRequest](#proto-todo-v1-ListTodosRequest)
    - [ListTodosResponse](#proto-todo-v1-ListTodosResponse)
    - [Todo](#proto-todo-v1-Todo)
    - [UpdateTodoRequest](#proto-todo-v1-UpdateTodoRequest)
    - [UpdateTodoResponse](#proto-todo-v1-UpdateTodoResponse)
  
    - [TodoService](#proto-todo-v1-TodoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_todo_v1_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/todo/v1/todo.proto



<a name="proto-todo-v1-CreateTodoRequest"></a>

### CreateTodoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="proto-todo-v1-CreateTodoResponse"></a>

### CreateTodoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#proto-todo-v1-Todo) |  |  |






<a name="proto-todo-v1-DeleteTodoRequest"></a>

### DeleteTodoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="proto-todo-v1-DeleteTodoResponse"></a>

### DeleteTodoResponse







<a name="proto-todo-v1-GetTodoRequest"></a>

### GetTodoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="proto-todo-v1-GetTodoResponse"></a>

### GetTodoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#proto-todo-v1-Todo) |  |  |






<a name="proto-todo-v1-ListTodosRequest"></a>

### ListTodosRequest







<a name="proto-todo-v1-ListTodosResponse"></a>

### ListTodosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todos | [Todo](#proto-todo-v1-Todo) | repeated |  |






<a name="proto-todo-v1-Todo"></a>

### Todo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| is_done | [bool](#bool) |  |  |






<a name="proto-todo-v1-UpdateTodoRequest"></a>

### UpdateTodoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| is_done | [bool](#bool) |  |  |






<a name="proto-todo-v1-UpdateTodoResponse"></a>

### UpdateTodoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#proto-todo-v1-Todo) |  |  |





 

 

 


<a name="proto-todo-v1-TodoService"></a>

### TodoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListTodos | [ListTodosRequest](#proto-todo-v1-ListTodosRequest) | [ListTodosResponse](#proto-todo-v1-ListTodosResponse) |  |
| GetTodo | [GetTodoRequest](#proto-todo-v1-GetTodoRequest) | [GetTodoResponse](#proto-todo-v1-GetTodoResponse) |  |
| CreateTodo | [CreateTodoRequest](#proto-todo-v1-CreateTodoRequest) | [CreateTodoResponse](#proto-todo-v1-CreateTodoResponse) |  |
| UpdateTodo | [UpdateTodoRequest](#proto-todo-v1-UpdateTodoRequest) | [UpdateTodoResponse](#proto-todo-v1-UpdateTodoResponse) |  |
| DeleteTodo | [DeleteTodoRequest](#proto-todo-v1-DeleteTodoRequest) | [DeleteTodoResponse](#proto-todo-v1-DeleteTodoResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

