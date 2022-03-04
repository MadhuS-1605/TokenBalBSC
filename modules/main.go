package modules
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	 bsctoken "TokenBalBSC/build"
)

// GetAddressBalance returns the given address balance =P
func GetAddressBalance(tokenaddress string, address string) (string, error) {
	
	client, err := ethclient.Dial("https://speedy-nodes-nyc.moralis.io/2abeb0fd5d1beb3b221a0f05/bsc/mainnet")

	if err != nil {
		return "0", err
	}

	tokenAddress := common.HexToAddress(tokenaddress)
	instance, err1 := bsctoken.NewBsctoken(tokenAddress, client)
	if err1 != nil {
		return "0", err1
	}
	
	account := common.HexToAddress(address)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return "0", err
	}
	

	return balance.String(), nil
}