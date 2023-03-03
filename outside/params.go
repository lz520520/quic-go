package outside

import (
	"github.com/quic-go/quic-go/internal/congestion"
	"github.com/quic-go/quic-go/internal/protocol"
)

func SetMinInitialPacketSize(size int) {
	protocol.MinInitialPacketSize = protocol.ByteCount(size)
}

// 设置最大包大小，默认MTU是按照1500计算的
func SetMaxPacketBufferSize(size int) {
	protocol.MaxPacketBufferSize = protocol.ByteCount(size)
}

func SetInitialPacketSizeIPv4(size int) {
	protocol.InitialPacketSizeIPv4 = protocol.ByteCount(size)
}

func SetInitialPacketSizeIPv6(size int) {
	protocol.InitialPacketSizeIPv6 = protocol.ByteCount(size)
}

func RefreshRelatedParams() {
	congestion.RefreshInitialMaxDatagramSize()
	congestion.RefreshMaxDatagramSize()
}
