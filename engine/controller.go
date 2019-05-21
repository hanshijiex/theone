package engine

type IController interface {
	Run(req *Request, res *Response) error
}