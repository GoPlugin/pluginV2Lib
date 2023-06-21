package protocol

import (
	"context"

	"github.com/GoPlugin/pluginV2Lib/libocr/offchainreporting2plus/types"
)

type Database interface {
	types.ConfigDatabase

	ReadCert(ctx context.Context, configDigest types.ConfigDigest) (CertifiedPrepareOrCommit, error)
	WriteCert(ctx context.Context, configDigest types.ConfigDigest, cert CertifiedPrepareOrCommit) error
}
