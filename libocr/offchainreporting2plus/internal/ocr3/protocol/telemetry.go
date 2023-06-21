package protocol

import (
	"github.com/GoPlugin/pluginV2Lib/libocr/commontypes"
	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting2plus/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint64,
		round uint8,
		leader commontypes.OracleID,
	)
}
