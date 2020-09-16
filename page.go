package wand

type Page struct {
	Current int
	Num     int
}

func (p *Page) Limit() (int, int) {
	if p.Current < 1 {
		p.Current = 1
	}
	if p.Num < 1 {
		p.Num = 10
	}
	offset := (p.Current - 1) * p.Num
	return p.Num, offset
}
