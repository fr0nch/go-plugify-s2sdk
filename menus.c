#include "shared.h"

PLUGIFY_EXPORT bool (*__s2sdk_RegisterMenuType)(String*, void*, void*, void*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_UnregisterMenuType)(String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsMenuTypeRegistered)(String*) = NULL;


PLUGIFY_EXPORT Vector (*__s2sdk_GetMenuTypes)() = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetDefaultMenuType)(String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetDefaultMenuType)() = NULL;


PLUGIFY_EXPORT uint32_t (*__s2sdk_CreateMenu)(String*, void*, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_DestroyMenu)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsValidMenu)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuTitle)(uint32_t, String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetMenuTitle)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuType)(uint32_t, String*) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetMenuType)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuPagination)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetMenuPagination)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuExitButton)(uint32_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_GetMenuExitButton)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuExitBackButton)(uint32_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_GetMenuExitBackButton)(uint32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuCloseOnSelect)(uint32_t, bool) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_GetMenuCloseOnSelect)(uint32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_AddMenuItem)(uint32_t, String*, String*, int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_InsertMenuItemAt)(uint32_t, int32_t, String*, String*, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveMenuItem)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_RemoveAllMenuItems)(uint32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetMenuItemsCount)(uint32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetMenuItemInfoText)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_GetMenuItemDisplay)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetMenuItemStyle)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_IsMenuItemSelectable)(uint32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuItemDisplay)(uint32_t, int32_t, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SetMenuItemStyle)(uint32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_DisplayMenu)(uint32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_DisplayMenuAtItem)(uint32_t, int32_t, int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_CancelClientMenu)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__s2sdk_GetClientMenu)(int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetClientMenuOffset)(int32_t) = NULL;


PLUGIFY_EXPORT double (*__s2sdk_GetClientMenuTime)(int32_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_GetClientMenuCursor)(int32_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_SetClientMenuCursor)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_ClientMenuHasPrevPage)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_ClientMenuHasNextPage)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_MenuNextPage)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_MenuPrevPage)(int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_SelectMenuItem)(int32_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_HandleDigitInput)(int32_t, int32_t) = NULL;


