package grpc

// Config is the configuration for the gRPC server.
type Config struct {
	Port uint16 `env:"GRPC_SERVER_PORT" envDefault:"50051"`
}
