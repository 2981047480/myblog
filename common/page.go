package common

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize: 10,
		PageNum:  1,
	}
}

type PageRequest struct {
	PageSize int
	PageNum  int
}

func (p *PageRequest) OffSet() int {
	return (p.PageNum - 1) * p.PageSize
}
