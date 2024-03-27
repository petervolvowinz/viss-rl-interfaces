module github.com/petervolvowinz/viss-rl-interfaces

go 1.22.1

replace github.com/petervolvowinz/viss-rl-interfaces/broker => ./broker

replace github.com/petervolvowinz/viss-rl-interfaces/proto_files => ./base

require (
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
)
