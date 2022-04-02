package keeper

import (
	"github.com/frank4g/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
