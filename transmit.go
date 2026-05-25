package s2sdk

/*
#include "transmit.h"
#cgo noescape SetTransmitInfoEntity
#cgo noescape ClearTransmitInfoEntity
#cgo noescape IsTransmitInfoEntitySet
#cgo noescape SetTransmitInfoEntityAll
#cgo noescape ClearTransmitInfoEntityAll
#cgo noescape SetTransmitInfoNonPlayer
#cgo noescape ClearTransmitInfoNonPlayer
#cgo noescape IsTransmitInfoNonPlayerSet
#cgo noescape SetTransmitInfoNonPlayerAll
#cgo noescape ClearTransmitInfoNonPlayerAll
#cgo noescape SetTransmitInfoOutOfPVS
#cgo noescape ClearTransmitInfoOutOfPVS
#cgo noescape IsTransmitInfoOutOfPVSSet
#cgo noescape SetTransmitInfoOutOfPVSAll
#cgo noescape ClearTransmitInfoOutOfPVSAll
#cgo noescape SetTransmitInfoHLTV
#cgo noescape ClearTransmitInfoHLTV
#cgo noescape IsTransmitInfoHLTVSet
#cgo noescape SetTransmitInfoHLTVAll
#cgo noescape ClearTransmitInfoHLTVAll
#cgo noescape GetTransmitInfoTargetSlotsCount
#cgo noescape GetTransmitInfoTargetSlot
#cgo noescape AddTransmitInfoTargetSlot
#cgo noescape RemoveTransmitInfoTargetSlot
#cgo noescape GetTransmitInfoTargetSlotsAll
#cgo noescape RemoveTransmitInfoTargetSlotsAll
#cgo noescape GetTransmitInfoPlayerSlot
#cgo noescape SetTransmitInfoPlayerSlot
#cgo noescape GetTransmitInfoFullUpdate
#cgo noescape SetTransmitInfoFullUpdate
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

// Generated from s2sdk (group: transmit)

// SetTransmitInfoEntity 
//  @brief Sets a bit in the TransmitEntity bitvec, marking an entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as transmittable.
func SetTransmitInfoEntity(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoEntity(__info, __entityHandle)
}

// ClearTransmitInfoEntity 
//  @brief Clears a bit in the TransmitEntity bitvec, marking an entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as not transmittable.
func ClearTransmitInfoEntity(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoEntity(__info, __entityHandle)
}

// IsTransmitInfoEntitySet 
//  @brief Checks if a bit is set in the TransmitEntity bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func IsTransmitInfoEntitySet(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoEntitySet(__info, __entityHandle))
	return __retVal
}

// SetTransmitInfoEntityAll 
//  @brief Sets all bits in the TransmitEntity bitvec, marking all entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoEntityAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoEntityAll(__info)
}

// ClearTransmitInfoEntityAll 
//  @brief Clears all bits in the TransmitEntity bitvec, marking all entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoEntityAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoEntityAll(__info)
}

// SetTransmitInfoNonPlayer 
//  @brief Sets a bit in the TransmitNonPlayers bitvec, marking a non-player entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as transmittable.
func SetTransmitInfoNonPlayer(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoNonPlayer(__info, __entityHandle)
}

// ClearTransmitInfoNonPlayer 
//  @brief Clears a bit in the TransmitNonPlayers bitvec, marking a non-player entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as not transmittable.
func ClearTransmitInfoNonPlayer(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoNonPlayer(__info, __entityHandle)
}

// IsTransmitInfoNonPlayerSet 
//  @brief Checks if a bit is set in the TransmitNonPlayers bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func IsTransmitInfoNonPlayerSet(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoNonPlayerSet(__info, __entityHandle))
	return __retVal
}

// SetTransmitInfoNonPlayerAll 
//  @brief Sets all bits in the TransmitNonPlayers bitvec, marking all non-player entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoNonPlayerAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoNonPlayerAll(__info)
}

// ClearTransmitInfoNonPlayerAll 
//  @brief Clears all bits in the TransmitNonPlayers bitvec, marking all non-player entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoNonPlayerAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoNonPlayerAll(__info)
}

// SetTransmitInfoOutOfPVS 
//  @brief Sets a bit in the TransmitOutOfPVS bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func SetTransmitInfoOutOfPVS(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoOutOfPVS(__info, __entityHandle)
}

// ClearTransmitInfoOutOfPVS 
//  @brief Clears a bit in the TransmitOutOfPVS bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func ClearTransmitInfoOutOfPVS(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoOutOfPVS(__info, __entityHandle)
}

// IsTransmitInfoOutOfPVSSet 
//  @brief Checks if a bit is set in the TransmitOutOfPVS bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func IsTransmitInfoOutOfPVSSet(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoOutOfPVSSet(__info, __entityHandle))
	return __retVal
}

// SetTransmitInfoOutOfPVSAll 
//  @brief Sets all bits in the TransmitOutOfPVS bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoOutOfPVSAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoOutOfPVSAll(__info)
}

// ClearTransmitInfoOutOfPVSAll 
//  @brief Clears all bits in the TransmitOutOfPVS bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoOutOfPVSAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoOutOfPVSAll(__info)
}

// SetTransmitInfoHLTV 
//  @brief Sets a bit in the TransmitHLTV bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func SetTransmitInfoHLTV(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoHLTV(__info, __entityHandle)
}

// ClearTransmitInfoHLTV 
//  @brief Clears a bit in the TransmitHLTV bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func ClearTransmitInfoHLTV(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoHLTV(__info, __entityHandle)
}

// IsTransmitInfoHLTVSet 
//  @brief Checks if a bit is set in the TransmitHLTV bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func IsTransmitInfoHLTVSet(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoHLTVSet(__info, __entityHandle))
	return __retVal
}

// SetTransmitInfoHLTVAll 
//  @brief Sets all bits in the TransmitHLTV bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoHLTVAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoHLTVAll(__info)
}

// ClearTransmitInfoHLTVAll 
//  @brief Clears all bits in the TransmitHLTV bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoHLTVAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoHLTVAll(__info)
}

// GetTransmitInfoTargetSlotsCount 
//  @brief Gets the count of target player slots.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The number of target player slots, or 0 if the info pointer is null.
func GetTransmitInfoTargetSlotsCount(info uintptr) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__retVal = int32(C.GetTransmitInfoTargetSlotsCount(__info))
	return __retVal
}

// GetTransmitInfoTargetSlot 
//  @brief Gets a player slot value at a specific index in the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param index: The index in the target slots vector.
//
//  @return The player slot value, or -1 if the index is invalid or info is null.
func GetTransmitInfoTargetSlot(info uintptr, index int32) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__index := C.int32_t(index)
	__retVal = int32(C.GetTransmitInfoTargetSlot(__info, __index))
	return __retVal
}

// AddTransmitInfoTargetSlot 
//  @brief Adds a player slot to the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to add.
func AddTransmitInfoTargetSlot(info uintptr, playerSlot int32) {
	__info := C.uintptr_t(info)
	__playerSlot := C.int32_t(playerSlot)
	C.AddTransmitInfoTargetSlot(__info, __playerSlot)
}

// RemoveTransmitInfoTargetSlot 
//  @brief Removes a player slot from the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param index: Index within the target slots vector to remove.
func RemoveTransmitInfoTargetSlot(info uintptr, index int32) {
	__info := C.uintptr_t(info)
	__index := C.int32_t(index)
	C.RemoveTransmitInfoTargetSlot(__info, __index)
}

// GetTransmitInfoTargetSlotsAll 
//  @brief Gets the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slots array.
func GetTransmitInfoTargetSlotsAll(info uintptr) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__info := C.uintptr_t(info)
	plugify.Block {
		Try: func() {
			__native := C.GetTransmitInfoTargetSlotsAll(__info)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// RemoveTransmitInfoTargetSlotsAll 
//  @brief Clears all target player slots from the vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func RemoveTransmitInfoTargetSlotsAll(info uintptr) {
	__info := C.uintptr_t(info)
	C.RemoveTransmitInfoTargetSlotsAll(__info)
}

// GetTransmitInfoPlayerSlot 
//  @brief Gets the player slot value from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slot value, or -1 if info is null.
func GetTransmitInfoPlayerSlot(info uintptr) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__retVal = int32(C.GetTransmitInfoPlayerSlot(__info))
	return __retVal
}

// SetTransmitInfoPlayerSlot 
//  @brief Sets the player slot value in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to set.
func SetTransmitInfoPlayerSlot(info uintptr, playerSlot int32) {
	__info := C.uintptr_t(info)
	__playerSlot := C.int32_t(playerSlot)
	C.SetTransmitInfoPlayerSlot(__info, __playerSlot)
}

// GetTransmitInfoFullUpdate 
//  @brief Gets the full update flag from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return True if full update is enabled, false otherwise.
func GetTransmitInfoFullUpdate(info uintptr) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__retVal = bool(C.GetTransmitInfoFullUpdate(__info))
	return __retVal
}

// SetTransmitInfoFullUpdate 
//  @brief Sets the full update flag in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param fullUpdate: The full update flag value to set.
func SetTransmitInfoFullUpdate(info uintptr, fullUpdate bool) {
	__info := C.uintptr_t(info)
	__fullUpdate := C.bool(fullUpdate)
	C.SetTransmitInfoFullUpdate(__info, __fullUpdate)
}

var (
	CheckTransmitInfoErrEmptyHandle = errors.New("CheckTransmitInfo: empty handle")
)

//  @brief RAII wrapper for CheckTransmitInfo pointer.
//
type CheckTransmitInfo struct {
	handle    uintptr
}

// NewCheckTransmitInfo creates a CheckTransmitInfo from a handle
func NewCheckTransmitInfo(handle uintptr) *CheckTransmitInfo {
	return &CheckTransmitInfo{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *CheckTransmitInfo) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *CheckTransmitInfo) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *CheckTransmitInfo) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *CheckTransmitInfo) IsValid() bool {
	return w.handle != 0
}

// SetEntity 
//  @brief Sets a bit in the TransmitEntity bitvec, marking an entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as transmittable.
func (w *CheckTransmitInfo) SetEntity(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoEntity(w.handle, entityHandle)
	return nil
}

// ClearEntity 
//  @brief Clears a bit in the TransmitEntity bitvec, marking an entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as not transmittable.
func (w *CheckTransmitInfo) ClearEntity(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoEntity(w.handle, entityHandle)
	return nil
}

// IsEntitySet 
//  @brief Checks if a bit is set in the TransmitEntity bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func (w *CheckTransmitInfo) IsEntitySet(entityHandle int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return IsTransmitInfoEntitySet(w.handle, entityHandle), nil
}

// SetEntityAll 
//  @brief Sets all bits in the TransmitEntity bitvec, marking all entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) SetEntityAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoEntityAll(w.handle)
	return nil
}

// ClearEntityAll 
//  @brief Clears all bits in the TransmitEntity bitvec, marking all entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) ClearEntityAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoEntityAll(w.handle)
	return nil
}

// SetNonPlayer 
//  @brief Sets a bit in the TransmitNonPlayers bitvec, marking a non-player entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as transmittable.
func (w *CheckTransmitInfo) SetNonPlayer(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoNonPlayer(w.handle, entityHandle)
	return nil
}

// ClearNonPlayer 
//  @brief Clears a bit in the TransmitNonPlayers bitvec, marking a non-player entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as not transmittable.
func (w *CheckTransmitInfo) ClearNonPlayer(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoNonPlayer(w.handle, entityHandle)
	return nil
}

// IsNonPlayerSet 
//  @brief Checks if a bit is set in the TransmitNonPlayers bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func (w *CheckTransmitInfo) IsNonPlayerSet(entityHandle int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return IsTransmitInfoNonPlayerSet(w.handle, entityHandle), nil
}

// SetNonPlayerAll 
//  @brief Sets all bits in the TransmitNonPlayers bitvec, marking all non-player entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) SetNonPlayerAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoNonPlayerAll(w.handle)
	return nil
}

// ClearNonPlayerAll 
//  @brief Clears all bits in the TransmitNonPlayers bitvec, marking all non-player entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) ClearNonPlayerAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoNonPlayerAll(w.handle)
	return nil
}

// SetOutOfPVS 
//  @brief Sets a bit in the TransmitOutOfPVS bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func (w *CheckTransmitInfo) SetOutOfPVS(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoOutOfPVS(w.handle, entityHandle)
	return nil
}

// ClearOutOfPVS 
//  @brief Clears a bit in the TransmitOutOfPVS bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func (w *CheckTransmitInfo) ClearOutOfPVS(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoOutOfPVS(w.handle, entityHandle)
	return nil
}

// IsOutOfPVSSet 
//  @brief Checks if a bit is set in the TransmitOutOfPVS bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func (w *CheckTransmitInfo) IsOutOfPVSSet(entityHandle int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return IsTransmitInfoOutOfPVSSet(w.handle, entityHandle), nil
}

// SetOutOfPVSAll 
//  @brief Sets all bits in the TransmitOutOfPVS bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) SetOutOfPVSAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoOutOfPVSAll(w.handle)
	return nil
}

// ClearOutOfPVSAll 
//  @brief Clears all bits in the TransmitOutOfPVS bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) ClearOutOfPVSAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoOutOfPVSAll(w.handle)
	return nil
}

// SetHLTV 
//  @brief Sets a bit in the TransmitHLTV bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func (w *CheckTransmitInfo) SetHLTV(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoHLTV(w.handle, entityHandle)
	return nil
}

// ClearHLTV 
//  @brief Clears a bit in the TransmitHLTV bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func (w *CheckTransmitInfo) ClearHLTV(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoHLTV(w.handle, entityHandle)
	return nil
}

// IsHLTVSet 
//  @brief Checks if a bit is set in the TransmitHLTV bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func (w *CheckTransmitInfo) IsHLTVSet(entityHandle int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return IsTransmitInfoHLTVSet(w.handle, entityHandle), nil
}

// SetHLTVAll 
//  @brief Sets all bits in the TransmitHLTV bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) SetHLTVAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoHLTVAll(w.handle)
	return nil
}

// ClearHLTVAll 
//  @brief Clears all bits in the TransmitHLTV bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) ClearHLTVAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoHLTVAll(w.handle)
	return nil
}

// GetTargetSlotsCount 
//  @brief Gets the count of target player slots.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The number of target player slots, or 0 if the info pointer is null.
func (w *CheckTransmitInfo) GetTargetSlotsCount() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return GetTransmitInfoTargetSlotsCount(w.handle), nil
}

// GetTargetSlot 
//  @brief Gets a player slot value at a specific index in the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param index: The index in the target slots vector.
//
//  @return The player slot value, or -1 if the index is invalid or info is null.
func (w *CheckTransmitInfo) GetTargetSlot(index int32) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return GetTransmitInfoTargetSlot(w.handle, index), nil
}

// AddTargetSlot 
//  @brief Adds a player slot to the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to add.
func (w *CheckTransmitInfo) AddTargetSlot(playerSlot int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	AddTransmitInfoTargetSlot(w.handle, playerSlot)
	return nil
}

// RemoveTargetSlot 
//  @brief Removes a player slot from the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param index: Index within the target slots vector to remove.
func (w *CheckTransmitInfo) RemoveTargetSlot(index int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	RemoveTransmitInfoTargetSlot(w.handle, index)
	return nil
}

// GetTargetSlotsAll 
//  @brief Gets the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slots array.
func (w *CheckTransmitInfo) GetTargetSlotsAll() ([]int32, error) {
	if w.handle == 0 {
		var zero []int32
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return GetTransmitInfoTargetSlotsAll(w.handle), nil
}

// RemoveTargetSlotsAll 
//  @brief Clears all target player slots from the vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) RemoveTargetSlotsAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	RemoveTransmitInfoTargetSlotsAll(w.handle)
	return nil
}

// GetPlayerSlot 
//  @brief Gets the player slot value from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slot value, or -1 if info is null.
func (w *CheckTransmitInfo) GetPlayerSlot() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return GetTransmitInfoPlayerSlot(w.handle), nil
}

// SetPlayerSlot 
//  @brief Sets the player slot value in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to set.
func (w *CheckTransmitInfo) SetPlayerSlot(playerSlot int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoPlayerSlot(w.handle, playerSlot)
	return nil
}

// GetFullUpdate 
//  @brief Gets the full update flag from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return True if full update is enabled, false otherwise.
func (w *CheckTransmitInfo) GetFullUpdate() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CheckTransmitInfoErrEmptyHandle
	}
	return GetTransmitInfoFullUpdate(w.handle), nil
}

// SetFullUpdate 
//  @brief Sets the full update flag in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param fullUpdate: The full update flag value to set.
func (w *CheckTransmitInfo) SetFullUpdate(fullUpdate bool) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoFullUpdate(w.handle, fullUpdate)
	return nil
}

