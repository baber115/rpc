# 如何生成代码

```shell
# 在当前目录下
protoc -I=. --go_out=./pb --go_opt=module="gitee.com/infraboard/go-course/day21/pb" ./hello.proto

-I：-IPATH, --proto_path=PATH, 指定proto文件搜索的路径, 如果有多个路径 可以多次使用-I 来指定, 如果不指定默认为当前目录
--go_out: --go指插件的名称, 我们安装的插件为: protoc-gen-go, 而protoc-gen是插件命名规范, go是插件名称, 因此这里是--go, 而--go_out 表示的是 go插件的 out参数, 这里指编译产物的存放目录
--go_opt: protoc-gen-go插件opt参数, 这里的module指定了go module, 生成的go pkg 会去除掉module路径，生成对应pkg
pb/hello.proto: 我们proto文件路径
```