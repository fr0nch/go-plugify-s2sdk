package s2sdk

/*
#include "addons.h"
#cgo noescape AddServerAddon
#cgo noescape RemoveServerAddon
#cgo noescape GetServerAddons
#cgo noescape AddPrivateAddon
#cgo noescape RemovePrivateAddon
#cgo noescape GetPrivateAddons
#cgo noescape AddClientOnlyAddon
#cgo noescape RemoveClientOnlyAddon
#cgo noescape GetClientOnlyAddons
#cgo noescape SendAddon
#cgo noescape RemoveSentAddon
#cgo noescape ClearSentAddons
#cgo noescape GetSentAddons
#cgo noescape GetClientAddons
#cgo noescape GetCurrentWorkshopMap
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

// Generated from s2sdk (group: addons)

var _AddServerAddon = func(addon uint64, refresh bool) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.AddServerAddon(__addon, __refresh))
	return __retVal
}

// AddServerAddon 
//  @brief Adds a workshop addon that is mounted on the server and sent to all clients. Same as adding the addon to s2_extra_addons. If the addon is not downloaded yet, the server downloads it and reloads the map.
//
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true, the addon is mounted right away, otherwise on the next map load.
//
//  @return True if the addon was added, false if it is already in the list.
func AddServerAddon(addon uint64, refresh bool) bool {
	return _AddServerAddon(addon, refresh)
}

var _RemoveServerAddon = func(addon uint64, refresh bool) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.RemoveServerAddon(__addon, __refresh))
	return __retVal
}

// RemoveServerAddon 
//  @brief Removes a workshop addon from the server addons.
//
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true, the addon is unmounted right away, otherwise on the next map load.
//
//  @return True if the addon was removed, false if it is not in the list.
func RemoveServerAddon(addon uint64, refresh bool) bool {
	return _RemoveServerAddon(addon, refresh)
}

var _GetServerAddons = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block {
		Try: func() {
			__native := C.GetServerAddons()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetServerAddons 
//  @brief Retrieves the workshop addons that are mounted on the server and sent to all clients.
//
//
//  @return An array of workshop IDs from s2_extra_addons.
func GetServerAddons() []uint64 {
	return _GetServerAddons()
}

var _AddPrivateAddon = func(addon uint64, refresh bool) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.AddPrivateAddon(__addon, __refresh))
	return __retVal
}

// AddPrivateAddon 
//  @brief Adds a workshop addon that is mounted on the server but sent only to the clients passed to SendAddon. Same as adding the addon to s2_private_addons. If the addon is not downloaded yet, the server downloads it and reloads the map.
//
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true, the addon is mounted right away, otherwise on the next map load.
//
//  @return True if the addon was added, false if it is already in the list.
func AddPrivateAddon(addon uint64, refresh bool) bool {
	return _AddPrivateAddon(addon, refresh)
}

var _RemovePrivateAddon = func(addon uint64, refresh bool) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.RemovePrivateAddon(__addon, __refresh))
	return __retVal
}

// RemovePrivateAddon 
//  @brief Removes a workshop addon from the private addons.
//
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true, the addon is unmounted right away, otherwise on the next map load.
//
//  @return True if the addon was removed, false if it is not in the list.
func RemovePrivateAddon(addon uint64, refresh bool) bool {
	return _RemovePrivateAddon(addon, refresh)
}

var _GetPrivateAddons = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block {
		Try: func() {
			__native := C.GetPrivateAddons()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetPrivateAddons 
//  @brief Retrieves the workshop addons that are mounted on the server but sent only to selected clients.
//
//
//  @return An array of workshop IDs from s2_private_addons.
func GetPrivateAddons() []uint64 {
	return _GetPrivateAddons()
}

var _AddClientOnlyAddon = func(addon uint64, refresh bool) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.AddClientOnlyAddon(__addon, __refresh))
	return __retVal
}

// AddClientOnlyAddon 
//  @brief Adds a workshop addon that is sent to all clients but not mounted on the server. Same as adding the addon to s2_client_extra_addons. Use it only for content the server never loads itself, e.g. replacements of the base game files.
//
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true, connected clients that do not have the addon yet are told to download it right away.
//
//  @return True if the addon was added, false if it is already in the list.
func AddClientOnlyAddon(addon uint64, refresh bool) bool {
	return _AddClientOnlyAddon(addon, refresh)
}

var _RemoveClientOnlyAddon = func(addon uint64) bool {
	var __retVal bool
	__addon := C.uint64_t(addon)
	__retVal = bool(C.RemoveClientOnlyAddon(__addon))
	return __retVal
}

// RemoveClientOnlyAddon 
//  @brief Removes a workshop addon from the client only addons. Clients that already mounted the addon keep it until they reconnect.
//
//  @param addon: The workshop ID of the addon.
//
//  @return True if the addon was removed, false if it is not in the list.
func RemoveClientOnlyAddon(addon uint64) bool {
	return _RemoveClientOnlyAddon(addon)
}

var _GetClientOnlyAddons = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block {
		Try: func() {
			__native := C.GetClientOnlyAddons()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientOnlyAddons 
//  @brief Retrieves the workshop addons that are sent to all clients but not mounted on the server.
//
//
//  @return An array of workshop IDs from s2_client_extra_addons.
func GetClientOnlyAddons() []uint64 {
	return _GetClientOnlyAddons()
}

var _SendAddon = func(steamId uint64, addon uint64, refresh bool) bool {
	var __retVal bool
	__steamId := C.uint64_t(steamId)
	__addon := C.uint64_t(addon)
	__refresh := C.bool(refresh)
	__retVal = bool(C.SendAddon(__steamId, __addon, __refresh))
	return __retVal
}

// SendAddon 
//  @brief Sends a workshop addon to a specific client. Use it for private addons, or for client only addons that only this client should get. The addon stays sent to this SteamID64 across reconnects and map changes until it is removed with RemoveSentAddon or ClearSentAddons, or the server restarts, so it is enough to call this once, e.g. when the client gets the rights for it.
//
//  @param steamId: The SteamID64 of the client.
//  @param addon: The workshop ID of the addon.
//  @param refresh: If true and the client is connected, it is told to download the addon right away, otherwise it gets the addon on the next connect.
//
//  @return True if the addon was sent, false if the client already has it or the SteamID64 is 0.
func SendAddon(steamId uint64, addon uint64, refresh bool) bool {
	return _SendAddon(steamId, addon, refresh)
}

var _RemoveSentAddon = func(steamId uint64, addon uint64) bool {
	var __retVal bool
	__steamId := C.uint64_t(steamId)
	__addon := C.uint64_t(addon)
	__retVal = bool(C.RemoveSentAddon(__steamId, __addon))
	return __retVal
}

// RemoveSentAddon 
//  @brief Removes a workshop addon sent to a specific client. The client keeps the mounted addon until it reconnects.
//
//  @param steamId: The SteamID64 of the client.
//  @param addon: The workshop ID of the addon.
//
//  @return True if the addon was removed, false if it was not sent to the client.
func RemoveSentAddon(steamId uint64, addon uint64) bool {
	return _RemoveSentAddon(steamId, addon)
}

var _ClearSentAddons = func(steamId uint64) {
	__steamId := C.uint64_t(steamId)
	C.ClearSentAddons(__steamId)
}

// ClearSentAddons 
//  @brief Removes all workshop addons sent to a specific client. The client keeps the mounted addons until it reconnects.
//
//  @param steamId: The SteamID64 of the client.
func ClearSentAddons(steamId uint64) {
	_ClearSentAddons(steamId)
}

var _GetSentAddons = func(steamId uint64) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__steamId := C.uint64_t(steamId)
	plugify.Block {
		Try: func() {
			__native := C.GetSentAddons(__steamId)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetSentAddons 
//  @brief Retrieves the workshop addons sent to a specific client.
//
//  @param steamId: The SteamID64 of the client.
//
//  @return An array of workshop IDs sent with SendAddon or s2_sent_addons.
func GetSentAddons(steamId uint64) []uint64 {
	return _GetSentAddons(steamId)
}

var _GetClientAddons = func(steamId uint64) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__steamId := C.uint64_t(steamId)
	plugify.Block {
		Try: func() {
			__native := C.GetClientAddons(__steamId)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientAddons 
//  @brief Retrieves all workshop addons a client downloads and mounts when it connects. Includes the current workshop map, the server addons, the client only addons and the addons sent to this client.
//
//  @param steamId: The SteamID64 of the client.
//
//  @return An array of workshop IDs in the order the client mounts them.
func GetClientAddons(steamId uint64) []uint64 {
	return _GetClientAddons(steamId)
}

var _GetCurrentWorkshopMap = func() uint64 {
	__retVal := uint64(C.GetCurrentWorkshopMap())
	return __retVal
}

// GetCurrentWorkshopMap 
//  @brief Retrieves the workshop ID of the current map.
//
//
//  @return The workshop ID of the current map, or 0 if the current map is not a workshop map.
func GetCurrentWorkshopMap() uint64 {
	return _GetCurrentWorkshopMap()
}

