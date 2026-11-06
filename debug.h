#pragma once

#include "shared.h"

extern void (*__s2sdk_DebugBreak)();

static void DebugBreak() {
	__s2sdk_DebugBreak();
}

extern void (*__s2sdk_DebugDrawBox)(Vector3*, Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugDrawBox(Vector3* center, Vector3* mins, Vector3* maxs, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugDrawBox(center, mins, maxs, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugDrawBoxDirection)(Vector3*, Vector3*, Vector3*, Vector3*, Vector3*, float, float);

static void DebugDrawBoxDirection(Vector3* center, Vector3* mins, Vector3* maxs, Vector3* forward, Vector3* color, float alpha, float duration) {
	__s2sdk_DebugDrawBoxDirection(center, mins, maxs, forward, color, alpha, duration);
}

extern void (*__s2sdk_DebugDrawCircle)(Vector3*, Vector3*, float, float, bool, float);

static void DebugDrawCircle(Vector3* center, Vector3* color, float alpha, float radius, bool zTest, float duration) {
	__s2sdk_DebugDrawCircle(center, color, alpha, radius, zTest, duration);
}

extern void (*__s2sdk_DebugDrawClear)();

static void DebugDrawClear() {
	__s2sdk_DebugDrawClear();
}

extern void (*__s2sdk_DebugDrawLine)(Vector3*, Vector3*, int32_t, int32_t, int32_t, bool, float);

static void DebugDrawLine(Vector3* origin, Vector3* target, int32_t r, int32_t g, int32_t b, bool zTest, float duration) {
	__s2sdk_DebugDrawLine(origin, target, r, g, b, zTest, duration);
}

extern void (*__s2sdk_DebugDrawLine_vCol)(Vector3*, Vector3*, Vector3*, bool, float);

static void DebugDrawLine_vCol(Vector3* start, Vector3* end, Vector3* color, bool zTest, float duration) {
	__s2sdk_DebugDrawLine_vCol(start, end, color, zTest, duration);
}

extern void (*__s2sdk_DebugDrawScreenTextLine)(float, float, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugDrawScreenTextLine(float x, float y, int32_t lineOffset, String* text, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugDrawScreenTextLine(x, y, lineOffset, text, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugDrawSphere)(Vector3*, Vector3*, float, float, bool, float);

static void DebugDrawSphere(Vector3* center, Vector3* color, float alpha, float radius, bool zTest, float duration) {
	__s2sdk_DebugDrawSphere(center, color, alpha, radius, zTest, duration);
}

extern void (*__s2sdk_DebugDrawText)(Vector3*, String*, bool, float);

static void DebugDrawText(Vector3* origin, String* text, bool viewCheck, float duration) {
	__s2sdk_DebugDrawText(origin, text, viewCheck, duration);
}

extern void (*__s2sdk_DebugScreenTextPretty)(float, float, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float, String*, int32_t, bool);

static void DebugScreenTextPretty(float x, float y, int32_t lineOffset, String* text, int32_t r, int32_t g, int32_t b, int32_t a, float duration, String* font, int32_t size, bool bold) {
	__s2sdk_DebugScreenTextPretty(x, y, lineOffset, text, r, g, b, a, duration, font, size, bold);
}

extern void (*__s2sdk_DebugScriptAssert)(bool, String*);

static void DebugScriptAssert(bool assertion, String* message) {
	__s2sdk_DebugScriptAssert(assertion, message);
}

extern void (*__s2sdk_DebugAxis)(Vector3*, Vector4*, float, bool, float);

static void DebugAxis(Vector3* origin, Vector4* orientation, float scale, bool zTest, float duration) {
	__s2sdk_DebugAxis(origin, orientation, scale, zTest, duration);
}

extern void (*__s2sdk_DebugBox)(Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugBox(Vector3* mins, Vector3* maxs, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugBox(mins, maxs, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugBoxAngles)(Vector3*, Vector3*, Vector3*, Vector4*, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugBoxAngles(Vector3* origin, Vector3* mins, Vector3* maxs, Vector4* orientation, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugBoxAngles(origin, mins, maxs, orientation, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCapsule)(Vector3*, Vector4*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCapsule(Vector3* origin, Vector4* orientation, float radius, float halfHeight, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCapsule(origin, orientation, radius, halfHeight, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCircle)(Vector3*, Vector4*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCircle(Vector3* center, Vector4* orientation, float radius, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCircle(center, orientation, radius, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCircleScreenOriented)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCircleScreenOriented(Vector3* center, float radius, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCircleScreenOriented(center, radius, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCone)(Vector3*, Vector3*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCone(Vector3* origin, Vector3* direction, float radius, float height, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCone(origin, direction, radius, height, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCross)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCross(Vector3* origin, float size, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCross(origin, size, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCross3D)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCross3D(Vector3* origin, float size, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCross3D(origin, size, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugCross3DOriented)(Vector3*, Vector4*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugCross3DOriented(Vector3* origin, Vector4* orientation, float size, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugCross3DOriented(origin, orientation, size, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugDrawTickMarkedLine)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugDrawTickMarkedLine(Vector3* start, Vector3* end, float tickDist, int32_t r, int32_t g, int32_t b, int32_t a, int32_t tickCount, bool zTest, float duration) {
	__s2sdk_DebugDrawTickMarkedLine(start, end, tickDist, r, g, b, a, tickCount, zTest, duration);
}

extern void (*__s2sdk_DebugEntityAttachments)(int32_t, float, float);

static void DebugEntityAttachments(int32_t entityHandle, float size, float duration) {
	__s2sdk_DebugEntityAttachments(entityHandle, size, duration);
}

extern void (*__s2sdk_DebugEntityAxis)(int32_t, float, bool, float);

static void DebugEntityAxis(int32_t entityHandle, float scale, bool zTest, float duration) {
	__s2sdk_DebugEntityAxis(entityHandle, scale, zTest, duration);
}

extern void (*__s2sdk_DebugEntityBounds)(int32_t, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugEntityBounds(int32_t entityHandle, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugEntityBounds(entityHandle, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugEntitySkeleton)(int32_t, float);

static void DebugEntitySkeleton(int32_t entityHandle, float duration) {
	__s2sdk_DebugEntitySkeleton(entityHandle, duration);
}

extern void (*__s2sdk_DebugEntityText)(int32_t, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugEntityText(int32_t entityHandle, int32_t textOffset, String* text, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugEntityText(entityHandle, textOffset, text, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugFilledRect2D)(Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugFilledRect2D(Vector2* min, Vector2* max, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugFilledRect2D(min, max, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugHorzArrow)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugHorzArrow(Vector3* start, Vector3* end, float width, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugHorzArrow(start, end, width, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugLine)(Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugLine(Vector3* start, Vector3* end, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugLine(start, end, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugLine2D)(Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugLine2D(Vector2* start, Vector2* end, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugLine2D(start, end, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugPopDebugOverlayScope)();

static void DebugPopDebugOverlayScope() {
	__s2sdk_DebugPopDebugOverlayScope();
}

extern void (*__s2sdk_DebugPushAndClearDebugOverlayScope)(String*);

static void DebugPushAndClearDebugOverlayScope(String* name) {
	__s2sdk_DebugPushAndClearDebugOverlayScope(name);
}

extern void (*__s2sdk_DebugPushDebugOverlayScope)(String*);

static void DebugPushDebugOverlayScope(String* name) {
	__s2sdk_DebugPushDebugOverlayScope(name);
}

extern void (*__s2sdk_DebugRemoveAllInScope)(String*);

static void DebugRemoveAllInScope(String* name) {
	__s2sdk_DebugRemoveAllInScope(name);
}

extern void (*__s2sdk_DebugSolidCone)(Vector3*, Vector3*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugSolidCone(Vector3* origin, Vector3* direction, float radius, float height, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugSolidCone(origin, direction, radius, height, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugSphere)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugSphere(Vector3* center, float radius, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugSphere(center, radius, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugSweptBox)(Vector3*, Vector3*, Vector3*, Vector3*, Vector4*, int32_t, int32_t, int32_t, int32_t, float);

static void DebugSweptBox(Vector3* start, Vector3* end, Vector3* mins, Vector3* maxs, Vector4* orientation, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugSweptBox(start, end, mins, maxs, orientation, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugText)(Vector3*, int32_t, String*, float, int32_t, int32_t, int32_t, int32_t, float);

static void DebugText(Vector3* origin, int32_t lineOffset, String* text, float scale, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
	__s2sdk_DebugText(origin, lineOffset, text, scale, r, g, b, a, duration);
}

extern void (*__s2sdk_DebugTexture)(String*, Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, Vector2*, Vector2*, float);

static void DebugTexture(String* textureName, Vector2* min, Vector2* max, int32_t r, int32_t g, int32_t b, int32_t a, Vector2* uvMin, Vector2* uvMax, float duration) {
	__s2sdk_DebugTexture(textureName, min, max, r, g, b, a, uvMin, uvMax, duration);
}

extern void (*__s2sdk_DebugTriangle)(Vector3*, Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugTriangle(Vector3* v1, Vector3* v2, Vector3* v3, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugTriangle(v1, v2, v3, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugVectorText3D)(Vector3*, Vector4*, String*, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugVectorText3D(Vector3* origin, Vector4* orientation, String* text, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugVectorText3D(origin, orientation, text, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugVertArrow)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugVertArrow(Vector3* start, Vector3* end, float width, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugVertArrow(start, end, width, r, g, b, a, zTest, duration);
}

extern void (*__s2sdk_DebugYawArrow)(Vector3*, float, float, float, int32_t, int32_t, int32_t, int32_t, bool, float);

static void DebugYawArrow(Vector3* origin, float yaw, float length, float width, int32_t r, int32_t g, int32_t b, int32_t a, bool zTest, float duration) {
	__s2sdk_DebugYawArrow(origin, yaw, length, width, r, g, b, a, zTest, duration);
}

