package s2sdk

/*
#include "schema.h"
#cgo noescape GetSchemaOffset
#cgo noescape GetSchemaChainOffset
#cgo noescape IsSchemaFieldNetworked
#cgo noescape GetSchemaClassSize
#cgo noescape GetEntData2
#cgo noescape SetEntData2
#cgo noescape GetEntDataFloat2
#cgo noescape SetEntDataFloat2
#cgo noescape GetEntDataColor2
#cgo noescape SetEntDataColor2
#cgo noescape GetEntDataString2
#cgo noescape SetEntDataString2
#cgo noescape GetEntDataCString2
#cgo noescape SetEntDataCString2
#cgo noescape GetEntDataVector3D2
#cgo noescape SetEntDataVector3D2
#cgo noescape GetEntDataVector4D2
#cgo noescape SetEntDataVector4D2
#cgo noescape GetEntDataVector2D2
#cgo noescape SetEntDataVector2D2
#cgo noescape GetEntDataEnt2
#cgo noescape SetEntDataEnt2
#cgo noescape ChangeEntityState2
#cgo noescape GetEntData
#cgo noescape SetEntData
#cgo noescape GetEntDataFloat
#cgo noescape SetEntDataFloat
#cgo noescape GetEntDataColor
#cgo noescape SetEntDataColor
#cgo noescape GetEntDataString
#cgo noescape SetEntDataString
#cgo noescape GetEntDataCString
#cgo noescape SetEntDataCString
#cgo noescape GetEntDataVector3D
#cgo noescape SetEntDataVector3D
#cgo noescape GetEntDataVector4D
#cgo noescape SetEntDataVector4D
#cgo noescape GetEntDataVector2D
#cgo noescape SetEntDataVector2D
#cgo noescape GetEntDataEnt
#cgo noescape SetEntDataEnt
#cgo noescape ChangeEntityState
#cgo noescape GetEntSchemaArraySize2
#cgo noescape GetEntSchema2
#cgo noescape SetEntSchema2
#cgo noescape GetEntSchemaFloat2
#cgo noescape SetEntSchemaFloat2
#cgo noescape GetEntSchemaColor2
#cgo noescape SetEntSchemaColor2
#cgo noescape GetEntSchemaString2
#cgo noescape SetEntSchemaString2
#cgo noescape GetEntSchemaVector3D2
#cgo noescape SetEntSchemaVector3D2
#cgo noescape GetEntSchemaVector2D2
#cgo noescape SetEntSchemaVector2D2
#cgo noescape GetEntSchemaVector4D2
#cgo noescape SetEntSchemaVector4D2
#cgo noescape GetEntSchemaEnt2
#cgo noescape SetEntSchemaEnt2
#cgo noescape PushEntSchemaEnt2
#cgo noescape EraseEntSchemaEnt2
#cgo noescape NetworkStateChanged2
#cgo noescape GetEntSchemaArraySize
#cgo noescape GetEntSchema
#cgo noescape SetEntSchema
#cgo noescape GetEntSchemaFloat
#cgo noescape SetEntSchemaFloat
#cgo noescape GetEntSchemaColor
#cgo noescape SetEntSchemaColor
#cgo noescape GetEntSchemaString
#cgo noescape SetEntSchemaString
#cgo noescape GetEntSchemaVector3D
#cgo noescape SetEntSchemaVector3D
#cgo noescape GetEntSchemaVector2D
#cgo noescape SetEntSchemaVector2D
#cgo noescape GetEntSchemaVector4D
#cgo noescape SetEntSchemaVector4D
#cgo noescape GetEntSchemaEnt
#cgo noescape SetEntSchemaEnt
#cgo noescape PushEntSchemaEnt
#cgo noescape EraseEntSchemaEnt
#cgo noescape NetworkStateChanged
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

// Generated from s2sdk (group: schema)

var _GetSchemaOffset = func(className string, memberName string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetSchemaOffset((*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetSchemaOffset 
//  @brief Get the offset of a member in a given schema class.
//
//  @param className: The name of the class.
//  @param memberName: The name of the member whose offset is to be retrieved.
//
//  @return The offset of the member in the class, or -1 if the offset is not found.
func GetSchemaOffset(className string, memberName string) int32 {
	return _GetSchemaOffset(className, memberName)
}

var _GetSchemaChainOffset = func(className string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetSchemaChainOffset((*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// GetSchemaChainOffset 
//  @brief Get the offset of a chain in a given schema class.
//
//  @param className: The name of the class.
//
//  @return The offset of the chain entity in the class, or -1 if the offset is not found.
func GetSchemaChainOffset(className string) int32 {
	return _GetSchemaChainOffset(className)
}

var _IsSchemaFieldNetworked = func(className string, memberName string) bool {
	var __retVal bool
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.IsSchemaFieldNetworked((*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// IsSchemaFieldNetworked 
//  @brief Check if a schema field is networked.
//
//  @param className: The name of the class.
//  @param memberName: The name of the member to check.
//
//  @return True if the member is networked, false otherwise.
func IsSchemaFieldNetworked(className string, memberName string) bool {
	return _IsSchemaFieldNetworked(className, memberName)
}

var _GetSchemaClassSize = func(className string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetSchemaClassSize((*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// GetSchemaClassSize 
//  @brief Get the size of a schema class.
//
//  @param className: The name of the class.
//
//  @return The size of the class in bytes, or -1 if the class is not found.
func GetSchemaClassSize(className string) int32 {
	return _GetSchemaClassSize(className)
}

var _GetEntData2 = func(entity uintptr, offset int32, size int32) int64 {
	var __retVal int64
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = int64(C.GetEntData2(__entity, __offset, __size))
	return __retVal
}

// GetEntData2 
//  @brief Peeks into an entity's object schema and retrieves the integer value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read (valid values are 1, 2, 4 or 8).
//
//  @return The integer value at the given memory location.
func GetEntData2(entity uintptr, offset int32, size int32) int64 {
	return _GetEntData2(entity, offset, size)
}

var _SetEntData2 = func(entity uintptr, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.int64_t(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntData2(__entity, __offset, __value, __size, __changeState, __chainOffset)
}

// SetEntData2 
//  @brief Peeks into an entity's object data and sets the integer value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The integer value to set.
//  @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntData2(entity uintptr, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	_SetEntData2(entity, offset, value, size, changeState, chainOffset)
}

var _GetEntDataFloat2 = func(entity uintptr, offset int32, size int32) float64 {
	var __retVal float64
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = float64(C.GetEntDataFloat2(__entity, __offset, __size))
	return __retVal
}

// GetEntDataFloat2 
//  @brief Peeks into an entity's object schema and retrieves the float value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read (valid values are 1, 2, 4 or 8).
//
//  @return The float value at the given memory location.
func GetEntDataFloat2(entity uintptr, offset int32, size int32) float64 {
	return _GetEntDataFloat2(entity, offset, size)
}

var _SetEntDataFloat2 = func(entity uintptr, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.double(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataFloat2(__entity, __offset, __value, __size, __changeState, __chainOffset)
}

// SetEntDataFloat2 
//  @brief Peeks into an entity's object data and sets the float value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The float value to set.
//  @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataFloat2(entity uintptr, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	_SetEntDataFloat2(entity, offset, value, size, changeState, chainOffset)
}

var _GetEntDataColor2 = func(entity uintptr, offset int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataColor2(__entity, __offset)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataColor2 
//  @brief Peeks into an entity's object schema and retrieves the color value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The color value at the given memory location.
func GetEntDataColor2(entity uintptr, offset int32) plugify.Vector4 {
	return _GetEntDataColor2(entity, offset)
}

var _SetEntDataColor2 = func(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataColor2(__entity, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataColor2 
//  @brief Peeks into an entity's object data and sets the color at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The color value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataColor2(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	_SetEntDataColor2(entity, offset, value, changeState, chainOffset)
}

var _GetEntDataString2 = func(entity uintptr, offset int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	plugify.Block {
		Try: func() {
			__native := C.GetEntDataString2(__entity, __offset)
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

// GetEntDataString2 
//  @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The string value at the given memory location.
func GetEntDataString2(entity uintptr, offset int32) string {
	return _GetEntDataString2(entity, offset)
}

var _SetEntDataString2 = func(entity uintptr, offset int32, value string, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block {
		Try: func() {
			C.SetEntDataString2(__entity, __offset, (*C.String)(unsafe.Pointer(&__value)), __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntDataString2 
//  @brief Peeks into an entity's object data and sets the string at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The string value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataString2(entity uintptr, offset int32, value string, changeState bool, chainOffset int32) {
	_SetEntDataString2(entity, offset, value, changeState, chainOffset)
}

var _GetEntDataCString2 = func(entity uintptr, offset int32, size int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	plugify.Block {
		Try: func() {
			__native := C.GetEntDataCString2(__entity, __offset, __size)
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

// GetEntDataCString2 
//  @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read.
//
//  @return The string value at the given memory location.
func GetEntDataCString2(entity uintptr, offset int32, size int32) string {
	return _GetEntDataCString2(entity, offset, size)
}

var _SetEntDataCString2 = func(entity uintptr, offset int32, value string, size int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block {
		Try: func() {
			C.SetEntDataCString2(__entity, __offset, (*C.String)(unsafe.Pointer(&__value)), __size, __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntDataCString2 
//  @brief Peeks into an entity's object data and sets the string at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The string value to set.
//  @param size: Number of bytes to write.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataCString2(entity uintptr, offset int32, value string, size int32, changeState bool, chainOffset int32) {
	_SetEntDataCString2(entity, offset, value, size, changeState, chainOffset)
}

var _GetEntDataVector3D2 = func(entity uintptr, offset int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector3D2(__entity, __offset)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector3D2 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector3D2(entity uintptr, offset int32) plugify.Vector3 {
	return _GetEntDataVector3D2(entity, offset)
}

var _SetEntDataVector3D2 = func(entity uintptr, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector3D2(__entity, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector3D2 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector3D2(entity uintptr, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	_SetEntDataVector3D2(entity, offset, value, changeState, chainOffset)
}

var _GetEntDataVector4D2 = func(entity uintptr, offset int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector4D2(__entity, __offset)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector4D2 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector4D2(entity uintptr, offset int32) plugify.Vector4 {
	return _GetEntDataVector4D2(entity, offset)
}

var _SetEntDataVector4D2 = func(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector4D2(__entity, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector4D2 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector4D2(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	_SetEntDataVector4D2(entity, offset, value, changeState, chainOffset)
}

var _GetEntDataVector2D2 = func(entity uintptr, offset int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector2D2(__entity, __offset)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector2D2 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector2D2(entity uintptr, offset int32) plugify.Vector2 {
	return _GetEntDataVector2D2(entity, offset)
}

var _SetEntDataVector2D2 = func(entity uintptr, offset int32, value plugify.Vector2, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector2D2(__entity, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector2D2 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector2D2(entity uintptr, offset int32, value plugify.Vector2, changeState bool, chainOffset int32) {
	_SetEntDataVector2D2(entity, offset, value, changeState, chainOffset)
}

var _GetEntDataEnt2 = func(entity uintptr, offset int32) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__retVal = int32(C.GetEntDataEnt2(__entity, __offset))
	return __retVal
}

// GetEntDataEnt2 
//  @brief Peeks into an entity's object data and retrieves the entity handle at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//
//  @return The entity handle at the given memory location.
func GetEntDataEnt2(entity uintptr, offset int32) int32 {
	return _GetEntDataEnt2(entity, offset)
}

var _SetEntDataEnt2 = func(entity uintptr, offset int32, value int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataEnt2(__entity, __offset, __value, __changeState, __chainOffset)
}

// SetEntDataEnt2 
//  @brief Peeks into an entity's object data and sets the entity handle at the given offset.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param value: The entity handle to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataEnt2(entity uintptr, offset int32, value int32, changeState bool, chainOffset int32) {
	_SetEntDataEnt2(entity, offset, value, changeState, chainOffset)
}

var _ChangeEntityState2 = func(entity uintptr, offset int32, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__chainOffset := C.int32_t(chainOffset)
	C.ChangeEntityState2(__entity, __offset, __chainOffset)
}

// ChangeEntityState2 
//  @brief Updates the networked state of a schema field for a given entity pointer.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param offset: The offset of the schema to use.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func ChangeEntityState2(entity uintptr, offset int32, chainOffset int32) {
	_ChangeEntityState2(entity, offset, chainOffset)
}

var _GetEntData = func(entityHandle int32, offset int32, size int32) int64 {
	var __retVal int64
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = int64(C.GetEntData(__entityHandle, __offset, __size))
	return __retVal
}

// GetEntData 
//  @brief Peeks into an entity's object schema and retrieves the integer value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read (valid values are 1, 2, 4 or 8).
//
//  @return The integer value at the given memory location.
func GetEntData(entityHandle int32, offset int32, size int32) int64 {
	return _GetEntData(entityHandle, offset, size)
}

var _SetEntData = func(entityHandle int32, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.int64_t(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntData(__entityHandle, __offset, __value, __size, __changeState, __chainOffset)
}

// SetEntData 
//  @brief Peeks into an entity's object data and sets the integer value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The integer value to set.
//  @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntData(entityHandle int32, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	_SetEntData(entityHandle, offset, value, size, changeState, chainOffset)
}

var _GetEntDataFloat = func(entityHandle int32, offset int32, size int32) float64 {
	var __retVal float64
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = float64(C.GetEntDataFloat(__entityHandle, __offset, __size))
	return __retVal
}

// GetEntDataFloat 
//  @brief Peeks into an entity's object schema and retrieves the float value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read (valid values are 1, 2, 4 or 8).
//
//  @return The float value at the given memory location.
func GetEntDataFloat(entityHandle int32, offset int32, size int32) float64 {
	return _GetEntDataFloat(entityHandle, offset, size)
}

var _SetEntDataFloat = func(entityHandle int32, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.double(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataFloat(__entityHandle, __offset, __value, __size, __changeState, __chainOffset)
}

// SetEntDataFloat 
//  @brief Peeks into an entity's object data and sets the float value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The float value to set.
//  @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataFloat(entityHandle int32, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	_SetEntDataFloat(entityHandle, offset, value, size, changeState, chainOffset)
}

var _GetEntDataColor = func(entityHandle int32, offset int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataColor(__entityHandle, __offset)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataColor 
//  @brief Peeks into an entity's object schema and retrieves the color value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The color value at the given memory location.
func GetEntDataColor(entityHandle int32, offset int32) plugify.Vector4 {
	return _GetEntDataColor(entityHandle, offset)
}

var _SetEntDataColor = func(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataColor(__entityHandle, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataColor 
//  @brief Peeks into an entity's object data and sets the color at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The color value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataColor(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	_SetEntDataColor(entityHandle, offset, value, changeState, chainOffset)
}

var _GetEntDataString = func(entityHandle int32, offset int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	plugify.Block {
		Try: func() {
			__native := C.GetEntDataString(__entityHandle, __offset)
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

// GetEntDataString 
//  @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The string value at the given memory location.
func GetEntDataString(entityHandle int32, offset int32) string {
	return _GetEntDataString(entityHandle, offset)
}

var _SetEntDataString = func(entityHandle int32, offset int32, value string, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block {
		Try: func() {
			C.SetEntDataString(__entityHandle, __offset, (*C.String)(unsafe.Pointer(&__value)), __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntDataString 
//  @brief Peeks into an entity's object data and sets the string at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The string value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataString(entityHandle int32, offset int32, value string, changeState bool, chainOffset int32) {
	_SetEntDataString(entityHandle, offset, value, changeState, chainOffset)
}

var _GetEntDataCString = func(entityHandle int32, offset int32, size int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	plugify.Block {
		Try: func() {
			__native := C.GetEntDataCString(__entityHandle, __offset, __size)
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

// GetEntDataCString 
//  @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param size: Number of bytes to read.
//
//  @return The string value at the given memory location.
func GetEntDataCString(entityHandle int32, offset int32, size int32) string {
	return _GetEntDataCString(entityHandle, offset, size)
}

var _SetEntDataCString = func(entityHandle int32, offset int32, value string, size int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block {
		Try: func() {
			C.SetEntDataCString(__entityHandle, __offset, (*C.String)(unsafe.Pointer(&__value)), __size, __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntDataCString 
//  @brief Peeks into an entity's object data and sets the string at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The string value to set.
//  @param size: Number of bytes to write.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataCString(entityHandle int32, offset int32, value string, size int32, changeState bool, chainOffset int32) {
	_SetEntDataCString(entityHandle, offset, value, size, changeState, chainOffset)
}

var _GetEntDataVector3D = func(entityHandle int32, offset int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector3D(__entityHandle, __offset)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector3D 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector3D(entityHandle int32, offset int32) plugify.Vector3 {
	return _GetEntDataVector3D(entityHandle, offset)
}

var _SetEntDataVector3D = func(entityHandle int32, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector3D(__entityHandle, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector3D 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector3D(entityHandle int32, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	_SetEntDataVector3D(entityHandle, offset, value, changeState, chainOffset)
}

var _GetEntDataVector4D = func(entityHandle int32, offset int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector4D(__entityHandle, __offset)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector4D 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector4D(entityHandle int32, offset int32) plugify.Vector4 {
	return _GetEntDataVector4D(entityHandle, offset)
}

var _SetEntDataVector4D = func(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector4D(__entityHandle, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector4D 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector4D(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32) {
	_SetEntDataVector4D(entityHandle, offset, value, changeState, chainOffset)
}

var _GetEntDataVector2D = func(entityHandle int32, offset int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector2D(__entityHandle, __offset)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// GetEntDataVector2D 
//  @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The vector value at the given memory location.
func GetEntDataVector2D(entityHandle int32, offset int32) plugify.Vector2 {
	return _GetEntDataVector2D(entityHandle, offset)
}

var _SetEntDataVector2D = func(entityHandle int32, offset int32, value plugify.Vector2, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector2D(__entityHandle, __offset, &__value, __changeState, __chainOffset)
}

// SetEntDataVector2D 
//  @brief Peeks into an entity's object data and sets the vector at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataVector2D(entityHandle int32, offset int32, value plugify.Vector2, changeState bool, chainOffset int32) {
	_SetEntDataVector2D(entityHandle, offset, value, changeState, chainOffset)
}

var _GetEntDataEnt = func(entityHandle int32, offset int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__retVal = int32(C.GetEntDataEnt(__entityHandle, __offset))
	return __retVal
}

// GetEntDataEnt 
//  @brief Peeks into an entity's object data and retrieves the entity handle at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//
//  @return The entity handle at the given memory location.
func GetEntDataEnt(entityHandle int32, offset int32) int32 {
	return _GetEntDataEnt(entityHandle, offset)
}

var _SetEntDataEnt = func(entityHandle int32, offset int32, value int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataEnt(__entityHandle, __offset, __value, __changeState, __chainOffset)
}

// SetEntDataEnt 
//  @brief Peeks into an entity's object data and sets the entity handle at the given offset.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param value: The entity handle to set.
//  @param changeState: If true, change will be sent over the network.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func SetEntDataEnt(entityHandle int32, offset int32, value int32, changeState bool, chainOffset int32) {
	_SetEntDataEnt(entityHandle, offset, value, changeState, chainOffset)
}

var _ChangeEntityState = func(entityHandle int32, offset int32, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__chainOffset := C.int32_t(chainOffset)
	C.ChangeEntityState(__entityHandle, __offset, __chainOffset)
}

// ChangeEntityState 
//  @brief Updates the networked state of a schema field for a given entity handle.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param offset: The offset of the schema to use.
//  @param chainOffset: The offset of the chain entity in the class (-2 for non-entity classes).
func ChangeEntityState(entityHandle int32, offset int32, chainOffset int32) {
	_ChangeEntityState(entityHandle, offset, chainOffset)
}

var _GetEntSchemaArraySize2 = func(entity uintptr, className string, memberName string) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetEntSchemaArraySize2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaArraySize2 
//  @brief Retrieves the count of values that an entity schema's array can store.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//
//  @return Size of array (in elements) or 0 if schema is not an array.
func GetEntSchemaArraySize2(entity uintptr, className string, memberName string) int32 {
	return _GetEntSchemaArraySize2(entity, className, memberName)
}

var _GetEntSchema2 = func(entity uintptr, className string, memberName string, element int32) int64 {
	var __retVal int64
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.GetEntSchema2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchema2 
//  @brief Retrieves an integer value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return An integer value at the given schema offset.
func GetEntSchema2(entity uintptr, className string, memberName string, element int32) int64 {
	return _GetEntSchema2(entity, className, memberName, element)
}

var _SetEntSchema2 = func(entity uintptr, className string, memberName string, value int64, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int64_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchema2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchema2 
//  @brief Sets an integer value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The integer value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchema2(entity uintptr, className string, memberName string, value int64, changeState bool, element int32) {
	_SetEntSchema2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaFloat2 = func(entity uintptr, className string, memberName string, element int32) float64 {
	var __retVal float64
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = float64(C.GetEntSchemaFloat2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaFloat2 
//  @brief Retrieves a float value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A float value at the given schema offset.
func GetEntSchemaFloat2(entity uintptr, className string, memberName string, element int32) float64 {
	return _GetEntSchemaFloat2(entity, className, memberName, element)
}

var _SetEntSchemaFloat2 = func(entity uintptr, className string, memberName string, value float64, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.double(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaFloat2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaFloat2 
//  @brief Sets a float value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The float value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaFloat2(entity uintptr, className string, memberName string, value float64, changeState bool, element int32) {
	_SetEntSchemaFloat2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaColor2 = func(entity uintptr, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaColor2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaColor2 
//  @brief Retrieves a color value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A color value at the given schema offset.
func GetEntSchemaColor2(entity uintptr, className string, memberName string, element int32) plugify.Vector4 {
	return _GetEntSchemaColor2(entity, className, memberName, element)
}

var _SetEntSchemaColor2 = func(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaColor2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaColor2 
//  @brief Sets a color value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The color value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaColor2(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	_SetEntSchemaColor2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaString2 = func(entity uintptr, className string, memberName string, element int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaString2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaString2 
//  @brief Retrieves a string value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaString2(entity uintptr, className string, memberName string, element int32) string {
	return _GetEntSchemaString2(entity, className, memberName, element)
}

var _SetEntSchemaString2 = func(entity uintptr, className string, memberName string, value string, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaString2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), (*C.String)(unsafe.Pointer(&__value)), __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntSchemaString2 
//  @brief Sets a string value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The string value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaString2(entity uintptr, className string, memberName string, value string, changeState bool, element int32) {
	_SetEntSchemaString2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector3D2 = func(entity uintptr, className string, memberName string, element int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector3D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector3D2 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A vector value at the given schema offset.
func GetEntSchemaVector3D2(entity uintptr, className string, memberName string, element int32) plugify.Vector3 {
	return _GetEntSchemaVector3D2(entity, className, memberName, element)
}

var _SetEntSchemaVector3D2 = func(entity uintptr, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector3D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector3D2 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector3D2(entity uintptr, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	_SetEntSchemaVector3D2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector2D2 = func(entity uintptr, className string, memberName string, element int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector2D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector2D2 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A vector value at the given schema offset.
func GetEntSchemaVector2D2(entity uintptr, className string, memberName string, element int32) plugify.Vector2 {
	return _GetEntSchemaVector2D2(entity, className, memberName, element)
}

var _SetEntSchemaVector2D2 = func(entity uintptr, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector2D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector2D2 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector2D2(entity uintptr, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	_SetEntSchemaVector2D2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector4D2 = func(entity uintptr, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector4D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector4D2 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A vector value at the given schema offset.
func GetEntSchemaVector4D2(entity uintptr, className string, memberName string, element int32) plugify.Vector4 {
	return _GetEntSchemaVector4D2(entity, className, memberName, element)
}

var _SetEntSchemaVector4D2 = func(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector4D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector4D2 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector4D2(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	_SetEntSchemaVector4D2(entity, className, memberName, value, changeState, element)
}

var _GetEntSchemaEnt2 = func(entity uintptr, className string, memberName string, element int32) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaEnt2 
//  @brief Retrieves an entity handle from an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaEnt2(entity uintptr, className string, memberName string, element int32) int32 {
	return _GetEntSchemaEnt2(entity, className, memberName, element)
}

var _SetEntSchemaEnt2 = func(entity uintptr, className string, memberName string, value int32, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaEnt2 
//  @brief Sets an entity handle in an entity's schema.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The entity handle to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaEnt2(entity uintptr, className string, memberName string, value int32, changeState bool, element int32) {
	_SetEntSchemaEnt2(entity, className, memberName, value, changeState, element)
}

var _PushEntSchemaEnt2 = func(entity uintptr, className string, memberName string, value int32, changeState bool) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	plugify.Block {
		Try: func() {
			C.PushEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// PushEntSchemaEnt2 
//  @brief Pushes an entity handle into an entity's schema collection.
//
//  @param entity: Pointer to the instance of the class where the value is to be pushed.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The entity handle to push.
//  @param changeState: If true, change will be sent over the network.
func PushEntSchemaEnt2(entity uintptr, className string, memberName string, value int32, changeState bool) {
	_PushEntSchemaEnt2(entity, className, memberName, value, changeState)
}

var _EraseEntSchemaEnt2 = func(entity uintptr, className string, memberName string, element int32, changeState bool) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	__changeState := C.bool(changeState)
	plugify.Block {
		Try: func() {
			C.EraseEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element, __changeState)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// EraseEntSchemaEnt2 
//  @brief Erases an entity handle from an entity's schema collection by index.
//
//  @param entity: Pointer to the instance of the class where the value is to be erased.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element index to erase (starting from 0).
//  @param changeState: If true, change will be sent over the network.
func EraseEntSchemaEnt2(entity uintptr, className string, memberName string, element int32, changeState bool) {
	_EraseEntSchemaEnt2(entity, className, memberName, element, changeState)
}

var _NetworkStateChanged2 = func(entity uintptr, className string, memberName string) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			C.NetworkStateChanged2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// NetworkStateChanged2 
//  @brief Updates the networked state of a schema field for a given entity pointer.
//
//  @param entity: Pointer to the instance of the class where the value is to be set.
//  @param className: The name of the class that contains the member.
//  @param memberName: The name of the member to be set.
func NetworkStateChanged2(entity uintptr, className string, memberName string) {
	_NetworkStateChanged2(entity, className, memberName)
}

var _GetEntSchemaArraySize = func(entityHandle int32, className string, memberName string) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetEntSchemaArraySize(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaArraySize 
//  @brief Retrieves the count of values that an entity schema's array can store.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//
//  @return Size of array (in elements) or 0 if schema is not an array.
func GetEntSchemaArraySize(entityHandle int32, className string, memberName string) int32 {
	return _GetEntSchemaArraySize(entityHandle, className, memberName)
}

var _GetEntSchema = func(entityHandle int32, className string, memberName string, element int32) int64 {
	var __retVal int64
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.GetEntSchema(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchema 
//  @brief Retrieves an integer value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return An integer value at the given schema offset.
func GetEntSchema(entityHandle int32, className string, memberName string, element int32) int64 {
	return _GetEntSchema(entityHandle, className, memberName, element)
}

var _SetEntSchema = func(entityHandle int32, className string, memberName string, value int64, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int64_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchema(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchema 
//  @brief Sets an integer value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The integer value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchema(entityHandle int32, className string, memberName string, value int64, changeState bool, element int32) {
	_SetEntSchema(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaFloat = func(entityHandle int32, className string, memberName string, element int32) float64 {
	var __retVal float64
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = float64(C.GetEntSchemaFloat(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaFloat 
//  @brief Retrieves a float value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A float value at the given schema offset.
func GetEntSchemaFloat(entityHandle int32, className string, memberName string, element int32) float64 {
	return _GetEntSchemaFloat(entityHandle, className, memberName, element)
}

var _SetEntSchemaFloat = func(entityHandle int32, className string, memberName string, value float64, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.double(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaFloat(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaFloat 
//  @brief Sets a float value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The float value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaFloat(entityHandle int32, className string, memberName string, value float64, changeState bool, element int32) {
	_SetEntSchemaFloat(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaColor = func(entityHandle int32, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaColor(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaColor 
//  @brief Retrieves a color value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A color value at the given schema offset.
func GetEntSchemaColor(entityHandle int32, className string, memberName string, element int32) plugify.Vector4 {
	return _GetEntSchemaColor(entityHandle, className, memberName, element)
}

var _SetEntSchemaColor = func(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaColor(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaColor 
//  @brief Sets a color value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The color value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaColor(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	_SetEntSchemaColor(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaString = func(entityHandle int32, className string, memberName string, element int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaString(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaString 
//  @brief Retrieves a string value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaString(entityHandle int32, className string, memberName string, element int32) string {
	return _GetEntSchemaString(entityHandle, className, memberName, element)
}

var _SetEntSchemaString = func(entityHandle int32, className string, memberName string, value string, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaString(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), (*C.String)(unsafe.Pointer(&__value)), __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEntSchemaString 
//  @brief Sets a string value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The string value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaString(entityHandle int32, className string, memberName string, value string, changeState bool, element int32) {
	_SetEntSchemaString(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector3D = func(entityHandle int32, className string, memberName string, element int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector3D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector3D 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaVector3D(entityHandle int32, className string, memberName string, element int32) plugify.Vector3 {
	return _GetEntSchemaVector3D(entityHandle, className, memberName, element)
}

var _SetEntSchemaVector3D = func(entityHandle int32, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector3D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector3D 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector3D(entityHandle int32, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	_SetEntSchemaVector3D(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector2D = func(entityHandle int32, className string, memberName string, element int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector2D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector2D 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaVector2D(entityHandle int32, className string, memberName string, element int32) plugify.Vector2 {
	return _GetEntSchemaVector2D(entityHandle, className, memberName, element)
}

var _SetEntSchemaVector2D = func(entityHandle int32, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector2D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector2D 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector2D(entityHandle int32, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	_SetEntSchemaVector2D(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaVector4D = func(entityHandle int32, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__native := C.GetEntSchemaVector4D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaVector4D 
//  @brief Retrieves a vector value from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaVector4D(entityHandle int32, className string, memberName string, element int32) plugify.Vector4 {
	return _GetEntSchemaVector4D(entityHandle, className, memberName, element)
}

var _SetEntSchemaVector4D = func(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaVector4D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaVector4D 
//  @brief Sets a vector value in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The vector value to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector4D(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	_SetEntSchemaVector4D(entityHandle, className, memberName, value, changeState, element)
}

var _GetEntSchemaEnt = func(entityHandle int32, className string, memberName string, element int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchemaEnt 
//  @brief Retrieves an entity handle from an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element # (starting from 0) if schema is an array.
//
//  @return A string value at the given schema offset.
func GetEntSchemaEnt(entityHandle int32, className string, memberName string, element int32) int32 {
	return _GetEntSchemaEnt(entityHandle, className, memberName, element)
}

var _SetEntSchemaEnt = func(entityHandle int32, className string, memberName string, value int32, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block {
		Try: func() {
			C.SetEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// SetEntSchemaEnt 
//  @brief Sets an entity handle in an entity's schema.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The entity handle to set.
//  @param changeState: If true, change will be sent over the network.
//  @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaEnt(entityHandle int32, className string, memberName string, value int32, changeState bool, element int32) {
	_SetEntSchemaEnt(entityHandle, className, memberName, value, changeState, element)
}

var _PushEntSchemaEnt = func(entityHandle int32, className string, memberName string, value int32, changeState bool) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	plugify.Block {
		Try: func() {
			C.PushEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// PushEntSchemaEnt 
//  @brief Pushes an entity handle into an entity's schema collection.
//
//  @param entityHandle: The handle of the entity whose schema collection is modified.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param value: The entity handle to push.
//  @param changeState: If true, change will be sent over the network.
func PushEntSchemaEnt(entityHandle int32, className string, memberName string, value int32, changeState bool) {
	_PushEntSchemaEnt(entityHandle, className, memberName, value, changeState)
}

var _EraseEntSchemaEnt = func(entityHandle int32, className string, memberName string, element int32, changeState bool) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	__changeState := C.bool(changeState)
	plugify.Block {
		Try: func() {
			C.EraseEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element, __changeState)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// EraseEntSchemaEnt 
//  @brief Erases an entity handle from an entity's schema collection by index.
//
//  @param entityHandle: The handle of the entity whose schema collection is modified.
//  @param className: The name of the class.
//  @param memberName: The name of the schema member.
//  @param element: Element index to erase (starting from 0).
//  @param changeState: If true, change will be sent over the network.
func EraseEntSchemaEnt(entityHandle int32, className string, memberName string, element int32, changeState bool) {
	_EraseEntSchemaEnt(entityHandle, className, memberName, element, changeState)
}

var _NetworkStateChanged = func(entityHandle int32, className string, memberName string) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block {
		Try: func() {
			C.NetworkStateChanged(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// NetworkStateChanged 
//  @brief Updates the networked state of a schema field for a given entity handle.
//
//  @param entityHandle: The handle of the entity from which the value is to be retrieved.
//  @param className: The name of the class that contains the member.
//  @param memberName: The name of the member to be set.
func NetworkStateChanged(entityHandle int32, className string, memberName string) {
	_NetworkStateChanged(entityHandle, className, memberName)
}

