
#include "main.h"
#include "chainparams.h"
#include "base58.h"

#include <boost/algorithm/string/replace.hpp>
#include <boost/algorithm/string/predicate.hpp>
#include <boost/algorithm/string/split.hpp>
#include <boost/algorithm/string.hpp>

using namespace std;
using namespace boost;

string scripttoaddress(const CScript& scriptPubKey,
                       bool* ptrIsNotary = nullptr,
                       string* ptrNotary = nullptr);

bool CTransaction::IsExchangeTx(int & nOut, uint256 & wid) const
{
    nOut = -1;
    wid = uint256();
    size_t n_out = vout.size();
    for(size_t i=0; i< n_out; i++) {
        string sNotary;
        bool fNotary = false;
        scripttoaddress(vout[i].scriptPubKey, &fNotary, &sNotary);
        if (!fNotary) continue;
        if (!boost::starts_with(sNotary, "XCH:")) continue;
        vector<string> vOutputArgs;
        boost::split(vOutputArgs, sNotary, boost::is_any_of(":"));
        if (vOutputArgs.size() <2) {
            return true; // no output, just exchange tx
        }
        string sOut = vOutputArgs[1];
        nOut = std::stoi(sOut);
        if (nOut <0 || nOut >= int(n_out)) {
            nOut = -1;
            return true;
        }
        
        if (vOutputArgs.size() <3) {
            nOut = -1;
            return true; // no control hash
        }
       
        string sControlHash = vOutputArgs[2];
        int64_t nAmount = vout[nOut].nValue;
        string sAddress = scripttoaddress(vout[nOut].scriptPubKey, &fNotary, &sNotary);
        {
            { // liquid withdraw control
                CDataStream ss(SER_GETHASH, 0);
                size_t n_inp = vin.size();
                for(size_t j=0; j< n_inp; j++) {
                    ss << vin[j].prevout.hash;
                    ss << vin[j].prevout.n;
                }
                ss << string("L");
                ss << sAddress;
                ss << nAmount;
                string sHash = Hash(ss.begin(), ss.end()).GetHex();
                if (sHash == sControlHash) {
                    wid = uint256(sHash);
                    return true;
                }
            }
            
            { // reserve withdraw control
                CDataStream ss(SER_GETHASH, 0);
                size_t n_inp = vin.size();
                for(size_t j=0; j< n_inp; j++) {
                    ss << vin[j].prevout.hash;
                    ss << vin[j].prevout.n;
                }
                ss << string("R");
                ss << sAddress;
                ss << nAmount;
                string sHash = Hash(ss.begin(), ss.end()).GetHex();
                if (sHash == sControlHash) {
                    wid = uint256(sHash);
                    return true;
                }
            }
            
            // control id does not match, reset output number
            nOut = -1;
        }
        
        return true;
    }
    return false;
}

string scripttoaddress(const CScript& scriptPubKey,
                       bool* ptrIsNotary,
                       string* ptrNotary) {
    int nRequired;
    txnouttype type;
    vector<CTxDestination> addresses;
    if (ExtractDestinations(scriptPubKey, type, addresses, nRequired)) {
        std::string str_addr_all;
        bool fNone = true;
        for(const CTxDestination& addr : addresses) {
            std::string str_addr = CBitcoinAddress(addr).ToString();
            if (!str_addr_all.empty())
                str_addr_all += "\n";
            str_addr_all += str_addr;
            fNone = false;
        }
        if (!fNone)
            return str_addr_all;
    }

    if (ptrNotary || ptrIsNotary) {
        if (ptrIsNotary) *ptrIsNotary = false;
        if (ptrNotary) *ptrNotary = "";

        opcodetype opcode1;
        vector<unsigned char> vch1;
        CScript::const_iterator pc1 = scriptPubKey.begin();
        if (scriptPubKey.GetOp(pc1, opcode1, vch1)) {
            if (opcode1 == OP_RETURN && scriptPubKey.size()>1) {
                if (ptrIsNotary) *ptrIsNotary = true;
                if (ptrNotary) {
                    unsigned long len_bytes = scriptPubKey[1];
                    if (len_bytes > scriptPubKey.size()-2)
                        len_bytes = scriptPubKey.size()-2;
                    for (uint32_t i=0; i< len_bytes; i++) {
                        ptrNotary->push_back(char(scriptPubKey[i+2]));
                    }
                }
            }
        }
    }

    string as_bytes;
    unsigned long len_bytes = scriptPubKey.size();
    for(unsigned int i=0; i< len_bytes; i++) {
        as_bytes += char(scriptPubKey[i]);
    }
    return as_bytes;
}
