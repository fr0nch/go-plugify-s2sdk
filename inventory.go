package s2sdk

/*
#include "inventory.h"
#cgo noescape GetClientInventory
#cgo noescape GetClientLoadoutItemId
#cgo noescape GetClientLoadoutItemDefinition
#cgo noescape SetClientLoadoutItem
#cgo noescape SetClientLoadoutDefinition
#cgo noescape FindClientLoadoutSlot
#cgo noescape RestoreClientLoadout
#cgo noescape GiveClientDefaultWearables
#cgo noescape GetClientInventoryItems
#cgo noescape CreateClientItem
#cgo noescape RemoveClientItem
#cgo noescape SetClientItemAttribute
#cgo noescape SetClientItemPaint
#cgo noescape SetClientItemSticker
#cgo noescape SetClientItemKeychain
#cgo noescape SetClientItemStatTrak
#cgo noescape SetClientItemName
#cgo noescape GetClientItemAttribute
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

// Generated from s2sdk (group: inventory)

var _GetClientInventory = func(playerSlot int32) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uintptr(C.GetClientInventory(__playerSlot))
	return __retVal
}

// GetClientInventory 
//  @brief Returns the client's game inventory (CCSPlayerInventory). Available once OnClientInventoryLoaded fires; bots have no inventory.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return A pointer to the inventory, or nullptr if it is not loaded.
func GetClientInventory(playerSlot int32) uintptr {
	return _GetClientInventory(playerSlot)
}

var _GetClientLoadoutItemId = func(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__loadoutSlot := C.int32_t(loadoutSlot)
	__retVal = uint64(C.GetClientLoadoutItemId(__playerSlot, __team, __loadoutSlot))
	return __retVal
}

// GetClientLoadoutItemId 
//  @brief Returns the item ID equipped in a loadout slot. An empty slot gives the default item of the slot, whose ID is 0.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//
//  @return The item ID, or 0.
func GetClientLoadoutItemId(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uint64 {
	return _GetClientLoadoutItemId(playerSlot, team, loadoutSlot)
}

var _GetClientLoadoutItemDefinition = func(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uint16 {
	var __retVal uint16
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__loadoutSlot := C.int32_t(loadoutSlot)
	__retVal = uint16(C.GetClientLoadoutItemDefinition(__playerSlot, __team, __loadoutSlot))
	return __retVal
}

// GetClientLoadoutItemDefinition 
//  @brief Returns the item definition index equipped in a loadout slot. An empty slot gives the default item of the slot.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//
//  @return The item definition index, or 0.
func GetClientLoadoutItemDefinition(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot) uint16 {
	return _GetClientLoadoutItemDefinition(playerSlot, team, loadoutSlot)
}

var _SetClientLoadoutItem = func(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot, itemId uint64) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__loadoutSlot := C.int32_t(loadoutSlot)
	__itemId := C.uint64_t(itemId)
	__retVal = bool(C.SetClientLoadoutItem(__playerSlot, __team, __loadoutSlot, __itemId))
	return __retVal
}

// SetClientLoadoutItem 
//  @brief Equips an item of the client's inventory in a loadout slot. The game reads the loadout when it gives items, so already given weapons keep their look; it rewrites the loadout when the inventory is loaded again or the player changes it. The inventory update event is sent afterwards, as after an equip from the game coordinator.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//  @param itemId: The item ID, it must be in the client's inventory and belong to the slot or its group.
//
//  @return True if the item was equipped.
func SetClientLoadoutItem(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot, itemId uint64) bool {
	return _SetClientLoadoutItem(playerSlot, team, loadoutSlot, itemId)
}

var _SetClientLoadoutDefinition = func(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot, itemDefinition uint16) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__loadoutSlot := C.int32_t(loadoutSlot)
	__itemDefinition := C.uint16_t(itemDefinition)
	__retVal = bool(C.SetClientLoadoutDefinition(__playerSlot, __team, __loadoutSlot, __itemDefinition))
	return __retVal
}

// SetClientLoadoutDefinition 
//  @brief Equips a stock item definition in a loadout slot without an owned item, e.g. picks USP-S over P2000. The definition must belong to the slot or its group. Works like SetClientLoadoutItem otherwise.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team; music kits and flair use CSTeam.None.
//  @param loadoutSlot: The loadout position.
//  @param itemDefinition: The item definition index.
//
//  @return True if the definition was equipped.
func SetClientLoadoutDefinition(playerSlot int32, team CSTeam, loadoutSlot LoadoutSlot, itemDefinition uint16) bool {
	return _SetClientLoadoutDefinition(playerSlot, team, loadoutSlot, itemDefinition)
}

var _FindClientLoadoutSlot = func(playerSlot int32, team CSTeam, itemDefinition uint16) LoadoutSlot {
	var __retVal LoadoutSlot
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	__itemDefinition := C.uint16_t(itemDefinition)
	__retVal = LoadoutSlot(C.FindClientLoadoutSlot(__playerSlot, __team, __itemDefinition))
	return __retVal
}

// FindClientLoadoutSlot 
//  @brief Finds the loadout slot that gives an item definition to a team, e.g. which rifle slot holds the AK-47. Checks the slots the definition can take: its own slot, or every slot of its group for pistols, mid-tier weapons and rifles. Equip into the returned slot to keep one item per weapon: with two items of one weapon the game picks one at random on every give.
//
//  @param playerSlot: The index of the player's slot.
//  @param team: The loadout team.
//  @param itemDefinition: The item definition index.
//
//  @return The slot whose equipped or stock item has the definition, or LoadoutSlot.Invalid.
func FindClientLoadoutSlot(playerSlot int32, team CSTeam, itemDefinition uint16) LoadoutSlot {
	return _FindClientLoadoutSlot(playerSlot, team, itemDefinition)
}

var _RestoreClientLoadout = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.RestoreClientLoadout(__playerSlot))
	return __retVal
}

// RestoreClientLoadout 
//  @brief Undoes the loadout changes made through s2sdk and puts back what the game equipped.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return True if the loadout was restored.
func RestoreClientLoadout(playerSlot int32) bool {
	return _RestoreClientLoadout(playerSlot)
}

var _GiveClientDefaultWearables = func(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.GiveClientDefaultWearables(__playerSlot))
	return __retVal
}

// GiveClientDefaultWearables 
//  @brief Puts the client's loadout wearables on the pawn at once, as the game does at spawn: agent patches from the agent slot and gloves from the gloves slot.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return True if the wearables were applied, false if the player has no pawn.
func GiveClientDefaultWearables(playerSlot int32) bool {
	return _GiveClientDefaultWearables(playerSlot)
}

var _GetClientInventoryItems = func(playerSlot int32) []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block {
		Try: func() {
			__native := C.GetClientInventoryItems(__playerSlot)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientInventoryItems 
//  @brief Returns the IDs of the items in the client's inventory.
//
//  @param playerSlot: The index of the player's slot.
//
//  @return The item IDs, empty if the inventory is not loaded.
func GetClientInventoryItems(playerSlot int32) []uint64 {
	return _GetClientInventoryItems(playerSlot)
}

var _CreateClientItem = func(playerSlot int32, itemDefinition uint16, quality ItemQuality) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__itemDefinition := C.uint16_t(itemDefinition)
	__quality := C.int32_t(quality)
	__retVal = uint64(C.CreateClientItem(__playerSlot, __itemDefinition, __quality))
	return __retVal
}

// CreateClientItem 
//  @brief Creates an item that the client does not own and puts it into the client's inventory. The item lives until the client leaves, the map changes or the game coordinator sends the inventory again, so create items again in OnClientInventoryLoaded. Equip it with SetClientLoadoutItem. Refused while the "FollowCS2ServerGuidelines" option is enabled: Valve's server rules forbid giving players items they do not own. Attributes can be set only on created items.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemDefinition: The item definition index.
//  @param quality: The item quality.
//
//  @return The new item ID, or 0 on failure.
func CreateClientItem(playerSlot int32, itemDefinition uint16, quality ItemQuality) uint64 {
	return _CreateClientItem(playerSlot, itemDefinition, quality)
}

var _RemoveClientItem = func(playerSlot int32, itemId uint64) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__retVal = bool(C.RemoveClientItem(__playerSlot, __itemId))
	return __retVal
}

// RemoveClientItem 
//  @brief Removes an item made by CreateClientItem from the client's inventory. Loadout slots that held the item get back what the game equipped there.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//
//  @return True if the item was removed.
func RemoveClientItem(playerSlot int32, itemId uint64) bool {
	return _RemoveClientItem(playerSlot, itemId)
}

var _SetClientItemAttribute = func(playerSlot int32, itemId uint64, attribute string, value string) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__attribute := plugify.ConstructString(attribute)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetClientItemAttribute(__playerSlot, __itemId, (*C.String)(unsafe.Pointer(&__attribute)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__attribute)
			plugify.DestroyString(&__value)
		},
	}.Do()
	return __retVal
}

// SetClientItemAttribute 
//  @brief Sets an attribute of an item made by CreateClientItem, e.g. "set item texture prefab" (paint kit), "set item texture seed", "set item texture wear", "kill eater" or "custom name attr". The value is text that the attribute's type parses, as in the item schema: "38", "0.01" or a name.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param attribute: The attribute definition name.
//  @param value: The value.
//
//  @return True if the attribute was set.
func SetClientItemAttribute(playerSlot int32, itemId uint64, attribute string, value string) bool {
	return _SetClientItemAttribute(playerSlot, itemId, attribute, value)
}

var _SetClientItemPaint = func(playerSlot int32, itemId uint64, paintKit int32, seed int32, wear float32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__paintKit := C.int32_t(paintKit)
	__seed := C.int32_t(seed)
	__wear := C.float(wear)
	__retVal = bool(C.SetClientItemPaint(__playerSlot, __itemId, __paintKit, __seed, __wear))
	return __retVal
}

// SetClientItemPaint 
//  @brief Sets the skin of an item made by CreateClientItem.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param paintKit: The paint kit index, e.g. 259 for Howl.
//  @param seed: The pattern seed, 0 to 1000.
//  @param wear: The wear, 0.0 (factory new) to 1.0.
//
//  @return True if the skin was set.
func SetClientItemPaint(playerSlot int32, itemId uint64, paintKit int32, seed int32, wear float32) bool {
	return _SetClientItemPaint(playerSlot, itemId, paintKit, seed, wear)
}

var _SetClientItemSticker = func(playerSlot int32, itemId uint64, stickerSlot int32, stickerId int32, wear float32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__stickerSlot := C.int32_t(stickerSlot)
	__stickerId := C.int32_t(stickerId)
	__wear := C.float(wear)
	__retVal = bool(C.SetClientItemSticker(__playerSlot, __itemId, __stickerSlot, __stickerId, __wear))
	return __retVal
}

// SetClientItemSticker 
//  @brief Puts a sticker on an item made by CreateClientItem. The client draws stickers of a created item after a full update, see ForceClientFullUpdate.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param stickerSlot: The sticker slot, from 0 to GetItemDefinitionStickerSlots - 1.
//  @param stickerId: The sticker kit index, e.g. 60 for iBUYPOWER (Holo) | Katowice 2014.
//  @param wear: The sticker wear, 0.0 (new) to 1.0 (scraped off).
//
//  @return True if the sticker was set.
func SetClientItemSticker(playerSlot int32, itemId uint64, stickerSlot int32, stickerId int32, wear float32) bool {
	return _SetClientItemSticker(playerSlot, itemId, stickerSlot, stickerId, wear)
}

var _SetClientItemKeychain = func(playerSlot int32, itemId uint64, keychainSlot int32, keychainId int32, seed int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__keychainSlot := C.int32_t(keychainSlot)
	__keychainId := C.int32_t(keychainId)
	__seed := C.int32_t(seed)
	__retVal = bool(C.SetClientItemKeychain(__playerSlot, __itemId, __keychainSlot, __keychainId, __seed))
	return __retVal
}

// SetClientItemKeychain 
//  @brief Puts a keychain (charm) on an item made by CreateClientItem.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param keychainSlot: The keychain slot, from 0 to GetItemDefinitionKeychainSlots - 1.
//  @param keychainId: The keychain definition index.
//  @param seed: The keychain pattern seed.
//
//  @return True if the keychain was set.
func SetClientItemKeychain(playerSlot int32, itemId uint64, keychainSlot int32, keychainId int32, seed int32) bool {
	return _SetClientItemKeychain(playerSlot, itemId, keychainSlot, keychainId, seed)
}

var _SetClientItemStatTrak = func(playerSlot int32, itemId uint64, kills int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__kills := C.int32_t(kills)
	__retVal = bool(C.SetClientItemStatTrak(__playerSlot, __itemId, __kills))
	return __retVal
}

// SetClientItemStatTrak 
//  @brief Makes an item made by CreateClientItem a StatTrak item with the given kill count. The game does not count kills for created items, the plugin counts them and calls this again. Create the item with quality 9 to show the StatTrak name.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param kills: The kill count.
//
//  @return True if the counter was set.
func SetClientItemStatTrak(playerSlot int32, itemId uint64, kills int32) bool {
	return _SetClientItemStatTrak(playerSlot, itemId, kills)
}

var _SetClientItemName = func(playerSlot int32, itemId uint64, name string) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__name := plugify.ConstructString(name)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.SetClientItemName(__playerSlot, __itemId, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// SetClientItemName 
//  @brief Sets the name tag of an item made by CreateClientItem.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param name: The name.
//
//  @return True if the name was set.
func SetClientItemName(playerSlot int32, itemId uint64, name string) bool {
	return _SetClientItemName(playerSlot, itemId, name)
}

var _GetClientItemAttribute = func(playerSlot int32, itemId uint64, attribute string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	__itemId := C.uint64_t(itemId)
	__attribute := plugify.ConstructString(attribute)
	plugify.Block {
		Try: func() {
			__native := C.GetClientItemAttribute(__playerSlot, __itemId, (*C.String)(unsafe.Pointer(&__attribute)))
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

// GetClientItemAttribute 
//  @brief Returns an attribute of an item in the client's inventory as text, e.g. "kill eater" after a plugin counted kills.
//
//  @param playerSlot: The index of the player's slot.
//  @param itemId: The item ID.
//  @param attribute: The attribute definition name.
//
//  @return The value, or an empty string if the item has no such attribute.
func GetClientItemAttribute(playerSlot int32, itemId uint64, attribute string) string {
	return _GetClientItemAttribute(playerSlot, itemId, attribute)
}

