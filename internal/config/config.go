package config

func Config() map[string]string {
	return map[string]string{
		"port":                 ":8080",
		"discount_grpc_server": "localhost:50051",
	}
}
