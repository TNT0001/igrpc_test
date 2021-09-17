# iGHTK data serving API

[[_TOC_]]

## REST API

### Authentication

Xác thực bằng token tạo bởi ID service.
Loại token `Bearer` scope `igdata:query.read`

### Make query
```
POST /api/v1/query/{ConnectionName}/{TableName}
```

> Các params:

| Parameter    | Mandatory | datatype         | Description                 |
| ------------ | --------- | ---------------- | --------------------------- |
| page         | no        | int              | Trang                       |
| limit        | no        | int              | Giới hạn kết quả mỗi trang  |
| sorts        | no        | array            | Sort kết quả                |
| omit         | no        | array            | Danh sách columns bị bỏ qua |
| select       | no        | array            | Danh sách columns cần lấy   |
| conjunctions | no        | `[]conjunctions` | Các nhóm điều kiện          |

> Data type structures

1. condition

   | Property | Data type      | Description          |
       | -------- | -------------- | -------------------- |
   | column   | string         | tên column cần query |
   | operator | string         | Comparison operator  |
   | value    | string, number | Comparison value     |

1. conjunctions

   | Property   | Data type     | Description          |
       | ---------- | ------------- | -------------------- |
   | type       | string        | `or` hoặc `and`      |
   | conditions | `[]condition` | Danh sách conditions |

> Comparison operators

| Comparison operator | Mô tả                          | Bắt buộc value | Value data type    |
| ------------------- | ------------------------------ | -------------- | ------------------ |
| eq                  | So sánh bằng                   | y              | string,number      |
| neq                 | So sánh khác                   | y              | string,number      |
| gt                  | So sánh lớn hơn                | y              | number             |
| gte                 | So sánh lớn hơn hoặc bằng      | y              | number             |
| lt                  | So sánh nhỏ hơn                | y              | number             |
| lte                 | So sánh nhỏ hơn hoặc bằng      | y              | number             |
| in                  | Tìm item trong danh sách       | y              | []string, []number |
| nin                 | Tìm item không trong danh sách | y              | []string, []number |
| lk                  | Query LIKE                     | y              | string             |
| nlk                 | Query NOT LIKE                 | y              | string             |
| null                | IS NULL                        | n              | không dùng         |
| not_null            | IS NOT NULL                    | n              | không dùng         |

> Example request body
```json
{
    "page": 1,
    "limit": 10,
    "sorts": [
        {
            "column": "order",
            "mode": "desc"
        },
        {
            "column": "id",
            "mode": "asc"
        }
    ],
    "omit": [
        "ingore_column_name"
    ],
    "select": [
        "id"
    ],
    "conjunctions": [
        {
            "type": "or",
            "conditions": [
                {
                    "column": "order",
                    "operator": "lt",
                    "value": "1000000000000"
                },
                {
                    "column": "order",
                    "operator": "gt",
                    "value": "10"
                }
            ]
        },
        {
            "type": "or",
            "conditions": [
                {
                    "column": "order",
                    "operator": "lt",
                    "value": "1000000000000"
                },
                {
                    "column": "order",
                    "operator": "gt",
                    "value": "10"
                }
            ]
        }
    ]
}
```

## GRPC

**Note:** Các ví dụ phần này đều dùng ngôn ngữ golang.Tuy nhiên bạn có thể sử dụng những ngôn ngữ khác bằng cách dùng
grpc để generate từ file proto để được client và server side code cho ngôn ngữ tương ứng.
File proto nằm ở [igrpc_project_protofile](https://git.ghtk.vn/gmicro/ig/igrpc-proto/-/tree/master/proto/igdata-service)

### Create GRPC Client
Server side và client side code cho golang nằm ở [igdata-service](https://git.ghtk.vn/gmicro/ig/igrpc-proto/-/tree/master/generated/igdata-service).
sử dụng gmicro/grpcclient để tạo client.
> Create igdata client example:
```
client := grpcclient.NewClient(target string, opts ...ClientOption)
conn, err := client.Connect(context.Background())
igdataClient = igrpcproto.NewIgdataClient(conn)
```
`target : địa chỉ của grpc server`

`opts : các option cho việc tạo client`

### Authentication

> Nếu server yêu cầu xác thực với tls/ssl credentials thì có thể tạo `TransportCredentials` từ cert file và thêm 
vào option khi tạo client, example:

```
cred := credentials.NewClientTLSFromFile("certfile", "")
client := grpcclient.NewClient(target string, 
      grpcclient.WithTransportCredentials(cred),
      opts ...,
      )
```

> Xác thực bằng token tạo bởi ID service.Loại token `Bearer` scope `igdata:query.read`.

Có thể sử dụng interceptor cho việc thêm token, metada vào header cho từng lời gọi rpc, example:
```
type middleware struct {
}

func (m *middleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor{
	return 	func(ctx context.Context, method string, req interface{},
		reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, {token or metadata key}, {token or metadata value}, ...)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func (m *middleware)  StreamClientInterceptor() grpc.StreamClientInterceptor{
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
		method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error){
		newCtx := metadata.AppendToOutgoingContext(ctx, {token or metadata key}, {token or metadata value}, ...)
		return streamer(newCtx, desc, cc, method, opts...)
	}
}
```
Và thêm interceptor này vào khi tạo client, ex:
```
client := grpcclient.NewClient(target string, 
      grpcclient.WithTransportCredentials(cred),
      grpcclient.WithInterceptor(*middleware),
      opts ...,
      )
```

### Make Query

Những request, response struct đều trong file ở [igdata-service](https://git.ghtk.vn/gmicro/ig/igrpc-proto/-/tree/master/generated/igdata-service).

> Create request with proto struct

Các request trong gprc đều tương ứng với các request trong http api. ví dụ request với SQLQuery:

```
type SQLQueryRequest struct {
	Paginate       *SQLPaginate      `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	ConnectionName string            `protobuf:"bytes,2,opt,name=connection_name,json=connectionName,proto3" json:"connection_name,omitempty"`
	TableName      string            `protobuf:"bytes,3,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Select         []string          `protobuf:"bytes,4,rep,name=select,proto3" json:"select,omitempty"`
	OmitColumns    []string          `protobuf:"bytes,5,rep,name=omit_columns,json=omitColumns,proto3" json:"omit_columns,omitempty"`
	Sorts          []*SQLSort        `protobuf:"bytes,6,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Conjunctions   []*SQLConjunction `protobuf:"bytes,7,rep,name=conjunctions,proto3" json:"conjunctions,omitempty"`
}
```
```
type SQLPaginate struct {
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}
```
```
type SQLSort struct {
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Mode   string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
}
```
```
type SQLConjunction struct {
	Type       string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Conditions []*SQLCondition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
}
```
```
type SQLCondition struct {
	Column             string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	ComparisonOperator string `protobuf:"bytes,2,opt,name=comparison_operator,json=comparisonOperator,proto3" json:"comparison_operator,omitempty"`
	Value              []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

```


> Call RPC

Sử dụng client đã tạo để gọi rpc. Hiện tại thì server đang hỗ trợ hai cách gọi rpc là 
1. unary call 
2. bidirection stream call

Ví dụ cho việc gọi unary call SQLQuery :

```
igdataClient.SQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption)
```

**Note:** if use haproxy with nginx behind and use tls/ssl.you can turn haproxy mode to tcp and let nginx handle tls/ssl.

