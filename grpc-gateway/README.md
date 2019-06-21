# 依赖
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

$ cd proto
$ bash gen.sh
$ bash gen-reverse-proxy.sh
// 生成后启动
$ cd ..
$ go run main.go

// 测试
$ curl -X POST -k "http://localhost:8080/v1/example/echo" -d '{"value": "CoreOS is hiring!"}'
$ curl -vX POST -k "http://localhost:8080/v1/example/addItem" -d '{"userId":2,"itemType":3,"itemQuantity":3,"orderId":"123","source":12323,"prptys":[{"itemId":"qwer","id":12,"value":123,"info":"info"}],"itemName":"12321"}'
