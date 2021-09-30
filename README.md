# Exemplo de RPC em Golang #
Repositório com o objetivo de compartilhar um exemplo de RPC com Golang

```sh
git clone github.com/andrebetiolo/grpc-rpc
```
## CLI ##

Fiz uma cli com `make` para ajudar

```sh
make
```

```sh
help                           Show this help
run-client                     Run the client
run-server                     Run the server
```

Pastas

```sh
.
├── client
│   └── main.go # Exemplo de um cliente se comunicando com o servidor
├── go.mod
├── Makefile # Arquivo do make que faço a CLI
└── server
    └── main.go # Exemplo de criação de um servidor RPC em Golang
```