// Author: xufei
// Date: 2019/4/30

package car

type BenzBuilder struct {
	*Benz
}

func NewBenzBuilder() Builder {
	return &BenzBuilder{
		Benz: &Benz{},
	}
}

func (b *BenzBuilder) Color(color Color) Builder {
	b.Benz.Color = color
	return b
}

func (b *BenzBuilder) Wheels(w Wheels) Builder {
	b.Benz.Wheels = w
	return b
}

func (b *BenzBuilder) TopSpeed(s Speed) Builder {
	b.Benz.Speed = s
	return b
}

func (b *BenzBuilder) Build() Interface {
	return b.Benz
}
