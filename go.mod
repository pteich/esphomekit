module github.com/pteich/esphomekit

go 1.14

require (
	github.com/brutella/hc v1.2.2
	github.com/golang/protobuf v1.3.5
	github.com/lucasb-eyer/go-colorful v1.0.3
	github.com/miekg/dns v1.1.31 // indirect
	github.com/pteich/configstruct v1.1.0
	github.com/pteich/go-timeout-httpclient v0.0.0-20200110111718-916aff4d9c82
	github.com/pteich/logger v1.1.2
	github.com/xiam/to v0.0.0-20200126224905-d60d31e03561 // indirect
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/net v0.0.0-20201009032441-dbdefad45b89 // indirect
	golang.org/x/sys v0.0.0-20201009025420-dfb3f7c4e634 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/grpc v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
