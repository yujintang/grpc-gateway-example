package main

import (
	"context"
	"log"
	"mastiff/controller"
	pb "mastiff/proto"
	"net"
	"net/http"
	"os"
	"sort"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
