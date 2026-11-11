package s2sdk

/*
#include "bodies.h"
#cgo noescape AddBodyImpulseAtPosition
#cgo noescape AddBodyVelocity
#cgo noescape DetachBodyFromParent
#cgo noescape IsBodyAttachedToParent
#cgo noescape SetBodyAngularVelocity
#cgo noescape SetBodyMaterialGroup
#cgo noescape SetBodyVelocity
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

// Generated from s2sdk (group: bodies)

var _AddBodyImpulseAtPosition = func(entityHandle int32, position plugify.Vector3, impulse plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__position := *(*C.Vector3)(unsafe.Pointer(&position))
	__impulse := *(*C.Vector3)(unsafe.Pointer(&impulse))
	C.AddBodyImpulseAtPosition(__entityHandle, &__position, &__impulse)
}

// AddBodyImpulseAtPosition 
//  @brief Applies an impulse to an entity at a specific world position.
//
//  @param entityHandle: The handle of the entity.
//  @param position: The world position where the impulse will be applied.
//  @param impulse: The impulse vector to apply.
func AddBodyImpulseAtPosition(entityHandle int32, position plugify.Vector3, impulse plugify.Vector3) {
	_AddBodyImpulseAtPosition(entityHandle, position, impulse)
}

var _AddBodyVelocity = func(entityHandle int32, linearVelocity plugify.Vector3, angularVelocity plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__linearVelocity := *(*C.Vector3)(unsafe.Pointer(&linearVelocity))
	__angularVelocity := *(*C.Vector3)(unsafe.Pointer(&angularVelocity))
	C.AddBodyVelocity(__entityHandle, &__linearVelocity, &__angularVelocity)
}

// AddBodyVelocity 
//  @brief Adds linear and angular velocity to the entity's physics object.
//
//  @param entityHandle: The handle of the entity.
//  @param linearVelocity: The linear velocity vector to add.
//  @param angularVelocity: The angular velocity vector to add.
func AddBodyVelocity(entityHandle int32, linearVelocity plugify.Vector3, angularVelocity plugify.Vector3) {
	_AddBodyVelocity(entityHandle, linearVelocity, angularVelocity)
}

var _DetachBodyFromParent = func(entityHandle int32) {
	__entityHandle := C.int32_t(entityHandle)
	C.DetachBodyFromParent(__entityHandle)
}

// DetachBodyFromParent 
//  @brief Detaches the entity from its parent.
//
//  @param entityHandle: The handle of the entity.
func DetachBodyFromParent(entityHandle int32) {
	_DetachBodyFromParent(entityHandle)
}

var _IsBodyAttachedToParent = func(entityHandle int32) bool {
	var __retVal bool
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsBodyAttachedToParent(__entityHandle))
	return __retVal
}

// IsBodyAttachedToParent 
//  @brief Checks whether the entity is attached to a parent.
//
//  @param entityHandle: The handle of the entity.
//
//  @return True if attached to a parent, false otherwise.
func IsBodyAttachedToParent(entityHandle int32) bool {
	return _IsBodyAttachedToParent(entityHandle)
}

var _SetBodyAngularVelocity = func(entityHandle int32, angVelocity plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__angVelocity := *(*C.Vector3)(unsafe.Pointer(&angVelocity))
	C.SetBodyAngularVelocity(__entityHandle, &__angVelocity)
}

// SetBodyAngularVelocity 
//  @brief Sets the angular velocity of the entity.
//
//  @param entityHandle: The handle of the entity.
//  @param angVelocity: The new angular velocity vector.
func SetBodyAngularVelocity(entityHandle int32, angVelocity plugify.Vector3) {
	_SetBodyAngularVelocity(entityHandle, angVelocity)
}

var _SetBodyMaterialGroup = func(entityHandle int32, materialGroup string) {
	__entityHandle := C.int32_t(entityHandle)
	__materialGroup := plugify.ConstructString(materialGroup)
	plugify.Block {
		Try: func() {
			C.SetBodyMaterialGroup(__entityHandle, (*C.String)(unsafe.Pointer(&__materialGroup)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__materialGroup)
		},
	}.Do()
}

// SetBodyMaterialGroup 
//  @brief Sets the material group of the entity.
//
//  @param entityHandle: The handle of the entity.
//  @param materialGroup: The material group token to assign.
func SetBodyMaterialGroup(entityHandle int32, materialGroup string) {
	_SetBodyMaterialGroup(entityHandle, materialGroup)
}

var _SetBodyVelocity = func(entityHandle int32, velocity plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.SetBodyVelocity(__entityHandle, &__velocity)
}

// SetBodyVelocity 
//  @brief Sets the linear velocity of the entity.
//
//  @param entityHandle: The handle of the entity.
//  @param velocity: The new velocity vector.
func SetBodyVelocity(entityHandle int32, velocity plugify.Vector3) {
	_SetBodyVelocity(entityHandle, velocity)
}

