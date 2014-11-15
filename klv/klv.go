// klv is interface to address hok nuks for deferred lookup
// cleave reality at the joints --> ham hocks


package klv

import (
	//"fmt"
	//"reflect"
	//"strconv"
)

type Sig string

type Nuk interface {
	Sig() Sig
}

// hoks hold nuks referenced by sig
type Hok interface {
	Sig() Sig  // unique for hoks in rent
	Ark() (Hok, bool) // archetype hok
	HasKids() bool
	Kid(Sig) (Hok, bool)
	Jnt(Hok) Hok // joint new kid hok into this one by sig
	HasNuks() bool
	Nuk(Sig) (Nuk, bool)
	Swlo(Nuk) X // swallow new nuk by sig, returns address for rent nuk
}

// retrieves sigs for rent and nuk from hok tree at given root by x address
func Dox(hk Hok, x X) (rentSg Sig, nukSg Sig, ok bool) {
	sg := 

	// if x is leaf check that sg is in nuks
	if x.IsLef() {
		if nk, okk := hk.Nuk(x.Sig()); okk {
			return hk.Sig(), x.Sig(), true
		}
	
	// otherwise if x is not lef try to recurse on kid hok
	} else {
		if kid, has := hk.Kid(x.Sig()); has {
			return Dox(kid, x.Kid())
		}
	}
	
	// otherwise try to recurse on ark
	if ark, okk := hk.Ark(); okk {
		return Dox(ark, x)
	}
	
	// otherwise fail
	ok = false
	return
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

