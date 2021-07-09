package snap7go

//#cgo CFLAGS: -I .
//#include "snap7.h"
import "C"
import (
	"errors"
	"unsafe"
)

type S7Object = C.ulong

func Cli_GetCpuInfo(cli S7Object) (info TS7CpuInfo, err error) {
	var code C.int = C.Cli_GetCpuInfo(cli, (*C.TS7CpuInfo)(unsafe.Pointer(&info)))
	err = cliErrorTables[int(code)]
	return
}

var cliErrorTables = map[int]error{
	0x001: errors.New("Error during PDU negotiation."),
}
