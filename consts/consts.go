package consts

import (
    "github.com/ava-labs/avalanchego/ids"
    "github.com/ava-labs/avalanchego/vms/platformvm/warp"
    "github.com/ava-labs/hypersdk/chain"
    "github.com/ava-labs/hypersdk/codec"
    "github.com/ava-labs/hypersdk/consts"
)

const (
    // Choose a human-readable part for your hyperchain (e.g., "mychain")
    HRP = "mychain"

    // Choose a name for your hyperchain (e.g., "MyCustomChain")
    Name = "MyCustomChain"

    // Choose a token symbol (e.g., "MCC")
    Symbol = "MCC"
)

var ID ids.ID

func init() {
    b := make([]byte, consts.IDLen)
    copy(b, []byte(Name))
    vmID, err := ids.ToID(b)
    if err != nil {
        panic(err)
    }
    ID = vmID
}

// Instantiate registry here so it can be imported by any package. We set these
// values in [controller/registry].
var (
    ActionRegistry *codec.TypeParser[chain.Action, *warp.Message, bool]
    AuthRegistry   *codec.TypeParser[chain.Auth, *warp.Message, bool]
)
