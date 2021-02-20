// Code generated by MockGen. DO NOT EDIT.
// Source: session.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	wire "github.com/lucas-clemente/quic-go/internal/wire"
)

// MockUnpacker is a mock of Unpacker interface.
type MockUnpacker struct {
	ctrl     *gomock.Controller
	recorder *MockUnpackerMockRecorder
}

// MockUnpackerMockRecorder is the mock recorder for MockUnpacker.
type MockUnpackerMockRecorder struct {
	mock *MockUnpacker
}

// NewMockUnpacker creates a new mock instance.
func NewMockUnpacker(ctrl *gomock.Controller) *MockUnpacker {
	mock := &MockUnpacker{ctrl: ctrl}
	mock.recorder = &MockUnpackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnpacker) EXPECT() *MockUnpackerMockRecorder {
	return m.recorder
}

// Unpack mocks base method.
func (m *MockUnpacker) Unpack(hdr *wire.Header, rcvTime time.Time, data []byte) (*unpackedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack", hdr, rcvTime, data)
	ret0, _ := ret[0].(*unpackedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unpack indicates an expected call of Unpack.
func (mr *MockUnpackerMockRecorder) Unpack(hdr, rcvTime, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockUnpacker)(nil).Unpack), hdr, rcvTime, data)
}
