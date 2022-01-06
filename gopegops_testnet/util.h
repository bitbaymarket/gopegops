
#ifndef GOPEGOPS_UTIL_H
#define GOPEGOPS_UTIL_H

#include "tinyformat.h"

#include <boost/thread/mutex.hpp>
#include <boost/thread.hpp>

#define LogPrintf(fmt, ...) printf(fmt, ##__VA_ARGS__)
#define GetArg(a1,a2) a2

static const unsigned int LOCKTIME_THRESHOLD = 500000000; // Tue Nov  5 00:53:20 1985 UTC

uint256 GetRandHash();

int64_t GetAdjustedTime();

int LogPrintStr(const std::string &str);

static inline bool error(const char* format)
{
    LogPrintStr(std::string("ERROR: ") + format + "\n");
    return false;
}

/* When we switch to C++11, this can be switched to variadic templates instead
 * of this macro-based construction (see tinyformat.h).
 */
#define MAKE_ERROR_AND_LOG_FUNC(n)                                        \
    /*   Print to debug.log if -debug=category switch is given OR category is NULL. */ \
    template<TINYFORMAT_ARGTYPES(n)>                                          \
    static inline int LogPrint(const char* category, const char* format, TINYFORMAT_VARARGS(n))  \
    {                                                                         \
        return LogPrintStr(tfm::format(format, TINYFORMAT_PASSARGS(n))); \
    }                                                                         \
    /*   Log error and return false */                                        \
    template<TINYFORMAT_ARGTYPES(n)>                                          \
    static inline bool error(const char* format, TINYFORMAT_VARARGS(n))                     \
    {                                                                         \
        LogPrintStr("ERROR: " + tfm::format(format, TINYFORMAT_PASSARGS(n)) + "\n"); \
        return false;                                                         \
    }

TINYFORMAT_FOREACH_ARGNUM(MAKE_ERROR_AND_LOG_FUNC)


#endif

