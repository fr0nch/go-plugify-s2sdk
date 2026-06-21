package s2sdk

/*
#include "clients.h"
#cgo noescape EntPointerToPlayerSlot
#cgo noescape PlayerSlotToEntPointer
#cgo noescape PlayerSlotToEntHandle
#cgo noescape PlayerSlotToClientPtr
#cgo noescape ClientPtrToPlayerSlot
#cgo noescape PlayerSlotToClientIndex
#cgo noescape ClientIndexToPlayerSlot
#cgo noescape PlayerServicesToPlayerSlot
#cgo noescape GetClientAuthId
#cgo noescape GetClientAccountId
#cgo noescape GetClientSteamID64
#cgo noescape GetClientIp
#cgo noescape GetClientLanguage
#cgo noescape GetClientOS
#cgo noescape GetClientName
#cgo noescape GetClientTime
#cgo noescape GetClientLatency
#cgo noescape GetUserFlagBits
#cgo noescape SetUserFlagBits
#cgo noescape AddUserFlags
#cgo noescape RemoveUserFlags
#cgo noescape IsClientAuthorized
#cgo noescape IsClientConnected
#cgo noescape IsClientInGame
#cgo noescape IsClientSourceTV
#cgo noescape IsClientAlive
#cgo noescape IsFakeClient
#cgo noescape GetClientMoveType
#cgo noescape SetClientMoveType
#cgo noescape GetClientGravity
#cgo noescape SetClientGravity
#cgo noescape GetClientFlags
#cgo noescape SetClientFlags
#cgo noescape GetClientRenderColor
#cgo noescape SetClientRenderColor
#cgo noescape GetClientRenderMode
#cgo noescape SetClientRenderMode
#cgo noescape GetClientMass
#cgo noescape SetClientMass
#cgo noescape GetClientFriction
#cgo noescape SetClientFriction
#cgo noescape GetClientHealth
#cgo noescape SetClientHealth
#cgo noescape GetClientMaxHealth
#cgo noescape SetClientMaxHealth
#cgo noescape GetClientTeam
#cgo noescape SetClientTeam
#cgo noescape GetClientAbsOrigin
#cgo noescape SetClientAbsOrigin
#cgo noescape GetClientAbsScale
#cgo noescape SetClientAbsScale
#cgo noescape GetClientAbsAngles
#cgo noescape SetClientAbsAngles
#cgo noescape GetClientLocalOrigin
#cgo noescape SetClientLocalOrigin
#cgo noescape GetClientLocalScale
#cgo noescape SetClientLocalScale
#cgo noescape GetClientLocalAngles
#cgo noescape SetClientLocalAngles
#cgo noescape GetClientAbsVelocity
#cgo noescape SetClientAbsVelocity
#cgo noescape GetClientBaseVelocity
#cgo noescape GetClientLocalAngVelocity
#cgo noescape GetClientAngVelocity
#cgo noescape SetClientAngVelocity
#cgo noescape GetClientLocalVelocity
#cgo noescape GetClientAngRotation
#cgo noescape SetClientAngRotation
#cgo noescape TransformPointClientToWorld
#cgo noescape TransformPointWorldToClient
#cgo noescape GetClientEyePosition
#cgo noescape GetClientEyeAngles
#cgo noescape SetClientForwardVector
#cgo noescape GetClientForwardVector
#cgo noescape GetClientLeftVector
#cgo noescape GetClientRightVector
#cgo noescape GetClientUpVector
#cgo noescape GetClientTransform
#cgo noescape GetClientModel
#cgo noescape SetClientModel
#cgo noescape GetClientWaterLevel
#cgo noescape GetClientGroundEntity
#cgo noescape GetClientEffects
#cgo noescape AddClientEffects
#cgo noescape RemoveClientEffects
#cgo noescape GetClientBoundingMaxs
#cgo noescape GetClientBoundingMins
#cgo noescape GetClientCenter
#cgo noescape TeleportClient
#cgo noescape ApplyAbsVelocityImpulseToClient
#cgo noescape ApplyLocalAngularVelocityImpulseToClient
#cgo noescape AcceptClientInput
#cgo noescape ConnectClientOutput
#cgo noescape DisconnectClientOutput
#cgo noescape DisconnectClientRedirectedOutput
#cgo noescape FireClientOutput
#cgo noescape RedirectClientOutput
#cgo noescape FollowClient
#cgo noescape FollowClientMerge
#cgo noescape TakeClientDamage
#cgo noescape GetClientPawn
#cgo noescape ProcessTargetString
#cgo noescape SwitchClientTeam
#cgo noescape ChangeClientTeam
#cgo noescape RespawnClient
#cgo noescape ForcePlayerSuicide
#cgo noescape KickClient
#cgo noescape BanClient
#cgo noescape BanIdentity
#cgo noescape GetClientActiveWeapon
#cgo noescape GetClientWeapons
#cgo noescape RemoveWeapons
#cgo noescape DropWeapon
#cgo noescape SelectWeapon
#cgo noescape SwitchWeapon
#cgo noescape RemoveWeapon
#cgo noescape GiveNamedItem
#cgo noescape GetClientButtons
#cgo noescape GetClientArmor
#cgo noescape SetClientArmor
#cgo noescape GetClientSpeed
#cgo noescape SetClientSpeed
#cgo noescape GetClientMoney
#cgo noescape SetClientMoney
#cgo noescape GetClientKills
#cgo noescape SetClientKills
#cgo noescape GetClientDeaths
#cgo noescape SetClientDeaths
#cgo noescape GetClientAssists
#cgo noescape SetClientAssists
#cgo noescape GetClientDamage
#cgo noescape SetClientDamage
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

// Generated from s2sdk (group: clients)

var P_EntPointerToPlayerSlot = func(entity uintptr) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__retVal = int32(C.EntPointerToPlayerSlot(__entity))
	return __retVal
}

// EntPointerToPlayerSlot 
//  @brief Retrieves the player slot from a given entity pointer.
//
//  @param entity: A pointer to the entity (CBaseEntity*).
//
//  @return The player slot if valid, otherwise -1.
func EntPointerToPlayerSlot(entity uintptr) int32 {
	return P_EntPointerToPlayerSlot(entity)
}

var P_PlayerSlotToEntPointer = func(playerSlot int32) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uintptr(C.PlayerSlotToEntPointer(__playerSlot))
	return __retVal
}

// PlayerSlotToEntPointer 
//  @brief Returns a pointer to the entity instance by player slot index.
//
//  @param playerSlot: Index of the player slot.
//
//  @return Pointer to the entity instance, or nullptr if the slot is invalid.
func PlayerSlotToEntPointer(playerSlot int32) uintptr {
	return P_PlayerSlotToEntPointer(playerSlot)
}

var P_PlayerSlotToEntHandle = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.PlayerSlotToEntHandle(__playerSlot))
	return __retVal
}

// PlayerSlotToEntHandle 
//  @brief Returns the entity handle associated with a player slot index.
//
//  @param playerSlot: Index of the player slot.
//
//  @return The index of the entity, or -1 if the handle is invalid.
func PlayerSlotToEntHandle(playerSlot int32) int32 {
	return P_PlayerSlotToEntHandle(playerSlot)
}

var P_PlayerSlotToClientPtr = func(playerSlot int32) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uintptr(C.PlayerSlotToClientPtr(__playerSlot))
	return __retVal
}

// PlayerSlotToClientPtr 
//  @brief Retrieves the client object from a given player slot.
//
//  @param playerSlot: The index of the player's slot (0-based).
//
//  @return A pointer to the client object if found, otherwise nullptr.
func PlayerSlotToClientPtr(playerSlot int32) uintptr {
	return P_PlayerSlotToClientPtr(playerSlot)
}

var P_ClientPtrToPlayerSlot = func(client uintptr) int32 {
	var __retVal int32
	__client := C.uintptr_t(client)
	__retVal = int32(C.ClientPtrToPlayerSlot(__client))
	return __retVal
}

// ClientPtrToPlayerSlot 
//  @brief Retrieves the index of a given client object.
//
//  @param client: A pointer to the client object (CServerSideClient*).
//
//  @return The player slot if found, otherwise -1.
func ClientPtrToPlayerSlot(client uintptr) int32 {
	return P_ClientPtrToPlayerSlot(client)
}

var P_PlayerSlotToClientIndex = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.PlayerSlotToClientIndex(__playerSlot))
	return __retVal
}

// PlayerSlotToClientIndex 
//  @brief Returns the entity index for a given player slot.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The entity index if valid, otherwise 0.
func PlayerSlotToClientIndex(playerSlot int32) int32 {
	return P_PlayerSlotToClientIndex(playerSlot)
}

var P_ClientIndexToPlayerSlot = func(clientIndex int32) int32 {
	var __retVal int32
	__clientIndex := C.int32_t(clientIndex)
	__retVal = int32(C.ClientIndexToPlayerSlot(__clientIndex))
	return __retVal
}

// ClientIndexToPlayerSlot 
//  @brief Retrieves the player slot from a given client index.
//
//  @param clientIndex: The index of the client.
//
//  @return The player slot if valid, otherwise -1.
func ClientIndexToPlayerSlot(clientIndex int32) int32 {
	return P_ClientIndexToPlayerSlot(clientIndex)
}

var P_PlayerServicesToPlayerSlot = func(service uintptr) int32 {
	var __retVal int32
	__service := C.uintptr_t(service)
	__retVal = int32(C.PlayerServicesToPlayerSlot(__service))
	return __retVal
}

// PlayerServicesToPlayerSlot 
//  @brief Retrieves the player slot from a given player service.
//
//  @param service: The service pointer. Like CCSPlayer_ItemServices, CCSPlayer_WeaponServices ect.
//
//  @return The player slot if valid, otherwise -1.
func PlayerServicesToPlayerSlot(service uintptr) int32 {
	return P_PlayerServicesToPlayerSlot(service)
}

var P_GetClientAuthId = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientAuthId(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientAuthId 
//  @brief Retrieves a client's authentication string (SteamID).
//
//  @param playerSlot: The index of the player's slot whose authentication string is being retrieved.
//
//  @return The authentication string.
func GetClientAuthId(playerSlot int32) string {
	return P_GetClientAuthId(playerSlot)
}

var P_GetClientAccountId = func(playerSlot int32) uint32 {
	var __retVal uint32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint32(C.GetClientAccountId(__playerSlot))
	return __retVal
}

// GetClientAccountId 
//  @brief Returns the client's Steam account ID, a unique number identifying a given Steam account.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return uint32_t The client's steam account ID.
func GetClientAccountId(playerSlot int32) uint32 {
	return P_GetClientAccountId(playerSlot)
}

var P_GetClientSteamID64 = func(playerSlot int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint64(C.GetClientSteamID64(__playerSlot))
	return __retVal
}

// GetClientSteamID64 
//  @brief Returns the client's SteamID64 - a unique 64-bit identifier of a Steam account.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return uint64_t The client's SteamID64.
func GetClientSteamID64(playerSlot int32) uint64 {
	return P_GetClientSteamID64(playerSlot)
}

var P_GetClientIp = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientIp(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientIp 
//  @brief Retrieves a client's IP address.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The client's IP address.
func GetClientIp(playerSlot int32) string {
	return P_GetClientIp(playerSlot)
}

var P_GetClientLanguage = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientLanguage(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientLanguage 
//  @brief Retrieves a client's language.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The client's language.
func GetClientLanguage(playerSlot int32) string {
	return P_GetClientLanguage(playerSlot)
}

var P_GetClientOS = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientOS(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientOS 
//  @brief Retrieves a client's operating system.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The client's operating system.
func GetClientOS(playerSlot int32) string {
	return P_GetClientOS(playerSlot)
}

var P_GetClientName = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientName(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientName 
//  @brief Returns the client's name.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The client's name.
func GetClientName(playerSlot int32) string {
	return P_GetClientName(playerSlot)
}

var P_GetClientTime = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientTime(__playerSlot))
	return __retVal
}

// GetClientTime 
//  @brief Returns the client's connection time in seconds.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return float Connection time in seconds.
func GetClientTime(playerSlot int32) float32 {
	return P_GetClientTime(playerSlot)
}

var P_GetClientLatency = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientLatency(__playerSlot))
	return __retVal
}

// GetClientLatency 
//  @brief Returns the client's current latency (RTT).
//
//  @param playerSlot: The index of the player's slot.
//
//  @return float Latency value.
func GetClientLatency(playerSlot int32) float32 {
	return P_GetClientLatency(playerSlot)
}

var P_GetUserFlagBits = func(playerSlot int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint64(C.GetUserFlagBits(__playerSlot))
	return __retVal
}

// GetUserFlagBits 
//  @brief Returns the client's access flags.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return uint64 Access flags as a bitmask.
func GetUserFlagBits(playerSlot int32) uint64 {
	return P_GetUserFlagBits(playerSlot)
}

var P_SetUserFlagBits = func(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.SetUserFlagBits(__playerSlot, __flags)
}

// SetUserFlagBits 
//  @brief Sets the access flags on a client using a bitmask.
//
//  @param playerSlot: The index of the player's slot.
//  @param flags: Bitmask representing the flags to be set.
func SetUserFlagBits(playerSlot int32, flags uint64) {
	P_SetUserFlagBits(playerSlot, flags)
}

var P_AddUserFlags = func(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.AddUserFlags(__playerSlot, __flags)
}

// AddUserFlags 
//  @brief Adds access flags to a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param flags: Bitmask representing the flags to be added.
func AddUserFlags(playerSlot int32, flags uint64) {
	P_AddUserFlags(playerSlot, flags)
}

var P_RemoveUserFlags = func(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.RemoveUserFlags(__playerSlot, __flags)
}

// RemoveUserFlags 
//  @brief Removes access flags from a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param flags: Bitmask representing the flags to be removed.
func RemoveUserFlags(playerSlot int32, flags uint64) {
	P_RemoveUserFlags(playerSlot, flags)
}

var P_IsClientAuthorized = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientAuthorized(__playerSlot))
	return __retVal
}

// IsClientAuthorized 
//  @brief Checks if a certain player has been authenticated.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the player is authenticated, false otherwise.
func IsClientAuthorized(playerSlot int32) bool {
	return P_IsClientAuthorized(playerSlot)
}

var P_IsClientConnected = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientConnected(__playerSlot))
	return __retVal
}

// IsClientConnected 
//  @brief Checks if a certain player is connected.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the player is connected, false otherwise.
func IsClientConnected(playerSlot int32) bool {
	return P_IsClientConnected(playerSlot)
}

var P_IsClientInGame = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientInGame(__playerSlot))
	return __retVal
}

// IsClientInGame 
//  @brief Checks if a certain player has entered the game.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the player is in the game, false otherwise.
func IsClientInGame(playerSlot int32) bool {
	return P_IsClientInGame(playerSlot)
}

var P_IsClientSourceTV = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientSourceTV(__playerSlot))
	return __retVal
}

// IsClientSourceTV 
//  @brief Checks if a certain player is the SourceTV bot.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the client is the SourceTV bot, false otherwise.
func IsClientSourceTV(playerSlot int32) bool {
	return P_IsClientSourceTV(playerSlot)
}

var P_IsClientAlive = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientAlive(__playerSlot))
	return __retVal
}

// IsClientAlive 
//  @brief Checks if the client is alive or dead.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the client is alive, false if dead.
func IsClientAlive(playerSlot int32) bool {
	return P_IsClientAlive(playerSlot)
}

var P_IsFakeClient = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsFakeClient(__playerSlot))
	return __retVal
}

// IsFakeClient 
//  @brief Checks if a certain player is a fake client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return true if the client is a fake client, false otherwise.
func IsFakeClient(playerSlot int32) bool {
	return P_IsFakeClient(playerSlot)
}

var P_GetClientMoveType = func(playerSlot int32) MoveType {
	var __retVal MoveType
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMoveType(__playerSlot))
	return __retVal
}

// GetClientMoveType 
//  @brief Retrieves the movement type of an client.
//
//  @param playerSlot: The index of the player's slot whose movement type is to be retrieved.
//
//  @return The movement type of the entity, or 0 if the entity is invalid.
func GetClientMoveType(playerSlot int32) MoveType {
	return P_GetClientMoveType(playerSlot)
}

var P_SetClientMoveType = func(playerSlot int32, moveType MoveType) {
	__playerSlot := C.int32_t(playerSlot)
	__moveType := C.int32_t(moveType)
	C.SetClientMoveType(__playerSlot, __moveType)
}

// SetClientMoveType 
//  @brief Sets the movement type of an client.
//
//  @param playerSlot: The index of the player's slot whose movement type is to be set.
//  @param moveType: The movement type of the entity, or 0 if the entity is invalid.
func SetClientMoveType(playerSlot int32, moveType MoveType) {
	P_SetClientMoveType(playerSlot, moveType)
}

var P_GetClientGravity = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientGravity(__playerSlot))
	return __retVal
}

// GetClientGravity 
//  @brief Retrieves the gravity scale of an client.
//
//  @param playerSlot: The index of the player's slot whose gravity scale is to be retrieved.
//
//  @return The gravity scale of the client, or 0.0f if the client is invalid.
func GetClientGravity(playerSlot int32) float32 {
	return P_GetClientGravity(playerSlot)
}

var P_SetClientGravity = func(playerSlot int32, gravity float32) {
	__playerSlot := C.int32_t(playerSlot)
	__gravity := C.float(gravity)
	C.SetClientGravity(__playerSlot, __gravity)
}

// SetClientGravity 
//  @brief Sets the gravity scale of an client.
//
//  @param playerSlot: The index of the player's slot whose gravity scale is to be set.
//  @param gravity: The new gravity scale to set for the client.
func SetClientGravity(playerSlot int32, gravity float32) {
	P_SetClientGravity(playerSlot, gravity)
}

var P_GetClientFlags = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientFlags(__playerSlot))
	return __retVal
}

// GetClientFlags 
//  @brief Retrieves the flags of an client.
//
//  @param playerSlot: The index of the player's slot whose flags are to be retrieved.
//
//  @return The flags of the client, or 0 if the client is invalid.
func GetClientFlags(playerSlot int32) int32 {
	return P_GetClientFlags(playerSlot)
}

var P_SetClientFlags = func(playerSlot int32, flags int32) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.int32_t(flags)
	C.SetClientFlags(__playerSlot, __flags)
}

// SetClientFlags 
//  @brief Sets the flags of an client.
//
//  @param playerSlot: The index of the player's slot whose flags are to be set.
//  @param flags: The new flags to set for the client.
func SetClientFlags(playerSlot int32, flags int32) {
	P_SetClientFlags(playerSlot, flags)
}

var P_GetClientRenderColor = func(playerSlot int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientRenderColor(__playerSlot)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientRenderColor 
//  @brief Retrieves the render color of an client.
//
//  @param playerSlot: The index of the player's slot whose render color is to be retrieved.
//
//  @return The raw color value of the client's render color, or 0 if the client is invalid.
func GetClientRenderColor(playerSlot int32) plugify.Vector4 {
	return P_GetClientRenderColor(playerSlot)
}

var P_SetClientRenderColor = func(playerSlot int32, color plugify.Vector4) {
	__playerSlot := C.int32_t(playerSlot)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	C.SetClientRenderColor(__playerSlot, &__color)
}

// SetClientRenderColor 
//  @brief Sets the render color of an client.
//
//  @param playerSlot: The index of the player's slot whose render color is to be set.
//  @param color: The new raw color value to set for the client's render color.
func SetClientRenderColor(playerSlot int32, color plugify.Vector4) {
	P_SetClientRenderColor(playerSlot, color)
}

var P_GetClientRenderMode = func(playerSlot int32) RenderMode {
	var __retVal RenderMode
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint8(C.GetClientRenderMode(__playerSlot))
	return __retVal
}

// GetClientRenderMode 
//  @brief Retrieves the render mode of an client.
//
//  @param playerSlot: The index of the player's slot whose render mode is to be retrieved.
//
//  @return The render mode of the client, or 0 if the client is invalid.
func GetClientRenderMode(playerSlot int32) RenderMode {
	return P_GetClientRenderMode(playerSlot)
}

var P_SetClientRenderMode = func(playerSlot int32, renderMode RenderMode) {
	__playerSlot := C.int32_t(playerSlot)
	__renderMode := C.uint8_t(renderMode)
	C.SetClientRenderMode(__playerSlot, __renderMode)
}

// SetClientRenderMode 
//  @brief Sets the render mode of an client.
//
//  @param playerSlot: The index of the player's slot whose render mode is to be set.
//  @param renderMode: The new render mode to set for the client.
func SetClientRenderMode(playerSlot int32, renderMode RenderMode) {
	P_SetClientRenderMode(playerSlot, renderMode)
}

var P_GetClientMass = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMass(__playerSlot))
	return __retVal
}

// GetClientMass 
//  @brief Retrieves the mass of an client.
//
//  @param playerSlot: The index of the player's slot whose mass is to be retrieved.
//
//  @return The mass of the client, or 0 if the client is invalid.
func GetClientMass(playerSlot int32) int32 {
	return P_GetClientMass(playerSlot)
}

var P_SetClientMass = func(playerSlot int32, mass int32) {
	__playerSlot := C.int32_t(playerSlot)
	__mass := C.int32_t(mass)
	C.SetClientMass(__playerSlot, __mass)
}

// SetClientMass 
//  @brief Sets the mass of an client.
//
//  @param playerSlot: The index of the player's slot whose mass is to be set.
//  @param mass: The new mass value to set for the client.
func SetClientMass(playerSlot int32, mass int32) {
	P_SetClientMass(playerSlot, mass)
}

var P_GetClientFriction = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientFriction(__playerSlot))
	return __retVal
}

// GetClientFriction 
//  @brief Retrieves the friction of an client.
//
//  @param playerSlot: The index of the player's slot whose friction is to be retrieved.
//
//  @return The friction of the client, or 0 if the client is invalid.
func GetClientFriction(playerSlot int32) float32 {
	return P_GetClientFriction(playerSlot)
}

var P_SetClientFriction = func(playerSlot int32, friction float32) {
	__playerSlot := C.int32_t(playerSlot)
	__friction := C.float(friction)
	C.SetClientFriction(__playerSlot, __friction)
}

// SetClientFriction 
//  @brief Sets the friction of an client.
//
//  @param playerSlot: The index of the player's slot whose friction is to be set.
//  @param friction: The new friction value to set for the client.
func SetClientFriction(playerSlot int32, friction float32) {
	P_SetClientFriction(playerSlot, friction)
}

var P_GetClientHealth = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientHealth(__playerSlot))
	return __retVal
}

// GetClientHealth 
//  @brief Retrieves the health of an client.
//
//  @param playerSlot: The index of the player's slot whose health is to be retrieved.
//
//  @return The health of the client, or 0 if the client is invalid.
func GetClientHealth(playerSlot int32) int32 {
	return P_GetClientHealth(playerSlot)
}

var P_SetClientHealth = func(playerSlot int32, health int32) {
	__playerSlot := C.int32_t(playerSlot)
	__health := C.int32_t(health)
	C.SetClientHealth(__playerSlot, __health)
}

// SetClientHealth 
//  @brief Sets the health of an client.
//
//  @param playerSlot: The index of the player's slot whose health is to be set.
//  @param health: The new health value to set for the client.
func SetClientHealth(playerSlot int32, health int32) {
	P_SetClientHealth(playerSlot, health)
}

var P_GetClientMaxHealth = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMaxHealth(__playerSlot))
	return __retVal
}

// GetClientMaxHealth 
//  @brief Retrieves the max health of an client.
//
//  @param playerSlot: The index of the player's slot whose max health is to be retrieved.
//
//  @return The max health of the client, or 0 if the client is invalid.
func GetClientMaxHealth(playerSlot int32) int32 {
	return P_GetClientMaxHealth(playerSlot)
}

var P_SetClientMaxHealth = func(playerSlot int32, maxHealth int32) {
	__playerSlot := C.int32_t(playerSlot)
	__maxHealth := C.int32_t(maxHealth)
	C.SetClientMaxHealth(__playerSlot, __maxHealth)
}

// SetClientMaxHealth 
//  @brief Sets the max health of an client.
//
//  @param playerSlot: The index of the player's slot whose max health is to be set.
//  @param maxHealth: The new max health value to set for the client.
func SetClientMaxHealth(playerSlot int32, maxHealth int32) {
	P_SetClientMaxHealth(playerSlot, maxHealth)
}

var P_GetClientTeam = func(playerSlot int32) CSTeam {
	var __retVal CSTeam
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientTeam(__playerSlot))
	return __retVal
}

// GetClientTeam 
//  @brief Retrieves the team number of an client.
//
//  @param playerSlot: The index of the player's slot whose team number is to be retrieved.
//
//  @return The team number of the client, or 0 if the client is invalid.
func GetClientTeam(playerSlot int32) CSTeam {
	return P_GetClientTeam(playerSlot)
}

var P_SetClientTeam = func(playerSlot int32, team CSTeam) {
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	C.SetClientTeam(__playerSlot, __team)
}

// SetClientTeam 
//  @brief Sets the team number of an client.
//
//  @param playerSlot: The index of the player's slot whose team number is to be set.
//  @param team: The new team number to set for the client.
func SetClientTeam(playerSlot int32, team CSTeam) {
	P_SetClientTeam(playerSlot, team)
}

var P_GetClientAbsOrigin = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAbsOrigin(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAbsOrigin 
//  @brief Retrieves the absolute origin of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute origin is to be retrieved.
//
//  @return A vector where the absolute origin will be stored.
func GetClientAbsOrigin(playerSlot int32) plugify.Vector3 {
	return P_GetClientAbsOrigin(playerSlot)
}

var P_SetClientAbsOrigin = func(playerSlot int32, origin plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	C.SetClientAbsOrigin(__playerSlot, &__origin)
}

// SetClientAbsOrigin 
//  @brief Sets the absolute origin of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute origin is to be set.
//  @param origin: The new absolute origin to set for the client.
func SetClientAbsOrigin(playerSlot int32, origin plugify.Vector3) {
	P_SetClientAbsOrigin(playerSlot, origin)
}

var P_GetClientAbsScale = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientAbsScale(__playerSlot))
	return __retVal
}

// GetClientAbsScale 
//  @brief Retrieves the absolute scale of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute scale is to be retrieved.
//
//  @return A vector where the absolute scale will be stored.
func GetClientAbsScale(playerSlot int32) float32 {
	return P_GetClientAbsScale(playerSlot)
}

var P_SetClientAbsScale = func(playerSlot int32, scale float32) {
	__playerSlot := C.int32_t(playerSlot)
	__scale := C.float(scale)
	C.SetClientAbsScale(__playerSlot, __scale)
}

// SetClientAbsScale 
//  @brief Sets the absolute scale of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute scale is to be set.
//  @param scale: The new absolute scale to set for the client.
func SetClientAbsScale(playerSlot int32, scale float32) {
	P_SetClientAbsScale(playerSlot, scale)
}

var P_GetClientAbsAngles = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAbsAngles(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAbsAngles 
//  @brief Retrieves the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be retrieved.
//
//  @return A QAngle where the angular rotation will be stored.
func GetClientAbsAngles(playerSlot int32) plugify.Vector3 {
	return P_GetClientAbsAngles(playerSlot)
}

var P_SetClientAbsAngles = func(playerSlot int32, angle plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__angle := *(*C.Vector3)(unsafe.Pointer(&angle))
	C.SetClientAbsAngles(__playerSlot, &__angle)
}

// SetClientAbsAngles 
//  @brief Sets the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be set.
//  @param angle: The new angular rotation to set for the client.
func SetClientAbsAngles(playerSlot int32, angle plugify.Vector3) {
	P_SetClientAbsAngles(playerSlot, angle)
}

var P_GetClientLocalOrigin = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientLocalOrigin(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientLocalOrigin 
//  @brief Retrieves the local origin of an client.
//
//  @param playerSlot: The index of the player's slot whose local origin is to be retrieved.
//
//  @return A vector where the local origin will be stored.
func GetClientLocalOrigin(playerSlot int32) plugify.Vector3 {
	return P_GetClientLocalOrigin(playerSlot)
}

var P_SetClientLocalOrigin = func(playerSlot int32, origin plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	C.SetClientLocalOrigin(__playerSlot, &__origin)
}

// SetClientLocalOrigin 
//  @brief Sets the local origin of an client.
//
//  @param playerSlot: The index of the player's slot whose local origin is to be set.
//  @param origin: The new local origin to set for the client.
func SetClientLocalOrigin(playerSlot int32, origin plugify.Vector3) {
	P_SetClientLocalOrigin(playerSlot, origin)
}

var P_GetClientLocalScale = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientLocalScale(__playerSlot))
	return __retVal
}

// GetClientLocalScale 
//  @brief Retrieves the local scale of an client.
//
//  @param playerSlot: The index of the player's slot whose local scale is to be retrieved.
//
//  @return A vector where the local scale will be stored.
func GetClientLocalScale(playerSlot int32) float32 {
	return P_GetClientLocalScale(playerSlot)
}

var P_SetClientLocalScale = func(playerSlot int32, scale float32) {
	__playerSlot := C.int32_t(playerSlot)
	__scale := C.float(scale)
	C.SetClientLocalScale(__playerSlot, __scale)
}

// SetClientLocalScale 
//  @brief Sets the local scale of an client.
//
//  @param playerSlot: The index of the player's slot whose local scale is to be set.
//  @param scale: The new local scale to set for the client.
func SetClientLocalScale(playerSlot int32, scale float32) {
	P_SetClientLocalScale(playerSlot, scale)
}

var P_GetClientLocalAngles = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientLocalAngles(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientLocalAngles 
//  @brief Retrieves the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be retrieved.
//
//  @return A QAngle where the angular rotation will be stored.
func GetClientLocalAngles(playerSlot int32) plugify.Vector3 {
	return P_GetClientLocalAngles(playerSlot)
}

var P_SetClientLocalAngles = func(playerSlot int32, angle plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__angle := *(*C.Vector3)(unsafe.Pointer(&angle))
	C.SetClientLocalAngles(__playerSlot, &__angle)
}

// SetClientLocalAngles 
//  @brief Sets the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be set.
//  @param angle: The new angular rotation to set for the client.
func SetClientLocalAngles(playerSlot int32, angle plugify.Vector3) {
	P_SetClientLocalAngles(playerSlot, angle)
}

var P_GetClientAbsVelocity = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAbsVelocity(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAbsVelocity 
//  @brief Retrieves the absolute velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute velocity is to be retrieved.
//
//  @return A vector where the absolute velocity will be stored.
func GetClientAbsVelocity(playerSlot int32) plugify.Vector3 {
	return P_GetClientAbsVelocity(playerSlot)
}

var P_SetClientAbsVelocity = func(playerSlot int32, velocity plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.SetClientAbsVelocity(__playerSlot, &__velocity)
}

// SetClientAbsVelocity 
//  @brief Sets the absolute velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose absolute velocity is to be set.
//  @param velocity: The new absolute velocity to set for the client.
func SetClientAbsVelocity(playerSlot int32, velocity plugify.Vector3) {
	P_SetClientAbsVelocity(playerSlot, velocity)
}

var P_GetClientBaseVelocity = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientBaseVelocity(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientBaseVelocity 
//  @brief Retrieves the base velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose base velocity is to be retrieved.
//
//  @return A vector where the base velocity will be stored.
func GetClientBaseVelocity(playerSlot int32) plugify.Vector3 {
	return P_GetClientBaseVelocity(playerSlot)
}

var P_GetClientLocalAngVelocity = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientLocalAngVelocity(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientLocalAngVelocity 
//  @brief Retrieves the local angular velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose local angular velocity is to be retrieved.
//
//  @return A vector where the local angular velocity will be stored.
func GetClientLocalAngVelocity(playerSlot int32) plugify.Vector3 {
	return P_GetClientLocalAngVelocity(playerSlot)
}

var P_GetClientAngVelocity = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAngVelocity(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAngVelocity 
//  @brief Retrieves the angular velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose angular velocity is to be retrieved.
//
//  @return A vector where the angular velocity will be stored.
func GetClientAngVelocity(playerSlot int32) plugify.Vector3 {
	return P_GetClientAngVelocity(playerSlot)
}

var P_SetClientAngVelocity = func(playerSlot int32, velocity plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.SetClientAngVelocity(__playerSlot, &__velocity)
}

// SetClientAngVelocity 
//  @brief Sets the angular velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose angular velocity is to be set.
//  @param velocity: The new angular velocity to set for the client.
func SetClientAngVelocity(playerSlot int32, velocity plugify.Vector3) {
	P_SetClientAngVelocity(playerSlot, velocity)
}

var P_GetClientLocalVelocity = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientLocalVelocity(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientLocalVelocity 
//  @brief Retrieves the local velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose local velocity is to be retrieved.
//
//  @return A vector where the local velocity will be stored.
func GetClientLocalVelocity(playerSlot int32) plugify.Vector3 {
	return P_GetClientLocalVelocity(playerSlot)
}

var P_GetClientAngRotation = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAngRotation(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAngRotation 
//  @brief Retrieves the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be retrieved.
//
//  @return A vector where the angular rotation will be stored.
func GetClientAngRotation(playerSlot int32) plugify.Vector3 {
	return P_GetClientAngRotation(playerSlot)
}

var P_SetClientAngRotation = func(playerSlot int32, rotation plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__rotation := *(*C.Vector3)(unsafe.Pointer(&rotation))
	C.SetClientAngRotation(__playerSlot, &__rotation)
}

// SetClientAngRotation 
//  @brief Sets the angular rotation of an client.
//
//  @param playerSlot: The index of the player's slot whose angular rotation is to be set.
//  @param rotation: The new angular rotation to set for the client.
func SetClientAngRotation(playerSlot int32, rotation plugify.Vector3) {
	P_SetClientAngRotation(playerSlot, rotation)
}

var P_TransformPointClientToWorld = func(playerSlot int32, point plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__point := *(*C.Vector3)(unsafe.Pointer(&point))
	__native := C.TransformPointClientToWorld(__playerSlot, &__point)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// TransformPointClientToWorld 
//  @brief Returns the input Vector transformed from client to world space.
//
//  @param playerSlot: The index of the player's slot
//  @param point: Point in client local space to transform
//
//  @return The point transformed to world space coordinates
func TransformPointClientToWorld(playerSlot int32, point plugify.Vector3) plugify.Vector3 {
	return P_TransformPointClientToWorld(playerSlot, point)
}

var P_TransformPointWorldToClient = func(playerSlot int32, point plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__point := *(*C.Vector3)(unsafe.Pointer(&point))
	__native := C.TransformPointWorldToClient(__playerSlot, &__point)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// TransformPointWorldToClient 
//  @brief Returns the input Vector transformed from world to client space.
//
//  @param playerSlot: The index of the player's slot
//  @param point: Point in world space to transform
//
//  @return The point transformed to client local space coordinates
func TransformPointWorldToClient(playerSlot int32, point plugify.Vector3) plugify.Vector3 {
	return P_TransformPointWorldToClient(playerSlot, point)
}

var P_GetClientEyePosition = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientEyePosition(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientEyePosition 
//  @brief Get vector to eye position - absolute coords.
//
//  @param playerSlot: The index of the player's slot
//
//  @return Eye position in absolute/world coordinates
func GetClientEyePosition(playerSlot int32) plugify.Vector3 {
	return P_GetClientEyePosition(playerSlot)
}

var P_GetClientEyeAngles = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientEyeAngles(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientEyeAngles 
//  @brief Get the qangles that this client is looking at.
//
//  @param playerSlot: The index of the player's slot
//
//  @return Eye angles as a vector (pitch, yaw, roll)
func GetClientEyeAngles(playerSlot int32) plugify.Vector3 {
	return P_GetClientEyeAngles(playerSlot)
}

var P_SetClientForwardVector = func(playerSlot int32, forward plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__forward := *(*C.Vector3)(unsafe.Pointer(&forward))
	C.SetClientForwardVector(__playerSlot, &__forward)
}

// SetClientForwardVector 
//  @brief Sets the forward velocity of an client.
//
//  @param playerSlot: The index of the player's slot whose forward velocity is to be set.
func SetClientForwardVector(playerSlot int32, forward plugify.Vector3) {
	P_SetClientForwardVector(playerSlot, forward)
}

var P_GetClientForwardVector = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientForwardVector(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientForwardVector 
//  @brief Get the forward vector of the client.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Forward-facing direction vector of the client
func GetClientForwardVector(playerSlot int32) plugify.Vector3 {
	return P_GetClientForwardVector(playerSlot)
}

var P_GetClientLeftVector = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientLeftVector(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientLeftVector 
//  @brief Get the left vector of the client.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Left-facing direction vector of the client (aligned with the y axis)
func GetClientLeftVector(playerSlot int32) plugify.Vector3 {
	return P_GetClientLeftVector(playerSlot)
}

var P_GetClientRightVector = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientRightVector(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientRightVector 
//  @brief Get the right vector of the client.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Right-facing direction vector of the client
func GetClientRightVector(playerSlot int32) plugify.Vector3 {
	return P_GetClientRightVector(playerSlot)
}

var P_GetClientUpVector = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientUpVector(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientUpVector 
//  @brief Get the up vector of the client.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Up-facing direction vector of the client
func GetClientUpVector(playerSlot int32) plugify.Vector3 {
	return P_GetClientUpVector(playerSlot)
}

var P_GetClientTransform = func(playerSlot int32) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientTransform(__playerSlot)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientTransform 
//  @brief Get the client-to-world transformation matrix.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return 4x4 transformation matrix representing client's position, rotation, and scale in world space
func GetClientTransform(playerSlot int32) plugify.Matrix4x4 {
	return P_GetClientTransform(playerSlot)
}

var P_GetClientModel = func(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientModel(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientModel 
//  @brief Retrieves the model name of an client.
//
//  @param playerSlot: The index of the player's slot whose model name is to be retrieved.
//
//  @return A string where the model name will be stored.
func GetClientModel(playerSlot int32) string {
	return P_GetClientModel(playerSlot)
}

var P_SetClientModel = func(playerSlot int32, model string) {
	__playerSlot := C.int32_t(playerSlot)
	__model := plugify.ConstructString(model)
	plugify.Block {
		Try: func() {
			C.SetClientModel(__playerSlot, (*C.String)(unsafe.Pointer(&__model)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
}

// SetClientModel 
//  @brief Sets the model name of an client.
//
//  @param playerSlot: The index of the player's slot whose model name is to be set.
//  @param model: The new model name to set for the client.
func SetClientModel(playerSlot int32, model string) {
	P_SetClientModel(playerSlot, model)
}

var P_GetClientWaterLevel = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientWaterLevel(__playerSlot))
	return __retVal
}

// GetClientWaterLevel 
//  @brief Retrieves the water level of an client.
//
//  @param playerSlot: The index of the player's slot whose water level is to be retrieved.
//
//  @return The water level of the client, or 0.0f if the client is invalid.
func GetClientWaterLevel(playerSlot int32) float32 {
	return P_GetClientWaterLevel(playerSlot)
}

var P_GetClientGroundEntity = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientGroundEntity(__playerSlot))
	return __retVal
}

// GetClientGroundEntity 
//  @brief Retrieves the ground client of an client.
//
//  @param playerSlot: The index of the player's slot whose ground client is to be retrieved.
//
//  @return The handle of the ground client, or INVALID_EHANDLE_INDEX if the client is invalid.
func GetClientGroundEntity(playerSlot int32) int32 {
	return P_GetClientGroundEntity(playerSlot)
}

var P_GetClientEffects = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientEffects(__playerSlot))
	return __retVal
}

// GetClientEffects 
//  @brief Retrieves the effects of an client.
//
//  @param playerSlot: The index of the player's slot whose effects are to be retrieved.
//
//  @return The effect flags of the client, or 0 if the client is invalid.
func GetClientEffects(playerSlot int32) int32 {
	return P_GetClientEffects(playerSlot)
}

var P_AddClientEffects = func(playerSlot int32, effects int32) {
	__playerSlot := C.int32_t(playerSlot)
	__effects := C.int32_t(effects)
	C.AddClientEffects(__playerSlot, __effects)
}

// AddClientEffects 
//  @brief Adds the render effect flag to an client.
//
//  @param playerSlot: The index of the player's slot to modify
//  @param effects: Render effect flags to add
func AddClientEffects(playerSlot int32, effects int32) {
	P_AddClientEffects(playerSlot, effects)
}

var P_RemoveClientEffects = func(playerSlot int32, effects int32) {
	__playerSlot := C.int32_t(playerSlot)
	__effects := C.int32_t(effects)
	C.RemoveClientEffects(__playerSlot, __effects)
}

// RemoveClientEffects 
//  @brief Removes the render effect flag from an client.
//
//  @param playerSlot: The index of the player's slot to modify
//  @param effects: Render effect flags to remove
func RemoveClientEffects(playerSlot int32, effects int32) {
	P_RemoveClientEffects(playerSlot, effects)
}

var P_GetClientBoundingMaxs = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientBoundingMaxs(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientBoundingMaxs 
//  @brief Get a vector containing max bounds, centered on object.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Vector containing the maximum bounds of the client's bounding box
func GetClientBoundingMaxs(playerSlot int32) plugify.Vector3 {
	return P_GetClientBoundingMaxs(playerSlot)
}

var P_GetClientBoundingMins = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientBoundingMins(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientBoundingMins 
//  @brief Get a vector containing min bounds, centered on object.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Vector containing the minimum bounds of the client's bounding box
func GetClientBoundingMins(playerSlot int32) plugify.Vector3 {
	return P_GetClientBoundingMins(playerSlot)
}

var P_GetClientCenter = func(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientCenter(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientCenter 
//  @brief Get vector to center of object - absolute coords.
//
//  @param playerSlot: The index of the player's slot to query
//
//  @return Vector pointing to the center of the client in absolute/world coordinates
func GetClientCenter(playerSlot int32) plugify.Vector3 {
	return P_GetClientCenter(playerSlot)
}

var P_TeleportClient = func(playerSlot int32, origin plugify.Vector3, angles plugify.Vector3, velocity plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__angles := *(*C.Vector3)(unsafe.Pointer(&angles))
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.TeleportClient(__playerSlot, &__origin, &__angles, &__velocity)
}

// TeleportClient 
//  @brief Teleports an client to a specified location and orientation.
//
//  @param playerSlot: The index of the player's slot to teleport.
//  @param origin: A pointer to a Vector representing the new absolute position. Use nan vector to not set.
//  @param angles: A pointer to a QAngle representing the new orientation. Use nan vector to not set.
//  @param velocity: A pointer to a Vector representing the new velocity. Use nan vector to not set.
func TeleportClient(playerSlot int32, origin plugify.Vector3, angles plugify.Vector3, velocity plugify.Vector3) {
	P_TeleportClient(playerSlot, origin, angles, velocity)
}

var P_ApplyAbsVelocityImpulseToClient = func(playerSlot int32, vecImpulse plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__vecImpulse := *(*C.Vector3)(unsafe.Pointer(&vecImpulse))
	C.ApplyAbsVelocityImpulseToClient(__playerSlot, &__vecImpulse)
}

// ApplyAbsVelocityImpulseToClient 
//  @brief Apply an absolute velocity impulse to an client.
//
//  @param playerSlot: The index of the player's slot to apply impulse to
//  @param vecImpulse: Velocity impulse vector to apply
func ApplyAbsVelocityImpulseToClient(playerSlot int32, vecImpulse plugify.Vector3) {
	P_ApplyAbsVelocityImpulseToClient(playerSlot, vecImpulse)
}

var P_ApplyLocalAngularVelocityImpulseToClient = func(playerSlot int32, angImpulse plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__angImpulse := *(*C.Vector3)(unsafe.Pointer(&angImpulse))
	C.ApplyLocalAngularVelocityImpulseToClient(__playerSlot, &__angImpulse)
}

// ApplyLocalAngularVelocityImpulseToClient 
//  @brief Apply a local angular velocity impulse to an client.
//
//  @param playerSlot: The index of the player's slot to apply impulse to
//  @param angImpulse: Angular velocity impulse vector to apply
func ApplyLocalAngularVelocityImpulseToClient(playerSlot int32, angImpulse plugify.Vector3) {
	P_ApplyLocalAngularVelocityImpulseToClient(playerSlot, angImpulse)
}

var P_AcceptClientInput = func(playerSlot int32, inputName string, activatorHandle int32, callerHandle int32, value any, type_ FieldType, outputId int32) {
	__playerSlot := C.int32_t(playerSlot)
	__inputName := plugify.ConstructString(inputName)
	__activatorHandle := C.int32_t(activatorHandle)
	__callerHandle := C.int32_t(callerHandle)
	__value := plugify.ConstructVariant(value)
	__type_ := C.int32_t(type_)
	__outputId := C.int32_t(outputId)
	plugify.Block {
		Try: func() {
			C.AcceptClientInput(__playerSlot, (*C.String)(unsafe.Pointer(&__inputName)), __activatorHandle, __callerHandle, (*C.Variant)(unsafe.Pointer(&__value)), __type_, __outputId)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__inputName)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// AcceptClientInput 
//  @brief Invokes a named input method on a specified client.
//
//  @param playerSlot: The handle of the target client that will receive the input.
//  @param inputName: The name of the input action to invoke.
//  @param activatorHandle: The index of the player's slot that initiated the sequence of actions.
//  @param callerHandle: The index of the player's slot sending this event. Use -1 to specify
//  @param value: The value associated with the input action.
//  @param type_: The type or classification of the value.
//  @param outputId: An identifier for tracking the output of this operation.
func AcceptClientInput(playerSlot int32, inputName string, activatorHandle int32, callerHandle int32, value any, type_ FieldType, outputId int32) {
	P_AcceptClientInput(playerSlot, inputName, activatorHandle, callerHandle, value, type_, outputId)
}

var P_ConnectClientOutput = func(playerSlot int32, output string, functionName string) {
	__playerSlot := C.int32_t(playerSlot)
	__output := plugify.ConstructString(output)
	__functionName := plugify.ConstructString(functionName)
	plugify.Block {
		Try: func() {
			C.ConnectClientOutput(__playerSlot, (*C.String)(unsafe.Pointer(&__output)), (*C.String)(unsafe.Pointer(&__functionName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__output)
			plugify.DestroyString(&__functionName)
		},
	}.Do()
}

// ConnectClientOutput 
//  @brief Connects a script function to an player output.
//
//  @param playerSlot: The handle of the player.
//  @param output: The name of the output to connect to.
//  @param functionName: The name of the script function to call.
func ConnectClientOutput(playerSlot int32, output string, functionName string) {
	P_ConnectClientOutput(playerSlot, output, functionName)
}

var P_DisconnectClientOutput = func(playerSlot int32, output string, functionName string) {
	__playerSlot := C.int32_t(playerSlot)
	__output := plugify.ConstructString(output)
	__functionName := plugify.ConstructString(functionName)
	plugify.Block {
		Try: func() {
			C.DisconnectClientOutput(__playerSlot, (*C.String)(unsafe.Pointer(&__output)), (*C.String)(unsafe.Pointer(&__functionName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__output)
			plugify.DestroyString(&__functionName)
		},
	}.Do()
}

// DisconnectClientOutput 
//  @brief Disconnects a script function from an player output.
//
//  @param playerSlot: The handle of the player.
//  @param output: The name of the output.
//  @param functionName: The name of the script function to disconnect.
func DisconnectClientOutput(playerSlot int32, output string, functionName string) {
	P_DisconnectClientOutput(playerSlot, output, functionName)
}

var P_DisconnectClientRedirectedOutput = func(playerSlot int32, output string, functionName string, targetHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__output := plugify.ConstructString(output)
	__functionName := plugify.ConstructString(functionName)
	__targetHandle := C.int32_t(targetHandle)
	plugify.Block {
		Try: func() {
			C.DisconnectClientRedirectedOutput(__playerSlot, (*C.String)(unsafe.Pointer(&__output)), (*C.String)(unsafe.Pointer(&__functionName)), __targetHandle)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__output)
			plugify.DestroyString(&__functionName)
		},
	}.Do()
}

// DisconnectClientRedirectedOutput 
//  @brief Disconnects a script function from an I/O event on a different player.
//
//  @param playerSlot: The handle of the calling player.
//  @param output: The name of the output.
//  @param functionName: The function name to disconnect.
//  @param targetHandle: The handle of the entity whose output is being disconnected.
func DisconnectClientRedirectedOutput(playerSlot int32, output string, functionName string, targetHandle int32) {
	P_DisconnectClientRedirectedOutput(playerSlot, output, functionName, targetHandle)
}

var P_FireClientOutput = func(playerSlot int32, outputName string, activatorHandle int32, callerHandle int32, value any, type_ FieldType, delay float32) {
	__playerSlot := C.int32_t(playerSlot)
	__outputName := plugify.ConstructString(outputName)
	__activatorHandle := C.int32_t(activatorHandle)
	__callerHandle := C.int32_t(callerHandle)
	__value := plugify.ConstructVariant(value)
	__type_ := C.int32_t(type_)
	__delay := C.float(delay)
	plugify.Block {
		Try: func() {
			C.FireClientOutput(__playerSlot, (*C.String)(unsafe.Pointer(&__outputName)), __activatorHandle, __callerHandle, (*C.Variant)(unsafe.Pointer(&__value)), __type_, __delay)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__outputName)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// FireClientOutput 
//  @brief Fires an player output.
//
//  @param playerSlot: The handle of the player firing the output.
//  @param outputName: The name of the output to fire.
//  @param activatorHandle: The entity activating the output.
//  @param callerHandle: The entity that called the output.
//  @param value: The value associated with the input action.
//  @param type_: The type or classification of the value.
//  @param delay: Delay in seconds before firing the output.
func FireClientOutput(playerSlot int32, outputName string, activatorHandle int32, callerHandle int32, value any, type_ FieldType, delay float32) {
	P_FireClientOutput(playerSlot, outputName, activatorHandle, callerHandle, value, type_, delay)
}

var P_RedirectClientOutput = func(playerSlot int32, output string, functionName string, targetHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__output := plugify.ConstructString(output)
	__functionName := plugify.ConstructString(functionName)
	__targetHandle := C.int32_t(targetHandle)
	plugify.Block {
		Try: func() {
			C.RedirectClientOutput(__playerSlot, (*C.String)(unsafe.Pointer(&__output)), (*C.String)(unsafe.Pointer(&__functionName)), __targetHandle)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__output)
			plugify.DestroyString(&__functionName)
		},
	}.Do()
}

// RedirectClientOutput 
//  @brief Redirects an player output to call a function on another player.
//
//  @param playerSlot: The handle of the player whose output is being redirected.
//  @param output: The name of the output to redirect.
//  @param functionName: The function name to call on the target player.
//  @param targetHandle: The handle of the entity that will receive the output call.
func RedirectClientOutput(playerSlot int32, output string, functionName string, targetHandle int32) {
	P_RedirectClientOutput(playerSlot, output, functionName, targetHandle)
}

var P_FollowClient = func(playerSlot int32, attachmentHandle int32, boneMerge bool) {
	__playerSlot := C.int32_t(playerSlot)
	__attachmentHandle := C.int32_t(attachmentHandle)
	__boneMerge := C.bool(boneMerge)
	C.FollowClient(__playerSlot, __attachmentHandle, __boneMerge)
}

// FollowClient 
//  @brief Makes an client follow another client with optional bone merging.
//
//  @param playerSlot: The index of the player's slot that will follow
//  @param attachmentHandle: The index of the player's slot to follow
//  @param boneMerge: If true, bones will be merged between entities
func FollowClient(playerSlot int32, attachmentHandle int32, boneMerge bool) {
	P_FollowClient(playerSlot, attachmentHandle, boneMerge)
}

var P_FollowClientMerge = func(playerSlot int32, attachmentHandle int32, boneOrAttachName string) {
	__playerSlot := C.int32_t(playerSlot)
	__attachmentHandle := C.int32_t(attachmentHandle)
	__boneOrAttachName := plugify.ConstructString(boneOrAttachName)
	plugify.Block {
		Try: func() {
			C.FollowClientMerge(__playerSlot, __attachmentHandle, (*C.String)(unsafe.Pointer(&__boneOrAttachName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__boneOrAttachName)
		},
	}.Do()
}

// FollowClientMerge 
//  @brief Makes an client follow another client and merge with a specific bone or attachment.
//
//  @param playerSlot: The index of the player's slot that will follow
//  @param attachmentHandle: The index of the player's slot to follow
//  @param boneOrAttachName: Name of the bone or attachment point to merge with
func FollowClientMerge(playerSlot int32, attachmentHandle int32, boneOrAttachName string) {
	P_FollowClientMerge(playerSlot, attachmentHandle, boneOrAttachName)
}

var P_TakeClientDamage = func(playerSlot int32, inflictorSlot int32, attackerSlot int32, force plugify.Vector3, hitPos plugify.Vector3, damage float32, damageTypes DamageTypes) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__inflictorSlot := C.int32_t(inflictorSlot)
	__attackerSlot := C.int32_t(attackerSlot)
	__force := *(*C.Vector3)(unsafe.Pointer(&force))
	__hitPos := *(*C.Vector3)(unsafe.Pointer(&hitPos))
	__damage := C.float(damage)
	__damageTypes := C.int32_t(damageTypes)
	__retVal = int32(C.TakeClientDamage(__playerSlot, __inflictorSlot, __attackerSlot, &__force, &__hitPos, __damage, __damageTypes))
	return __retVal
}

// TakeClientDamage 
//  @brief Apply damage to an client.
//
//  @param playerSlot: The index of the player's slot receiving damage
//  @param inflictorSlot: The index of the player's slot inflicting damage (e.g., projectile)
//  @param attackerSlot: The index of the attacking client
//  @param force: Direction and magnitude of force to apply
//  @param hitPos: Position where the damage hit occurred
//  @param damage: Amount of damage to apply
//  @param damageTypes: Bitfield of damage type flags
//
//  @return Amount of damage actually applied to the client
func TakeClientDamage(playerSlot int32, inflictorSlot int32, attackerSlot int32, force plugify.Vector3, hitPos plugify.Vector3, damage float32, damageTypes DamageTypes) int32 {
	return P_TakeClientDamage(playerSlot, inflictorSlot, attackerSlot, force, hitPos, damage, damageTypes)
}

var P_GetClientPawn = func(playerSlot int32) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uintptr(C.GetClientPawn(__playerSlot))
	return __retVal
}

// GetClientPawn 
//  @brief Retrieves the pawn entity pointer associated with a client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return A pointer to the client's pawn entity, or nullptr if the client or controller is invalid.
func GetClientPawn(playerSlot int32) uintptr {
	return P_GetClientPawn(playerSlot)
}

var P_ProcessTargetString = func(caller int32, target string) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__caller := C.int32_t(caller)
	__target := plugify.ConstructString(target)
	plugify.Block {
		Try: func() {
			__native := C.ProcessTargetString(__caller, (*C.String)(unsafe.Pointer(&__target)))
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32[int32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
			plugify.DestroyString(&__target)
		},
	}.Do()
	return __retVal
}

// ProcessTargetString 
//  @brief Processes the target string to determine if one user can target another.
//
//  @param caller: The index of the player's slot making the target request.
//  @param target: The target string specifying the player or players to be targeted.
//
//  @return A vector where the result of the targeting operation will be stored.
func ProcessTargetString(caller int32, target string) []int32 {
	return P_ProcessTargetString(caller, target)
}

var P_SwitchClientTeam = func(playerSlot int32, team CSTeam) {
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	C.SwitchClientTeam(__playerSlot, __team)
}

// SwitchClientTeam 
//  @brief Switches the player's team.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The team index to switch the client to.
func SwitchClientTeam(playerSlot int32, team CSTeam) {
	P_SwitchClientTeam(playerSlot, team)
}

var P_ChangeClientTeam = func(playerSlot int32, team CSTeam) {
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	C.ChangeClientTeam(__playerSlot, __team)
}

// ChangeClientTeam 
//  @brief Changes the player's team.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The team index to change the client to.
func ChangeClientTeam(playerSlot int32, team CSTeam) {
	P_ChangeClientTeam(playerSlot, team)
}

var P_RespawnClient = func(playerSlot int32) {
	__playerSlot := C.int32_t(playerSlot)
	C.RespawnClient(__playerSlot)
}

// RespawnClient 
//  @brief Respawns a player.
//
//  @param playerSlot: The index of the player's slot to respawn.
func RespawnClient(playerSlot int32) {
	P_RespawnClient(playerSlot)
}

var P_ForcePlayerSuicide = func(playerSlot int32, explode bool, force bool) {
	__playerSlot := C.int32_t(playerSlot)
	__explode := C.bool(explode)
	__force := C.bool(force)
	C.ForcePlayerSuicide(__playerSlot, __explode, __force)
}

// ForcePlayerSuicide 
//  @brief Forces a player to commit suicide.
//
//  @param playerSlot: The index of the player's slot.
//  @param explode: If true, the client will explode upon death.
//  @param force: If true, the suicide will be forced.
func ForcePlayerSuicide(playerSlot int32, explode bool, force bool) {
	P_ForcePlayerSuicide(playerSlot, explode, force)
}

var P_KickClient = func(playerSlot int32, reason NetworkDisconnectionReason, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__reason := C.int32_t(reason)
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			C.KickClient(__playerSlot, __reason, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// KickClient 
//  @brief Disconnects a client from the server as soon as the next frame starts.
//
//  @param playerSlot: The index of the player's slot to be kicked.
//  @param reason: The network-level reason code describing why the client is being disconnected.
//  @param message: The optional internal diagnostic message. If empty, no message is passed to the engine.
func KickClient(playerSlot int32, reason NetworkDisconnectionReason, message string) {
	P_KickClient(playerSlot, reason, message)
}

var P_BanClient = func(playerSlot int32, duration float32, kick bool) {
	__playerSlot := C.int32_t(playerSlot)
	__duration := C.float(duration)
	__kick := C.bool(kick)
	C.BanClient(__playerSlot, __duration, __kick)
}

// BanClient 
//  @brief Bans a client for a specified duration.
//
//  @param playerSlot: The index of the player's slot to be banned.
//  @param duration: Duration of the ban in seconds.
//  @param kick: If true, the client will be kicked immediately after being banned.
func BanClient(playerSlot int32, duration float32, kick bool) {
	P_BanClient(playerSlot, duration, kick)
}

var P_BanIdentity = func(steamId uint64, duration float32, kick bool) {
	__steamId := C.uint64_t(steamId)
	__duration := C.float(duration)
	__kick := C.bool(kick)
	C.BanIdentity(__steamId, __duration, __kick)
}

// BanIdentity 
//  @brief Bans an identity (either an IP address or a Steam authentication string).
//
//  @param steamId: The Steam ID to ban.
//  @param duration: Duration of the ban in seconds.
//  @param kick: If true, the client will be kicked immediately after being banned.
func BanIdentity(steamId uint64, duration float32, kick bool) {
	P_BanIdentity(steamId, duration, kick)
}

var P_GetClientActiveWeapon = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientActiveWeapon(__playerSlot))
	return __retVal
}

// GetClientActiveWeapon 
//  @brief Retrieves the handle of the client's currently active weapon.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The entity handle of the active weapon, or INVALID_EHANDLE_INDEX if the client is invalid or has no active weapon.
func GetClientActiveWeapon(playerSlot int32) int32 {
	return P_GetClientActiveWeapon(playerSlot)
}

var P_GetClientWeapons = func(playerSlot int32) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientWeapons(__playerSlot)
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

// GetClientWeapons 
//  @brief Retrieves a list of weapon handles owned by the client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return A vector of entity handles for the client's weapons, or an empty vector if the client is invalid or has no weapons.
func GetClientWeapons(playerSlot int32) []int32 {
	return P_GetClientWeapons(playerSlot)
}

var P_RemoveWeapons = func(playerSlot int32, removeSuit bool) {
	__playerSlot := C.int32_t(playerSlot)
	__removeSuit := C.bool(removeSuit)
	C.RemoveWeapons(__playerSlot, __removeSuit)
}

// RemoveWeapons 
//  @brief Removes all weapons from a client, with an option to remove the suit as well.
//
//  @param playerSlot: The index of the player's slot.
//  @param removeSuit: A boolean indicating whether to also remove the client's suit.
func RemoveWeapons(playerSlot int32, removeSuit bool) {
	P_RemoveWeapons(playerSlot, removeSuit)
}

var P_DropWeapon = func(playerSlot int32, weaponHandle int32, target plugify.Vector3, velocity plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	__target := *(*C.Vector3)(unsafe.Pointer(&target))
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.DropWeapon(__playerSlot, __weaponHandle, &__target, &__velocity)
}

// DropWeapon 
//  @brief Forces a player to drop their weapon.
//
//  @param playerSlot: The index of the player's slot.
//  @param weaponHandle: The handle of weapon to drop.
//  @param target: Target direction.
//  @param velocity: Velocity to toss weapon or zero to just drop weapon.
func DropWeapon(playerSlot int32, weaponHandle int32, target plugify.Vector3, velocity plugify.Vector3) {
	P_DropWeapon(playerSlot, weaponHandle, target, velocity)
}

var P_SelectWeapon = func(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.SelectWeapon(__playerSlot, __weaponHandle)
}

// SelectWeapon 
//  @brief Selects a player's weapon.
//
//  @param playerSlot: The index of the player's slot.
//  @param weaponHandle: The handle of weapon to bump.
func SelectWeapon(playerSlot int32, weaponHandle int32) {
	P_SelectWeapon(playerSlot, weaponHandle)
}

var P_SwitchWeapon = func(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.SwitchWeapon(__playerSlot, __weaponHandle)
}

// SwitchWeapon 
//  @brief Switches a player's weapon.
//
//  @param playerSlot: The index of the player's slot.
//  @param weaponHandle: The handle of weapon to switch.
func SwitchWeapon(playerSlot int32, weaponHandle int32) {
	P_SwitchWeapon(playerSlot, weaponHandle)
}

var P_RemoveWeapon = func(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.RemoveWeapon(__playerSlot, __weaponHandle)
}

// RemoveWeapon 
//  @brief Removes a player's weapon.
//
//  @param playerSlot: The index of the player's slot.
//  @param weaponHandle: The handle of weapon to remove.
func RemoveWeapon(playerSlot int32, weaponHandle int32) {
	P_RemoveWeapon(playerSlot, weaponHandle)
}

var P_GiveNamedItem = func(playerSlot int32, itemName string) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__itemName := plugify.ConstructString(itemName)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GiveNamedItem(__playerSlot, (*C.String)(unsafe.Pointer(&__itemName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__itemName)
		},
	}.Do()
	return __retVal
}

// GiveNamedItem 
//  @brief Gives a named item (e.g., weapon) to a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemName: The name of the item to give.
//
//  @return The entity handle of the created item, or INVALID_EHANDLE_INDEX if the client or item is invalid.
func GiveNamedItem(playerSlot int32, itemName string) int32 {
	return P_GiveNamedItem(playerSlot, itemName)
}

var P_GetClientButtons = func(playerSlot int32, buttonIndex int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__buttonIndex := C.int32_t(buttonIndex)
	__retVal = uint64(C.GetClientButtons(__playerSlot, __buttonIndex))
	return __retVal
}

// GetClientButtons 
//  @brief Retrieves the state of a specific button for a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param buttonIndex: The index of the button (0-2).
//
//  @return uint64_t The state of the specified button, or 0 if the client or button index is invalid.
func GetClientButtons(playerSlot int32, buttonIndex int32) uint64 {
	return P_GetClientButtons(playerSlot, buttonIndex)
}

var P_GetClientArmor = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientArmor(__playerSlot))
	return __retVal
}

// GetClientArmor 
//  @brief Returns the client's armor value.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The armor value of the client.
func GetClientArmor(playerSlot int32) int32 {
	return P_GetClientArmor(playerSlot)
}

var P_SetClientArmor = func(playerSlot int32, armor int32) {
	__playerSlot := C.int32_t(playerSlot)
	__armor := C.int32_t(armor)
	C.SetClientArmor(__playerSlot, __armor)
}

// SetClientArmor 
//  @brief Sets the client's armor value.
//
//  @param playerSlot: The index of the player's slot.
//  @param armor: The armor value to set.
func SetClientArmor(playerSlot int32, armor int32) {
	P_SetClientArmor(playerSlot, armor)
}

var P_GetClientSpeed = func(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientSpeed(__playerSlot))
	return __retVal
}

// GetClientSpeed 
//  @brief Returns the client's speed value.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The speed value of the client.
func GetClientSpeed(playerSlot int32) float32 {
	return P_GetClientSpeed(playerSlot)
}

var P_SetClientSpeed = func(playerSlot int32, speed float32) {
	__playerSlot := C.int32_t(playerSlot)
	__speed := C.float(speed)
	C.SetClientSpeed(__playerSlot, __speed)
}

// SetClientSpeed 
//  @brief Sets the client's speed value.
//
//  @param playerSlot: The index of the player's slot.
//  @param speed: The speed value to set.
func SetClientSpeed(playerSlot int32, speed float32) {
	P_SetClientSpeed(playerSlot, speed)
}

var P_GetClientMoney = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMoney(__playerSlot))
	return __retVal
}

// GetClientMoney 
//  @brief Retrieves the amount of money a client has.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The amount of money the client has, or 0 if the player slot is invalid.
func GetClientMoney(playerSlot int32) int32 {
	return P_GetClientMoney(playerSlot)
}

var P_SetClientMoney = func(playerSlot int32, money int32) {
	__playerSlot := C.int32_t(playerSlot)
	__money := C.int32_t(money)
	C.SetClientMoney(__playerSlot, __money)
}

// SetClientMoney 
//  @brief Sets the amount of money for a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param money: The amount of money to set.
func SetClientMoney(playerSlot int32, money int32) {
	P_SetClientMoney(playerSlot, money)
}

var P_GetClientKills = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientKills(__playerSlot))
	return __retVal
}

// GetClientKills 
//  @brief Retrieves the number of kills for a client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The number of kills the client has, or 0 if the player slot is invalid.
func GetClientKills(playerSlot int32) int32 {
	return P_GetClientKills(playerSlot)
}

var P_SetClientKills = func(playerSlot int32, kills int32) {
	__playerSlot := C.int32_t(playerSlot)
	__kills := C.int32_t(kills)
	C.SetClientKills(__playerSlot, __kills)
}

// SetClientKills 
//  @brief Sets the number of kills for a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param kills: The number of kills to set.
func SetClientKills(playerSlot int32, kills int32) {
	P_SetClientKills(playerSlot, kills)
}

var P_GetClientDeaths = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientDeaths(__playerSlot))
	return __retVal
}

// GetClientDeaths 
//  @brief Retrieves the number of deaths for a client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The number of deaths the client has, or 0 if the player slot is invalid.
func GetClientDeaths(playerSlot int32) int32 {
	return P_GetClientDeaths(playerSlot)
}

var P_SetClientDeaths = func(playerSlot int32, deaths int32) {
	__playerSlot := C.int32_t(playerSlot)
	__deaths := C.int32_t(deaths)
	C.SetClientDeaths(__playerSlot, __deaths)
}

// SetClientDeaths 
//  @brief Sets the number of deaths for a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param deaths: The number of deaths to set.
func SetClientDeaths(playerSlot int32, deaths int32) {
	P_SetClientDeaths(playerSlot, deaths)
}

var P_GetClientAssists = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientAssists(__playerSlot))
	return __retVal
}

// GetClientAssists 
//  @brief Retrieves the number of assists for a client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The number of assists the client has, or 0 if the player slot is invalid.
func GetClientAssists(playerSlot int32) int32 {
	return P_GetClientAssists(playerSlot)
}

var P_SetClientAssists = func(playerSlot int32, assists int32) {
	__playerSlot := C.int32_t(playerSlot)
	__assists := C.int32_t(assists)
	C.SetClientAssists(__playerSlot, __assists)
}

// SetClientAssists 
//  @brief Sets the number of assists for a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param assists: The number of assists to set.
func SetClientAssists(playerSlot int32, assists int32) {
	P_SetClientAssists(playerSlot, assists)
}

var P_GetClientDamage = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientDamage(__playerSlot))
	return __retVal
}

// GetClientDamage 
//  @brief Retrieves the total damage dealt by a client.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The total damage dealt by the client, or 0 if the player slot is invalid.
func GetClientDamage(playerSlot int32) int32 {
	return P_GetClientDamage(playerSlot)
}

var P_SetClientDamage = func(playerSlot int32, damage int32) {
	__playerSlot := C.int32_t(playerSlot)
	__damage := C.int32_t(damage)
	C.SetClientDamage(__playerSlot, __damage)
}

// SetClientDamage 
//  @brief Sets the total damage dealt by a client.
//
//  @param playerSlot: The index of the player's slot.
//  @param damage: The amount of damage to set.
func SetClientDamage(playerSlot int32, damage int32) {
	P_SetClientDamage(playerSlot, damage)
}

