package main

import (
	"fmt"

	//gopegops_mainnet "../gopegops_mainnet"
	gopegops "../gopegops_testnet"
)

func main() {
	fmt.Println(gopegops.Makepeglevel(12, 11, 0, 10, 20, 30, "", ""))

	balanceBase64 := "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"
	pegpoolBase64 := "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"
	peglevelHex1 := "6e2d0000000000006d2d000000000000f700fa00fd00000000000000000000000000000000000000"

	empty, err := gopegops.Isempty(balanceBase64)
	if err != nil {
		fmt.Println("Isempty failed: ", err)
		return
	}
	if empty == true {
		fmt.Println("Isempty not passed: true")
		return
	}

	fmt.Println("Isempty passed")

	positive, err := gopegops.Ispositive(balanceBase64)
	if err != nil {
		fmt.Println("Ispositive failed: ", err)
		return
	}
	if positive == false {
		fmt.Println("Ispositive not passed: false")
		return
	}

	fmt.Println("Ispositive passed")

	/*
		outLeftPegdata64 string,
		outLeftLiquid int64,
		outLeftReserve int64,
		err error
	*/

	zeroBase64, zeroLiquid, zeroReserve, err := gopegops.Removecoins(balanceBase64, balanceBase64)
	if err != nil {
		fmt.Println("Removecoins failed: ", err)
		return
	}
	if zeroLiquid != 0 || zeroReserve != 0 {
		fmt.Println("Removecoins not passed: false")
		return
	}
	empty, err = gopegops.Isempty(zeroBase64)
	if err != nil {
		fmt.Println("Isempty failed: ", err)
		return
	}
	if empty != true {
		fmt.Println("Isempty not passed: false")
		return
	}

	fmt.Println("Removecoins passed, substract itself")

	minusBase64, minusLiquid, minusReserve, err := gopegops.Removecoins(zeroBase64, balanceBase64)
	if err != nil {
		fmt.Println("Removecoins failed: ", err)
		return
	}
	if minusLiquid != -1160175057156 || minusReserve != 0 {
		fmt.Println("Removecoins not passed: false", minusLiquid, minusReserve)
		return
	}
	empty, err = gopegops.Isempty(minusBase64)
	if err != nil {
		fmt.Println("Isempty failed: ", err)
		return
	}
	if empty != false {
		fmt.Println("Isempty not passed: true")
		return
	}
	positive, err = gopegops.Ispositive(minusBase64)
	if err != nil {
		fmt.Println("Ispositive failed: ", err)
		return
	}
	if positive == true {
		fmt.Println("Ispositive not passed: true")
		return
	}

	fmt.Println("Removecoins passed, substract second time itself")

	_, _, _, _, _, err = gopegops.Updatepegbalances(balanceBase64, pegpoolBase64, peglevelHex1)
	if err != nil {
		fmt.Println("Updatepegbalances failed: ", err)
		return
	}
	fmt.Println("Updatepegbalances passed")

	user_src := "AgACAAAAAAAAAAAAUgAAAAAAAAB42u3WMQ4AIQgEQOX/j/YDWhgKhcw0G4srLi6RMQAAAIDfTdk6Md+7PmT79LrP+k1mHrLz0f39jKL/GZf3F4fz7fdV9wJ7kpT2ZwCgjgWk0QEZAgAAAAAAAAABAAAAAAAAAPcB9wH3AQAAtAAAAAAAAABLAQAAAAAAAAAAAAAAAAAAOwEAAAAAAAA="
	user_dst := "AgACAAAAAAAAAAAAIAAAAAAAAAB42u3BMQEAAADCoPVPbQdvoAAAAAAAAAAAAHgMJYAAAQIAAAAAAAAAAQAAAAAAAAD3AfcB9wEAALQAAAAAAAAASwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	peglevel_hex := "02000000000000000100000000000000f701f701f7010000b4000000000000004b01000000000000"

	_, _, _, _, _, _, err = gopegops.Moveliquid(283, user_src, user_dst, peglevel_hex)
	if err != nil {
		fmt.Println("Moveliquid failed: ", err)
		return
	}
	fmt.Println("Moveliquid passed")

	// "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"
	// "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"
	// ""
	// 2000065
	// "mzRseDcAPwP8HjPvSM1cyCLHSuVH4sEbDE"
	// "02775a000000000000765a0000000000000300af04af04af04000000000000000000000000000000000000"
	// "5659c5650958009b40cc493d75af1f3d84fa95b2763f5cacc30ef1462fc71aa6:0"
	// "tprv8itwXtxKu2di1Mj76WvAHTnb883sCE1orJsKhBFEQ6VU8SxChxSpNy6PYaDP5MrWfgejvmFLJzAHgGEJ4sKPUXnuyHsUqELxpiBwmpjeezG"

	balance := "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"
	exchange := "AQAAAAAAAAABAgACAAAAAAAAAAAAAEoLAAAAAAAAeNptmneYT1cax+8dYgZRJkpCZhkk+iayiIQoUROijDIR3Si7q2eVaIMhWghhU8YTZRNL1MUwo3dCIoXRmZAMo4dg2QRz7/7mdz/fPHvOs79/3nnLKfect58Z0XR6Vj/Hcbo4we+m54d/CXerOumu42T/FuAV3YDf/X6AlwCf+PW8MFzN+GehjwZm/hrIf/AogK2yA9i5WtScfCGZM5cCvMu5+EbRIV5T5vkKOK58wsMc/mT20fvo3PDMab1LLX0uwnGqsd89rHeE+YtkpmzLgZu2Lo4PiTmJNZpc2BHCHw/Y/uoTY1JmhcbkZdyMh8EG24BXZb89vnFz5+B12gaM3ezrA2Cvm8UHRIfgmIeB/MA8AT2qVeUw7BsR4AX/E/CTj70ThpceC+gD7gb0KL6jPuu2gT+J8asiy4ThHvip7HMcsAX0n1knLndAX5S/X3jmJOavCb1lj81JJUPwVc5rAesM5pyXIVcW/gjgmcHJCTn0r/n+06w/k/VPutMO1g6t1bjNr/tvh/AU5h0PfyXzdmY/vZgHtfLHQ38fuLhiwF8Lf+m6w/tzWPHgTVj/KPKDmG8d+01Y0uDjeyHY9NH5ojn0dORrw49lnvzsM+mZAMYNOdXzeIj+JvLfIn+XA8/N/ZRg/O4D/86Vc/Tzke8J3QNfy/6aQK8BHFo9MXwih5HTOT/amRyZg0fmCuj34Kf3qvdKVoifwX7XAlNHXQzsIHJjGDbnHDYx7hzrvQld6//GvqKgH0B/dK+Noe9FvgT4IMalPhYX7O/TB2H4DeO6ABsCX2O87vE04+9L79ELD7wj6ywCf4Fxx5gvFvprWu9BgG+GnntEABswzgG+y7ylOefvof+ReTpEb8gxC+dt5HSOQ5Fjm85dxu+A3hTYdVAA+3MvFZGbzfw1gHmZZy7jCiPXhHGXwZcCn2bcUOAPfGdl8OeZ58+OOf8w4CeeeT6YuT8a/m30zGO969Jr5CtEBvxd4PfgYwbOlIaBAywDvzB03L9/GZgCPxfr1YQ+t9YF475aIneCeeKAS+FvAvrsdwn88pxHHOMzEcvAwdTlAv+KXGXGDUF+FXgnYAfml9+f6ZvfX032AD4DfAV4HuB+4OeWPUSDF+PeMzzTTiOgT4YezbhI6LPAPwZeY78X4V+QnQDbMm8t1v3SM+N7W/Btjjn/++CxyL0Arn3KvrbDlx9vAb6bfU1FfCH0Jdqf4j2wHPRY8DnAd1lnEbA48EXmGwbszHoNOQflCdXlh6x7Pgh9JHghxj0P3go4A7l0+SPmfQr+eOj3tV/207FmAJXfXAVOYrzykOPgpBHOd+DJ8PcxP2HK2Qb/KPidbPM+80HPo/ih8Ywbw3fWQ64Y/D2WPcwHHwy+RXkQ9PfABzJ+HfZ9RfEDOdmvvveOY8bHw/Dlt74DPwR+iO+bBz2f9Ah8i2/q7+fAvsAqrHOJcU2s7xLUrzgw2Tf9pvToFeC/FG+Zrx/zZ8A/A38UsLdnji/JPfwF/p/gzwaX35bd1bT2VZ314mV30LvlMs9X+jlScY95dO595F+lD0Dle6Xgz2M9xc82+k7Wkd7JHmeQ/z3rmPlMPb57ouJrhKnnJZHfA/5I5wZ9H/O0g94duYXwG4FLvy8Af0aP/ik90P37ZjxR/udmm/r2knWfW1h/LvjLwBHAosz7BXCHpZ9HGH9kTQC/ZJ0U1m0v/YWeqDgJ/ZrsH/oInZfiFOMUj7YDv5If1Xrsg3D/e95YVXEDPIp7yo/8R/DbAfWdiv/tgNLbO4yTfsxUfIWeIT3xTTtQHI0BVz72keo6xu9iXsWHdMf0v/LzCfArKG9iXH/m+QV+Fyuvyg39Kd/MV7sCK0F/2TXz8L9Dr3W4q/e/91SJcVnKZ6GPteJbquzN0qsy4Gngc9h/Fvc0DXwd8HXXPM+NjulHC/tmfaY89zQwC1gH/mrL3too/wVvzrrHwJNc85xaqY6V/QHLKj91zDowk/kWKI8FDmfd6cj3sPTqE/CVjqkf8lfngbukv+CUob/7KfU5pmJgdWicyO+kaT7Vpcif8s142Ry8AXh98AmO6QdVT8nP6Pv2VXox7CFKQR/H93cDP2/57XWyE+U9yJVi3A3wgvDfYt2x0Msg1w6/VNs143yy9Ax8PXCD8iTl54rTlv2UlP9SPca4htATWb8jeCr8Fa75nUus+CW9eQOoPKega9bxO5F/UnUm9FlW/qW8f7Zn6u8NzmUY+4wBbwX+nGvGn0LW/U9zzXirfH669X3yo7qnZRRYU5Erj5zqoavs86yVb5YDVpCeAVeqv6K6CvgTsBn8k0Dl7QXAi1j+sC30FuDyP5Us/dR6Lay6Q3qie+kAPYnziYH+IXTVoX2hF2H8cfijre+6AV/+U3Y+wzX1TfvS949V3aG8SfUO+EPFCc88p0zL/yiel0ZP2ntmfdOYeVVXJ8JvB/2WZdfSc9VPMcqLkJ/A+OGyN+QmuGZdoviVqHOz+nSZrmm3J3yzzyV/qLqjrvrGVh4WCX2NzifbrJelv6mcz1X4G6XHzL/eM+u4aOBW9R+UZzJO+epV5u0DrrqwCrCsZ+Zl6odtBC5W/QFf9YnykmTL/66XH1Ie6Jt+V/svYvWb1Rf/yTP9mOLpKt/MW95AfjjyryL3BHC+4pn6SZadd1J+lm3G9bnIP8H480DVzQcZ/xL0W4zXeRZj3h2uWe/fUj+Jca1V30aY9Xdr+UfPPFfFTdr1juwvzTPtWnai+r2ZlY9LTxaAy69WR769/Cb8B56Z78Za9684fNs160TlVcrLrsHf6pn+9YaV319XPwh6PeW70Ie4Zn0ySfoJ/q3yOPkZ+WvX9D+bwdOs+K/40gz6euR2ghdlvrrInQJf5Zr9ou5W/v7AMeN+Vat+1j6nW331aOuelG+oTpB/Vv1XzNr3FfDl8P9hxf+R8P/A9/dSv82KN7K/+noHg54fuc+gZ1l+6irjJin+K88D7nXNOKq+ifLqAuxL9a/6fknIaf01sjfs8dMIU2/kzxU/7jGui9XXlV5UtuyrP/zqVr+0NHANcgOtPstZ+AWY957VJ1Ge8Av8PlZfX/dXzqoL1W96h3Gym63K51XnuWaeqPcl5edzrP6k6oEfrXxB9doBnZPyHdZpzHmv4fyXqQ+iugt8OfKyX/Uvrlj52RRLT09b+fcoS0++R17+Uv26lspLkVvmmXar+kjvPxlWH3ailUeWtOxQ+Xhd+Sfwk1Y8nWX1C7Qv+fkfdQ+embc0suLlD+rLWf0u6UGMlU/K38o/qj7r5lh9ceWRVp36hVXXdHDN9yz19dOt+nOy6qOGhcJQeWIB5UHMd5H5NqM38eDquyne5QVXfiC/uVd5hfq+4PGqO63zu8k856y8Re/O7bPN8xxp+e/b1nefhb7dsrtV2EMLz8z/9f7zgPnVd71o9a/0zrIBudbsa4xr9iP0Tqh+ifIP9RHOKT4pz5T9WPuMs95RlF8pD1A/NQY5vYcqjqda9qvznKb3yWyzn6v42Rr8acsvqS+eH/5Nx9Rn2XEKdPWp5V/0/wnS73jrva2+9X2l1UdV3RVhxi/5teuumXcvVxzXO6Vr9lMGWHlsa3Ddn96Rqvhm3aw6Q/1CxQ/pv+LzTksPdS7qT2yy/Ir6RHovnmL58d6qw/j+yp6Zbz8O3hy+4tKTwIvZZh/7lPycb75brkTuMvztlh4p31AfLA252f7/7/f3980+wVu+aS96l2wMfNt6r1He18g134uiFE+gn3bMflKW1Xc44Zn9o6XqM1p+XvWc+iF6914YYfabdH+qaz8D/5v1PjaQcT2hT7Lqrdet+JQA1Pt7J+s9sJJnvi/IXmooLnnme+2KaTXCf713qFv4iP4LyayWxAFaTgAAAAAAAFlOAAAAAAAAAwADAAMAAAAAAAAAAAAAAAAAAAAAAAAAUxBhfj0BAACtOGciZ7YAAAAAAAA="
	txinppegdata := "AQAAAAAAAAABAgACAAAAAAAAAAAAMjc2YTkxNDgwYWNmZjE1NzJjZTcwMjg5NDVhMDVhYmFjNDc1ZGEwNGU2MzRmZTQ4OGFjzgIAAAAAAAB42u1ZMXIVMQy15CTkALlAhhuEhqEiRY7AAdLAIeACKXIMLkBDSQE1B6ChTsp0yfy/y/CzzszXrPLe86ZcNzvr9ciy9PQka0tZxzrWsY51LBlnH1/7/+c4jYNp3qZnmy/JvCVyT+/GvXEI5JmodyZXlWfi+0HY7/TV5d45s/3b95uvPqvH5Y/t7vu7MP/t77Cbr9P7sG/WJzu278eJ/kdn97v1HuRsE3+cf/o+u34zbdTm2/Pk9zi7/iHYpen9Z/P4bOuiVbL5EvZt4/bqeTy29R9O5uWj/Zrcw8S/8bwe1mXyDZwrjg2In5rIseSZ4TbO1wTPcb2H918JT6jxGf1kpP1inCD/VrA/e/5CntdJfyC+RLxnZBwVEj+s/Mw+2zd952Dt0ParSfxluBjIPHIE4noE+yG9o14Zvn5ec/oinGZ4riS+/WKY1QPxA9rPQJwO5HnUPM/qbSTPDuA8anxmvIj0yHgG5R12PcIrm3/QeZzEOarHKvATsh+bty3hezX+jfSLi3YxgMdsfwf+j/WpWqc7yadO2hHFPRvfKt8jO6j3KxN5wUFeRPIM+CXj8QLs9OUzh3c2j2Z86J11J8s3hYwjte5n8xAbf0XEXxVxx+Y7FE9sfafeK1SeXRqPKB+q9Qdbv6h5Q+U59T7J2gPxp5G8FPUZk3s5wlHGn1Eem2+dxFPmB/YewfKxijs2L1tnHCM/sfpYZ1yzuHip/qKaN9g6QuWR3vsOy0u9/IDkobgwsh6LfVU2v7D1LBtnCH8OeKqS/fsi2qcAXmXvhUNnPsv6Ayyu1b6Seq9k8yHb31HzIJpneXXpfxQDfIX8i/q4yM5V/H+0tK7yzrqXjXvWnmq+YterfS71P56KI7V/rdbbBfCb2pdm+0TeGd+sv1E/gO3PqnUoWy++1Duqm52837J1unovHJP/jibiaen/GDYeop7v3z6+/wOHt3R1ApBTAAAAAAAAj1MAAAAAAAADADAAMwA2AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJQ1dwAAAAAAAAAA"
	peglevel := "0190530000000000008f53000000000000300033003600000000000000000000000000000000000000"

	pl, err := gopegops.Decodepeglevel(peglevel)
	if err != nil {
		fmt.Println("Decodepeglevel failed: ", err)
		return
	}
	fmt.Printf("Peglevel: %+v\n", pl)
	fmt.Println("---------")

	pd, err := gopegops.Decodepegdata(balance)
	if err != nil {
		fmt.Println("Decodepegdata failed: ", err)
		return
	}
	fmt.Printf("Balance: %+v\n", pd)
	fmt.Println("---------")

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
	txinp1.TxoutId = "4790bcc0947d5861b13f2503a9b784493ecfa1969a4332d0131d2fe055450556:0"
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
