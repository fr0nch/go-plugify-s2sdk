package s2sdk

/*
#include "menus.h"
#cgo noescape RegisterMenuType
#cgo noescape UnregisterMenuType
#cgo noescape IsMenuTypeRegistered
#cgo noescape GetMenuTypes
#cgo noescape SetDefaultMenuType
#cgo noescape GetDefaultMenuType
#cgo noescape CreateMenu
#cgo noescape DestroyMenu
#cgo noescape IsValidMenu
#cgo noescape SetMenuTitle
#cgo noescape GetMenuTitle
#cgo noescape SetMenuType
#cgo noescape GetMenuType
#cgo noescape SetMenuPagination
#cgo noescape GetMenuPagination
#cgo noescape SetMenuExitButton
#cgo noescape GetMenuExitButton
#cgo noescape SetMenuExitBackButton
#cgo noescape GetMenuExitBackButton
#cgo noescape SetMenuCloseOnSelect
#cgo noescape GetMenuCloseOnSelect
#cgo noescape AddMenuItem
#cgo noescape InsertMenuItemAt
#cgo noescape RemoveMenuItem
#cgo noescape RemoveAllMenuItems
#cgo noescape GetMenuItemsCount
#cgo noescape GetMenuItemInfoText
#cgo noescape GetMenuItemDisplay
#cgo noescape GetMenuItemStyle
#cgo noescape IsMenuItemSelectable
#cgo noescape SetMenuItemDisplay
#cgo noescape SetMenuItemStyle
#cgo noescape DisplayMenu
#cgo noescape DisplayMenuAtItem
#cgo noescape CancelClientMenu
#cgo noescape GetClientMenu
#cgo noescape GetClientMenuOffset
#cgo noescape GetClientMenuTime
#cgo noescape GetClientMenuCursor
#cgo noescape SetClientMenuCursor
#cgo noescape ClientMenuHasPrevPage
#cgo noescape ClientMenuHasNextPage
#cgo noescape MenuNextPage
#cgo noescape MenuPrevPage
#cgo noescape SelectMenuItem
#cgo noescape HandleDigitInput
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from s2sdk (group: menus)

var _RegisterMenuType = func(name string, display MenuDisplayCallback, close_ MenuCloseCallback, frame MenuFrameCallback) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__display := plugify.GetFunctionPointerForDelegate(display)
	__close_ := plugify.GetFunctionPointerForDelegate(close_)
	__frame := plugify.GetFunctionPointerForDelegate(frame)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.RegisterMenuType((*C.String)(unsafe.Pointer(&__name)), __display, __close_, __frame))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// RegisterMenuType 
//  @brief Registers a new menu type backend (e.g. a custom rendering style), so it can be selected via CreateMenu/SetMenuType. Built-in types "chat", "console", "centerhtml" and "button" are registered automatically at startup.
//
//  @param name: The unique, case-insensitive name of the menu type.
//  @param display: Callback invoked to render the menu's current state to a client.
//  @param close_: Callback invoked to hide/clean up whatever UI was shown to a client.
//  @param frame: Optional callback invoked every server frame while a client has this menu type open (e.g. for input polling). Pass null if not needed.
//
//  @return True if the type was registered; false if the name is empty, a callback is null, or the name is already taken.
func RegisterMenuType(name string, display MenuDisplayCallback, close_ MenuCloseCallback, frame MenuFrameCallback) bool {
	return _RegisterMenuType(name, display, close_, frame)
}

var _UnregisterMenuType = func(name string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.UnregisterMenuType((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// UnregisterMenuType 
//  @brief Unregisters a previously registered menu type.
//
//  @param name: The name of the menu type to remove.
//
//  @return True if a type with that name was found and removed.
func UnregisterMenuType(name string) bool {
	return _UnregisterMenuType(name)
}

var _IsMenuTypeRegistered = func(name string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.IsMenuTypeRegistered((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// IsMenuTypeRegistered 
//  @brief Checks whether a menu type with the given name is currently registered.
//
//  @param name: The name of the menu type.
//
//  @return True if the menu type is registered.
func IsMenuTypeRegistered(name string) bool {
	return _IsMenuTypeRegistered(name)
}

var _GetMenuTypes = func() []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	plugify.Block {
		Try: func() {
			__native := C.GetMenuTypes()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetMenuTypes 
//  @brief Returns the names of all currently registered menu types.
//
//
//  @return The vector of menu type names.
func GetMenuTypes() []string {
	return _GetMenuTypes()
}

var _SetDefaultMenuType = func(name string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetDefaultMenuType((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// SetDefaultMenuType 
//  @brief Sets the menu type used by menus that don't specify one explicitly.
//
//  @param name: The name of an already-registered menu type.
//
//  @return True if the type exists and was set as default.
func SetDefaultMenuType(name string) bool {
	return _SetDefaultMenuType(name)
}

var _GetDefaultMenuType = func() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block {
		Try: func() {
			__native := C.GetDefaultMenuType()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetDefaultMenuType 
//  @brief Returns the current default menu type name.
//
//
//  @return The default menu type name.
func GetDefaultMenuType() string {
	return _GetDefaultMenuType()
}

var _CreateMenu = func(title string, handler MenuHandlerCallback, menuType string) MenuId {
	var __retVal MenuId
	__title := plugify.ConstructString(title)
	__handler := plugify.GetFunctionPointerForDelegate(handler)
	__menuType := plugify.ConstructString(menuType)
	plugify.Block {
		Try: func() {
			__retVal = MenuId(C.CreateMenu((*C.String)(unsafe.Pointer(&__title)), __handler, (*C.String)(unsafe.Pointer(&__menuType))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__title)
			plugify.DestroyString(&__menuType)
		},
	}.Do()
	return __retVal
}

// CreateMenu 
//  @brief Creates a new menu.
//
//  @param title: The title shown at the top of the menu.
//  @param handler: Callback invoked with Start/Select/Cancel/End actions as the menu is displayed and interacted with.
//  @param menuType: The name of the menu type backend to render with. Empty uses the current default menu type.
//
//  @return A handle to the created menu.
func CreateMenu(title string, handler MenuHandlerCallback, menuType string) MenuId {
	return _CreateMenu(title, handler, menuType)
}

var _DestroyMenu = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.DestroyMenu(__id))
	return __retVal
}

// DestroyMenu 
//  @brief Destroys a menu. Any client currently viewing it is cancelled first (MenuCancelReason::Destroyed).
//
//  @param id: The handle to the menu.
//
//  @return True if the menu existed and was destroyed.
func DestroyMenu(id MenuId) bool {
	return _DestroyMenu(id)
}

var _IsValidMenu = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.IsValidMenu(__id))
	return __retVal
}

// IsValidMenu 
//  @brief Checks whether a menu handle refers to an existing menu.
//
//  @param id: The handle to the menu.
//
//  @return True if the handle is valid.
func IsValidMenu(id MenuId) bool {
	return _IsValidMenu(id)
}

var _SetMenuTitle = func(id MenuId, title string) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__title := plugify.ConstructString(title)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetMenuTitle(__id, (*C.String)(unsafe.Pointer(&__title))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__title)
		},
	}.Do()
	return __retVal
}

// SetMenuTitle 
//  @brief Sets a menu's title.
//
//  @param id: The handle to the menu.
//  @param title: The new title.
//
//  @return True if the menu exists.
func SetMenuTitle(id MenuId, title string) bool {
	return _SetMenuTitle(id, title)
}

var _GetMenuTitle = func(id MenuId) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__id := C.uint32_t(id)
	plugify.Block {
		Try: func() {
			__native := C.GetMenuTitle(__id)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetMenuTitle 
//  @brief Gets a menu's title.
//
//  @param id: The handle to the menu.
//
//  @return The menu's title, or an empty string if the handle is invalid.
func GetMenuTitle(id MenuId) string {
	return _GetMenuTitle(id)
}

var _SetMenuType = func(id MenuId, typeName string) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__typeName := plugify.ConstructString(typeName)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetMenuType(__id, (*C.String)(unsafe.Pointer(&__typeName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__typeName)
		},
	}.Do()
	return __retVal
}

// SetMenuType 
//  @brief Sets which registered menu type backend renders this menu.
//
//  @param id: The handle to the menu.
//  @param typeName: The name of a registered menu type, or empty to use the default menu type.
//
//  @return True if the menu exists.
func SetMenuType(id MenuId, typeName string) bool {
	return _SetMenuType(id, typeName)
}

var _GetMenuType = func(id MenuId) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__id := C.uint32_t(id)
	plugify.Block {
		Try: func() {
			__native := C.GetMenuType(__id)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetMenuType 
//  @brief Gets the menu type backend name assigned to this menu.
//
//  @param id: The handle to the menu.
//
//  @return The menu type name (may be empty, meaning "use the default").
func GetMenuType(id MenuId) string {
	return _GetMenuType(id)
}

var _SetMenuPagination = func(id MenuId, itemsPerPage int32) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__itemsPerPage := C.int32_t(itemsPerPage)
	__retVal = bool(C.SetMenuPagination(__id, __itemsPerPage))
	return __retVal
}

// SetMenuPagination 
//  @brief Sets how many items are shown per page.
//
//  @param id: The handle to the menu.
//  @param itemsPerPage: The number of items per page, or 0 to disable pagination (show every item on one page).
//
//  @return True if the menu exists and itemsPerPage is not negative.
func SetMenuPagination(id MenuId, itemsPerPage int32) bool {
	return _SetMenuPagination(id, itemsPerPage)
}

var _GetMenuPagination = func(id MenuId) int32 {
	var __retVal int32
	__id := C.uint32_t(id)
	__retVal = int32(C.GetMenuPagination(__id))
	return __retVal
}

// GetMenuPagination 
//  @brief Gets how many items are shown per page.
//
//  @param id: The handle to the menu.
//
//  @return The items-per-page value, 0 meaning pagination is disabled.
func GetMenuPagination(id MenuId) int32 {
	return _GetMenuPagination(id)
}

var _SetMenuExitButton = func(id MenuId, enabled bool) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__enabled := C.bool(enabled)
	__retVal = bool(C.SetMenuExitButton(__id, __enabled))
	return __retVal
}

// SetMenuExitButton 
//  @brief Sets whether the menu shows an exit option.
//
//  @param id: The handle to the menu.
//  @param enabled: True to show an exit option.
//
//  @return True if the menu exists.
func SetMenuExitButton(id MenuId, enabled bool) bool {
	return _SetMenuExitButton(id, enabled)
}

var _GetMenuExitButton = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.GetMenuExitButton(__id))
	return __retVal
}

// GetMenuExitButton 
//  @brief Gets whether the menu shows an exit option.
//
//  @param id: The handle to the menu.
//
//  @return True if the exit option is enabled.
func GetMenuExitButton(id MenuId) bool {
	return _GetMenuExitButton(id)
}

var _SetMenuExitBackButton = func(id MenuId, enabled bool) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__enabled := C.bool(enabled)
	__retVal = bool(C.SetMenuExitBackButton(__id, __enabled))
	return __retVal
}

// SetMenuExitBackButton 
//  @brief Sets whether the menu shows a "back" option in place of the exit option. Selecting it cancels the display with MenuCancelReason::ExitBack instead of MenuCancelReason::Exit, which a handler can use to redisplay a parent menu (SourceMod-style ExitBack).
//
//  @param id: The handle to the menu.
//  @param enabled: True to show a back option instead of the exit option.
//
//  @return True if the menu exists.
func SetMenuExitBackButton(id MenuId, enabled bool) bool {
	return _SetMenuExitBackButton(id, enabled)
}

var _GetMenuExitBackButton = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.GetMenuExitBackButton(__id))
	return __retVal
}

// GetMenuExitBackButton 
//  @brief Gets whether the menu shows a "back" option in place of the exit option.
//
//  @param id: The handle to the menu.
//
//  @return True if the back option is enabled.
func GetMenuExitBackButton(id MenuId) bool {
	return _GetMenuExitBackButton(id)
}

var _SetMenuCloseOnSelect = func(id MenuId, enabled bool) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__enabled := C.bool(enabled)
	__retVal = bool(C.SetMenuCloseOnSelect(__id, __enabled))
	return __retVal
}

// SetMenuCloseOnSelect 
//  @brief Sets whether selecting an item automatically closes the menu display for that client. When disabled, the display stays open after MenuAction::Select and the handler is responsible for closing/redisplaying it if desired.
//
//  @param id: The handle to the menu.
//  @param enabled: True to auto-close on selection (the default).
//
//  @return True if the menu exists.
func SetMenuCloseOnSelect(id MenuId, enabled bool) bool {
	return _SetMenuCloseOnSelect(id, enabled)
}

var _GetMenuCloseOnSelect = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.GetMenuCloseOnSelect(__id))
	return __retVal
}

// GetMenuCloseOnSelect 
//  @brief Gets whether selecting an item automatically closes the menu display for that client.
//
//  @param id: The handle to the menu.
//
//  @return True if close-on-select is enabled.
func GetMenuCloseOnSelect(id MenuId) bool {
	return _GetMenuCloseOnSelect(id)
}

var _AddMenuItem = func(id MenuId, info string, display string, style MenuItemStyle) int32 {
	var __retVal int32
	__id := C.uint32_t(id)
	__info := plugify.ConstructString(info)
	__display := plugify.ConstructString(display)
	__style := C.int32_t(style)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.AddMenuItem(__id, (*C.String)(unsafe.Pointer(&__info)), (*C.String)(unsafe.Pointer(&__display)), __style))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__info)
			plugify.DestroyString(&__display)
		},
	}.Do()
	return __retVal
}

// AddMenuItem 
//  @brief Appends an item to the end of a menu.
//
//  @param id: The handle to the menu.
//  @param info: An internal identifier for the item, not shown to the client; retrieve it with GetMenuItemInfo from within the handler callback.
//  @param display: The text shown to the client.
//  @param style: The item's draw style (Default/Disabled/Spacer).
//
//  @return The index of the newly added item, or -1 if the menu handle is invalid.
func AddMenuItem(id MenuId, info string, display string, style MenuItemStyle) int32 {
	return _AddMenuItem(id, info, display, style)
}

var _InsertMenuItemAt = func(id MenuId, index int32, info string, display string, style MenuItemStyle) int32 {
	var __retVal int32
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__info := plugify.ConstructString(info)
	__display := plugify.ConstructString(display)
	__style := C.int32_t(style)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.InsertMenuItemAt(__id, __index, (*C.String)(unsafe.Pointer(&__info)), (*C.String)(unsafe.Pointer(&__display)), __style))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__info)
			plugify.DestroyString(&__display)
		},
	}.Do()
	return __retVal
}

// InsertMenuItemAt 
//  @brief Inserts an item into a menu at a specific index.
//
//  @param id: The handle to the menu.
//  @param index: The index to insert at; must be within [0, item count].
//  @param info: An internal identifier for the item, not shown to the client.
//  @param display: The text shown to the client.
//  @param style: The item's draw style (Default/Disabled/Spacer).
//
//  @return The index the item was inserted at, or -1 on failure.
func InsertMenuItemAt(id MenuId, index int32, info string, display string, style MenuItemStyle) int32 {
	return _InsertMenuItemAt(id, index, info, display, style)
}

var _RemoveMenuItem = func(id MenuId, index int32) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__retVal = bool(C.RemoveMenuItem(__id, __index))
	return __retVal
}

// RemoveMenuItem 
//  @brief Removes an item from a menu.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item to remove.
//
//  @return True if the item existed and was removed.
func RemoveMenuItem(id MenuId, index int32) bool {
	return _RemoveMenuItem(id, index)
}

var _RemoveAllMenuItems = func(id MenuId) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__retVal = bool(C.RemoveAllMenuItems(__id))
	return __retVal
}

// RemoveAllMenuItems 
//  @brief Removes every item from a menu.
//
//  @param id: The handle to the menu.
//
//  @return True if the menu exists.
func RemoveAllMenuItems(id MenuId) bool {
	return _RemoveAllMenuItems(id)
}

var _GetMenuItemsCount = func(id MenuId) int32 {
	var __retVal int32
	__id := C.uint32_t(id)
	__retVal = int32(C.GetMenuItemsCount(__id))
	return __retVal
}

// GetMenuItemsCount 
//  @brief Gets the number of items in a menu.
//
//  @param id: The handle to the menu.
//
//  @return The item count, or 0 if the handle is invalid.
func GetMenuItemsCount(id MenuId) int32 {
	return _GetMenuItemsCount(id)
}

var _GetMenuItemInfoText = func(id MenuId, index int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	plugify.Block {
		Try: func() {
			__native := C.GetMenuItemInfoText(__id, __index)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetMenuItemInfoText 
//  @brief Gets an item's internal info string.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's info string, or empty if out of range.
func GetMenuItemInfoText(id MenuId, index int32) string {
	return _GetMenuItemInfoText(id, index)
}

var _GetMenuItemDisplay = func(id MenuId, index int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	plugify.Block {
		Try: func() {
			__native := C.GetMenuItemDisplay(__id, __index)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetMenuItemDisplay 
//  @brief Gets an item's display text.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's display text, or empty if out of range.
func GetMenuItemDisplay(id MenuId, index int32) string {
	return _GetMenuItemDisplay(id, index)
}

var _GetMenuItemStyle = func(id MenuId, index int32) MenuItemStyle {
	var __retVal MenuItemStyle
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__retVal = MenuItemStyle(C.GetMenuItemStyle(__id, __index))
	return __retVal
}

// GetMenuItemStyle 
//  @brief Gets an item's draw style.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's style; MenuItemStyle::Disabled if out of range.
func GetMenuItemStyle(id MenuId, index int32) MenuItemStyle {
	return _GetMenuItemStyle(id, index)
}

var _IsMenuItemSelectable = func(id MenuId, index int32) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__retVal = bool(C.IsMenuItemSelectable(__id, __index))
	return __retVal
}

// IsMenuItemSelectable 
//  @brief Checks whether an item can currently be selected (style is Default).
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return True if the item is selectable.
func IsMenuItemSelectable(id MenuId, index int32) bool {
	return _IsMenuItemSelectable(id, index)
}

var _SetMenuItemDisplay = func(id MenuId, index int32, display string) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__display := plugify.ConstructString(display)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetMenuItemDisplay(__id, __index, (*C.String)(unsafe.Pointer(&__display))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__display)
		},
	}.Do()
	return __retVal
}

// SetMenuItemDisplay 
//  @brief Changes an item's display text.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//  @param display: The new display text.
//
//  @return True if the item exists.
func SetMenuItemDisplay(id MenuId, index int32, display string) bool {
	return _SetMenuItemDisplay(id, index, display)
}

var _SetMenuItemStyle = func(id MenuId, index int32, style MenuItemStyle) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__index := C.int32_t(index)
	__style := C.int32_t(style)
	__retVal = bool(C.SetMenuItemStyle(__id, __index, __style))
	return __retVal
}

// SetMenuItemStyle 
//  @brief Changes an item's draw style.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//  @param style: The new style.
//
//  @return True if the item exists.
func SetMenuItemStyle(id MenuId, index int32, style MenuItemStyle) bool {
	return _SetMenuItemStyle(id, index, style)
}

var _DisplayMenu = func(id MenuId, playerSlot int32, time int32) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__playerSlot := C.int32_t(playerSlot)
	__time := C.int32_t(time)
	__retVal = bool(C.DisplayMenu(__id, __playerSlot, __time))
	return __retVal
}

// DisplayMenu 
//  @brief Displays a menu to a client, starting at the first item. Replaces whatever menu the client currently has open, if any.
//
//  @param id: The handle to the menu.
//  @param playerSlot: The client's player slot.
//  @param time: How long, in seconds, before the menu auto-closes (MenuCancelReason::Timeout). 0 or negative means no timeout.
//
//  @return True if the menu was displayed.
func DisplayMenu(id MenuId, playerSlot int32, time int32) bool {
	return _DisplayMenu(id, playerSlot, time)
}

var _DisplayMenuAtItem = func(id MenuId, playerSlot int32, firstItem int32, time int32) bool {
	var __retVal bool
	__id := C.uint32_t(id)
	__playerSlot := C.int32_t(playerSlot)
	__firstItem := C.int32_t(firstItem)
	__time := C.int32_t(time)
	__retVal = bool(C.DisplayMenuAtItem(__id, __playerSlot, __firstItem, __time))
	return __retVal
}

// DisplayMenuAtItem 
//  @brief Displays a menu to a client, starting at a specific item.
//
//  @param id: The handle to the menu.
//  @param playerSlot: The client's player slot.
//  @param firstItem: The index of the first item to show.
//  @param time: How long, in seconds, before the menu auto-closes. 0 or negative means no timeout.
//
//  @return True if the menu was displayed.
func DisplayMenuAtItem(id MenuId, playerSlot int32, firstItem int32, time int32) bool {
	return _DisplayMenuAtItem(id, playerSlot, firstItem, time)
}

var _CancelClientMenu = func(playerSlot int32, reason MenuCancelReason) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__reason := C.int32_t(reason)
	__retVal = bool(C.CancelClientMenu(__playerSlot, __reason))
	return __retVal
}

// CancelClientMenu 
//  @brief Cancels whatever menu a client currently has open.
//
//  @param playerSlot: The client's player slot.
//  @param reason: The reason reported to the menu's handler via MenuAction::Cancel.
//
//  @return True if the client had a menu open and it was cancelled.
func CancelClientMenu(playerSlot int32, reason MenuCancelReason) bool {
	return _CancelClientMenu(playerSlot, reason)
}

var _GetClientMenu = func(playerSlot int32) MenuId {
	var __retVal MenuId
	__playerSlot := C.int32_t(playerSlot)
	__retVal = MenuId(C.GetClientMenu(__playerSlot))
	return __retVal
}

// GetClientMenu 
//  @brief Gets the menu a client currently has open.
//
//  @param playerSlot: The client's player slot.
//
//  @return The open menu's handle, or 0 if none.
func GetClientMenu(playerSlot int32) MenuId {
	return _GetClientMenu(playerSlot)
}

var _GetClientMenuOffset = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMenuOffset(__playerSlot))
	return __retVal
}

// GetClientMenuOffset 
//  @brief Gets the index of the first item shown on the client's current page.
//
//  @param playerSlot: The client's player slot.
//
//  @return The current page offset.
func GetClientMenuOffset(playerSlot int32) int32 {
	return _GetClientMenuOffset(playerSlot)
}

var _GetClientMenuTime = func(playerSlot int32) float64 {
	var __retVal float64
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float64(C.GetClientMenuTime(__playerSlot))
	return __retVal
}

// GetClientMenuTime 
//  @brief Gets the `time` value the client's current menu was displayed with.
//
//  @param playerSlot: The client's player slot.
//
//  @return The display time in seconds, 0 meaning no timeout.
func GetClientMenuTime(playerSlot int32) float64 {
	return _GetClientMenuTime(playerSlot)
}

var _GetClientMenuCursor = func(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMenuCursor(__playerSlot))
	return __retVal
}

// GetClientMenuCursor 
//  @brief Gets the absolute item index highlighted by cursor-driven menu types (e.g. the button menu).
//
//  @param playerSlot: The client's player slot.
//
//  @return The cursor's item index.
func GetClientMenuCursor(playerSlot int32) int32 {
	return _GetClientMenuCursor(playerSlot)
}

var _SetClientMenuCursor = func(playerSlot int32, index int32) {
	__playerSlot := C.int32_t(playerSlot)
	__index := C.int32_t(index)
	C.SetClientMenuCursor(__playerSlot, __index)
}

// SetClientMenuCursor 
//  @brief Sets the absolute item index highlighted by cursor-driven menu types.
//
//  @param playerSlot: The client's player slot.
//  @param index: The item index to highlight.
func SetClientMenuCursor(playerSlot int32, index int32) {
	_SetClientMenuCursor(playerSlot, index)
}

var _ClientMenuHasPrevPage = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.ClientMenuHasPrevPage(__playerSlot))
	return __retVal
}

// ClientMenuHasPrevPage 
//  @brief Checks whether the client's current menu display has a previous page to go back to.
//
//  @param playerSlot: The client's player slot.
//
//  @return True if a previous page exists.
func ClientMenuHasPrevPage(playerSlot int32) bool {
	return _ClientMenuHasPrevPage(playerSlot)
}

var _ClientMenuHasNextPage = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.ClientMenuHasNextPage(__playerSlot))
	return __retVal
}

// ClientMenuHasNextPage 
//  @brief Checks whether the client's current menu display has a next page.
//
//  @param playerSlot: The client's player slot.
//
//  @return True if a next page exists.
func ClientMenuHasNextPage(playerSlot int32) bool {
	return _ClientMenuHasNextPage(playerSlot)
}

var _MenuNextPage = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.MenuNextPage(__playerSlot))
	return __retVal
}

// MenuNextPage 
//  @brief Advances the client's current menu display to the next page and redraws it.
//
//  @param playerSlot: The client's player slot.
//
//  @return True if there was a next page to move to.
func MenuNextPage(playerSlot int32) bool {
	return _MenuNextPage(playerSlot)
}

var _MenuPrevPage = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.MenuPrevPage(__playerSlot))
	return __retVal
}

// MenuPrevPage 
//  @brief Moves the client's current menu display back to the previous page and redraws it.
//
//  @param playerSlot: The client's player slot.
//
//  @return True if there was a previous page to move to.
func MenuPrevPage(playerSlot int32) bool {
	return _MenuPrevPage(playerSlot)
}

var _SelectMenuItem = func(playerSlot int32, itemIndex int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemIndex := C.int32_t(itemIndex)
	__retVal = bool(C.SelectMenuItem(__playerSlot, __itemIndex))
	return __retVal
}

// SelectMenuItem 
//  @brief Selects an item on the client's current menu display by its absolute index. Intended to be called by menu type backends once they've resolved raw input into an item index.
//
//  @param playerSlot: The client's player slot.
//  @param itemIndex: The absolute index of the item to select.
//
//  @return True if the item existed and was selectable.
func SelectMenuItem(playerSlot int32, itemIndex int32) bool {
	return _SelectMenuItem(playerSlot, itemIndex)
}

var _HandleDigitInput = func(playerSlot int32, digit int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__digit := C.int32_t(digit)
	__retVal = bool(C.HandleDigitInput(__playerSlot, __digit))
	return __retVal
}

// HandleDigitInput 
//  @brief Shared input path for digit-driven menu types (chat/console/centerhtml): 1-7 select an item on the current page, 8 goes to the previous page, 9 to the next page, 0 exits.
//
//  @param playerSlot: The client's player slot.
//  @param digit: The digit (0-9) that was pressed.
//
//  @return True if the digit resulted in an action.
func HandleDigitInput(playerSlot int32, digit int32) bool {
	return _HandleDigitInput(playerSlot, digit)
}

var (
	MenuErrEmptyHandle = errors.New("Menu: empty handle")
)

//  @brief RAII wrapper for Menu handle.
//
type Menu struct {
	handle    MenuId
}

// NewMenuCreateMenu 
//  @brief Creates a new menu.
//
//  @param title: The title shown at the top of the menu.
//  @param handler: Callback invoked with Start/Select/Cancel/End actions as the menu is displayed and interacted with.
//  @param menuType: The name of the menu type backend to render with. Empty uses the current default menu type.
func NewMenuCreateMenu(title string, handler MenuHandlerCallback, menuType string) *Menu {
	return &Menu{
		handle: CreateMenu(title, handler, menuType),
	}
}

// NewMenu creates a Menu from a handle
func NewMenu(handle MenuId) *Menu {
	return &Menu{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *Menu) Get() MenuId {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Menu) Release() MenuId {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Menu) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Menu) IsValid() bool {
	return w.handle != 0
}

// Destroy 
//  @brief Destroys a menu. Any client currently viewing it is cancelled first (MenuCancelReason::Destroyed).
//
//  @param id: The handle to the menu.
//
//  @return True if the menu existed and was destroyed.
func (w *Menu) Destroy() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return DestroyMenu(w.handle), nil
}

// IsValid 
//  @brief Checks whether a menu handle refers to an existing menu.
//
//  @param id: The handle to the menu.
//
//  @return True if the handle is valid.
func (w *Menu) IsValid2() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return IsValidMenu(w.handle), nil
}

// SetTitle 
//  @brief Sets a menu's title.
//
//  @param id: The handle to the menu.
//  @param title: The new title.
//
//  @return True if the menu exists.
func (w *Menu) SetTitle(title string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuTitle(w.handle, title), nil
}

// GetTitle 
//  @brief Gets a menu's title.
//
//  @param id: The handle to the menu.
//
//  @return The menu's title, or an empty string if the handle is invalid.
func (w *Menu) GetTitle() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, MenuErrEmptyHandle
	}
	return GetMenuTitle(w.handle), nil
}

// SetType 
//  @brief Sets which registered menu type backend renders this menu.
//
//  @param id: The handle to the menu.
//  @param typeName: The name of a registered menu type, or empty to use the default menu type.
//
//  @return True if the menu exists.
func (w *Menu) SetType(typeName string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuType(w.handle, typeName), nil
}

// GetType 
//  @brief Gets the menu type backend name assigned to this menu.
//
//  @param id: The handle to the menu.
//
//  @return The menu type name (may be empty, meaning "use the default").
func (w *Menu) GetType() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, MenuErrEmptyHandle
	}
	return GetMenuType(w.handle), nil
}

// SetPagination 
//  @brief Sets how many items are shown per page.
//
//  @param id: The handle to the menu.
//  @param itemsPerPage: The number of items per page, or 0 to disable pagination (show every item on one page).
//
//  @return True if the menu exists and itemsPerPage is not negative.
func (w *Menu) SetPagination(itemsPerPage int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuPagination(w.handle, itemsPerPage), nil
}

// GetPagination 
//  @brief Gets how many items are shown per page.
//
//  @param id: The handle to the menu.
//
//  @return The items-per-page value, 0 meaning pagination is disabled.
func (w *Menu) GetPagination() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, MenuErrEmptyHandle
	}
	return GetMenuPagination(w.handle), nil
}

// SetExitButton 
//  @brief Sets whether the menu shows an exit option.
//
//  @param id: The handle to the menu.
//  @param enabled: True to show an exit option.
//
//  @return True if the menu exists.
func (w *Menu) SetExitButton(enabled bool) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuExitButton(w.handle, enabled), nil
}

// GetExitButton 
//  @brief Gets whether the menu shows an exit option.
//
//  @param id: The handle to the menu.
//
//  @return True if the exit option is enabled.
func (w *Menu) GetExitButton() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return GetMenuExitButton(w.handle), nil
}

// SetExitBackButton 
//  @brief Sets whether the menu shows a "back" option in place of the exit option. Selecting it cancels the display with MenuCancelReason::ExitBack instead of MenuCancelReason::Exit, which a handler can use to redisplay a parent menu (SourceMod-style ExitBack).
//
//  @param id: The handle to the menu.
//  @param enabled: True to show a back option instead of the exit option.
//
//  @return True if the menu exists.
func (w *Menu) SetExitBackButton(enabled bool) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuExitBackButton(w.handle, enabled), nil
}

// GetExitBackButton 
//  @brief Gets whether the menu shows a "back" option in place of the exit option.
//
//  @param id: The handle to the menu.
//
//  @return True if the back option is enabled.
func (w *Menu) GetExitBackButton() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return GetMenuExitBackButton(w.handle), nil
}

// SetCloseOnSelect 
//  @brief Sets whether selecting an item automatically closes the menu display for that client. When disabled, the display stays open after MenuAction::Select and the handler is responsible for closing/redisplaying it if desired.
//
//  @param id: The handle to the menu.
//  @param enabled: True to auto-close on selection (the default).
//
//  @return True if the menu exists.
func (w *Menu) SetCloseOnSelect(enabled bool) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuCloseOnSelect(w.handle, enabled), nil
}

// GetCloseOnSelect 
//  @brief Gets whether selecting an item automatically closes the menu display for that client.
//
//  @param id: The handle to the menu.
//
//  @return True if close-on-select is enabled.
func (w *Menu) GetCloseOnSelect() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return GetMenuCloseOnSelect(w.handle), nil
}

// AddItem 
//  @brief Appends an item to the end of a menu.
//
//  @param id: The handle to the menu.
//  @param info: An internal identifier for the item, not shown to the client; retrieve it with GetMenuItemInfo from within the handler callback.
//  @param display: The text shown to the client.
//  @param style: The item's draw style (Default/Disabled/Spacer).
//
//  @return The index of the newly added item, or -1 if the menu handle is invalid.
func (w *Menu) AddItem(info string, display string, style MenuItemStyle) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, MenuErrEmptyHandle
	}
	return AddMenuItem(w.handle, info, display, style), nil
}

// InsertItemAt 
//  @brief Inserts an item into a menu at a specific index.
//
//  @param id: The handle to the menu.
//  @param index: The index to insert at; must be within [0, item count].
//  @param info: An internal identifier for the item, not shown to the client.
//  @param display: The text shown to the client.
//  @param style: The item's draw style (Default/Disabled/Spacer).
//
//  @return The index the item was inserted at, or -1 on failure.
func (w *Menu) InsertItemAt(index int32, info string, display string, style MenuItemStyle) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, MenuErrEmptyHandle
	}
	return InsertMenuItemAt(w.handle, index, info, display, style), nil
}

// RemoveItem 
//  @brief Removes an item from a menu.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item to remove.
//
//  @return True if the item existed and was removed.
func (w *Menu) RemoveItem(index int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return RemoveMenuItem(w.handle, index), nil
}

// RemoveAllItems 
//  @brief Removes every item from a menu.
//
//  @param id: The handle to the menu.
//
//  @return True if the menu exists.
func (w *Menu) RemoveAllItems() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return RemoveAllMenuItems(w.handle), nil
}

// GetItemsCount 
//  @brief Gets the number of items in a menu.
//
//  @param id: The handle to the menu.
//
//  @return The item count, or 0 if the handle is invalid.
func (w *Menu) GetItemsCount() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, MenuErrEmptyHandle
	}
	return GetMenuItemsCount(w.handle), nil
}

// GetItemInfoText 
//  @brief Gets an item's internal info string.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's info string, or empty if out of range.
func (w *Menu) GetItemInfoText(index int32) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, MenuErrEmptyHandle
	}
	return GetMenuItemInfoText(w.handle, index), nil
}

// GetItemDisplay 
//  @brief Gets an item's display text.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's display text, or empty if out of range.
func (w *Menu) GetItemDisplay(index int32) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, MenuErrEmptyHandle
	}
	return GetMenuItemDisplay(w.handle, index), nil
}

// GetItemStyle 
//  @brief Gets an item's draw style.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return The item's style; MenuItemStyle::Disabled if out of range.
func (w *Menu) GetItemStyle(index int32) (MenuItemStyle, error) {
	if w.handle == 0 {
		var zero MenuItemStyle
		return zero, MenuErrEmptyHandle
	}
	return GetMenuItemStyle(w.handle, index), nil
}

// IsItemSelectable 
//  @brief Checks whether an item can currently be selected (style is Default).
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//
//  @return True if the item is selectable.
func (w *Menu) IsItemSelectable(index int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return IsMenuItemSelectable(w.handle, index), nil
}

// SetItemDisplay 
//  @brief Changes an item's display text.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//  @param display: The new display text.
//
//  @return True if the item exists.
func (w *Menu) SetItemDisplay(index int32, display string) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuItemDisplay(w.handle, index, display), nil
}

// SetItemStyle 
//  @brief Changes an item's draw style.
//
//  @param id: The handle to the menu.
//  @param index: The index of the item.
//  @param style: The new style.
//
//  @return True if the item exists.
func (w *Menu) SetItemStyle(index int32, style MenuItemStyle) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return SetMenuItemStyle(w.handle, index, style), nil
}

// Display 
//  @brief Displays a menu to a client, starting at the first item. Replaces whatever menu the client currently has open, if any.
//
//  @param id: The handle to the menu.
//  @param playerSlot: The client's player slot.
//  @param time: How long, in seconds, before the menu auto-closes (MenuCancelReason::Timeout). 0 or negative means no timeout.
//
//  @return True if the menu was displayed.
func (w *Menu) Display(playerSlot int32, time int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return DisplayMenu(w.handle, playerSlot, time), nil
}

// DisplayAtItem 
//  @brief Displays a menu to a client, starting at a specific item.
//
//  @param id: The handle to the menu.
//  @param playerSlot: The client's player slot.
//  @param firstItem: The index of the first item to show.
//  @param time: How long, in seconds, before the menu auto-closes. 0 or negative means no timeout.
//
//  @return True if the menu was displayed.
func (w *Menu) DisplayAtItem(playerSlot int32, firstItem int32, time int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, MenuErrEmptyHandle
	}
	return DisplayMenuAtItem(w.handle, playerSlot, firstItem, time), nil
}

