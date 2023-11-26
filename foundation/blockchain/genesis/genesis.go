// Package genesis maintains access to the genesis file.
package genesis

import (
	"encoding/json"
	"os"
	"time"
)

// Genesis represents the genesis file.
type Genesis struct {
	Date          time.Time         `json:"date"`
	ChainID       uint16            `json:"chain_id"`        // ChainID represents an unique id for the running instance.
	TransPerBlock uint16            `json:"trans_per_block"` // TransPerBlock maximum number of transaction that can be in a block.
	Difficulty    uint16            `json:"difficulty"`      // Difficulty it needs to be able to solve the work problem.
	MiningReward  uint64            `json:"mining_reward"`   // MiningReward for mining a block.
	GasPrice      uint64            `json:"gas_price"`       // GasPrice Fee paid for each transaction mined in a block.
	Balances      map[string]uint64 `json:"balances"`
}

// =============================================================================================

// Load opens an consumes the genesis file.
func Load() (Genesis, error) {
	path := "zblock/genesis.json"
	content, err := os.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var genesis Genesis
	err = json.Unmarshal(content, &genesis)
	if err != nil {
		return Genesis{}, err
	}
	return genesis, nil
}
