package onchain_test

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/KyberNetwork/go-project-template/pkg/onchain/multiutil"
	"github.com/KyberNetwork/go-project-template/pkg/onchain/simclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMultiUtil(t *testing.T) {
	t.Skip()
	_ = godotenv.Load("../../.env")

	multiContract := common.HexToAddress("0x1111111111111111111111111111111111111100")
	binCodeIdx := strings.Index(multiutil.MultiUtilMetaData.Bin, "f3fe") // RET,INVALID follow by runtime code
	require.Greater(t, binCodeIdx, -1)
	sclient, err := simclient.NewSimClient(os.Getenv("MAINNET_RPC"), http.DefaultClient, simclient.OverrideAccounts{
		multiContract: {
			Code: "0x" + multiutil.MultiUtilMetaData.Bin[binCodeIdx+4:],
		},
	})
	require.NoError(t, err)
	mclient, err := multiutil.NewMultiUtil(multiContract, sclient)
	require.NoError(t, err)
	balances, err := mclient.BalanceOfMultiTokens(&bind.CallOpts{},
		common.HexToAddress("0xBC33a1F908612640F2849b56b67a4De4d179C151"),
		[]common.Address{
			common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
			common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202"),
		})
	require.NoError(t, err)
	t.Log("USDT balance", balances[0].String())
	t.Log("KNC balance", balances[1].String())
}
