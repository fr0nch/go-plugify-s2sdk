#pragma once

#include "shared.h"

extern bool (*__s2sdk_AddServerAddon)(uint64_t, bool);

static bool AddServerAddon(uint64_t addon, bool refresh) {
	return __s2sdk_AddServerAddon(addon, refresh);
}

extern bool (*__s2sdk_RemoveServerAddon)(uint64_t, bool);

static bool RemoveServerAddon(uint64_t addon, bool refresh) {
	return __s2sdk_RemoveServerAddon(addon, refresh);
}

extern Vector (*__s2sdk_GetServerAddons)();

static Vector GetServerAddons() {
	return __s2sdk_GetServerAddons();
}

extern bool (*__s2sdk_AddPrivateAddon)(uint64_t, bool);

static bool AddPrivateAddon(uint64_t addon, bool refresh) {
	return __s2sdk_AddPrivateAddon(addon, refresh);
}

extern bool (*__s2sdk_RemovePrivateAddon)(uint64_t, bool);

static bool RemovePrivateAddon(uint64_t addon, bool refresh) {
	return __s2sdk_RemovePrivateAddon(addon, refresh);
}

extern Vector (*__s2sdk_GetPrivateAddons)();

static Vector GetPrivateAddons() {
	return __s2sdk_GetPrivateAddons();
}

extern bool (*__s2sdk_AddClientOnlyAddon)(uint64_t, bool);

static bool AddClientOnlyAddon(uint64_t addon, bool refresh) {
	return __s2sdk_AddClientOnlyAddon(addon, refresh);
}

extern bool (*__s2sdk_RemoveClientOnlyAddon)(uint64_t);

static bool RemoveClientOnlyAddon(uint64_t addon) {
	return __s2sdk_RemoveClientOnlyAddon(addon);
}

extern Vector (*__s2sdk_GetClientOnlyAddons)();

static Vector GetClientOnlyAddons() {
	return __s2sdk_GetClientOnlyAddons();
}

extern bool (*__s2sdk_SendAddon)(uint64_t, uint64_t, bool);

static bool SendAddon(uint64_t steamId, uint64_t addon, bool refresh) {
	return __s2sdk_SendAddon(steamId, addon, refresh);
}

extern bool (*__s2sdk_RemoveSentAddon)(uint64_t, uint64_t);

static bool RemoveSentAddon(uint64_t steamId, uint64_t addon) {
	return __s2sdk_RemoveSentAddon(steamId, addon);
}

extern void (*__s2sdk_ClearSentAddons)(uint64_t);

static void ClearSentAddons(uint64_t steamId) {
	__s2sdk_ClearSentAddons(steamId);
}

extern Vector (*__s2sdk_GetSentAddons)(uint64_t);

static Vector GetSentAddons(uint64_t steamId) {
	return __s2sdk_GetSentAddons(steamId);
}

extern Vector (*__s2sdk_GetClientAddons)(uint64_t);

static Vector GetClientAddons(uint64_t steamId) {
	return __s2sdk_GetClientAddons(steamId);
}

extern uint64_t (*__s2sdk_GetCurrentWorkshopMap)();

static uint64_t GetCurrentWorkshopMap() {
	return __s2sdk_GetCurrentWorkshopMap();
}

