package errors

//StatusMsgError represents the status and message of the error response
type StatusMsgError struct {
	Status  int
	Message string
}

//ErrorResponse represents all errors response from the API
var ErrorResponse = map[string]StatusMsgError{
	"notAuthorized": {
		Status:  401,
		Message: "Unauthorized user",
	},
	"postNotExist": {
		Status:  404,
		Message: "Post does not exist",
	},
	"categoryAlreadyExists": {
		Status:  409,
		Message: "Category already exists",
	},
	"categoryNotFound": {
		Status:  400,
		Message: `"categoryIds" not found`,
	},
	"nameIsRequired": {
		Status:  400,
		Message: `"name" is required`,
	},
	"userNotExist": {
		Status:  404,
		Message: "User does not exist",
	},
	"userAlreadyExists": {
		Status:  409,
		Message: "User already registered",
	},
	"missingFields": {
		Status:  400,
		Message: "Some required fields are missing",
	},
	"invalidFields": {
		Status:  400,
		Message: "Invalid fields",
	},
	"emailIsRequired": {
		Status:  400,
		Message: `"email" must be a valid email`,
	},
	"invalidPassword": {
		Status:  400,
		Message: `"password" length must be at least 6 characters long`,
	},
	"passwordIsRequired": {
		Status:  400,
		Message: `"password" is required`,
	},
	"tokenNotFound": {
		Status:  401,
		Message: "Token not found",
	},
	"invalidToken": {
		Status:  401,
		Message: "Expired or invalid token",
	},
	"displayNameIsRequired": {
		Status:  400,
		Message: `"displayName" is required`,
	},
	"displayNameInvalid": {
		Status:  400,
		Message: `"displayName" length must be at least 8 characters long`,
	},
}
