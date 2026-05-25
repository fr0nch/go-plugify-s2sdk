package s2sdk

/*
#include "filesystem.h"
#cgo noescape ReadFileVPK
#cgo noescape FindFileAbsoluteList
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin()

// Generated from s2sdk (group: filesystem)

// ReadFileVPK 
//  @brief Reads a file and returns its contents as a string.
//
//  @param localFileName: The relative path of the file to read.
//  @param pathId: The filesystem search path ID (e.g., "GAME"). If empty, uses "GAME".
//
//  @return The file contents, or an empty string on failure.
func ReadFileVPK(localFileName string, pathId string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__localFileName := plugify.ConstructString(localFileName)
	__pathId := plugify.ConstructString(pathId)
	plugify.Block {
		Try: func() {
			__native := C.ReadFileVPK((*C.String)(unsafe.Pointer(&__localFileName)), (*C.String)(unsafe.Pointer(&__pathId)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__localFileName)
			plugify.DestroyString(&__pathId)
		},
	}.Do()
	return __retVal
}

// FindFileAbsoluteList 
//  @brief Finds all files matching the given wildcard and path ID.
//
//  @param wildcard: The wildcard pattern to match.
//  @param pathId: The filesystem search path ID (e.g., "GAME"). If empty, uses "GAME".
//
//  @return The list of absolute file paths matching the wildcard.
func FindFileAbsoluteList(wildcard string, pathId string) []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	__wildcard := plugify.ConstructString(wildcard)
	__pathId := plugify.ConstructString(pathId)
	plugify.Block {
		Try: func() {
			__native := C.FindFileAbsoluteList((*C.String)(unsafe.Pointer(&__wildcard)), (*C.String)(unsafe.Pointer(&__pathId)))
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
			plugify.DestroyString(&__wildcard)
			plugify.DestroyString(&__pathId)
		},
	}.Do()
	return __retVal
}

