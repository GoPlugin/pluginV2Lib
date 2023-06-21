package protocol

import (
	"github.com/GoPlugin/pluginV2Lib/libocr/commontypes"
	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting2plus/types"
)

// Used only for testing
type XXXUnknownMessageType[RI any] struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType[RI]) CheckSize(types.ReportingPluginLimits) bool     { return true }
func (XXXUnknownMessageType[RI]) process(*oracleState[RI], commontypes.OracleID) {}
