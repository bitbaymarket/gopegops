
#include "chainparams.h"

#include <boost/assign/list_of.hpp>

using namespace boost::assign;

struct CChainParams testNetParams;

CChainParams::CChainParams() {
    // testnet
    base58Prefixes[PUBKEY_ADDRESS] = std::vector<unsigned char>(1,111);
    base58Prefixes[SCRIPT_ADDRESS] = std::vector<unsigned char>(1,196);
    base58Prefixes[SECRET_KEY] =     std::vector<unsigned char>(1,239);
    base58Prefixes[EXT_PUBLIC_KEY] = boost::assign::list_of(0x04)(0x35)(0x87)(0xCF).convert_to_container<std::vector<unsigned char> >();
    base58Prefixes[EXT_SECRET_KEY] = boost::assign::list_of(0x04)(0x35)(0x83)(0x94).convert_to_container<std::vector<unsigned char> >();
}

CChainParams & Params() {
    return testNetParams;
}


