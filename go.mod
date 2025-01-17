module github.com/interchained/gpn

go 1.16

require (
	github.com/interchained/cosmos-sdk v0.43.1 // bump
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/interchained/cosmos-proto v0.3.2 // bump
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/interchained/gpm v0.0.0-20210625155357-5a2c8d79013b
	github.com/interchained/genesismint v0.34.13 // bump
	github.com/interchained/gm-db v0.6.6 // bump 
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/interchained/protobuf v1.3.3-alpha.regen.1
