package util

type HandleResult struct {
	Result    interface{}
	IsSuccess bool
	Error     error
}

func (h *HandleResult) GetResult() interface{} {
	return h.Result
}

func (h *HandleResult) GetIsSuccess() bool {
	return h.IsSuccess
}

func (h *HandleResult) GetError() error {
	return h.Error
}
