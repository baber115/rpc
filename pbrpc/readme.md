# 如何生成代码

# 生成service pb编译文件
```shell
protoc -I=. --go_out=./ ./hello.proto
```

# 生成codec的protobuf编译文件
```shell
 protoc -I=./codec/pb/ --go_out=./codec/pb/ --go_opt=module="rpc-g7/pbrpc/codec/pb" ./codec/pb/header.proto
```