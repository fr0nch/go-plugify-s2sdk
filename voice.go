package s2sdk

/*
#include "voice.h"
#cgo noescape HookClientVoiceData
#cgo noescape UnhookClientVoiceData
#cgo noescape VoiceDataGetFormat
#cgo noescape VoiceDataSetFormat
#cgo noescape VoiceDataGetAudio
#cgo noescape VoiceDataSetAudio
#cgo noescape VoiceDataGetSampleRate
#cgo noescape VoiceDataSetSampleRate
#cgo noescape VoiceDataGetNumPackets
#cgo noescape VoiceDataSetNumPackets
#cgo noescape VoiceDataGetPacketOffsets
#cgo noescape VoiceDataSetPacketOffsets
#cgo noescape VoiceDataGetVoiceLevel
#cgo noescape VoiceDataSetVoiceLevel
#cgo noescape SendVoiceData
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

// Generated from s2sdk (group: voice)

var _HookClientVoiceData = func(callback ClientVoiceDataHookCallback) bool {
	var __retVal bool
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = bool(C.HookClientVoiceData(__callback))
	return __retVal
}

// HookClientVoiceData 
//  @brief Hooks every voice packet a client sends, before the engine passes it on. The packet can be changed or blocked.
//
//  @param callback: The callback function to invoke.
//
//  @return True if the hook was successfully added, false otherwise.
func HookClientVoiceData(callback ClientVoiceDataHookCallback) bool {
	return _HookClientVoiceData(callback)
}

var _UnhookClientVoiceData = func(callback ClientVoiceDataHookCallback) bool {
	var __retVal bool
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = bool(C.UnhookClientVoiceData(__callback))
	return __retVal
}

// UnhookClientVoiceData 
//  @brief Unhooks a previously hooked callback from HookClientVoiceData.
//
//  @param callback: The callback function to remove.
//
//  @return True if the hook was successfully removed, false otherwise.
func UnhookClientVoiceData(callback ClientVoiceDataHookCallback) bool {
	return _UnhookClientVoiceData(callback)
}

var _VoiceDataGetFormat = func(voiceData uintptr) VoiceDataFormat {
	var __retVal VoiceDataFormat
	__voiceData := C.uintptr_t(voiceData)
	__retVal = VoiceDataFormat(C.VoiceDataGetFormat(__voiceData))
	return __retVal
}

// VoiceDataGetFormat 
//  @brief Gets the codec of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The voice data format.
func VoiceDataGetFormat(voiceData uintptr) VoiceDataFormat {
	return _VoiceDataGetFormat(voiceData)
}

var _VoiceDataSetFormat = func(voiceData uintptr, format VoiceDataFormat) {
	__voiceData := C.uintptr_t(voiceData)
	__format := C.int32_t(format)
	C.VoiceDataSetFormat(__voiceData, __format)
}

// VoiceDataSetFormat 
//  @brief Sets the codec of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param format: The voice data format.
func VoiceDataSetFormat(voiceData uintptr, format VoiceDataFormat) {
	_VoiceDataSetFormat(voiceData, format)
}

var _VoiceDataGetAudio = func(voiceData uintptr) []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	__voiceData := C.uintptr_t(voiceData)
	plugify.Block {
		Try: func() {
			__native := C.VoiceDataGetAudio(__voiceData)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt8[uint8](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// VoiceDataGetAudio 
//  @brief Gets the compressed audio of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The compressed audio bytes.
func VoiceDataGetAudio(voiceData uintptr) []uint8 {
	return _VoiceDataGetAudio(voiceData)
}

var _VoiceDataSetAudio = func(voiceData uintptr, audio []uint8) {
	__voiceData := C.uintptr_t(voiceData)
	__audio := plugify.ConstructVectorUInt8(audio)
	plugify.Block {
		Try: func() {
			C.VoiceDataSetAudio(__voiceData, (*C.Vector)(unsafe.Pointer(&__audio)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__audio)
		},
	}.Do()
}

// VoiceDataSetAudio 
//  @brief Replaces the compressed audio of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param audio: The compressed audio bytes.
func VoiceDataSetAudio(voiceData uintptr, audio []uint8) {
	_VoiceDataSetAudio(voiceData, audio)
}

var _VoiceDataGetSampleRate = func(voiceData uintptr) uint32 {
	var __retVal uint32
	__voiceData := C.uintptr_t(voiceData)
	__retVal = uint32(C.VoiceDataGetSampleRate(__voiceData))
	return __retVal
}

// VoiceDataGetSampleRate 
//  @brief Gets the sample rate of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The sample rate in Hz.
func VoiceDataGetSampleRate(voiceData uintptr) uint32 {
	return _VoiceDataGetSampleRate(voiceData)
}

var _VoiceDataSetSampleRate = func(voiceData uintptr, sampleRate uint32) {
	__voiceData := C.uintptr_t(voiceData)
	__sampleRate := C.uint32_t(sampleRate)
	C.VoiceDataSetSampleRate(__voiceData, __sampleRate)
}

// VoiceDataSetSampleRate 
//  @brief Sets the sample rate of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param sampleRate: The sample rate in Hz.
func VoiceDataSetSampleRate(voiceData uintptr, sampleRate uint32) {
	_VoiceDataSetSampleRate(voiceData, sampleRate)
}

var _VoiceDataGetNumPackets = func(voiceData uintptr) uint32 {
	var __retVal uint32
	__voiceData := C.uintptr_t(voiceData)
	__retVal = uint32(C.VoiceDataGetNumPackets(__voiceData))
	return __retVal
}

// VoiceDataGetNumPackets 
//  @brief Gets the number of codec packets in a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The number of codec packets.
func VoiceDataGetNumPackets(voiceData uintptr) uint32 {
	return _VoiceDataGetNumPackets(voiceData)
}

var _VoiceDataSetNumPackets = func(voiceData uintptr, numPackets uint32) {
	__voiceData := C.uintptr_t(voiceData)
	__numPackets := C.uint32_t(numPackets)
	C.VoiceDataSetNumPackets(__voiceData, __numPackets)
}

// VoiceDataSetNumPackets 
//  @brief Sets the number of codec packets in a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param numPackets: The number of codec packets.
func VoiceDataSetNumPackets(voiceData uintptr, numPackets uint32) {
	_VoiceDataSetNumPackets(voiceData, numPackets)
}

var _VoiceDataGetPacketOffsets = func(voiceData uintptr) []uint32 {
	var __retVal []uint32
	var __retVal_native plugify.PlgVector
	__voiceData := C.uintptr_t(voiceData)
	plugify.Block {
		Try: func() {
			__native := C.VoiceDataGetPacketOffsets(__voiceData)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt32[uint32](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// VoiceDataGetPacketOffsets 
//  @brief Gets the offsets of the codec packets inside the audio of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The packet offsets in bytes.
func VoiceDataGetPacketOffsets(voiceData uintptr) []uint32 {
	return _VoiceDataGetPacketOffsets(voiceData)
}

var _VoiceDataSetPacketOffsets = func(voiceData uintptr, offsets []uint32) {
	__voiceData := C.uintptr_t(voiceData)
	__offsets := plugify.ConstructVectorUInt32(offsets)
	plugify.Block {
		Try: func() {
			C.VoiceDataSetPacketOffsets(__voiceData, (*C.Vector)(unsafe.Pointer(&__offsets)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt32(&__offsets)
		},
	}.Do()
}

// VoiceDataSetPacketOffsets 
//  @brief Replaces the offsets of the codec packets inside the audio of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param offsets: The packet offsets in bytes.
func VoiceDataSetPacketOffsets(voiceData uintptr, offsets []uint32) {
	_VoiceDataSetPacketOffsets(voiceData, offsets)
}

var _VoiceDataGetVoiceLevel = func(voiceData uintptr) float32 {
	var __retVal float32
	__voiceData := C.uintptr_t(voiceData)
	__retVal = float32(C.VoiceDataGetVoiceLevel(__voiceData))
	return __retVal
}

// VoiceDataGetVoiceLevel 
//  @brief Gets the voice level the client reported for a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//
//  @return The voice level.
func VoiceDataGetVoiceLevel(voiceData uintptr) float32 {
	return _VoiceDataGetVoiceLevel(voiceData)
}

var _VoiceDataSetVoiceLevel = func(voiceData uintptr, voiceLevel float32) {
	__voiceData := C.uintptr_t(voiceData)
	__voiceLevel := C.float(voiceLevel)
	C.VoiceDataSetVoiceLevel(__voiceData, __voiceLevel)
}

// VoiceDataSetVoiceLevel 
//  @brief Sets the voice level of a voice packet.
//
//  @param voiceData: The voice packet from HookClientVoiceData.
//  @param voiceLevel: The voice level.
func VoiceDataSetVoiceLevel(voiceData uintptr, voiceLevel float32) {
	_VoiceDataSetVoiceLevel(voiceData, voiceLevel)
}

var _SendVoiceData = func(recipients []int32, speakerSlot int32, audio []uint8, packetOffsets []uint32, sectionNumber uint32) bool {
	var __retVal bool
	__recipients := plugify.ConstructVectorInt32(recipients)
	__speakerSlot := C.int32_t(speakerSlot)
	__audio := plugify.ConstructVectorUInt8(audio)
	__packetOffsets := plugify.ConstructVectorUInt32(packetOffsets)
	__sectionNumber := C.uint32_t(sectionNumber)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SendVoiceData((*C.Vector)(unsafe.Pointer(&__recipients)), __speakerSlot, (*C.Vector)(unsafe.Pointer(&__audio)), (*C.Vector)(unsafe.Pointer(&__packetOffsets)), __sectionNumber))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__recipients)
			plugify.DestroyVectorUInt8(&__audio)
			plugify.DestroyVectorUInt32(&__packetOffsets)
		},
	}.Do()
	return __retVal
}

// SendVoiceData 
//  @brief Sends Opus voice frames to clients as if a player was talking.
//
//  @param recipients: The slots of the players who receive the voice.
//  @param speakerSlot: The slot of the player shown as talking, or -1 for none.
//  @param audio: Raw Opus packets encoded at 48 kHz, written back to back without a container (no Ogg) or length prefixes. Frame boundaries are given by packetOffsets.
//  @param packetOffsets: For every Opus packet in audio, the byte offset where it ends, in ascending order; the last offset equals the size of audio. Example: packets of 60, 55 and 62 bytes give [60, 115, 177].
//  @param sectionNumber: A number that grows with every call while one stream plays.
//
//  @return True if the voice was sent to at least one client, false otherwise.
func SendVoiceData(recipients []int32, speakerSlot int32, audio []uint8, packetOffsets []uint32, sectionNumber uint32) bool {
	return _SendVoiceData(recipients, speakerSlot, audio, packetOffsets, sectionNumber)
}

