#ifndef gopegops_H
#define gopegops_H

#include <stdint.h>

struct Out_getpeglevel {
    int     out_ok;
    char*   out_peglevel_hex;
    int64_t out_exchange_liquid;
    int64_t out_exchange_reserve;
    char*   out_pegpool_pegdata64;
    int64_t out_pegpool_value;
    char*   out_err;
};

struct Out_updatepegbalances {
    int     out_ok;
    char*   out_balance_pegdata64;
    int64_t out_balance_liquid;
    int64_t out_balance_reserve;
    char*   out_pegpool_pegdata64;
    int64_t out_pegpool_value;
    char*   out_err;
};

struct Out_movecoins {
    int     out_ok;
    char*   out_src_pegdata64;
    int64_t out_src_liquid;
    int64_t out_src_reserve;
    char*   out_dst_pegdata64;
    int64_t out_dst_liquid;
    int64_t out_dst_reserve;
    char*   out_err;
};

struct Out_removecoins {
    int     out_ok;
    char*   out_left_pegdata64;
    int64_t out_left_liquid;
    int64_t out_left_reserve;
    char*   out_err;
};

struct Out_txoutupdate {
    int     out_ok;
    int64_t out_txout_value;
    int64_t out_txout_next_cycle_available_liquid;
    int64_t out_txout_next_cycle_available_reserve;
    int16_t out_txout_value_hli;
    int16_t out_txout_next_cycle_available_liquid_hli;
    int16_t out_txout_next_cycle_available_reserve_hli;
    char*   out_err;
};

struct Inp_txinp {
    char*   inp_txout;
    char*   inp_pegdata64;
    char*   inp_privkey_bip32;
};

struct Out_txout {
    char*   out_txout;
    char*   out_pegdata64;
    char*   out_privkey_bip32;
};

struct Out_preparewithdraw {
    int                 out_ok;
    char*               out_balance_pegdata64;
    int64_t             out_balance_liquid;
    int64_t             out_balance_reserve;
    char*               out_exchange_pegdata64;
    int64_t             out_exchange_liquid;
    int64_t             out_exchange_reserve;
    char*               out_pegshift_pegdata64;
    char*               out_requested_pegdata64;
    char*               out_processed_pegdata64;
    char*               out_withdraw_idxch;
    char*               out_withdraw_txout;
    char*               out_rawtx;
    struct Out_txout*   out_txouts;
    int                 out_txouts_len;
    char*               out_err;
};

struct Out_decodepeglevel {
    int     out_ok;
    int     out_version;
    int     out_cycle_now;
    int     out_cycle_prev;
    int     out_buffer;
    int     out_peg_now;
    int     out_peg_next;
    int     out_peg_next_next;
    int     out_shift;
    int64_t out_shiftlastpart;
    int64_t out_shiftlasttotal;
    char*   out_err;
};

struct Out_decodepegdata {
    int     out_ok;
    int     out_version;
    int64_t out_value;
    int64_t out_liquid;
    int64_t out_reserve;
    int16_t out_value_hli;
    int16_t out_liquid_hli;
    int16_t out_reserve_hli;
    int32_t out_id;
    int     out_level_version;
    int     out_cycle_now;
    int     out_cycle_prev;
    int     out_buffer;
    int     out_peg_now;
    int     out_peg_next;
    int     out_peg_next_next;
    int     out_shift;
    int64_t out_shiftlastpart;
    int64_t out_shiftlasttotal;
    char*   out_err;
};

struct Out_ispositive {
    int     out_ok;
    int     out_positive;
    char*   out_err;
};

struct Out_isempty {
    int     out_ok;
    int     out_empty;
    char*   out_err;
};

#ifdef __cplusplus
extern "C" {
#endif

struct Out_getpeglevel* getpeglevel(
    int   inp_cycle_now,
    int   inp_cycle_prev,
    int   inp_buffer,
    int   inp_peg_now,
    int   inp_peg_next,
    int   inp_peg_next_next,
    char* inp_exchange_pegdata64,
    char* inp_pegshift_pegdata64
);

struct Out_decodepeglevel* decodepeglevel(
    char* inp_peglevel_hex
);

struct Out_decodepegdata* decodepegdata(
    char* inp_pegdata64
);

struct Out_ispositive* ispositive(
    char* inp_pegdata64
);

struct Out_isempty* isempty(
    char* inp_pegdata64
);

struct Out_updatepegbalances* updatepegbalances(
    char* inp_balance_pegdata64,
    char* inp_pegpool_pegdata64,
    char* inp_peglevel_hex
);

struct Out_movecoins* movecoins(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex,
    int     inp_cross_cycles
);

struct Out_movecoins* moveliquid(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex
);

struct Out_movecoins* movereserve(
    int64_t inp_move_amount,
    char*   inp_src_pegdata64,
    char*   inp_dst_pegdata64,
    char*   inp_peglevel_hex
);

struct Out_removecoins* removecoins(
    char*   inp_from_pegdata64,
    char*   inp_remove_pegdata64
);

struct Out_txoutupdate* txoutupdate(
    char*   inp_txout_pegdata64,
    char*   inp_peglevel_hex
);

struct Inp_txinp* maketxinps(int n);

void filltxinp(
    struct Inp_txinp* inps,
    int     i, 
    char*   inp_txout,
    char*   inp_pegdata64,
    char*   inp_privkey_bip32);

struct Out_preparewithdraw* prepareliquidwithdraw(
    struct Inp_txinp*   inp_txinps,
    int                 inp_txinps_len,
    char*               inp_balance_pegdata64,
    char*               inp_exchange_pegdata64,
    char*               inp_pegshift_pegdata64,
    int64_t             inp_amount_with_fee,
    char*               inp_address,
    char*               inp_peglevel_hex
);

struct Out_preparewithdraw* preparereservewithdraw(
    struct Inp_txinp*   inp_txinps,
    int                 inp_txinps_len,
    char*               inp_balance_pegdata64,
    char*               inp_exchange_pegdata64,
    char*               inp_pegshift_pegdata64,
    int64_t             inp_amount_with_fee,
    char*               inp_address,
    char*               inp_peglevel_hex
);

struct Out_txout* gettxout(struct Out_txout* out_txouts, int i);

void Out_getpeglevel_del(struct Out_getpeglevel*);
void Out_updatepegbalances_del(struct Out_updatepegbalances*);
void Out_movecoins_del(struct Out_movecoins*);
void Out_removecoins_del(struct Out_removecoins*);
void Out_txoutupdate_del(struct Out_txoutupdate*);
void Out_preparewithdraw_del(struct Out_preparewithdraw*);
void Inp_txinps_del(struct Inp_txinp*, int n);
void Out_txouts_del(struct Out_txout*, int n);
void Out_decodepeglevel_del(struct Out_decodepeglevel*);
void Out_decodepegdata_del(struct Out_decodepegdata*);
void Out_ispositive_del(struct Out_ispositive*);
void Out_isempty_del(struct Out_isempty*);

#ifdef __cplusplus
}
#endif

#endif // gopegops_H
