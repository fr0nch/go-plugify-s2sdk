package s2sdk

/*
#include "items.h"
#cgo noescape GetItemDefinitionName
#cgo noescape GetItemDefinitionByName
#cgo noescape GetItemDefinitionClassName
#cgo noescape GetItemDefinitionLoadoutSlot
#cgo noescape GetItemDefinitionTeam
#cgo noescape GetItemDefinitionStickerSlots
#cgo noescape GetItemDefinitionKeychainSlots
#cgo noescape GetItemDefinitionImage
#cgo noescape GetItemDefinitionModel
#cgo noescape GetAttributeDefinitionIndex
#cgo noescape GetAttributeDefinitionName
#cgo noescape GetPaintKitName
#cgo noescape GetPaintKitLocName
#cgo noescape FindPaintKit
#cgo noescape GetStickerKitName
#cgo noescape GetStickerKitLocName
#cgo noescape FindStickerKit
#cgo noescape GetKeychainName
#cgo noescape GetKeychainLocName
#cgo noescape FindKeychain
#cgo noescape GetClientItemView
#cgo noescape GetClientLoadoutItemView
#cgo noescape GetEntityItemView
#cgo noescape GetItemViewDefinition
#cgo noescape GetItemViewId
#cgo noescape GetItemViewAccountId
#cgo noescape GetItemViewQuality
#cgo noescape GetItemViewRarity
#cgo noescape GetItemViewOrigin
#cgo noescape IsItemViewTradable
#cgo noescape IsItemViewMarketable
#cgo noescape GetItemViewPaintKit
#cgo noescape GetItemViewPaintKitSeed
#cgo noescape GetItemViewPaintKitWear
#cgo noescape GetItemViewStickerAttribute
#cgo noescape GetItemViewKeychainAttribute
#cgo noescape GetItemViewAttribute
#cgo noescape GiveClientItem
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

// Generated from s2sdk (group: items)

var _GetItemDefinitionName = func(itemDefinition uint16) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__itemDefinition := C.uint16_t(itemDefinition)
	plugify.Block {
		Try: func() {
			__native := C.GetItemDefinitionName(__itemDefinition)
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

// GetItemDefinitionName 
//  @brief Returns the name of an item definition, e.g. "weapon_ak47" or "weapon_knife_karambit".
//
//  @param itemDefinition: The item definition index.
//
//  @return The definition name, or an empty string.
func GetItemDefinitionName(itemDefinition uint16) string {
	return _GetItemDefinitionName(itemDefinition)
}

var _GetItemDefinitionByName = func(name string) uint16 {
	var __retVal uint16
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = uint16(C.GetItemDefinitionByName((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetItemDefinitionByName 
//  @brief Finds an item definition by its name, e.g. "weapon_ak47", "studded_brokenfang_gloves" or "customplayer_tm_leet_varianta".
//
//  @param name: The definition name.
//
//  @return The item definition index, or 0.
func GetItemDefinitionByName(name string) uint16 {
	return _GetItemDefinitionByName(name)
}

var _GetItemDefinitionClassName = func(itemDefinition uint16) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__itemDefinition := C.uint16_t(itemDefinition)
	plugify.Block {
		Try: func() {
			__native := C.GetItemDefinitionClassName(__itemDefinition)
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

// GetItemDefinitionClassName 
//  @brief Returns the item class of an item definition, e.g. "weapon_knife" for every knife.
//
//  @param itemDefinition: The item definition index.
//
//  @return The item class, or an empty string.
func GetItemDefinitionClassName(itemDefinition uint16) string {
	return _GetItemDefinitionClassName(itemDefinition)
}

var _GetItemDefinitionLoadoutSlot = func(itemDefinition uint16, team CSTeam) LoadoutSlot {
	var __retVal LoadoutSlot
	__itemDefinition := C.uint16_t(itemDefinition)
	__team := C.int32_t(team)
	__retVal = LoadoutSlot(C.GetItemDefinitionLoadoutSlot(__itemDefinition, __team))
	return __retVal
}

// GetItemDefinitionLoadoutSlot 
//  @brief Returns the loadout slot that an item definition takes for a team.
//
//  @param itemDefinition: The item definition index.
//  @param team: The team.
//
//  @return The loadout position, or Invalid.
func GetItemDefinitionLoadoutSlot(itemDefinition uint16, team CSTeam) LoadoutSlot {
	return _GetItemDefinitionLoadoutSlot(itemDefinition, team)
}

var _GetItemDefinitionTeam = func(itemDefinition uint16) CSTeam {
	var __retVal CSTeam
	__itemDefinition := C.uint16_t(itemDefinition)
	__retVal = CSTeam(C.GetItemDefinitionTeam(__itemDefinition))
	return __retVal
}

// GetItemDefinitionTeam 
//  @brief Returns the only team that can use an item definition.
//
//  @param itemDefinition: The item definition index.
//
//  @return The team, or CSTeam.None when both teams or neither can use it.
func GetItemDefinitionTeam(itemDefinition uint16) CSTeam {
	return _GetItemDefinitionTeam(itemDefinition)
}

var _GetItemDefinitionStickerSlots = func(itemDefinition uint16) int32 {
	var __retVal int32
	__itemDefinition := C.uint16_t(itemDefinition)
	__retVal = int32(C.GetItemDefinitionStickerSlots(__itemDefinition))
	return __retVal
}

// GetItemDefinitionStickerSlots 
//  @brief Returns how many stickers (or patches on an agent) an item definition takes.
//
//  @param itemDefinition: The item definition index.
//
//  @return The number of sticker slots.
func GetItemDefinitionStickerSlots(itemDefinition uint16) int32 {
	return _GetItemDefinitionStickerSlots(itemDefinition)
}

var _GetItemDefinitionKeychainSlots = func(itemDefinition uint16) int32 {
	var __retVal int32
	__itemDefinition := C.uint16_t(itemDefinition)
	__retVal = int32(C.GetItemDefinitionKeychainSlots(__itemDefinition))
	return __retVal
}

// GetItemDefinitionKeychainSlots 
//  @brief Returns how many keychains an item definition takes.
//
//  @param itemDefinition: The item definition index.
//
//  @return The number of keychain slots.
func GetItemDefinitionKeychainSlots(itemDefinition uint16) int32 {
	return _GetItemDefinitionKeychainSlots(itemDefinition)
}

var _GetItemDefinitionImage = func(itemDefinition uint16) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__itemDefinition := C.uint16_t(itemDefinition)
	plugify.Block {
		Try: func() {
			__native := C.GetItemDefinitionImage(__itemDefinition)
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

// GetItemDefinitionImage 
//  @brief Returns the inventory image of an item definition.
//
//  @param itemDefinition: The item definition index.
//
//  @return The image path, or an empty string.
func GetItemDefinitionImage(itemDefinition uint16) string {
	return _GetItemDefinitionImage(itemDefinition)
}

var _GetItemDefinitionModel = func(itemDefinition uint16, world bool) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__itemDefinition := C.uint16_t(itemDefinition)
	__world := C.bool(world)
	plugify.Block {
		Try: func() {
			__native := C.GetItemDefinitionModel(__itemDefinition, __world)
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

// GetItemDefinitionModel 
//  @brief Returns the model of an item definition.
//
//  @param itemDefinition: The item definition index.
//  @param world: True for the world model, false for the one the player holds.
//
//  @return The model path, or an empty string.
func GetItemDefinitionModel(itemDefinition uint16, world bool) string {
	return _GetItemDefinitionModel(itemDefinition, world)
}

var _GetAttributeDefinitionIndex = func(attribute string) int32 {
	var __retVal int32
	__attribute := plugify.ConstructString(attribute)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.GetAttributeDefinitionIndex((*C.String)(unsafe.Pointer(&__attribute))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__attribute)
		},
	}.Do()
	return __retVal
}

// GetAttributeDefinitionIndex 
//  @brief Returns the index of an attribute definition, e.g. 6 for "set item texture prefab".
//
//  @param attribute: The attribute definition name.
//
//  @return The attribute definition index, or -1.
func GetAttributeDefinitionIndex(attribute string) int32 {
	return _GetAttributeDefinitionIndex(attribute)
}

var _GetAttributeDefinitionName = func(attributeIndex int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__attributeIndex := C.int32_t(attributeIndex)
	plugify.Block {
		Try: func() {
			__native := C.GetAttributeDefinitionName(__attributeIndex)
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

// GetAttributeDefinitionName 
//  @brief Returns the name of an attribute definition.
//
//  @param attributeIndex: The attribute definition index.
//
//  @return The attribute definition name, or an empty string.
func GetAttributeDefinitionName(attributeIndex int32) string {
	return _GetAttributeDefinitionName(attributeIndex)
}

var _GetPaintKitName = func(paintKit int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__paintKit := C.int32_t(paintKit)
	plugify.Block {
		Try: func() {
			__native := C.GetPaintKitName(__paintKit)
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

// GetPaintKitName 
//  @brief Returns the schema name of a paint kit, e.g. "cu_medieval_dragon_awp".
//
//  @param paintKit: The paint kit index.
//
//  @return The name, or an empty string.
func GetPaintKitName(paintKit int32) string {
	return _GetPaintKitName(paintKit)
}

var _GetPaintKitLocName = func(paintKit int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__paintKit := C.int32_t(paintKit)
	plugify.Block {
		Try: func() {
			__native := C.GetPaintKitLocName(__paintKit)
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

// GetPaintKitLocName 
//  @brief Returns the localization token of a paint kit's name, e.g. "#PaintKit_cu_medieval_dragon_awp_Tag" (Dragon Lore). The server has no translations; the client shows the text for a token sent to it.
//
//  @param paintKit: The paint kit index.
//
//  @return The token, or an empty string.
func GetPaintKitLocName(paintKit int32) string {
	return _GetPaintKitLocName(paintKit)
}

var _FindPaintKit = func(name string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.FindPaintKit((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FindPaintKit 
//  @brief Finds a paint kit by its schema name.
//
//  @param name: The schema name, e.g. "cu_medieval_dragon_awp".
//
//  @return The paint kit index, or -1.
func FindPaintKit(name string) int32 {
	return _FindPaintKit(name)
}

var _GetStickerKitName = func(stickerKit int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__stickerKit := C.int32_t(stickerKit)
	plugify.Block {
		Try: func() {
			__native := C.GetStickerKitName(__stickerKit)
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

// GetStickerKitName 
//  @brief Returns the schema name of a sticker kit, e.g. "kat2014_ibuypower_holo".
//
//  @param stickerKit: The sticker kit index.
//
//  @return The name, or an empty string.
func GetStickerKitName(stickerKit int32) string {
	return _GetStickerKitName(stickerKit)
}

var _GetStickerKitLocName = func(stickerKit int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__stickerKit := C.int32_t(stickerKit)
	plugify.Block {
		Try: func() {
			__native := C.GetStickerKitLocName(__stickerKit)
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

// GetStickerKitLocName 
//  @brief Returns the localization token of a sticker kit's name, e.g. "#StickerKit_kat2014_ibuypower_holo". The server has no translations; the client shows the text for a token sent to it.
//
//  @param stickerKit: The sticker kit index.
//
//  @return The token, or an empty string.
func GetStickerKitLocName(stickerKit int32) string {
	return _GetStickerKitLocName(stickerKit)
}

var _FindStickerKit = func(name string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.FindStickerKit((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FindStickerKit 
//  @brief Finds a sticker kit by its schema name. The name is compared without case, as the game does.
//
//  @param name: The schema name, e.g. "kat2014_ibuypower_holo".
//
//  @return The sticker kit index, or -1.
func FindStickerKit(name string) int32 {
	return _FindStickerKit(name)
}

var _GetKeychainName = func(keychain int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__keychain := C.int32_t(keychain)
	plugify.Block {
		Try: func() {
			__native := C.GetKeychainName(__keychain)
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

// GetKeychainName 
//  @brief Returns the schema name of a keychain, e.g. "kc_missinglink_ava".
//
//  @param keychain: The keychain index.
//
//  @return The name, or an empty string.
func GetKeychainName(keychain int32) string {
	return _GetKeychainName(keychain)
}

var _GetKeychainLocName = func(keychain int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__keychain := C.int32_t(keychain)
	plugify.Block {
		Try: func() {
			__native := C.GetKeychainLocName(__keychain)
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

// GetKeychainLocName 
//  @brief Returns the localization token of a keychain's name, e.g. "#keychain_kc_missinglink_ava". The server has no translations; the client shows the text for a token sent to it.
//
//  @param keychain: The keychain index.
//
//  @return The token, or an empty string.
func GetKeychainLocName(keychain int32) string {
	return _GetKeychainLocName(keychain)
}

var _FindKeychain = func(name string) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.FindKeychain((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FindKeychain 
//  @brief Finds a keychain by its schema name.
//
//  @param name: The schema name, e.g. "kc_missinglink_ava".
//
//  @return The keychain index, or -1.
func FindKeychain(name string) int32 {
	return _FindKeychain(name)
}

var _GetClientItemView = func(playerSlot int32, itemId uint64) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__retVal = uintptr(C.GetClientItemView(__playerSlot, __itemId))
	return __retVal
}

// GetClientItemView 
//  @brief Returns the item view (CEconItemView) of an item in the client's inventory. Valid while the item stays in the inventory.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//
//  @return A pointer to the item view, or nullptr.
func GetClientItemView(playerSlot int32, itemId uint64) uintptr {
	return _GetClientItemView(playerSlot, itemId)
}

var _GetClientLoadoutItemView = func(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__loadoutSlot := C.int32_t(loadoutSlot)
	__retVal = uintptr(C.GetClientLoadoutItemView(__playerSlot, __team, __loadoutSlot))
	return __retVal
}

// GetClientLoadoutItemView 
//  @brief Returns the item view (CEconItemView) equipped in a loadout slot. An empty slot gives the default item of the slot. Valid while the inventory is loaded.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//
//  @return A pointer to the item view, or nullptr.
func GetClientLoadoutItemView(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uintptr {
	return _GetClientLoadoutItemView(playerSlot, team, loadoutSlot)
}

var _GetEntityItemView = func(entityHandle int32) uintptr {
	var __retVal uintptr
	__entityHandle := C.int32_t(entityHandle)
	__retVal = uintptr(C.GetEntityItemView(__entityHandle))
	return __retVal
}

// GetEntityItemView 
//  @brief Returns the item view (CEconItemView) of a weapon or another econ entity. Valid while the entity exists.
//
//  @param entityHandle: The handle of the entity.
//
//  @return A pointer to the item view, or nullptr.
func GetEntityItemView(entityHandle int32) uintptr {
	return _GetEntityItemView(entityHandle)
}

var _GetItemViewDefinition = func(itemView uintptr) uint16 {
	var __retVal uint16
	__itemView := C.uintptr_t(itemView)
	__retVal = uint16(C.GetItemViewDefinition(__itemView))
	return __retVal
}

// GetItemViewDefinition 
//  @brief Returns the item definition index of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The item definition index.
func GetItemViewDefinition(itemView uintptr) uint16 {
	return _GetItemViewDefinition(itemView)
}

var _GetItemViewId = func(itemView uintptr) uint64 {
	var __retVal uint64
	__itemView := C.uintptr_t(itemView)
	__retVal = uint64(C.GetItemViewId(__itemView))
	return __retVal
}

// GetItemViewId 
//  @brief Returns the item ID of an item view, 0 for a default item.
//
//  @param itemView: A pointer to the item view.
//
//  @return The item ID.
func GetItemViewId(itemView uintptr) uint64 {
	return _GetItemViewId(itemView)
}

var _GetItemViewAccountId = func(itemView uintptr) uint32 {
	var __retVal uint32
	__itemView := C.uintptr_t(itemView)
	__retVal = uint32(C.GetItemViewAccountId(__itemView))
	return __retVal
}

// GetItemViewAccountId 
//  @brief Returns the Steam account ID of the owner of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The account ID.
func GetItemViewAccountId(itemView uintptr) uint32 {
	return _GetItemViewAccountId(itemView)
}

var _GetItemViewQuality = func(itemView uintptr) ItemQuality {
	var __retVal ItemQuality
	__itemView := C.uintptr_t(itemView)
	__retVal = ItemQuality(C.GetItemViewQuality(__itemView))
	return __retVal
}

// GetItemViewQuality 
//  @brief Returns the quality of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The quality.
func GetItemViewQuality(itemView uintptr) ItemQuality {
	return _GetItemViewQuality(itemView)
}

var _GetItemViewRarity = func(itemView uintptr) ItemRarity {
	var __retVal ItemRarity
	__itemView := C.uintptr_t(itemView)
	__retVal = ItemRarity(C.GetItemViewRarity(__itemView))
	return __retVal
}

// GetItemViewRarity 
//  @brief Returns the rarity of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The rarity.
func GetItemViewRarity(itemView uintptr) ItemRarity {
	return _GetItemViewRarity(itemView)
}

var _GetItemViewOrigin = func(itemView uintptr) ItemOrigin {
	var __retVal ItemOrigin
	__itemView := C.uintptr_t(itemView)
	__retVal = ItemOrigin(C.GetItemViewOrigin(__itemView))
	return __retVal
}

// GetItemViewOrigin 
//  @brief Returns how the item of an item view was obtained (eEconItemOrigin).
//
//  @param itemView: A pointer to the item view.
//
//  @return The origin.
func GetItemViewOrigin(itemView uintptr) ItemOrigin {
	return _GetItemViewOrigin(itemView)
}

var _IsItemViewTradable = func(itemView uintptr) bool {
	var __retVal bool
	__itemView := C.uintptr_t(itemView)
	__retVal = bool(C.IsItemViewTradable(__itemView))
	return __retVal
}

// IsItemViewTradable 
//  @brief Returns whether the item of an item view can be traded.
//
//  @param itemView: A pointer to the item view.
//
//  @return True if it is tradable.
func IsItemViewTradable(itemView uintptr) bool {
	return _IsItemViewTradable(itemView)
}

var _IsItemViewMarketable = func(itemView uintptr) bool {
	var __retVal bool
	__itemView := C.uintptr_t(itemView)
	__retVal = bool(C.IsItemViewMarketable(__itemView))
	return __retVal
}

// IsItemViewMarketable 
//  @brief Returns whether the item of an item view can be sold on the market.
//
//  @param itemView: A pointer to the item view.
//
//  @return True if it is marketable.
func IsItemViewMarketable(itemView uintptr) bool {
	return _IsItemViewMarketable(itemView)
}

var _GetItemViewPaintKit = func(itemView uintptr) int32 {
	var __retVal int32
	__itemView := C.uintptr_t(itemView)
	__retVal = int32(C.GetItemViewPaintKit(__itemView))
	return __retVal
}

// GetItemViewPaintKit 
//  @brief Returns the paint kit of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The paint kit index, 0 without a paint.
func GetItemViewPaintKit(itemView uintptr) int32 {
	return _GetItemViewPaintKit(itemView)
}

var _GetItemViewPaintKitSeed = func(itemView uintptr) int32 {
	var __retVal int32
	__itemView := C.uintptr_t(itemView)
	__retVal = int32(C.GetItemViewPaintKitSeed(__itemView))
	return __retVal
}

// GetItemViewPaintKitSeed 
//  @brief Returns the paint seed of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The seed.
func GetItemViewPaintKitSeed(itemView uintptr) int32 {
	return _GetItemViewPaintKitSeed(itemView)
}

var _GetItemViewPaintKitWear = func(itemView uintptr) float32 {
	var __retVal float32
	__itemView := C.uintptr_t(itemView)
	__retVal = float32(C.GetItemViewPaintKitWear(__itemView))
	return __retVal
}

// GetItemViewPaintKitWear 
//  @brief Returns the paint wear of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The wear, 0 without a paint.
func GetItemViewPaintKitWear(itemView uintptr) float32 {
	return _GetItemViewPaintKitWear(itemView)
}

var _GetItemViewStickerAttribute = func(itemView uintptr, stickerSlot int32, attribute StickerAttribute) float64 {
	var __retVal float64
	__itemView := C.uintptr_t(itemView)
	__stickerSlot := C.int32_t(stickerSlot)
	__attribute := C.int32_t(attribute)
	__retVal = float64(C.GetItemViewStickerAttribute(__itemView, __stickerSlot, __attribute))
	return __retVal
}

// GetItemViewStickerAttribute 
//  @brief Returns an attribute of a sticker on an item view.
//
//  @param itemView: A pointer to the item view.
//  @param stickerSlot: The sticker slot.
//  @param attribute: The sticker attribute.
//
//  @return The value, 0 if the slot is empty.
func GetItemViewStickerAttribute(itemView uintptr, stickerSlot int32, attribute StickerAttribute) float64 {
	return _GetItemViewStickerAttribute(itemView, stickerSlot, attribute)
}

var _GetItemViewKeychainAttribute = func(itemView uintptr, keychainSlot int32, attribute StickerAttribute) float64 {
	var __retVal float64
	__itemView := C.uintptr_t(itemView)
	__keychainSlot := C.int32_t(keychainSlot)
	__attribute := C.int32_t(attribute)
	__retVal = float64(C.GetItemViewKeychainAttribute(__itemView, __keychainSlot, __attribute))
	return __retVal
}

// GetItemViewKeychainAttribute 
//  @brief Returns an attribute of a keychain on an item view.
//
//  @param itemView: A pointer to the item view.
//  @param keychainSlot: The keychain slot.
//  @param attribute: The keychain attribute.
//
//  @return The value, 0 if the slot is empty.
func GetItemViewKeychainAttribute(itemView uintptr, keychainSlot int32, attribute StickerAttribute) float64 {
	return _GetItemViewKeychainAttribute(itemView, keychainSlot, attribute)
}

var _GetItemViewAttribute = func(itemView uintptr, attribute string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__itemView := C.uintptr_t(itemView)
	__attribute := plugify.ConstructString(attribute)
	plugify.Block {
		Try: func() {
			__native := C.GetItemViewAttribute(__itemView, (*C.String)(unsafe.Pointer(&__attribute)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__attribute)
		},
	}.Do()
	return __retVal
}

// GetItemViewAttribute 
//  @brief Returns any attribute of an item view as text, e.g. "kill eater" (StatTrak), "custom name attr" or "sticker slot 0 id".
//
//  @param itemView: A pointer to the item view.
//  @param attribute: The attribute definition name.
//
//  @return The value, or an empty string if the item has no such attribute.
func GetItemViewAttribute(itemView uintptr, attribute string) string {
	return _GetItemViewAttribute(itemView, attribute)
}

var _GiveClientItem = func(playerSlot int32, itemId uint64) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__retVal = int32(C.GiveClientItem(__playerSlot, __itemId))
	return __retVal
}

// GiveClientItem 
//  @brief Gives the client a weapon made from an item of their inventory, with its paint, stickers and name.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//
//  @return The weapon's entity handle, or INVALID_EHANDLE_INDEX.
func GiveClientItem(playerSlot int32, itemId uint64) int32 {
	return _GiveClientItem(playerSlot, itemId)
}

var (
	ItemDefinitionErrEmptyHandle = errors.New("ItemDefinition: empty handle")
)

//  @brief Wrapper for an item definition index from the item schema.
//
type ItemDefinition struct {
	handle    uint16
}

// NewItemDefinition creates a ItemDefinition from a handle
func NewItemDefinition(handle uint16) *ItemDefinition {
	return &ItemDefinition{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *ItemDefinition) Get() uint16 {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *ItemDefinition) Release() uint16 {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *ItemDefinition) Reset() {
	w.handle = 0
}

// Valid returns true if handle is not nil
func (w *ItemDefinition) Valid() bool {
	return w.handle != 0
}

// FindByName 
//  @brief Finds an item definition by its name, e.g. "weapon_ak47", "studded_brokenfang_gloves" or "customplayer_tm_leet_varianta".
//
//  @param name: The definition name.
//
//  @return The item definition index, or 0.
func (w *ItemDefinition) FindByName(name string) *ItemDefinition {
	return NewItemDefinition(GetItemDefinitionByName(name))
}

// GetName 
//  @brief Returns the name of an item definition, e.g. "weapon_ak47" or "weapon_knife_karambit".
//
//  @param itemDefinition: The item definition index.
//
//  @return The definition name, or an empty string.
func (w *ItemDefinition) GetName() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionName(w.handle), nil
}

// GetClassName 
//  @brief Returns the item class of an item definition, e.g. "weapon_knife" for every knife.
//
//  @param itemDefinition: The item definition index.
//
//  @return The item class, or an empty string.
func (w *ItemDefinition) GetClassName() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionClassName(w.handle), nil
}

// GetLoadoutSlot 
//  @brief Returns the loadout slot that an item definition takes for a team.
//
//  @param itemDefinition: The item definition index.
//  @param team: The team.
//
//  @return The loadout position, or Invalid.
func (w *ItemDefinition) GetLoadoutSlot(team CSTeam) (LoadoutSlot, error) {
	if w.handle == 0 {
		var zero LoadoutSlot
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionLoadoutSlot(w.handle, team), nil
}

// GetTeam 
//  @brief Returns the only team that can use an item definition.
//
//  @param itemDefinition: The item definition index.
//
//  @return The team, or CSTeam.None when both teams or neither can use it.
func (w *ItemDefinition) GetTeam() (CSTeam, error) {
	if w.handle == 0 {
		var zero CSTeam
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionTeam(w.handle), nil
}

// GetStickerSlots 
//  @brief Returns how many stickers (or patches on an agent) an item definition takes.
//
//  @param itemDefinition: The item definition index.
//
//  @return The number of sticker slots.
func (w *ItemDefinition) GetStickerSlots() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionStickerSlots(w.handle), nil
}

// GetKeychainSlots 
//  @brief Returns how many keychains an item definition takes.
//
//  @param itemDefinition: The item definition index.
//
//  @return The number of keychain slots.
func (w *ItemDefinition) GetKeychainSlots() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionKeychainSlots(w.handle), nil
}

// GetImage 
//  @brief Returns the inventory image of an item definition.
//
//  @param itemDefinition: The item definition index.
//
//  @return The image path, or an empty string.
func (w *ItemDefinition) GetImage() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionImage(w.handle), nil
}

// GetModel 
//  @brief Returns the model of an item definition.
//
//  @param itemDefinition: The item definition index.
//  @param world: True for the world model, false for the one the player holds.
//
//  @return The model path, or an empty string.
func (w *ItemDefinition) GetModel(world bool) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ItemDefinitionErrEmptyHandle
	}
	return GetItemDefinitionModel(w.handle, world), nil
}

var (
	ItemViewErrEmptyHandle = errors.New("ItemView: empty handle")
)

//  @brief Wrapper for a CEconItemView pointer. It does not own the view: a view of an inventory item is gone when the item is removed or the inventory is unloaded, a view of an entity is gone with the entity.
//
type ItemView struct {
	handle    uintptr
}

// NewItemView creates a ItemView from a handle
func NewItemView(handle uintptr) *ItemView {
	return &ItemView{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *ItemView) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *ItemView) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *ItemView) Reset() {
	w.handle = 0
}

// Valid returns true if handle is not nil
func (w *ItemView) Valid() bool {
	return w.handle != 0
}

// FromClient 
//  @brief Returns the item view (CEconItemView) of an item in the client's inventory. Valid while the item stays in the inventory.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//
//  @return A pointer to the item view, or nullptr.
func (w *ItemView) FromClient(playerSlot int32, itemId uint64) *ItemView {
	return NewItemView(GetClientItemView(playerSlot, itemId))
}

// FromLoadout 
//  @brief Returns the item view (CEconItemView) equipped in a loadout slot. An empty slot gives the default item of the slot. Valid while the inventory is loaded.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//
//  @return A pointer to the item view, or nullptr.
func (w *ItemView) FromLoadout(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) *ItemView {
	return NewItemView(GetClientLoadoutItemView(playerSlot, team, loadoutSlot))
}

// FromEntity 
//  @brief Returns the item view (CEconItemView) of a weapon or another econ entity. Valid while the entity exists.
//
//  @param entityHandle: The handle of the entity.
//
//  @return A pointer to the item view, or nullptr.
func (w *ItemView) FromEntity(entityHandle int32) *ItemView {
	return NewItemView(GetEntityItemView(entityHandle))
}

// GetDefinition 
//  @brief Returns the item definition index of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The item definition index.
func (w *ItemView) GetDefinition() (*ItemDefinition, error) {
	if w.handle == 0 {
		var zero *ItemDefinition
		return zero, ItemViewErrEmptyHandle
	}
	return NewItemDefinition(GetItemViewDefinition(w.handle)), nil
}

// GetId 
//  @brief Returns the item ID of an item view, 0 for a default item.
//
//  @param itemView: A pointer to the item view.
//
//  @return The item ID.
func (w *ItemView) GetId() (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewId(w.handle), nil
}

// GetAccountId 
//  @brief Returns the Steam account ID of the owner of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The account ID.
func (w *ItemView) GetAccountId() (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewAccountId(w.handle), nil
}

// GetQuality 
//  @brief Returns the quality of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The quality.
func (w *ItemView) GetQuality() (ItemQuality, error) {
	if w.handle == 0 {
		var zero ItemQuality
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewQuality(w.handle), nil
}

// GetRarity 
//  @brief Returns the rarity of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The rarity.
func (w *ItemView) GetRarity() (ItemRarity, error) {
	if w.handle == 0 {
		var zero ItemRarity
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewRarity(w.handle), nil
}

// GetOrigin 
//  @brief Returns how the item of an item view was obtained (eEconItemOrigin).
//
//  @param itemView: A pointer to the item view.
//
//  @return The origin.
func (w *ItemView) GetOrigin() (ItemOrigin, error) {
	if w.handle == 0 {
		var zero ItemOrigin
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewOrigin(w.handle), nil
}

// IsTradable 
//  @brief Returns whether the item of an item view can be traded.
//
//  @param itemView: A pointer to the item view.
//
//  @return True if it is tradable.
func (w *ItemView) IsTradable() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ItemViewErrEmptyHandle
	}
	return IsItemViewTradable(w.handle), nil
}

// IsMarketable 
//  @brief Returns whether the item of an item view can be sold on the market.
//
//  @param itemView: A pointer to the item view.
//
//  @return True if it is marketable.
func (w *ItemView) IsMarketable() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ItemViewErrEmptyHandle
	}
	return IsItemViewMarketable(w.handle), nil
}

// GetPaintKit 
//  @brief Returns the paint kit of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The paint kit index, 0 without a paint.
func (w *ItemView) GetPaintKit() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewPaintKit(w.handle), nil
}

// GetPaintKitSeed 
//  @brief Returns the paint seed of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The seed.
func (w *ItemView) GetPaintKitSeed() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewPaintKitSeed(w.handle), nil
}

// GetPaintKitWear 
//  @brief Returns the paint wear of an item view.
//
//  @param itemView: A pointer to the item view.
//
//  @return The wear, 0 without a paint.
func (w *ItemView) GetPaintKitWear() (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewPaintKitWear(w.handle), nil
}

// GetStickerAttribute 
//  @brief Returns an attribute of a sticker on an item view.
//
//  @param itemView: A pointer to the item view.
//  @param stickerSlot: The sticker slot.
//  @param attribute: The sticker attribute.
//
//  @return The value, 0 if the slot is empty.
func (w *ItemView) GetStickerAttribute(stickerSlot int32, attribute StickerAttribute) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewStickerAttribute(w.handle, stickerSlot, attribute), nil
}

// GetKeychainAttribute 
//  @brief Returns an attribute of a keychain on an item view.
//
//  @param itemView: A pointer to the item view.
//  @param keychainSlot: The keychain slot.
//  @param attribute: The keychain attribute.
//
//  @return The value, 0 if the slot is empty.
func (w *ItemView) GetKeychainAttribute(keychainSlot int32, attribute StickerAttribute) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewKeychainAttribute(w.handle, keychainSlot, attribute), nil
}

// GetAttribute 
//  @brief Returns any attribute of an item view as text, e.g. "kill eater" (StatTrak), "custom name attr" or "sticker slot 0 id".
//
//  @param itemView: A pointer to the item view.
//  @param attribute: The attribute definition name.
//
//  @return The value, or an empty string if the item has no such attribute.
func (w *ItemView) GetAttribute(attribute string) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ItemViewErrEmptyHandle
	}
	return GetItemViewAttribute(w.handle, attribute), nil
}

