#pragma once

#include "shared.h"

extern bool (*__s2sdk_RegisterMenuType)(String*, void*, void*, void*);

static bool RegisterMenuType(String* name, void* display, void* close_, void* frame) {
	return __s2sdk_RegisterMenuType(name, display, close_, frame);
}

extern bool (*__s2sdk_UnregisterMenuType)(String*);

static bool UnregisterMenuType(String* name) {
	return __s2sdk_UnregisterMenuType(name);
}

extern bool (*__s2sdk_IsMenuTypeRegistered)(String*);

static bool IsMenuTypeRegistered(String* name) {
	return __s2sdk_IsMenuTypeRegistered(name);
}

extern Vector (*__s2sdk_GetMenuTypes)();

static Vector GetMenuTypes() {
	return __s2sdk_GetMenuTypes();
}

extern bool (*__s2sdk_SetDefaultMenuType)(String*);

static bool SetDefaultMenuType(String* name) {
	return __s2sdk_SetDefaultMenuType(name);
}

extern String (*__s2sdk_GetDefaultMenuType)();

static String GetDefaultMenuType() {
	return __s2sdk_GetDefaultMenuType();
}

extern uint32_t (*__s2sdk_CreateMenu)(String*, void*, String*);

static uint32_t CreateMenu(String* title, void* handler, String* menuType) {
	return __s2sdk_CreateMenu(title, handler, menuType);
}

extern bool (*__s2sdk_DestroyMenu)(uint32_t);

static bool DestroyMenu(uint32_t id) {
	return __s2sdk_DestroyMenu(id);
}

extern bool (*__s2sdk_IsValidMenu)(uint32_t);

static bool IsValidMenu(uint32_t id) {
	return __s2sdk_IsValidMenu(id);
}

extern bool (*__s2sdk_SetMenuTitle)(uint32_t, String*);

static bool SetMenuTitle(uint32_t id, String* title) {
	return __s2sdk_SetMenuTitle(id, title);
}

extern String (*__s2sdk_GetMenuTitle)(uint32_t);

static String GetMenuTitle(uint32_t id) {
	return __s2sdk_GetMenuTitle(id);
}

extern bool (*__s2sdk_SetMenuType)(uint32_t, String*);

static bool SetMenuType(uint32_t id, String* typeName) {
	return __s2sdk_SetMenuType(id, typeName);
}

extern String (*__s2sdk_GetMenuType)(uint32_t);

static String GetMenuType(uint32_t id) {
	return __s2sdk_GetMenuType(id);
}

extern bool (*__s2sdk_SetMenuPagination)(uint32_t, int32_t);

static bool SetMenuPagination(uint32_t id, int32_t itemsPerPage) {
	return __s2sdk_SetMenuPagination(id, itemsPerPage);
}

extern int32_t (*__s2sdk_GetMenuPagination)(uint32_t);

static int32_t GetMenuPagination(uint32_t id) {
	return __s2sdk_GetMenuPagination(id);
}

extern bool (*__s2sdk_SetMenuExitButton)(uint32_t, bool);

static bool SetMenuExitButton(uint32_t id, bool enabled) {
	return __s2sdk_SetMenuExitButton(id, enabled);
}

extern bool (*__s2sdk_GetMenuExitButton)(uint32_t);

static bool GetMenuExitButton(uint32_t id) {
	return __s2sdk_GetMenuExitButton(id);
}

extern bool (*__s2sdk_SetMenuExitBackButton)(uint32_t, bool);

static bool SetMenuExitBackButton(uint32_t id, bool enabled) {
	return __s2sdk_SetMenuExitBackButton(id, enabled);
}

extern bool (*__s2sdk_GetMenuExitBackButton)(uint32_t);

static bool GetMenuExitBackButton(uint32_t id) {
	return __s2sdk_GetMenuExitBackButton(id);
}

extern bool (*__s2sdk_SetMenuCloseOnSelect)(uint32_t, bool);

static bool SetMenuCloseOnSelect(uint32_t id, bool enabled) {
	return __s2sdk_SetMenuCloseOnSelect(id, enabled);
}

extern bool (*__s2sdk_GetMenuCloseOnSelect)(uint32_t);

static bool GetMenuCloseOnSelect(uint32_t id) {
	return __s2sdk_GetMenuCloseOnSelect(id);
}

extern int32_t (*__s2sdk_AddMenuItem)(uint32_t, String*, String*, int32_t);

static int32_t AddMenuItem(uint32_t id, String* info, String* display, int32_t style) {
	return __s2sdk_AddMenuItem(id, info, display, style);
}

extern int32_t (*__s2sdk_InsertMenuItemAt)(uint32_t, int32_t, String*, String*, int32_t);

static int32_t InsertMenuItemAt(uint32_t id, int32_t index, String* info, String* display, int32_t style) {
	return __s2sdk_InsertMenuItemAt(id, index, info, display, style);
}

extern bool (*__s2sdk_RemoveMenuItem)(uint32_t, int32_t);

static bool RemoveMenuItem(uint32_t id, int32_t index) {
	return __s2sdk_RemoveMenuItem(id, index);
}

extern bool (*__s2sdk_RemoveAllMenuItems)(uint32_t);

static bool RemoveAllMenuItems(uint32_t id) {
	return __s2sdk_RemoveAllMenuItems(id);
}

extern int32_t (*__s2sdk_GetMenuItemsCount)(uint32_t);

static int32_t GetMenuItemsCount(uint32_t id) {
	return __s2sdk_GetMenuItemsCount(id);
}

extern String (*__s2sdk_GetMenuItemInfoText)(uint32_t, int32_t);

static String GetMenuItemInfoText(uint32_t id, int32_t index) {
	return __s2sdk_GetMenuItemInfoText(id, index);
}

extern String (*__s2sdk_GetMenuItemDisplay)(uint32_t, int32_t);

static String GetMenuItemDisplay(uint32_t id, int32_t index) {
	return __s2sdk_GetMenuItemDisplay(id, index);
}

extern int32_t (*__s2sdk_GetMenuItemStyle)(uint32_t, int32_t);

static int32_t GetMenuItemStyle(uint32_t id, int32_t index) {
	return __s2sdk_GetMenuItemStyle(id, index);
}

extern bool (*__s2sdk_IsMenuItemSelectable)(uint32_t, int32_t);

static bool IsMenuItemSelectable(uint32_t id, int32_t index) {
	return __s2sdk_IsMenuItemSelectable(id, index);
}

extern bool (*__s2sdk_SetMenuItemDisplay)(uint32_t, int32_t, String*);

static bool SetMenuItemDisplay(uint32_t id, int32_t index, String* display) {
	return __s2sdk_SetMenuItemDisplay(id, index, display);
}

extern bool (*__s2sdk_SetMenuItemStyle)(uint32_t, int32_t, int32_t);

static bool SetMenuItemStyle(uint32_t id, int32_t index, int32_t style) {
	return __s2sdk_SetMenuItemStyle(id, index, style);
}

extern bool (*__s2sdk_DisplayMenu)(uint32_t, int32_t, int32_t);

static bool DisplayMenu(uint32_t id, int32_t playerSlot, int32_t time) {
	return __s2sdk_DisplayMenu(id, playerSlot, time);
}

extern bool (*__s2sdk_DisplayMenuAtItem)(uint32_t, int32_t, int32_t, int32_t);

static bool DisplayMenuAtItem(uint32_t id, int32_t playerSlot, int32_t firstItem, int32_t time) {
	return __s2sdk_DisplayMenuAtItem(id, playerSlot, firstItem, time);
}

extern bool (*__s2sdk_CancelClientMenu)(int32_t, int32_t);

static bool CancelClientMenu(int32_t playerSlot, int32_t reason) {
	return __s2sdk_CancelClientMenu(playerSlot, reason);
}

extern uint32_t (*__s2sdk_GetClientMenu)(int32_t);

static uint32_t GetClientMenu(int32_t playerSlot) {
	return __s2sdk_GetClientMenu(playerSlot);
}

extern int32_t (*__s2sdk_GetClientMenuOffset)(int32_t);

static int32_t GetClientMenuOffset(int32_t playerSlot) {
	return __s2sdk_GetClientMenuOffset(playerSlot);
}

extern double (*__s2sdk_GetClientMenuTime)(int32_t);

static double GetClientMenuTime(int32_t playerSlot) {
	return __s2sdk_GetClientMenuTime(playerSlot);
}

extern int32_t (*__s2sdk_GetClientMenuCursor)(int32_t);

static int32_t GetClientMenuCursor(int32_t playerSlot) {
	return __s2sdk_GetClientMenuCursor(playerSlot);
}

extern void (*__s2sdk_SetClientMenuCursor)(int32_t, int32_t);

static void SetClientMenuCursor(int32_t playerSlot, int32_t index) {
	__s2sdk_SetClientMenuCursor(playerSlot, index);
}

extern bool (*__s2sdk_ClientMenuHasPrevPage)(int32_t);

static bool ClientMenuHasPrevPage(int32_t playerSlot) {
	return __s2sdk_ClientMenuHasPrevPage(playerSlot);
}

extern bool (*__s2sdk_ClientMenuHasNextPage)(int32_t);

static bool ClientMenuHasNextPage(int32_t playerSlot) {
	return __s2sdk_ClientMenuHasNextPage(playerSlot);
}

extern bool (*__s2sdk_MenuNextPage)(int32_t);

static bool MenuNextPage(int32_t playerSlot) {
	return __s2sdk_MenuNextPage(playerSlot);
}

extern bool (*__s2sdk_MenuPrevPage)(int32_t);

static bool MenuPrevPage(int32_t playerSlot) {
	return __s2sdk_MenuPrevPage(playerSlot);
}

extern bool (*__s2sdk_SelectMenuItem)(int32_t, int32_t);

static bool SelectMenuItem(int32_t playerSlot, int32_t itemIndex) {
	return __s2sdk_SelectMenuItem(playerSlot, itemIndex);
}

extern bool (*__s2sdk_HandleDigitInput)(int32_t, int32_t);

static bool HandleDigitInput(int32_t playerSlot, int32_t digit) {
	return __s2sdk_HandleDigitInput(playerSlot, digit);
}

