module day12_ectd

go 1.16

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc v1.37.0 => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/prometheus/client_golang v1.10.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210503195802-e9a32991a82e // indirect
	google.golang.org/grpc v1.37.0 // indirect
)
