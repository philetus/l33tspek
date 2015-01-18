package main

import (
	"fmt"
	"github.com/philetus/eezl"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/flat/fig"
)

func main() {
	xscreen := eezl.Xconnect() // get connection to x server screen
	ez := xscreen.NewEezl(600, 600) // open window
	fmt.Printf("created new eezl!\n")
	
	fg := &fig.Fig{
		Hndl: fig.Handl{"bands"},
	}
	
	// set up swags
	bndSwg := fig.Swag{
		Vz: true,
		Dpth: 1.0,
		Wt: 6.0,
		Klr: [4]float64{0.0, 0.0, 0.0, 0.4}, // translucent gray
	}
	
	linSwg := fig.Swag{
		Vz: true,
		Dpth: 2.0,
		Wt: 10.0,
		Klr: [4]float64{1.0, 0.0, 0.0, 0.6}, // translucent red
	}
	
	// set up band
	bildFlit(
		fg, 
		fig.Handl{"bnd"},
		fig.Swag{}, fig.Swag{}, // default swag not visible
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
						fig.Handl{"bnd"},
						fig.Swag{}, bndSwg,
						frst, lst,
					)
					
					pressed_flag = true

				case eezl.PointerRelease:

					// build new lin in current band position
					s := fmt.Sprintf("_lin%v", lnKnt)
					lnKnt++
					bildFlit(
						fg, 
						fig.Handl{fig.Sig(s)},
						fig.Swag{}, linSwg,
						frst, lst,
					)
					
					// set band visibility to false
					bildFlit(
						fg, 
						fig.Handl{"bnd"},
						fig.Swag{}, fig.Swag{}, // default is not visible
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
							fig.Handl{"bnd"},
							fig.Swag{}, bndSwg,
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
		hndl fig.Handl,
		ykSwg, fltSwg fig.Swag,
		frst, lst flat.Vek,
	) fig.Flit {

	// build marks
	yks := []fig.Yok{
		fig.Yok{
			Hndl: append(hndl, "_frst"), // append yok sig to handl
			Swg: ykSwg,
			Spt: frst,
		},
		fig.Yok{
			Hndl: append(hndl, "_lst"),
			Swg: ykSwg,
			Spt: lst,
		},
	}
	flt := fig.Flit{
			Hndl: hndl,
			Swg: fltSwg,
			Yoks: []fig.Handl{yks[0].Hndl, yks[1].Hndl},
	}
	
	// deks marks with fig
	fg.Add(yks[0])
	fg.Add(yks[1])
	fg.Add(flt)
	
	return flt
}

