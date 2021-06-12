package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func apnggif_syscall(png, gif string, tlevel int, bcolor string) (err error) {
	apng2gifHandler, err := syscall.LoadLibrary("apng2gif.dll")
	if err != nil {
		return err
	}

	apng2gif, err := syscall.GetProcAddress(apng2gifHandler, "apng2gif")
	if err != nil {
		return err
	}

	var narg uintptr = 4
	pngB, err := syscall.BytePtrFromString(png)
	if err != nil {
		return err
	}
	gifB, err := syscall.BytePtrFromString(gif)
	if err != nil {
		return err
	}
	bcolB, err := syscall.BytePtrFromString(bcolor)
	if err != nil {
		return err
	}
	ret, ret2, callErr := syscall.Syscall9(apng2gif,
		narg,
		uintptr(unsafe.Pointer(pngB)),
		uintptr(unsafe.Pointer(gifB)),
		uintptr(0),
		uintptr(unsafe.Pointer(bcolB)),
		0, 0, 0, 0, 0,
	)
	fmt.Printf("%+#v %+#v callErr=%+#v 0\n", ret, ret2, callErr)
	return nil
}

func main() {
	apnggif_syscall("anim.png", "anim.gif", 0, "")
}
