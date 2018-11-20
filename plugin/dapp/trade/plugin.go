// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trade

import (
	"github.com/33cn/chain33/pluginmgr"
	_ "github.com/33cn/plugin/plugin/dapp/trade/autotest"
	"github.com/33cn/plugin/plugin/dapp/trade/commands"
	"github.com/33cn/plugin/plugin/dapp/trade/executor"
	"github.com/33cn/plugin/plugin/dapp/trade/rpc"
	"github.com/33cn/plugin/plugin/dapp/trade/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TradeX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TradeCmd,
		RPC:      rpc.Init,
	})
}
