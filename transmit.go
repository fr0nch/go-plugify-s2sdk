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
#cgo noescape SetTransmitInfoAlways
#cgo noescape ClearTransmitInfoAlways
#cgo noescape IsTransmitInfoAlwaysSet
#cgo noescape SetTransmitInfoAlwaysAll
#cgo noescape ClearTransmitInfoAlwaysAll
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
#cgo noescape HideTransmitEntities
#cgo noescape ShowTransmitEntities
#cgo noescape GetHiddenTransmitEntities
#cgo noescape HideTransmitEntity
#cgo noescape ShowTransmitEntity
#cgo noescape HideTransmitEntityFromOtherPlayers
#cgo noescape ShowTransmitEntityToOtherPlayers
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

// Generated from s2sdk (group: transmit)

var _SetTransmitInfoEntity = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoEntity(__info, __entityHandle)
}

// SetTransmitInfoEntity 
//  @brief Sets a bit in the TransmitEntity bitvec, marking an entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as transmittable.
func SetTransmitInfoEntity(info uintptr, entityHandle int32) {
	_SetTransmitInfoEntity(info, entityHandle)
}

var _ClearTransmitInfoEntity = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoEntity(__info, __entityHandle)
}

// ClearTransmitInfoEntity 
//  @brief Clears a bit in the TransmitEntity bitvec, marking an entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as not transmittable.
func ClearTransmitInfoEntity(info uintptr, entityHandle int32) {
	_ClearTransmitInfoEntity(info, entityHandle)
}

var _IsTransmitInfoEntitySet = func(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoEntitySet(__info, __entityHandle))
	return __retVal
}

// IsTransmitInfoEntitySet 
//  @brief Checks if a bit is set in the TransmitEntity bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func IsTransmitInfoEntitySet(info uintptr, entityHandle int32) bool {
	return _IsTransmitInfoEntitySet(info, entityHandle)
}

var _SetTransmitInfoEntityAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoEntityAll(__info)
}

// SetTransmitInfoEntityAll 
//  @brief Sets all bits in the TransmitEntity bitvec, marking all entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoEntityAll(info uintptr) {
	_SetTransmitInfoEntityAll(info)
}

var _ClearTransmitInfoEntityAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoEntityAll(__info)
}

// ClearTransmitInfoEntityAll 
//  @brief Clears all bits in the TransmitEntity bitvec, marking all entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoEntityAll(info uintptr) {
	_ClearTransmitInfoEntityAll(info)
}

var _SetTransmitInfoNonPlayer = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoNonPlayer(__info, __entityHandle)
}

// SetTransmitInfoNonPlayer 
//  @brief Sets a bit in the TransmitNonPlayers bitvec, marking a non-player entity as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as transmittable.
func SetTransmitInfoNonPlayer(info uintptr, entityHandle int32) {
	_SetTransmitInfoNonPlayer(info, entityHandle)
}

var _ClearTransmitInfoNonPlayer = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoNonPlayer(__info, __entityHandle)
}

// ClearTransmitInfoNonPlayer 
//  @brief Clears a bit in the TransmitNonPlayers bitvec, marking a non-player entity as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to mark as not transmittable.
func ClearTransmitInfoNonPlayer(info uintptr, entityHandle int32) {
	_ClearTransmitInfoNonPlayer(info, entityHandle)
}

var _IsTransmitInfoNonPlayerSet = func(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoNonPlayerSet(__info, __entityHandle))
	return __retVal
}

// IsTransmitInfoNonPlayerSet 
//  @brief Checks if a bit is set in the TransmitNonPlayers bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The index of the non-player entity to check.
//
//  @return True if the entity is marked as transmittable, false otherwise.
func IsTransmitInfoNonPlayerSet(info uintptr, entityHandle int32) bool {
	return _IsTransmitInfoNonPlayerSet(info, entityHandle)
}

var _SetTransmitInfoNonPlayerAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoNonPlayerAll(__info)
}

// SetTransmitInfoNonPlayerAll 
//  @brief Sets all bits in the TransmitNonPlayers bitvec, marking all non-player entities as transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoNonPlayerAll(info uintptr) {
	_SetTransmitInfoNonPlayerAll(info)
}

var _ClearTransmitInfoNonPlayerAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoNonPlayerAll(__info)
}

// ClearTransmitInfoNonPlayerAll 
//  @brief Clears all bits in the TransmitNonPlayers bitvec, marking all non-player entities as not transmittable.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoNonPlayerAll(info uintptr) {
	_ClearTransmitInfoNonPlayerAll(info)
}

var _SetTransmitInfoOutOfPVS = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoOutOfPVS(__info, __entityHandle)
}

// SetTransmitInfoOutOfPVS 
//  @brief Sets a bit in the TransmitOutOfPVS bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func SetTransmitInfoOutOfPVS(info uintptr, entityHandle int32) {
	_SetTransmitInfoOutOfPVS(info, entityHandle)
}

var _ClearTransmitInfoOutOfPVS = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoOutOfPVS(__info, __entityHandle)
}

// ClearTransmitInfoOutOfPVS 
//  @brief Clears a bit in the TransmitOutOfPVS bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func ClearTransmitInfoOutOfPVS(info uintptr, entityHandle int32) {
	_ClearTransmitInfoOutOfPVS(info, entityHandle)
}

var _IsTransmitInfoOutOfPVSSet = func(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoOutOfPVSSet(__info, __entityHandle))
	return __retVal
}

// IsTransmitInfoOutOfPVSSet 
//  @brief Checks if a bit is set in the TransmitOutOfPVS bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func IsTransmitInfoOutOfPVSSet(info uintptr, entityHandle int32) bool {
	return _IsTransmitInfoOutOfPVSSet(info, entityHandle)
}

var _SetTransmitInfoOutOfPVSAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoOutOfPVSAll(__info)
}

// SetTransmitInfoOutOfPVSAll 
//  @brief Sets all bits in the TransmitOutOfPVS bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoOutOfPVSAll(info uintptr) {
	_SetTransmitInfoOutOfPVSAll(info)
}

var _ClearTransmitInfoOutOfPVSAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoOutOfPVSAll(__info)
}

// ClearTransmitInfoOutOfPVSAll 
//  @brief Clears all bits in the TransmitOutOfPVS bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoOutOfPVSAll(info uintptr) {
	_ClearTransmitInfoOutOfPVSAll(info)
}

var _SetTransmitInfoAlways = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.SetTransmitInfoAlways(__info, __entityHandle)
}

// SetTransmitInfoAlways 
//  @brief Sets a bit in the TransmitAlways bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func SetTransmitInfoAlways(info uintptr, entityHandle int32) {
	_SetTransmitInfoAlways(info, entityHandle)
}

var _ClearTransmitInfoAlways = func(info uintptr, entityHandle int32) {
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	C.ClearTransmitInfoAlways(__info, __entityHandle)
}

// ClearTransmitInfoAlways 
//  @brief Clears a bit in the TransmitAlways bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func ClearTransmitInfoAlways(info uintptr, entityHandle int32) {
	_ClearTransmitInfoAlways(info, entityHandle)
}

var _IsTransmitInfoAlwaysSet = func(info uintptr, entityHandle int32) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsTransmitInfoAlwaysSet(__info, __entityHandle))
	return __retVal
}

// IsTransmitInfoAlwaysSet 
//  @brief Checks if a bit is set in the TransmitAlways bitvec.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to check.
//
//  @return True if the entity is marked to always transmit, false otherwise.
func IsTransmitInfoAlwaysSet(info uintptr, entityHandle int32) bool {
	return _IsTransmitInfoAlwaysSet(info, entityHandle)
}

var _SetTransmitInfoAlwaysAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.SetTransmitInfoAlwaysAll(__info)
}

// SetTransmitInfoAlwaysAll 
//  @brief Sets all bits in the TransmitAlways bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func SetTransmitInfoAlwaysAll(info uintptr) {
	_SetTransmitInfoAlwaysAll(info)
}

var _ClearTransmitInfoAlwaysAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.ClearTransmitInfoAlwaysAll(__info)
}

// ClearTransmitInfoAlwaysAll 
//  @brief Clears all bits in the TransmitAlways bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func ClearTransmitInfoAlwaysAll(info uintptr) {
	_ClearTransmitInfoAlwaysAll(info)
}

var _GetTransmitInfoTargetSlotsCount = func(info uintptr) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__retVal = int32(C.GetTransmitInfoTargetSlotsCount(__info))
	return __retVal
}

// GetTransmitInfoTargetSlotsCount 
//  @brief Gets the count of target player slots.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The number of target player slots, or 0 if the info pointer is null.
func GetTransmitInfoTargetSlotsCount(info uintptr) int32 {
	return _GetTransmitInfoTargetSlotsCount(info)
}

var _GetTransmitInfoTargetSlot = func(info uintptr, index int32) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__index := C.int32_t(index)
	__retVal = int32(C.GetTransmitInfoTargetSlot(__info, __index))
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
	return _GetTransmitInfoTargetSlot(info, index)
}

var _AddTransmitInfoTargetSlot = func(info uintptr, playerSlot int32) {
	__info := C.uintptr_t(info)
	__playerSlot := C.int32_t(playerSlot)
	C.AddTransmitInfoTargetSlot(__info, __playerSlot)
}

// AddTransmitInfoTargetSlot 
//  @brief Adds a player slot to the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to add.
func AddTransmitInfoTargetSlot(info uintptr, playerSlot int32) {
	_AddTransmitInfoTargetSlot(info, playerSlot)
}

var _RemoveTransmitInfoTargetSlot = func(info uintptr, index int32) {
	__info := C.uintptr_t(info)
	__index := C.int32_t(index)
	C.RemoveTransmitInfoTargetSlot(__info, __index)
}

// RemoveTransmitInfoTargetSlot 
//  @brief Removes a player slot from the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param index: Index within the target slots vector to remove.
func RemoveTransmitInfoTargetSlot(info uintptr, index int32) {
	_RemoveTransmitInfoTargetSlot(info, index)
}

var _GetTransmitInfoTargetSlotsAll = func(info uintptr) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__info := C.uintptr_t(info)
	plugify.Block {
		Try: func() {
			__native := C.GetTransmitInfoTargetSlotsAll(__info)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetTransmitInfoTargetSlotsAll 
//  @brief Gets the target slots vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slots array.
func GetTransmitInfoTargetSlotsAll(info uintptr) []int32 {
	return _GetTransmitInfoTargetSlotsAll(info)
}

var _RemoveTransmitInfoTargetSlotsAll = func(info uintptr) {
	__info := C.uintptr_t(info)
	C.RemoveTransmitInfoTargetSlotsAll(__info)
}

// RemoveTransmitInfoTargetSlotsAll 
//  @brief Clears all target player slots from the vector.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func RemoveTransmitInfoTargetSlotsAll(info uintptr) {
	_RemoveTransmitInfoTargetSlotsAll(info)
}

var _GetTransmitInfoPlayerSlot = func(info uintptr) int32 {
	var __retVal int32
	__info := C.uintptr_t(info)
	__retVal = int32(C.GetTransmitInfoPlayerSlot(__info))
	return __retVal
}

// GetTransmitInfoPlayerSlot 
//  @brief Gets the player slot value from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return The player slot value, or -1 if info is null.
func GetTransmitInfoPlayerSlot(info uintptr) int32 {
	return _GetTransmitInfoPlayerSlot(info)
}

var _SetTransmitInfoPlayerSlot = func(info uintptr, playerSlot int32) {
	__info := C.uintptr_t(info)
	__playerSlot := C.int32_t(playerSlot)
	C.SetTransmitInfoPlayerSlot(__info, __playerSlot)
}

// SetTransmitInfoPlayerSlot 
//  @brief Sets the player slot value in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param playerSlot: The player slot value to set.
func SetTransmitInfoPlayerSlot(info uintptr, playerSlot int32) {
	_SetTransmitInfoPlayerSlot(info, playerSlot)
}

var _GetTransmitInfoFullUpdate = func(info uintptr) bool {
	var __retVal bool
	__info := C.uintptr_t(info)
	__retVal = bool(C.GetTransmitInfoFullUpdate(__info))
	return __retVal
}

// GetTransmitInfoFullUpdate 
//  @brief Gets the full update flag from the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//
//  @return True if full update is enabled, false otherwise.
func GetTransmitInfoFullUpdate(info uintptr) bool {
	return _GetTransmitInfoFullUpdate(info)
}

var _SetTransmitInfoFullUpdate = func(info uintptr, fullUpdate bool) {
	__info := C.uintptr_t(info)
	__fullUpdate := C.bool(fullUpdate)
	C.SetTransmitInfoFullUpdate(__info, __fullUpdate)
}

// SetTransmitInfoFullUpdate 
//  @brief Sets the full update flag in the CCheckTransmitInfo.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param fullUpdate: The full update flag value to set.
func SetTransmitInfoFullUpdate(info uintptr, fullUpdate bool) {
	_SetTransmitInfoFullUpdate(info, fullUpdate)
}

var _HideTransmitEntities = func(playerSlot int32, entHandles []int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entHandles := plugify.ConstructVectorInt32(entHandles)
	plugify.Block {
		Try: func() {
			C.HideTransmitEntities(__playerSlot, (*C.Vector)(unsafe.Pointer(&__entHandles)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__entHandles)
		},
	}.Do()
}

// HideTransmitEntities 
//  @brief Hides entities from a player's transmit list.
//
//  @param playerSlot: The player slot to hide entities from.
//  @param entHandles: Entity handles to hide.
func HideTransmitEntities(playerSlot int32, entHandles []int32) {
	_HideTransmitEntities(playerSlot, entHandles)
}

var _ShowTransmitEntities = func(playerSlot int32, entHandles []int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entHandles := plugify.ConstructVectorInt32(entHandles)
	plugify.Block {
		Try: func() {
			C.ShowTransmitEntities(__playerSlot, (*C.Vector)(unsafe.Pointer(&__entHandles)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__entHandles)
		},
	}.Do()
}

// ShowTransmitEntities 
//  @brief Shows previously hidden entities to a player.
//
//  @param playerSlot: The player slot to show entities to.
//  @param entHandles: Entity handles to show.
func ShowTransmitEntities(playerSlot int32, entHandles []int32) {
	_ShowTransmitEntities(playerSlot, entHandles)
}

var _GetHiddenTransmitEntities = func(playerSlot int32) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetHiddenTransmitEntities(__playerSlot)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetHiddenTransmitEntities 
//  @brief Gets all hidden entity handles for a player.
//
//  @param playerSlot: The player slot to query.
//
//  @return Array of hidden entity handles.
func GetHiddenTransmitEntities(playerSlot int32) []int32 {
	return _GetHiddenTransmitEntities(playerSlot)
}

var _HideTransmitEntity = func(playerSlot int32, entityHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entityHandle := C.int32_t(entityHandle)
	C.HideTransmitEntity(__playerSlot, __entityHandle)
}

// HideTransmitEntity 
//  @brief Hides a single entity from a player's transmit list.
//
//  @param playerSlot: The player slot to hide the entity from.
//  @param entityHandle: Entity handle to hide.
func HideTransmitEntity(playerSlot int32, entityHandle int32) {
	_HideTransmitEntity(playerSlot, entityHandle)
}

var _ShowTransmitEntity = func(playerSlot int32, entityHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entityHandle := C.int32_t(entityHandle)
	C.ShowTransmitEntity(__playerSlot, __entityHandle)
}

// ShowTransmitEntity 
//  @brief Shows a previously hidden entity to a player.
//
//  @param playerSlot: The player slot to show the entity to.
//  @param entityHandle: Entity handle to show.
func ShowTransmitEntity(playerSlot int32, entityHandle int32) {
	_ShowTransmitEntity(playerSlot, entityHandle)
}

var _HideTransmitEntityFromOtherPlayers = func(playerSlot int32, entityHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entityHandle := C.int32_t(entityHandle)
	C.HideTransmitEntityFromOtherPlayers(__playerSlot, __entityHandle)
}

// HideTransmitEntityFromOtherPlayers 
//  @brief Hides an entity from all players except the owner.
//
//  @param playerSlot: The owner player slot who will still see the entity.
//  @param entityHandle: Entity handle to hide from other players.
func HideTransmitEntityFromOtherPlayers(playerSlot int32, entityHandle int32) {
	_HideTransmitEntityFromOtherPlayers(playerSlot, entityHandle)
}

var _ShowTransmitEntityToOtherPlayers = func(playerSlot int32, entityHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__entityHandle := C.int32_t(entityHandle)
	C.ShowTransmitEntityToOtherPlayers(__playerSlot, __entityHandle)
}

// ShowTransmitEntityToOtherPlayers 
//  @brief Shows a previously hidden entity to all players except the owner.
//
//  @param playerSlot: The owner player slot who was excluded from hiding.
//  @param entityHandle: Entity handle to show to other players.
func ShowTransmitEntityToOtherPlayers(playerSlot int32, entityHandle int32) {
	_ShowTransmitEntityToOtherPlayers(playerSlot, entityHandle)
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
//  @brief Sets a bit in the TransmitAlways bitvec, marking an entity to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to mark as always transmittable.
func (w *CheckTransmitInfo) SetHLTV(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoAlways(w.handle, entityHandle)
	return nil
}

// ClearHLTV 
//  @brief Clears a bit in the TransmitAlways bitvec, unmarking an entity from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
//  @param entityHandle: The handle of the entity to unmark from always transmit.
func (w *CheckTransmitInfo) ClearHLTV(entityHandle int32) error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoAlways(w.handle, entityHandle)
	return nil
}

// IsHLTVSet 
//  @brief Checks if a bit is set in the TransmitAlways bitvec.
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
	return IsTransmitInfoAlwaysSet(w.handle, entityHandle), nil
}

// SetHLTVAll 
//  @brief Sets all bits in the TransmitAlways bitvec, marking all entities to always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) SetHLTVAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	SetTransmitInfoAlwaysAll(w.handle)
	return nil
}

// ClearHLTVAll 
//  @brief Clears all bits in the TransmitAlways bitvec, unmarking all entities from always transmit.
//
//  @param info: Pointer to the CCheckTransmitInfo structure.
func (w *CheckTransmitInfo) ClearHLTVAll() error {
	if w.handle == 0 {
		return CheckTransmitInfoErrEmptyHandle
	}
	ClearTransmitInfoAlwaysAll(w.handle)
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

