// 2d hok geometry model

package fig

import (
	//"fmt"
	//"reflect"
	"github.com/philetus/l33tspek/klv"
)

//
// *** fig implements hok for 2d geometry models
//
type Fig struct {
	S klv.Sig
	A klv.Hok
	Kids map[klv.Sig]klv.Hok
	Nuks map[klv.Sig]klv.Nuk
	// deks["Joint"] --> Warp
}
func (self Fig) Sig() klv.Sig {
	return self.S
}
func (self Fig) Ark() (klv.Hok, bool) {
	if self.A == nil {
		return nil, false
	}
	return self.A, true
}
func (self Fig) HasKids() bool {
	return self.kids != nil
}
func (self Fig) Kid(sig klv.Sig) (kid klv.Hok, has bool) {
	if !self.HasKids() {
		return nil, false
	}
	kid, has = self.Kids[sig]
	return
}
func (self Fig) Jnt(kid klv.Hok) klv.Hok {
	if !self.HasKids() {
		self.kids = make(map[klv.Sig]klv.Hok)
	}
	self.kids[kid.Sig()] = kid
	return kid
}
func (self Fig) Nuk(sig klv.Sig) (nuk klv.Nuk, has bool) {
	if !self.HasNuks() {
		return nil, false
	}
	nuk, has = self.deks[sig]
	return
}
func (self Fig) Swlo(nuk klv.Nuk) klv.X {
	if !self.HasNuks() {
		self.deks = make(map[klv.Sig]klv.Nuk)
	}
	sig := nuk.Sig()
	self.deks[sig] = nuk
	return klv.Xd{S: sig}
}

// ****************
// *** fig nuks
// ****************

// *** delta
type Delta struct {
	Sg klv.Sig
	Cmps [2]float64
}
func (self Delta) Sig() klv.Sig {
	return self.Sg
}

// *** warps
type LatWarp struct {
	S klv.Sig
	Lat klv.X
}
func (self LatWarp) Sig() klv.Sig {
	return self.S
}
type FlktWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Sym klv.X
}
func (self FlktWarp) Sig() klv.Sig {
	return self.S
}
type RotWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Hed klv.X
}
func (self RotWarp) Sig() klv.Sig {
	return self.S
}
type SkalWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Skal klv.X
}
func (self SkalWarp) Sig() klv.Sig {
	return self.S
}
type CmboWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Wrps [2]klv.X
}
func (self CmboWarp) Sig() klv.Sig {
	return self.S
}

// gid
type Gid struct {
	Sg klv.Sig
	Dlta klv.X
}
func (self Gid) Sig() klv.Sig {
	return self.Sg
}

// pins
type LatPin struct {
	Sg klv.Sig
	Dlta klv.X
}
func (self DeltaPin) Sig() klv.Sig {
	return self.Sg
}

type SektPin struct {
	Sg klv.Sig
	Gids [2]klv.X
}
func (self SektPin) Sig() klv.Sig {
	return self.Sg
}


// *** flit
type Flit struct {
	klv.Guk // provides nal func
	S klv.Sig
	Yoks []klv.X
}
func (self Flit) Sig() klv.Sig {
	return self.S
}

