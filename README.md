# introduce
dockerコンテナの中でdelve debuggerを回せるための設定を一番簡単にまとめる。

# Resources
## Image
- golang1:22
## Delve
```shell
go install github.com/go-delve/delve/cmd/dlv@latest
```
## Port Mapping
- delve: `2345:2345`
- server: `8081:8080`
- 詳しくは`docker-compose.yml`

# Create Docker Container
```shell
docker-compose up --build -d
```

# Testing
- breakpointを好きなところ設定した、後以下のコマンドを実施させる
    ```shell
    curl 'http://localhost:8081/unix-time?date=2024-02-14&time=11:11:11'
    ```
- 結果
    ```json
    {"unix_time":1707909071}
    ```
- breakpointで止まるか確認する。

# Settings: Goland
1. Run > Edit Configurations...
2. ＋　> Go remote
3. host: localhost / port: 2345



