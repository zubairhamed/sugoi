package sugoi
import (
	"net/http"
	"encoding/json"
	"text/template"
	"strconv"
)

func msgContent(defaultMsg string, msg ... string) string {
	var content string

	if len(msg) > 0 {
		content = msg[0]
	} else {
		content = defaultMsg
	}
	return content
}

func InternalServerError(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusInternalServerError,
		content: msgContent("500 - Internal Server Error", msg...),
	}
}

func NotFound(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusNotFound,
		content: msgContent("404 - Not Found", msg...),
	}
}

func OK(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusOK,
		content: msgContent("200 - OK", msg...),
	}
}

func NoContent(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusNoContent,
		content: msgContent("204 - No Content", msg...),
	}
}

func Accepted(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusAccepted,
		content: msgContent("202 - Accepted", msg...),
	}
}

func ServiceUnavailable(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusServiceUnavailable,
		content: msgContent("503 - Service Unavailable", msg...),
	}
}

func BadRequest(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusBadRequest,
		content: msgContent("400 - Bad Request", msg...),
	}
}

func Unauthorized(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusUnauthorized,
		content: msgContent("401 -Unauthorized", msg...),
	}
}

func Forbidden(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusForbidden,
		content: msgContent("403 - Forbidden", msg...),
	}
}

func MethodNotAllowed(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusMethodNotAllowed,
		content: msgContent("405 - Not Allowed", msg...),
	}
}

func NotImplemented(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusNotImplemented,
		content: msgContent("501 - Not Implemented", msg...),
	}
}

func NotModified(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusNotModified,
		content: msgContent("304 - Not Modified", msg...),
	}
}

func UnsupportedMediaType(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusUnsupportedMediaType,
		content: msgContent("415 - Unsupported Media Type", msg...),
	}
}

func Conflict(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusConflict,
		content: msgContent("409 - Conflict", msg...),
	}
}

func NotAcceptable(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusNotAcceptable,
		content: msgContent("406 - Not Acceptable", msg...),
	}
}

func Created(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusCreated,
		content: msgContent("201 - Created", msg...),
	}
}

func Gone(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusGone,
		content: msgContent("410 - Gone", msg...),
	}
}

func Found(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusFound,
		content: msgContent("302 - Found", msg...),
	}
}

func MovedPermanently(msg ... string) HttpCode {
	return HttpCode{
		code: http.StatusMovedPermanently,
		content: msgContent("301 - Moved Permanently", msg...),
	}
}

//func SendHttpCodeResponse(codeResponse HttpCode, w http.ResponseWriter) {
//	w.WriteHeader(codeResponse.GetCode())
//
//	if codeResponse.GetContent() != "" {
//		w.Write([]byte(codeResponse.GetContent()))
//	}
//}

func ResponseHandler(response Response, w http.ResponseWriter) {
	content := response.content
	httpCode := response.httpCode
	w.WriteHeader(httpCode)

	if val, ok := content.(string); ok {
		w.Write([]byte(val))
	} else
	if val, ok := content.(int); ok {
		w.Write([]byte(strconv.Atoi(val)))
	} else
	if val, ok := content.(*HtmlContent); ok {
		tpl := template.New(val.tpl)
		t, err := tpl.ParseFiles(val.tpl)
		if err != nil {
			ise := InternalServerError(err.Error())
			SendResponse(InternalServerError(err.Error()), 500, w)
		}
		err = t.Execute(w, val.model)
	} else {
		b, err := json.Marshal(content)

		if err != nil {
			ise := InternalServerError(err.Error())
			SendResponse(InternalServerError(err.Error()), 500, w)
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

type Response struct {
	httpCode 	int
	content 	interface{}
}

func Html(tpl string, model interface{}) *HtmlContent {
	return &HtmlContent{
		tpl: tpl,
		model: model,
	}
}

type HtmlContent struct {
	tpl 	string
	model 	interface{}
}
