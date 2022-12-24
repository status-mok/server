> Not working yet =)

# StatusMOK mock server
This is a mock server written in Go.

## The Goal
The main goal is to make an all-in-one mock server for Web developers and testers which is easy to configure, use and with high throughput.

### Goal decomposition
1. It provides an admin API via REST and GRPC.
2. It runs one or several mock servers on one instance.
3. Mock servers can be created manually or using config files.
4. Mock server configuration can be exported and imported (own format).
5. It supports the most common protocols:
   * HTTP (REST, GraphQL, WebSockets)
   * GRPC
   * Thrift
   * TCP/UDP
6. You can add expectations for each route/rpc based on input values.
7. You can customize response of each route/rpc using:
   * fixed values
   * random values
   * values from imported lists
8. The server logs all the requests and responses. It keeps the most recent values in storage and prints all of them to STDOUT.
9. Configs can be provisioned at server start up, and they will be validated and applied. In case of errors or conflicts during start up the server shall exit with non-zero code. 
10. It supports only own format of configuration files. But it also provides tools (cli or generators) to create config files by converting some common API specification formats: OpenAPI, proto, thrift.
11. It collects metrics of server work-load and received requests, and it provides them in Prometheus format. 
12. The server may use in-memory database or Redis to run as a multi-instance service in cloud.
13. The server shall be distributed as docker image and binaries.

### Custom integrations
1. ngrok - make mock servers publicly available via ngrok reverse-proxy.
