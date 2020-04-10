module github.com/pteich/esphomekit

go 1.13

require (
	github.com/brutella/hc v1.2.2-0.20200406141821-64571bae3f05
	github.com/golang/protobuf v1.3.5
	github.com/lucasb-eyer/go-colorful v1.0.3
	github.com/pteich/configstruct v1.1.0
	github.com/pteich/go-timeout-httpclient v0.0.0-20200110111718-916aff4d9c82
	github.com/pteich/logger v1.1.2
	google.golang.org/grpc v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
