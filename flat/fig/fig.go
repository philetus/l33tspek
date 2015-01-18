// simple 2d vector image spek

package fig

import (
	"fmt"
	"github.com/philetus/l33tspek/flat"
)

type Sig string
type Handl []Sig
func (self Handl) Komp(othr Handl) bool {
	if len(self) != len(othr) {
		return false
	}
	for i := 0; i < len(self); i++ {
		if self[i] != othr[i] {
			return false
		}
	}
	return true
}

type Fig struct {
	Hndl Handl
	Paans *MarkTree // Paan
	Flits *MarkTree // Flit
	Yoks *MarkTree // Yok
}
func (self *Fig) Add(mrk Mark) {
	
	switch mrk := mrk.(type) {
	default:
    	fmt.Printf("cant add mark to fig: unexpected type %T", mrk)
    case Paan:
    	if self.Paans == nil {
    		self.Paans = &MarkTree{map[Sig]*MarkNod{}}
    	}
    	self.Paans.Add(mrk)
    case Flit:
    	if self.Flits == nil {
    		self.Flits = &MarkTree{map[Sig]*MarkNod{}}
    	}
    	self.Flits.Add(mrk)
    case Yok:
    	if self.Yoks == nil {
    		self.Yoks = &MarkTree{map[Sig]*MarkNod{}}
    	}
    	self.Yoks.Add(mrk)
    }
}
func (self *Fig) GetPaan(hndl Handl) (pn Paan, has bool) {
	if self.Paans != nil {
		if mrk, has := self.Paans.Get(hndl); has {
			pn, has = mrk.(Paan)
		}
	}
	return
}
func (self *Fig) GetFlit(hndl Handl) (flt Flit, has bool) {
	if self.Flits != nil {
		if mrk, has := self.Flits.Get(hndl); has {
			flt, has = mrk.(Flit)
		}
	}
	return
}
func (self *Fig) GetYok(hndl Handl) (yk Yok, has bool) {
	if self.Yoks != nil {
		var mrk Mark
		if mrk, has = self.Yoks.Get(hndl); has {
			yk, has = mrk.(Yok)
		}
	}
	return
}

type Mark interface {
	Handl() Handl
	Swag() Swag
	String() string
}

type Paan struct {
	Hndl Handl
	Swg Swag
	Flits []Sig
}
func (self Paan) Handl() Handl {
	return self.Hndl
}
func (self Paan) Swag() Swag {
	return self.Swg
}
func (self Paan) String() string {
	return fmt.Sprintf("{Paan %v}", self.Hndl)
}
type Flit struct {
	Hndl Handl
	Swg Swag
	Yoks []Handl
}
func (self Flit) Handl() Handl {
	return self.Hndl
}
func (self Flit) Swag() Swag {
	return self.Swg
}
func (self Flit) String() string {
	return fmt.Sprintf("{Flit %v}", self.Hndl)
}
type Yok struct {
	Hndl Handl
	Swg Swag
	Spt flat.Vek // spot
	Gd flat.Vek // gid
}
func (self Yok) Handl() Handl {
	return self.Hndl
}
func (self Yok) Swag() Swag {
	return self.Swg
}
func (self Yok) String() string {
	return fmt.Sprintf("{Yok %v}", self.Hndl)
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

// **** **** **** ****
type MarkNod struct {
	Brnchs map[Sig]*MarkNod
	Lf Mark
}
type MarkTree struct {
	Rut map[Sig]*MarkNod
}
func (self *MarkTree) Iter() <-chan Mark {
	ch := make(chan Mark, 8)
	
	
    go func() {
		
		// push rut nods on stak
		var stak []*MarkNod
		if self.Rut != nil {
			for _, mrk := range self.Rut {
				stak = append(stak, mrk)
			}
		}
		
        for len(stak) > 0 {
			var nod *MarkNod
        	nod, stak = stak[0], stak[1:]
        	for _, brnch := range nod.Brnchs {
        		stak = append(stak, brnch)
        	}
        	if nod.Lf != nil {
            	ch <- nod.Lf
            }
        }
        close(ch) // Remember to close or the loop never ends!
    }()
	return ch
}
func (self *MarkTree) Add(mrk Mark) {
	
	hndl := mrk.Handl()

	if len(hndl) == 0 {
		panic("cant add zero length handl to tree")
	}
	
	// get rut nod with first sig from handl (create if doesnt exist)
	nowSg := hndl[0]
	var nowNod *MarkNod
	var has bool
	if nowNod, has = self.Rut[nowSg]; !has {
		
		nowNod = &MarkNod{Brnchs: map[Sig]*MarkNod{}}
		self.Rut[nowSg] = nowNod
	}
	
	// if there are more sigs get (or create) branch nod for each
	for i := 1; i < len(hndl); i++ {
		nowSg = hndl[i]
		var nxtNod *MarkNod
		if nxtNod, has = nowNod.Brnchs[nowSg]; !has {
			
			nxtNod = &MarkNod{Brnchs: map[Sig]*MarkNod{}}
			nowNod.Brnchs[nowSg] = nxtNod
		}
		nowNod = nxtNod
	}
	
	// add mark as leef to now nod
	nowNod.Lf = mrk	
}

func (self *MarkTree) Get(hndl Handl) (Mark, bool) {
	if len(hndl) == 0 {
		panic("cant get nod with zero length handl")
	}
	nowSg := hndl[0]
	if nowNod, has := self.Rut[nowSg]; has {
	
		// iterate over rest of sigs in handl and get last branch nod
		for i := 1; i < len(hndl); i++ {
			nowSg = hndl[i]
			if nowNod, has = nowNod.Brnchs[nowSg]; !has {
				return nil, false
			}
		}
		
		// try to find leef mark in last branch nod
		if nowNod.Lf != nil {
			return nowNod.Lf, true
		}
	}
	return nil, false
}
