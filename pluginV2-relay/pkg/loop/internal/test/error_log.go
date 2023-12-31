package test

import (
	"context"
	"fmt"

	"github.com/GoPlugin/pluginV2Lib/pluginV2-relay/pkg/loop/internal"
)

var _ internal.ErrorLog = (*StaticErrorLog)(nil)

type StaticErrorLog struct{}

func (s *StaticErrorLog) SaveError(ctx context.Context, msg string) error {
	if msg != errMsg {
		return fmt.Errorf("expected %q but got %q", errMsg, msg)
	}
	return nil
}
