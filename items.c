#include "shared.h"

PLUGIFY_EXPORT String (*__s2sdk_GetItemDefinitionName)(uint16_t) = NULL;


PLUGIFY_EXPORT uint16_t (*__s2sdk_GetItemDefinitionByName)(String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetItemDefinitionClassName)(uint16_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemDefinitionLoadoutSlot)(uint16_t, int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemDefinitionTeam)(uint16_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemDefinitionStickerSlots)(uint16_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemDefinitionKeychainSlots)(uint16_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetItemDefinitionImage)(uint16_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetItemDefinitionModel)(uint16_t, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetAttributeDefinitionIndex)(String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetAttributeDefinitionName)(int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetPaintKitName)(int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetPaintKitLocName)(int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_FindPaintKit)(String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetStickerKitName)(int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetStickerKitLocName)(int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_FindStickerKit)(String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetKeychainName)(int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetKeychainLocName)(int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_FindKeychain)(String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_GetClientItemView)(int32_t, uint64_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_GetClientLoadoutItemView)(int32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_GetEntityItemView)(int32_t) = NULL;


PLUGIFY_EXPORT uint16_t (*__s2sdk_GetItemViewDefinition)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__s2sdk_GetItemViewId)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__s2sdk_GetItemViewAccountId)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemViewQuality)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemViewRarity)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemViewOrigin)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsItemViewTradable)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsItemViewMarketable)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemViewPaintKit)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetItemViewPaintKitSeed)(uintptr_t) = NULL;


PLUGIFY_EXPORT float (*__s2sdk_GetItemViewPaintKitWear)(uintptr_t) = NULL;


PLUGIFY_EXPORT double (*__s2sdk_GetItemViewStickerAttribute)(uintptr_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT double (*__s2sdk_GetItemViewKeychainAttribute)(uintptr_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetItemViewAttribute)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GiveClientItem)(int32_t, uint64_t) = NULL;


