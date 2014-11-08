// geometry for l33tspek

package fig

import (
	//"fmt"
	//"reflect"
)

// unique name for tracking instances
type Sig uint

func makeKrystr(nxt uint) func() Sig {
       return func() Sig {
               sig := Sig(nxt)
               nxt++
               return sig
       }
}

// implements flake interface
type Ting struct {
	sig Sig
}
func (self *Ting) Kryst() {
       self.sig = krystr() // use closure to gen unique id
}

// every ting is a special snowflake with a unique sig ;)
type Flake interface {
	Sig() Sig
}
func (self *Ting) Sig() Sig {
	return self.sig
}

type Ava struct {
	sigs []Sig
}
// sig
func (self *Ava) Peek() (sg Sig, last bool) {
	switch len(self.sigs) {
	case 0: // wtf?
		sg, last = nil, true
	case 1: // this is sig of geometry
		sg = self.sigs[0]
		last = true
	default: // this is sig of containing hok
		sg = self.sigs[0]
		last = false
	}
}
// returns new ava without the top sig to pass to hok kid
func (self *Ava) Spawn() Ava {
	return Ava{self.sigs[1:]}
}

// represents bezier curve for eezl
type Bez struct {
	Coords [2]Vek
	Gids [2]Vek
}

// position in 2d plane; implements spot
type Pin struct {
	Ting
	coords Vek
}

// edge between two pins
type Ra struct {
	Ting
	ends [2]PinTr
}


// a (possibly closed) path across a series of ras sharing pins
type Flit struct {
	Ting
	edges []RaTr
}

type Fig struct {
	Ting
	marks map[Sig]Ting
}
func (self *Fig) Mark(tng Ting) {
	self.marks[tng.Sig()] = tng
}
func (self *Fig) Dox(av Ava) (tng Ting, ok bool) {
	sigs = av.Sigs()
	if len(sigs) == 1 {
		tng, ok = self.marks[sigs[0]]
		return
	} 
	tng, ok = self.marks[sigs[0]]
	if ok {
		hk = Hok(tng)
		tng, ok = hk.Dox(Tar{sigs=sigs[1:]})
	}
}

type Hok struct {
	Fig
	rent Fig
	palmps map[Sig]Sig
	joint Warp
}
func (self *Hok) Palmp(trgt Sig, tng Ting) {
	self.marks[tng.Sig()] = tng
	self.palmps[trgt Sig] = tng.Sig()
}
func (self *Hok) Dox(av Ava) (tng Ting, ok bool) {
	sigs = av.Sigs()
	// check if ava is palmpd
	if len(sigs) == 1 {
		tng, ok = self.palmps[sigs[0]]
		if ok {
			tng, ok = self.marks[sigs[0]]
			return
		} 
	} else {
		tng, ok = self.palmps[sigs[0]]
		if ok {
			tng, ok = self.marks[sigs[0]]
			if ok {
				hk = Hok(tng)
				tng, ok = hk.Dox(Tar{sigs=sigs[1:]})
				return
			}
		}
	}
		tng, ok = self.marks[sigs[0]]
		if ok {
			hok = Hok(tng)
			tng, ok = hok.Dox(Tar{sigs=sigs[1:]})
		}
	}
}
