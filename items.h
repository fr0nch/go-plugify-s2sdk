#pragma once

#include "shared.h"

extern String (*__s2sdk_GetItemDefinitionName)(uint16_t);

static String GetItemDefinitionName(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionName(itemDefinition);
}

extern uint16_t (*__s2sdk_GetItemDefinitionByName)(String*);

static uint16_t GetItemDefinitionByName(String* name) {
	return __s2sdk_GetItemDefinitionByName(name);
}

extern String (*__s2sdk_GetItemDefinitionClassName)(uint16_t);

static String GetItemDefinitionClassName(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionClassName(itemDefinition);
}

extern int32_t (*__s2sdk_GetItemDefinitionLoadoutSlot)(uint16_t, int32_t);

static int32_t GetItemDefinitionLoadoutSlot(uint16_t itemDefinition, int32_t team) {
	return __s2sdk_GetItemDefinitionLoadoutSlot(itemDefinition, team);
}

extern int32_t (*__s2sdk_GetItemDefinitionTeam)(uint16_t);

static int32_t GetItemDefinitionTeam(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionTeam(itemDefinition);
}

extern int32_t (*__s2sdk_GetItemDefinitionStickerSlots)(uint16_t);

static int32_t GetItemDefinitionStickerSlots(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionStickerSlots(itemDefinition);
}

extern int32_t (*__s2sdk_GetItemDefinitionKeychainSlots)(uint16_t);

static int32_t GetItemDefinitionKeychainSlots(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionKeychainSlots(itemDefinition);
}

extern String (*__s2sdk_GetItemDefinitionImage)(uint16_t);

static String GetItemDefinitionImage(uint16_t itemDefinition) {
	return __s2sdk_GetItemDefinitionImage(itemDefinition);
}

extern String (*__s2sdk_GetItemDefinitionModel)(uint16_t, bool);

static String GetItemDefinitionModel(uint16_t itemDefinition, bool world) {
	return __s2sdk_GetItemDefinitionModel(itemDefinition, world);
}

extern int32_t (*__s2sdk_GetAttributeDefinitionIndex)(String*);

static int32_t GetAttributeDefinitionIndex(String* attribute) {
	return __s2sdk_GetAttributeDefinitionIndex(attribute);
}

extern String (*__s2sdk_GetAttributeDefinitionName)(int32_t);

static String GetAttributeDefinitionName(int32_t attributeIndex) {
	return __s2sdk_GetAttributeDefinitionName(attributeIndex);
}

extern String (*__s2sdk_GetPaintKitName)(int32_t);

static String GetPaintKitName(int32_t paintKit) {
	return __s2sdk_GetPaintKitName(paintKit);
}

extern String (*__s2sdk_GetPaintKitLocName)(int32_t);

static String GetPaintKitLocName(int32_t paintKit) {
	return __s2sdk_GetPaintKitLocName(paintKit);
}

extern int32_t (*__s2sdk_FindPaintKit)(String*);

static int32_t FindPaintKit(String* name) {
	return __s2sdk_FindPaintKit(name);
}

extern String (*__s2sdk_GetStickerKitName)(int32_t);

static String GetStickerKitName(int32_t stickerKit) {
	return __s2sdk_GetStickerKitName(stickerKit);
}

extern String (*__s2sdk_GetStickerKitLocName)(int32_t);

static String GetStickerKitLocName(int32_t stickerKit) {
	return __s2sdk_GetStickerKitLocName(stickerKit);
}

extern int32_t (*__s2sdk_FindStickerKit)(String*);

static int32_t FindStickerKit(String* name) {
	return __s2sdk_FindStickerKit(name);
}

extern String (*__s2sdk_GetKeychainName)(int32_t);

static String GetKeychainName(int32_t keychain) {
	return __s2sdk_GetKeychainName(keychain);
}

extern String (*__s2sdk_GetKeychainLocName)(int32_t);

static String GetKeychainLocName(int32_t keychain) {
	return __s2sdk_GetKeychainLocName(keychain);
}

extern int32_t (*__s2sdk_FindKeychain)(String*);

static int32_t FindKeychain(String* name) {
	return __s2sdk_FindKeychain(name);
}

extern uintptr_t (*__s2sdk_GetClientItemView)(int32_t, uint64_t);

static uintptr_t GetClientItemView(int32_t playerSlot, uint64_t itemId) {
	return __s2sdk_GetClientItemView(playerSlot, itemId);
}

extern uintptr_t (*__s2sdk_GetClientLoadoutItemView)(int32_t, int32_t, int32_t);

static uintptr_t GetClientLoadoutItemView(int32_t playerSlot, int32_t team, int32_t loadoutSlot) {
	return __s2sdk_GetClientLoadoutItemView(playerSlot, team, loadoutSlot);
}

extern uintptr_t (*__s2sdk_GetEntityItemView)(int32_t);

static uintptr_t GetEntityItemView(int32_t entityHandle) {
	return __s2sdk_GetEntityItemView(entityHandle);
}

extern uint16_t (*__s2sdk_GetItemViewDefinition)(uintptr_t);

static uint16_t GetItemViewDefinition(uintptr_t itemView) {
	return __s2sdk_GetItemViewDefinition(itemView);
}

extern uint64_t (*__s2sdk_GetItemViewId)(uintptr_t);

static uint64_t GetItemViewId(uintptr_t itemView) {
	return __s2sdk_GetItemViewId(itemView);
}

extern uint32_t (*__s2sdk_GetItemViewAccountId)(uintptr_t);

static uint32_t GetItemViewAccountId(uintptr_t itemView) {
	return __s2sdk_GetItemViewAccountId(itemView);
}

extern int32_t (*__s2sdk_GetItemViewQuality)(uintptr_t);

static int32_t GetItemViewQuality(uintptr_t itemView) {
	return __s2sdk_GetItemViewQuality(itemView);
}

extern int32_t (*__s2sdk_GetItemViewRarity)(uintptr_t);

static int32_t GetItemViewRarity(uintptr_t itemView) {
	return __s2sdk_GetItemViewRarity(itemView);
}

extern int32_t (*__s2sdk_GetItemViewOrigin)(uintptr_t);

static int32_t GetItemViewOrigin(uintptr_t itemView) {
	return __s2sdk_GetItemViewOrigin(itemView);
}

extern bool (*__s2sdk_IsItemViewTradable)(uintptr_t);

static bool IsItemViewTradable(uintptr_t itemView) {
	return __s2sdk_IsItemViewTradable(itemView);
}

extern bool (*__s2sdk_IsItemViewMarketable)(uintptr_t);

static bool IsItemViewMarketable(uintptr_t itemView) {
	return __s2sdk_IsItemViewMarketable(itemView);
}

extern int32_t (*__s2sdk_GetItemViewPaintKit)(uintptr_t);

static int32_t GetItemViewPaintKit(uintptr_t itemView) {
	return __s2sdk_GetItemViewPaintKit(itemView);
}

extern int32_t (*__s2sdk_GetItemViewPaintKitSeed)(uintptr_t);

static int32_t GetItemViewPaintKitSeed(uintptr_t itemView) {
	return __s2sdk_GetItemViewPaintKitSeed(itemView);
}

extern float (*__s2sdk_GetItemViewPaintKitWear)(uintptr_t);

static float GetItemViewPaintKitWear(uintptr_t itemView) {
	return __s2sdk_GetItemViewPaintKitWear(itemView);
}

extern double (*__s2sdk_GetItemViewStickerAttribute)(uintptr_t, int32_t, int32_t);

static double GetItemViewStickerAttribute(uintptr_t itemView, int32_t stickerSlot, int32_t attribute) {
	return __s2sdk_GetItemViewStickerAttribute(itemView, stickerSlot, attribute);
}

extern double (*__s2sdk_GetItemViewKeychainAttribute)(uintptr_t, int32_t, int32_t);

static double GetItemViewKeychainAttribute(uintptr_t itemView, int32_t keychainSlot, int32_t attribute) {
	return __s2sdk_GetItemViewKeychainAttribute(itemView, keychainSlot, attribute);
}

extern String (*__s2sdk_GetItemViewAttribute)(uintptr_t, String*);

static String GetItemViewAttribute(uintptr_t itemView, String* attribute) {
	return __s2sdk_GetItemViewAttribute(itemView, attribute);
}

extern int32_t (*__s2sdk_GiveClientItem)(int32_t, uint64_t);

static int32_t GiveClientItem(int32_t playerSlot, uint64_t itemId) {
	return __s2sdk_GiveClientItem(playerSlot, itemId);
}

