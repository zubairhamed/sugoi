package sugoi
import (
	"net/http"
	"encoding/json"
	"log"
)

func InternalServerError(msg ... string) HttpCode {
	var content string

	if len(msg) > 0 {
		content = msg[0]
	} else {
		content = "500 - Internal Server Error"
	}

	return HttpCode{
		code: http.StatusInternalServerError,
		content: content,
	}
}

func NotFound(msg ... string) HttpCode {
	var content string

	if len(msg) > 0 {
		content = msg[0]
	} else {
		content = "404 - Not Found"
	}

	return HttpCode{
		code: http.StatusNotFound,
		content: content,
	}
}

func SendHttpResponse(response interface{}, w http.ResponseWriter) {
	ResponseHandler(response, w)
}

func SendHttpCodeResponse(codeResponse HttpCode, w http.ResponseWriter) {
	w.WriteHeader(codeResponse.GetCode())

	if codeResponse.GetContent() != "" {
		w.Write([]byte(codeResponse.GetContent()))
	}
}

func ResponseHandler(response interface{}, w http.ResponseWriter) {
	if val, ok := response.(string); ok {
		w.Write([]byte(val))
	} else
	if val, ok := response.(int); ok {
		log.Println("int", val)
	} else
	if val, ok := response.(HttpCode); ok {
		SendHttpCodeResponse(val, w)
	} else {
		log.Println("Handling Object >> JSON", response)
		b, err := json.Marshal(response)

		log.Println(err)
		if err != nil {
			errorHttpCode := HttpCode{
				code: 500,
				content: "An error occured processing request",
			}
			SendHttpCodeResponse(errorHttpCode, w)
		} else {
			w.Write(b)
		}
	}
}


type HttpCode struct {
	code 	int
	content string
}

func (h *HttpCode) GetCode() int {
	return h.code
}

func (h *HttpCode) GetContent() string {
	return h.content
}

