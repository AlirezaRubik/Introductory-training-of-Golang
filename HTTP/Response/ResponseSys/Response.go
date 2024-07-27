package responsesys

import (
	"encoding/json"
	"net/http"
)

type ResponseHandler struct {
	Success    bool
	StatusCode int
	Message    interface{}
	Err        error
}

func (R * ResponseHandler)MakeResponseHander(StatucCode int, Message interface{}, Err error, w http.ResponseWriter) {
	ResponseHandler := new(ResponseHandler)
	if StatucCode != http.StatusOK {
		ResponseHandler.StatusCode = StatucCode
	} else {
		ResponseHandler.Success=true
		ResponseHandler.StatusCode = 200
	}
	ResponseHandler.Message = Message
	ResponseHandler.Err = Err
	ResponseJson, err := json.Marshal(ResponseHandler)
	if err != nil {
		ResponseHandler.Success = false
		ResponseHandler.Err=err
	}
    w.WriteHeader(StatucCode)
	w.Write(ResponseJson)
}