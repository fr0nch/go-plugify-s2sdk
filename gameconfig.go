package s2sdk

/*
#include "gameconfig.h"
#cgo noescape CloseGameConfigFile
#cgo noescape LoadGameConfigFile
#cgo noescape GetGameConfigPatch
#cgo noescape GetGameConfigOffset
#cgo noescape GetGameConfigAddress
#cgo noescape GetGameConfigVTable
#cgo noescape GetGameConfigSignature
#cgo noescape GetGameConfigPatchAll
#cgo noescape GetGameConfigOffsetAll
#cgo noescape GetGameConfigAddressAll
#cgo noescape GetGameConfigVTableAll
#cgo noescape GetGameConfigSignatureAll
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
var _ = plugify.ApiVersion

// Generated from s2sdk (group: gameconfig)

var _CloseGameConfigFile = func(id ConfigId) {
	__id := C.uint32_t(id)
	C.CloseGameConfigFile(__id)
}

// CloseGameConfigFile 
//  @brief Closes a game configuration file.
//
//  @param id: An id to the game configuration to be closed.
func CloseGameConfigFile(id ConfigId) {
	_CloseGameConfigFile(id)
}

var _LoadGameConfigFile = func(paths []string) ConfigId {
	var __retVal ConfigId
	__paths := plugify.ConstructVectorString(paths)
	plugify.Block {
		Try: func() {
			__retVal = ConfigId(C.LoadGameConfigFile((*C.Vector)(unsafe.Pointer(&__paths))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__paths)
		},
	}.Do()
	return __retVal
}

// LoadGameConfigFile 
//  @brief Loads a game configuration file.
//
//  @param paths: The paths to the game configuration file to be loaded.
//
//  @return A id to the loaded game configuration object, or -1 if loading fails.
func LoadGameConfigFile(paths []string) ConfigId {
	return _LoadGameConfigFile(paths)
}

var _GetGameConfigPatch = func(id ConfigId, name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__id := C.uint32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__native := C.GetGameConfigPatch(__id, (*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigPatch 
//  @brief Retrieves a patch associated with the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the patch.
//  @param name: The name of the patch to be retrieved.
//
//  @return A string where the patch will be stored.
func GetGameConfigPatch(id ConfigId, name string) string {
	return _GetGameConfigPatch(id, name)
}

var _GetGameConfigOffset = func(id ConfigId, name string) int32 {
	var __retVal int32
	__id := C.uint32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetGameConfigOffset(__id, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigOffset 
//  @brief Retrieves the offset associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the offset.
//  @param name: The name whose offset is to be retrieved.
//
//  @return The offset associated with the specified name.
func GetGameConfigOffset(id ConfigId, name string) int32 {
	return _GetGameConfigOffset(id, name)
}

var _GetGameConfigAddress = func(id ConfigId, name string) uintptr {
	var __retVal uintptr
	__id := C.uint32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigAddress(__id, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigAddress 
//  @brief Retrieves the address associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the address.
//  @param name: The name whose address is to be retrieved.
//
//  @return A pointer to the address associated with the specified name.
func GetGameConfigAddress(id ConfigId, name string) uintptr {
	return _GetGameConfigAddress(id, name)
}

var _GetGameConfigVTable = func(id ConfigId, name string) uintptr {
	var __retVal uintptr
	__id := C.uint32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigVTable(__id, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigVTable 
//  @brief Retrieves a vtable associated with the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the vtable.
//  @param name: The name of the vtable to be retrieved.
//
//  @return A pointer to the vtable associated with the specified name
func GetGameConfigVTable(id ConfigId, name string) uintptr {
	return _GetGameConfigVTable(id, name)
}

var _GetGameConfigSignature = func(id ConfigId, name string) uintptr {
	var __retVal uintptr
	__id := C.uint32_t(id)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigSignature(__id, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigSignature 
//  @brief Retrieves the signature associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the signature.
//  @param name: The name whose signature is to be resolved and retrieved.
//
//  @return A pointer to the signature associated with the specified name.
func GetGameConfigSignature(id ConfigId, name string) uintptr {
	return _GetGameConfigSignature(id, name)
}

var _GetGameConfigPatchAll = func(name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__native := C.GetGameConfigPatchAll((*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigPatchAll 
//  @brief Retrieves a patch by scanning all loaded game configurations.
//
//  @param name: The name of the patch to be retrieved.
//
//  @return A string containing the patch, or an empty string if not found.
func GetGameConfigPatchAll(name string) string {
	return _GetGameConfigPatchAll(name)
}

var _GetGameConfigOffsetAll = func(name string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetGameConfigOffsetAll((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigOffsetAll 
//  @brief Retrieves an offset by scanning all loaded game configurations.
//
//  @param name: The name whose offset is to be retrieved.
//
//  @return The offset associated with the specified name, or -1 if not found.
func GetGameConfigOffsetAll(name string) int32 {
	return _GetGameConfigOffsetAll(name)
}

var _GetGameConfigAddressAll = func(name string) uintptr {
	var __retVal uintptr
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigAddressAll((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigAddressAll 
//  @brief Retrieves an address by scanning all loaded game configurations.
//
//  @param name: The name whose address is to be retrieved.
//
//  @return A pointer to the address associated with the specified name, or nullptr if not found.
func GetGameConfigAddressAll(name string) uintptr {
	return _GetGameConfigAddressAll(name)
}

var _GetGameConfigVTableAll = func(name string) uintptr {
	var __retVal uintptr
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigVTableAll((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigVTableAll 
//  @brief Retrieves a vtable by scanning all loaded game configurations.
//
//  @param name: The name of the vtable to be retrieved.
//
//  @return A pointer to the vtable associated with the specified name, or nullptr if not found.
func GetGameConfigVTableAll(name string) uintptr {
	return _GetGameConfigVTableAll(name)
}

var _GetGameConfigSignatureAll = func(name string) uintptr {
	var __retVal uintptr
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.GetGameConfigSignatureAll((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigSignatureAll 
//  @brief Retrieves a signature by scanning all loaded game configurations.
//
//  @param name: The name whose signature is to be resolved and retrieved.
//
//  @return A pointer to the signature associated with the specified name, or nullptr if not found.
func GetGameConfigSignatureAll(name string) uintptr {
	return _GetGameConfigSignatureAll(name)
}

var (
	GameConfigErrEmptyHandle = errors.New("GameConfig: empty handle")
)

//  @brief RAII wrapper for GameConfig handle.
//
type GameConfig struct {
	handle    ConfigId
	cleanup   runtime.Cleanup
	ownership ownership
	noCopy    noCopy
}

// NewGameConfigLoadGameConfigFile 
//  @brief Loads a game configuration file.
//
//  @param paths: The paths to the game configuration file to be loaded.
func NewGameConfigLoadGameConfigFile(paths []string) *GameConfig {
	return NewGameConfigOwned(LoadGameConfigFile(paths))
}

// NewGameConfigBorrowed creates a GameConfig from a borrowed handle
func NewGameConfigBorrowed(handle ConfigId) *GameConfig {
	if handle == 0 {
		return &GameConfig{}
	}
	return &GameConfig{
		handle:    handle,
		ownership: Borrowed,
	}
}

// NewGameConfigOwned creates a GameConfig from an owned handle
func NewGameConfigOwned(handle ConfigId) *GameConfig {
	if handle == 0 {
		return &GameConfig{}
	}
	w := &GameConfig{
		handle:    handle,
		ownership: Owned,
	}
	w.cleanup = runtime.AddCleanup(w, destroyGameConfigHandle, handle)
	return w
}

// destroyGameConfigHandle destroys an owned handle.
func destroyGameConfigHandle(handle ConfigId) {
	if handle != 0 {
		CloseGameConfigFile(handle)
	}
}

// destroy cleans up owned handles
func (w *GameConfig) destroy() {
	if w.handle != 0 && w.ownership == Owned {
		CloseGameConfigFile(w.handle)
	}
}

// nullify resets the handle
func (w *GameConfig) nullify() {
	w.handle = 0
	w.ownership = Borrowed
}

// Close explicitly destroys the handle (like C++ destructor, but manual)
func (w *GameConfig) Close() {
	w.Reset()
}

// Get returns the underlying handle
func (w *GameConfig) Get() ConfigId {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *GameConfig) Release() ConfigId {
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	handle := w.handle
	w.nullify()
	return handle
}

// Reset destroys and resets the handle
func (w *GameConfig) Reset() {
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	w.destroy()
	w.nullify()
}

// Valid returns true if handle is not nil
func (w *GameConfig) Valid() bool {
	return w.handle != 0
}

// GetPatch 
//  @brief Retrieves a patch associated with the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the patch.
//  @param name: The name of the patch to be retrieved.
//
//  @return A string where the patch will be stored.
func (w *GameConfig) GetPatch(name string) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, GameConfigErrEmptyHandle
	}
	return GetGameConfigPatch(w.handle, name), nil
}

// GetOffset 
//  @brief Retrieves the offset associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the offset.
//  @param name: The name whose offset is to be retrieved.
//
//  @return The offset associated with the specified name.
func (w *GameConfig) GetOffset(name string) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, GameConfigErrEmptyHandle
	}
	return GetGameConfigOffset(w.handle, name), nil
}

// GetAddress 
//  @brief Retrieves the address associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the address.
//  @param name: The name whose address is to be retrieved.
//
//  @return A pointer to the address associated with the specified name.
func (w *GameConfig) GetAddress(name string) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, GameConfigErrEmptyHandle
	}
	return GetGameConfigAddress(w.handle, name), nil
}

// GetVTable 
//  @brief Retrieves a vtable associated with the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the vtable.
//  @param name: The name of the vtable to be retrieved.
//
//  @return A pointer to the vtable associated with the specified name
func (w *GameConfig) GetVTable(name string) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, GameConfigErrEmptyHandle
	}
	return GetGameConfigVTable(w.handle, name), nil
}

// GetSignature 
//  @brief Retrieves the signature associated with a name from the game configuration.
//
//  @param id: An id to the game configuration from which to retrieve the signature.
//  @param name: The name whose signature is to be resolved and retrieved.
//
//  @return A pointer to the signature associated with the specified name.
func (w *GameConfig) GetSignature(name string) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, GameConfigErrEmptyHandle
	}
	return GetGameConfigSignature(w.handle, name), nil
}

