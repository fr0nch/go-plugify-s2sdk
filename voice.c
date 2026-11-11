#include "shared.h"

PLUGIFY_EXPORT bool (*__s2sdk_HookClientVoiceData)(void*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_UnhookClientVoiceData)(void*) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_VoiceDataGetFormat)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetFormat)(uintptr_t, int32_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_VoiceDataGetAudio)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetAudio)(uintptr_t, Vector*) = NULL;


PLUGIFY_EXPORT uint32_t (*__s2sdk_VoiceDataGetSampleRate)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetSampleRate)(uintptr_t, uint32_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__s2sdk_VoiceDataGetNumPackets)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetNumPackets)(uintptr_t, uint32_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_VoiceDataGetPacketOffsets)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetPacketOffsets)(uintptr_t, Vector*) = NULL;


PLUGIFY_EXPORT float (*__s2sdk_VoiceDataGetVoiceLevel)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_VoiceDataSetVoiceLevel)(uintptr_t, float) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SendVoiceData)(Vector*, int32_t, Vector*, Vector*, uint32_t) = NULL;


