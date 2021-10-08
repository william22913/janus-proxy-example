package error_model

type ErrorModel struct{
	ErrorCode int
	Message string
	CausedBy error
}

func GenerateNonErrorModel()ErrorModel{
	return ErrorModel{
		ErrorCode: 200,
		Message:   "Success",
		CausedBy:  nil,
	}
}

func GenerateUnsupportedServiceModel()ErrorModel{
	return ErrorModel{
		ErrorCode: 400,
		Message:   "Unsupported Service",
		CausedBy:  nil,
	}
}

func GenerateJSONInvalidModel()ErrorModel{
	return ErrorModel{
		ErrorCode: 400,
		Message:   "Invalid JSON",
		CausedBy:  nil,
	}
}

func GenerateEmptyFieldError(fieldName string)ErrorModel{
	return ErrorModel{
		ErrorCode: 400,
		Message:   "Empty Field " + fieldName,
		CausedBy:  nil,
	}
}

func GenerateUnknownErrorModel(causedBy error)ErrorModel{
	return ErrorModel{
		ErrorCode: 500,
		Message:   "Unknown Error",
		CausedBy:  causedBy,
	}
}