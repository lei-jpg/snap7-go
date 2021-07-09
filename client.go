package snap7go

//#cgo CFLAGS: -I .
//#include "snap7.h"
import "C"
import (
	"errors"
	"unsafe"
)

type S7Object = C.S7Object

func Cli_GetCpuInfo(cli S7Object) (info TS7CpuInfo, err error) {
	var code C.int = C.Cli_GetCpuInfo(cli, (*C.TS7CpuInfo)(unsafe.Pointer(&info)))
	err = cliErrorsTable[int(code)]
	return
}

var cliErrorsTable = map[int]error{
	0x001: errors.New("Error during PDU negotiation."),
}
