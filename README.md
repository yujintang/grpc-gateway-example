## 本例文件结构:

```
├── README.md
├── bin
│   ├── linux_amd64
│   └── mac_amd64
├── controller
│   └── mastiff.go
├── example
│   └── client.go
├── go.mod
├── go.sum
├── main.go
├── makefile
└── proto
    ├── mastiff.pb.go
    ├── mastiff.pb.gw.go
    ├── mastiff.proto
    └── mastiff.swagger.json
```



##Protobuf3 语法

[proto3简介](https://developers.google.com/protocol-buffers/docs/proto3#simple)

```protobuf
syntax = "proto3";

package pb_mastiff;
option go_package = "proto;mastiff";

import "google/api/annotations.proto";

// 创建Id
message CreateIdReq {
    string Prefix = 1;
}

message CreateIdRes {
    string Data = 1;
}

service Mastiff  {
    rpc createId(CreateIdReq) returns (CreateIdRes){
        option (google.api.http) = {
        post: "/v1/createId"
        body: "*"
        };
    }
}
```

1. 第一行`syntax`用来指定 `pb`版本，必须位于第一行，不添加默认为`protobuf2`
2. 消息体定义由 `message`开始; 参数建议使用下划线分割; 每个参数需要添加唯一编号，用于标识`消息二进制格式字段`; `label` 默认使用 `singular`(0-1次), `repeated`(0-n次) `optional`和`required`被弃用。
3. 注释使用 `//` 或 `/* ... */`
4. 更新消息类型过程中，切勿删除现有字段与字段编号，使用保留字段`reserved`, 直接删除或后面重用该字段，可能会导致各种问题。
5. 参数类型为枚举类型，枚举需要从`0`开始。
6. 其他消息类型，可嵌套使用，如要在父消息类型外使用， 则使用 `_Parent_._Type_`
7. 导入其他消息类型，可使用`import` 关键字， 若引入文件目录不确定，则直接引用pb文件，编译期间使用`-I`/`--proto_path`是更好的选择。pb2与pb3可互相引用，但是不能在proto3语法中使用proto2枚举。
8. `package` 用于防止出现命名冲突，相当于一个命名空间的作用
9. `option` 的`go_package` 用于声明生成的`go`代码package 包名。`option go_package`  >  `package` > `文件名`

## pb编译

[protobuf简介](https://github.com/protocolbuffers/protobuf)

### pb在Mac环境下，golang语言的使用

1. 安装protobuf
   ```shell
   brew install protobuf
   ```
   
   
   
2. 安装golang代码生成依赖

   ```shell
   go get -u -v github.com/golang/protobuf/proto
   go get -u -v github.com/golang/protobuf/protoc-gen-go
   go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
   go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
   ```

3. makefile 如下：

   > --proto_path= 或 -I ：用于指定proto依赖。grpc-gateway 文件根据不同版本，可能需要调整其路径
   >
   > --xx_out= : 根据我们上面安装protoc-gen-xx, 生成对应的xx文件。

   ```makefile
   .PHONY:all
   all:clean build compile
   
   build:
   	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.2/third_party/googleapis --go_out=plugins=grpc:. proto/mastiff.proto
   	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.2/third_party/googleapis --grpc-gateway_out=logtostderr=true:. proto/mastiff.proto
   	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.2/third_party/googleapis --swagger_out=logtostderr=true:. proto/mastiff.proto
   
   compile:
   	GOOS=darwin GOARCH=amd64 go build -o bin/mac_amd64 main.go
   	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64 main.go
   
   clean:
   	rm -rf ./proto/*.go ./proto/*.json ./bin
   ```

   

4. 执行`make`,生成对应文件
		```shell
      make
   ```

    

## RPC服务与Gateway服务

### 启动service

> 分别启动`rpc`与`http`服务

```shell
go run main.go rpc
go run main.go http
```

### 调用RPC

```shell
go run example/client.go
```

### 调用Http

```shell
curl --location --request POST 'localhost:1728/v1/createId' --header 'Content-Type: application/json' --data-raw '{"Prefix": "dml:"}'
```



## 代码

### controller/mastiff

```go
package controller

import (
	"context"
	pb "mastiff/proto"
)

type Service struct{}

// 创建ID
func (s Service) CreateId(ctx context.Context, req *pb.CreateIdReq) (*pb.CreateIdRes, error) {
	return &pb.CreateIdRes{
		Data: req.Prefix + "123456",
	}, nil
}
```

### main.go

```
package main

import (
   "context"
   grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
   grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
   "github.com/grpc-ecosystem/grpc-gateway/runtime"
   "github.com/urfave/cli"
   "google.golang.org/grpc"
   "google.golang.org/grpc/reflection"
   "log"
   "mastiff/controller"
   pb "mastiff/proto"
   "net"
   "net/http"
   "os"
   "sort"
)

const (
   RpcPort  = ":1718"
   HttpPort = ":1728"
)

func main() {
   mastiff := cli.NewApp()
   mastiff.Name = "controller"
   mastiff.Version = "1.0.0"
   mastiff.Email = "shanquan54@gmail.com"
   mastiff.Usage = "用于Token管理,Auth授权检测的服务系统"
   mastiff.Commands = []cli.Command{
      {
         Name:        "http",
         Usage:       "http 服务",
         Description: "启动http服务",
         Action:      httpServe,
      },
      {
         Name:        "rpc",
         Usage:       "rpc 服务",
         Description: "启动rpc服务",
         Action:      rpcSereve,
      },
   }
   sort.Sort(cli.FlagsByName(mastiff.Flags))
   sort.Sort(cli.CommandsByName(mastiff.Commands))
   err := mastiff.Run(os.Args)
   if err != nil {
      log.Fatal(err.Error())
   }
}

func httpServe(c *cli.Context) {
   ctx := context.Background()
   ctx, cancel := context.WithCancel(ctx)
   defer cancel()

   mux := runtime.NewServeMux()
   opts := []grpc.DialOption{grpc.WithInsecure()}
   err := pb.RegisterMastiffHandlerFromEndpoint(ctx, mux, RpcPort, opts)
   if err != nil {
      log.Fatal(err)
   }
   log.Print("Listen Rpc Service: " + RpcPort)
   log.Print("init Http Service: " + HttpPort)
   if err := http.ListenAndServe(HttpPort, mux); err != nil {
      log.Fatal("Http Gateway failed")
   }
}

func rpcSereve(c *cli.Context) {

   // 启动服务监听
   lis, err := net.Listen("tcp", RpcPort)
   if err != nil {
      log.Fatal(err)
   }
   s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
      grpc_recovery.UnaryServerInterceptor(),
   )))
   pb.RegisterMastiffServer(s, &controller.MastiffService{})
   reflection.Register(s)

   log.Print("init Rpc Service: " + RpcPort)
   if err := s.Serve(lis); err != nil {
      log.Fatal("Rpc failed")
   }
}
```

### example/client.go

```
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "mastiff/proto"
)

func main() {
	conn, err := grpc.Dial(":1718", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	mastiff := pb.NewMastiffClient(conn)

	data, err := mastiff.CreateId(context.Background(), &pb.CreateIdReq{Prefix: "dml:"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(data)
}
```