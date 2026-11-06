#include "shared.h"

PLUGIFY_EXPORT void (*__s2sdk_DebugBreak)() = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawBox)(Vector3*, Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawBoxDirection)(Vector3*, Vector3*, Vector3*, Vector3*, Vector3*, float, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawCircle)(Vector3*, Vector3*, float, float, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawClear)() = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawLine)(Vector3*, Vector3*, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawLine_vCol)(Vector3*, Vector3*, Vector3*, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawScreenTextLine)(float, float, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawSphere)(Vector3*, Vector3*, float, float, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawText)(Vector3*, String*, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugScreenTextPretty)(float, float, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float, String*, int32_t, bool) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugScriptAssert)(bool, String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugAxis)(Vector3*, Vector4*, float, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugBox)(Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugBoxAngles)(Vector3*, Vector3*, Vector3*, Vector4*, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCapsule)(Vector3*, Vector4*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCircle)(Vector3*, Vector4*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCircleScreenOriented)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCone)(Vector3*, Vector3*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCross)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCross3D)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugCross3DOriented)(Vector3*, Vector4*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugDrawTickMarkedLine)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugEntityAttachments)(int32_t, float, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugEntityAxis)(int32_t, float, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugEntityBounds)(int32_t, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugEntitySkeleton)(int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugEntityText)(int32_t, int32_t, String*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugFilledRect2D)(Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugHorzArrow)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugLine)(Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugLine2D)(Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugPopDebugOverlayScope)() = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugPushAndClearDebugOverlayScope)(String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugPushDebugOverlayScope)(String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugRemoveAllInScope)(String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugSolidCone)(Vector3*, Vector3*, float, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugSphere)(Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugSweptBox)(Vector3*, Vector3*, Vector3*, Vector3*, Vector4*, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugText)(Vector3*, int32_t, String*, float, int32_t, int32_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugTexture)(String*, Vector2*, Vector2*, int32_t, int32_t, int32_t, int32_t, Vector2*, Vector2*, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugTriangle)(Vector3*, Vector3*, Vector3*, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugVectorText3D)(Vector3*, Vector4*, String*, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugVertArrow)(Vector3*, Vector3*, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DebugYawArrow)(Vector3*, float, float, float, int32_t, int32_t, int32_t, int32_t, bool, float) = NULL;


