#pragma once

#include "shared.h"

extern String (*__s2sdk_ReadFileVPK)(String*, String*);

static String ReadFileVPK(String* localFileName, String* pathId) {
	return __s2sdk_ReadFileVPK(localFileName, pathId);
}

extern Vector (*__s2sdk_FindFileAbsoluteList)(String*, String*);

static Vector FindFileAbsoluteList(String* wildcard, String* pathId) {
	return __s2sdk_FindFileAbsoluteList(wildcard, pathId);
}

