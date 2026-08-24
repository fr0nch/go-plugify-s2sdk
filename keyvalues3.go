package s2sdk

/*
#include "keyvalues3.h"
#cgo noescape Kv3Create
#cgo noescape Kv3CreateWithCluster
#cgo noescape Kv3CreateCopy
#cgo noescape Kv3Destroy
#cgo noescape Kv3CopyFrom
#cgo noescape Kv3OverlayKeysFrom
#cgo noescape Kv3GetContext
#cgo noescape Kv3GetMetaData
#cgo noescape Kv3HasFlag
#cgo noescape Kv3HasAnyFlags
#cgo noescape Kv3GetAllFlags
#cgo noescape Kv3SetAllFlags
#cgo noescape Kv3SetFlag
#cgo noescape Kv3GetType
#cgo noescape Kv3GetTypeEx
#cgo noescape Kv3GetSubType
#cgo noescape Kv3HasInvalidMemberNames
#cgo noescape Kv3SetHasInvalidMemberNames
#cgo noescape Kv3GetTypeAsString
#cgo noescape Kv3GetSubTypeAsString
#cgo noescape Kv3ToString
#cgo noescape Kv3IsNull
#cgo noescape Kv3SetToNull
#cgo noescape Kv3IsArray
#cgo noescape Kv3IsKV3Array
#cgo noescape Kv3IsTable
#cgo noescape Kv3IsString
#cgo noescape Kv3GetBool
#cgo noescape Kv3GetChar
#cgo noescape Kv3GetUChar32
#cgo noescape Kv3GetInt8
#cgo noescape Kv3GetUInt8
#cgo noescape Kv3GetShort
#cgo noescape Kv3GetUShort
#cgo noescape Kv3GetInt
#cgo noescape Kv3GetUInt
#cgo noescape Kv3GetInt64
#cgo noescape Kv3GetUInt64
#cgo noescape Kv3GetFloat
#cgo noescape Kv3GetDouble
#cgo noescape Kv3SetBool
#cgo noescape Kv3SetChar
#cgo noescape Kv3SetUChar32
#cgo noescape Kv3SetInt8
#cgo noescape Kv3SetUInt8
#cgo noescape Kv3SetShort
#cgo noescape Kv3SetUShort
#cgo noescape Kv3SetInt
#cgo noescape Kv3SetUInt
#cgo noescape Kv3SetInt64
#cgo noescape Kv3SetUInt64
#cgo noescape Kv3SetFloat
#cgo noescape Kv3SetDouble
#cgo noescape Kv3GetPointer
#cgo noescape Kv3SetPointer
#cgo noescape Kv3GetStringToken
#cgo noescape Kv3SetStringToken
#cgo noescape Kv3GetEHandle
#cgo noescape Kv3SetEHandle
#cgo noescape Kv3GetString
#cgo noescape Kv3SetString
#cgo noescape Kv3SetStringExternal
#cgo noescape Kv3GetBinaryBlob
#cgo noescape Kv3GetBinaryBlobSize
#cgo noescape Kv3SetToBinaryBlob
#cgo noescape Kv3SetToBinaryBlobExternal
#cgo noescape Kv3GetColor
#cgo noescape Kv3SetColor
#cgo noescape Kv3GetVector
#cgo noescape Kv3GetVector2D
#cgo noescape Kv3GetVector4D
#cgo noescape Kv3GetQuaternion
#cgo noescape Kv3GetQAngle
#cgo noescape Kv3GetMatrix3x4
#cgo noescape Kv3SetVector
#cgo noescape Kv3SetVector2D
#cgo noescape Kv3SetVector4D
#cgo noescape Kv3SetQuaternion
#cgo noescape Kv3SetQAngle
#cgo noescape Kv3SetMatrix3x4
#cgo noescape Kv3GetArrayElementCount
#cgo noescape Kv3SetArrayElementCount
#cgo noescape Kv3SetToEmptyKV3Array
#cgo noescape Kv3GetArrayElement
#cgo noescape Kv3ArrayInsertElementBefore
#cgo noescape Kv3ArrayInsertElementAfter
#cgo noescape Kv3ArrayAddElementToTail
#cgo noescape Kv3ArraySwapItems
#cgo noescape Kv3ArrayRemoveElement
#cgo noescape Kv3SetToEmptyTable
#cgo noescape Kv3GetMemberCount
#cgo noescape Kv3HasMember
#cgo noescape Kv3FindMember
#cgo noescape Kv3FindOrCreateMember
#cgo noescape Kv3RemoveMember
#cgo noescape Kv3GetMemberName
#cgo noescape Kv3GetMemberByIndex
#cgo noescape Kv3GetMemberBool
#cgo noescape Kv3GetMemberChar
#cgo noescape Kv3GetMemberUChar32
#cgo noescape Kv3GetMemberInt8
#cgo noescape Kv3GetMemberUInt8
#cgo noescape Kv3GetMemberShort
#cgo noescape Kv3GetMemberUShort
#cgo noescape Kv3GetMemberInt
#cgo noescape Kv3GetMemberUInt
#cgo noescape Kv3GetMemberInt64
#cgo noescape Kv3GetMemberUInt64
#cgo noescape Kv3GetMemberFloat
#cgo noescape Kv3GetMemberDouble
#cgo noescape Kv3GetMemberPointer
#cgo noescape Kv3GetMemberStringToken
#cgo noescape Kv3GetMemberEHandle
#cgo noescape Kv3GetMemberString
#cgo noescape Kv3GetMemberColor
#cgo noescape Kv3GetMemberVector
#cgo noescape Kv3GetMemberVector2D
#cgo noescape Kv3GetMemberVector4D
#cgo noescape Kv3GetMemberQuaternion
#cgo noescape Kv3GetMemberQAngle
#cgo noescape Kv3GetMemberMatrix3x4
#cgo noescape Kv3SetMemberToNull
#cgo noescape Kv3SetMemberToEmptyArray
#cgo noescape Kv3SetMemberToEmptyTable
#cgo noescape Kv3SetMemberToBinaryBlob
#cgo noescape Kv3SetMemberToBinaryBlobExternal
#cgo noescape Kv3SetMemberToCopyOfValue
#cgo noescape Kv3SetMemberBool
#cgo noescape Kv3SetMemberChar
#cgo noescape Kv3SetMemberUChar32
#cgo noescape Kv3SetMemberInt8
#cgo noescape Kv3SetMemberUInt8
#cgo noescape Kv3SetMemberShort
#cgo noescape Kv3SetMemberUShort
#cgo noescape Kv3SetMemberInt
#cgo noescape Kv3SetMemberUInt
#cgo noescape Kv3SetMemberInt64
#cgo noescape Kv3SetMemberUInt64
#cgo noescape Kv3SetMemberFloat
#cgo noescape Kv3SetMemberDouble
#cgo noescape Kv3SetMemberPointer
#cgo noescape Kv3SetMemberStringToken
#cgo noescape Kv3SetMemberEHandle
#cgo noescape Kv3SetMemberString
#cgo noescape Kv3SetMemberStringExternal
#cgo noescape Kv3SetMemberColor
#cgo noescape Kv3SetMemberVector
#cgo noescape Kv3SetMemberVector2D
#cgo noescape Kv3SetMemberVector4D
#cgo noescape Kv3SetMemberQuaternion
#cgo noescape Kv3SetMemberQAngle
#cgo noescape Kv3SetMemberMatrix3x4
#cgo noescape Kv3DebugPrint
#cgo noescape Kv3LoadFromBuffer
#cgo noescape Kv3Load
#cgo noescape Kv3LoadFromText
#cgo noescape Kv3LoadFromFileToContext
#cgo noescape Kv3LoadFromFile
#cgo noescape Kv3LoadFromJSON
#cgo noescape Kv3LoadFromJSONFile
#cgo noescape Kv3LoadFromKV1File
#cgo noescape Kv3LoadFromKV1Text
#cgo noescape Kv3LoadFromKV1TextTranslated
#cgo noescape Kv3LoadFromKV3OrKV1
#cgo noescape Kv3LoadFromOldSchemaText
#cgo noescape Kv3LoadTextNoHeader
#cgo noescape Kv3Save
#cgo noescape Kv3SaveAsJSON
#cgo noescape Kv3SaveAsJSONString
#cgo noescape Kv3SaveAsKV1Text
#cgo noescape Kv3SaveAsKV1TextTranslated
#cgo noescape Kv3SaveTextNoHeaderToBuffer
#cgo noescape Kv3SaveTextNoHeader
#cgo noescape Kv3SaveTextToString
#cgo noescape Kv3SaveToFile
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

// Generated from s2sdk (group: keyvalues3)

var _Kv3Create = func(type_ int32, subtype int32) uintptr {
	var __retVal uintptr
	__type_ := C.int32_t(type_)
	__subtype := C.int32_t(subtype)
	__retVal = uintptr(C.Kv3Create(__type_, __subtype))
	return __retVal
}

// Kv3Create 
//  @brief Creates a new KeyValues3 object with specified type and subtype
//
//  @param type_: The KV3 type enumeration value
//  @param subtype: The KV3 subtype enumeration value
//
//  @return Pointer to the newly created KeyValues3 object
func Kv3Create(type_ int32, subtype int32) uintptr {
	return _Kv3Create(type_, subtype)
}

var _Kv3CreateWithCluster = func(cluster_elem int32, type_ int32, subtype int32) uintptr {
	var __retVal uintptr
	__cluster_elem := C.int32_t(cluster_elem)
	__type_ := C.int32_t(type_)
	__subtype := C.int32_t(subtype)
	__retVal = uintptr(C.Kv3CreateWithCluster(__cluster_elem, __type_, __subtype))
	return __retVal
}

// Kv3CreateWithCluster 
//  @brief Creates a new KeyValues3 object with cluster element, type, and subtype
//
//  @param cluster_elem: The cluster element index
//  @param type_: The KV3 type enumeration value
//  @param subtype: The KV3 subtype enumeration value
//
//  @return Pointer to the newly created KeyValues3 object
func Kv3CreateWithCluster(cluster_elem int32, type_ int32, subtype int32) uintptr {
	return _Kv3CreateWithCluster(cluster_elem, type_, subtype)
}

var _Kv3CreateCopy = func(other uintptr) uintptr {
	var __retVal uintptr
	__other := C.uintptr_t(other)
	__retVal = uintptr(C.Kv3CreateCopy(__other))
	return __retVal
}

// Kv3CreateCopy 
//  @brief Creates a copy of an existing KeyValues3 object
//
//  @param other: Pointer to the KeyValues3 object to copy
//
//  @return Pointer to the newly created copy, or nullptr if other is null
func Kv3CreateCopy(other uintptr) uintptr {
	return _Kv3CreateCopy(other)
}

var _Kv3Destroy = func(kv uintptr) {
	__kv := C.uintptr_t(kv)
	C.Kv3Destroy(__kv)
}

// Kv3Destroy 
//  @brief Destroys a KeyValues3 object and frees its memory
//
//  @param kv: Pointer to the KeyValues3 object to destroy
func Kv3Destroy(kv uintptr) {
	_Kv3Destroy(kv)
}

var _Kv3CopyFrom = func(kv uintptr, other uintptr) {
	__kv := C.uintptr_t(kv)
	__other := C.uintptr_t(other)
	C.Kv3CopyFrom(__kv, __other)
}

// Kv3CopyFrom 
//  @brief Copies data from another KeyValues3 object
//
//  @param kv: Pointer to the destination KeyValues3 object
//  @param other: Pointer to the source KeyValues3 object
func Kv3CopyFrom(kv uintptr, other uintptr) {
	_Kv3CopyFrom(kv, other)
}

var _Kv3OverlayKeysFrom = func(kv uintptr, other uintptr, depth bool) {
	__kv := C.uintptr_t(kv)
	__other := C.uintptr_t(other)
	__depth := C.bool(depth)
	C.Kv3OverlayKeysFrom(__kv, __other, __depth)
}

// Kv3OverlayKeysFrom 
//  @brief Overlays keys from another KeyValues3 object
//
//  @param kv: Pointer to the destination KeyValues3 object
//  @param other: Pointer to the source KeyValues3 object
//  @param depth: Whether to perform a deep overlay
func Kv3OverlayKeysFrom(kv uintptr, other uintptr, depth bool) {
	_Kv3OverlayKeysFrom(kv, other, depth)
}

var _Kv3GetContext = func(kv uintptr) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__retVal = uintptr(C.Kv3GetContext(__kv))
	return __retVal
}

// Kv3GetContext 
//  @brief Gets the context associated with a KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Pointer to the CKV3Arena, or nullptr if kv is null
func Kv3GetContext(kv uintptr) uintptr {
	return _Kv3GetContext(kv)
}

var _Kv3GetMetaData = func(kv uintptr, ppCtx uintptr) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__ppCtx := C.uintptr_t(ppCtx)
	__retVal = uintptr(C.Kv3GetMetaData(__kv, __ppCtx))
	return __retVal
}

// Kv3GetMetaData 
//  @brief Gets the metadata associated with a KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ppCtx: Pointer to store the context pointer
//
//  @return Pointer to the KV3MetaData_t structure, or nullptr if kv is null
func Kv3GetMetaData(kv uintptr, ppCtx uintptr) uintptr {
	return _Kv3GetMetaData(kv, ppCtx)
}

var _Kv3HasFlag = func(kv uintptr, flag uint8) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__flag := C.uint8_t(flag)
	__retVal = bool(C.Kv3HasFlag(__kv, __flag))
	return __retVal
}

// Kv3HasFlag 
//  @brief Checks if a specific flag is set
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flag: The flag to check
//
//  @return true if the flag is set, false otherwise
func Kv3HasFlag(kv uintptr, flag uint8) bool {
	return _Kv3HasFlag(kv, flag)
}

var _Kv3HasAnyFlags = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3HasAnyFlags(__kv))
	return __retVal
}

// Kv3HasAnyFlags 
//  @brief Checks if any flags are set
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if any flags are set, false otherwise
func Kv3HasAnyFlags(kv uintptr) bool {
	return _Kv3HasAnyFlags(kv)
}

var _Kv3GetAllFlags = func(kv uintptr) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__retVal = uint8(C.Kv3GetAllFlags(__kv))
	return __retVal
}

// Kv3GetAllFlags 
//  @brief Gets all flags as a bitmask
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Bitmask of all flags, or 0 if kv is null
func Kv3GetAllFlags(kv uintptr) uint8 {
	return _Kv3GetAllFlags(kv)
}

var _Kv3SetAllFlags = func(kv uintptr, flags uint8) {
	__kv := C.uintptr_t(kv)
	__flags := C.uint8_t(flags)
	C.Kv3SetAllFlags(__kv, __flags)
}

// Kv3SetAllFlags 
//  @brief Sets all flags from a bitmask
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flags: Bitmask of flags to set
func Kv3SetAllFlags(kv uintptr, flags uint8) {
	_Kv3SetAllFlags(kv, flags)
}

var _Kv3SetFlag = func(kv uintptr, flag uint8, state bool) {
	__kv := C.uintptr_t(kv)
	__flag := C.uint8_t(flag)
	__state := C.bool(state)
	C.Kv3SetFlag(__kv, __flag, __state)
}

// Kv3SetFlag 
//  @brief Sets or clears a specific flag
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flag: The flag to modify
//  @param state: true to set the flag, false to clear it
func Kv3SetFlag(kv uintptr, flag uint8, state bool) {
	_Kv3SetFlag(kv, flag, state)
}

var _Kv3GetType = func(kv uintptr) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__retVal = uint8(C.Kv3GetType(__kv))
	return __retVal
}

// Kv3GetType 
//  @brief Gets the basic type of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The type enumeration value, or 0 if kv is null
func Kv3GetType(kv uintptr) uint8 {
	return _Kv3GetType(kv)
}

var _Kv3GetTypeEx = func(kv uintptr) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__retVal = uint8(C.Kv3GetTypeEx(__kv))
	return __retVal
}

// Kv3GetTypeEx 
//  @brief Gets the extended type of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The extended type enumeration value, or 0 if kv is null
func Kv3GetTypeEx(kv uintptr) uint8 {
	return _Kv3GetTypeEx(kv)
}

var _Kv3GetSubType = func(kv uintptr) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__retVal = uint8(C.Kv3GetSubType(__kv))
	return __retVal
}

// Kv3GetSubType 
//  @brief Gets the subtype of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The subtype enumeration value, or 0 if kv is null
func Kv3GetSubType(kv uintptr) uint8 {
	return _Kv3GetSubType(kv)
}

var _Kv3HasInvalidMemberNames = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3HasInvalidMemberNames(__kv))
	return __retVal
}

// Kv3HasInvalidMemberNames 
//  @brief Checks if the object has invalid member names
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if invalid member names exist, false otherwise
func Kv3HasInvalidMemberNames(kv uintptr) bool {
	return _Kv3HasInvalidMemberNames(kv)
}

var _Kv3SetHasInvalidMemberNames = func(kv uintptr, bValue bool) {
	__kv := C.uintptr_t(kv)
	__bValue := C.bool(bValue)
	C.Kv3SetHasInvalidMemberNames(__kv, __bValue)
}

// Kv3SetHasInvalidMemberNames 
//  @brief Sets the invalid member names flag
//
//  @param kv: Pointer to the KeyValues3 object
//  @param bValue: true to mark as having invalid member names, false otherwise
func Kv3SetHasInvalidMemberNames(kv uintptr, bValue bool) {
	_Kv3SetHasInvalidMemberNames(kv, bValue)
}

var _Kv3GetTypeAsString = func(kv uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetTypeAsString(__kv)
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

// Kv3GetTypeAsString 
//  @brief Gets the type as a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return String representation of the type, or empty string if kv is null
func Kv3GetTypeAsString(kv uintptr) string {
	return _Kv3GetTypeAsString(kv)
}

var _Kv3GetSubTypeAsString = func(kv uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetSubTypeAsString(__kv)
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

// Kv3GetSubTypeAsString 
//  @brief Gets the subtype as a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return String representation of the subtype, or empty string if kv is null
func Kv3GetSubTypeAsString(kv uintptr) string {
	return _Kv3GetSubTypeAsString(kv)
}

var _Kv3ToString = func(kv uintptr, flags uint32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__native := C.Kv3ToString(__kv, __flags)
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

// Kv3ToString 
//  @brief Converts the KeyValues3 object to a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flags: Formatting flags for the string conversion
//
//  @return String representation of the object, or empty string if kv is null
func Kv3ToString(kv uintptr, flags uint32) string {
	return _Kv3ToString(kv, flags)
}

var _Kv3IsNull = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3IsNull(__kv))
	return __retVal
}

// Kv3IsNull 
//  @brief Checks if the KeyValues3 object is null
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is null or the pointer is null, false otherwise
func Kv3IsNull(kv uintptr) bool {
	return _Kv3IsNull(kv)
}

var _Kv3SetToNull = func(kv uintptr) {
	__kv := C.uintptr_t(kv)
	C.Kv3SetToNull(__kv)
}

// Kv3SetToNull 
//  @brief Sets the KeyValues3 object to null
//
//  @param kv: Pointer to the KeyValues3 object
func Kv3SetToNull(kv uintptr) {
	_Kv3SetToNull(kv)
}

var _Kv3IsArray = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3IsArray(__kv))
	return __retVal
}

// Kv3IsArray 
//  @brief Checks if the KeyValues3 object is an array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is an array, false otherwise
func Kv3IsArray(kv uintptr) bool {
	return _Kv3IsArray(kv)
}

var _Kv3IsKV3Array = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3IsKV3Array(__kv))
	return __retVal
}

// Kv3IsKV3Array 
//  @brief Checks if the KeyValues3 object is a KV3 array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a KV3 array, false otherwise
func Kv3IsKV3Array(kv uintptr) bool {
	return _Kv3IsKV3Array(kv)
}

var _Kv3IsTable = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3IsTable(__kv))
	return __retVal
}

// Kv3IsTable 
//  @brief Checks if the KeyValues3 object is a table
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a table, false otherwise
func Kv3IsTable(kv uintptr) bool {
	return _Kv3IsTable(kv)
}

var _Kv3IsString = func(kv uintptr) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__retVal = bool(C.Kv3IsString(__kv))
	return __retVal
}

// Kv3IsString 
//  @brief Checks if the KeyValues3 object is a string
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a string, false otherwise
func Kv3IsString(kv uintptr) bool {
	return _Kv3IsString(kv)
}

var _Kv3GetBool = func(kv uintptr, defaultValue bool) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__defaultValue := C.bool(defaultValue)
	__retVal = bool(C.Kv3GetBool(__kv, __defaultValue))
	return __retVal
}

// Kv3GetBool 
//  @brief Gets the boolean value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Boolean value or defaultValue
func Kv3GetBool(kv uintptr, defaultValue bool) bool {
	return _Kv3GetBool(kv, defaultValue)
}

var _Kv3GetChar = func(kv uintptr, defaultValue int8) int8 {
	var __retVal int8
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int8_t(defaultValue)
	__retVal = int8(C.Kv3GetChar(__kv, __defaultValue))
	return __retVal
}

// Kv3GetChar 
//  @brief Gets the char value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Char value or defaultValue
func Kv3GetChar(kv uintptr, defaultValue int8) int8 {
	return _Kv3GetChar(kv, defaultValue)
}

var _Kv3GetUChar32 = func(kv uintptr, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint32_t(defaultValue)
	__retVal = uint32(C.Kv3GetUChar32(__kv, __defaultValue))
	return __retVal
}

// Kv3GetUChar32 
//  @brief Gets the 32-bit Unicode character value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return 32-bit Unicode character value or defaultValue
func Kv3GetUChar32(kv uintptr, defaultValue uint32) uint32 {
	return _Kv3GetUChar32(kv, defaultValue)
}

var _Kv3GetInt8 = func(kv uintptr, defaultValue int8) int8 {
	var __retVal int8
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int8_t(defaultValue)
	__retVal = int8(C.Kv3GetInt8(__kv, __defaultValue))
	return __retVal
}

// Kv3GetInt8 
//  @brief Gets the signed 8-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int8_t value or defaultValue
func Kv3GetInt8(kv uintptr, defaultValue int8) int8 {
	return _Kv3GetInt8(kv, defaultValue)
}

var _Kv3GetUInt8 = func(kv uintptr, defaultValue uint8) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint8_t(defaultValue)
	__retVal = uint8(C.Kv3GetUInt8(__kv, __defaultValue))
	return __retVal
}

// Kv3GetUInt8 
//  @brief Gets the unsigned 8-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint8_t value or defaultValue
func Kv3GetUInt8(kv uintptr, defaultValue uint8) uint8 {
	return _Kv3GetUInt8(kv, defaultValue)
}

var _Kv3GetShort = func(kv uintptr, defaultValue int16) int16 {
	var __retVal int16
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int16_t(defaultValue)
	__retVal = int16(C.Kv3GetShort(__kv, __defaultValue))
	return __retVal
}

// Kv3GetShort 
//  @brief Gets the signed 16-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int16_t value or defaultValue
func Kv3GetShort(kv uintptr, defaultValue int16) int16 {
	return _Kv3GetShort(kv, defaultValue)
}

var _Kv3GetUShort = func(kv uintptr, defaultValue uint16) uint16 {
	var __retVal uint16
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint16_t(defaultValue)
	__retVal = uint16(C.Kv3GetUShort(__kv, __defaultValue))
	return __retVal
}

// Kv3GetUShort 
//  @brief Gets the unsigned 16-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint16_t value or defaultValue
func Kv3GetUShort(kv uintptr, defaultValue uint16) uint16 {
	return _Kv3GetUShort(kv, defaultValue)
}

var _Kv3GetInt = func(kv uintptr, defaultValue int32) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int32_t(defaultValue)
	__retVal = int32(C.Kv3GetInt(__kv, __defaultValue))
	return __retVal
}

// Kv3GetInt 
//  @brief Gets the signed 32-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int32_t value or defaultValue
func Kv3GetInt(kv uintptr, defaultValue int32) int32 {
	return _Kv3GetInt(kv, defaultValue)
}

var _Kv3GetUInt = func(kv uintptr, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint32_t(defaultValue)
	__retVal = uint32(C.Kv3GetUInt(__kv, __defaultValue))
	return __retVal
}

// Kv3GetUInt 
//  @brief Gets the unsigned 32-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint32_t value or defaultValue
func Kv3GetUInt(kv uintptr, defaultValue uint32) uint32 {
	return _Kv3GetUInt(kv, defaultValue)
}

var _Kv3GetInt64 = func(kv uintptr, defaultValue int64) int64 {
	var __retVal int64
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int64_t(defaultValue)
	__retVal = int64(C.Kv3GetInt64(__kv, __defaultValue))
	return __retVal
}

// Kv3GetInt64 
//  @brief Gets the signed 64-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int64_t value or defaultValue
func Kv3GetInt64(kv uintptr, defaultValue int64) int64 {
	return _Kv3GetInt64(kv, defaultValue)
}

var _Kv3GetUInt64 = func(kv uintptr, defaultValue uint64) uint64 {
	var __retVal uint64
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint64_t(defaultValue)
	__retVal = uint64(C.Kv3GetUInt64(__kv, __defaultValue))
	return __retVal
}

// Kv3GetUInt64 
//  @brief Gets the unsigned 64-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint64_t value or defaultValue
func Kv3GetUInt64(kv uintptr, defaultValue uint64) uint64 {
	return _Kv3GetUInt64(kv, defaultValue)
}

var _Kv3GetFloat = func(kv uintptr, defaultValue float32) float32 {
	var __retVal float32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.float(defaultValue)
	__retVal = float32(C.Kv3GetFloat(__kv, __defaultValue))
	return __retVal
}

// Kv3GetFloat 
//  @brief Gets the float value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Float value or defaultValue
func Kv3GetFloat(kv uintptr, defaultValue float32) float32 {
	return _Kv3GetFloat(kv, defaultValue)
}

var _Kv3GetDouble = func(kv uintptr, defaultValue float64) float64 {
	var __retVal float64
	__kv := C.uintptr_t(kv)
	__defaultValue := C.double(defaultValue)
	__retVal = float64(C.Kv3GetDouble(__kv, __defaultValue))
	return __retVal
}

// Kv3GetDouble 
//  @brief Gets the double value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Double value or defaultValue
func Kv3GetDouble(kv uintptr, defaultValue float64) float64 {
	return _Kv3GetDouble(kv, defaultValue)
}

var _Kv3SetBool = func(kv uintptr, value bool) {
	__kv := C.uintptr_t(kv)
	__value := C.bool(value)
	C.Kv3SetBool(__kv, __value)
}

// Kv3SetBool 
//  @brief Sets the KeyValues3 object to a boolean value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Boolean value to set
func Kv3SetBool(kv uintptr, value bool) {
	_Kv3SetBool(kv, value)
}

var _Kv3SetChar = func(kv uintptr, value int8) {
	__kv := C.uintptr_t(kv)
	__value := C.int8_t(value)
	C.Kv3SetChar(__kv, __value)
}

// Kv3SetChar 
//  @brief Sets the KeyValues3 object to a char value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Char value to set
func Kv3SetChar(kv uintptr, value int8) {
	_Kv3SetChar(kv, value)
}

var _Kv3SetUChar32 = func(kv uintptr, value uint32) {
	__kv := C.uintptr_t(kv)
	__value := C.uint32_t(value)
	C.Kv3SetUChar32(__kv, __value)
}

// Kv3SetUChar32 
//  @brief Sets the KeyValues3 object to a 32-bit Unicode character value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: 32-bit Unicode character value to set
func Kv3SetUChar32(kv uintptr, value uint32) {
	_Kv3SetUChar32(kv, value)
}

var _Kv3SetInt8 = func(kv uintptr, value int8) {
	__kv := C.uintptr_t(kv)
	__value := C.int8_t(value)
	C.Kv3SetInt8(__kv, __value)
}

// Kv3SetInt8 
//  @brief Sets the KeyValues3 object to a signed 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int8_t value to set
func Kv3SetInt8(kv uintptr, value int8) {
	_Kv3SetInt8(kv, value)
}

var _Kv3SetUInt8 = func(kv uintptr, value uint8) {
	__kv := C.uintptr_t(kv)
	__value := C.uint8_t(value)
	C.Kv3SetUInt8(__kv, __value)
}

// Kv3SetUInt8 
//  @brief Sets the KeyValues3 object to an unsigned 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint8_t value to set
func Kv3SetUInt8(kv uintptr, value uint8) {
	_Kv3SetUInt8(kv, value)
}

var _Kv3SetShort = func(kv uintptr, value int16) {
	__kv := C.uintptr_t(kv)
	__value := C.int16_t(value)
	C.Kv3SetShort(__kv, __value)
}

// Kv3SetShort 
//  @brief Sets the KeyValues3 object to a signed 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int16_t value to set
func Kv3SetShort(kv uintptr, value int16) {
	_Kv3SetShort(kv, value)
}

var _Kv3SetUShort = func(kv uintptr, value uint16) {
	__kv := C.uintptr_t(kv)
	__value := C.uint16_t(value)
	C.Kv3SetUShort(__kv, __value)
}

// Kv3SetUShort 
//  @brief Sets the KeyValues3 object to an unsigned 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint16_t value to set
func Kv3SetUShort(kv uintptr, value uint16) {
	_Kv3SetUShort(kv, value)
}

var _Kv3SetInt = func(kv uintptr, value int32) {
	__kv := C.uintptr_t(kv)
	__value := C.int32_t(value)
	C.Kv3SetInt(__kv, __value)
}

// Kv3SetInt 
//  @brief Sets the KeyValues3 object to a signed 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int32_t value to set
func Kv3SetInt(kv uintptr, value int32) {
	_Kv3SetInt(kv, value)
}

var _Kv3SetUInt = func(kv uintptr, value uint32) {
	__kv := C.uintptr_t(kv)
	__value := C.uint32_t(value)
	C.Kv3SetUInt(__kv, __value)
}

// Kv3SetUInt 
//  @brief Sets the KeyValues3 object to an unsigned 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint32_t value to set
func Kv3SetUInt(kv uintptr, value uint32) {
	_Kv3SetUInt(kv, value)
}

var _Kv3SetInt64 = func(kv uintptr, value int64) {
	__kv := C.uintptr_t(kv)
	__value := C.int64_t(value)
	C.Kv3SetInt64(__kv, __value)
}

// Kv3SetInt64 
//  @brief Sets the KeyValues3 object to a signed 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int64_t value to set
func Kv3SetInt64(kv uintptr, value int64) {
	_Kv3SetInt64(kv, value)
}

var _Kv3SetUInt64 = func(kv uintptr, value uint64) {
	__kv := C.uintptr_t(kv)
	__value := C.uint64_t(value)
	C.Kv3SetUInt64(__kv, __value)
}

// Kv3SetUInt64 
//  @brief Sets the KeyValues3 object to an unsigned 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint64_t value to set
func Kv3SetUInt64(kv uintptr, value uint64) {
	_Kv3SetUInt64(kv, value)
}

var _Kv3SetFloat = func(kv uintptr, value float32) {
	__kv := C.uintptr_t(kv)
	__value := C.float(value)
	C.Kv3SetFloat(__kv, __value)
}

// Kv3SetFloat 
//  @brief Sets the KeyValues3 object to a float value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Float value to set
func Kv3SetFloat(kv uintptr, value float32) {
	_Kv3SetFloat(kv, value)
}

var _Kv3SetDouble = func(kv uintptr, value float64) {
	__kv := C.uintptr_t(kv)
	__value := C.double(value)
	C.Kv3SetDouble(__kv, __value)
}

// Kv3SetDouble 
//  @brief Sets the KeyValues3 object to a double value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Double value to set
func Kv3SetDouble(kv uintptr, value float64) {
	_Kv3SetDouble(kv, value)
}

var _Kv3GetPointer = func(kv uintptr, defaultValue uintptr) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uintptr_t(defaultValue)
	__retVal = uintptr(C.Kv3GetPointer(__kv, __defaultValue))
	return __retVal
}

// Kv3GetPointer 
//  @brief Gets the pointer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null
//
//  @return Pointer value as uintptr_t or defaultValue
func Kv3GetPointer(kv uintptr, defaultValue uintptr) uintptr {
	return _Kv3GetPointer(kv, defaultValue)
}

var _Kv3SetPointer = func(kv uintptr, ptr uintptr) {
	__kv := C.uintptr_t(kv)
	__ptr := C.uintptr_t(ptr)
	C.Kv3SetPointer(__kv, __ptr)
}

// Kv3SetPointer 
//  @brief Sets the KeyValues3 object to a pointer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ptr: Pointer value as uintptr_t to set
func Kv3SetPointer(kv uintptr, ptr uintptr) {
	_Kv3SetPointer(kv, ptr)
}

var _Kv3GetStringToken = func(kv uintptr, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.uint32_t(defaultValue)
	__retVal = uint32(C.Kv3GetStringToken(__kv, __defaultValue))
	return __retVal
}

// Kv3GetStringToken 
//  @brief Gets the string token value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default token value to return if kv is null
//
//  @return String token hash code or defaultValue
func Kv3GetStringToken(kv uintptr, defaultValue uint32) uint32 {
	return _Kv3GetStringToken(kv, defaultValue)
}

var _Kv3SetStringToken = func(kv uintptr, token uint32) {
	__kv := C.uintptr_t(kv)
	__token := C.uint32_t(token)
	C.Kv3SetStringToken(__kv, __token)
}

// Kv3SetStringToken 
//  @brief Sets the KeyValues3 object to a string token value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param token: String token hash code to set
func Kv3SetStringToken(kv uintptr, token uint32) {
	_Kv3SetStringToken(kv, token)
}

var _Kv3GetEHandle = func(kv uintptr, defaultValue int32) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__defaultValue := C.int32_t(defaultValue)
	__retVal = int32(C.Kv3GetEHandle(__kv, __defaultValue))
	return __retVal
}

// Kv3GetEHandle 
//  @brief Gets the entity handle value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default entity handle value to return if kv is null
//
//  @return Entity handle as int32_t or defaultValue
func Kv3GetEHandle(kv uintptr, defaultValue int32) int32 {
	return _Kv3GetEHandle(kv, defaultValue)
}

var _Kv3SetEHandle = func(kv uintptr, ehandle int32) {
	__kv := C.uintptr_t(kv)
	__ehandle := C.int32_t(ehandle)
	C.Kv3SetEHandle(__kv, __ehandle)
}

// Kv3SetEHandle 
//  @brief Sets the KeyValues3 object to an entity handle value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ehandle: Entity handle value to set
func Kv3SetEHandle(kv uintptr, ehandle int32) {
	_Kv3SetEHandle(kv, ehandle)
}

var _Kv3GetString = func(kv uintptr, defaultValue string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	__defaultValue := plugify.ConstructString(defaultValue)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetString(__kv, (*C.String)(unsafe.Pointer(&__defaultValue)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__defaultValue)
		},
	}.Do()
	return __retVal
}

// Kv3GetString 
//  @brief Gets the string value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default string to return if kv is null or value is empty
//
//  @return String value or defaultValue
func Kv3GetString(kv uintptr, defaultValue string) string {
	return _Kv3GetString(kv, defaultValue)
}

var _Kv3SetString = func(kv uintptr, str string, subtype uint8) {
	__kv := C.uintptr_t(kv)
	__str := plugify.ConstructString(str)
	__subtype := C.uint8_t(subtype)
	plugify.Block {
		Try: func() {
			C.Kv3SetString(__kv, (*C.String)(unsafe.Pointer(&__str)), __subtype)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__str)
		},
	}.Do()
}

// Kv3SetString 
//  @brief Sets the KeyValues3 object to a string value (copies the string)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param str: String value to set
//  @param subtype: String subtype enumeration value
func Kv3SetString(kv uintptr, str string, subtype uint8) {
	_Kv3SetString(kv, str, subtype)
}

var _Kv3SetStringExternal = func(kv uintptr, str string, subtype uint8) {
	__kv := C.uintptr_t(kv)
	__str := plugify.ConstructString(str)
	__subtype := C.uint8_t(subtype)
	plugify.Block {
		Try: func() {
			C.Kv3SetStringExternal(__kv, (*C.String)(unsafe.Pointer(&__str)), __subtype)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__str)
		},
	}.Do()
}

// Kv3SetStringExternal 
//  @brief Sets the KeyValues3 object to an external string value (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param str: External string value to reference
//  @param subtype: String subtype enumeration value
func Kv3SetStringExternal(kv uintptr, str string, subtype uint8) {
	_Kv3SetStringExternal(kv, str, subtype)
}

var _Kv3GetBinaryBlob = func(kv uintptr) []uint8 {
	var __retVal []uint8
	var __retVal_native plugify.PlgVector
	__kv := C.uintptr_t(kv)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetBinaryBlob(__kv)
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

// Kv3GetBinaryBlob 
//  @brief Gets the binary blob from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Vector containing the binary blob data, or empty vector if kv is null
func Kv3GetBinaryBlob(kv uintptr) []uint8 {
	return _Kv3GetBinaryBlob(kv)
}

var _Kv3GetBinaryBlobSize = func(kv uintptr) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__retVal = int32(C.Kv3GetBinaryBlobSize(__kv))
	return __retVal
}

// Kv3GetBinaryBlobSize 
//  @brief Gets the size of the binary blob in the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Size of the binary blob in bytes, or 0 if kv is null
func Kv3GetBinaryBlobSize(kv uintptr) int32 {
	return _Kv3GetBinaryBlobSize(kv)
}

var _Kv3SetToBinaryBlob = func(kv uintptr, blob []uint8) {
	__kv := C.uintptr_t(kv)
	__blob := plugify.ConstructVectorUInt8(blob)
	plugify.Block {
		Try: func() {
			C.Kv3SetToBinaryBlob(__kv, (*C.Vector)(unsafe.Pointer(&__blob)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__blob)
		},
	}.Do()
}

// Kv3SetToBinaryBlob 
//  @brief Sets the KeyValues3 object to a binary blob (copies the data)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param blob: Vector containing the binary blob data
func Kv3SetToBinaryBlob(kv uintptr, blob []uint8) {
	_Kv3SetToBinaryBlob(kv, blob)
}

var _Kv3SetToBinaryBlobExternal = func(kv uintptr, blob []uint8, free_mem bool) {
	__kv := C.uintptr_t(kv)
	__blob := plugify.ConstructVectorUInt8(blob)
	__free_mem := C.bool(free_mem)
	plugify.Block {
		Try: func() {
			C.Kv3SetToBinaryBlobExternal(__kv, (*C.Vector)(unsafe.Pointer(&__blob)), __free_mem)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__blob)
		},
	}.Do()
}

// Kv3SetToBinaryBlobExternal 
//  @brief Sets the KeyValues3 object to an external binary blob (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param blob: Vector containing the external binary blob data
//  @param free_mem: Whether to free the memory when the object is destroyed
func Kv3SetToBinaryBlobExternal(kv uintptr, blob []uint8, free_mem bool) {
	_Kv3SetToBinaryBlobExternal(kv, blob, free_mem)
}

var _Kv3GetColor = func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetColor(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetColor 
//  @brief Gets the color value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default color value to return if kv is null
//
//  @return Color value as vec4 or defaultValue
func Kv3GetColor(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetColor(kv, defaultValue)
}

var _Kv3SetColor = func(kv uintptr, color plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	C.Kv3SetColor(__kv, &__color)
}

// Kv3SetColor 
//  @brief Sets the KeyValues3 object to a color value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param color: Color value as vec4 to set
func Kv3SetColor(kv uintptr, color plugify.Vector4) {
	_Kv3SetColor(kv, color)
}

var _Kv3GetVector = func(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetVector(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetVector 
//  @brief Gets the 3D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default vector to return if kv is null
//
//  @return 3D vector or defaultValue
func Kv3GetVector(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3 {
	return _Kv3GetVector(kv, defaultValue)
}

var _Kv3GetVector2D = func(kv uintptr, defaultValue plugify.Vector2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector2)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetVector2D(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetVector2D 
//  @brief Gets the 2D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default 2D vector to return if kv is null
//
//  @return 2D vector or defaultValue
func Kv3GetVector2D(kv uintptr, defaultValue plugify.Vector2) plugify.Vector2 {
	return _Kv3GetVector2D(kv, defaultValue)
}

var _Kv3GetVector4D = func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetVector4D(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetVector4D 
//  @brief Gets the 4D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default 4D vector to return if kv is null
//
//  @return 4D vector or defaultValue
func Kv3GetVector4D(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetVector4D(kv, defaultValue)
}

var _Kv3GetQuaternion = func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetQuaternion(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetQuaternion 
//  @brief Gets the quaternion value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default quaternion to return if kv is null
//
//  @return Quaternion as vec4 or defaultValue
func Kv3GetQuaternion(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetQuaternion(kv, defaultValue)
}

var _Kv3GetQAngle = func(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetQAngle(__kv, &__defaultValue)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetQAngle 
//  @brief Gets the angle (QAngle) value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default angle to return if kv is null
//
//  @return QAngle as vec3 or defaultValue
func Kv3GetQAngle(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3 {
	return _Kv3GetQAngle(kv, defaultValue)
}

var _Kv3GetMatrix3x4 = func(kv uintptr, defaultValue plugify.Matrix4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__kv := C.uintptr_t(kv)
	__defaultValue := *(*C.Matrix4x4)(unsafe.Pointer(&defaultValue))
	__native := C.Kv3GetMatrix3x4(__kv, &__defaultValue)
	__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
	return __retVal
}

// Kv3GetMatrix3x4 
//  @brief Gets the 3x4 matrix value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default matrix to return if kv is null
//
//  @return 3x4 matrix as mat4x4 or defaultValue
func Kv3GetMatrix3x4(kv uintptr, defaultValue plugify.Matrix4x4) plugify.Matrix4x4 {
	return _Kv3GetMatrix3x4(kv, defaultValue)
}

var _Kv3SetVector = func(kv uintptr, vec plugify.Vector3) {
	__kv := C.uintptr_t(kv)
	__vec := *(*C.Vector3)(unsafe.Pointer(&vec))
	C.Kv3SetVector(__kv, &__vec)
}

// Kv3SetVector 
//  @brief Sets the KeyValues3 object to a 3D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec: 3D vector to set
func Kv3SetVector(kv uintptr, vec plugify.Vector3) {
	_Kv3SetVector(kv, vec)
}

var _Kv3SetVector2D = func(kv uintptr, vec2d plugify.Vector2) {
	__kv := C.uintptr_t(kv)
	__vec2d := *(*C.Vector2)(unsafe.Pointer(&vec2d))
	C.Kv3SetVector2D(__kv, &__vec2d)
}

// Kv3SetVector2D 
//  @brief Sets the KeyValues3 object to a 2D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec2d: 2D vector to set
func Kv3SetVector2D(kv uintptr, vec2d plugify.Vector2) {
	_Kv3SetVector2D(kv, vec2d)
}

var _Kv3SetVector4D = func(kv uintptr, vec4d plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__vec4d := *(*C.Vector4)(unsafe.Pointer(&vec4d))
	C.Kv3SetVector4D(__kv, &__vec4d)
}

// Kv3SetVector4D 
//  @brief Sets the KeyValues3 object to a 4D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec4d: 4D vector to set
func Kv3SetVector4D(kv uintptr, vec4d plugify.Vector4) {
	_Kv3SetVector4D(kv, vec4d)
}

var _Kv3SetQuaternion = func(kv uintptr, quat plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__quat := *(*C.Vector4)(unsafe.Pointer(&quat))
	C.Kv3SetQuaternion(__kv, &__quat)
}

// Kv3SetQuaternion 
//  @brief Sets the KeyValues3 object to a quaternion value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param quat: Quaternion to set (as vec4)
func Kv3SetQuaternion(kv uintptr, quat plugify.Vector4) {
	_Kv3SetQuaternion(kv, quat)
}

var _Kv3SetQAngle = func(kv uintptr, ang plugify.Vector3) {
	__kv := C.uintptr_t(kv)
	__ang := *(*C.Vector3)(unsafe.Pointer(&ang))
	C.Kv3SetQAngle(__kv, &__ang)
}

// Kv3SetQAngle 
//  @brief Sets the KeyValues3 object to an angle (QAngle) value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ang: QAngle to set (as vec3)
func Kv3SetQAngle(kv uintptr, ang plugify.Vector3) {
	_Kv3SetQAngle(kv, ang)
}

var _Kv3SetMatrix3x4 = func(kv uintptr, matrix plugify.Matrix4x4) {
	__kv := C.uintptr_t(kv)
	__matrix := *(*C.Matrix4x4)(unsafe.Pointer(&matrix))
	C.Kv3SetMatrix3x4(__kv, &__matrix)
}

// Kv3SetMatrix3x4 
//  @brief Sets the KeyValues3 object to a 3x4 matrix value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param matrix: 3x4 matrix to set (as mat4x4)
func Kv3SetMatrix3x4(kv uintptr, matrix plugify.Matrix4x4) {
	_Kv3SetMatrix3x4(kv, matrix)
}

var _Kv3GetArrayElementCount = func(kv uintptr) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__retVal = int32(C.Kv3GetArrayElementCount(__kv))
	return __retVal
}

// Kv3GetArrayElementCount 
//  @brief Gets the number of elements in the array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Number of array elements, or 0 if kv is null or not an array
func Kv3GetArrayElementCount(kv uintptr) int32 {
	return _Kv3GetArrayElementCount(kv)
}

var _Kv3SetArrayElementCount = func(kv uintptr, count int32, type_ uint8, subtype uint8) {
	__kv := C.uintptr_t(kv)
	__count := C.int32_t(count)
	__type_ := C.uint8_t(type_)
	__subtype := C.uint8_t(subtype)
	C.Kv3SetArrayElementCount(__kv, __count, __type_, __subtype)
}

// Kv3SetArrayElementCount 
//  @brief Sets the number of elements in the array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param count: Number of elements to set
//  @param type_: Type of array elements
//  @param subtype: Subtype of array elements
func Kv3SetArrayElementCount(kv uintptr, count int32, type_ uint8, subtype uint8) {
	_Kv3SetArrayElementCount(kv, count, type_, subtype)
}

var _Kv3SetToEmptyKV3Array = func(kv uintptr) {
	__kv := C.uintptr_t(kv)
	C.Kv3SetToEmptyKV3Array(__kv)
}

// Kv3SetToEmptyKV3Array 
//  @brief Sets the KeyValues3 object to an empty KV3 array
//
//  @param kv: Pointer to the KeyValues3 object
func Kv3SetToEmptyKV3Array(kv uintptr) {
	_Kv3SetToEmptyKV3Array(kv)
}

var _Kv3GetArrayElement = func(kv uintptr, elem int32) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__elem := C.int32_t(elem)
	__retVal = uintptr(C.Kv3GetArrayElement(__kv, __elem))
	return __retVal
}

// Kv3GetArrayElement 
//  @brief Gets an array element at the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index of the element to get
//
//  @return Pointer to the element KeyValues3 object, or nullptr if invalid
func Kv3GetArrayElement(kv uintptr, elem int32) uintptr {
	return _Kv3GetArrayElement(kv, elem)
}

var _Kv3ArrayInsertElementBefore = func(kv uintptr, elem int32) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__elem := C.int32_t(elem)
	__retVal = uintptr(C.Kv3ArrayInsertElementBefore(__kv, __elem))
	return __retVal
}

// Kv3ArrayInsertElementBefore 
//  @brief Inserts a new element before the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index before which to insert
//
//  @return Pointer to the newly inserted element, or nullptr if invalid
func Kv3ArrayInsertElementBefore(kv uintptr, elem int32) uintptr {
	return _Kv3ArrayInsertElementBefore(kv, elem)
}

var _Kv3ArrayInsertElementAfter = func(kv uintptr, elem int32) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__elem := C.int32_t(elem)
	__retVal = uintptr(C.Kv3ArrayInsertElementAfter(__kv, __elem))
	return __retVal
}

// Kv3ArrayInsertElementAfter 
//  @brief Inserts a new element after the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index after which to insert
//
//  @return Pointer to the newly inserted element, or nullptr if invalid
func Kv3ArrayInsertElementAfter(kv uintptr, elem int32) uintptr {
	return _Kv3ArrayInsertElementAfter(kv, elem)
}

var _Kv3ArrayAddElementToTail = func(kv uintptr) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__retVal = uintptr(C.Kv3ArrayAddElementToTail(__kv))
	return __retVal
}

// Kv3ArrayAddElementToTail 
//  @brief Adds a new element to the end of the array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Pointer to the newly added element, or nullptr if invalid
func Kv3ArrayAddElementToTail(kv uintptr) uintptr {
	return _Kv3ArrayAddElementToTail(kv)
}

var _Kv3ArraySwapItems = func(kv uintptr, idx1 int32, idx2 int32) {
	__kv := C.uintptr_t(kv)
	__idx1 := C.int32_t(idx1)
	__idx2 := C.int32_t(idx2)
	C.Kv3ArraySwapItems(__kv, __idx1, __idx2)
}

// Kv3ArraySwapItems 
//  @brief Swaps two array elements
//
//  @param kv: Pointer to the KeyValues3 object
//  @param idx1: Index of the first element
//  @param idx2: Index of the second element
func Kv3ArraySwapItems(kv uintptr, idx1 int32, idx2 int32) {
	_Kv3ArraySwapItems(kv, idx1, idx2)
}

var _Kv3ArrayRemoveElement = func(kv uintptr, elem int32) {
	__kv := C.uintptr_t(kv)
	__elem := C.int32_t(elem)
	C.Kv3ArrayRemoveElement(__kv, __elem)
}

// Kv3ArrayRemoveElement 
//  @brief Removes an element from the array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index of the element to remove
func Kv3ArrayRemoveElement(kv uintptr, elem int32) {
	_Kv3ArrayRemoveElement(kv, elem)
}

var _Kv3SetToEmptyTable = func(kv uintptr) {
	__kv := C.uintptr_t(kv)
	C.Kv3SetToEmptyTable(__kv)
}

// Kv3SetToEmptyTable 
//  @brief Sets the KeyValues3 object to an empty table
//
//  @param kv: Pointer to the KeyValues3 object
func Kv3SetToEmptyTable(kv uintptr) {
	_Kv3SetToEmptyTable(kv)
}

var _Kv3GetMemberCount = func(kv uintptr) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__retVal = int32(C.Kv3GetMemberCount(__kv))
	return __retVal
}

// Kv3GetMemberCount 
//  @brief Gets the number of members in the table
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Number of table members, or 0 if kv is null or not a table
func Kv3GetMemberCount(kv uintptr) int32 {
	return _Kv3GetMemberCount(kv)
}

var _Kv3HasMember = func(kv uintptr, name string) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3HasMember(__kv, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3HasMember 
//  @brief Checks if a member with the specified name exists
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to check
//
//  @return true if the member exists, false otherwise
func Kv3HasMember(kv uintptr, name string) bool {
	return _Kv3HasMember(kv, name)
}

var _Kv3FindMember = func(kv uintptr, name string) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.Kv3FindMember(__kv, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3FindMember 
//  @brief Finds a member by name
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to find
//
//  @return Pointer to the member KeyValues3 object, or nullptr if not found
func Kv3FindMember(kv uintptr, name string) uintptr {
	return _Kv3FindMember(kv, name)
}

var _Kv3FindOrCreateMember = func(kv uintptr, name string) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.Kv3FindOrCreateMember(__kv, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3FindOrCreateMember 
//  @brief Finds a member by name, or creates it if it doesn't exist
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to find or create
//
//  @return Pointer to the member KeyValues3 object, or nullptr if kv is null
func Kv3FindOrCreateMember(kv uintptr, name string) uintptr {
	return _Kv3FindOrCreateMember(kv, name)
}

var _Kv3RemoveMember = func(kv uintptr, name string) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3RemoveMember(__kv, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3RemoveMember 
//  @brief Removes a member from the table
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to remove
//
//  @return true if the member was removed, false otherwise
func Kv3RemoveMember(kv uintptr, name string) bool {
	return _Kv3RemoveMember(kv, name)
}

var _Kv3GetMemberName = func(kv uintptr, index int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	__index := C.int32_t(index)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberName(__kv, __index)
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

// Kv3GetMemberName 
//  @brief Gets the name of a member at the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param index: Index of the member
//
//  @return Name of the member, or empty string if invalid
func Kv3GetMemberName(kv uintptr, index int32) string {
	return _Kv3GetMemberName(kv, index)
}

var _Kv3GetMemberByIndex = func(kv uintptr, index int32) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__index := C.int32_t(index)
	__retVal = uintptr(C.Kv3GetMemberByIndex(__kv, __index))
	return __retVal
}

// Kv3GetMemberByIndex 
//  @brief Gets a member by index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param index: Index of the member to get
//
//  @return Pointer to the member KeyValues3 object, or nullptr if invalid
func Kv3GetMemberByIndex(kv uintptr, index int32) uintptr {
	return _Kv3GetMemberByIndex(kv, index)
}

var _Kv3GetMemberBool = func(kv uintptr, name string, defaultValue bool) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.bool(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3GetMemberBool(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberBool 
//  @brief Gets a boolean value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Boolean value or defaultValue
func Kv3GetMemberBool(kv uintptr, name string, defaultValue bool) bool {
	return _Kv3GetMemberBool(kv, name, defaultValue)
}

var _Kv3GetMemberChar = func(kv uintptr, name string, defaultValue int8) int8 {
	var __retVal int8
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int8_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int8(C.Kv3GetMemberChar(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberChar 
//  @brief Gets a char value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Char value or defaultValue
func Kv3GetMemberChar(kv uintptr, name string, defaultValue int8) int8 {
	return _Kv3GetMemberChar(kv, name, defaultValue)
}

var _Kv3GetMemberUChar32 = func(kv uintptr, name string, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint32_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint32(C.Kv3GetMemberUChar32(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberUChar32 
//  @brief Gets a 32-bit Unicode character value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return 32-bit Unicode character value or defaultValue
func Kv3GetMemberUChar32(kv uintptr, name string, defaultValue uint32) uint32 {
	return _Kv3GetMemberUChar32(kv, name, defaultValue)
}

var _Kv3GetMemberInt8 = func(kv uintptr, name string, defaultValue int8) int8 {
	var __retVal int8
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int8_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int8(C.Kv3GetMemberInt8(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberInt8 
//  @brief Gets a signed 8-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int8_t value or defaultValue
func Kv3GetMemberInt8(kv uintptr, name string, defaultValue int8) int8 {
	return _Kv3GetMemberInt8(kv, name, defaultValue)
}

var _Kv3GetMemberUInt8 = func(kv uintptr, name string, defaultValue uint8) uint8 {
	var __retVal uint8
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint8_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint8(C.Kv3GetMemberUInt8(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberUInt8 
//  @brief Gets an unsigned 8-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint8_t value or defaultValue
func Kv3GetMemberUInt8(kv uintptr, name string, defaultValue uint8) uint8 {
	return _Kv3GetMemberUInt8(kv, name, defaultValue)
}

var _Kv3GetMemberShort = func(kv uintptr, name string, defaultValue int16) int16 {
	var __retVal int16
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int16_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int16(C.Kv3GetMemberShort(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberShort 
//  @brief Gets a signed 16-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int16_t value or defaultValue
func Kv3GetMemberShort(kv uintptr, name string, defaultValue int16) int16 {
	return _Kv3GetMemberShort(kv, name, defaultValue)
}

var _Kv3GetMemberUShort = func(kv uintptr, name string, defaultValue uint16) uint16 {
	var __retVal uint16
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint16_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint16(C.Kv3GetMemberUShort(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberUShort 
//  @brief Gets an unsigned 16-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint16_t value or defaultValue
func Kv3GetMemberUShort(kv uintptr, name string, defaultValue uint16) uint16 {
	return _Kv3GetMemberUShort(kv, name, defaultValue)
}

var _Kv3GetMemberInt = func(kv uintptr, name string, defaultValue int32) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int32_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.Kv3GetMemberInt(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberInt 
//  @brief Gets a signed 32-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int32_t value or defaultValue
func Kv3GetMemberInt(kv uintptr, name string, defaultValue int32) int32 {
	return _Kv3GetMemberInt(kv, name, defaultValue)
}

var _Kv3GetMemberUInt = func(kv uintptr, name string, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint32_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint32(C.Kv3GetMemberUInt(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberUInt 
//  @brief Gets an unsigned 32-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint32_t value or defaultValue
func Kv3GetMemberUInt(kv uintptr, name string, defaultValue uint32) uint32 {
	return _Kv3GetMemberUInt(kv, name, defaultValue)
}

var _Kv3GetMemberInt64 = func(kv uintptr, name string, defaultValue int64) int64 {
	var __retVal int64
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int64_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.Kv3GetMemberInt64(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberInt64 
//  @brief Gets a signed 64-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int64_t value or defaultValue
func Kv3GetMemberInt64(kv uintptr, name string, defaultValue int64) int64 {
	return _Kv3GetMemberInt64(kv, name, defaultValue)
}

var _Kv3GetMemberUInt64 = func(kv uintptr, name string, defaultValue uint64) uint64 {
	var __retVal uint64
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint64_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint64(C.Kv3GetMemberUInt64(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberUInt64 
//  @brief Gets an unsigned 64-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint64_t value or defaultValue
func Kv3GetMemberUInt64(kv uintptr, name string, defaultValue uint64) uint64 {
	return _Kv3GetMemberUInt64(kv, name, defaultValue)
}

var _Kv3GetMemberFloat = func(kv uintptr, name string, defaultValue float32) float32 {
	var __retVal float32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.float(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = float32(C.Kv3GetMemberFloat(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberFloat 
//  @brief Gets a float value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Float value or defaultValue
func Kv3GetMemberFloat(kv uintptr, name string, defaultValue float32) float32 {
	return _Kv3GetMemberFloat(kv, name, defaultValue)
}

var _Kv3GetMemberDouble = func(kv uintptr, name string, defaultValue float64) float64 {
	var __retVal float64
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.double(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = float64(C.Kv3GetMemberDouble(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberDouble 
//  @brief Gets a double value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Double value or defaultValue
func Kv3GetMemberDouble(kv uintptr, name string, defaultValue float64) float64 {
	return _Kv3GetMemberDouble(kv, name, defaultValue)
}

var _Kv3GetMemberPointer = func(kv uintptr, name string, defaultValue uintptr) uintptr {
	var __retVal uintptr
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uintptr_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.Kv3GetMemberPointer(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberPointer 
//  @brief Gets a pointer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Pointer value as uintptr_t or defaultValue
func Kv3GetMemberPointer(kv uintptr, name string, defaultValue uintptr) uintptr {
	return _Kv3GetMemberPointer(kv, name, defaultValue)
}

var _Kv3GetMemberStringToken = func(kv uintptr, name string, defaultValue uint32) uint32 {
	var __retVal uint32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint32_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = uint32(C.Kv3GetMemberStringToken(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberStringToken 
//  @brief Gets a string token value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default token value to return if member not found
//
//  @return String token hash code or defaultValue
func Kv3GetMemberStringToken(kv uintptr, name string, defaultValue uint32) uint32 {
	return _Kv3GetMemberStringToken(kv, name, defaultValue)
}

var _Kv3GetMemberEHandle = func(kv uintptr, name string, defaultValue int32) int32 {
	var __retVal int32
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := C.int32_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.Kv3GetMemberEHandle(__kv, (*C.String)(unsafe.Pointer(&__name)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberEHandle 
//  @brief Gets an entity handle value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default entity handle value to return if member not found
//
//  @return Entity handle as int32_t or defaultValue
func Kv3GetMemberEHandle(kv uintptr, name string, defaultValue int32) int32 {
	return _Kv3GetMemberEHandle(kv, name, defaultValue)
}

var _Kv3GetMemberString = func(kv uintptr, name string, defaultValue string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := plugify.ConstructString(defaultValue)
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberString(__kv, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__defaultValue)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__defaultValue)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberString 
//  @brief Gets a string value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default string to return if member not found
//
//  @return String value or defaultValue
func Kv3GetMemberString(kv uintptr, name string, defaultValue string) string {
	return _Kv3GetMemberString(kv, name, defaultValue)
}

var _Kv3GetMemberColor = func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberColor(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberColor 
//  @brief Gets a color value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default color value to return if member not found
//
//  @return Color value as vec4 or defaultValue
func Kv3GetMemberColor(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetMemberColor(kv, name, defaultValue)
}

var _Kv3GetMemberVector = func(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberVector(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberVector 
//  @brief Gets a 3D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default vector to return if member not found
//
//  @return 3D vector or defaultValue
func Kv3GetMemberVector(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3 {
	return _Kv3GetMemberVector(kv, name, defaultValue)
}

var _Kv3GetMemberVector2D = func(kv uintptr, name string, defaultValue plugify.Vector2) plugify.Vector2 {
	var __retVal plugify.Vector2
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector2)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberVector2D(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberVector2D 
//  @brief Gets a 2D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default 2D vector to return if member not found
//
//  @return 2D vector or defaultValue
func Kv3GetMemberVector2D(kv uintptr, name string, defaultValue plugify.Vector2) plugify.Vector2 {
	return _Kv3GetMemberVector2D(kv, name, defaultValue)
}

var _Kv3GetMemberVector4D = func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberVector4D(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberVector4D 
//  @brief Gets a 4D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default 4D vector to return if member not found
//
//  @return 4D vector or defaultValue
func Kv3GetMemberVector4D(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetMemberVector4D(kv, name, defaultValue)
}

var _Kv3GetMemberQuaternion = func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	var __retVal plugify.Vector4
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberQuaternion(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberQuaternion 
//  @brief Gets a quaternion value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default quaternion to return if member not found
//
//  @return Quaternion as vec4 or defaultValue
func Kv3GetMemberQuaternion(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4 {
	return _Kv3GetMemberQuaternion(kv, name, defaultValue)
}

var _Kv3GetMemberQAngle = func(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3 {
	var __retVal plugify.Vector3
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberQAngle(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberQAngle 
//  @brief Gets an angle (QAngle) value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default angle to return if member not found
//
//  @return QAngle as vec3 or defaultValue
func Kv3GetMemberQAngle(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3 {
	return _Kv3GetMemberQAngle(kv, name, defaultValue)
}

var _Kv3GetMemberMatrix3x4 = func(kv uintptr, name string, defaultValue plugify.Matrix4x4) plugify.Matrix4x4 {
	var __retVal plugify.Matrix4x4
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Matrix4x4)(unsafe.Pointer(&defaultValue))
	plugify.Block {
		Try: func() {
			__native := C.Kv3GetMemberMatrix3x4(__kv, (*C.String)(unsafe.Pointer(&__name)), &__defaultValue)
			__retVal = *(*plugify.Matrix4x4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// Kv3GetMemberMatrix3x4 
//  @brief Gets a 3x4 matrix value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default matrix to return if member not found
//
//  @return 3x4 matrix as mat4x4 or defaultValue
func Kv3GetMemberMatrix3x4(kv uintptr, name string, defaultValue plugify.Matrix4x4) plugify.Matrix4x4 {
	return _Kv3GetMemberMatrix3x4(kv, name, defaultValue)
}

var _Kv3SetMemberToNull = func(kv uintptr, name string) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToNull(__kv, (*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberToNull 
//  @brief Sets a table member to null
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func Kv3SetMemberToNull(kv uintptr, name string) {
	_Kv3SetMemberToNull(kv, name)
}

var _Kv3SetMemberToEmptyArray = func(kv uintptr, name string) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToEmptyArray(__kv, (*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberToEmptyArray 
//  @brief Sets a table member to an empty array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func Kv3SetMemberToEmptyArray(kv uintptr, name string) {
	_Kv3SetMemberToEmptyArray(kv, name)
}

var _Kv3SetMemberToEmptyTable = func(kv uintptr, name string) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToEmptyTable(__kv, (*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberToEmptyTable 
//  @brief Sets a table member to an empty table
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func Kv3SetMemberToEmptyTable(kv uintptr, name string) {
	_Kv3SetMemberToEmptyTable(kv, name)
}

var _Kv3SetMemberToBinaryBlob = func(kv uintptr, name string, blob []uint8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__blob := plugify.ConstructVectorUInt8(blob)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToBinaryBlob(__kv, (*C.String)(unsafe.Pointer(&__name)), (*C.Vector)(unsafe.Pointer(&__blob)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVectorUInt8(&__blob)
		},
	}.Do()
}

// Kv3SetMemberToBinaryBlob 
//  @brief Sets a table member to a binary blob (copies the data)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param blob: Vector containing the binary blob data
func Kv3SetMemberToBinaryBlob(kv uintptr, name string, blob []uint8) {
	_Kv3SetMemberToBinaryBlob(kv, name, blob)
}

var _Kv3SetMemberToBinaryBlobExternal = func(kv uintptr, name string, blob []uint8, free_mem bool) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__blob := plugify.ConstructVectorUInt8(blob)
	__free_mem := C.bool(free_mem)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToBinaryBlobExternal(__kv, (*C.String)(unsafe.Pointer(&__name)), (*C.Vector)(unsafe.Pointer(&__blob)), __free_mem)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVectorUInt8(&__blob)
		},
	}.Do()
}

// Kv3SetMemberToBinaryBlobExternal 
//  @brief Sets a table member to an external binary blob (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param blob: Vector containing the external binary blob data
//  @param free_mem: Whether to free the memory when the object is destroyed
func Kv3SetMemberToBinaryBlobExternal(kv uintptr, name string, blob []uint8, free_mem bool) {
	_Kv3SetMemberToBinaryBlobExternal(kv, name, blob, free_mem)
}

var _Kv3SetMemberToCopyOfValue = func(kv uintptr, name string, other uintptr) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__other := C.uintptr_t(other)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberToCopyOfValue(__kv, (*C.String)(unsafe.Pointer(&__name)), __other)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberToCopyOfValue 
//  @brief Sets a table member to a copy of another KeyValues3 value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param other: Pointer to the KeyValues3 object to copy
func Kv3SetMemberToCopyOfValue(kv uintptr, name string, other uintptr) {
	_Kv3SetMemberToCopyOfValue(kv, name, other)
}

var _Kv3SetMemberBool = func(kv uintptr, name string, value bool) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.bool(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberBool(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberBool 
//  @brief Sets a table member to a boolean value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Boolean value to set
func Kv3SetMemberBool(kv uintptr, name string, value bool) {
	_Kv3SetMemberBool(kv, name, value)
}

var _Kv3SetMemberChar = func(kv uintptr, name string, value int8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.int8_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberChar(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberChar 
//  @brief Sets a table member to a char value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Char value to set
func Kv3SetMemberChar(kv uintptr, name string, value int8) {
	_Kv3SetMemberChar(kv, name, value)
}

var _Kv3SetMemberUChar32 = func(kv uintptr, name string, value uint32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.uint32_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberUChar32(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberUChar32 
//  @brief Sets a table member to a 32-bit Unicode character value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: 32-bit Unicode character value to set
func Kv3SetMemberUChar32(kv uintptr, name string, value uint32) {
	_Kv3SetMemberUChar32(kv, name, value)
}

var _Kv3SetMemberInt8 = func(kv uintptr, name string, value int8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.int8_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberInt8(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberInt8 
//  @brief Sets a table member to a signed 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int8_t value to set
func Kv3SetMemberInt8(kv uintptr, name string, value int8) {
	_Kv3SetMemberInt8(kv, name, value)
}

var _Kv3SetMemberUInt8 = func(kv uintptr, name string, value uint8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.uint8_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberUInt8(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberUInt8 
//  @brief Sets a table member to an unsigned 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint8_t value to set
func Kv3SetMemberUInt8(kv uintptr, name string, value uint8) {
	_Kv3SetMemberUInt8(kv, name, value)
}

var _Kv3SetMemberShort = func(kv uintptr, name string, value int16) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.int16_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberShort(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberShort 
//  @brief Sets a table member to a signed 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int16_t value to set
func Kv3SetMemberShort(kv uintptr, name string, value int16) {
	_Kv3SetMemberShort(kv, name, value)
}

var _Kv3SetMemberUShort = func(kv uintptr, name string, value uint16) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.uint16_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberUShort(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberUShort 
//  @brief Sets a table member to an unsigned 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint16_t value to set
func Kv3SetMemberUShort(kv uintptr, name string, value uint16) {
	_Kv3SetMemberUShort(kv, name, value)
}

var _Kv3SetMemberInt = func(kv uintptr, name string, value int32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.int32_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberInt(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberInt 
//  @brief Sets a table member to a signed 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int32_t value to set
func Kv3SetMemberInt(kv uintptr, name string, value int32) {
	_Kv3SetMemberInt(kv, name, value)
}

var _Kv3SetMemberUInt = func(kv uintptr, name string, value uint32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.uint32_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberUInt(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberUInt 
//  @brief Sets a table member to an unsigned 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint32_t value to set
func Kv3SetMemberUInt(kv uintptr, name string, value uint32) {
	_Kv3SetMemberUInt(kv, name, value)
}

var _Kv3SetMemberInt64 = func(kv uintptr, name string, value int64) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.int64_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberInt64(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberInt64 
//  @brief Sets a table member to a signed 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int64_t value to set
func Kv3SetMemberInt64(kv uintptr, name string, value int64) {
	_Kv3SetMemberInt64(kv, name, value)
}

var _Kv3SetMemberUInt64 = func(kv uintptr, name string, value uint64) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.uint64_t(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberUInt64(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberUInt64 
//  @brief Sets a table member to an unsigned 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint64_t value to set
func Kv3SetMemberUInt64(kv uintptr, name string, value uint64) {
	_Kv3SetMemberUInt64(kv, name, value)
}

var _Kv3SetMemberFloat = func(kv uintptr, name string, value float32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.float(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberFloat(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberFloat 
//  @brief Sets a table member to a float value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Float value to set
func Kv3SetMemberFloat(kv uintptr, name string, value float32) {
	_Kv3SetMemberFloat(kv, name, value)
}

var _Kv3SetMemberDouble = func(kv uintptr, name string, value float64) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__value := C.double(value)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberDouble(__kv, (*C.String)(unsafe.Pointer(&__name)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberDouble 
//  @brief Sets a table member to a double value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Double value to set
func Kv3SetMemberDouble(kv uintptr, name string, value float64) {
	_Kv3SetMemberDouble(kv, name, value)
}

var _Kv3SetMemberPointer = func(kv uintptr, name string, ptr uintptr) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__ptr := C.uintptr_t(ptr)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberPointer(__kv, (*C.String)(unsafe.Pointer(&__name)), __ptr)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberPointer 
//  @brief Sets a table member to a pointer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ptr: Pointer value as uintptr_t to set
func Kv3SetMemberPointer(kv uintptr, name string, ptr uintptr) {
	_Kv3SetMemberPointer(kv, name, ptr)
}

var _Kv3SetMemberStringToken = func(kv uintptr, name string, token uint32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__token := C.uint32_t(token)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberStringToken(__kv, (*C.String)(unsafe.Pointer(&__name)), __token)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberStringToken 
//  @brief Sets a table member to a string token value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param token: String token hash code to set
func Kv3SetMemberStringToken(kv uintptr, name string, token uint32) {
	_Kv3SetMemberStringToken(kv, name, token)
}

var _Kv3SetMemberEHandle = func(kv uintptr, name string, ehandle int32) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__ehandle := C.int32_t(ehandle)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberEHandle(__kv, (*C.String)(unsafe.Pointer(&__name)), __ehandle)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberEHandle 
//  @brief Sets a table member to an entity handle value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ehandle: Entity handle value to set
func Kv3SetMemberEHandle(kv uintptr, name string, ehandle int32) {
	_Kv3SetMemberEHandle(kv, name, ehandle)
}

var _Kv3SetMemberString = func(kv uintptr, name string, str string, subtype uint8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__str := plugify.ConstructString(str)
	__subtype := C.uint8_t(subtype)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberString(__kv, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__str)), __subtype)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__str)
		},
	}.Do()
}

// Kv3SetMemberString 
//  @brief Sets a table member to a string value (copies the string)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param str: String value to set
//  @param subtype: String subtype enumeration value
func Kv3SetMemberString(kv uintptr, name string, str string, subtype uint8) {
	_Kv3SetMemberString(kv, name, str, subtype)
}

var _Kv3SetMemberStringExternal = func(kv uintptr, name string, str string, subtype uint8) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__str := plugify.ConstructString(str)
	__subtype := C.uint8_t(subtype)
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberStringExternal(__kv, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__str)), __subtype)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__str)
		},
	}.Do()
}

// Kv3SetMemberStringExternal 
//  @brief Sets a table member to an external string value (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param str: External string value to reference
//  @param subtype: String subtype enumeration value
func Kv3SetMemberStringExternal(kv uintptr, name string, str string, subtype uint8) {
	_Kv3SetMemberStringExternal(kv, name, str, subtype)
}

var _Kv3SetMemberColor = func(kv uintptr, name string, color plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberColor(__kv, (*C.String)(unsafe.Pointer(&__name)), &__color)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberColor 
//  @brief Sets a table member to a color value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param color: Color value as vec4 to set
func Kv3SetMemberColor(kv uintptr, name string, color plugify.Vector4) {
	_Kv3SetMemberColor(kv, name, color)
}

var _Kv3SetMemberVector = func(kv uintptr, name string, vec plugify.Vector3) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__vec := *(*C.Vector3)(unsafe.Pointer(&vec))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberVector(__kv, (*C.String)(unsafe.Pointer(&__name)), &__vec)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberVector 
//  @brief Sets a table member to a 3D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec: 3D vector to set
func Kv3SetMemberVector(kv uintptr, name string, vec plugify.Vector3) {
	_Kv3SetMemberVector(kv, name, vec)
}

var _Kv3SetMemberVector2D = func(kv uintptr, name string, vec2d plugify.Vector2) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__vec2d := *(*C.Vector2)(unsafe.Pointer(&vec2d))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberVector2D(__kv, (*C.String)(unsafe.Pointer(&__name)), &__vec2d)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberVector2D 
//  @brief Sets a table member to a 2D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec2d: 2D vector to set
func Kv3SetMemberVector2D(kv uintptr, name string, vec2d plugify.Vector2) {
	_Kv3SetMemberVector2D(kv, name, vec2d)
}

var _Kv3SetMemberVector4D = func(kv uintptr, name string, vec4d plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__vec4d := *(*C.Vector4)(unsafe.Pointer(&vec4d))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberVector4D(__kv, (*C.String)(unsafe.Pointer(&__name)), &__vec4d)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberVector4D 
//  @brief Sets a table member to a 4D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec4d: 4D vector to set
func Kv3SetMemberVector4D(kv uintptr, name string, vec4d plugify.Vector4) {
	_Kv3SetMemberVector4D(kv, name, vec4d)
}

var _Kv3SetMemberQuaternion = func(kv uintptr, name string, quat plugify.Vector4) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__quat := *(*C.Vector4)(unsafe.Pointer(&quat))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberQuaternion(__kv, (*C.String)(unsafe.Pointer(&__name)), &__quat)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberQuaternion 
//  @brief Sets a table member to a quaternion value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param quat: Quaternion to set (as vec4)
func Kv3SetMemberQuaternion(kv uintptr, name string, quat plugify.Vector4) {
	_Kv3SetMemberQuaternion(kv, name, quat)
}

var _Kv3SetMemberQAngle = func(kv uintptr, name string, ang plugify.Vector3) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__ang := *(*C.Vector3)(unsafe.Pointer(&ang))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberQAngle(__kv, (*C.String)(unsafe.Pointer(&__name)), &__ang)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberQAngle 
//  @brief Sets a table member to an angle (QAngle) value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ang: QAngle to set (as vec3)
func Kv3SetMemberQAngle(kv uintptr, name string, ang plugify.Vector3) {
	_Kv3SetMemberQAngle(kv, name, ang)
}

var _Kv3SetMemberMatrix3x4 = func(kv uintptr, name string, matrix plugify.Matrix4x4) {
	__kv := C.uintptr_t(kv)
	__name := plugify.ConstructString(name)
	__matrix := *(*C.Matrix4x4)(unsafe.Pointer(&matrix))
	plugify.Block {
		Try: func() {
			C.Kv3SetMemberMatrix3x4(__kv, (*C.String)(unsafe.Pointer(&__name)), &__matrix)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// Kv3SetMemberMatrix3x4 
//  @brief Sets a table member to a 3x4 matrix value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param matrix: 3x4 matrix to set (as mat4x4)
func Kv3SetMemberMatrix3x4(kv uintptr, name string, matrix plugify.Matrix4x4) {
	_Kv3SetMemberMatrix3x4(kv, name, matrix)
}

var _Kv3DebugPrint = func(kv uintptr) {
	__kv := C.uintptr_t(kv)
	C.Kv3DebugPrint(__kv)
}

// Kv3DebugPrint 
//  @brief Prints debug information about the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
func Kv3DebugPrint(kv uintptr) {
	_Kv3DebugPrint(kv)
}

var _Kv3LoadFromBuffer = func(context uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	var __retVal bool
	__context := C.uintptr_t(context)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructVectorUInt8(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromBuffer(__context, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromBuffer 
//  @brief Loads KeyValues3 data from a buffer into a context
//
//  @param context: Pointer to the KeyValues3 context
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromBuffer(context uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	return _Kv3LoadFromBuffer(context, error_, input, kv_name, flags)
}

var _Kv3Load = func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructVectorUInt8(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3Load(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3Load 
//  @brief Loads KeyValues3 data from a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3Load(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	return _Kv3Load(kv, error_, input, kv_name, flags)
}

var _Kv3LoadFromText = func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructString(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromText(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromText 
//  @brief Loads KeyValues3 data from a text string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Text string containing KV3 data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromText(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	return _Kv3LoadFromText(kv, error_, input, kv_name, flags)
}

var _Kv3LoadFromFileToContext = func(context uintptr, error_ *string, filename string, path string, flags uint32) bool {
	var __retVal bool
	__context := C.uintptr_t(context)
	__error_ := plugify.ConstructString(*error_)
	__filename := plugify.ConstructString(filename)
	__path := plugify.ConstructString(path)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromFileToContext(__context, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__filename)), (*C.String)(unsafe.Pointer(&__path)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__filename)
			plugify.DestroyString(&__path)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromFileToContext 
//  @brief Loads KeyValues3 data from a file into a context
//
//  @param context: Pointer to the KeyValues3 context
//  @param error_: Output string for error messages
//  @param filename: Name of the file to load
//  @param path: Path to the file
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromFileToContext(context uintptr, error_ *string, filename string, path string, flags uint32) bool {
	return _Kv3LoadFromFileToContext(context, error_, filename, path, flags)
}

var _Kv3LoadFromFile = func(kv uintptr, error_ *string, filename string, path string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__filename := plugify.ConstructString(filename)
	__path := plugify.ConstructString(path)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromFile(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__filename)), (*C.String)(unsafe.Pointer(&__path)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__filename)
			plugify.DestroyString(&__path)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromFile 
//  @brief Loads KeyValues3 data from a file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param filename: Name of the file to load
//  @param path: Path to the file
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromFile(kv uintptr, error_ *string, filename string, path string, flags uint32) bool {
	return _Kv3LoadFromFile(kv, error_, filename, path, flags)
}

var _Kv3LoadFromJSON = func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructString(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromJSON(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromJSON 
//  @brief Loads KeyValues3 data from a JSON string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: JSON string
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromJSON(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	return _Kv3LoadFromJSON(kv, error_, input, kv_name, flags)
}

var _Kv3LoadFromJSONFile = func(kv uintptr, error_ *string, path string, filename string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__path := plugify.ConstructString(path)
	__filename := plugify.ConstructString(filename)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromJSONFile(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__path)), (*C.String)(unsafe.Pointer(&__filename)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__path)
			plugify.DestroyString(&__filename)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromJSONFile 
//  @brief Loads KeyValues3 data from a JSON file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param path: Path to the file
//  @param filename: Name of the file to load
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromJSONFile(kv uintptr, error_ *string, path string, filename string, flags uint32) bool {
	return _Kv3LoadFromJSONFile(kv, error_, path, filename, flags)
}

var _Kv3LoadFromKV1File = func(kv uintptr, error_ *string, path string, filename string, esc_behavior uint8, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__path := plugify.ConstructString(path)
	__filename := plugify.ConstructString(filename)
	__esc_behavior := C.uint8_t(esc_behavior)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromKV1File(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__path)), (*C.String)(unsafe.Pointer(&__filename)), __esc_behavior, __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__path)
			plugify.DestroyString(&__filename)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromKV1File 
//  @brief Loads KeyValues3 data from a KeyValues1 file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param path: Path to the file
//  @param filename: Name of the file to load
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromKV1File(kv uintptr, error_ *string, path string, filename string, esc_behavior uint8, flags uint32) bool {
	return _Kv3LoadFromKV1File(kv, error_, path, filename, esc_behavior, flags)
}

var _Kv3LoadFromKV1Text = func(kv uintptr, error_ *string, input string, esc_behavior uint8, kv_name string, unk bool, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructString(input)
	__esc_behavior := C.uint8_t(esc_behavior)
	__kv_name := plugify.ConstructString(kv_name)
	__unk := C.bool(unk)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromKV1Text(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__input)), __esc_behavior, (*C.String)(unsafe.Pointer(&__kv_name)), __unk, __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromKV1Text 
//  @brief Loads KeyValues3 data from a KeyValues1 text string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: KV1 text string
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param kv_name: Name for the KeyValues3 object
//  @param unk: Unknown boolean parameter
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromKV1Text(kv uintptr, error_ *string, input string, esc_behavior uint8, kv_name string, unk bool, flags uint32) bool {
	return _Kv3LoadFromKV1Text(kv, error_, input, esc_behavior, kv_name, unk, flags)
}

var _Kv3LoadFromKV1TextTranslated = func(kv uintptr, error_ *string, input string, esc_behavior uint8, translation uintptr, unk1 int32, kv_name string, unk2 bool, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructString(input)
	__esc_behavior := C.uint8_t(esc_behavior)
	__translation := C.uintptr_t(translation)
	__unk1 := C.int32_t(unk1)
	__kv_name := plugify.ConstructString(kv_name)
	__unk2 := C.bool(unk2)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromKV1TextTranslated(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__input)), __esc_behavior, __translation, __unk1, (*C.String)(unsafe.Pointer(&__kv_name)), __unk2, __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromKV1TextTranslated 
//  @brief Loads KeyValues3 data from a KeyValues1 text string with translation
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: KV1 text string
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param translation: Pointer to translation table
//  @param unk1: Unknown integer parameter
//  @param kv_name: Name for the KeyValues3 object
//  @param unk2: Unknown boolean parameter
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromKV1TextTranslated(kv uintptr, error_ *string, input string, esc_behavior uint8, translation uintptr, unk1 int32, kv_name string, unk2 bool, flags uint32) bool {
	return _Kv3LoadFromKV1TextTranslated(kv, error_, input, esc_behavior, translation, unk1, kv_name, unk2, flags)
}

var _Kv3LoadFromKV3OrKV1 = func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructVectorUInt8(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromKV3OrKV1(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromKV3OrKV1 
//  @brief Loads data from a buffer that may be KV3 or KV1 format
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromKV3OrKV1(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	return _Kv3LoadFromKV3OrKV1(kv, error_, input, kv_name, flags)
}

var _Kv3LoadFromOldSchemaText = func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructVectorUInt8(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadFromOldSchemaText(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadFromOldSchemaText 
//  @brief Loads KeyValues3 data from old schema text format
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadFromOldSchemaText(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool {
	return _Kv3LoadFromOldSchemaText(kv, error_, input, kv_name, flags)
}

var _Kv3LoadTextNoHeader = func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__input := plugify.ConstructString(input)
	__kv_name := plugify.ConstructString(kv_name)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3LoadTextNoHeader(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__input)), (*C.String)(unsafe.Pointer(&__kv_name)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__input)
			plugify.DestroyString(&__kv_name)
		},
	}.Do()
	return __retVal
}

// Kv3LoadTextNoHeader 
//  @brief Loads KeyValues3 text without a header
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Text string containing KV3 data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func Kv3LoadTextNoHeader(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool {
	return _Kv3LoadTextNoHeader(kv, error_, input, kv_name, flags)
}

var _Kv3Save = func(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructVectorUInt8(*output)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3Save(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__output)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			plugify.GetVectorDataUInt8To(&__output, output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3Save 
//  @brief Saves KeyValues3 data to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output buffer data
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func Kv3Save(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool {
	return _Kv3Save(kv, error_, output, flags)
}

var _Kv3SaveAsJSON = func(kv uintptr, error_ *string, output *[]uint8) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructVectorUInt8(*output)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveAsJSON(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__output))))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			plugify.GetVectorDataUInt8To(&__output, output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveAsJSON 
//  @brief Saves KeyValues3 data as JSON to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output JSON data
//
//  @return true if successful, false otherwise
func Kv3SaveAsJSON(kv uintptr, error_ *string, output *[]uint8) bool {
	return _Kv3SaveAsJSON(kv, error_, output)
}

var _Kv3SaveAsJSONString = func(kv uintptr, error_ *string, output *string) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructString(*output)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveAsJSONString(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__output))))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			*output = plugify.GetStringData[string](&__output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveAsJSONString 
//  @brief Saves KeyValues3 data as a JSON string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the JSON output
//
//  @return true if successful, false otherwise
func Kv3SaveAsJSONString(kv uintptr, error_ *string, output *string) bool {
	return _Kv3SaveAsJSONString(kv, error_, output)
}

var _Kv3SaveAsKV1Text = func(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructVectorUInt8(*output)
	__esc_behavior := C.uint8_t(esc_behavior)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveAsKV1Text(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__output)), __esc_behavior))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			plugify.GetVectorDataUInt8To(&__output, output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveAsKV1Text 
//  @brief Saves KeyValues3 data as KeyValues1 text to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output KV1 text data
//  @param esc_behavior: Escape sequence behavior for KV1 text
//
//  @return true if successful, false otherwise
func Kv3SaveAsKV1Text(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8) bool {
	return _Kv3SaveAsKV1Text(kv, error_, output, esc_behavior)
}

var _Kv3SaveAsKV1TextTranslated = func(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8, translation uintptr, unk int32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructVectorUInt8(*output)
	__esc_behavior := C.uint8_t(esc_behavior)
	__translation := C.uintptr_t(translation)
	__unk := C.int32_t(unk)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveAsKV1TextTranslated(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__output)), __esc_behavior, __translation, __unk))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			plugify.GetVectorDataUInt8To(&__output, output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveAsKV1TextTranslated 
//  @brief Saves KeyValues3 data as KeyValues1 text with translation to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output KV1 text data
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param translation: Pointer to translation table
//  @param unk: Unknown integer parameter
//
//  @return true if successful, false otherwise
func Kv3SaveAsKV1TextTranslated(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8, translation uintptr, unk int32) bool {
	return _Kv3SaveAsKV1TextTranslated(kv, error_, output, esc_behavior, translation, unk)
}

var _Kv3SaveTextNoHeaderToBuffer = func(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructVectorUInt8(*output)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveTextNoHeaderToBuffer(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.Vector)(unsafe.Pointer(&__output)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			plugify.GetVectorDataUInt8To(&__output, output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyVectorUInt8(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveTextNoHeaderToBuffer 
//  @brief Saves KeyValues3 text without a header to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output text data
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func Kv3SaveTextNoHeaderToBuffer(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool {
	return _Kv3SaveTextNoHeaderToBuffer(kv, error_, output, flags)
}

var _Kv3SaveTextNoHeader = func(kv uintptr, error_ *string, output *string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructString(*output)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveTextNoHeader(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__output)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			*output = plugify.GetStringData[string](&__output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveTextNoHeader 
//  @brief Saves KeyValues3 text without a header to a string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the text output
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func Kv3SaveTextNoHeader(kv uintptr, error_ *string, output *string, flags uint32) bool {
	return _Kv3SaveTextNoHeader(kv, error_, output, flags)
}

var _Kv3SaveTextToString = func(kv uintptr, error_ *string, output *string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__output := plugify.ConstructString(*output)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveTextToString(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__output)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
			*output = plugify.GetStringData[string](&__output)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__output)
		},
	}.Do()
	return __retVal
}

// Kv3SaveTextToString 
//  @brief Saves KeyValues3 text to a string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the text output
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func Kv3SaveTextToString(kv uintptr, error_ *string, output *string, flags uint32) bool {
	return _Kv3SaveTextToString(kv, error_, output, flags)
}

var _Kv3SaveToFile = func(kv uintptr, error_ *string, filename string, path string, flags uint32) bool {
	var __retVal bool
	__kv := C.uintptr_t(kv)
	__error_ := plugify.ConstructString(*error_)
	__filename := plugify.ConstructString(filename)
	__path := plugify.ConstructString(path)
	__flags := C.uint32_t(flags)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.Kv3SaveToFile(__kv, (*C.String)(unsafe.Pointer(&__error_)), (*C.String)(unsafe.Pointer(&__filename)), (*C.String)(unsafe.Pointer(&__path)), __flags))
			// Unmarshal - Convert native data to managed data.
			*error_ = plugify.GetStringData[string](&__error_)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
			plugify.DestroyString(&__filename)
			plugify.DestroyString(&__path)
		},
	}.Do()
	return __retVal
}

// Kv3SaveToFile 
//  @brief Saves KeyValues3 data to a file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param filename: Name of the file to save
//  @param path: Path to save the file
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func Kv3SaveToFile(kv uintptr, error_ *string, filename string, path string, flags uint32) bool {
	return _Kv3SaveToFile(kv, error_, filename, path, flags)
}

var (
	KeyValues3ErrEmptyHandle = errors.New("KeyValues3: empty handle")
)

//  @brief RAII wrapper for KeyValues3 handle.
//
type KeyValues3 struct {
	handle    uintptr
	cleanup   runtime.Cleanup
	ownership ownership
	noCopy    noCopy
}

// NewKeyValues3Kv3Create 
//  @brief Creates a new KeyValues3 object with specified type and subtype
//
//  @param type_: The KV3 type enumeration value
//  @param subtype: The KV3 subtype enumeration value
func NewKeyValues3Kv3Create(type_ int32, subtype int32) *KeyValues3 {
	return NewKeyValues3Owned(Kv3Create(type_, subtype))
}

// NewKeyValues3Kv3CreateWithCluster 
//  @brief Creates a new KeyValues3 object with cluster element, type, and subtype
//
//  @param cluster_elem: The cluster element index
//  @param type_: The KV3 type enumeration value
//  @param subtype: The KV3 subtype enumeration value
func NewKeyValues3Kv3CreateWithCluster(cluster_elem int32, type_ int32, subtype int32) *KeyValues3 {
	return NewKeyValues3Owned(Kv3CreateWithCluster(cluster_elem, type_, subtype))
}

// NewKeyValues3Kv3CreateCopy 
//  @brief Creates a copy of an existing KeyValues3 object
//
//  @param other: Pointer to the KeyValues3 object to copy
func NewKeyValues3Kv3CreateCopy(other uintptr) *KeyValues3 {
	return NewKeyValues3Owned(Kv3CreateCopy(other))
}

// NewKeyValues3Borrowed creates a KeyValues3 from a borrowed handle
func NewKeyValues3Borrowed(handle uintptr) *KeyValues3 {
	if handle == 0 {
		return &KeyValues3{}
	}
	return &KeyValues3{
		handle:    handle,
		ownership: Borrowed,
	}
}

// NewKeyValues3Owned creates a KeyValues3 from an owned handle
func NewKeyValues3Owned(handle uintptr) *KeyValues3 {
	if handle == 0 {
		return &KeyValues3{}
	}
	w := &KeyValues3{
		handle:    handle,
		ownership: Owned,
	}
	w.cleanup = runtime.AddCleanup(w, destroyKeyValues3Handle, handle)
	return w
}

// destroyKeyValues3Handle destroys an owned handle.
func destroyKeyValues3Handle(handle uintptr) {
	if handle != 0 {
		Kv3Destroy(handle)
	}
}

// destroy cleans up owned handles
func (w *KeyValues3) destroy() {
	if w.handle != 0 && w.ownership == Owned {
		Kv3Destroy(w.handle)
	}
}

// nullify resets the handle
func (w *KeyValues3) nullify() {
	w.handle = 0
	w.ownership = Borrowed
}

// Close explicitly destroys the handle (like C++ destructor, but manual)
func (w *KeyValues3) Close() {
	w.Reset()
}

// Get returns the underlying handle
func (w *KeyValues3) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *KeyValues3) Release() uintptr {
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	handle := w.handle
	w.nullify()
	return handle
}

// Reset destroys and resets the handle
func (w *KeyValues3) Reset() {
	if w.ownership == Owned && w.handle != 0 {
		w.cleanup.Stop()
	}
	w.destroy()
	w.nullify()
}

// Valid returns true if handle is not nil
func (w *KeyValues3) Valid() bool {
	return w.handle != 0
}

// CopyFrom 
//  @brief Copies data from another KeyValues3 object
//
//  @param kv: Pointer to the destination KeyValues3 object
//  @param other: Pointer to the source KeyValues3 object
func (w *KeyValues3) CopyFrom(other *KeyValues3) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3CopyFrom(w.handle, other.Get())
	return nil
}

// OverlayKeysFrom 
//  @brief Overlays keys from another KeyValues3 object
//
//  @param kv: Pointer to the destination KeyValues3 object
//  @param other: Pointer to the source KeyValues3 object
//  @param depth: Whether to perform a deep overlay
func (w *KeyValues3) OverlayKeysFrom(other *KeyValues3, depth bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3OverlayKeysFrom(w.handle, other.Get(), depth)
	return nil
}

// GetContext 
//  @brief Gets the context associated with a KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Pointer to the CKV3Arena, or nullptr if kv is null
func (w *KeyValues3) GetContext() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetContext(w.handle), nil
}

// GetMetaData 
//  @brief Gets the metadata associated with a KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ppCtx: Pointer to store the context pointer
//
//  @return Pointer to the KV3MetaData_t structure, or nullptr if kv is null
func (w *KeyValues3) GetMetaData(ppCtx uintptr) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMetaData(w.handle, ppCtx), nil
}

// HasFlag 
//  @brief Checks if a specific flag is set
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flag: The flag to check
//
//  @return true if the flag is set, false otherwise
func (w *KeyValues3) HasFlag(flag uint8) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3HasFlag(w.handle, flag), nil
}

// HasAnyFlags 
//  @brief Checks if any flags are set
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if any flags are set, false otherwise
func (w *KeyValues3) HasAnyFlags() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3HasAnyFlags(w.handle), nil
}

// GetAllFlags 
//  @brief Gets all flags as a bitmask
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Bitmask of all flags, or 0 if kv is null
func (w *KeyValues3) GetAllFlags() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetAllFlags(w.handle), nil
}

// SetAllFlags 
//  @brief Sets all flags from a bitmask
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flags: Bitmask of flags to set
func (w *KeyValues3) SetAllFlags(flags uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetAllFlags(w.handle, flags)
	return nil
}

// SetFlag 
//  @brief Sets or clears a specific flag
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flag: The flag to modify
//  @param state: true to set the flag, false to clear it
func (w *KeyValues3) SetFlag(flag uint8, state bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetFlag(w.handle, flag, state)
	return nil
}

// GetType 
//  @brief Gets the basic type of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The type enumeration value, or 0 if kv is null
func (w *KeyValues3) GetType() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetType(w.handle), nil
}

// GetTypeEx 
//  @brief Gets the extended type of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The extended type enumeration value, or 0 if kv is null
func (w *KeyValues3) GetTypeEx() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetTypeEx(w.handle), nil
}

// GetSubType 
//  @brief Gets the subtype of the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return The subtype enumeration value, or 0 if kv is null
func (w *KeyValues3) GetSubType() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetSubType(w.handle), nil
}

// HasInvalidMemberNames 
//  @brief Checks if the object has invalid member names
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if invalid member names exist, false otherwise
func (w *KeyValues3) HasInvalidMemberNames() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3HasInvalidMemberNames(w.handle), nil
}

// SetHasInvalidMemberNames 
//  @brief Sets the invalid member names flag
//
//  @param kv: Pointer to the KeyValues3 object
//  @param bValue: true to mark as having invalid member names, false otherwise
func (w *KeyValues3) SetHasInvalidMemberNames(bValue bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetHasInvalidMemberNames(w.handle, bValue)
	return nil
}

// GetTypeAsString 
//  @brief Gets the type as a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return String representation of the type, or empty string if kv is null
func (w *KeyValues3) GetTypeAsString() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetTypeAsString(w.handle), nil
}

// GetSubTypeAsString 
//  @brief Gets the subtype as a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return String representation of the subtype, or empty string if kv is null
func (w *KeyValues3) GetSubTypeAsString() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetSubTypeAsString(w.handle), nil
}

// ToString 
//  @brief Converts the KeyValues3 object to a string representation
//
//  @param kv: Pointer to the KeyValues3 object
//  @param flags: Formatting flags for the string conversion
//
//  @return String representation of the object, or empty string if kv is null
func (w *KeyValues3) ToString(flags uint32) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3ToString(w.handle, flags), nil
}

// IsNull 
//  @brief Checks if the KeyValues3 object is null
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is null or the pointer is null, false otherwise
func (w *KeyValues3) IsNull() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3IsNull(w.handle), nil
}

// SetToNull 
//  @brief Sets the KeyValues3 object to null
//
//  @param kv: Pointer to the KeyValues3 object
func (w *KeyValues3) SetToNull() error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetToNull(w.handle)
	return nil
}

// IsArray 
//  @brief Checks if the KeyValues3 object is an array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is an array, false otherwise
func (w *KeyValues3) IsArray() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3IsArray(w.handle), nil
}

// IsKV3Array 
//  @brief Checks if the KeyValues3 object is a KV3 array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a KV3 array, false otherwise
func (w *KeyValues3) IsKV3Array() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3IsKV3Array(w.handle), nil
}

// IsTable 
//  @brief Checks if the KeyValues3 object is a table
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a table, false otherwise
func (w *KeyValues3) IsTable() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3IsTable(w.handle), nil
}

// IsString 
//  @brief Checks if the KeyValues3 object is a string
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return true if the object is a string, false otherwise
func (w *KeyValues3) IsString() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3IsString(w.handle), nil
}

// GetBool 
//  @brief Gets the boolean value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Boolean value or defaultValue
func (w *KeyValues3) GetBool(defaultValue bool) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetBool(w.handle, defaultValue), nil
}

// GetChar 
//  @brief Gets the char value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Char value or defaultValue
func (w *KeyValues3) GetChar(defaultValue int8) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetChar(w.handle, defaultValue), nil
}

// GetUChar32 
//  @brief Gets the 32-bit Unicode character value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return 32-bit Unicode character value or defaultValue
func (w *KeyValues3) GetUChar32(defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetUChar32(w.handle, defaultValue), nil
}

// GetInt8 
//  @brief Gets the signed 8-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int8_t value or defaultValue
func (w *KeyValues3) GetInt8(defaultValue int8) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetInt8(w.handle, defaultValue), nil
}

// GetUInt8 
//  @brief Gets the unsigned 8-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint8_t value or defaultValue
func (w *KeyValues3) GetUInt8(defaultValue uint8) (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetUInt8(w.handle, defaultValue), nil
}

// GetShort 
//  @brief Gets the signed 16-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int16_t value or defaultValue
func (w *KeyValues3) GetShort(defaultValue int16) (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetShort(w.handle, defaultValue), nil
}

// GetUShort 
//  @brief Gets the unsigned 16-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint16_t value or defaultValue
func (w *KeyValues3) GetUShort(defaultValue uint16) (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetUShort(w.handle, defaultValue), nil
}

// GetInt 
//  @brief Gets the signed 32-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int32_t value or defaultValue
func (w *KeyValues3) GetInt(defaultValue int32) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetInt(w.handle, defaultValue), nil
}

// GetUInt 
//  @brief Gets the unsigned 32-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint32_t value or defaultValue
func (w *KeyValues3) GetUInt(defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetUInt(w.handle, defaultValue), nil
}

// GetInt64 
//  @brief Gets the signed 64-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return int64_t value or defaultValue
func (w *KeyValues3) GetInt64(defaultValue int64) (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetInt64(w.handle, defaultValue), nil
}

// GetUInt64 
//  @brief Gets the unsigned 64-bit integer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return uint64_t value or defaultValue
func (w *KeyValues3) GetUInt64(defaultValue uint64) (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetUInt64(w.handle, defaultValue), nil
}

// GetFloat 
//  @brief Gets the float value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Float value or defaultValue
func (w *KeyValues3) GetFloat(defaultValue float32) (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetFloat(w.handle, defaultValue), nil
}

// GetDouble 
//  @brief Gets the double value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null or conversion fails
//
//  @return Double value or defaultValue
func (w *KeyValues3) GetDouble(defaultValue float64) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetDouble(w.handle, defaultValue), nil
}

// SetBool 
//  @brief Sets the KeyValues3 object to a boolean value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Boolean value to set
func (w *KeyValues3) SetBool(value bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetBool(w.handle, value)
	return nil
}

// SetChar 
//  @brief Sets the KeyValues3 object to a char value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Char value to set
func (w *KeyValues3) SetChar(value int8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetChar(w.handle, value)
	return nil
}

// SetUChar32 
//  @brief Sets the KeyValues3 object to a 32-bit Unicode character value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: 32-bit Unicode character value to set
func (w *KeyValues3) SetUChar32(value uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetUChar32(w.handle, value)
	return nil
}

// SetInt8 
//  @brief Sets the KeyValues3 object to a signed 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int8_t value to set
func (w *KeyValues3) SetInt8(value int8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetInt8(w.handle, value)
	return nil
}

// SetUInt8 
//  @brief Sets the KeyValues3 object to an unsigned 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint8_t value to set
func (w *KeyValues3) SetUInt8(value uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetUInt8(w.handle, value)
	return nil
}

// SetShort 
//  @brief Sets the KeyValues3 object to a signed 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int16_t value to set
func (w *KeyValues3) SetShort(value int16) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetShort(w.handle, value)
	return nil
}

// SetUShort 
//  @brief Sets the KeyValues3 object to an unsigned 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint16_t value to set
func (w *KeyValues3) SetUShort(value uint16) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetUShort(w.handle, value)
	return nil
}

// SetInt 
//  @brief Sets the KeyValues3 object to a signed 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int32_t value to set
func (w *KeyValues3) SetInt(value int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetInt(w.handle, value)
	return nil
}

// SetUInt 
//  @brief Sets the KeyValues3 object to an unsigned 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint32_t value to set
func (w *KeyValues3) SetUInt(value uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetUInt(w.handle, value)
	return nil
}

// SetInt64 
//  @brief Sets the KeyValues3 object to a signed 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: int64_t value to set
func (w *KeyValues3) SetInt64(value int64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetInt64(w.handle, value)
	return nil
}

// SetUInt64 
//  @brief Sets the KeyValues3 object to an unsigned 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: uint64_t value to set
func (w *KeyValues3) SetUInt64(value uint64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetUInt64(w.handle, value)
	return nil
}

// SetFloat 
//  @brief Sets the KeyValues3 object to a float value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Float value to set
func (w *KeyValues3) SetFloat(value float32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetFloat(w.handle, value)
	return nil
}

// SetDouble 
//  @brief Sets the KeyValues3 object to a double value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param value: Double value to set
func (w *KeyValues3) SetDouble(value float64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetDouble(w.handle, value)
	return nil
}

// GetPointer 
//  @brief Gets the pointer value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default value to return if kv is null
//
//  @return Pointer value as uintptr_t or defaultValue
func (w *KeyValues3) GetPointer(defaultValue uintptr) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetPointer(w.handle, defaultValue), nil
}

// SetPointer 
//  @brief Sets the KeyValues3 object to a pointer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ptr: Pointer value as uintptr_t to set
func (w *KeyValues3) SetPointer(ptr uintptr) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetPointer(w.handle, ptr)
	return nil
}

// GetStringToken 
//  @brief Gets the string token value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default token value to return if kv is null
//
//  @return String token hash code or defaultValue
func (w *KeyValues3) GetStringToken(defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetStringToken(w.handle, defaultValue), nil
}

// SetStringToken 
//  @brief Sets the KeyValues3 object to a string token value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param token: String token hash code to set
func (w *KeyValues3) SetStringToken(token uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetStringToken(w.handle, token)
	return nil
}

// GetEHandle 
//  @brief Gets the entity handle value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default entity handle value to return if kv is null
//
//  @return Entity handle as int32_t or defaultValue
func (w *KeyValues3) GetEHandle(defaultValue int32) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetEHandle(w.handle, defaultValue), nil
}

// SetEHandle 
//  @brief Sets the KeyValues3 object to an entity handle value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ehandle: Entity handle value to set
func (w *KeyValues3) SetEHandle(ehandle int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetEHandle(w.handle, ehandle)
	return nil
}

// GetString 
//  @brief Gets the string value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default string to return if kv is null or value is empty
//
//  @return String value or defaultValue
func (w *KeyValues3) GetString(defaultValue string) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetString(w.handle, defaultValue), nil
}

// SetString 
//  @brief Sets the KeyValues3 object to a string value (copies the string)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param str: String value to set
//  @param subtype: String subtype enumeration value
func (w *KeyValues3) SetString(str string, subtype uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetString(w.handle, str, subtype)
	return nil
}

// SetStringExternal 
//  @brief Sets the KeyValues3 object to an external string value (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param str: External string value to reference
//  @param subtype: String subtype enumeration value
func (w *KeyValues3) SetStringExternal(str string, subtype uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetStringExternal(w.handle, str, subtype)
	return nil
}

// GetBinaryBlob 
//  @brief Gets the binary blob from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Vector containing the binary blob data, or empty vector if kv is null
func (w *KeyValues3) GetBinaryBlob() ([]uint8, error) {
	if w.handle == 0 {
		var zero []uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetBinaryBlob(w.handle), nil
}

// GetBinaryBlobSize 
//  @brief Gets the size of the binary blob in the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Size of the binary blob in bytes, or 0 if kv is null
func (w *KeyValues3) GetBinaryBlobSize() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetBinaryBlobSize(w.handle), nil
}

// SetToBinaryBlob 
//  @brief Sets the KeyValues3 object to a binary blob (copies the data)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param blob: Vector containing the binary blob data
func (w *KeyValues3) SetToBinaryBlob(blob []uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetToBinaryBlob(w.handle, blob)
	return nil
}

// SetToBinaryBlobExternal 
//  @brief Sets the KeyValues3 object to an external binary blob (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param blob: Vector containing the external binary blob data
//  @param free_mem: Whether to free the memory when the object is destroyed
func (w *KeyValues3) SetToBinaryBlobExternal(blob []uint8, free_mem bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetToBinaryBlobExternal(w.handle, blob, free_mem)
	return nil
}

// GetColor 
//  @brief Gets the color value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default color value to return if kv is null
//
//  @return Color value as vec4 or defaultValue
func (w *KeyValues3) GetColor(defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetColor(w.handle, defaultValue), nil
}

// SetColor 
//  @brief Sets the KeyValues3 object to a color value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param color: Color value as vec4 to set
func (w *KeyValues3) SetColor(color plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetColor(w.handle, color)
	return nil
}

// GetVector 
//  @brief Gets the 3D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default vector to return if kv is null
//
//  @return 3D vector or defaultValue
func (w *KeyValues3) GetVector(defaultValue plugify.Vector3) (plugify.Vector3, error) {
	if w.handle == 0 {
		var zero plugify.Vector3
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetVector(w.handle, defaultValue), nil
}

// GetVector2D 
//  @brief Gets the 2D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default 2D vector to return if kv is null
//
//  @return 2D vector or defaultValue
func (w *KeyValues3) GetVector2D(defaultValue plugify.Vector2) (plugify.Vector2, error) {
	if w.handle == 0 {
		var zero plugify.Vector2
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetVector2D(w.handle, defaultValue), nil
}

// GetVector4D 
//  @brief Gets the 4D vector value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default 4D vector to return if kv is null
//
//  @return 4D vector or defaultValue
func (w *KeyValues3) GetVector4D(defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetVector4D(w.handle, defaultValue), nil
}

// GetQuaternion 
//  @brief Gets the quaternion value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default quaternion to return if kv is null
//
//  @return Quaternion as vec4 or defaultValue
func (w *KeyValues3) GetQuaternion(defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetQuaternion(w.handle, defaultValue), nil
}

// GetQAngle 
//  @brief Gets the angle (QAngle) value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default angle to return if kv is null
//
//  @return QAngle as vec3 or defaultValue
func (w *KeyValues3) GetQAngle(defaultValue plugify.Vector3) (plugify.Vector3, error) {
	if w.handle == 0 {
		var zero plugify.Vector3
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetQAngle(w.handle, defaultValue), nil
}

// GetMatrix3x4 
//  @brief Gets the 3x4 matrix value from the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
//  @param defaultValue: Default matrix to return if kv is null
//
//  @return 3x4 matrix as mat4x4 or defaultValue
func (w *KeyValues3) GetMatrix3x4(defaultValue plugify.Matrix4x4) (plugify.Matrix4x4, error) {
	if w.handle == 0 {
		var zero plugify.Matrix4x4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMatrix3x4(w.handle, defaultValue), nil
}

// SetVector 
//  @brief Sets the KeyValues3 object to a 3D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec: 3D vector to set
func (w *KeyValues3) SetVector(vec plugify.Vector3) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetVector(w.handle, vec)
	return nil
}

// SetVector2D 
//  @brief Sets the KeyValues3 object to a 2D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec2d: 2D vector to set
func (w *KeyValues3) SetVector2D(vec2d plugify.Vector2) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetVector2D(w.handle, vec2d)
	return nil
}

// SetVector4D 
//  @brief Sets the KeyValues3 object to a 4D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param vec4d: 4D vector to set
func (w *KeyValues3) SetVector4D(vec4d plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetVector4D(w.handle, vec4d)
	return nil
}

// SetQuaternion 
//  @brief Sets the KeyValues3 object to a quaternion value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param quat: Quaternion to set (as vec4)
func (w *KeyValues3) SetQuaternion(quat plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetQuaternion(w.handle, quat)
	return nil
}

// SetQAngle 
//  @brief Sets the KeyValues3 object to an angle (QAngle) value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param ang: QAngle to set (as vec3)
func (w *KeyValues3) SetQAngle(ang plugify.Vector3) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetQAngle(w.handle, ang)
	return nil
}

// SetMatrix3x4 
//  @brief Sets the KeyValues3 object to a 3x4 matrix value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param matrix: 3x4 matrix to set (as mat4x4)
func (w *KeyValues3) SetMatrix3x4(matrix plugify.Matrix4x4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMatrix3x4(w.handle, matrix)
	return nil
}

// GetArrayElementCount 
//  @brief Gets the number of elements in the array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Number of array elements, or 0 if kv is null or not an array
func (w *KeyValues3) GetArrayElementCount() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetArrayElementCount(w.handle), nil
}

// SetArrayElementCount 
//  @brief Sets the number of elements in the array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param count: Number of elements to set
//  @param type_: Type of array elements
//  @param subtype: Subtype of array elements
func (w *KeyValues3) SetArrayElementCount(count int32, type_ uint8, subtype uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetArrayElementCount(w.handle, count, type_, subtype)
	return nil
}

// SetToEmptyKV3Array 
//  @brief Sets the KeyValues3 object to an empty KV3 array
//
//  @param kv: Pointer to the KeyValues3 object
func (w *KeyValues3) SetToEmptyKV3Array() error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetToEmptyKV3Array(w.handle)
	return nil
}

// GetArrayElement 
//  @brief Gets an array element at the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index of the element to get
//
//  @return Pointer to the element KeyValues3 object, or nullptr if invalid
func (w *KeyValues3) GetArrayElement(elem int32) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3GetArrayElement(w.handle, elem)), nil
}

// ArrayInsertElementBefore 
//  @brief Inserts a new element before the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index before which to insert
//
//  @return Pointer to the newly inserted element, or nullptr if invalid
func (w *KeyValues3) ArrayInsertElementBefore(elem int32) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3ArrayInsertElementBefore(w.handle, elem)), nil
}

// ArrayInsertElementAfter 
//  @brief Inserts a new element after the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index after which to insert
//
//  @return Pointer to the newly inserted element, or nullptr if invalid
func (w *KeyValues3) ArrayInsertElementAfter(elem int32) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3ArrayInsertElementAfter(w.handle, elem)), nil
}

// ArrayAddElementToTail 
//  @brief Adds a new element to the end of the array
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Pointer to the newly added element, or nullptr if invalid
func (w *KeyValues3) ArrayAddElementToTail() (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3ArrayAddElementToTail(w.handle)), nil
}

// ArraySwapItems 
//  @brief Swaps two array elements
//
//  @param kv: Pointer to the KeyValues3 object
//  @param idx1: Index of the first element
//  @param idx2: Index of the second element
func (w *KeyValues3) ArraySwapItems(idx1 int32, idx2 int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3ArraySwapItems(w.handle, idx1, idx2)
	return nil
}

// ArrayRemoveElement 
//  @brief Removes an element from the array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param elem: Index of the element to remove
func (w *KeyValues3) ArrayRemoveElement(elem int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3ArrayRemoveElement(w.handle, elem)
	return nil
}

// SetToEmptyTable 
//  @brief Sets the KeyValues3 object to an empty table
//
//  @param kv: Pointer to the KeyValues3 object
func (w *KeyValues3) SetToEmptyTable() error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetToEmptyTable(w.handle)
	return nil
}

// GetMemberCount 
//  @brief Gets the number of members in the table
//
//  @param kv: Pointer to the KeyValues3 object
//
//  @return Number of table members, or 0 if kv is null or not a table
func (w *KeyValues3) GetMemberCount() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberCount(w.handle), nil
}

// HasMember 
//  @brief Checks if a member with the specified name exists
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to check
//
//  @return true if the member exists, false otherwise
func (w *KeyValues3) HasMember(name string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3HasMember(w.handle, name), nil
}

// FindMember 
//  @brief Finds a member by name
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to find
//
//  @return Pointer to the member KeyValues3 object, or nullptr if not found
func (w *KeyValues3) FindMember(name string) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3FindMember(w.handle, name)), nil
}

// FindOrCreateMember 
//  @brief Finds a member by name, or creates it if it doesn't exist
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to find or create
//
//  @return Pointer to the member KeyValues3 object, or nullptr if kv is null
func (w *KeyValues3) FindOrCreateMember(name string) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3FindOrCreateMember(w.handle, name)), nil
}

// RemoveMember 
//  @brief Removes a member from the table
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member to remove
//
//  @return true if the member was removed, false otherwise
func (w *KeyValues3) RemoveMember(name string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3RemoveMember(w.handle, name), nil
}

// GetMemberName 
//  @brief Gets the name of a member at the specified index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param index: Index of the member
//
//  @return Name of the member, or empty string if invalid
func (w *KeyValues3) GetMemberName(index int32) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberName(w.handle, index), nil
}

// GetMemberByIndex 
//  @brief Gets a member by index
//
//  @param kv: Pointer to the KeyValues3 object
//  @param index: Index of the member to get
//
//  @return Pointer to the member KeyValues3 object, or nullptr if invalid
func (w *KeyValues3) GetMemberByIndex(index int32) (*KeyValues3, error) {
	if w.handle == 0 {
		var zero *KeyValues3
		return zero, KeyValues3ErrEmptyHandle
	}
	return NewKeyValues3Borrowed(Kv3GetMemberByIndex(w.handle, index)), nil
}

// GetMemberBool 
//  @brief Gets a boolean value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Boolean value or defaultValue
func (w *KeyValues3) GetMemberBool(name string, defaultValue bool) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberBool(w.handle, name, defaultValue), nil
}

// GetMemberChar 
//  @brief Gets a char value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Char value or defaultValue
func (w *KeyValues3) GetMemberChar(name string, defaultValue int8) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberChar(w.handle, name, defaultValue), nil
}

// GetMemberUChar32 
//  @brief Gets a 32-bit Unicode character value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return 32-bit Unicode character value or defaultValue
func (w *KeyValues3) GetMemberUChar32(name string, defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberUChar32(w.handle, name, defaultValue), nil
}

// GetMemberInt8 
//  @brief Gets a signed 8-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int8_t value or defaultValue
func (w *KeyValues3) GetMemberInt8(name string, defaultValue int8) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberInt8(w.handle, name, defaultValue), nil
}

// GetMemberUInt8 
//  @brief Gets an unsigned 8-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint8_t value or defaultValue
func (w *KeyValues3) GetMemberUInt8(name string, defaultValue uint8) (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberUInt8(w.handle, name, defaultValue), nil
}

// GetMemberShort 
//  @brief Gets a signed 16-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int16_t value or defaultValue
func (w *KeyValues3) GetMemberShort(name string, defaultValue int16) (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberShort(w.handle, name, defaultValue), nil
}

// GetMemberUShort 
//  @brief Gets an unsigned 16-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint16_t value or defaultValue
func (w *KeyValues3) GetMemberUShort(name string, defaultValue uint16) (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberUShort(w.handle, name, defaultValue), nil
}

// GetMemberInt 
//  @brief Gets a signed 32-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int32_t value or defaultValue
func (w *KeyValues3) GetMemberInt(name string, defaultValue int32) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberInt(w.handle, name, defaultValue), nil
}

// GetMemberUInt 
//  @brief Gets an unsigned 32-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint32_t value or defaultValue
func (w *KeyValues3) GetMemberUInt(name string, defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberUInt(w.handle, name, defaultValue), nil
}

// GetMemberInt64 
//  @brief Gets a signed 64-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return int64_t value or defaultValue
func (w *KeyValues3) GetMemberInt64(name string, defaultValue int64) (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberInt64(w.handle, name, defaultValue), nil
}

// GetMemberUInt64 
//  @brief Gets an unsigned 64-bit integer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return uint64_t value or defaultValue
func (w *KeyValues3) GetMemberUInt64(name string, defaultValue uint64) (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberUInt64(w.handle, name, defaultValue), nil
}

// GetMemberFloat 
//  @brief Gets a float value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Float value or defaultValue
func (w *KeyValues3) GetMemberFloat(name string, defaultValue float32) (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberFloat(w.handle, name, defaultValue), nil
}

// GetMemberDouble 
//  @brief Gets a double value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Double value or defaultValue
func (w *KeyValues3) GetMemberDouble(name string, defaultValue float64) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberDouble(w.handle, name, defaultValue), nil
}

// GetMemberPointer 
//  @brief Gets a pointer value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default value to return if member not found
//
//  @return Pointer value as uintptr_t or defaultValue
func (w *KeyValues3) GetMemberPointer(name string, defaultValue uintptr) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberPointer(w.handle, name, defaultValue), nil
}

// GetMemberStringToken 
//  @brief Gets a string token value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default token value to return if member not found
//
//  @return String token hash code or defaultValue
func (w *KeyValues3) GetMemberStringToken(name string, defaultValue uint32) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberStringToken(w.handle, name, defaultValue), nil
}

// GetMemberEHandle 
//  @brief Gets an entity handle value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default entity handle value to return if member not found
//
//  @return Entity handle as int32_t or defaultValue
func (w *KeyValues3) GetMemberEHandle(name string, defaultValue int32) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberEHandle(w.handle, name, defaultValue), nil
}

// GetMemberString 
//  @brief Gets a string value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default string to return if member not found
//
//  @return String value or defaultValue
func (w *KeyValues3) GetMemberString(name string, defaultValue string) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberString(w.handle, name, defaultValue), nil
}

// GetMemberColor 
//  @brief Gets a color value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default color value to return if member not found
//
//  @return Color value as vec4 or defaultValue
func (w *KeyValues3) GetMemberColor(name string, defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberColor(w.handle, name, defaultValue), nil
}

// GetMemberVector 
//  @brief Gets a 3D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default vector to return if member not found
//
//  @return 3D vector or defaultValue
func (w *KeyValues3) GetMemberVector(name string, defaultValue plugify.Vector3) (plugify.Vector3, error) {
	if w.handle == 0 {
		var zero plugify.Vector3
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberVector(w.handle, name, defaultValue), nil
}

// GetMemberVector2D 
//  @brief Gets a 2D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default 2D vector to return if member not found
//
//  @return 2D vector or defaultValue
func (w *KeyValues3) GetMemberVector2D(name string, defaultValue plugify.Vector2) (plugify.Vector2, error) {
	if w.handle == 0 {
		var zero plugify.Vector2
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberVector2D(w.handle, name, defaultValue), nil
}

// GetMemberVector4D 
//  @brief Gets a 4D vector value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default 4D vector to return if member not found
//
//  @return 4D vector or defaultValue
func (w *KeyValues3) GetMemberVector4D(name string, defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberVector4D(w.handle, name, defaultValue), nil
}

// GetMemberQuaternion 
//  @brief Gets a quaternion value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default quaternion to return if member not found
//
//  @return Quaternion as vec4 or defaultValue
func (w *KeyValues3) GetMemberQuaternion(name string, defaultValue plugify.Vector4) (plugify.Vector4, error) {
	if w.handle == 0 {
		var zero plugify.Vector4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberQuaternion(w.handle, name, defaultValue), nil
}

// GetMemberQAngle 
//  @brief Gets an angle (QAngle) value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default angle to return if member not found
//
//  @return QAngle as vec3 or defaultValue
func (w *KeyValues3) GetMemberQAngle(name string, defaultValue plugify.Vector3) (plugify.Vector3, error) {
	if w.handle == 0 {
		var zero plugify.Vector3
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberQAngle(w.handle, name, defaultValue), nil
}

// GetMemberMatrix3x4 
//  @brief Gets a 3x4 matrix value from a table member
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param defaultValue: Default matrix to return if member not found
//
//  @return 3x4 matrix as mat4x4 or defaultValue
func (w *KeyValues3) GetMemberMatrix3x4(name string, defaultValue plugify.Matrix4x4) (plugify.Matrix4x4, error) {
	if w.handle == 0 {
		var zero plugify.Matrix4x4
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3GetMemberMatrix3x4(w.handle, name, defaultValue), nil
}

// SetMemberToNull 
//  @brief Sets a table member to null
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func (w *KeyValues3) SetMemberToNull(name string) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToNull(w.handle, name)
	return nil
}

// SetMemberToEmptyArray 
//  @brief Sets a table member to an empty array
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func (w *KeyValues3) SetMemberToEmptyArray(name string) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToEmptyArray(w.handle, name)
	return nil
}

// SetMemberToEmptyTable 
//  @brief Sets a table member to an empty table
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
func (w *KeyValues3) SetMemberToEmptyTable(name string) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToEmptyTable(w.handle, name)
	return nil
}

// SetMemberToBinaryBlob 
//  @brief Sets a table member to a binary blob (copies the data)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param blob: Vector containing the binary blob data
func (w *KeyValues3) SetMemberToBinaryBlob(name string, blob []uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToBinaryBlob(w.handle, name, blob)
	return nil
}

// SetMemberToBinaryBlobExternal 
//  @brief Sets a table member to an external binary blob (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param blob: Vector containing the external binary blob data
//  @param free_mem: Whether to free the memory when the object is destroyed
func (w *KeyValues3) SetMemberToBinaryBlobExternal(name string, blob []uint8, free_mem bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToBinaryBlobExternal(w.handle, name, blob, free_mem)
	return nil
}

// SetMemberToCopyOfValue 
//  @brief Sets a table member to a copy of another KeyValues3 value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param other: Pointer to the KeyValues3 object to copy
func (w *KeyValues3) SetMemberToCopyOfValue(name string, other uintptr) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberToCopyOfValue(w.handle, name, other)
	return nil
}

// SetMemberBool 
//  @brief Sets a table member to a boolean value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Boolean value to set
func (w *KeyValues3) SetMemberBool(name string, value bool) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberBool(w.handle, name, value)
	return nil
}

// SetMemberChar 
//  @brief Sets a table member to a char value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Char value to set
func (w *KeyValues3) SetMemberChar(name string, value int8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberChar(w.handle, name, value)
	return nil
}

// SetMemberUChar32 
//  @brief Sets a table member to a 32-bit Unicode character value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: 32-bit Unicode character value to set
func (w *KeyValues3) SetMemberUChar32(name string, value uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberUChar32(w.handle, name, value)
	return nil
}

// SetMemberInt8 
//  @brief Sets a table member to a signed 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int8_t value to set
func (w *KeyValues3) SetMemberInt8(name string, value int8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberInt8(w.handle, name, value)
	return nil
}

// SetMemberUInt8 
//  @brief Sets a table member to an unsigned 8-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint8_t value to set
func (w *KeyValues3) SetMemberUInt8(name string, value uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberUInt8(w.handle, name, value)
	return nil
}

// SetMemberShort 
//  @brief Sets a table member to a signed 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int16_t value to set
func (w *KeyValues3) SetMemberShort(name string, value int16) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberShort(w.handle, name, value)
	return nil
}

// SetMemberUShort 
//  @brief Sets a table member to an unsigned 16-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint16_t value to set
func (w *KeyValues3) SetMemberUShort(name string, value uint16) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberUShort(w.handle, name, value)
	return nil
}

// SetMemberInt 
//  @brief Sets a table member to a signed 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int32_t value to set
func (w *KeyValues3) SetMemberInt(name string, value int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberInt(w.handle, name, value)
	return nil
}

// SetMemberUInt 
//  @brief Sets a table member to an unsigned 32-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint32_t value to set
func (w *KeyValues3) SetMemberUInt(name string, value uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberUInt(w.handle, name, value)
	return nil
}

// SetMemberInt64 
//  @brief Sets a table member to a signed 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: int64_t value to set
func (w *KeyValues3) SetMemberInt64(name string, value int64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberInt64(w.handle, name, value)
	return nil
}

// SetMemberUInt64 
//  @brief Sets a table member to an unsigned 64-bit integer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: uint64_t value to set
func (w *KeyValues3) SetMemberUInt64(name string, value uint64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberUInt64(w.handle, name, value)
	return nil
}

// SetMemberFloat 
//  @brief Sets a table member to a float value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Float value to set
func (w *KeyValues3) SetMemberFloat(name string, value float32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberFloat(w.handle, name, value)
	return nil
}

// SetMemberDouble 
//  @brief Sets a table member to a double value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param value: Double value to set
func (w *KeyValues3) SetMemberDouble(name string, value float64) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberDouble(w.handle, name, value)
	return nil
}

// SetMemberPointer 
//  @brief Sets a table member to a pointer value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ptr: Pointer value as uintptr_t to set
func (w *KeyValues3) SetMemberPointer(name string, ptr uintptr) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberPointer(w.handle, name, ptr)
	return nil
}

// SetMemberStringToken 
//  @brief Sets a table member to a string token value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param token: String token hash code to set
func (w *KeyValues3) SetMemberStringToken(name string, token uint32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberStringToken(w.handle, name, token)
	return nil
}

// SetMemberEHandle 
//  @brief Sets a table member to an entity handle value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ehandle: Entity handle value to set
func (w *KeyValues3) SetMemberEHandle(name string, ehandle int32) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberEHandle(w.handle, name, ehandle)
	return nil
}

// SetMemberString 
//  @brief Sets a table member to a string value (copies the string)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param str: String value to set
//  @param subtype: String subtype enumeration value
func (w *KeyValues3) SetMemberString(name string, str string, subtype uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberString(w.handle, name, str, subtype)
	return nil
}

// SetMemberStringExternal 
//  @brief Sets a table member to an external string value (does not copy)
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param str: External string value to reference
//  @param subtype: String subtype enumeration value
func (w *KeyValues3) SetMemberStringExternal(name string, str string, subtype uint8) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberStringExternal(w.handle, name, str, subtype)
	return nil
}

// SetMemberColor 
//  @brief Sets a table member to a color value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param color: Color value as vec4 to set
func (w *KeyValues3) SetMemberColor(name string, color plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberColor(w.handle, name, color)
	return nil
}

// SetMemberVector 
//  @brief Sets a table member to a 3D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec: 3D vector to set
func (w *KeyValues3) SetMemberVector(name string, vec plugify.Vector3) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberVector(w.handle, name, vec)
	return nil
}

// SetMemberVector2D 
//  @brief Sets a table member to a 2D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec2d: 2D vector to set
func (w *KeyValues3) SetMemberVector2D(name string, vec2d plugify.Vector2) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberVector2D(w.handle, name, vec2d)
	return nil
}

// SetMemberVector4D 
//  @brief Sets a table member to a 4D vector value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param vec4d: 4D vector to set
func (w *KeyValues3) SetMemberVector4D(name string, vec4d plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberVector4D(w.handle, name, vec4d)
	return nil
}

// SetMemberQuaternion 
//  @brief Sets a table member to a quaternion value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param quat: Quaternion to set (as vec4)
func (w *KeyValues3) SetMemberQuaternion(name string, quat plugify.Vector4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberQuaternion(w.handle, name, quat)
	return nil
}

// SetMemberQAngle 
//  @brief Sets a table member to an angle (QAngle) value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param ang: QAngle to set (as vec3)
func (w *KeyValues3) SetMemberQAngle(name string, ang plugify.Vector3) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberQAngle(w.handle, name, ang)
	return nil
}

// SetMemberMatrix3x4 
//  @brief Sets a table member to a 3x4 matrix value
//
//  @param kv: Pointer to the KeyValues3 object
//  @param name: Name of the member
//  @param matrix: 3x4 matrix to set (as mat4x4)
func (w *KeyValues3) SetMemberMatrix3x4(name string, matrix plugify.Matrix4x4) error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3SetMemberMatrix3x4(w.handle, name, matrix)
	return nil
}

// DebugPrint 
//  @brief Prints debug information about the KeyValues3 object
//
//  @param kv: Pointer to the KeyValues3 object
func (w *KeyValues3) DebugPrint() error {
	if w.handle == 0 {
		return KeyValues3ErrEmptyHandle
	}
	Kv3DebugPrint(w.handle)
	return nil
}

// Load 
//  @brief Loads KeyValues3 data from a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) Load(error_ *string, input []uint8, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3Load(w.handle, error_, input, kv_name, flags), nil
}

// LoadFromText 
//  @brief Loads KeyValues3 data from a text string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Text string containing KV3 data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromText(error_ *string, input string, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromText(w.handle, error_, input, kv_name, flags), nil
}

// LoadFromFile 
//  @brief Loads KeyValues3 data from a file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param filename: Name of the file to load
//  @param path: Path to the file
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromFile(error_ *string, filename string, path string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromFile(w.handle, error_, filename, path, flags), nil
}

// LoadFromJSON 
//  @brief Loads KeyValues3 data from a JSON string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: JSON string
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromJSON(error_ *string, input string, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromJSON(w.handle, error_, input, kv_name, flags), nil
}

// LoadFromJSONFile 
//  @brief Loads KeyValues3 data from a JSON file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param path: Path to the file
//  @param filename: Name of the file to load
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromJSONFile(error_ *string, path string, filename string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromJSONFile(w.handle, error_, path, filename, flags), nil
}

// LoadFromKV1File 
//  @brief Loads KeyValues3 data from a KeyValues1 file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param path: Path to the file
//  @param filename: Name of the file to load
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromKV1File(error_ *string, path string, filename string, esc_behavior uint8, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromKV1File(w.handle, error_, path, filename, esc_behavior, flags), nil
}

// LoadFromKV1Text 
//  @brief Loads KeyValues3 data from a KeyValues1 text string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: KV1 text string
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param kv_name: Name for the KeyValues3 object
//  @param unk: Unknown boolean parameter
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromKV1Text(error_ *string, input string, esc_behavior uint8, kv_name string, unk bool, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromKV1Text(w.handle, error_, input, esc_behavior, kv_name, unk, flags), nil
}

// LoadFromKV1TextTranslated 
//  @brief Loads KeyValues3 data from a KeyValues1 text string with translation
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: KV1 text string
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param translation: Pointer to translation table
//  @param unk1: Unknown integer parameter
//  @param kv_name: Name for the KeyValues3 object
//  @param unk2: Unknown boolean parameter
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromKV1TextTranslated(error_ *string, input string, esc_behavior uint8, translation uintptr, unk1 int32, kv_name string, unk2 bool, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromKV1TextTranslated(w.handle, error_, input, esc_behavior, translation, unk1, kv_name, unk2, flags), nil
}

// LoadFromKV3OrKV1 
//  @brief Loads data from a buffer that may be KV3 or KV1 format
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromKV3OrKV1(error_ *string, input []uint8, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromKV3OrKV1(w.handle, error_, input, kv_name, flags), nil
}

// LoadFromOldSchemaText 
//  @brief Loads KeyValues3 data from old schema text format
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Vector containing the input buffer data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadFromOldSchemaText(error_ *string, input []uint8, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadFromOldSchemaText(w.handle, error_, input, kv_name, flags), nil
}

// LoadTextNoHeader 
//  @brief Loads KeyValues3 text without a header
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param input: Text string containing KV3 data
//  @param kv_name: Name for the KeyValues3 object
//  @param flags: Loading flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) LoadTextNoHeader(error_ *string, input string, kv_name string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3LoadTextNoHeader(w.handle, error_, input, kv_name, flags), nil
}

// Save 
//  @brief Saves KeyValues3 data to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output buffer data
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) Save(error_ *string, output *[]uint8, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3Save(w.handle, error_, output, flags), nil
}

// SaveAsJSON 
//  @brief Saves KeyValues3 data as JSON to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output JSON data
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveAsJSON(error_ *string, output *[]uint8) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveAsJSON(w.handle, error_, output), nil
}

// SaveAsJSONString 
//  @brief Saves KeyValues3 data as a JSON string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the JSON output
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveAsJSONString(error_ *string, output *string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveAsJSONString(w.handle, error_, output), nil
}

// SaveAsKV1Text 
//  @brief Saves KeyValues3 data as KeyValues1 text to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output KV1 text data
//  @param esc_behavior: Escape sequence behavior for KV1 text
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveAsKV1Text(error_ *string, output *[]uint8, esc_behavior uint8) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveAsKV1Text(w.handle, error_, output, esc_behavior), nil
}

// SaveAsKV1TextTranslated 
//  @brief Saves KeyValues3 data as KeyValues1 text with translation to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output KV1 text data
//  @param esc_behavior: Escape sequence behavior for KV1 text
//  @param translation: Pointer to translation table
//  @param unk: Unknown integer parameter
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveAsKV1TextTranslated(error_ *string, output *[]uint8, esc_behavior uint8, translation uintptr, unk int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveAsKV1TextTranslated(w.handle, error_, output, esc_behavior, translation, unk), nil
}

// SaveTextNoHeaderToBuffer 
//  @brief Saves KeyValues3 text without a header to a buffer
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: Vector to store the output text data
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveTextNoHeaderToBuffer(error_ *string, output *[]uint8, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveTextNoHeaderToBuffer(w.handle, error_, output, flags), nil
}

// SaveTextNoHeader 
//  @brief Saves KeyValues3 text without a header to a string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the text output
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveTextNoHeader(error_ *string, output *string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveTextNoHeader(w.handle, error_, output, flags), nil
}

// SaveTextToString 
//  @brief Saves KeyValues3 text to a string
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param output: String to store the text output
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveTextToString(error_ *string, output *string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveTextToString(w.handle, error_, output, flags), nil
}

// SaveToFile 
//  @brief Saves KeyValues3 data to a file
//
//  @param kv: Pointer to the KeyValues3 object
//  @param error_: Output string for error messages
//  @param filename: Name of the file to save
//  @param path: Path to save the file
//  @param flags: Saving flags
//
//  @return true if successful, false otherwise
func (w *KeyValues3) SaveToFile(error_ *string, filename string, path string, flags uint32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, KeyValues3ErrEmptyHandle
	}
	return Kv3SaveToFile(w.handle, error_, filename, path, flags), nil
}

