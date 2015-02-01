package main

import (
	"fmt"
	"github.com/philetus/eezl"
	"github.com/philetus/l33tspek/tag"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/fig"
)

func main() {
	xscreen := eezl.Xconnect() // get connection to x server screen
	ez := xscreen.NewEezl(600, 600) // open window
	fmt.Printf("created new eezl!\n")
	
	fg := &fig.Fig{
		Dfl: tag.Dufl{"bands"},
	}
	
	// set up swags
	bndSwg := fig.Swag{
		Dpth: 1.0,
		Wt: 6.0,
		Klr: [4]float64{0.0, 0.0, 0.0, 0.4}, // translucent gray
	}
	
	linSwg := fig.Swag{
		Dpth: 2.0,
		Wt: 10.0,
		Klr: [4]float64{1.0, 0.0, 0.0, 0.6}, // translucent red
	}
	
	// set up band
	bildFlit(
		fg, 
		tag.Dufl{"bnd"},
		fig.Swag{Hd: true}, fig.Swag{Hd: true}, // set swag hide flag
		flat.Vek{}, flat.Vek{}, // init with veks at 0, 0
	)
	var frst, lst flat.Vek
	
	// global transform
	wrp := flat.IandI
		
	var pressed_flag bool = false
	var lnKnt int = 1
	
    bg_clr := []float64{1.0, 1.0, 1.0, 1.0} // opaque white

	fmt.Println("entering mainloop")
	
	for {
		select {
			
			case inpt := <-ez.InputPipe:				
				switch inpt.Flavr {
				
				case eezl.PointerPress:
					frst = flat.Vek{float64(inpt.Y), float64(inpt.X)}
					lst = frst
					bildFlit(
						fg, 
						tag.Dufl{"bnd"},
						fig.Swag{Hd: true}, bndSwg,
						frst, lst,
					)
					
					pressed_flag = true

				case eezl.PointerRelease:

					// build new lin in current band position
					s := fmt.Sprintf("_lin%v", lnKnt)
					lnKnt++
					bildFlit(
						fg, 
						tag.Dufl{tag.Handl(s)},
						fig.Swag{Hd: true}, linSwg,
						frst, lst,
					)
					
					// set band visibility to false
					bildFlit(
						fg, 
						tag.Dufl{"bnd"},
						fig.Swag{Hd: true}, fig.Swag{Hd: true}, // default is not visible
						frst, lst,
					)
					
					pressed_flag = false
								   
					ez.Stain() // trigger eezl redraw
					
				case eezl.PointerMotion:
					if pressed_flag {
						lst = flat.Vek{float64(inpt.Y), float64(inpt.X)}
						
						// rebuild band flit in current position
						bildFlit(
							fg, 
							tag.Dufl{"bnd"},
							fig.Swag{Hd: true}, bndSwg,
							frst, lst,
						)
						
						ez.Stain() // trigger eezl redraw
					}
				
				case eezl.KeyPress:
					fmt.Println(inpt.Stroke.Name)
				}
				
			case gel := <- ez.GelPipe:
							
				// fill background
				gel.SetColor(bg_clr[0], bg_clr[1], bg_clr[2], bg_clr[3])
				gel.Coat()
				
				// render fig with gelr
				fg.Gelr(gel, wrp)
								
				// send trigger sig
				gel.Trigger()
		}
	}
}

func bildFlit(
		fg *fig.Fig,
		dfl tag.Dufl,
		ykSwg, fltSwg fig.Swag,
		frst, lst flat.Vek,
	) fig.Flit {

	// build marks
	yks := []fig.Yok{
		fig.Yok{
			Dfl: append(dfl, "_frst"), // append yok sig to handl
			Swg: ykSwg,
			Spt: frst,
		},
		fig.Yok{
			Dfl: append(dfl, "_lst"),
			Swg: ykSwg,
			Spt: lst,
		},
	}
	flt := fig.Flit{
			Dfl: dfl,
			Swg: fltSwg,
			Yks: []tag.Dufl{yks[0].Dfl, yks[1].Dfl},
	}
	
	// deks marks with fig
	fg.Skrib(yks[0])
	fg.Skrib(yks[1])
	fg.Skrib(flt)
	
	return flt
}

