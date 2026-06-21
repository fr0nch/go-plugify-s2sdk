package s2sdk

/*
#include "listeners.h"
#cgo noescape OnClientConnect_Register
#cgo noescape OnClientConnect_Unregister
#cgo noescape OnClientConnect_Post_Register
#cgo noescape OnClientConnect_Post_Unregister
#cgo noescape OnClientConnected_Register
#cgo noescape OnClientConnected_Unregister
#cgo noescape OnClientPutInServer_Register
#cgo noescape OnClientPutInServer_Unregister
#cgo noescape OnClientDisconnect_Register
#cgo noescape OnClientDisconnect_Unregister
#cgo noescape OnClientDisconnect_Post_Register
#cgo noescape OnClientDisconnect_Post_Unregister
#cgo noescape OnClientActive_Register
#cgo noescape OnClientActive_Unregister
#cgo noescape OnClientFullyConnect_Register
#cgo noescape OnClientFullyConnect_Unregister
#cgo noescape OnClientSettingsChanged_Register
#cgo noescape OnClientSettingsChanged_Unregister
#cgo noescape OnClientAuthenticated_Register
#cgo noescape OnClientAuthenticated_Unregister
#cgo noescape OnRoundTerminated_Register
#cgo noescape OnRoundTerminated_Unregister
#cgo noescape OnEntityCreated_Register
#cgo noescape OnEntityCreated_Unregister
#cgo noescape OnEntitySpawned_Register
#cgo noescape OnEntitySpawned_Unregister
#cgo noescape OnEntityDeleted_Register
#cgo noescape OnEntityDeleted_Unregister
#cgo noescape OnEntityParentChanged_Register
#cgo noescape OnEntityParentChanged_Unregister
#cgo noescape OnServerCheckTransmit_Register
#cgo noescape OnServerCheckTransmit_Unregister
#cgo noescape OnServerStartup_Register
#cgo noescape OnServerStartup_Unregister
#cgo noescape OnBuildGameSessionManifest_Register
#cgo noescape OnBuildGameSessionManifest_Unregister
#cgo noescape OnServerActivate_Register
#cgo noescape OnServerActivate_Unregister
#cgo noescape OnServerSpawn_Register
#cgo noescape OnServerSpawn_Unregister
#cgo noescape OnServerStarted_Register
#cgo noescape OnServerStarted_Unregister
#cgo noescape OnMapStart_Register
#cgo noescape OnMapStart_Unregister
#cgo noescape OnMapEnd_Register
#cgo noescape OnMapEnd_Unregister
#cgo noescape OnGameFrame_Register
#cgo noescape OnGameFrame_Unregister
#cgo noescape OnUpdateWhenNotInGame_Register
#cgo noescape OnUpdateWhenNotInGame_Unregister
#cgo noescape OnPreWorldUpdate_Register
#cgo noescape OnPreWorldUpdate_Unregister
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

// Generated from s2sdk (group: listeners)

var P_OnClientConnect_Register = func(callback OnClientConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Register(__callback)
}

// OnClientConnect_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientConnect_Register(callback OnClientConnectCallback) {
	P_OnClientConnect_Register(callback)
}

var P_OnClientConnect_Unregister = func(callback OnClientConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Unregister(__callback)
}

// OnClientConnect_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientConnect_Unregister(callback OnClientConnectCallback) {
	P_OnClientConnect_Unregister(callback)
}

var P_OnClientConnect_Post_Register = func(callback OnClientConnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Post_Register(__callback)
}

// OnClientConnect_Post_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientConnect_Post_Register(callback OnClientConnect_PostCallback) {
	P_OnClientConnect_Post_Register(callback)
}

var P_OnClientConnect_Post_Unregister = func(callback OnClientConnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Post_Unregister(__callback)
}

// OnClientConnect_Post_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientConnect_Post_Unregister(callback OnClientConnect_PostCallback) {
	P_OnClientConnect_Post_Unregister(callback)
}

var P_OnClientConnected_Register = func(callback OnClientConnectedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnected_Register(__callback)
}

// OnClientConnected_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientConnected_Register(callback OnClientConnectedCallback) {
	P_OnClientConnected_Register(callback)
}

var P_OnClientConnected_Unregister = func(callback OnClientConnectedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnected_Unregister(__callback)
}

// OnClientConnected_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientConnected_Unregister(callback OnClientConnectedCallback) {
	P_OnClientConnected_Unregister(callback)
}

var P_OnClientPutInServer_Register = func(callback OnClientPutInServerCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientPutInServer_Register(__callback)
}

// OnClientPutInServer_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientPutInServer_Register(callback OnClientPutInServerCallback) {
	P_OnClientPutInServer_Register(callback)
}

var P_OnClientPutInServer_Unregister = func(callback OnClientPutInServerCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientPutInServer_Unregister(__callback)
}

// OnClientPutInServer_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientPutInServer_Unregister(callback OnClientPutInServerCallback) {
	P_OnClientPutInServer_Unregister(callback)
}

var P_OnClientDisconnect_Register = func(callback OnClientDisconnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Register(__callback)
}

// OnClientDisconnect_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientDisconnect_Register(callback OnClientDisconnectCallback) {
	P_OnClientDisconnect_Register(callback)
}

var P_OnClientDisconnect_Unregister = func(callback OnClientDisconnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Unregister(__callback)
}

// OnClientDisconnect_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientDisconnect_Unregister(callback OnClientDisconnectCallback) {
	P_OnClientDisconnect_Unregister(callback)
}

var P_OnClientDisconnect_Post_Register = func(callback OnClientDisconnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Post_Register(__callback)
}

// OnClientDisconnect_Post_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientDisconnect_Post_Register(callback OnClientDisconnect_PostCallback) {
	P_OnClientDisconnect_Post_Register(callback)
}

var P_OnClientDisconnect_Post_Unregister = func(callback OnClientDisconnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Post_Unregister(__callback)
}

// OnClientDisconnect_Post_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientDisconnect_Post_Unregister(callback OnClientDisconnect_PostCallback) {
	P_OnClientDisconnect_Post_Unregister(callback)
}

var P_OnClientActive_Register = func(callback OnClientActiveCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientActive_Register(__callback)
}

// OnClientActive_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientActive_Register(callback OnClientActiveCallback) {
	P_OnClientActive_Register(callback)
}

var P_OnClientActive_Unregister = func(callback OnClientActiveCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientActive_Unregister(__callback)
}

// OnClientActive_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientActive_Unregister(callback OnClientActiveCallback) {
	P_OnClientActive_Unregister(callback)
}

var P_OnClientFullyConnect_Register = func(callback OnClientFullyConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientFullyConnect_Register(__callback)
}

// OnClientFullyConnect_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientFullyConnect_Register(callback OnClientFullyConnectCallback) {
	P_OnClientFullyConnect_Register(callback)
}

var P_OnClientFullyConnect_Unregister = func(callback OnClientFullyConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientFullyConnect_Unregister(__callback)
}

// OnClientFullyConnect_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientFullyConnect_Unregister(callback OnClientFullyConnectCallback) {
	P_OnClientFullyConnect_Unregister(callback)
}

var P_OnClientSettingsChanged_Register = func(callback OnClientSettingsChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientSettingsChanged_Register(__callback)
}

// OnClientSettingsChanged_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientSettingsChanged_Register(callback OnClientSettingsChangedCallback) {
	P_OnClientSettingsChanged_Register(callback)
}

var P_OnClientSettingsChanged_Unregister = func(callback OnClientSettingsChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientSettingsChanged_Unregister(__callback)
}

// OnClientSettingsChanged_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientSettingsChanged_Unregister(callback OnClientSettingsChangedCallback) {
	P_OnClientSettingsChanged_Unregister(callback)
}

var P_OnClientAuthenticated_Register = func(callback OnClientAuthenticatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientAuthenticated_Register(__callback)
}

// OnClientAuthenticated_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnClientAuthenticated_Register(callback OnClientAuthenticatedCallback) {
	P_OnClientAuthenticated_Register(callback)
}

var P_OnClientAuthenticated_Unregister = func(callback OnClientAuthenticatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientAuthenticated_Unregister(__callback)
}

// OnClientAuthenticated_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnClientAuthenticated_Unregister(callback OnClientAuthenticatedCallback) {
	P_OnClientAuthenticated_Unregister(callback)
}

var P_OnRoundTerminated_Register = func(callback OnRoundTerminatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnRoundTerminated_Register(__callback)
}

// OnRoundTerminated_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnRoundTerminated_Register(callback OnRoundTerminatedCallback) {
	P_OnRoundTerminated_Register(callback)
}

var P_OnRoundTerminated_Unregister = func(callback OnRoundTerminatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnRoundTerminated_Unregister(__callback)
}

// OnRoundTerminated_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnRoundTerminated_Unregister(callback OnRoundTerminatedCallback) {
	P_OnRoundTerminated_Unregister(callback)
}

var P_OnEntityCreated_Register = func(callback OnEntityCreatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityCreated_Register(__callback)
}

// OnEntityCreated_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnEntityCreated_Register(callback OnEntityCreatedCallback) {
	P_OnEntityCreated_Register(callback)
}

var P_OnEntityCreated_Unregister = func(callback OnEntityCreatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityCreated_Unregister(__callback)
}

// OnEntityCreated_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnEntityCreated_Unregister(callback OnEntityCreatedCallback) {
	P_OnEntityCreated_Unregister(callback)
}

var P_OnEntitySpawned_Register = func(callback OnEntitySpawnedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntitySpawned_Register(__callback)
}

// OnEntitySpawned_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnEntitySpawned_Register(callback OnEntitySpawnedCallback) {
	P_OnEntitySpawned_Register(callback)
}

var P_OnEntitySpawned_Unregister = func(callback OnEntitySpawnedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntitySpawned_Unregister(__callback)
}

// OnEntitySpawned_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnEntitySpawned_Unregister(callback OnEntitySpawnedCallback) {
	P_OnEntitySpawned_Unregister(callback)
}

var P_OnEntityDeleted_Register = func(callback OnEntityDeletedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityDeleted_Register(__callback)
}

// OnEntityDeleted_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnEntityDeleted_Register(callback OnEntityDeletedCallback) {
	P_OnEntityDeleted_Register(callback)
}

var P_OnEntityDeleted_Unregister = func(callback OnEntityDeletedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityDeleted_Unregister(__callback)
}

// OnEntityDeleted_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnEntityDeleted_Unregister(callback OnEntityDeletedCallback) {
	P_OnEntityDeleted_Unregister(callback)
}

var P_OnEntityParentChanged_Register = func(callback OnEntityParentChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityParentChanged_Register(__callback)
}

// OnEntityParentChanged_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnEntityParentChanged_Register(callback OnEntityParentChangedCallback) {
	P_OnEntityParentChanged_Register(callback)
}

var P_OnEntityParentChanged_Unregister = func(callback OnEntityParentChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityParentChanged_Unregister(__callback)
}

// OnEntityParentChanged_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnEntityParentChanged_Unregister(callback OnEntityParentChangedCallback) {
	P_OnEntityParentChanged_Unregister(callback)
}

var P_OnServerCheckTransmit_Register = func(callback OnServerCheckTransmitCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerCheckTransmit_Register(__callback)
}

// OnServerCheckTransmit_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnServerCheckTransmit_Register(callback OnServerCheckTransmitCallback) {
	P_OnServerCheckTransmit_Register(callback)
}

var P_OnServerCheckTransmit_Unregister = func(callback OnServerCheckTransmitCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerCheckTransmit_Unregister(__callback)
}

// OnServerCheckTransmit_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnServerCheckTransmit_Unregister(callback OnServerCheckTransmitCallback) {
	P_OnServerCheckTransmit_Unregister(callback)
}

var P_OnServerStartup_Register = func(callback OnServerStartupCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStartup_Register(__callback)
}

// OnServerStartup_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnServerStartup_Register(callback OnServerStartupCallback) {
	P_OnServerStartup_Register(callback)
}

var P_OnServerStartup_Unregister = func(callback OnServerStartupCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStartup_Unregister(__callback)
}

// OnServerStartup_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnServerStartup_Unregister(callback OnServerStartupCallback) {
	P_OnServerStartup_Unregister(callback)
}

var P_OnBuildGameSessionManifest_Register = func(callback OnBuildGameSessionManifestCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnBuildGameSessionManifest_Register(__callback)
}

// OnBuildGameSessionManifest_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnBuildGameSessionManifest_Register(callback OnBuildGameSessionManifestCallback) {
	P_OnBuildGameSessionManifest_Register(callback)
}

var P_OnBuildGameSessionManifest_Unregister = func(callback OnBuildGameSessionManifestCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnBuildGameSessionManifest_Unregister(__callback)
}

// OnBuildGameSessionManifest_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnBuildGameSessionManifest_Unregister(callback OnBuildGameSessionManifestCallback) {
	P_OnBuildGameSessionManifest_Unregister(callback)
}

var P_OnServerActivate_Register = func(callback OnServerActivateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerActivate_Register(__callback)
}

// OnServerActivate_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnServerActivate_Register(callback OnServerActivateCallback) {
	P_OnServerActivate_Register(callback)
}

var P_OnServerActivate_Unregister = func(callback OnServerActivateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerActivate_Unregister(__callback)
}

// OnServerActivate_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnServerActivate_Unregister(callback OnServerActivateCallback) {
	P_OnServerActivate_Unregister(callback)
}

var P_OnServerSpawn_Register = func(callback OnServerSpawnCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerSpawn_Register(__callback)
}

// OnServerSpawn_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnServerSpawn_Register(callback OnServerSpawnCallback) {
	P_OnServerSpawn_Register(callback)
}

var P_OnServerSpawn_Unregister = func(callback OnServerSpawnCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerSpawn_Unregister(__callback)
}

// OnServerSpawn_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnServerSpawn_Unregister(callback OnServerSpawnCallback) {
	P_OnServerSpawn_Unregister(callback)
}

var P_OnServerStarted_Register = func(callback OnServerStartedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStarted_Register(__callback)
}

// OnServerStarted_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnServerStarted_Register(callback OnServerStartedCallback) {
	P_OnServerStarted_Register(callback)
}

var P_OnServerStarted_Unregister = func(callback OnServerStartedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStarted_Unregister(__callback)
}

// OnServerStarted_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnServerStarted_Unregister(callback OnServerStartedCallback) {
	P_OnServerStarted_Unregister(callback)
}

var P_OnMapStart_Register = func(callback OnMapStartCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnMapStart_Register(__callback)
}

// OnMapStart_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnMapStart_Register(callback OnMapStartCallback) {
	P_OnMapStart_Register(callback)
}

var P_OnMapStart_Unregister = func(callback OnMapStartCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnMapStart_Unregister(__callback)
}

// OnMapStart_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnMapStart_Unregister(callback OnMapStartCallback) {
	P_OnMapStart_Unregister(callback)
}

var P_OnMapEnd_Register = func(callback OnMapEndCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnMapEnd_Register(__callback)
}

// OnMapEnd_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnMapEnd_Register(callback OnMapEndCallback) {
	P_OnMapEnd_Register(callback)
}

var P_OnMapEnd_Unregister = func(callback OnMapEndCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnMapEnd_Unregister(__callback)
}

// OnMapEnd_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnMapEnd_Unregister(callback OnMapEndCallback) {
	P_OnMapEnd_Unregister(callback)
}

var P_OnGameFrame_Register = func(callback OnGameFrameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnGameFrame_Register(__callback)
}

// OnGameFrame_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnGameFrame_Register(callback OnGameFrameCallback) {
	P_OnGameFrame_Register(callback)
}

var P_OnGameFrame_Unregister = func(callback OnGameFrameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnGameFrame_Unregister(__callback)
}

// OnGameFrame_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnGameFrame_Unregister(callback OnGameFrameCallback) {
	P_OnGameFrame_Unregister(callback)
}

var P_OnUpdateWhenNotInGame_Register = func(callback OnUpdateWhenNotInGameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnUpdateWhenNotInGame_Register(__callback)
}

// OnUpdateWhenNotInGame_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnUpdateWhenNotInGame_Register(callback OnUpdateWhenNotInGameCallback) {
	P_OnUpdateWhenNotInGame_Register(callback)
}

var P_OnUpdateWhenNotInGame_Unregister = func(callback OnUpdateWhenNotInGameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnUpdateWhenNotInGame_Unregister(__callback)
}

// OnUpdateWhenNotInGame_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnUpdateWhenNotInGame_Unregister(callback OnUpdateWhenNotInGameCallback) {
	P_OnUpdateWhenNotInGame_Unregister(callback)
}

var P_OnPreWorldUpdate_Register = func(callback OnPreWorldUpdateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnPreWorldUpdate_Register(__callback)
}

// OnPreWorldUpdate_Register 
//  @brief Register callback to event.
//
//  @param callback: Function callback.
func OnPreWorldUpdate_Register(callback OnPreWorldUpdateCallback) {
	P_OnPreWorldUpdate_Register(callback)
}

var P_OnPreWorldUpdate_Unregister = func(callback OnPreWorldUpdateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnPreWorldUpdate_Unregister(__callback)
}

// OnPreWorldUpdate_Unregister 
//  @brief Unregister callback to event.
//
//  @param callback: Function callback.
func OnPreWorldUpdate_Unregister(callback OnPreWorldUpdateCallback) {
	P_OnPreWorldUpdate_Unregister(callback)
}

