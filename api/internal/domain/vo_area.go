package domain

import "errors"

type Area struct {
	lx float64
	rx float64
	ty float64
	by float64
}

// NewAreaはバリデーション付きでAreaを生成
func NewArea(lx, rx, ty, by float64) (Area, error) {
	if lx >= rx || by >= ty {
		return Area{}, errors.New("invalid area coordinates")
	}
	return Area{lx: lx, rx: rx, ty: ty, by: by}, nil
}

func (a Area) LX() float64 { return a.lx }
func (a Area) RX() float64 { return a.rx }
func (a Area) TY() float64 { return a.ty }
func (a Area) BY() float64 { return a.by }
