package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Models "TokenBalBSC/models"
	Modules "TokenBalBSC/modules"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get parameter from url request
	vars := mux.Vars(r)
	module := vars["module"]

	// Get the query parameters from url request
	tokenaddress := r.URL.Query().Get("tokenaddress")
	address := r.URL.Query().Get("address")
	name := r.URL.Query().Get("name")
	symbol := r.URL.Query().Get("symbol")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")
	switch module {

	case "get-balance":
		
		if tokenaddress == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		if address == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		balance, err := Modules.GetAddressBalance(tokenaddress, address)
		//name, err := Modules.Name(&_Bsctoken.CallOpts)
		//symbol, err := Modules.Symbol(&_Bsctoken.CallOpts)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&Models.BalanceResponse{
			Name: name,
			Symbol: symbol,
			Address: address,
			Balance: balance,
		})


	}

}