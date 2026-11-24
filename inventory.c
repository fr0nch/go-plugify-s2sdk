#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__s2sdk_GetClientInventory)(int32_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__s2sdk_GetClientLoadoutItemId)(int32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT uint16_t (*__s2sdk_GetClientLoadoutItemDefinition)(int32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientLoadoutItem)(int32_t, int32_t, int32_t, uint64_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientLoadoutDefinition)(int32_t, int32_t, int32_t, uint16_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_FindClientLoadoutSlot)(int32_t, int32_t, uint16_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RestoreClientLoadout)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_GiveClientDefaultWearables)(int32_t) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetClientInventoryItems)(int32_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__s2sdk_CreateClientItem)(int32_t, uint16_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveClientItem)(int32_t, uint64_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemAttribute)(int32_t, uint64_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemPaint)(int32_t, uint64_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemSticker)(int32_t, uint64_t, int32_t, int32_t, float) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemKeychain)(int32_t, uint64_t, int32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemStatTrak)(int32_t, uint64_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetClientItemName)(int32_t, uint64_t, String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetClientItemAttribute)(int32_t, uint64_t, String*) = NULL;


