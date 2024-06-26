package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize: 10,
		PageNum:  1,
	}
}

func NewPageRequestFromGin(c *gin.Context) *PageRequest {
	p := NewPageRequest()

	Psstr := c.Query("page_size")
	Pnstr := c.Query("page_number")

	if Psstr != "" {
		ps, _ := strconv.Atoi(Psstr)
		if ps != 0 {
			p.PageSize = ps
		}
	}

	if Pnstr != "" {
		pn, _ := strconv.Atoi(Pnstr)
		if pn != 0 {
			p.PageNum = pn
		}
	}
	return p
}

type PageRequest struct {
	PageSize int
	PageNum  int
}

func (p *PageRequest) OffSet() int {
	return (p.PageNum - 1) * p.PageSize
}
