# APIs

## 引入/更新 submodule
1. 初始化本地子模块配置文件
    ```shell
    git submodule init
    ```

2. 更新项目，抓取子模块内容。
    ```shell
    git submodule update
    ```

## 生成代码
```shell
make zapis
```

## Generate

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/envoyproxy/protoc-gen-validate@latest
```

