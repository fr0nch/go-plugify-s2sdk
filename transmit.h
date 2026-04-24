#pragma once

#include "shared.h"

extern void (*__s2sdk_SetTransmitInfoEntity)(uintptr_t, int32_t);

static void SetTransmitInfoEntity(uintptr_t info, int32_t entityHandle) {
	__s2sdk_SetTransmitInfoEntity(info, entityHandle);
}

extern void (*__s2sdk_ClearTransmitInfoEntity)(uintptr_t, int32_t);

static void ClearTransmitInfoEntity(uintptr_t info, int32_t entityHandle) {
	__s2sdk_ClearTransmitInfoEntity(info, entityHandle);
}

extern bool (*__s2sdk_IsTransmitInfoEntitySet)(uintptr_t, int32_t);

static bool IsTransmitInfoEntitySet(uintptr_t info, int32_t entityHandle) {
	return __s2sdk_IsTransmitInfoEntitySet(info, entityHandle);
}

extern void (*__s2sdk_SetTransmitInfoEntityAll)(uintptr_t);

static void SetTransmitInfoEntityAll(uintptr_t info) {
	__s2sdk_SetTransmitInfoEntityAll(info);
}

extern void (*__s2sdk_ClearTransmitInfoEntityAll)(uintptr_t);

static void ClearTransmitInfoEntityAll(uintptr_t info) {
	__s2sdk_ClearTransmitInfoEntityAll(info);
}

extern void (*__s2sdk_SetTransmitInfoNonPlayer)(uintptr_t, int32_t);

static void SetTransmitInfoNonPlayer(uintptr_t info, int32_t entityHandle) {
	__s2sdk_SetTransmitInfoNonPlayer(info, entityHandle);
}

extern void (*__s2sdk_ClearTransmitInfoNonPlayer)(uintptr_t, int32_t);

static void ClearTransmitInfoNonPlayer(uintptr_t info, int32_t entityHandle) {
	__s2sdk_ClearTransmitInfoNonPlayer(info, entityHandle);
}

extern bool (*__s2sdk_IsTransmitInfoNonPlayerSet)(uintptr_t, int32_t);

static bool IsTransmitInfoNonPlayerSet(uintptr_t info, int32_t entityHandle) {
	return __s2sdk_IsTransmitInfoNonPlayerSet(info, entityHandle);
}

extern void (*__s2sdk_SetTransmitInfoNonPlayerAll)(uintptr_t);

static void SetTransmitInfoNonPlayerAll(uintptr_t info) {
	__s2sdk_SetTransmitInfoNonPlayerAll(info);
}

extern void (*__s2sdk_ClearTransmitInfoNonPlayerAll)(uintptr_t);

static void ClearTransmitInfoNonPlayerAll(uintptr_t info) {
	__s2sdk_ClearTransmitInfoNonPlayerAll(info);
}

extern void (*__s2sdk_SetTransmitInfoOutOfPVS)(uintptr_t, int32_t);

static void SetTransmitInfoOutOfPVS(uintptr_t info, int32_t entityHandle) {
	__s2sdk_SetTransmitInfoOutOfPVS(info, entityHandle);
}

extern void (*__s2sdk_ClearTransmitInfoOutOfPVS)(uintptr_t, int32_t);

static void ClearTransmitInfoOutOfPVS(uintptr_t info, int32_t entityHandle) {
	__s2sdk_ClearTransmitInfoOutOfPVS(info, entityHandle);
}

extern bool (*__s2sdk_IsTransmitInfoOutOfPVSSet)(uintptr_t, int32_t);

static bool IsTransmitInfoOutOfPVSSet(uintptr_t info, int32_t entityHandle) {
	return __s2sdk_IsTransmitInfoOutOfPVSSet(info, entityHandle);
}

extern void (*__s2sdk_SetTransmitInfoOutOfPVSAll)(uintptr_t);

static void SetTransmitInfoOutOfPVSAll(uintptr_t info) {
	__s2sdk_SetTransmitInfoOutOfPVSAll(info);
}

extern void (*__s2sdk_ClearTransmitInfoOutOfPVSAll)(uintptr_t);

static void ClearTransmitInfoOutOfPVSAll(uintptr_t info) {
	__s2sdk_ClearTransmitInfoOutOfPVSAll(info);
}

extern void (*__s2sdk_SetTransmitInfoHLTV)(uintptr_t, int32_t);

static void SetTransmitInfoHLTV(uintptr_t info, int32_t entityHandle) {
	__s2sdk_SetTransmitInfoHLTV(info, entityHandle);
}

extern void (*__s2sdk_ClearTransmitInfoHLTV)(uintptr_t, int32_t);

static void ClearTransmitInfoHLTV(uintptr_t info, int32_t entityHandle) {
	__s2sdk_ClearTransmitInfoHLTV(info, entityHandle);
}

extern bool (*__s2sdk_IsTransmitInfoHLTVSet)(uintptr_t, int32_t);

static bool IsTransmitInfoHLTVSet(uintptr_t info, int32_t entityHandle) {
	return __s2sdk_IsTransmitInfoHLTVSet(info, entityHandle);
}

extern void (*__s2sdk_SetTransmitInfoHLTVAll)(uintptr_t);

static void SetTransmitInfoHLTVAll(uintptr_t info) {
	__s2sdk_SetTransmitInfoHLTVAll(info);
}

extern void (*__s2sdk_ClearTransmitInfoHLTVAll)(uintptr_t);

static void ClearTransmitInfoHLTVAll(uintptr_t info) {
	__s2sdk_ClearTransmitInfoHLTVAll(info);
}

extern int32_t (*__s2sdk_GetTransmitInfoTargetSlotsCount)(uintptr_t);

static int32_t GetTransmitInfoTargetSlotsCount(uintptr_t info) {
	return __s2sdk_GetTransmitInfoTargetSlotsCount(info);
}

extern int32_t (*__s2sdk_GetTransmitInfoTargetSlot)(uintptr_t, int32_t);

static int32_t GetTransmitInfoTargetSlot(uintptr_t info, int32_t index) {
	return __s2sdk_GetTransmitInfoTargetSlot(info, index);
}

extern void (*__s2sdk_AddTransmitInfoTargetSlot)(uintptr_t, int32_t);

static void AddTransmitInfoTargetSlot(uintptr_t info, int32_t playerSlot) {
	__s2sdk_AddTransmitInfoTargetSlot(info, playerSlot);
}

extern void (*__s2sdk_RemoveTransmitInfoTargetSlot)(uintptr_t, int32_t);

static void RemoveTransmitInfoTargetSlot(uintptr_t info, int32_t index) {
	__s2sdk_RemoveTransmitInfoTargetSlot(info, index);
}

extern Vector (*__s2sdk_GetTransmitInfoTargetSlotsAll)(uintptr_t);

static Vector GetTransmitInfoTargetSlotsAll(uintptr_t info) {
	return __s2sdk_GetTransmitInfoTargetSlotsAll(info);
}

extern void (*__s2sdk_RemoveTransmitInfoTargetSlotsAll)(uintptr_t);

static void RemoveTransmitInfoTargetSlotsAll(uintptr_t info) {
	__s2sdk_RemoveTransmitInfoTargetSlotsAll(info);
}

extern int32_t (*__s2sdk_GetTransmitInfoPlayerSlot)(uintptr_t);

static int32_t GetTransmitInfoPlayerSlot(uintptr_t info) {
	return __s2sdk_GetTransmitInfoPlayerSlot(info);
}

extern void (*__s2sdk_SetTransmitInfoPlayerSlot)(uintptr_t, int32_t);

static void SetTransmitInfoPlayerSlot(uintptr_t info, int32_t playerSlot) {
	__s2sdk_SetTransmitInfoPlayerSlot(info, playerSlot);
}

extern bool (*__s2sdk_GetTransmitInfoFullUpdate)(uintptr_t);

static bool GetTransmitInfoFullUpdate(uintptr_t info) {
	return __s2sdk_GetTransmitInfoFullUpdate(info);
}

extern void (*__s2sdk_SetTransmitInfoFullUpdate)(uintptr_t, bool);

static void SetTransmitInfoFullUpdate(uintptr_t info, bool fullUpdate) {
	__s2sdk_SetTransmitInfoFullUpdate(info, fullUpdate);
}

