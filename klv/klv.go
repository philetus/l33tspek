// klv is interface to address hok nuks for deferred lookup
// cleave reality at the joints --> ham hocks


package klv

import (
	//"fmt"
	//"reflect"
	//"strconv"
)

type Sig string
type Xsig []Sig
type Nal float64

type Nuk interface {
	Sig() Sig
	Nal() (Nal, bool)
}

// mixin for non-nal nuks
type Guk struct {}
func (self Guk) Nal() (Nal, bool) {
	return 0.0, false
}

// hoks hold nuks referenced by sig
type Hok interface {
	Sig() Sig  // unique for hoks in rent
	Ark() Hok // archetype hok
	HasKids() bool
	Kid(Sig) (Hok, bool)
	Jnt(Hok) Hok // joint new kid hok into this one by sig
	HasNuks() bool
	Nuk(Sig) (Nuk, bool)
	Swlo(Nuk) X // swallow new nuk by sig, returns address for rent nuk
}

// retrieves nuk from hok tree at given root by x address
func Dox(hk Hok, x X) (nk Nuk, ss Xsig, ok bool) {
	sg := x.Sig()

	// if x is leaf try to return nuk by x.sig
	if x.IsLef() {
		if nk, ok = hk.Nuk(sg); ok {
			ss = Xsig{sg}
			return
		}
	
	// otherwise if x is not lef try to recurse on kid hok
	} else {
		if kid, has := hk.Kid(sg); has {
			if nk, ss, ok = Dox(kid, x.Kid()); ok {
				ss = append(ss, sg)
			}
			return
		}
	}
	
	// otherwise try to dox with ark
	if ark := self.Ark(); ark != nil {
		return Dox(ark, x)
	}
	
	// otherwise fail
	return nil, nil, false
}

// x is lookup address of nuk
type X interface { 
	Kid() X
	Sig() Sig
	IsLef() bool
}

//
//  *** Xd has sig indexed in current fig
//
type Xd struct {
	S Sig
}
func (self Xd) Kid() X {
	panic("klv.Xd.Kid() called!")
}
func (self Xd) Sig() Sig {
	return self.S
}
func (self Xd) IsLef() bool {
	return true
}

//
// *** XX has sig of kid node to apply 
//
type XX struct {
	S Sig
	K X
}
func (self XX) Kid() X {
	return self.K
}
func (self XX) Sig() Sig {
	return self.S
}
func (self XX) IsLef() bool {
	return false
}

