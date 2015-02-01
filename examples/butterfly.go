package main

import (
	"fmt"
	"github.com/philetus/eezl"
	"github.com/philetus/l33tspek/tag"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/fig"
	"github.com/philetus/l33tspek/zib"
)

func butterfly(sss *zib.Sess) {

	// to hold butterfly
	bfly := &zib.Zib{
		Hndl: "bfly",
		Swg: fig.Swag{
			Dpth: 2.0,
			Wt: 5.0,
			Klr: [4]float64{1.0, 0.0, 0.0, 0.6}, // translucent red
		},
	}
	sss.Sup(bfly) // intro to sess
	sss.Zut(bfly) // set as rut zib
	
	// default down delta vek [1.0, 0.0]
	bfly.Skrib(zib.Delta{
		Hndl: "d_down", 
		Crds: flat.Down,
	})
	
	// butterfly spine yoks
	bfly.Skrib(zib.Delta{
		Hndl: "d0",
		Crds: flat.Vek{0.0, 0.0},
	})
	bfly.Skrib(zib.Yok{
		Hndl: "y0",
		Spt: zib.XDelta{"bfly", "d0"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})
	bfly.Skrib(zib.Delta{
		Hndl: "d1",
		Crds: flat.Vek{-33.0, 0.0},
	})
	bfly.Skrib(zib.Yok{
		Hndl: "y1",
		Spt: zib.XDelta{"bfly", "d1"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})
	
	// left wing
	lwng := &zib.Zib{
		Hndl: "lwng",
		Wrp: zib.XWarp{"lwng", "hip"}, // defined below
	}
	sss.Sup(lwng) // intro to sess
	bfly.Nyun(lwng) // sub to bfly
	
	// hip translation warp
	lwng.Skrib(zib.Delta{
		Hndl: "d_hip", 
		Crds: flat.Vek{8.0, 15.0},
	})
	lwng.Skrib(zib.LatWarp{ // transform to first yok of left wing profile
		Hndl: "hip",
		Dlta: zib.XDelta{"lwng", "d_hip"},
	})
	
	// left wing profile yoks
	lwng.Skrib(zib.Delta{
		Hndl: "d0",
		Crds: flat.Vek{0.0, 0.0},
	})
	lwng.Skrib(zib.Yok{
		Hndl: "y0",
		Spt: zib.XDelta{"lwng", "d0"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})
	lwng.Skrib(zib.Delta{
		Hndl: "d1",
		Crds: flat.Vek{7.0, 15.0},
	})
	lwng.Skrib(zib.Yok{
		Hndl: "y1",
		Spt: zib.XDelta{"lwng", "d1"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})
	lwng.Skrib(zib.Delta{
		Hndl: "d2",
		Crds: flat.Vek{-26.0, 0.0},
	})
	lwng.Skrib(zib.Yok{
		Hndl: "y2",
		Spt: zib.XDelta{"lwng", "d2"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})
	lwng.Skrib(zib.Delta{
		Hndl: "d3",
		Crds: flat.Vek{-47.0, 5.0},
	})
	lwng.Skrib(zib.Yok{
		Hndl: "y3",
		Spt: zib.XDelta{"lwng", "d3"},
		Gd: zib.XDelta{"bfly", "d_down"},
	})

	// right wing
	rwng := &zib.Zib{
		Hndl: "rwng",
		Ark: zib.XZib{"lwng"}, // inherits marks from lwng
		Wrp: zib.XWarp{"rwng", "flip_hip"}, // defined below
	}
	sss.Sup(rwng) // intro to sess
	bfly.Nyun(rwng) // sub to bfly
	
	// build combowarp that reflects across spine then translates to hip
	lwng.Skrib(zib.FlektWarp{ 
		Hndl: "flip",
		Sym: zib.XDelta{"bfly", "d_down"},
	})
	lwng.Skrib(zib.ComboWarp{ 
		Hndl: "flip_hip",
		Nyns: []zib.XWarp{
			{"rwng", "flip"},
			{"lwng", "hip"},
		},
	})
    
    bfly.Skrib(zib.Flit{
    	Hndl: "wngs",
    	Yks: []zib.XYok{
    		{"bfly", "y0"},
    		{"lwng", "y0"},
    		{"lwng", "y1"},
    		{"lwng", "y2"},
    		{"lwng", "y3"},
    		{"bfly", "y1"},
    		{"rwng", "y3"},
    		{"rwng", "y2"},
    		{"rwng", "y1"},
    		{"rwng", "y0"},
    		{"bfly", "y0"},
    	},
    })
    
    return
}

func main() {
	xscreen := eezl.Xconnect() // get connection to x server screen
	ez := xscreen.NewEezl(600, 600) // open window
	fmt.Printf("created new eezl!\n")
	
	var pressed_flag bool = false
	
    //var ln_thk float64 = 10.0
    //ln_clr := [4]float64{1.0, 0.0, 0.0, 0.6} // slightly translucent red
    //var bnd_thk float64 = 6.0
    //bnd_clr := [4]float64{0.0, 0.0, 0.0, 0.4} // translucent gray
    bg_clr := [4]float64{1.0, 1.0, 1.0, 1.0} // opaque white
    
    // start zib sess and add butterfly to it
    sss := &zib.Sess{Dfl: tag.Dufl{"examples", "butterfly"}}
    butterfly(sss)
    
    // warp to put butterfly in window
    wrp := flat.ComboWarp(
    	flat.LatWarp(flat.Vek{60.0, 40.0}),
    	flat.SkalWarp(flat.Vek{6.0, 6.0}),
    )
    
	for {
		select {
			
			case inpt := <-ez.InputPipe:
				//fmt.Printf(".%d", inpt.Flavr)
				switch inpt.Flavr {
				
				case eezl.PointerPress:
					pressed_flag = true
					fmt.Printf("pressed pointer!\n")


				case eezl.PointerRelease:
					pressed_flag = false
					fmt.Printf("pressed pointer!\n")
					
					//ez.Stain() // trigger eezl redraw
					
				case eezl.PointerMotion:
					if pressed_flag {
						fmt.Printf("dragged pointer!\n")
						
						//ez.Stain() // trigger eezl redraw
					}
				
				case eezl.KeyPress:
					fmt.Println(inpt.Stroke.Name)
				}
				
			case gel := <- ez.GelPipe:
			
				// fill background
				gel.SetColor(bg_clr[0], bg_clr[1], bg_clr[2], bg_clr[3])
				gel.Coat()
								
				// collapse butterfly sess into fig spek
				bfig := sss.Klaps()
				
				// render to gel
				bfig.Gelr(gel, wrp)
								
				// draw band
				//if pressed_flag {
				//}
				
				// send trigger sig
				gel.Trigger()
		}
	}
}
