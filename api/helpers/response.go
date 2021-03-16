package helpers

type response struct {
	statusCode int16
	message    string
	data       interface{}
	err        error
}

// func ResponseBody(statusCode int16, message string, data interface{}, err error) interface{} {

// }
