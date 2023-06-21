package protocol

import (
	"github.com/GoPlugin/pluginV2Lib/libocr/commontypes"
	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint32,
		round uint8,
		leader commontypes.OracleID,
	)
}
