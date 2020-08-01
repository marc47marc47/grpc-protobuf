



# grpc-golang-server

sync from Building an Basic API with gRPC and Protobuf

https://www.youtube.com/watch?v=Y92WWaZJl24

protobuff are used for faster communications between consumers and providers.

To generate code form proto config 
 - protoc --proto_path=proto --go_out=plugins=grpc:proto proto\service.proto

Run the both client and Provider.

```bash
sh make.sh
server.exe
client.exe
```

test in browser url:
http://localhost:8080/add/2/3
