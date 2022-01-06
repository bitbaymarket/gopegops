package gopegops_mainnet

import (
	"fmt"
	"unsafe"
)

/*
#cgo CXXFLAGS: -std=gnu++11
#cgo LDFLAGS: -lboost_system -lboost_thread -lssl -lcrypto -lz
#include <stdlib.h>
#include "gopegops.h"
*/
import "C"

// Makepeglevel return peglevel
func Makepeglevel(
	inpCycleNow int,
	inpCyclePrev int,

	inpBuffer int,
	inpPeg int,
	inpPegNext int,
	inpPegNextNext int,

	inpExchangePegdata64 string,
	inpPegshiftPegdata64 string) (

	outPeglevelHex string,
	outExchangeLiquid int64,
	outExchangeReserve int64,
	outPegpoolPegdata64 string,
	outPegpoolValue int64,
	err error) {

	if inpCyclePrev >= inpCycleNow {
		err = fmt.Errorf("inpCyclePrev %d is equal or greater inpCycleNow %d",
			inpCyclePrev, inpCycleNow)
		return
	}

	cInpExchangePegdata64 := C.CString(inpExchangePegdata64)
	cInpPegshiftPegdata64 := C.CString(inpPegshiftPegdata64)

	ptr := C.getpeglevel(
		C.int(inpCycleNow),
		C.int(inpCyclePrev),
		C.int(inpBuffer),
		C.int(inpPeg),
		C.int(inpPegNext),
		C.int(inpPegNextNext),
		cInpExchangePegdata64,
		cInpPegshiftPegdata64)

	outPeglevelHex = C.GoString(ptr.out_peglevel_hex)
	outPegpoolPegdata64 = C.GoString(ptr.out_pegpool_pegdata64)
	outExchangeLiquid = int64(ptr.out_exchange_liquid)
	outExchangeReserve = int64(ptr.out_exchange_reserve)
	outPegpoolValue = int64(ptr.out_pegpool_value)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpExchangePegdata64))
	C.free(unsafe.Pointer(cInpPegshiftPegdata64))
	C.Out_getpeglevel_del(ptr)
	return
}

// Decoding peglevel hex
type Peglevel struct {
	CycleNow       int
	CyclePrev      int
	Buffer         int
	PegNow         int
	PegNext        int
	PegNextNext    int
	Shift          int
	ShiftLastPart  int64
	ShiftLastTotal int64
}

func Decodepeglevel(inpPeglevelHex string) (

	out *Peglevel,
	err error) {

	cInpPeglevelHex := C.CString(inpPeglevelHex)

	ptr := C.decodepeglevel(cInpPeglevelHex)

	out = nil
	if ptr.out_ok == 1 {
		out = new(Peglevel)

		out.CycleNow = int(ptr.out_cycle_now)
		out.CyclePrev = int(ptr.out_cycle_prev)
		out.Buffer = int(ptr.out_buffer)
		out.PegNow = int(ptr.out_peg_now)
		out.PegNext = int(ptr.out_peg_next)
		out.PegNextNext = int(ptr.out_peg_next_next)
		out.Shift = int(ptr.out_shift)
		out.ShiftLastPart = int64(ptr.out_shiftlastpart)
		out.ShiftLastTotal = int64(ptr.out_shiftlasttotal)
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_decodepeglevel_del(ptr)
	return
}

// Decoding pegdata base64
type Pegdata struct {
	Level      *Peglevel
	Value      int64
	Liquid     int64
	Reserve    int64
	ValueHLI   int16
	LiquidHLI  int16
	ReserveHLI int16
	Id         int32
	//Fractions [1200]int64 // todo
}

func Decodepegdata(inpPegdata64 string) (

	out *Pegdata,
	err error) {

	cInpPegdata64 := C.CString(inpPegdata64)

	ptr := C.decodepegdata(cInpPegdata64)

	out = nil
	if ptr.out_ok == 1 {
		out = new(Pegdata)
		out.Level = new(Peglevel)

		out.Value = int64(ptr.out_value)
		out.Liquid = int64(ptr.out_liquid)
		out.Reserve = int64(ptr.out_reserve)
		out.ValueHLI = int16(ptr.out_value_hli)
		out.LiquidHLI = int16(ptr.out_liquid_hli)
		out.ReserveHLI = int16(ptr.out_reserve_hli)
		out.Id = int32(ptr.out_id)

		out.Level.CycleNow = int(ptr.out_cycle_now)
		out.Level.CyclePrev = int(ptr.out_cycle_prev)
		out.Level.Buffer = int(ptr.out_buffer)
		out.Level.PegNow = int(ptr.out_peg_now)
		out.Level.PegNext = int(ptr.out_peg_next)
		out.Level.PegNextNext = int(ptr.out_peg_next_next)
		out.Level.Shift = int(ptr.out_shift)
		out.Level.ShiftLastPart = int64(ptr.out_shiftlastpart)
		out.Level.ShiftLastTotal = int64(ptr.out_shiftlasttotal)
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpPegdata64))
	C.Out_decodepegdata_del(ptr)
	return
}

// All fractions are non-negative
func Ispositive(inpPegdata64 string) (

	positive bool,
	err error) {

	cInpPegdata64 := C.CString(inpPegdata64)

	ptr := C.ispositive(cInpPegdata64)

	if ptr.out_ok == 1 {
		if ptr.out_positive == 0 {
			positive = false
		} else {
			positive = true
		}
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpPegdata64))
	C.Out_ispositive_del(ptr)
	return
}

// All fractions are zeros
func Isempty(inpPegdata64 string) (

	empty bool,
	err error) {

	cInpPegdata64 := C.CString(inpPegdata64)

	ptr := C.isempty(cInpPegdata64)

	if ptr.out_ok == 1 {
		if ptr.out_empty == 0 {
			empty = false
		} else {
			empty = true
		}
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpPegdata64))
	C.Out_isempty_del(ptr)
	return
}

// Updatepegbalances to update peg balances
func Updatepegbalances(
	inpBalancePegdata64 string,
	inpPegpoolPegdata64 string,
	inpPeglevelHex string) (

	outBalancePegdata64 string,
	outBalanceLiquid int64,
	outBalanceReserve int64,
	outPegpoolPegdata64 string,
	outPegpoolValue int64,
	err error) {

	cInpBalancePegdata64 := C.CString(inpBalancePegdata64)
	cInpPegpoolPegdata64 := C.CString(inpPegpoolPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)

	ptr := C.updatepegbalances(
		cInpBalancePegdata64,
		cInpPegpoolPegdata64,
		cInpPeglevelHex)

	outBalancePegdata64 = C.GoString(ptr.out_balance_pegdata64)
	outPegpoolPegdata64 = C.GoString(ptr.out_pegpool_pegdata64)
	outBalanceLiquid = int64(ptr.out_balance_liquid)
	outBalanceReserve = int64(ptr.out_balance_reserve)
	outPegpoolValue = int64(ptr.out_pegpool_value)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpBalancePegdata64))
	C.free(unsafe.Pointer(cInpPegpoolPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_updatepegbalances_del(ptr)
	return
}

// Movecoins to move coins beetween balances (pegdata)
// requires equal cycle in src, dst and peglevel (inpCrossCycles=false)
// auxiliary operations - move coins without cycle check (inpCrossCycles=false)
// resulting src and dst pegdata have cycle of peglevel
func Movecoins(
	inpMoveAmount int64,
	inpSrcPegdata64 string,
	inpDstPegdata64 string,
	inpPeglevelHex string,
	inpCrossCycles bool) (

	outSrcPegdata64 string,
	outSrcLiquid int64,
	outSrcReserve int64,
	outDstPegdata64 string,
	outDstLiquid int64,
	outDstReserve int64,
	err error) {

	cInpSrcPegdata64 := C.CString(inpSrcPegdata64)
	cInpDstPegdata64 := C.CString(inpDstPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)
	iInpCrossCycles := 0
	if inpCrossCycles {
		iInpCrossCycles = 1
	}

	ptr := C.movecoins(
		C.int64_t(inpMoveAmount),
		cInpSrcPegdata64,
		cInpDstPegdata64,
		cInpPeglevelHex,
		C.int(iInpCrossCycles))

	outSrcPegdata64 = C.GoString(ptr.out_src_pegdata64)
	outSrcLiquid = int64(ptr.out_src_liquid)
	outSrcReserve = int64(ptr.out_src_reserve)
	outDstPegdata64 = C.GoString(ptr.out_dst_pegdata64)
	outDstLiquid = int64(ptr.out_dst_liquid)
	outDstReserve = int64(ptr.out_dst_reserve)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpSrcPegdata64))
	C.free(unsafe.Pointer(cInpDstPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_movecoins_del(ptr)
	return
}

// Moveliquid to move liquid coins beetween balances (pegdata)
// requires equal cycle in src, dst and peglevel
func Moveliquid(
	inpMoveAmount int64,
	inpSrcPegdata64 string,
	inpDstPegdata64 string,
	inpPeglevelHex string) (

	outSrcPegdata64 string,
	outSrcLiquid int64,
	outSrcReserve int64,
	outDstPegdata64 string,
	outDstLiquid int64,
	outDstReserve int64,
	err error) {

	cInpSrcPegdata64 := C.CString(inpSrcPegdata64)
	cInpDstPegdata64 := C.CString(inpDstPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)

	ptr := C.moveliquid(
		C.int64_t(inpMoveAmount),
		cInpSrcPegdata64,
		cInpDstPegdata64,
		cInpPeglevelHex)

	outSrcPegdata64 = C.GoString(ptr.out_src_pegdata64)
	outSrcLiquid = int64(ptr.out_src_liquid)
	outSrcReserve = int64(ptr.out_src_reserve)
	outDstPegdata64 = C.GoString(ptr.out_dst_pegdata64)
	outDstLiquid = int64(ptr.out_dst_liquid)
	outDstReserve = int64(ptr.out_dst_reserve)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpSrcPegdata64))
	C.free(unsafe.Pointer(cInpDstPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_movecoins_del(ptr)
	return
}

// Movereserve to move reserve coins beetween balances (pegdata)
// requires equal cycle in src, dst and peglevel
func Movereserve(
	inpMoveAmount int64,
	inpSrcPegdata64 string,
	inpDstPegdata64 string,
	inpPeglevelHex string) (

	outSrcPegdata64 string,
	outSrcLiquid int64,
	outSrcReserve int64,
	outDstPegdata64 string,
	outDstLiquid int64,
	outDstReserve int64,
	err error) {

	cInpSrcPegdata64 := C.CString(inpSrcPegdata64)
	cInpDstPegdata64 := C.CString(inpDstPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)

	ptr := C.movereserve(
		C.int64_t(inpMoveAmount),
		cInpSrcPegdata64,
		cInpDstPegdata64,
		cInpPeglevelHex)

	outSrcPegdata64 = C.GoString(ptr.out_src_pegdata64)
	outSrcLiquid = int64(ptr.out_src_liquid)
	outSrcReserve = int64(ptr.out_src_reserve)
	outDstPegdata64 = C.GoString(ptr.out_dst_pegdata64)
	outDstLiquid = int64(ptr.out_dst_liquid)
	outDstReserve = int64(ptr.out_dst_reserve)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpSrcPegdata64))
	C.free(unsafe.Pointer(cInpDstPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_movecoins_del(ptr)
	return
}

// Removecoins to substitute 'remove' pegdata 'from' pegdata
// requires equal cycle in src, dst and peglevel
func Removecoins(
	inpFromPegdata64 string,
	inpRemovePegdata64 string) (

	outLeftPegdata64 string,
	outLeftLiquid int64,
	outLeftReserve int64,
	err error) {

	cInpFromPegdata64 := C.CString(inpFromPegdata64)
	cInpRemovePegdata64 := C.CString(inpRemovePegdata64)

	ptr := C.removecoins(
		cInpFromPegdata64,
		cInpRemovePegdata64)

	outLeftPegdata64 = C.GoString(ptr.out_left_pegdata64)
	outLeftLiquid = int64(ptr.out_left_liquid)
	outLeftReserve = int64(ptr.out_left_reserve)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpFromPegdata64))
	C.free(unsafe.Pointer(cInpRemovePegdata64))
	C.Out_removecoins_del(ptr)
	return
}

// Updatetxout returns next cycle information about txout
// available liquid and reserve (for withdraw) and HLI indexes
// (Half Liquidity Index - level of peg when half of coins turns into reserve)
func Updatetxout(
	inpTxoutPegdata64 string,
	inpPeglevelHex string) (

	outTxoutValue int64,
	outTxoutNextCycleAvailableLiquid int64,
	outTxoutNextCycleAvailableReserve int64,
	outTxoutValueHLI int16,
	outTxoutNextCycleAvailableLiquidHLI int16,
	outTxoutNextCycleAvailableReserveHLI int16,
	err error) {

	cInpTxoutPegdata64 := C.CString(inpTxoutPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)

	ptr := C.txoutupdate(
		cInpTxoutPegdata64,
		cInpPeglevelHex)

	outTxoutValue = int64(ptr.out_txout_value)
	outTxoutNextCycleAvailableLiquid = int64(ptr.out_txout_next_cycle_available_liquid)
	outTxoutNextCycleAvailableReserve = int64(ptr.out_txout_next_cycle_available_reserve)
	outTxoutValueHLI = int16(ptr.out_txout_value_hli)
	outTxoutNextCycleAvailableLiquidHLI = int16(ptr.out_txout_next_cycle_available_liquid_hli)
	outTxoutNextCycleAvailableReserveHLI = int16(ptr.out_txout_next_cycle_available_reserve_hli)

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.free(unsafe.Pointer(cInpTxoutPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.Out_txoutupdate_del(ptr)
	return
}

type TxOut struct {
	TxoutId      string // txhash:nout
	Pegdata64    string
	PrivKeyBip32 string
}

func Prepareliquidwithdraw(
	inpTxInps []*TxOut,
	inpBalancePegdata64 string,
	inpExchangePegdata64 string,
	inpPegshiftPegdata64 string,
	inpAmountWithFee int64,
	inpAddress string,
	inpPeglevelHex string) (

	outBalancePegdata64 string,
	outBalanceLiquid int64,
	outBalanceReserve int64,
	outExchangePegdata64 string,
	outExchangeLiquid int64,
	outExchangeReserve int64,
	outPegshiftPegdata64 string,
	outRequestedPegdata64 string,
	outProcessedPegdata64 string,
	outWithdrawIdXch string,
	outWithdrawTxout string,
	outRawTx string,
	outTxOuts []*TxOut,
	err error) {

	cInpBalancePegdata64 := C.CString(inpBalancePegdata64)
	cInpExchangePegdata64 := C.CString(inpExchangePegdata64)
	cInpPegshiftPegdata64 := C.CString(inpPegshiftPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)
	cInpAddress := C.CString(inpAddress)

	txinps := C.maketxinps(C.int(len(inpTxInps)))
	for i := 0; i < len(inpTxInps); i++ {
		inpTxInp := inpTxInps[i]
		cInpTxoutId := C.CString(inpTxInp.TxoutId)
		cInpPegdata64 := C.CString(inpTxInp.Pegdata64)
		cInpPrivKeyBip32 := C.CString(inpTxInp.PrivKeyBip32)
		C.filltxinp(txinps, C.int(i), cInpTxoutId, cInpPegdata64, cInpPrivKeyBip32)
		C.free(unsafe.Pointer(cInpTxoutId))
		C.free(unsafe.Pointer(cInpPegdata64))
		C.free(unsafe.Pointer(cInpPrivKeyBip32))
	}

	ptr := C.prepareliquidwithdraw(
		txinps,
		C.int(len(inpTxInps)),
		cInpBalancePegdata64,
		cInpExchangePegdata64,
		cInpPegshiftPegdata64,
		C.int64_t(inpAmountWithFee),
		cInpAddress,
		cInpPeglevelHex)

	outBalancePegdata64 = C.GoString(ptr.out_balance_pegdata64)
	outBalanceLiquid = int64(ptr.out_balance_liquid)
	outBalanceReserve = int64(ptr.out_balance_reserve)
	outExchangePegdata64 = C.GoString(ptr.out_exchange_pegdata64)
	outExchangeLiquid = int64(ptr.out_exchange_liquid)
	outExchangeReserve = int64(ptr.out_exchange_reserve)
	outPegshiftPegdata64 = C.GoString(ptr.out_pegshift_pegdata64)
	outRequestedPegdata64 = C.GoString(ptr.out_requested_pegdata64)
	outProcessedPegdata64 = C.GoString(ptr.out_processed_pegdata64)
	outWithdrawIdXch = C.GoString(ptr.out_withdraw_idxch)
	outWithdrawTxout = C.GoString(ptr.out_withdraw_txout)
	outRawTx = C.GoString(ptr.out_rawtx)

	outTxOuts = make([]*TxOut, 0)
	for i := 0; i < int(ptr.out_txouts_len); i++ {
		txout_ptr := C.gettxout(ptr.out_txouts, C.int(i))
		txout := new(TxOut)
		txout.TxoutId = C.GoString(txout_ptr.out_txout)
		txout.Pegdata64 = C.GoString(txout_ptr.out_pegdata64)
		txout.PrivKeyBip32 = C.GoString(txout_ptr.out_privkey_bip32)
		outTxOuts = append(outTxOuts, txout)
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.Inp_txinps_del(txinps, C.int(len(inpTxInps)))
	C.free(unsafe.Pointer(cInpBalancePegdata64))
	C.free(unsafe.Pointer(cInpExchangePegdata64))
	C.free(unsafe.Pointer(cInpPegshiftPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.free(unsafe.Pointer(cInpAddress))
	C.Out_preparewithdraw_del(ptr)
	return
}

func Preparereservewithdraw(
	inpTxInps []*TxOut,
	inpBalancePegdata64 string,
	inpExchangePegdata64 string,
	inpPegshiftPegdata64 string,
	inpAmountWithFee int64,
	inpAddress string,
	inpPeglevelHex string) (

	outBalancePegdata64 string,
	outBalanceLiquid int64,
	outBalanceReserve int64,
	outExchangePegdata64 string,
	outExchangeLiquid int64,
	outExchangeReserve int64,
	outPegshiftPegdata64 string,
	outRequestedPegdata64 string,
	outProcessedPegdata64 string,
	outWithdrawIdXch string,
	outWithdrawTxout string,
	outRawTx string,
	outTxOuts []*TxOut,
	err error) {

	cInpBalancePegdata64 := C.CString(inpBalancePegdata64)
	cInpExchangePegdata64 := C.CString(inpExchangePegdata64)
	cInpPegshiftPegdata64 := C.CString(inpPegshiftPegdata64)
	cInpPeglevelHex := C.CString(inpPeglevelHex)
	cInpAddress := C.CString(inpAddress)

	txinps := C.maketxinps(C.int(len(inpTxInps)))
	for i := 0; i < len(inpTxInps); i++ {
		inpTxInp := inpTxInps[i]
		cInpTxoutId := C.CString(inpTxInp.TxoutId)
		cInpPegdata64 := C.CString(inpTxInp.Pegdata64)
		cInpPrivKeyBip32 := C.CString(inpTxInp.PrivKeyBip32)
		C.filltxinp(txinps, C.int(i), cInpTxoutId, cInpPegdata64, cInpPrivKeyBip32)
		C.free(unsafe.Pointer(cInpTxoutId))
		C.free(unsafe.Pointer(cInpPegdata64))
		C.free(unsafe.Pointer(cInpPrivKeyBip32))
	}

	ptr := C.preparereservewithdraw(
		txinps,
		C.int(len(inpTxInps)),
		cInpBalancePegdata64,
		cInpExchangePegdata64,
		cInpPegshiftPegdata64,
		C.int64_t(inpAmountWithFee),
		cInpAddress,
		cInpPeglevelHex)

	outBalancePegdata64 = C.GoString(ptr.out_balance_pegdata64)
	outBalanceLiquid = int64(ptr.out_balance_liquid)
	outBalanceReserve = int64(ptr.out_balance_reserve)
	outExchangePegdata64 = C.GoString(ptr.out_exchange_pegdata64)
	outExchangeLiquid = int64(ptr.out_exchange_liquid)
	outExchangeReserve = int64(ptr.out_exchange_reserve)
	outPegshiftPegdata64 = C.GoString(ptr.out_pegshift_pegdata64)
	outRequestedPegdata64 = C.GoString(ptr.out_requested_pegdata64)
	outProcessedPegdata64 = C.GoString(ptr.out_processed_pegdata64)
	outWithdrawIdXch = C.GoString(ptr.out_withdraw_idxch)
	outWithdrawTxout = C.GoString(ptr.out_withdraw_txout)
	outRawTx = C.GoString(ptr.out_rawtx)

	outTxOuts = make([]*TxOut, 0)
	for i := 0; i < int(ptr.out_txouts_len); i++ {
		txout_ptr := C.gettxout(ptr.out_txouts, C.int(i))
		txout := new(TxOut)
		txout.TxoutId = C.GoString(txout_ptr.out_txout)
		txout.Pegdata64 = C.GoString(txout_ptr.out_pegdata64)
		txout.PrivKeyBip32 = C.GoString(txout_ptr.out_privkey_bip32)
		outTxOuts = append(outTxOuts, txout)
	}

	if ptr.out_ok != 1 {
		err = fmt.Errorf(C.GoString(ptr.out_err))
	} else {
		err = nil
	}

	C.Inp_txinps_del(txinps, C.int(len(inpTxInps)))
	C.free(unsafe.Pointer(cInpBalancePegdata64))
	C.free(unsafe.Pointer(cInpExchangePegdata64))
	C.free(unsafe.Pointer(cInpPegshiftPegdata64))
	C.free(unsafe.Pointer(cInpPeglevelHex))
	C.free(unsafe.Pointer(cInpAddress))
	C.Out_preparewithdraw_del(ptr)
	return
}
