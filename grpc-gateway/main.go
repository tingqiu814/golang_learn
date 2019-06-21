package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "example/grpc-gateway/proto"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterAddItemServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	var svc Svc
	gw.RegisterYourServiceServer(s, svc)
	var addSvc AddItemSvc
	gw.RegisterAddItemServiceServer(s, addSvc)

	go func() {
		l, err := net.Listen("tcp", ":9090")
		if err != nil {
			panic(fmt.Sprintf("listen err: %v", err))
		}
		s.Serve(l)
	}()

	return http.ListenAndServe(":8080", mux)
}

type Svc struct{}

func (s Svc) Echo(c context.Context, g *gw.StringMessage) (*gw.StringMessage, error) {
	fmt.Printf("input c: %v, \ng: %v", c, g)
	return g, nil
}

type AddItemSvc struct{}

func (s AddItemSvc) AddItem(c context.Context, g *gw.AddItemRequest) (rg *gw.AddItemRequest, err error) {
	fmt.Printf("======================\n")
	//	x := gw.AddItemRequest{
	//		UserId:       uint64(2),
	//		ItemType:     uint64(3),
	//		ItemQuantity: int64(3),
	//		OrderId:      string("123"),
	//		Source:       uint64(12323),
	//		Prptys: []*gw.Property{
	//			&gw.Property{
	//				ItemId: string("qwer"),
	//				Id:     uint64(12),
	//				Value:  int64(123),
	//				Info:   string("info"),
	//			},
	//		},
	//		ItemName: string("12321"),
	//	}
	b, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("input: %v\n", string(b))
	fmt.Printf("input c: %v, \ng: %v", c, string(b))
	return g, nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
