
#ifndef BITBAY_CHAINPARAMS_STUB_H
#define BITBAY_CHAINPARAMS_STUB_H

#include <vector>

/** TESTNET Params **/

struct CChainParams {
    
    enum Base58Type {
        PUBKEY_ADDRESS,
        SCRIPT_ADDRESS,
        SECRET_KEY,
        EXT_PUBLIC_KEY,
        EXT_SECRET_KEY,

        MAX_BASE58_TYPES
    };
    
    std::vector<unsigned char> base58Prefixes[MAX_BASE58_TYPES];
    const std::vector<unsigned char> &Base58Prefix(Base58Type type) const { return base58Prefixes[type]; }

    int nPegFrozenTime  = (3600 * 24 * 30);
    int nPegVFrozenTime = (3600 * 24 * 30 *4);
    int nPegInterval = 200;
    int nPegIntervalProbeHeight = 100000;
    
    int PegFrozenTime() const { return nPegFrozenTime; }
    int PegVFrozenTime() const { return nPegVFrozenTime; }
    
    virtual int PegInterval(int /*nHeight*/) const { return nPegInterval; }
    
    CChainParams();
};

CChainParams & Params();

#endif
