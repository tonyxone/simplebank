module github.com/tonyxone/simplebank

go 1.18

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.6
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
)

require github.com/dgrijalva/jwt-go v3.2.0+incompatible

require github.com/o1egl/paseto v1.0.0

require (
	github.com/go-playground/validator/v10 v10.11.0
	github.com/goccy/go-json v0.9.10 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/rs/zerolog v1.28.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.48.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.28.1
)
