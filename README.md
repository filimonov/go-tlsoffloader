Introduction
================

Very simple (Thanks to Go!) TCP socket SSL/TLS offloader proxy.
Client can connect to the unsecure socket, and go-tlsoffloader will tunnel the connection to a secure socket.

Quick Start
================

```
go run go-sslterminator.go
clickhouse-client --port 44300 --user explorer
```

Help
================
```
go run go-sslterminator.go --help
  -b string
    	backend address (default "gh-api.clickhouse.tech:9440")
  -l string
    	local address (default "localhost:44300")
```

License
================

Licensed under the New BSD License.

Author
================

Author of go-sslterminator: Uri Shamay (shamayuri@gmail.com)
Rewritten by: filimonov
