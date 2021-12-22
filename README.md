# WMB GRPC Server

## Protobuf compile

```shell
protoc --go_out=./api --go-grpc_out=./api api/omzet.proto
```

## Build Project
```shell
GOOS=linux GOARCH=386 go build -o build/opo-server enigmacamp.com/opo
```

## Run Background
```shell
 GRPC_HOST=0.0.0.0 GRPC_PORT=9999 nohup ./omzet-server &
```
