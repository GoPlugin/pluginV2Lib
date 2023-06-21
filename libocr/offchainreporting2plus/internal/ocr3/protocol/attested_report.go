package protocol

import (
	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting2plus/ocr3types"
	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting2plus/types"
)

type AttestedReportMany[RI any] struct {
	ReportWithInfo       ocr3types.ReportWithInfo[RI]
	AttributedSignatures []types.AttributedOnchainSignature
}
