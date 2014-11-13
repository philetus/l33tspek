// 2d vector and homogeneous transformation matrix lib

package flat

import (
	"fmt"
	"math"
	//"reflect"
)

var (
	// identity matrix warp
	IandI = Warp{
		Wek{1.0, 0.0, 0.0},
		Wek{0.0, 1.0, 0.0},
		Wek{0.0, 0.0, 1.0},
	}
)

// 2d vector
type Vek [2]float64

// 3 scalar vector for homogenous coordinates
type Wek [3]float64

// 3x3 matrix for homogeneous coordinate transformations
type Warp [3]Wek

// string methods
func (vk *Vek) String() string {
	return fmt.Sprint(vk[0], vk[1])
}
func (wk *Wek) String() string {
	return fmt.Sprint(wk[0], wk[1], wk[2])
}
func (wrp *Warp) String() string {
	return fmt.Sprint(
		"%s %s %s", 
		wrp[0].String(), wrp[1].String(), wrp[2].String())
}

// generates a translation matrix from a delta vector
func LatWarp(dlta Vek) Warp {
	return Warp{
		Wek{1.0, 0.0, 0.0},
		Wek{0.0, 1.0, 0.0},
		Wek{dlta[0], dlta[1], 1.0},
	}
}

// generates a rotation matrix from a heading vector
// if heading vector is zero magnitude returns identity matrix
func RotWarp(hdng Vek) Warp {
	if MagSkwr(hdng) == 0.0 {
		return IandI
	}
	a := Hed(hdng)
	sn, cs := math.Sin(a), math.Cos(a)
	return Warp{
		Wek{cs, sn, 0.0},
		Wek{-sn, cs, 0.0},
		Wek{0.0, 0.0, 1.0},
	}
}

// generates a translation matrix from a symmetry line
// if symmetry line is zero magnitude returns identity matrix
func FlktWarp(sym Vek) Warp {
	u, ok := Normlz(sym)
	if !ok {
		return IandI
	}
	return Warp{
		Wek{u[0]*u[0] - u[1]*u[1], 2.0*u[0]*u[1], 0.0},
		Wek{2.0*u[0]*u[1], u[1]*u[1] - u[0]*u[0], 0.0},
		Wek{0.0, 0.0, 1.0},
	}
}

// generates a scaling matrix from a scale vector
func SkalWarp(skl Vek) Warp {
	return Warp{
		Wek{skl[0], 0.0, 0.0},
		Wek{0.0, skl[1], 0.0},
		Wek{0.0, 0.0, 1.0},
	}
}

// returns angle of vek from y axis
func Hed(v Vek) float64 {
	return math.Atan2(v[1], v[0])
}

// return unit length vek with same heading, and boolean for zero fail 
func Normlz(v Vek) (u Vek, ok bool) {
	s := MagSkwr(v)
	switch {
	case s == 0.0:
		u, ok = v, false
	case s == 1.0:
		u, ok = v, true
	default:
		u = Skal(v, 1.0 / math.Sqrt(s))
		ok = true
	}
	return
}

func Skal(v Vek, s float64) Vek {
	return Vek{v[0]*s, v[1]*s}
}

// returns magnitude (length) of vek
func Mag(v Vek) float64 {
	return float64(math.Sqrt(MagSkwr(v)))
}

// returns square of the magnitude of vek
func MagSkwr(v Vek) float64 {
	return v[0]*v[0] + v[1]*v[1]
}

// add and subtract veks
func Add(a, b Vek) Vek {
	return Vek{a[0] + b[0], a[1] + b[1]}
}
func Sub(a, b Vek) Vek {
	return Vek{a[0] - b[0], a[1] - b[1]}
}

// return dot product of two veks
func Dot(a, b Vek) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// return cross product of two veks
func Kros(a, b Vek) Vek {
	return Vek{
		a[1]*b[0] - a[0]*b[1],
		a[0]*b[1] - a[1]*b[0],
	}
}

// multiply two veks and return result as new vek
func Mult(a, b Vek) Vek {
	return Vek{a[0]*b[0], a[1]*b[1]}
}

// multiply vek by warp and return result as new vek
func Mog(w Warp, v Vek) Vek {
	return Vek{
		v[0]*w[0][0] + v[1]*w[1][0] + w[2][0],
		v[0]*w[0][1] + v[1]*w[1][1] + w[2][1],
	}
}

// multiply two warps and return result as new warp
func CmboWarp(a, b Warp) Warp {
	return Warp{
		Wek{
			a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0],
			a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1],
			a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2],
		},
		Wek{
			a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0],
			a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1],
			a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2],
		},
		Wek{
			a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0],
			a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1],
			a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2],
		},
	}
}

// transposes warp and returns it as new warp
func Tpos(w Warp) Warp {
	return Warp{
		Wek{w[0][0], w[1][0], w[2][0]},
		Wek{w[0][1], w[1][1], w[2][1]},
		Wek{w[0][2], w[1][2], w[2][2]},
	}
}
