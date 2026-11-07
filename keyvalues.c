#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1Create)(String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1Destroy)(uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_Kv1GetName)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetName)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1FindKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1FindOrCreateKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1CreateKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1CreateNewKey)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1AddSubKey)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetFirstSubKey)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetNextKey)(uintptr_t) = NULL;


PLUGIFY_EXPORT Vector4 (*__s2sdk_Kv1GetColor)(uintptr_t, String*, Vector4*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetColor)(uintptr_t, String*, Vector4*) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_Kv1GetInt)(uintptr_t, String*, int32_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetInt)(uintptr_t, String*, int32_t) = NULL;


PLUGIFY_EXPORT float (*__s2sdk_Kv1GetFloat)(uintptr_t, String*, float) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetFloat)(uintptr_t, String*, float) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_Kv1GetString)(uintptr_t, String*, String*) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetString)(uintptr_t, String*, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetPtr)(uintptr_t, String*, uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetPtr)(uintptr_t, String*, uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1GetBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetBool)(uintptr_t, String*, bool) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1MakeCopy)(uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1Clear)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1LoadFromFile)(uintptr_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1LoadFromString)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1SaveToFile)(uintptr_t, String*, String*) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1RemoveKey)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uint64_t (*__s2sdk_Kv1GetUint64)(uintptr_t, String*, uint64_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetUint64)(uintptr_t, String*, uint64_t) = NULL;


PLUGIFY_EXPORT int32_t (*__s2sdk_Kv1GetDataType)(uintptr_t, String*) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetFirstValue)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetNextValue)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetFirstTrueSubKey)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__s2sdk_Kv1GetNextTrueSubKey)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1RemoveSubKey)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1SetEscapeSequences)(uintptr_t, bool) = NULL;


PLUGIFY_EXPORT void (*__s2sdk_Kv1CopySubkeys)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__s2sdk_Kv1ExportToString)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__s2sdk_Kv1IsEmpty)(uintptr_t, String*) = NULL;


