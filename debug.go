package s2sdk

/*
#include "debug.h"
#cgo noescape DebugBreak
#cgo noescape DebugDrawBox
#cgo noescape DebugDrawBoxDirection
#cgo noescape DebugDrawCircle
#cgo noescape DebugDrawClear
#cgo noescape DebugDrawLine
#cgo noescape DebugDrawLine_vCol
#cgo noescape DebugDrawScreenTextLine
#cgo noescape DebugDrawSphere
#cgo noescape DebugDrawText
#cgo noescape DebugScreenTextPretty
#cgo noescape DebugScriptAssert
#cgo noescape DebugAxis
#cgo noescape DebugBox
#cgo noescape DebugBoxAngles
#cgo noescape DebugCapsule
#cgo noescape DebugCircle
#cgo noescape DebugCircleScreenOriented
#cgo noescape DebugCone
#cgo noescape DebugCross
#cgo noescape DebugCross3D
#cgo noescape DebugCross3DOriented
#cgo noescape DebugDrawTickMarkedLine
#cgo noescape DebugEntityAttachments
#cgo noescape DebugEntityAxis
#cgo noescape DebugEntityBounds
#cgo noescape DebugEntitySkeleton
#cgo noescape DebugEntityText
#cgo noescape DebugFilledRect2D
#cgo noescape DebugHorzArrow
#cgo noescape DebugLine
#cgo noescape DebugLine2D
#cgo noescape DebugPopDebugOverlayScope
#cgo noescape DebugPushAndClearDebugOverlayScope
#cgo noescape DebugPushDebugOverlayScope
#cgo noescape DebugRemoveAllInScope
#cgo noescape DebugSolidCone
#cgo noescape DebugSphere
#cgo noescape DebugSweptBox
#cgo noescape DebugText
#cgo noescape DebugTexture
#cgo noescape DebugTriangle
#cgo noescape DebugVectorText3D
#cgo noescape DebugVertArrow
#cgo noescape DebugYawArrow
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

// Generated from s2sdk (group: debug)

var _DebugBreak = func() {
	C.DebugBreak()
}

// DebugBreak 
//  @brief Triggers a breakpoint in the debugger.
//
func DebugBreak() {
	_DebugBreak()
}

var _DebugDrawBox = func(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, r int32, g int32, b int32, a int32, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__mins := *(*C.Vector3)(unsafe.Pointer(&mins))
	__maxs := *(*C.Vector3)(unsafe.Pointer(&maxs))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	C.DebugDrawBox(&__center, &__mins, &__maxs, __r, __g, __b, __a, __duration)
}

// DebugDrawBox 
//  @brief Draws a debug overlay box.
//
//  @param center: Center of the box in world space.
//  @param mins: Minimum bounds relative to the center.
//  @param maxs: Maximum bounds relative to the center.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display the box.
func DebugDrawBox(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, r int32, g int32, b int32, a int32, duration float32) {
	_DebugDrawBox(center, mins, maxs, r, g, b, a, duration)
}

var _DebugDrawBoxDirection = func(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, forward plugify.Vector3, color plugify.Vector3, alpha float32, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__mins := *(*C.Vector3)(unsafe.Pointer(&mins))
	__maxs := *(*C.Vector3)(unsafe.Pointer(&maxs))
	__forward := *(*C.Vector3)(unsafe.Pointer(&forward))
	__color := *(*C.Vector3)(unsafe.Pointer(&color))
	__alpha := C.float(alpha)
	__duration := C.float(duration)
	C.DebugDrawBoxDirection(&__center, &__mins, &__maxs, &__forward, &__color, __alpha, __duration)
}

// DebugDrawBoxDirection 
//  @brief Draws a debug box oriented in the direction of a forward vector.
//
//  @param center: Center of the box.
//  @param mins: Minimum bounds.
//  @param maxs: Maximum bounds.
//  @param forward: Forward direction vector.
//  @param color: RGB color vector.
//  @param alpha: Alpha transparency.
//  @param duration: Duration (in seconds) to display the box.
func DebugDrawBoxDirection(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, forward plugify.Vector3, color plugify.Vector3, alpha float32, duration float32) {
	_DebugDrawBoxDirection(center, mins, maxs, forward, color, alpha, duration)
}

var _DebugDrawCircle = func(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__color := *(*C.Vector3)(unsafe.Pointer(&color))
	__alpha := C.float(alpha)
	__radius := C.float(radius)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugDrawCircle(&__center, &__color, __alpha, __radius, __zTest, __duration)
}

// DebugDrawCircle 
//  @brief Draws a debug circle.
//
//  @param center: Center of the circle.
//  @param color: RGB color vector.
//  @param alpha: Alpha transparency.
//  @param radius: Circle radius.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the circle.
func DebugDrawCircle(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32) {
	_DebugDrawCircle(center, color, alpha, radius, zTest, duration)
}

var _DebugDrawClear = func() {
	C.DebugDrawClear()
}

// DebugDrawClear 
//  @brief Clears all debug overlays.
//
func DebugDrawClear() {
	_DebugDrawClear()
}

var _DebugDrawLine = func(origin plugify.Vector3, target plugify.Vector3, r int32, g int32, b int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__target := *(*C.Vector3)(unsafe.Pointer(&target))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugDrawLine(&__origin, &__target, __r, __g, __b, __zTest, __duration)
}

// DebugDrawLine 
//  @brief Draws a debug overlay line.
//
//  @param origin: Start point of the line.
//  @param target: End point of the line.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the line.
func DebugDrawLine(origin plugify.Vector3, target plugify.Vector3, r int32, g int32, b int32, zTest bool, duration float32) {
	_DebugDrawLine(origin, target, r, g, b, zTest, duration)
}

var _DebugDrawLine_vCol = func(start plugify.Vector3, end plugify.Vector3, color plugify.Vector3, zTest bool, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__color := *(*C.Vector3)(unsafe.Pointer(&color))
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugDrawLine_vCol(&__start, &__end, &__color, __zTest, __duration)
}

// DebugDrawLine_vCol 
//  @brief Draws a debug line using a color vector.
//
//  @param start: Start point of the line.
//  @param end: End point of the line.
//  @param color: RGB color vector.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the line.
func DebugDrawLine_vCol(start plugify.Vector3, end plugify.Vector3, color plugify.Vector3, zTest bool, duration float32) {
	_DebugDrawLine_vCol(start, end, color, zTest, duration)
}

var _DebugDrawScreenTextLine = func(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32) {
	__x := C.float(x)
	__y := C.float(y)
	__lineOffset := C.int32_t(lineOffset)
	__text := plugify.ConstructString(text)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugDrawScreenTextLine(__x, __y, __lineOffset, (*C.String)(unsafe.Pointer(&__text)), __r, __g, __b, __a, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// DebugDrawScreenTextLine 
//  @brief Draws text at a specified screen position with line offset.
//
//  @param x: X coordinate in screen space.
//  @param y: Y coordinate in screen space.
//  @param lineOffset: Line offset value.
//  @param text: The text string to display.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha transparency value.
//  @param duration: Duration (in seconds) to display the text.
func DebugDrawScreenTextLine(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32) {
	_DebugDrawScreenTextLine(x, y, lineOffset, text, r, g, b, a, duration)
}

var _DebugDrawSphere = func(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__color := *(*C.Vector3)(unsafe.Pointer(&color))
	__alpha := C.float(alpha)
	__radius := C.float(radius)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugDrawSphere(&__center, &__color, __alpha, __radius, __zTest, __duration)
}

// DebugDrawSphere 
//  @brief Draws a debug sphere.
//
//  @param center: Center of the sphere.
//  @param color: RGB color vector.
//  @param alpha: Alpha transparency.
//  @param radius: Radius of the sphere.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the sphere.
func DebugDrawSphere(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32) {
	_DebugDrawSphere(center, color, alpha, radius, zTest, duration)
}

var _DebugDrawText = func(origin plugify.Vector3, text string, viewCheck bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__text := plugify.ConstructString(text)
	__viewCheck := C.bool(viewCheck)
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugDrawText(&__origin, (*C.String)(unsafe.Pointer(&__text)), __viewCheck, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// DebugDrawText 
//  @brief Draws text in 3D space.
//
//  @param origin: World-space position to draw the text at.
//  @param text: The text string to display.
//  @param viewCheck: If true, only draws when visible to camera.
//  @param duration: Duration (in seconds) to display the text.
func DebugDrawText(origin plugify.Vector3, text string, viewCheck bool, duration float32) {
	_DebugDrawText(origin, text, viewCheck, duration)
}

var _DebugScreenTextPretty = func(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32, font string, size int32, bold bool) {
	__x := C.float(x)
	__y := C.float(y)
	__lineOffset := C.int32_t(lineOffset)
	__text := plugify.ConstructString(text)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	__font := plugify.ConstructString(font)
	__size := C.int32_t(size)
	__bold := C.bool(bold)
	plugify.Block {
		Try: func() {
			C.DebugScreenTextPretty(__x, __y, __lineOffset, (*C.String)(unsafe.Pointer(&__text)), __r, __g, __b, __a, __duration, (*C.String)(unsafe.Pointer(&__font)), __size, __bold)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
			plugify.DestroyString(&__font)
		},
	}.Do()
}

// DebugScreenTextPretty 
//  @brief Draws styled debug text on screen.
//
//  @param x: X coordinate.
//  @param y: Y coordinate.
//  @param lineOffset: Line offset value.
//  @param text: Text string.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha transparency.
//  @param duration: Duration (in seconds) to display the text.
//  @param font: Font name.
//  @param size: Font size.
//  @param bold: Whether text should be bold.
func DebugScreenTextPretty(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32, font string, size int32, bold bool) {
	_DebugScreenTextPretty(x, y, lineOffset, text, r, g, b, a, duration, font, size, bold)
}

var _DebugScriptAssert = func(assertion bool, message string) {
	__assertion := C.bool(assertion)
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			C.DebugScriptAssert(__assertion, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// DebugScriptAssert 
//  @brief Performs an assertion and logs a message if the assertion fails.
//
//  @param assertion: Boolean value to test.
//  @param message: Message to display if the assertion fails.
func DebugScriptAssert(assertion bool, message string) {
	_DebugScriptAssert(assertion, message)
}

var _DebugAxis = func(origin plugify.Vector3, orientation plugify.Vector4, scale float32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__scale := C.float(scale)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugAxis(&__origin, &__orientation, __scale, __zTest, __duration)
}

// DebugAxis 
//  @brief Draws an axis at a specified world-space origin.
//
//  @param origin: The origin of the axis in world space.
//  @param orientation: The orientation as a quaternion.
//  @param scale: The scale of the axis.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the axis.
func DebugAxis(origin plugify.Vector3, orientation plugify.Vector4, scale float32, zTest bool, duration float32) {
	_DebugAxis(origin, orientation, scale, zTest, duration)
}

var _DebugBox = func(mins plugify.Vector3, maxs plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__mins := *(*C.Vector3)(unsafe.Pointer(&mins))
	__maxs := *(*C.Vector3)(unsafe.Pointer(&maxs))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugBox(&__mins, &__maxs, __r, __g, __b, __a, __zTest, __duration)
}

// DebugBox 
//  @brief Draws a world-space axis-aligned box.
//
//  @param mins: The minimum corner of the box in world space.
//  @param maxs: The maximum corner of the box in world space.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the box.
func DebugBox(mins plugify.Vector3, maxs plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugBox(mins, maxs, r, g, b, a, zTest, duration)
}

var _DebugBoxAngles = func(origin plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, orientation plugify.Vector4, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__mins := *(*C.Vector3)(unsafe.Pointer(&mins))
	__maxs := *(*C.Vector3)(unsafe.Pointer(&maxs))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugBoxAngles(&__origin, &__mins, &__maxs, &__orientation, __r, __g, __b, __a, __zTest, __duration)
}

// DebugBoxAngles 
//  @brief Draws an oriented box in world space.
//
//  @param origin: The box origin in world space.
//  @param mins: The local-space minimum corner.
//  @param maxs: The local-space maximum corner.
//  @param orientation: The box orientation as a quaternion.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the box.
func DebugBoxAngles(origin plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, orientation plugify.Vector4, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugBoxAngles(origin, mins, maxs, orientation, r, g, b, a, zTest, duration)
}

var _DebugCapsule = func(origin plugify.Vector3, orientation plugify.Vector4, radius float32, halfHeight float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__radius := C.float(radius)
	__halfHeight := C.float(halfHeight)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCapsule(&__origin, &__orientation, __radius, __halfHeight, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCapsule 
//  @brief Draws a capsule in world space.
//
//  @param origin: The capsule origin in world space.
//  @param orientation: The capsule orientation as a quaternion.
//  @param radius: The capsule radius.
//  @param halfHeight: Half the capsule height.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the capsule.
func DebugCapsule(origin plugify.Vector3, orientation plugify.Vector4, radius float32, halfHeight float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCapsule(origin, orientation, radius, halfHeight, r, g, b, a, zTest, duration)
}

var _DebugCircle = func(center plugify.Vector3, orientation plugify.Vector4, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__radius := C.float(radius)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCircle(&__center, &__orientation, __radius, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCircle 
//  @brief Draws a circle in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param center: The center of the circle in world space.
//  @param orientation: The orientation as a quaternion.
//  @param radius: The circle radius.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the circle.
func DebugCircle(center plugify.Vector3, orientation plugify.Vector4, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCircle(center, orientation, radius, r, g, b, a, zTest, duration)
}

var _DebugCircleScreenOriented = func(center plugify.Vector3, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__radius := C.float(radius)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCircleScreenOriented(&__center, __radius, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCircleScreenOriented 
//  @brief Draws a screen-oriented circle in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param center: The center of the circle in world space.
//  @param radius: The radius of the circle.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the circle.
func DebugCircleScreenOriented(center plugify.Vector3, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCircleScreenOriented(center, radius, r, g, b, a, zTest, duration)
}

var _DebugCone = func(origin plugify.Vector3, direction plugify.Vector3, radius float32, height float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__direction := *(*C.Vector3)(unsafe.Pointer(&direction))
	__radius := C.float(radius)
	__height := C.float(height)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCone(&__origin, &__direction, __radius, __height, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCone 
//  @brief Draws a wireframe cone in world space.
//
//  @param origin: The base position of the cone in world space.
//  @param direction: The direction vector of the cone in world space.
//  @param radius: The cone base radius.
//  @param height: The cone height.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the cone.
func DebugCone(origin plugify.Vector3, direction plugify.Vector3, radius float32, height float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCone(origin, direction, radius, height, r, g, b, a, zTest, duration)
}

var _DebugCross = func(origin plugify.Vector3, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__size := C.float(size)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCross(&__origin, __size, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCross 
//  @brief Draws a screen-aligned cross in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param origin: The center position of the cross in world space.
//  @param size: The size of the cross.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the cross.
func DebugCross(origin plugify.Vector3, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCross(origin, size, r, g, b, a, zTest, duration)
}

var _DebugCross3D = func(origin plugify.Vector3, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__size := C.float(size)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCross3D(&__origin, __size, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCross3D 
//  @brief Draws a world-aligned cross in world space.
//
//  @param origin: The center of the cross in world space.
//  @param size: The size of the cross.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the cross.
func DebugCross3D(origin plugify.Vector3, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCross3D(origin, size, r, g, b, a, zTest, duration)
}

var _DebugCross3DOriented = func(origin plugify.Vector3, orientation plugify.Vector4, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__size := C.float(size)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugCross3DOriented(&__origin, &__orientation, __size, __r, __g, __b, __a, __zTest, __duration)
}

// DebugCross3DOriented 
//  @brief Draws an oriented cross in world space.
//
//  @param origin: The center position of the cross.
//  @param orientation: The orientation of the cross as a quaternion.
//  @param size: The size of the cross.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the cross.
func DebugCross3DOriented(origin plugify.Vector3, orientation plugify.Vector4, size float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugCross3DOriented(origin, orientation, size, r, g, b, a, zTest, duration)
}

var _DebugDrawTickMarkedLine = func(start plugify.Vector3, end plugify.Vector3, tickDist float32, r int32, g int32, b int32, a int32, tickCount int32, zTest bool, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__tickDist := C.float(tickDist)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__tickCount := C.int32_t(tickCount)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugDrawTickMarkedLine(&__start, &__end, __tickDist, __r, __g, __b, __a, __tickCount, __zTest, __duration)
}

// DebugDrawTickMarkedLine 
//  @brief Draws a dashed (tick-marked) line in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param start: The start point in world space.
//  @param end: The end point in world space.
//  @param tickDist: Distance between tick marks.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param tickCount: tickCount
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the line.
func DebugDrawTickMarkedLine(start plugify.Vector3, end plugify.Vector3, tickDist float32, r int32, g int32, b int32, a int32, tickCount int32, zTest bool, duration float32) {
	_DebugDrawTickMarkedLine(start, end, tickDist, r, g, b, a, tickCount, zTest, duration)
}

var _DebugEntityAttachments = func(entityHandle int32, size float32, duration float32) {
	__entityHandle := C.int32_t(entityHandle)
	__size := C.float(size)
	__duration := C.float(duration)
	C.DebugEntityAttachments(__entityHandle, __size, __duration)
}

// DebugEntityAttachments 
//  @brief Draws the attachments of an entity.
//
//  @param entityHandle: Entity handle.
//  @param size: Size of attachment markers.
//  @param duration: Duration (in seconds) to display.
func DebugEntityAttachments(entityHandle int32, size float32, duration float32) {
	_DebugEntityAttachments(entityHandle, size, duration)
}

var _DebugEntityAxis = func(entityHandle int32, scale float32, zTest bool, duration float32) {
	__entityHandle := C.int32_t(entityHandle)
	__scale := C.float(scale)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugEntityAxis(__entityHandle, __scale, __zTest, __duration)
}

// DebugEntityAxis 
//  @brief Draws the local axis of an entity at its origin.
//
//  @param entityHandle: Entity handle.
//  @param scale: Axis size.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display.
func DebugEntityAxis(entityHandle int32, scale float32, zTest bool, duration float32) {
	_DebugEntityAxis(entityHandle, scale, zTest, duration)
}

var _DebugEntityBounds = func(entityHandle int32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__entityHandle := C.int32_t(entityHandle)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugEntityBounds(__entityHandle, __r, __g, __b, __a, __zTest, __duration)
}

// DebugEntityBounds 
//  @brief Draws an entity’s world-space bounding box.
//
//  @param entityHandle: Entity handle.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display.
func DebugEntityBounds(entityHandle int32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugEntityBounds(entityHandle, r, g, b, a, zTest, duration)
}

var _DebugEntitySkeleton = func(entityHandle int32, duration float32) {
	__entityHandle := C.int32_t(entityHandle)
	__duration := C.float(duration)
	C.DebugEntitySkeleton(__entityHandle, __duration)
}

// DebugEntitySkeleton 
//  @brief Draws the skeleton of a model for a given entity. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param entityHandle: Entity handle.
//  @param duration: Duration (in seconds) to display.
func DebugEntitySkeleton(entityHandle int32, duration float32) {
	_DebugEntitySkeleton(entityHandle, duration)
}

var _DebugEntityText = func(entityHandle int32, textOffset int32, text string, r int32, g int32, b int32, a int32, duration float32) {
	__entityHandle := C.int32_t(entityHandle)
	__textOffset := C.int32_t(textOffset)
	__text := plugify.ConstructString(text)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugEntityText(__entityHandle, __textOffset, (*C.String)(unsafe.Pointer(&__text)), __r, __g, __b, __a, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// DebugEntityText 
//  @brief Draws text attached to an entity.
//
//  @param entityHandle: Entity handle.
//  @param textOffset: Offset index for text position.
//  @param text: Text to display.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display.
func DebugEntityText(entityHandle int32, textOffset int32, text string, r int32, g int32, b int32, a int32, duration float32) {
	_DebugEntityText(entityHandle, textOffset, text, r, g, b, a, duration)
}

var _DebugFilledRect2D = func(min plugify.Vector2, max plugify.Vector2, r int32, g int32, b int32, a int32, duration float32) {
	__min := *(*C.Vector2)(unsafe.Pointer(&min))
	__max := *(*C.Vector2)(unsafe.Pointer(&max))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	C.DebugFilledRect2D(&__min, &__max, __r, __g, __b, __a, __duration)
}

// DebugFilledRect2D 
//  @brief Draws a filled 2D rectangle in screen space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param min: The top-left corner in pixels.
//  @param max: The bottom-right corner in pixels.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display.
func DebugFilledRect2D(min plugify.Vector2, max plugify.Vector2, r int32, g int32, b int32, a int32, duration float32) {
	_DebugFilledRect2D(min, max, r, g, b, a, duration)
}

var _DebugHorzArrow = func(start plugify.Vector3, end plugify.Vector3, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__width := C.float(width)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugHorzArrow(&__start, &__end, __width, __r, __g, __b, __a, __zTest, __duration)
}

// DebugHorzArrow 
//  @brief Draws a horizontal arrow between two points in world space.
//
//  @param start: The start point in world space.
//  @param end: The end point in world space.
//  @param width: Arrow width.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display.
func DebugHorzArrow(start plugify.Vector3, end plugify.Vector3, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugHorzArrow(start, end, width, r, g, b, a, zTest, duration)
}

var _DebugLine = func(start plugify.Vector3, end plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugLine(&__start, &__end, __r, __g, __b, __a, __zTest, __duration)
}

// DebugLine 
//  @brief Draws a 3D line between two world-space points.
//
//  @param start: The start point.
//  @param end: The end point.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display.
func DebugLine(start plugify.Vector3, end plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugLine(start, end, r, g, b, a, zTest, duration)
}

var _DebugLine2D = func(start plugify.Vector2, end plugify.Vector2, r int32, g int32, b int32, a int32, duration float32) {
	__start := *(*C.Vector2)(unsafe.Pointer(&start))
	__end := *(*C.Vector2)(unsafe.Pointer(&end))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	C.DebugLine2D(&__start, &__end, __r, __g, __b, __a, __duration)
}

// DebugLine2D 
//  @brief Draws a line between two points in screen space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param start: The start position in pixels.
//  @param end: The end position in pixels.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display.
func DebugLine2D(start plugify.Vector2, end plugify.Vector2, r int32, g int32, b int32, a int32, duration float32) {
	_DebugLine2D(start, end, r, g, b, a, duration)
}

var _DebugPopDebugOverlayScope = func() {
	C.DebugPopDebugOverlayScope()
}

// DebugPopDebugOverlayScope 
//  @brief Pops the current debug overlay scope, ending a grouped overlay batch.
//
func DebugPopDebugOverlayScope() {
	_DebugPopDebugOverlayScope()
}

var _DebugPushAndClearDebugOverlayScope = func(name string) {
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.DebugPushAndClearDebugOverlayScope((*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// DebugPushAndClearDebugOverlayScope 
//  @brief Pushes a new debug overlay scope and clears all overlays with this identifier.
//
//  @param name: Overlay scope name.
func DebugPushAndClearDebugOverlayScope(name string) {
	_DebugPushAndClearDebugOverlayScope(name)
}

var _DebugPushDebugOverlayScope = func(name string) {
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.DebugPushDebugOverlayScope((*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// DebugPushDebugOverlayScope 
//  @brief Pushes a new debug overlay scope without clearing.
//
//  @param name: Overlay scope name.
func DebugPushDebugOverlayScope(name string) {
	_DebugPushDebugOverlayScope(name)
}

var _DebugRemoveAllInScope = func(name string) {
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			C.DebugRemoveAllInScope((*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// DebugRemoveAllInScope 
//  @brief Removes all debug overlays marked with a specific identifier.
//
//  @param name: Overlay scope name.
func DebugRemoveAllInScope(name string) {
	_DebugRemoveAllInScope(name)
}

var _DebugSolidCone = func(origin plugify.Vector3, direction plugify.Vector3, radius float32, height float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__direction := *(*C.Vector3)(unsafe.Pointer(&direction))
	__radius := C.float(radius)
	__height := C.float(height)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugSolidCone(&__origin, &__direction, __radius, __height, __r, __g, __b, __a, __zTest, __duration)
}

// DebugSolidCone 
//  @brief Draws a solid cone in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param origin: Base position of the cone in world space.
//  @param direction: Direction vector of the cone in world space.
//  @param radius: Base radius of the cone.
//  @param height: Cone height.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the cone.
func DebugSolidCone(origin plugify.Vector3, direction plugify.Vector3, radius float32, height float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugSolidCone(origin, direction, radius, height, r, g, b, a, zTest, duration)
}

var _DebugSphere = func(center plugify.Vector3, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__center := *(*C.Vector3)(unsafe.Pointer(&center))
	__radius := C.float(radius)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugSphere(&__center, __radius, __r, __g, __b, __a, __zTest, __duration)
}

// DebugSphere 
//  @brief Draws a wireframe sphere in world space.
//
//  @param center: Center of the sphere in world space.
//  @param radius: Sphere radius.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the sphere.
func DebugSphere(center plugify.Vector3, radius float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugSphere(center, radius, r, g, b, a, zTest, duration)
}

var _DebugSweptBox = func(start plugify.Vector3, end plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, orientation plugify.Vector4, r int32, g int32, b int32, a int32, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__mins := *(*C.Vector3)(unsafe.Pointer(&mins))
	__maxs := *(*C.Vector3)(unsafe.Pointer(&maxs))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	C.DebugSweptBox(&__start, &__end, &__mins, &__maxs, &__orientation, __r, __g, __b, __a, __duration)
}

// DebugSweptBox 
//  @brief Draws a swept box between two points in world space.
//
//  @param start: Start point in world space.
//  @param end: End point in world space.
//  @param mins: Local-space minimum bounds.
//  @param maxs: Local-space maximum bounds.
//  @param orientation: Orientation of the box.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display the swept box.
func DebugSweptBox(start plugify.Vector3, end plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, orientation plugify.Vector4, r int32, g int32, b int32, a int32, duration float32) {
	_DebugSweptBox(start, end, mins, maxs, orientation, r, g, b, a, duration)
}

var _DebugText = func(origin plugify.Vector3, lineOffset int32, text string, scale float32, r int32, g int32, b int32, a int32, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__lineOffset := C.int32_t(lineOffset)
	__text := plugify.ConstructString(text)
	__scale := C.float(scale)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugText(&__origin, __lineOffset, (*C.String)(unsafe.Pointer(&__text)), __scale, __r, __g, __b, __a, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// DebugText 
//  @brief Draws 2D text in world space.
//
//  @param origin: Origin of the text in world space.
//  @param lineOffset: Line offset index.
//  @param text: Text to display.
//  @param scale: Text scale factor.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param duration: Duration (in seconds) to display.
func DebugText(origin plugify.Vector3, lineOffset int32, text string, scale float32, r int32, g int32, b int32, a int32, duration float32) {
	_DebugText(origin, lineOffset, text, scale, r, g, b, a, duration)
}

var _DebugTexture = func(textureName string, min plugify.Vector2, max plugify.Vector2, r int32, g int32, b int32, a int32, uvMin plugify.Vector2, uvMax plugify.Vector2, duration float32) {
	__textureName := plugify.ConstructString(textureName)
	__min := *(*C.Vector2)(unsafe.Pointer(&min))
	__max := *(*C.Vector2)(unsafe.Pointer(&max))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__uvMin := *(*C.Vector2)(unsafe.Pointer(&uvMin))
	__uvMax := *(*C.Vector2)(unsafe.Pointer(&uvMax))
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugTexture((*C.String)(unsafe.Pointer(&__textureName)), &__min, &__max, __r, __g, __b, __a, &__uvMin, &__uvMax, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__textureName)
		},
	}.Do()
}

// DebugTexture 
//  @brief Draws a textured rectangle on screen.
//
//  @param textureName: Name of the texture to draw.
//  @param min: Top-left screen-space coordinate in pixels.
//  @param max: Bottom-right screen-space coordinate in pixels.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param uvMin: Minimum UV coordinate.
//  @param uvMax: Maximum UV coordinate.
//  @param duration: Duration (in seconds) to display the texture.
func DebugTexture(textureName string, min plugify.Vector2, max plugify.Vector2, r int32, g int32, b int32, a int32, uvMin plugify.Vector2, uvMax plugify.Vector2, duration float32) {
	_DebugTexture(textureName, min, max, r, g, b, a, uvMin, uvMax, duration)
}

var _DebugTriangle = func(v1 plugify.Vector3, v2 plugify.Vector3, v3 plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__v1 := *(*C.Vector3)(unsafe.Pointer(&v1))
	__v2 := *(*C.Vector3)(unsafe.Pointer(&v2))
	__v3 := *(*C.Vector3)(unsafe.Pointer(&v3))
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugTriangle(&__v1, &__v2, &__v3, __r, __g, __b, __a, __zTest, __duration)
}

// DebugTriangle 
//  @brief Draws a filled triangle in world space. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param v1: First vertex.
//  @param v2: Second vertex.
//  @param v3: Third vertex.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the triangle.
func DebugTriangle(v1 plugify.Vector3, v2 plugify.Vector3, v3 plugify.Vector3, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugTriangle(v1, v2, v3, r, g, b, a, zTest, duration)
}

var _DebugVectorText3D = func(origin plugify.Vector3, orientation plugify.Vector4, text string, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__orientation := *(*C.Vector4)(unsafe.Pointer(&orientation))
	__text := plugify.ConstructString(text)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	plugify.Block {
		Try: func() {
			C.DebugVectorText3D(&__origin, &__orientation, (*C.String)(unsafe.Pointer(&__text)), __r, __g, __b, __a, __zTest, __duration)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__text)
		},
	}.Do()
}

// DebugVectorText3D 
//  @brief Draws 3D text in world space with orientation. Not displayed on clients of a dedicated server: the server does not send this primitive.
//
//  @param origin: Origin position in world space.
//  @param orientation: Orientation of the text.
//  @param text: Text to display.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the text.
func DebugVectorText3D(origin plugify.Vector3, orientation plugify.Vector4, text string, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugVectorText3D(origin, orientation, text, r, g, b, a, zTest, duration)
}

var _DebugVertArrow = func(start plugify.Vector3, end plugify.Vector3, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__start := *(*C.Vector3)(unsafe.Pointer(&start))
	__end := *(*C.Vector3)(unsafe.Pointer(&end))
	__width := C.float(width)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugVertArrow(&__start, &__end, __width, __r, __g, __b, __a, __zTest, __duration)
}

// DebugVertArrow 
//  @brief Draws a vertical arrow between two points in world space.
//
//  @param start: Start point in world space.
//  @param end: End point in world space.
//  @param width: Arrow width.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display.
func DebugVertArrow(start plugify.Vector3, end plugify.Vector3, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugVertArrow(start, end, width, r, g, b, a, zTest, duration)
}

var _DebugYawArrow = func(origin plugify.Vector3, yaw float32, length float32, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__yaw := C.float(yaw)
	__length := C.float(length)
	__width := C.float(width)
	__r := C.int32_t(r)
	__g := C.int32_t(g)
	__b := C.int32_t(b)
	__a := C.int32_t(a)
	__zTest := C.bool(zTest)
	__duration := C.float(duration)
	C.DebugYawArrow(&__origin, __yaw, __length, __width, __r, __g, __b, __a, __zTest, __duration)
}

// DebugYawArrow 
//  @brief Draws a yaw-based directional arrow in world space.
//
//  @param origin: Origin position in world space.
//  @param yaw: Yaw angle in degrees.
//  @param length: Length of the arrow.
//  @param width: Arrow width.
//  @param r: Red color value.
//  @param g: Green color value.
//  @param b: Blue color value.
//  @param a: Alpha (transparency) value.
//  @param zTest: Whether to perform depth testing.
//  @param duration: Duration (in seconds) to display the arrow.
func DebugYawArrow(origin plugify.Vector3, yaw float32, length float32, width float32, r int32, g int32, b int32, a int32, zTest bool, duration float32) {
	_DebugYawArrow(origin, yaw, length, width, r, g, b, a, zTest, duration)
}

