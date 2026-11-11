#pragma once

#include "shared.h"

extern bool (*__s2sdk_HookClientVoiceData)(void*);

static bool HookClientVoiceData(void* callback) {
	return __s2sdk_HookClientVoiceData(callback);
}

extern bool (*__s2sdk_UnhookClientVoiceData)(void*);

static bool UnhookClientVoiceData(void* callback) {
	return __s2sdk_UnhookClientVoiceData(callback);
}

extern int32_t (*__s2sdk_VoiceDataGetFormat)(uintptr_t);

static int32_t VoiceDataGetFormat(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetFormat(voiceData);
}

extern void (*__s2sdk_VoiceDataSetFormat)(uintptr_t, int32_t);

static void VoiceDataSetFormat(uintptr_t voiceData, int32_t format) {
	__s2sdk_VoiceDataSetFormat(voiceData, format);
}

extern Vector (*__s2sdk_VoiceDataGetAudio)(uintptr_t);

static Vector VoiceDataGetAudio(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetAudio(voiceData);
}

extern void (*__s2sdk_VoiceDataSetAudio)(uintptr_t, Vector*);

static void VoiceDataSetAudio(uintptr_t voiceData, Vector* audio) {
	__s2sdk_VoiceDataSetAudio(voiceData, audio);
}

extern uint32_t (*__s2sdk_VoiceDataGetSampleRate)(uintptr_t);

static uint32_t VoiceDataGetSampleRate(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetSampleRate(voiceData);
}

extern void (*__s2sdk_VoiceDataSetSampleRate)(uintptr_t, uint32_t);

static void VoiceDataSetSampleRate(uintptr_t voiceData, uint32_t sampleRate) {
	__s2sdk_VoiceDataSetSampleRate(voiceData, sampleRate);
}

extern uint32_t (*__s2sdk_VoiceDataGetNumPackets)(uintptr_t);

static uint32_t VoiceDataGetNumPackets(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetNumPackets(voiceData);
}

extern void (*__s2sdk_VoiceDataSetNumPackets)(uintptr_t, uint32_t);

static void VoiceDataSetNumPackets(uintptr_t voiceData, uint32_t numPackets) {
	__s2sdk_VoiceDataSetNumPackets(voiceData, numPackets);
}

extern Vector (*__s2sdk_VoiceDataGetPacketOffsets)(uintptr_t);

static Vector VoiceDataGetPacketOffsets(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetPacketOffsets(voiceData);
}

extern void (*__s2sdk_VoiceDataSetPacketOffsets)(uintptr_t, Vector*);

static void VoiceDataSetPacketOffsets(uintptr_t voiceData, Vector* offsets) {
	__s2sdk_VoiceDataSetPacketOffsets(voiceData, offsets);
}

extern float (*__s2sdk_VoiceDataGetVoiceLevel)(uintptr_t);

static float VoiceDataGetVoiceLevel(uintptr_t voiceData) {
	return __s2sdk_VoiceDataGetVoiceLevel(voiceData);
}

extern void (*__s2sdk_VoiceDataSetVoiceLevel)(uintptr_t, float);

static void VoiceDataSetVoiceLevel(uintptr_t voiceData, float voiceLevel) {
	__s2sdk_VoiceDataSetVoiceLevel(voiceData, voiceLevel);
}

extern bool (*__s2sdk_SendVoiceData)(Vector*, int32_t, Vector*, Vector*, uint32_t);

static bool SendVoiceData(Vector* recipients, int32_t speakerSlot, Vector* audio, Vector* packetOffsets, uint32_t sectionNumber) {
	return __s2sdk_SendVoiceData(recipients, speakerSlot, audio, packetOffsets, sectionNumber);
}

