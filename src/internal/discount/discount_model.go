package discount

import (
	context "context"
	"log"
	"new-backend-challenge/internal/config"

	"google.golang.org/grpc"
)

type IDiscountModel interface {
	GetDiscount(productId int32) float32
}

type DiscountModel struct {
}

func NewDiscountModel() *DiscountModel {
	return &DiscountModel{}
}

func (model DiscountModel) GetDiscount(productId int32) float32 {
	ctx := context.Background()
	conn, err := grpc.Dial(config.Config()["discount_grpc_server"], grpc.WithInsecure())
	if err != nil {
		log.Printf("Connect failed: %s", err.Error())
	}
	defer conn.Close()
	client := NewDiscountClient(conn)
	if client != nil {
		response, err := client.GetDiscount(ctx, &GetDiscountRequest{ProductID: productId})
		if err != nil {
			log.Printf("Connect failed: %s", err.Error())
			return 0
		}
		return response.GetPercentage()
	}
	return 0
}
