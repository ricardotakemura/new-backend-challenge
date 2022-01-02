package error

import (
	"testing"
)

type MockErrorModel struct {
}

func (model MockErrorModel) GetById(id string, lang string, params map[string]string) Error {
	switch id {
	case "product_not_found":
		return Error{Code: "0001", Message: "Product not found."}
	case "invalid_quantity":
		return Error{Code: "0002", Message: "Invalid quantity."}
	case "invalid_body":
		return Error{Code: "0003", Message: "Invalid body."}
	case "invalid_path":
		return Error{Code: "0004", Message: "Invalid path."}
	case "product_already_in_the_cart":
		return Error{Code: "0005", Message: "Product already in the cart."}
	default:
		return Error{Code: "0000", Message: "Internal server error."}
	}
}

func Test_ErrorService_GENERIC_ERROR_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.GENERIC_ERROR("en").Code != "0000" {
		t.Errorf("errorService.GENERIC_ERROR.Code: expected %s obtained %s", "0000", errorService.GENERIC_ERROR("en").Code)
	}
}

func Test_ErrorService_PRODUCT_NOT_FOUND_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.PRODUCT_NOT_FOUND("en", "1").Code != "0001" {
		t.Errorf("errorService.PRODUCT_NOT_FOUND.Code: expected %s obtained %s", "0002", errorService.PRODUCT_NOT_FOUND("en", "1").Code)
	}
}

func Test_ErrorService_INVALID_QUANTITY_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.INVALID_QUANTITY("en").Code != "0002" {
		t.Errorf("errorService.INVALID_QUANTITY.Code: expected %s obtained %s", "0002", errorService.INVALID_QUANTITY("en").Code)
	}
}

func Test_ErrorService_INVALID_BODY_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.INVALID_BODY("en").Code != "0003" {
		t.Errorf("errorService.INVALID_BODY.Code: expected %s obtained %s", "0003", errorService.INVALID_BODY("en").Code)
	}
}

func Test_ErrorService_PRODUCT_ALREADY_IN_THE_CART_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.PRODUCT_ALREADY_IN_THE_CART("en", "1").Code != "0005" {
		t.Errorf("errorService.PRODUCT_ALREADY_IN_THE_CART.Code: expected %s obtained %s", "0005", errorService.PRODUCT_ALREADY_IN_THE_CART("en", "1").Code)
	}
}

func Test_ErrorService_INVALID_PATH_Ok(t *testing.T) {
	errorService := ErrorService{ErrorModel: MockErrorModel{}}
	if errorService.INVALID_PATH("en").Code != "0004" {
		t.Errorf("errorService.INVALID_PATH.Code: expected %s obtained %s", "0004", errorService.INVALID_PATH("en").Code)
	}
}
