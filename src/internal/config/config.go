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
	var discountServer = os.Getenv("DISCOUNT_SERVER")
	if len(discountServer) == 0 {
		discountServer = "new-backend-challenge_discount_1:50051"
	}
	return map[string]string{
		"port":                 port,
		"blackFridayDay":       blackFridayDay,
		"discount_grpc_server": discountServer,
	}
}
