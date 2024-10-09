# Server

环境变量模版文件 [template](../template.env)

启动服务

```shell
dotenvx run -f .dev.env -- go run ./garage_server
```

```bash
swag init -d ./garage_server/router -g init.go --outputTypes yaml --pdl 3
```
