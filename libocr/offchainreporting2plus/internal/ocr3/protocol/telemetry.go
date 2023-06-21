package protocol

import (
	"github.com/Goplugin/PluginV2Lib/libocr/commontypes"
	"github.com/Goplugin/PluginV2Lib/libocr/offchainreporting2plus/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint64,
		round uint8,
		leader commontypes.OracleID,
	)
}
