
#include <map>
#include <vector>
#include <stdarg.h>
#include <string.h>

#include "uint256.h"

#include <openssl/rand.h>
using namespace std;

#include <allocators.h>
LockedPageManager LockedPageManager::instance;


uint256 GetRandHash()
{
    uint256 hash;
    RAND_bytes((unsigned char*)&hash, sizeof(hash));
    return hash;
}

int LogPrintStr(const std::string &str)
{
    int ret = 0; // Returns total number of characters written
    ret = fwrite(str.data(), 1, str.size(), stdout);
    return ret;
}

static int64_t nMockTime = 0;  // For unit testing

int64_t GetAdjustedTime()
{
    if (nMockTime) return nMockTime;

    return time(NULL);
}
