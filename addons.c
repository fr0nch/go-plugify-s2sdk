#include "shared.h"

PLUGIFY_EXPORT bool (*__s2sdk_AddServerAddon)(uint64_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveServerAddon)(uint64_t, bool) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetServerAddons)() = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_AddPrivateAddon)(uint64_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemovePrivateAddon)(uint64_t, bool) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetPrivateAddons)() = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_AddClientOnlyAddon)(uint64_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveClientOnlyAddon)(uint64_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetClientOnlyAddons)() = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SendAddon)(uint64_t, uint64_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveSentAddon)(uint64_t, uint64_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_ClearSentAddons)(uint64_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetSentAddons)(uint64_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetClientAddons)(uint64_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__s2sdk_GetCurrentWorkshopMap)() = NULL;


