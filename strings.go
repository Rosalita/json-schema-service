package main

const (
	// Mongo strings
	validationDbName = "validation_service"
	schemaCollection = "schemas"

	// Error strings
	errMsgSchemaNotFound = "schema not found"
	errMsgInternalError  = "internal error"
	errMsgDatabaseError  = "database error"

	// Action strings
	actionUpload   = "uploadSchema"
	actionValidate = "validateDocument"

	//Status strings
	statusSuccess = "success"
	statusError   = "error"

	//Message strings
	msgInvalidJson = "Invalid JSON"
)
