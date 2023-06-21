package protocol

import "github.com/GoPlugin/pluginV2Lib/libocr/commontypes"

// Used only for testing
type XXXUnknownMessageType struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType) process(*oracleState, commontypes.OracleID) {}
