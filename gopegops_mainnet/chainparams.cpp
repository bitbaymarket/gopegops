
#include "chainparams.h"

#include <boost/assign/list_of.hpp>

using namespace boost::assign;

struct CChainParams mainNetParams;

CChainParams::CChainParams() {
    // mainnet
    base58Prefixes[PUBKEY_ADDRESS] = std::vector<unsigned char>(1,25);
    base58Prefixes[SCRIPT_ADDRESS] = std::vector<unsigned char>(1,85);
    base58Prefixes[SECRET_KEY] =     std::vector<unsigned char>(1,153);
    base58Prefixes[EXT_PUBLIC_KEY] = boost::assign::list_of(0x04)(0x88)(0xB2)(0x1E).convert_to_container<std::vector<unsigned char> >();
    base58Prefixes[EXT_SECRET_KEY] = boost::assign::list_of(0x04)(0x88)(0xAD)(0xE4).convert_to_container<std::vector<unsigned char> >();
}

CChainParams & Params() {
    return mainNetParams;
}


