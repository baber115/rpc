# 如何生成代码

# 生成service pb编译文件
```shell
protoc -I=./grpc/simple/pb/ --go_out=./grpc/simple/pb/ --go_opt=module="rpc-g7/grpc/simple/pb" ./grpc/simple/pb/hello.proto
```

# 补充rpc 接口定义protobuf的代码生成
```shell
protoc -I=./grpc/simple/pb/ --go_out=./grpc/simple/pb/ --go_opt=module="rpc-g7/grpc/simple/pb" --go-grpc_out=./grpc/simple/pb/ --go-grpc_opt=module="rpc-g7/grpc/simple/pb" ./grpc/simple/pb/hello.proto
```
