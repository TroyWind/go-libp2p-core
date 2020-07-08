package dlcp2plog

import (
	"github.com/libp2p/go-libp2p-core/tools/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("lcp2p")
}
