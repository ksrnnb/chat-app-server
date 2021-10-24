package response

type Response struct {
	Code   int
	Params map[string]interface{}
}

func (res *Response) IsSuccessful() bool {
	return res.Code >= 200 && res.Code <= 299
}
