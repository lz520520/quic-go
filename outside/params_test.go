package outside

import (
	"testing"
)

func TestSet(t *testing.T) {
	SetMinInitialPacketSize(512)
	SetInitialPacketSizeIPv4(512 + 52)
	SetInitialPacketSizeIPv6(512 + 32)
	SetMaxPacketBufferSize(800)

	RefreshRelatedParams()
}
