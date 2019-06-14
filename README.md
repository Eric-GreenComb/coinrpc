# CoinRPC

## UsdtAPI
实现一个简单的omni核心RPC接口。 
用于帮助USDT转移和监控地址记帐记录。

Examples:

```golang
package main

import (
	"github.com/fanguanghui/coinrpc/usdt"
	"github.com/fanguanghui/coinrpc/rpc"
	"log"
)

var (
	connCfg = &rpc.ConnConfig{
		Host: "localhost:19031",
		User: "admin",
		Pass: "123456",
	}
)

func main() {
    	omni := usdt.NewOmniClient(connCfg)
    
    	// Bitcoin手续费地址
    	// mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ
    
    	//mfysfuTB8QcnCnbR65Y91ycZS4RLbg1TGF
    	//mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME
    	//mwuavEM8Xmu1nkCmgku4iPbrpKR4L5jM4W
    
    	//address := omni.GetNewAddress("")
    	//log.Printf("获取钱包地址：%v\n", address)
    
    	index := omni.GetBlockCount()
    	log.Printf("区块高度：%v\n", index)
    
    	b, r := omni.GetBalance("mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ", 2)
    	log.Printf("查询余额：%s, %s\n", b, r)
    
    	b1, r1 := omni.GetBalance("mfysfuTB8QcnCnbR65Y91ycZS4RLbg1TGF", 2)
    	log.Printf("查询余额：%s, %s\n", b1, r1)
    
    	b2, r2 := omni.GetBalance("mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME", 2)
    	log.Printf("查询余额：%s, %s\n", b2, r2)
    
    
    	txs := omni.ListTransactions()
    	log.Printf("交易列表：%v\n", txs)
    
    	ids := omni.GetBlockTransactions(index)
    	log.Printf("按区块查交易：%v\n", ids)
    
    	tx := omni.GetTransaction("9512b20d9dc88c385328ce4a4fac0d6dfad7150a30649e28a27a7c9b4ab26c49")
    	log.Printf("按Txid查询交易：%v\n", tx)
    
    	ptxs := omni.GetPendingTransactions("")
    	log.Printf("零确认交易列表：%v\n", ptxs)
    
    
    	// fromaddress需有Bitcion做手续费
    	h ,err:= omni.Send("mfysfuTB8QcnCnbR65Y91ycZS4RLbg1TGF", "mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME", 2, "0.001")
    	if err !=nil{
    		fmt.Println(err)
    	}
    	log.Printf("固定数量交易：%v\n", h)
    
    
    	// v0.3.1开始可以使用
    	h1 ,err1:= omni.FundedSend("mfysfuTB8QcnCnbR65Y91ycZS4RLbg1TGF", "mwuavEM8Xmu1nkCmgku4iPbrpKR4L5jM4W", 2, "0.001","mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ")
    	if err1 !=nil{
    		fmt.Println(err1)
    	}
    	log.Printf("固定数量交易：%v\n", h1)
    	
    	h2 ,err2:= omni.FundedSendall("mfysfuTB8QcnCnbR65Y91ycZS4RLbg1TGF", "mipvXNuvsoD9JvxCiG6MeKB4noyp3e88ME", 2, "mt8Lvqmik6w4ZimnqDb8pKUupYyXFpjBQJ")
    	if err2 !=nil{
    		fmt.Println(err2)
    	}
    	log.Printf("所有数量交易：%v\n", h2)

}
```