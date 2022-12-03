// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// SignalServiceServer is an autogenerated mock type for the SignalServiceServer type
type SignalServiceServer struct {
	mock.Mock
}

type SignalServiceServer_GetOrCreateSignal struct {
	*mock.Call
}

func (_m SignalServiceServer_GetOrCreateSignal) Return(_a0 *admin.Signal, _a1 error) *SignalServiceServer_GetOrCreateSignal {
	return &SignalServiceServer_GetOrCreateSignal{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SignalServiceServer) OnGetOrCreateSignal(_a0 context.Context, _a1 *admin.SignalGetOrCreateRequest) *SignalServiceServer_GetOrCreateSignal {
	c_call := _m.On("GetOrCreateSignal", _a0, _a1)
	return &SignalServiceServer_GetOrCreateSignal{Call: c_call}
}

func (_m *SignalServiceServer) OnGetOrCreateSignalMatch(matchers ...interface{}) *SignalServiceServer_GetOrCreateSignal {
	c_call := _m.On("GetOrCreateSignal", matchers...)
	return &SignalServiceServer_GetOrCreateSignal{Call: c_call}
}

// GetOrCreateSignal provides a mock function with given fields: _a0, _a1
func (_m *SignalServiceServer) GetOrCreateSignal(_a0 context.Context, _a1 *admin.SignalGetOrCreateRequest) (*admin.Signal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *admin.Signal
	if rf, ok := ret.Get(0).(func(context.Context, *admin.SignalGetOrCreateRequest) *admin.Signal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Signal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.SignalGetOrCreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SignalServiceServer_ListSignals struct {
	*mock.Call
}

func (_m SignalServiceServer_ListSignals) Return(_a0 *admin.SignalList, _a1 error) *SignalServiceServer_ListSignals {
	return &SignalServiceServer_ListSignals{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SignalServiceServer) OnListSignals(_a0 context.Context, _a1 *admin.SignalListRequest) *SignalServiceServer_ListSignals {
	c_call := _m.On("ListSignals", _a0, _a1)
	return &SignalServiceServer_ListSignals{Call: c_call}
}

func (_m *SignalServiceServer) OnListSignalsMatch(matchers ...interface{}) *SignalServiceServer_ListSignals {
	c_call := _m.On("ListSignals", matchers...)
	return &SignalServiceServer_ListSignals{Call: c_call}
}

// ListSignals provides a mock function with given fields: _a0, _a1
func (_m *SignalServiceServer) ListSignals(_a0 context.Context, _a1 *admin.SignalListRequest) (*admin.SignalList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *admin.SignalList
	if rf, ok := ret.Get(0).(func(context.Context, *admin.SignalListRequest) *admin.SignalList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SignalList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.SignalListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SignalServiceServer_SetSignal struct {
	*mock.Call
}

func (_m SignalServiceServer_SetSignal) Return(_a0 *admin.SignalSetResponse, _a1 error) *SignalServiceServer_SetSignal {
	return &SignalServiceServer_SetSignal{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SignalServiceServer) OnSetSignal(_a0 context.Context, _a1 *admin.SignalSetRequest) *SignalServiceServer_SetSignal {
	c_call := _m.On("SetSignal", _a0, _a1)
	return &SignalServiceServer_SetSignal{Call: c_call}
}

func (_m *SignalServiceServer) OnSetSignalMatch(matchers ...interface{}) *SignalServiceServer_SetSignal {
	c_call := _m.On("SetSignal", matchers...)
	return &SignalServiceServer_SetSignal{Call: c_call}
}

// SetSignal provides a mock function with given fields: _a0, _a1
func (_m *SignalServiceServer) SetSignal(_a0 context.Context, _a1 *admin.SignalSetRequest) (*admin.SignalSetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *admin.SignalSetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.SignalSetRequest) *admin.SignalSetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SignalSetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.SignalSetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
