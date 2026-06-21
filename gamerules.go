package s2sdk

/*
#include "gamerules.h"
#cgo noescape GetGameRulesProxy
#cgo noescape GetGameRules
#cgo noescape GetGameTeamManager
#cgo noescape GetGameTeamScore
#cgo noescape GetGamePlayerCount
#cgo noescape GetGameTotalRoundsPlayed
#cgo noescape TerminateRound
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

// Generated from s2sdk (group: gamerules)

var P_GetGameRulesProxy = func() uintptr {
	__retVal := uintptr(C.GetGameRulesProxy())
	return __retVal
}

// GetGameRulesProxy 
//  @brief Retrieves the pointer to the current game rules proxy instance.
//
//
//  @return A pointer to the game rules entity instance.
func GetGameRulesProxy() uintptr {
	return P_GetGameRulesProxy()
}

var P_GetGameRules = func() uintptr {
	__retVal := uintptr(C.GetGameRules())
	return __retVal
}

// GetGameRules 
//  @brief Retrieves the pointer to the current game rules instance.
//
//
//  @return A pointer to the game rules object.
func GetGameRules() uintptr {
	return P_GetGameRules()
}

var P_GetGameTeamManager = func(team CSTeam) uintptr {
	var __retVal uintptr
	__team := C.int32_t(team)
	__retVal = uintptr(C.GetGameTeamManager(__team))
	return __retVal
}

// GetGameTeamManager 
//  @brief Retrieves the team manager instance for a specified team.
//
//  @param team: The numeric identifier of the team.
//
//  @return A pointer to the corresponding CTeam instance, or nullptr if the team was not found.
func GetGameTeamManager(team CSTeam) uintptr {
	return P_GetGameTeamManager(team)
}

var P_GetGameTeamScore = func(team CSTeam) int32 {
	var __retVal int32
	__team := C.int32_t(team)
	__retVal = int32(C.GetGameTeamScore(__team))
	return __retVal
}

// GetGameTeamScore 
//  @brief Retrieves the current score of a specified team.
//
//  @param team: The numeric identifier of the team.
//
//  @return The current score of the team, or -1 if the team could not be found.
func GetGameTeamScore(team CSTeam) int32 {
	return P_GetGameTeamScore(team)
}

var P_GetGamePlayerCount = func(team CSTeam) int32 {
	var __retVal int32
	__team := C.int32_t(team)
	__retVal = int32(C.GetGamePlayerCount(__team))
	return __retVal
}

// GetGamePlayerCount 
//  @brief Retrieves the number of players on a specified team.
//
//  @param team: The numeric identifier of the team (e.g., CS_TEAM_T, CS_TEAM_CT, CS_TEAM_SPECTATOR).
//
//  @return The number of players on the team, or -1 if game rules are unavailable.
func GetGamePlayerCount(team CSTeam) int32 {
	return P_GetGamePlayerCount(team)
}

var P_GetGameTotalRoundsPlayed = func() int32 {
	__retVal := int32(C.GetGameTotalRoundsPlayed())
	return __retVal
}

// GetGameTotalRoundsPlayed 
//  @brief Returns the total number of rounds played in the current match.
//
//
//  @return The total number of rounds played, or -1 if the game rules are unavailable.
func GetGameTotalRoundsPlayed() int32 {
	return P_GetGameTotalRoundsPlayed()
}

var P_TerminateRound = func(delay float32, reason CSRoundEndReason) {
	__delay := C.float(delay)
	__reason := C.int32_t(reason)
	C.TerminateRound(__delay, __reason)
}

// TerminateRound 
//  @brief Forces the round to end with a specified reason and delay.
//
//  @param delay: Time (in seconds) to delay before the next round starts.
//  @param reason: The reason for ending the round, defined by the CSRoundEndReason enum.
func TerminateRound(delay float32, reason CSRoundEndReason) {
	P_TerminateRound(delay, reason)
}

