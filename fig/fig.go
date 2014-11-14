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
	kids map[klv.Sig]klv.Hok
	deks map[klv.Sig]klv.Nuk
	// deks["Joint"] --> Warp
}
func (self Fig) Sig() klv.Sig {
	return self.S
}
func (self Fig) Ark() klv.Hok {
	return self.A
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
func (self Fig) Nuk(sig Sig) (nuk Nuk, has bool) {
	if !self.HasNuks() {
		return nil, false
	}
	nuk, has = self.deks[sig]
	return
}
func (self Fig) Swlo(nuk Nuk) X {
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

// *** skalr
type Skalr struct {
	S klv.Sig
	N klv.Nal // float64 value
}
func (self Skalr) Sig() klv.Sig {
	return self.S
}
func (self Skalr) Nal() (klv.Nal, bool) {
	return self.N, true
}

// *** neg - negates kid value
type Neg struct {
	klv.Guk // provides nal func
	S klv.Sig
	Valu klv.X // value
}
func (self Neg) Sig() klv.Sig {
	return self.S
}

// *** vek
type Vek struct {
	klv.Guk // provides nal func
	S klv.Sig
	V klv.X
	U klv.X
}
func (self Vek) Sig() klv.Sig {
	return self.S
}

// *** warps
type LatWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Dlta klv.X // delta vector
}
func (self LatWarp) Sig() klv.Sig {
	return self.S
}
type FlktWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Sym klv.X // line of symmetry 
}
func (self FlktWarp) Sig() klv.Sig {
	return self.S
}
type RotWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Hed klv.X // heading to rotate by
}
func (self RotWarp) Sig() klv.Sig {
	return self.S
}
type SkalWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	Skal klv.X // vek to scale by
}
func (self SkalWarp) Sig() klv.Sig {
	return self.S
}
type CmboWarp struct {
	klv.Guk // provides nal func
	S klv.Sig
	A klv.X
	B klv.X
}
func (self Vek) Sig() klv.Sig {
	return self.S
}

// *** pin
type Pin struct {
	klv.Guk // provides nal func
	S klv.Sig
	At klv.X // pin position
}
func (self Pin) Sig() klv.Sig {
	return self.S
}

// *** ra
type Ra struct {
	klv.Guk // provides nal func
	S klv.Sig
	A klv.X // first pin
	B klv.X // second pin
}
func (self Ra) Sig() klv.Sig {
	return self.S
}

// *** Yok
type Yok struct {
	klv.Guk // provides nal func
	S klv.Sig
	A klv.X // first ra
	B klv.X // second ra
	Van klv.X // vane vek
}
func (self Yok) Sig() klv.Sig {
	return self.S
}

// *** flit
type Flit struct {
	klv.Guk // provides nal func
	S klv.Sig
	Yoks []X
}
func (self Flit) Sig() klv.Sig {
	return self.S
}

