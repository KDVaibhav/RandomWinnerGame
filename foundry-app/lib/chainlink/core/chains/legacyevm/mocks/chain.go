// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	client "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"

	commontypes "github.com/smartcontractkit/chainlink/v2/common/types"

	config "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"

	context "context"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"

	gas "github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"

	log "github.com/smartcontractkit/chainlink/v2/core/chains/evm/log"

	logger "github.com/smartcontractkit/chainlink/v2/core/logger"

	logpoller "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"

	mock "github.com/stretchr/testify/mock"

	monitor "github.com/smartcontractkit/chainlink/v2/core/chains/evm/monitor"

	txmgr "github.com/smartcontractkit/chainlink/v2/common/txmgr"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

// BalanceMonitor provides a mock function with given fields:
func (_m *Chain) BalanceMonitor() monitor.BalanceMonitor {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BalanceMonitor")
	}

	var r0 monitor.BalanceMonitor
	if rf, ok := ret.Get(0).(func() monitor.BalanceMonitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monitor.BalanceMonitor)
		}
	}

	return r0
}

// Client provides a mock function with given fields:
func (_m *Chain) Client() client.Client {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Chain) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *Chain) Config() config.ChainScopedConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Config")
	}

	var r0 config.ChainScopedConfig
	if rf, ok := ret.Get(0).(func() config.ChainScopedConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ChainScopedConfig)
		}
	}

	return r0
}

// GasEstimator provides a mock function with given fields:
func (_m *Chain) GasEstimator() gas.EvmFeeEstimator {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GasEstimator")
	}

	var r0 gas.EvmFeeEstimator
	if rf, ok := ret.Get(0).(func() gas.EvmFeeEstimator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.EvmFeeEstimator)
		}
	}

	return r0
}

// GetChainStatus provides a mock function with given fields: ctx
func (_m *Chain) GetChainStatus(ctx context.Context) (types.ChainStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetChainStatus")
	}

	var r0 types.ChainStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.ChainStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.ChainStatus); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.ChainStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadBroadcaster provides a mock function with given fields:
func (_m *Chain) HeadBroadcaster() commontypes.HeadBroadcaster[*evmtypes.Head, common.Hash] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HeadBroadcaster")
	}

	var r0 commontypes.HeadBroadcaster[*evmtypes.Head, common.Hash]
	if rf, ok := ret.Get(0).(func() commontypes.HeadBroadcaster[*evmtypes.Head, common.Hash]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(commontypes.HeadBroadcaster[*evmtypes.Head, common.Hash])
		}
	}

	return r0
}

// HeadTracker provides a mock function with given fields:
func (_m *Chain) HeadTracker() commontypes.HeadTracker[*evmtypes.Head, common.Hash] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HeadTracker")
	}

	var r0 commontypes.HeadTracker[*evmtypes.Head, common.Hash]
	if rf, ok := ret.Get(0).(func() commontypes.HeadTracker[*evmtypes.Head, common.Hash]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(commontypes.HeadTracker[*evmtypes.Head, common.Hash])
		}
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *Chain) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Chain) ID() *big.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// ListNodeStatuses provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *Chain) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) ([]types.NodeStatus, string, int, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListNodeStatuses")
	}

	var r0 []types.NodeStatus
	var r1 string
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) []types.NodeStatus); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NodeStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, string) string); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int32, string) int); ok {
		r2 = rf(ctx, pageSize, pageToken)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(context.Context, int32, string) error); ok {
		r3 = rf(ctx, pageSize, pageToken)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// LogBroadcaster provides a mock function with given fields:
func (_m *Chain) LogBroadcaster() log.Broadcaster {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LogBroadcaster")
	}

	var r0 log.Broadcaster
	if rf, ok := ret.Get(0).(func() log.Broadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Broadcaster)
		}
	}

	return r0
}

// LogPoller provides a mock function with given fields:
func (_m *Chain) LogPoller() logpoller.LogPoller {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LogPoller")
	}

	var r0 logpoller.LogPoller
	if rf, ok := ret.Get(0).(func() logpoller.LogPoller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logpoller.LogPoller)
		}
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *Chain) Logger() logger.Logger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Chain) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *Chain) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *Chain) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transact provides a mock function with given fields: ctx, from, to, amount, balanceCheck
func (_m *Chain) Transact(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool) error {
	ret := _m.Called(ctx, from, to, amount, balanceCheck)

	if len(ret) == 0 {
		panic("no return value specified for Transact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int, bool) error); ok {
		r0 = rf(ctx, from, to, amount, balanceCheck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxManager provides a mock function with given fields:
func (_m *Chain) TxManager() txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxManager")
	}

	var r0 txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]
	if rf, ok := ret.Get(0).(func() txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(txmgr.TxManager[*big.Int, *evmtypes.Head, common.Address, common.Hash, common.Hash, evmtypes.Nonce, gas.EvmFee])
		}
	}

	return r0
}

// NewChain creates a new instance of Chain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChain(t interface {
	mock.TestingT
	Cleanup(func())
}) *Chain {
	mock := &Chain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
