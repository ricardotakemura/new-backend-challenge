package error

type ErrorService struct {
	errorModel *ErrorModel
}

var _errorService *ErrorService

func GetErrorService() *ErrorService {
	if _errorService == nil {
		_errorService = NewErrorService()
	}
	return _errorService
}

func NewErrorService() *ErrorService {
	return &ErrorService{errorModel: NewErrorModel()}
}

func (service ErrorService) GetById(id string, lang string, params map[string]string) Error {
	return (*service.errorModel).GetById(id, lang, params)
}

func (service ErrorService) PRODUCT_NOT_FOUND(lang string, productId string) Error {
	var params = map[string]string{}
	params["ProductId"] = productId
	return service.GetById("product_not_found", lang, params)
}

func (service ErrorService) PRODUCT_ALREADY_IN_THE_CART(lang string, productId string) Error {
	var params = map[string]string{}
	params["ProductId"] = productId
	return service.GetById("product_already_in_the_cart", lang, params)
}

func (service ErrorService) INVALID_QUANTITY(lang string) Error {
	return service.GetById("invalid_quantity", lang, nil)
}

func (service ErrorService) INSUFFICIENT_STOCK(lang string) Error {
	return service.GetById("insufficient_stock", lang, nil)
}

func (service ErrorService) INVALID_BODY(lang string) Error {
	return service.GetById("invalid_body", lang, nil)
}

func (service ErrorService) INVALID_PATH(lang string) Error {
	return service.GetById("invalid_path", lang, nil)
}

func (service ErrorService) GENERIC_ERROR(lang string) Error {
	return service.GetById("generic_error", lang, nil)
}
