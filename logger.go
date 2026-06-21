package s2sdk

/*
#include "logger.h"
#cgo noescape RegisterLoggingChannel
#cgo noescape AddLoggerTagToChannel
#cgo noescape HasLoggerTag
#cgo noescape IsLoggerChannelEnabledBySeverity
#cgo noescape IsLoggerChannelEnabledByVerbosity
#cgo noescape GetLoggerChannelVerbosity
#cgo noescape SetLoggerChannelVerbosity
#cgo noescape SetLoggerChannelVerbosityByName
#cgo noescape SetLoggerChannelVerbosityByTag
#cgo noescape GetLoggerChannelColor
#cgo noescape SetLoggerChannelColor
#cgo noescape GetLoggerChannelFlags
#cgo noescape SetLoggerChannelFlags
#cgo noescape Log
#cgo noescape LogColored
#cgo noescape LogFull
#cgo noescape LogFullColored
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

// Generated from s2sdk (group: logger)

var P_RegisterLoggingChannel = func(name string, flags int32, verbosity LoggingVerbosity, color plugify.Vector4) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	__flags := C.int32_t(flags)
	__verbosity := C.int32_t(verbosity)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	plugify.Block {
		Try: func() {
			__retVal = int32(C.RegisterLoggingChannel((*C.String)(unsafe.Pointer(&__name)), __flags, __verbosity, &__color))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// RegisterLoggingChannel 
//  @brief Registers a new logging channel with specified properties.
//
//  @param name: The name of the logging channel.
//  @param flags: Flags associated with the logging channel.
//  @param verbosity: The verbosity level for the logging channel.
//  @param color: The color for messages logged to this channel.
//
//  @return The ID of the newly created logging channel.
func RegisterLoggingChannel(name string, flags int32, verbosity LoggingVerbosity, color plugify.Vector4) int32 {
	return P_RegisterLoggingChannel(name, flags, verbosity, color)
}

var P_AddLoggerTagToChannel = func(channelID int32, tagName string) {
	__channelID := C.int32_t(channelID)
	__tagName := plugify.ConstructString(tagName)
	plugify.Block {
		Try: func() {
			C.AddLoggerTagToChannel(__channelID, (*C.String)(unsafe.Pointer(&__tagName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tagName)
		},
	}.Do()
}

// AddLoggerTagToChannel 
//  @brief Adds a tag to a specified logging channel.
//
//  @param channelID: The ID of the logging channel to which the tag will be added.
//  @param tagName: The name of the tag to add to the channel.
func AddLoggerTagToChannel(channelID int32, tagName string) {
	P_AddLoggerTagToChannel(channelID, tagName)
}

var P_HasLoggerTag = func(channelID int32, tag string) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__tag := plugify.ConstructString(tag)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.HasLoggerTag(__channelID, (*C.String)(unsafe.Pointer(&__tag))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tag)
		},
	}.Do()
	return __retVal
}

// HasLoggerTag 
//  @brief Checks if a specified tag exists in a logging channel.
//
//  @param channelID: The ID of the logging channel.
//  @param tag: The name of the tag to check for.
//
//  @return True if the tag exists in the channel, otherwise false.
func HasLoggerTag(channelID int32, tag string) bool {
	return P_HasLoggerTag(channelID, tag)
}

var P_IsLoggerChannelEnabledBySeverity = func(channelID int32, severity LoggingSeverity) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__retVal = bool(C.IsLoggerChannelEnabledBySeverity(__channelID, __severity))
	return __retVal
}

// IsLoggerChannelEnabledBySeverity 
//  @brief Checks if a logging channel is enabled based on severity.
//
//  @param channelID: The ID of the logging channel.
//  @param severity: The severity of a logging operation.
//
//  @return True if the channel is enabled for the specified severity, otherwise false.
func IsLoggerChannelEnabledBySeverity(channelID int32, severity LoggingSeverity) bool {
	return P_IsLoggerChannelEnabledBySeverity(channelID, severity)
}

var P_IsLoggerChannelEnabledByVerbosity = func(channelID int32, verbosity LoggingVerbosity) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__verbosity := C.int32_t(verbosity)
	__retVal = bool(C.IsLoggerChannelEnabledByVerbosity(__channelID, __verbosity))
	return __retVal
}

// IsLoggerChannelEnabledByVerbosity 
//  @brief Checks if a logging channel is enabled based on verbosity.
//
//  @param channelID: The ID of the logging channel.
//  @param verbosity: The verbosity level to check.
//
//  @return True if the channel is enabled for the specified verbosity, otherwise false.
func IsLoggerChannelEnabledByVerbosity(channelID int32, verbosity LoggingVerbosity) bool {
	return P_IsLoggerChannelEnabledByVerbosity(channelID, verbosity)
}

var P_GetLoggerChannelVerbosity = func(channelID int32) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__retVal = int32(C.GetLoggerChannelVerbosity(__channelID))
	return __retVal
}

// GetLoggerChannelVerbosity 
//  @brief Retrieves the verbosity level of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//
//  @return The verbosity level of the specified logging channel.
func GetLoggerChannelVerbosity(channelID int32) int32 {
	return P_GetLoggerChannelVerbosity(channelID)
}

var P_SetLoggerChannelVerbosity = func(channelID int32, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__verbosity := C.int32_t(verbosity)
	C.SetLoggerChannelVerbosity(__channelID, __verbosity)
}

// SetLoggerChannelVerbosity 
//  @brief Sets the verbosity level of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//  @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosity(channelID int32, verbosity LoggingVerbosity) {
	P_SetLoggerChannelVerbosity(channelID, verbosity)
}

var P_SetLoggerChannelVerbosityByName = func(channelID int32, name string, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__name := plugify.ConstructString(name)
	__verbosity := C.int32_t(verbosity)
	plugify.Block {
		Try: func() {
			C.SetLoggerChannelVerbosityByName(__channelID, (*C.String)(unsafe.Pointer(&__name)), __verbosity)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// SetLoggerChannelVerbosityByName 
//  @brief Sets the verbosity level of a logging channel by name.
//
//  @param channelID: The ID of the logging channel.
//  @param name: The name of the logging channel.
//  @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosityByName(channelID int32, name string, verbosity LoggingVerbosity) {
	P_SetLoggerChannelVerbosityByName(channelID, name, verbosity)
}

var P_SetLoggerChannelVerbosityByTag = func(channelID int32, tag string, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__tag := plugify.ConstructString(tag)
	__verbosity := C.int32_t(verbosity)
	plugify.Block {
		Try: func() {
			C.SetLoggerChannelVerbosityByTag(__channelID, (*C.String)(unsafe.Pointer(&__tag)), __verbosity)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tag)
		},
	}.Do()
}

// SetLoggerChannelVerbosityByTag 
//  @brief Sets the verbosity level of a logging channel by tag.
//
//  @param channelID: The ID of the logging channel.
//  @param tag: The name of the tag.
//  @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosityByTag(channelID int32, tag string, verbosity LoggingVerbosity) {
	P_SetLoggerChannelVerbosityByTag(channelID, tag, verbosity)
}

var P_GetLoggerChannelColor = func(channelID int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__channelID := C.int32_t(channelID)
	__native := C.GetLoggerChannelColor(__channelID)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetLoggerChannelColor 
//  @brief Retrieves the color setting of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//
//  @return The color value of the specified logging channel.
func GetLoggerChannelColor(channelID int32) plugify.Vector4 {
	return P_GetLoggerChannelColor(channelID)
}

var P_SetLoggerChannelColor = func(channelID int32, color plugify.Vector4) {
	__channelID := C.int32_t(channelID)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	C.SetLoggerChannelColor(__channelID, &__color)
}

// SetLoggerChannelColor 
//  @brief Sets the color setting of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//  @param color: The new color value to set for the channel.
func SetLoggerChannelColor(channelID int32, color plugify.Vector4) {
	P_SetLoggerChannelColor(channelID, color)
}

var P_GetLoggerChannelFlags = func(channelID int32) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__retVal = int32(C.GetLoggerChannelFlags(__channelID))
	return __retVal
}

// GetLoggerChannelFlags 
//  @brief Retrieves the flags of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//
//  @return The flags of the specified logging channel.
func GetLoggerChannelFlags(channelID int32) int32 {
	return P_GetLoggerChannelFlags(channelID)
}

var P_SetLoggerChannelFlags = func(channelID int32, eFlags LoggingChannelFlags) {
	__channelID := C.int32_t(channelID)
	__eFlags := C.int32_t(eFlags)
	C.SetLoggerChannelFlags(__channelID, __eFlags)
}

// SetLoggerChannelFlags 
//  @brief Sets the flags of a logging channel.
//
//  @param channelID: The ID of the logging channel.
//  @param eFlags: The new flags to set for the channel.
func SetLoggerChannelFlags(channelID int32, eFlags LoggingChannelFlags) {
	P_SetLoggerChannelFlags(channelID, eFlags)
}

var P_Log = func(channelID int32, severity LoggingSeverity, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.Log(__channelID, __severity, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// Log 
//  @brief Logs a message to a specified channel with a severity level.
//
//  @param channelID: The ID of the logging channel.
//  @param severity: The severity level for the log message.
//  @param message: The message to log.
//
//  @return An integer indicating the result of the logging operation.
func Log(channelID int32, severity LoggingSeverity, message string) int32 {
	return P_Log(channelID, severity, message)
}

var P_LogColored = func(channelID int32, severity LoggingSeverity, color plugify.Vector4, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.LogColored(__channelID, __severity, &__color, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogColored 
//  @brief Logs a colored message to a specified channel with a severity level.
//
//  @param channelID: The ID of the logging channel.
//  @param severity: The severity level for the log message.
//  @param color: The color for the log message.
//  @param message: The message to log.
//
//  @return An integer indicating the result of the logging operation.
func LogColored(channelID int32, severity LoggingSeverity, color plugify.Vector4, message string) int32 {
	return P_LogColored(channelID, severity, color, message)
}

var P_LogFull = func(channelID int32, severity LoggingSeverity, file string, line int32, function string, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__file := plugify.ConstructString(file)
	__line := C.int32_t(line)
	__function := plugify.ConstructString(function)
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.LogFull(__channelID, __severity, (*C.String)(unsafe.Pointer(&__file)), __line, (*C.String)(unsafe.Pointer(&__function)), (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__file)
			plugify.DestroyString(&__function)
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogFull 
//  @brief Logs a detailed message to a specified channel, including source code info.
//
//  @param channelID: The ID of the logging channel.
//  @param severity: The severity level for the log message.
//  @param file: The file name where the log call occurred.
//  @param line: The line number where the log call occurred.
//  @param function: The name of the function where the log call occurred.
//  @param message: The message to log.
//
//  @return An integer indicating the result of the logging operation.
func LogFull(channelID int32, severity LoggingSeverity, file string, line int32, function string, message string) int32 {
	return P_LogFull(channelID, severity, file, line, function, message)
}

var P_LogFullColored = func(channelID int32, severity LoggingSeverity, file string, line int32, function string, color plugify.Vector4, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__file := plugify.ConstructString(file)
	__line := C.int32_t(line)
	__function := plugify.ConstructString(function)
	__color := *(*C.Vector4)(unsafe.Pointer(&color))
	__message := plugify.ConstructString(message)
	plugify.Block {
		Try: func() {
			__retVal = int32(C.LogFullColored(__channelID, __severity, (*C.String)(unsafe.Pointer(&__file)), __line, (*C.String)(unsafe.Pointer(&__function)), &__color, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__file)
			plugify.DestroyString(&__function)
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogFullColored 
//  @brief Logs a detailed colored message to a specified channel, including source code info.
//
//  @param channelID: The ID of the logging channel.
//  @param severity: The severity level for the log message.
//  @param file: The file name where the log call occurred.
//  @param line: The line number where the log call occurred.
//  @param function: The name of the function where the log call occurred.
//  @param color: The color for the log message.
//  @param message: The message to log.
//
//  @return An integer indicating the result of the logging operation.
func LogFullColored(channelID int32, severity LoggingSeverity, file string, line int32, function string, color plugify.Vector4, message string) int32 {
	return P_LogFullColored(channelID, severity, file, line, function, color, message)
}

