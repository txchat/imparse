module github.com/txchat/imparse

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/txchat/im v0.0.1
	google.golang.org/protobuf v1.28.1
)

replace (
	github.com/txchat/im => ../im
)