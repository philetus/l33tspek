// simple 2d vector image spek

package fig

import (
	"fmt"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/tag"
)

type Fig struct {
	Dfl tag.Dufl
	PnBg tag.DuflBag
	FltBg tag.DuflBag
	YkBg tag.DuflBag
}
func (self *Fig) Skrib(mrk Mark) {
	
	switch mrk := mrk.(type) {
	default:
    	fmt.Printf("cant add mark to fig: unexpected type %T", mrk)
    case Paan:
    	self.PnBg.Stass(mrk)
    case Flit:
    	self.FltBg.Stass(mrk)
    case Yok:
    	self.YkBg.Stass(mrk)
    }
}
func (self *Fig) FekkPaan(dfl tag.Dufl) (pn Paan, has bool) {
	if dflr, hs := self.PnBg.Fekk(dfl); hs {
		pn, has = dflr.(Paan)
	}
	return
}
func (self *Fig) FekkFlit(dfl tag.Dufl) (flt Flit, has bool) {
	if dflr, hs := self.FltBg.Fekk(dfl); hs {
		flt, has = dflr.(Flit)
	}
	return
}
func (self *Fig) FekkYok(dfl tag.Dufl) (yk Yok, has bool) {
	if dflr, hs := self.YkBg.Fekk(dfl); hs {
		yk, has = dflr.(Yok)
	}
	return
}

type Mark interface {
	Dufl() tag.Dufl
	Swag() Swag
	String() string
}

type Paan struct {
	Dfl tag.Dufl
	Swg Swag
	Flits []tag.Dufl
}
func (self Paan) Dufl() tag.Dufl {
	return self.Dfl
}
func (self Paan) Swag() Swag {
	return self.Swg
}
func (self Paan) String() string {
	return fmt.Sprintf("{Paan %v}", self.Dfl)
}
type Flit struct {
	Dfl tag.Dufl
	Swg Swag
	Yoks []tag.Dufl
}
func (self Flit) Dufl() tag.Dufl {
	return self.Dfl
}
func (self Flit) Swag() Swag {
	return self.Swg
}
func (self Flit) String() string {
	return fmt.Sprintf("{Flit %v}", self.Dfl)
}
type Yok struct {
	Dfl tag.Dufl
	Swg Swag
	Spt flat.Vek // spot
	Gd flat.Vek // gid
}
func (self Yok) Dufl() tag.Dufl {
	return self.Dfl
}
func (self Yok) Swag() Swag {
	return self.Swg
}
func (self Yok) String() string {
	return fmt.Sprintf("{Yok %v}", self.Dfl)
}

type Swag struct {
	Vz bool
	Dpth float64
	Wt float64
	Klr [4]float64 // r g b a
}
func (self Swag) Komp(othr Swag) bool {
	if self.Vz != othr.Vz {
		return false
	}
	if self.Dpth != othr.Dpth {
		return false
	}
	if self.Wt != othr.Wt {
		return false
	}
	if self.Klr != othr.Klr {
		return false
	}
	return true
}

