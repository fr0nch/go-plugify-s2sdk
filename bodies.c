#include "shared.h"

PLUGIFY_EXPORT void (*__s2sdk_AddBodyImpulseAtPosition)(int32_t, Vector3*, Vector3*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_AddBodyVelocity)(int32_t, Vector3*, Vector3*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_DetachBodyFromParent)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsBodyAttachedToParent)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_SetBodyAngularVelocity)(int32_t, Vector3*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_SetBodyMaterialGroup)(int32_t, String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_SetBodyVelocity)(int32_t, Vector3*) = NULL;


