#pragma once

#include "shared.h"

extern uintptr_t (*__s2sdk_GetClientInventory)(int32_t);

static uintptr_t GetClientInventory(int32_t playerSlot) {
	return __s2sdk_GetClientInventory(playerSlot);
}

extern uint64_t (*__s2sdk_GetClientLoadoutItemId)(int32_t, int32_t, int32_t);

static uint64_t GetClientLoadoutItemId(int32_t playerSlot, int32_t team, int32_t loadoutSlot) {
	return __s2sdk_GetClientLoadoutItemId(playerSlot, team, loadoutSlot);
}

extern uint16_t (*__s2sdk_GetClientLoadoutItemDefinition)(int32_t, int32_t, int32_t);

static uint16_t GetClientLoadoutItemDefinition(int32_t playerSlot, int32_t team, int32_t loadoutSlot) {
	return __s2sdk_GetClientLoadoutItemDefinition(playerSlot, team, loadoutSlot);
}

extern bool (*__s2sdk_SetClientLoadoutItem)(int32_t, int32_t, int32_t, uint64_t);

static bool SetClientLoadoutItem(int32_t playerSlot, int32_t team, int32_t loadoutSlot, uint64_t itemId) {
	return __s2sdk_SetClientLoadoutItem(playerSlot, team, loadoutSlot, itemId);
}

extern bool (*__s2sdk_SetClientLoadoutDefinition)(int32_t, int32_t, int32_t, uint16_t);

static bool SetClientLoadoutDefinition(int32_t playerSlot, int32_t team, int32_t loadoutSlot, uint16_t itemDefinition) {
	return __s2sdk_SetClientLoadoutDefinition(playerSlot, team, loadoutSlot, itemDefinition);
}

extern int32_t (*__s2sdk_FindClientLoadoutSlot)(int32_t, int32_t, uint16_t);

static int32_t FindClientLoadoutSlot(int32_t playerSlot, int32_t team, uint16_t itemDefinition) {
	return __s2sdk_FindClientLoadoutSlot(playerSlot, team, itemDefinition);
}

extern bool (*__s2sdk_RestoreClientLoadout)(int32_t);

static bool RestoreClientLoadout(int32_t playerSlot) {
	return __s2sdk_RestoreClientLoadout(playerSlot);
}

extern bool (*__s2sdk_GiveClientDefaultWearables)(int32_t);

static bool GiveClientDefaultWearables(int32_t playerSlot) {
	return __s2sdk_GiveClientDefaultWearables(playerSlot);
}

extern Vector (*__s2sdk_GetClientInventoryItems)(int32_t);

static Vector GetClientInventoryItems(int32_t playerSlot) {
	return __s2sdk_GetClientInventoryItems(playerSlot);
}

extern uint64_t (*__s2sdk_CreateClientItem)(int32_t, uint16_t, int32_t);

static uint64_t CreateClientItem(int32_t playerSlot, uint16_t itemDefinition, int32_t quality) {
	return __s2sdk_CreateClientItem(playerSlot, itemDefinition, quality);
}

extern bool (*__s2sdk_RemoveClientItem)(int32_t, uint64_t);

static bool RemoveClientItem(int32_t playerSlot, uint64_t itemId) {
	return __s2sdk_RemoveClientItem(playerSlot, itemId);
}

extern bool (*__s2sdk_SetClientItemAttribute)(int32_t, uint64_t, String*, String*);

static bool SetClientItemAttribute(int32_t playerSlot, uint64_t itemId, String* attribute, String* value) {
	return __s2sdk_SetClientItemAttribute(playerSlot, itemId, attribute, value);
}

extern bool (*__s2sdk_SetClientItemPaint)(int32_t, uint64_t, int32_t, int32_t, float);

static bool SetClientItemPaint(int32_t playerSlot, uint64_t itemId, int32_t paintKit, int32_t seed, float wear) {
	return __s2sdk_SetClientItemPaint(playerSlot, itemId, paintKit, seed, wear);
}

extern bool (*__s2sdk_SetClientItemSticker)(int32_t, uint64_t, int32_t, int32_t, float);

static bool SetClientItemSticker(int32_t playerSlot, uint64_t itemId, int32_t stickerSlot, int32_t stickerId, float wear) {
	return __s2sdk_SetClientItemSticker(playerSlot, itemId, stickerSlot, stickerId, wear);
}

extern bool (*__s2sdk_SetClientItemKeychain)(int32_t, uint64_t, int32_t, int32_t, int32_t);

static bool SetClientItemKeychain(int32_t playerSlot, uint64_t itemId, int32_t keychainSlot, int32_t keychainId, int32_t seed) {
	return __s2sdk_SetClientItemKeychain(playerSlot, itemId, keychainSlot, keychainId, seed);
}

extern bool (*__s2sdk_SetClientItemStatTrak)(int32_t, uint64_t, int32_t);

static bool SetClientItemStatTrak(int32_t playerSlot, uint64_t itemId, int32_t kills) {
	return __s2sdk_SetClientItemStatTrak(playerSlot, itemId, kills);
}

extern bool (*__s2sdk_SetClientItemName)(int32_t, uint64_t, String*);

static bool SetClientItemName(int32_t playerSlot, uint64_t itemId, String* name) {
	return __s2sdk_SetClientItemName(playerSlot, itemId, name);
}

extern String (*__s2sdk_GetClientItemAttribute)(int32_t, uint64_t, String*);

static String GetClientItemAttribute(int32_t playerSlot, uint64_t itemId, String* attribute) {
	return __s2sdk_GetClientItemAttribute(playerSlot, itemId, attribute);
}

