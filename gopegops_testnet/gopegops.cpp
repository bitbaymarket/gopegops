
#include <string>
#include <cstring>
#include <iostream>

#include <gopegops.h>

#include "pegops.h"

struct Out_getpeglevel* getpeglevel(
    int   inp_cycle_now,
    int   inp_cycle_prev,
    int   inp_buffer,
    int   inp_peg_now,
    int   inp_peg_next,
    int   inp_peg_next_next,
    char* inp_exchange_pegdata64,
    char* inp_pegshift_pegdata64
    )
{
    struct Out_getpeglevel* out = new Out_getpeglevel;

    std::string out_peglevel_hex;
    std::string out_pegpool_pegdata64;
    std::string out_err;

    bool ok = pegops::getpeglevel(
        inp_cycle_now,
        inp_cycle_prev,
        inp_buffer,
        inp_peg_now,
        inp_peg_next,
        inp_peg_next_next,
        std::string(inp_exchange_pegdata64),
        std::string(inp_pegshift_pegdata64),

        out_peglevel_hex,
        out->out_exchange_liquid,
        out->out_exchange_reserve,
        out_pegpool_pegdata64,
        out->out_pegpool_value,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_peglevel_hex = new char[out_peglevel_hex.length()+1];
    out->out_pegpool_pegdata64 = new char[out_pegpool_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_peglevel_hex, out_peglevel_hex.c_str());
    std::strcpy(out->out_pegpool_pegdata64, out_pegpool_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_getpeglevel_del(struct Out_getpeglevel* out)
{
    if (!out) return;
    delete[] out->out_peglevel_hex;
    delete[] out->out_pegpool_pegdata64;
    delete[] out->out_err;
    delete out;
}

struct Out_updatepegbalances* updatepegbalances(
    char* inp_balance_pegdata64,
    char* inp_pegpool_pegdata64,
    char* inp_peglevel_hex
    )
{
    struct Out_updatepegbalances* out = new Out_updatepegbalances;
    out->out_balance_liquid  =0;
    out->out_balance_reserve =0;
    out->out_pegpool_value =0;

    std::string out_balance_pegdata64;
    std::string out_pegpool_pegdata64;
    std::string out_err;

    bool ok = pegops::updatepegbalances(
        std::string(inp_balance_pegdata64),
        std::string(inp_pegpool_pegdata64),
        std::string(inp_peglevel_hex),

        out_balance_pegdata64,
        out->out_balance_liquid,
        out->out_balance_reserve,
        out_pegpool_pegdata64,
        out->out_pegpool_value,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_balance_pegdata64 = new char[out_balance_pegdata64.length()+1];
    out->out_pegpool_pegdata64 = new char[out_pegpool_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_balance_pegdata64, out_balance_pegdata64.c_str());
    std::strcpy(out->out_pegpool_pegdata64, out_pegpool_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_updatepegbalances_del(struct Out_updatepegbalances* out)
{
    if (!out) return;
    delete[] out->out_balance_pegdata64;
    delete[] out->out_pegpool_pegdata64;
    delete[] out->out_err;
    delete out;
}

struct Out_movecoins* movecoins(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex,
    int	 inp_cross_cycles
    )
{
    struct Out_movecoins* out = new Out_movecoins;
    out->out_src_liquid  =0;
    out->out_src_reserve =0;
    out->out_dst_liquid  =0;
    out->out_dst_reserve =0;

    std::string out_src_pegdata64;
    std::string out_dst_pegdata64;
    std::string out_err;

    bool ok = pegops::movecoins(
        inp_move_amount,
        std::string(inp_src_pegdata64),
        std::string(inp_dst_pegdata64),
        std::string(inp_peglevel_hex),
        inp_cross_cycles >0 ? true : false,

        out_src_pegdata64,
        out->out_src_liquid,
        out->out_src_reserve,
        out_dst_pegdata64,
        out->out_dst_liquid,
        out->out_dst_reserve,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_src_pegdata64 = new char[out_src_pegdata64.length()+1];
    out->out_dst_pegdata64 = new char[out_dst_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_src_pegdata64, out_src_pegdata64.c_str());
    std::strcpy(out->out_dst_pegdata64, out_dst_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

struct Out_movecoins* moveliquid(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex
    )
{
    struct Out_movecoins* out = new Out_movecoins;
    out->out_src_liquid  =0;
    out->out_src_reserve =0;
    out->out_dst_liquid  =0;
    out->out_dst_reserve =0;

    std::string out_src_pegdata64;
    std::string out_dst_pegdata64;
    std::string out_err;

    bool ok = pegops::moveliquid(
        inp_move_amount,
        std::string(inp_src_pegdata64),
        std::string(inp_dst_pegdata64),
        std::string(inp_peglevel_hex),

        out_src_pegdata64,
        out->out_src_liquid,
        out->out_src_reserve,
        out_dst_pegdata64,
        out->out_dst_liquid,
        out->out_dst_reserve,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_src_pegdata64 = new char[out_src_pegdata64.length()+1];
    out->out_dst_pegdata64 = new char[out_dst_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_src_pegdata64, out_src_pegdata64.c_str());
    std::strcpy(out->out_dst_pegdata64, out_dst_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

struct Out_movecoins* movereserve(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex
    )
{
    struct Out_movecoins* out = new Out_movecoins;
    out->out_src_liquid  =0;
    out->out_src_reserve =0;
    out->out_dst_liquid  =0;
    out->out_dst_reserve =0;

    std::string out_src_pegdata64;
    std::string out_dst_pegdata64;
    std::string out_err;

    bool ok = pegops::movereserve(
        inp_move_amount,
        std::string(inp_src_pegdata64),
        std::string(inp_dst_pegdata64),
        std::string(inp_peglevel_hex),

        out_src_pegdata64,
        out->out_src_liquid,
        out->out_src_reserve,
        out_dst_pegdata64,
        out->out_dst_liquid,
        out->out_dst_reserve,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_src_pegdata64 = new char[out_src_pegdata64.length()+1];
    out->out_dst_pegdata64 = new char[out_dst_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_src_pegdata64, out_src_pegdata64.c_str());
    std::strcpy(out->out_dst_pegdata64, out_dst_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_movecoins_del(struct Out_movecoins* out)
{
    if (!out) return;
    delete[] out->out_src_pegdata64;
    delete[] out->out_dst_pegdata64;
    delete[] out->out_err;
    delete out;
}

struct Out_removecoins* removecoins(
    char*   inp_from_pegdata64,
    char*   inp_remove_pegdata64
    )
{
    struct Out_removecoins* out = new Out_removecoins;
    out->out_left_liquid  =0;
    out->out_left_reserve =0;

    std::string out_left_pegdata64;
    std::string out_err;

    bool ok = pegops::removecoins(
        std::string(inp_from_pegdata64),
        std::string(inp_remove_pegdata64),

        out_left_pegdata64,
        out->out_left_liquid,
        out->out_left_reserve,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_left_pegdata64 = new char[out_left_pegdata64.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_left_pegdata64, out_left_pegdata64.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_removecoins_del(struct Out_removecoins* out)
{
    if (!out) return;
    delete[] out->out_left_pegdata64;
    delete[] out->out_err;
    delete out;
}

struct Out_txoutupdate* txoutupdate(
    char*   inp_txout_pegdata64,
    char*   inp_peglevel_hex
    )
{
    struct Out_txoutupdate* out = new Out_txoutupdate;
    out->out_txout_value =0;
    out->out_txout_next_cycle_available_liquid =0;
    out->out_txout_next_cycle_available_reserve =0;
    out->out_txout_value_hli =0;
    out->out_txout_next_cycle_available_liquid_hli =0;
    out->out_txout_next_cycle_available_reserve_hli =0;

    std::string out_err;

    bool ok = pegops::updatetxout(
        std::string(inp_txout_pegdata64),
        std::string(inp_peglevel_hex),

        out->out_txout_value,
        out->out_txout_next_cycle_available_liquid,
        out->out_txout_next_cycle_available_reserve,
        out->out_txout_value_hli,
        out->out_txout_next_cycle_available_liquid_hli,
        out->out_txout_next_cycle_available_reserve_hli,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_txoutupdate_del(struct Out_txoutupdate* out)
{
    if (!out) return;
    delete[] out->out_err;
    delete out;
}

struct Inp_txinp* maketxinps(int n)
{
    struct Inp_txinp* inps = new Inp_txinp[n];
    return inps;
}

void filltxinp(
    struct Inp_txinp* inps,
    int     i,
    char*   inp_txout,
    char*   inp_pegdata64,
    char*   inp_privkey_bip32)
{
    struct Inp_txinp* inp = inps+i;

    inp->inp_txout = new char[strlen(inp_txout)+1];
    inp->inp_pegdata64 = new char[strlen(inp_pegdata64)+1];
    inp->inp_privkey_bip32 = new char[strlen(inp_privkey_bip32)+1];

    std::strcpy(inp->inp_txout, inp_txout);
    std::strcpy(inp->inp_pegdata64, inp_pegdata64);
    std::strcpy(inp->inp_privkey_bip32, inp_privkey_bip32);
}

struct Out_preparewithdraw* prepareliquidwithdraw(
    struct Inp_txinp*   inp_txinps,
    int                 inp_txinps_len,
    char*               inp_balance_pegdata64,
    char*               inp_exchange_pegdata64,
    char*               inp_pegshift_pegdata64,
    int64_t             inp_amount_with_fee,
    char*               inp_address,
    char*               inp_peglevel_hex
    )
{
    struct Out_preparewithdraw* out = new Out_preparewithdraw;
    out->out_balance_liquid  =0;
    out->out_balance_reserve =0;
    out->out_exchange_liquid  =0;
    out->out_exchange_reserve =0;

    std::string out_balance_pegdata64;
    std::string out_exchange_pegdata64;
    std::string out_pegshift_pegdata64;
    std::string out_processed_pegdata64;
    std::string out_requested_pegdata64;
    std::string out_withdraw_idxch;
    std::string out_withdraw_txout;
    std::string out_rawtx;
    std::string out_err;

    pegops::txinps inp_txinps_str;
    pegops::txouts out_txouts_str;

    for(int i=0; i<inp_txinps_len; i++) {
        struct Inp_txinp* inp_txinp = inp_txinps+i;
        auto inp_txinp_str = make_tuple(
            std::string(inp_txinp->inp_txout),
            std::string(inp_txinp->inp_pegdata64),
            std::string(inp_txinp->inp_privkey_bip32)
        );
        inp_txinps_str.push_back(inp_txinp_str);
    }

    bool ok = pegops::prepareliquidwithdraw(
        inp_txinps_str,
        std::string(inp_balance_pegdata64),
        std::string(inp_exchange_pegdata64),
        std::string(inp_pegshift_pegdata64),
        inp_amount_with_fee,
        std::string(inp_address),
        std::string(inp_peglevel_hex),

        out_balance_pegdata64,
        out->out_balance_liquid,
        out->out_balance_reserve,
        out_exchange_pegdata64,
        out->out_exchange_liquid,
        out->out_exchange_reserve,
        out_pegshift_pegdata64,
        out_requested_pegdata64,
        out_processed_pegdata64,
        out_withdraw_idxch,
        out_withdraw_txout,
        out_rawtx,
        out_txouts_str,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_balance_pegdata64 = new char[out_balance_pegdata64.length()+1];
    out->out_exchange_pegdata64 = new char[out_exchange_pegdata64.length()+1];
    out->out_pegshift_pegdata64 = new char[out_pegshift_pegdata64.length()+1];
    out->out_requested_pegdata64 = new char[out_requested_pegdata64.length()+1];
    out->out_processed_pegdata64 = new char[out_processed_pegdata64.length()+1];
    out->out_withdraw_idxch = new char[out_withdraw_idxch.length()+1];
    out->out_withdraw_txout = new char[out_withdraw_txout.length()+1];
    out->out_rawtx = new char[out_rawtx.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_balance_pegdata64, out_balance_pegdata64.c_str());
    std::strcpy(out->out_exchange_pegdata64, out_exchange_pegdata64.c_str());
    std::strcpy(out->out_pegshift_pegdata64, out_pegshift_pegdata64.c_str());
    std::strcpy(out->out_requested_pegdata64, out_requested_pegdata64.c_str());
    std::strcpy(out->out_processed_pegdata64, out_processed_pegdata64.c_str());
    std::strcpy(out->out_withdraw_idxch, out_withdraw_idxch.c_str());
    std::strcpy(out->out_withdraw_txout, out_withdraw_txout.c_str());
    std::strcpy(out->out_rawtx, out_rawtx.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    out->out_txouts_len = out_txouts_str.size();
    if (out_txouts_str.empty()) {
        out->out_txouts = nullptr;
    } else {
        out->out_txouts = new Out_txout[out->out_txouts_len];
        struct Out_txout * out_txout_ptr = out->out_txouts;
        for(const auto & out_txout_str : out_txouts_str) {
            std::string out_txout;
            std::string out_pegdata64;
            std::string out_privkey_bip32;
            std::tie(out_txout, out_pegdata64, out_privkey_bip32) = out_txout_str;
            out_txout_ptr->out_txout = new char[out_txout.length()+1];
            out_txout_ptr->out_pegdata64 = new char[out_pegdata64.length()+1];
            out_txout_ptr->out_privkey_bip32 = new char[out_privkey_bip32.length()+1];
            std::strcpy(out_txout_ptr->out_txout, out_txout.c_str());
            std::strcpy(out_txout_ptr->out_pegdata64, out_pegdata64.c_str());
            std::strcpy(out_txout_ptr->out_privkey_bip32, out_privkey_bip32.c_str());
            out_txout_ptr++;
        }
    }

    return out;
}

struct Out_preparewithdraw* preparereservewithdraw(
    struct Inp_txinp*   inp_txinps,
    int                 inp_txinps_len,
    char*               inp_balance_pegdata64,
    char*               inp_exchange_pegdata64,
    char*               inp_pegshift_pegdata64,
    int64_t             inp_amount_with_fee,
    char*               inp_address,
    char*               inp_peglevel_hex
    )
{
    struct Out_preparewithdraw* out = new Out_preparewithdraw;
    out->out_balance_liquid  =0;
    out->out_balance_reserve =0;
    out->out_exchange_liquid  =0;
    out->out_exchange_reserve =0;

    std::string out_balance_pegdata64;
    std::string out_exchange_pegdata64;
    std::string out_pegshift_pegdata64;
    std::string out_processed_pegdata64;
    std::string out_requested_pegdata64;
    std::string out_withdraw_idxch;
    std::string out_withdraw_txout;
    std::string out_rawtx;
    std::string out_err;

    pegops::txinps inp_txinps_str;
    pegops::txouts out_txouts_str;

    for(int i=0; i<inp_txinps_len; i++) {
        struct Inp_txinp* inp_txinp = inp_txinps+i;
        auto inp_txinp_str = make_tuple(
            std::string(inp_txinp->inp_txout),
            std::string(inp_txinp->inp_pegdata64),
            std::string(inp_txinp->inp_privkey_bip32)
        );
        inp_txinps_str.push_back(inp_txinp_str);
    }

    bool ok = pegops::preparereservewithdraw(
        inp_txinps_str,
        std::string(inp_balance_pegdata64),
        std::string(inp_exchange_pegdata64),
        std::string(inp_pegshift_pegdata64),
        inp_amount_with_fee,
        std::string(inp_address),
        std::string(inp_peglevel_hex),

        out_balance_pegdata64,
        out->out_balance_liquid,
        out->out_balance_reserve,
        out_exchange_pegdata64,
        out->out_exchange_liquid,
        out->out_exchange_reserve,
        out_pegshift_pegdata64,
        out_requested_pegdata64,
        out_processed_pegdata64,
        out_withdraw_idxch,
        out_withdraw_txout,
        out_rawtx,
        out_txouts_str,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_balance_pegdata64 = new char[out_balance_pegdata64.length()+1];
    out->out_exchange_pegdata64 = new char[out_exchange_pegdata64.length()+1];
    out->out_pegshift_pegdata64 = new char[out_pegshift_pegdata64.length()+1];
    out->out_requested_pegdata64 = new char[out_requested_pegdata64.length()+1];
    out->out_processed_pegdata64 = new char[out_processed_pegdata64.length()+1];
    out->out_withdraw_idxch = new char[out_withdraw_idxch.length()+1];
    out->out_withdraw_txout = new char[out_withdraw_txout.length()+1];
    out->out_rawtx = new char[out_rawtx.length()+1];
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_balance_pegdata64, out_balance_pegdata64.c_str());
    std::strcpy(out->out_exchange_pegdata64, out_exchange_pegdata64.c_str());
    std::strcpy(out->out_pegshift_pegdata64, out_pegshift_pegdata64.c_str());
    std::strcpy(out->out_requested_pegdata64, out_requested_pegdata64.c_str());
    std::strcpy(out->out_processed_pegdata64, out_processed_pegdata64.c_str());
    std::strcpy(out->out_withdraw_idxch, out_withdraw_idxch.c_str());
    std::strcpy(out->out_withdraw_txout, out_withdraw_txout.c_str());
    std::strcpy(out->out_rawtx, out_rawtx.c_str());
    std::strcpy(out->out_err, out_err.c_str());

    out->out_txouts_len = out_txouts_str.size();
    if (out_txouts_str.empty()) {
        out->out_txouts = nullptr;
    } else {
        out->out_txouts = new Out_txout[out->out_txouts_len];
        struct Out_txout * out_txout_ptr = out->out_txouts;
        for(const auto & out_txout_str : out_txouts_str) {
            std::string out_txout;
            std::string out_pegdata64;
            std::string out_privkey_bip32;
            std::tie(out_txout, out_pegdata64, out_privkey_bip32) = out_txout_str;
            out_txout_ptr->out_txout = new char[out_txout.length()+1];
            out_txout_ptr->out_pegdata64 = new char[out_pegdata64.length()+1];
            out_txout_ptr->out_privkey_bip32 = new char[out_privkey_bip32.length()+1];
            std::strcpy(out_txout_ptr->out_txout, out_txout.c_str());
            std::strcpy(out_txout_ptr->out_pegdata64, out_pegdata64.c_str());
            std::strcpy(out_txout_ptr->out_privkey_bip32, out_privkey_bip32.c_str());
            out_txout_ptr++;
        }
    }

    return out;
}

void Out_preparewithdraw_del(struct Out_preparewithdraw* out)
{
    if (!out) return;
    delete[] out->out_balance_pegdata64;
    delete[] out->out_exchange_pegdata64;
    delete[] out->out_pegshift_pegdata64;
    delete[] out->out_requested_pegdata64;
    delete[] out->out_processed_pegdata64;
    delete[] out->out_rawtx;
    delete[] out->out_err;
    if (out->out_txouts) {
        Out_txouts_del(out->out_txouts, out->out_txouts_len);
    }
    delete out;
}

void Inp_txinps_del(struct Inp_txinp* inps, int n)
{
    if (!inps) return;
    for(int i=0; i<n; i++) {
        delete[] (inps+i)->inp_txout;
        delete[] (inps+i)->inp_pegdata64;
        delete[] (inps+i)->inp_privkey_bip32;
    }
    delete[] inps;
}

struct Out_txout* gettxout(struct Out_txout* outs, int i)
{
    if (!outs) return nullptr;
    return outs + i;
}

void Out_txouts_del(struct Out_txout* outs, int n)
{
    if (!outs) return;
    for(int i=0; i<n; i++) {
        delete[] (outs+i)->out_txout;
        delete[] (outs+i)->out_pegdata64;
        delete[] (outs+i)->out_privkey_bip32;
    }
    delete[] outs;
}

struct Out_decodepeglevel* decodepeglevel(
    char* inp_peglevel_hex
    )
{
    struct Out_decodepeglevel* out = new Out_decodepeglevel;

    std::string out_err;

    bool ok = pegops::getpeglevelinfo(
        std::string(inp_peglevel_hex),

        out->out_version,
        out->out_cycle_now,
        out->out_cycle_prev,
        out->out_buffer,
        out->out_peg_now,
        out->out_peg_next,
        out->out_peg_next_next,
        out->out_shift,
        out->out_shiftlastpart,
        out->out_shiftlasttotal,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

struct Out_decodepegdata* decodepegdata(
    char* inp_pegdata64
    )
{
    struct Out_decodepegdata* out = new Out_decodepegdata;

    std::string out_err;

    bool ok = pegops::getpegdatainfo(
        std::string(inp_pegdata64),

        out->out_version,
        out->out_value,
        out->out_liquid,
        out->out_reserve,
        out->out_value_hli,
        out->out_liquid_hli,
        out->out_reserve_hli,
        out->out_id,

        out->out_level_version,
        out->out_cycle_now,
        out->out_cycle_prev,
        out->out_buffer,
        out->out_peg_now,
        out->out_peg_next,
        out->out_peg_next_next,
        out->out_shift,
        out->out_shiftlastpart,
        out->out_shiftlasttotal,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_decodepeglevel_del(struct Out_decodepeglevel* out)
{
    if (!out) return;
    delete[] out->out_err;
    delete out;
}

void Out_decodepegdata_del(struct Out_decodepegdata* out)
{
    if (!out) return;
    delete[] out->out_err;
    delete out;
}

struct Out_ispositive* ispositive(
    char* inp_pegdata64
    )
{
    struct Out_ispositive* out = new Out_ispositive;

    bool out_positive;
    std::string out_err;

    bool ok = pegops::ispositive(
        std::string(inp_pegdata64),

        out_positive,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_positive = out_positive ? 1 : 0;
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_ispositive_del(struct Out_ispositive* out)
{
    if (!out) return;
    delete[] out->out_err;
    delete out;
}

struct Out_isempty* isempty(
    char* inp_pegdata64
    )
{
    struct Out_isempty* out = new Out_isempty;

    bool out_empty;
    std::string out_err;

    bool ok = pegops::isempty(
        std::string(inp_pegdata64),

        out_empty,
        out_err);

    out->out_ok = ok ? 1 : 0;
    out->out_empty = out_empty ? 1 : 0;
    out->out_err = new char[out_err.length()+1];

    std::strcpy(out->out_err, out_err.c_str());

    return out;
}

void Out_isempty_del(struct Out_isempty* out)
{
    if (!out) return;
    delete[] out->out_err;
    delete out;
}
