package main

/*
#cgo CFLAGS:
#cgo LDFLAGS: -L. -lapng2gif
#include "apng2gif.h"
*/
import "C"

import "fmt"

func apnggif_cgo(png, gif string, tlevel int, bcolor string) (err error) {
	cpng := C.CString(png)
	cgif := C.CString(gif)
	cbcol := C.CString(bcolor)
	ret := C.apng2gif(cpng, cgif, C.int(tlevel), cbcol)
	fmt.Printf("%+#v\n", ret)
	return
}

func main() {
	apnggif_cgo("anim.png", "anim.gif", 0, "")
}
