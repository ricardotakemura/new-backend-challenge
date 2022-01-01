package config

import "os"

func Config() map[string]string {
	var port = os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	}
	var blackFridayDay = os.Getenv("BLACK_FRIDAY_DAY")
	if len(blackFridayDay) == 0 {
		blackFridayDay = "12-31"
	}
	return map[string]string{
		"port":                 port,
		"blackFridayDay":       blackFridayDay,
		"discount_grpc_server": "localhost:50051",
	}
}
