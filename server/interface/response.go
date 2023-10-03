package interfaces

type Response interface {
	Marshal()
	Status() int
	Message() string
}

type BaseResponse struct{}

func (res BaseResponse) Marshal() {}

func (res BaseResponse) Status() int {
	return 200
}

func (res BaseResponse) Message() string {
	return ""
}

type ErrorBaseResponse struct {
	BaseResponse
}

func (res ErrorBaseResponse) Status() int {
	return 500
}

type EmptyRequest struct{}
