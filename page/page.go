package page

import "math"

type Page struct {
	Size int //页大小
	Page int //当前页
	Count int //总页数
	Total int //总记录
	OffSet int //
}

func (p *Page)SetTotal(total int) {
	p.Total = total
	if p.Size == 0 {
		p.Size = 10
	}
	count := int(math.Ceil(float64(p.Total / p.Size)))
	p.Count = count
}

func (p *Page)GetCount() int {

	if p.Count > 0 {
		return p.Count
	}

	if p.Size == 0 {
		p.Size = 10
	}
	count := int(math.Ceil(float64(p.Total / p.Size)))

	p.Count = count

	return count
}

func (p *Page)Offset()int{
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Page > p.Count {
		p.Page = p.Count
	}

	p.OffSet = (p.Page - 1) * p.Size

	return p.OffSet
}

func (p *Page)Limit() int {
	return p.Size
}

func Test() {

}