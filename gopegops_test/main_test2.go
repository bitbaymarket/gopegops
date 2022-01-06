package main

import (
	"fmt"

	gopegops "../gopegops_testnet"
)

func main_test2() {
	balance := "AQAAAAAAAAABAgACAAAAAAAAAAAAAEgBAAAAAAAAeNrtlT1OAzEQhT32cpKIsyFIkFJQ0yHOwQFoIpIqXbgTVeIUaxc86ek5K6V7X2N51p7/8aZkjDHGGGOMMcYYY4wxxhhjjLk3b9/zeqr/Ke37B8ijyT+FHO38Nnlu++PjvD6DvPPS5Gf4fgB7fX0ietbEnw3Iu54fov8Vzve4tyRve9j37xeiP8A/dq+QOvb7mcgD9GU4z/zCuM5wrhD/E5Hjvop8VFJX1If+Z3J/B/H0e3tSH+Yf86efn4h93KO9JPoiFurFeJOIQ/Wn6r8q7FzAXxYX1rcM6g9yjulXfcrqxOqfRN4y8buKfmBxB9HP5oLNYRFxsflU8xoiDtanVeT9QbyTak4wPhbnJPqD1T8Nvo8sH3mwrqPvrurHuvD/FAv7n9lLoq9ZXVn+1DyFmPsyOOdL381b517V4+tvXt9Xs/wKsQGxDAJ2XAAAAAAAAHVcAAAAAAAAAwBwA20DagMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEBLTAAAAAAAAAAAAA=="
	exchange := "AQAAAAAAAAABAgACAAAAAAAAAAAAAEgBAAAAAAAAeNrtlT1OAzEQhT32cpKIsyFIkFJQ0yHOwQFoIpIqXbgTVeIUaxc86ek5K6V7X2N51p7/8aZkjDHGGGOMMcYYY4wxxhhjjLk3b9/zeqr/Ke37B8ijyT+FHO38Nnlu++PjvD6DvPPS5Gf4fgB7fX0ietbEnw3Iu54fov8Vzve4tyRve9j37xeiP8A/dq+QOvb7mcgD9GU4z/zCuM5wrhD/E5Hjvop8VFJX1If+Z3J/B/H0e3tSH+Yf86efn4h93KO9JPoiFurFeJOIQ/Wn6r8q7FzAXxYX1rcM6g9yjulXfcrqxOqfRN4y8buKfmBxB9HP5oLNYRFxsflU8xoiDtanVeT9QbyTak4wPhbnJPqD1T8Nvo8sH3mwrqPvrurHuvD/FAv7n9lLoq9ZXVn+1DyFmPsyOOdL381b517V4+tvXt9Xs/wKsQGxDAJ2XAAAAAAAAHVcAAAAAAAAAwBwA20DagMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEBLTAAAAAAAAAAAAA=="
	txinppegdata := "AQAAAAAAAAABAgACAAAAAAAAAAAAMjc2YTkxNDE2NzZmNzYzNDBkNDlhZmYwY2FkOGU3MWU0MTA0MmQzNzdlNDVjNzQ4OGFjMAAAAAAAAAB42u3FMQ0AIAwAsJnBBEr4MbWDeSWgYALapxEAAAAAAAAAAAB0nJ3/UWu+L0YCAuwCdlwAAAAAAAB1XAAAAAAAAAMAcANtA2oDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/3oc7AAAAAAAAAAA="
	peglevel := "02765c000000000000755c000000000000030070036d036a03000000000000000000000000000000000000"

	fmt.Println("Updatetxout:")
	txout_value, txout_nc_liquid, txout_nc_reserve,
		txout_value_hli, txout_liquid_hli, txout_reserve_hli,
		err := gopegops.Updatetxout(txinppegdata, peglevel)
	if err != nil {
		fmt.Println("Updatetxout failed: ", err)
		return
	}
	fmt.Println("Updatetxout passed")
	fmt.Println("---------")
	fmt.Println("txout value:", txout_value)
	fmt.Println("txout next cycle liquid:", txout_nc_liquid)
	fmt.Println("txout next cycle reserve:", txout_nc_reserve)
	fmt.Println("---------")
	fmt.Println("txout value hli:", txout_value_hli)
	fmt.Println("txout next cycle liquid hli:", txout_liquid_hli)
	fmt.Println("txout next cycle reserve hli:", txout_reserve_hli)
	fmt.Println("---------")

	txinps := make([]*gopegops.TxOut, 0)
	txinp1 := new(gopegops.TxOut)
	txinp1.TxoutId = "0191045773206006fd4938fc97ef389d5d2b13a8f2287206360abdd70c7964b9:1"
	txinp1.Pegdata64 = txinppegdata
	txinp1.PrivKeyBip32 = "tprv8ezDmewmXkNvKU8dZbpfRigkWSPLaATumcZwMZ4vufnthHo7LmZGuTBtMUL28iynLqVfQH3EkbMz4KxD3Bv7zYRtxgedbp2YdwtNPeVqjBj"
	txinps = append(txinps, txinp1)

	out_balance, out_balance_liquid, out_balance_reserve,
		out_exchange, out_exchange_liquid, out_exchange_reserve,
		out_pegshift, out_requested, out_processed,
		withdraw_idxch, withdraw_txout, rawtx, change_txouts, err := gopegops.Prepareliquidwithdraw(
		txinps,
		balance,
		exchange,
		"",
		2000000,
		"mhZjgCf1mJsL7cnCe225jsoZsKPUWpsrf7",
		peglevel)

	if err != nil {
		fmt.Println("Prepareliquidwithdraw failed: ", err)
		return
	}
	fmt.Println("Prepareliquidwithdraw passed")

	fmt.Println("---------")
	fmt.Println("withdraw_idxch:", withdraw_idxch)
	fmt.Println("withdraw_txout:", withdraw_txout)
	fmt.Println("---------")
	fmt.Println("rawtx:", rawtx)
	for i := 0; i < len(change_txouts); i++ {
		change_txout := change_txouts[i]
		fmt.Println("---------")
		fmt.Println("change:", change_txout)
	}
	fmt.Println("---------")
	fmt.Println("balance:", out_balance)
	fmt.Println("balance liquid:", out_balance_liquid)
	fmt.Println("balance reserve:", out_balance_reserve)
	fmt.Println("---------")
	fmt.Println("exchange:", out_exchange)
	fmt.Println("exchange liquid:", out_exchange_liquid)
	fmt.Println("exchange reserve:", out_exchange_reserve)
	fmt.Println("---------")
	fmt.Println("pegshift:", out_pegshift)
	fmt.Println("---------")
	fmt.Println("requested:", out_requested)
	fmt.Println("---------")
	fmt.Println("processed:", out_processed)
	fmt.Println("---------")
}
