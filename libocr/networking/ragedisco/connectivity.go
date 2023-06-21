package ragedisco

import ragetypes "github.com/GoPlugin/pluginV2Lib/libocr/ragep2p/types"

type connectivityMsgType int

const (
	_ connectivityMsgType = iota
	connectivityAdd
	connectivityRemove
)

type connectivityMsg struct {
	msgType connectivityMsgType
	peerID  ragetypes.PeerID
}
