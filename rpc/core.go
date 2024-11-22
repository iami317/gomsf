package rpc

import (
	"fmt"
	"reflect"
)

type core struct {
	rpc *RPC
}

type CoreAddModulePathReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
	Path   string `json:"path"`
}

type CoreAddModulePathRes struct {
	Exploits  uint32 `msgpack:"exploits" json:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary" json:"auxiliary"`
	Post      uint32 `msgpack:"post" json:"post"`
	Encoders  uint32 `msgpack:"encoders" json:"encoders"`
	Nops      uint32 `msgpack:"nops" json:"nops"`
	Payloads  uint32 `msgpack:"payloads" json:"payloads"`
}

// AddModulePath adds a new local file system path (local to the server) as a module path
func (c *core) AddModulePath(path string) (*CoreAddModulePathRes, error) {
	req := &CoreAddModulePathReq{
		Method: "core.add_module_path",
		Token:  c.rpc.Token(),
		Path:   path,
	}

	var res *CoreAddModulePathRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreGetgReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	OptionName string `json:"option_name"`
}

type CoreGetgRes struct {
	Result string `json:"result"`
}

// Getg returns a global datastore option
func (c *core) Getg(optionName string) (*CoreGetgRes, error) {
	req := &CoreGetgReq{
		Method:     "core.getg",
		Token:      c.rpc.Token(),
		OptionName: optionName,
	}

	var res interface{}
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return &CoreGetgRes{Result: reflect.ValueOf(res).MapIndex(reflect.ValueOf(optionName)).String()}, nil
}

type CoreModuleStatsReq struct {
	Method string
	Token  string
}

type CoreModuleStatsRes struct {
	Exploits  uint32 `msgpack:"exploits" json:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary" json:"auxiliary"`
	Post      uint32 `msgpack:"post" json:"post"`
	Encoders  uint32 `msgpack:"encoders" json:"encoders"`
	Nops      uint32 `msgpack:"nops" json:"nops"`
	Payloads  uint32 `msgpack:"payloads" json:"payloads"`
}

// ModuleStats returns the module stats
func (c *core) ModuleStats() (*CoreModuleStatsRes, error) {
	req := &CoreModuleStatsReq{
		Method: "core.module_stats",
		Token:  c.rpc.Token(),
	}

	var res *CoreModuleStatsRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreReloadModulesReq struct {
	Method string
	Token  string
}

type CoreReloadModulesRes struct {
	Exploits  uint32 `msgpack:"exploits" json:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary" json:"auxiliary"`
	Post      uint32 `msgpack:"post" json:"post"`
	Encoders  uint32 `msgpack:"encoders" json:"encoders"`
	Nops      uint32 `msgpack:"nops" json:"nops"`
	Payloads  uint32 `msgpack:"payloads" json:"payloads"`
}

// ReloadModules reloads framework modules
func (c *core) ReloadModules() (*CoreReloadModulesRes, error) {
	req := &CoreReloadModulesReq{
		Method: "core.reload_modules",
		Token:  c.rpc.Token(),
	}

	var res *CoreReloadModulesRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreSaveReq struct {
	Method string
	Token  string
}

type CoreSaveRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// Save saves current framework settings
func (c *core) Save() (*CoreSaveRes, error) {
	req := &CoreSaveReq{
		Method: "core.save",
		Token:  c.rpc.Token(),
	}

	var res *CoreSaveRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreSetgReq struct {
	Method      string `json:"method"`
	Token       string `json:"token"`
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
}

type CoreSetgRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// Setg sets a global datastore option
func (c *core) Setg(optionName, optionValue string) (*CoreSetgRes, error) {
	req := &CoreSetgReq{
		Method:      "core.setg",
		Token:       c.rpc.Token(),
		OptionName:  optionName,
		OptionValue: optionValue,
	}

	var res *CoreSetgRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreStopReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type CoreStopRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// Stop stops the RPC service
func (c *core) Stop() (*CoreStopRes, error) {
	req := &CoreStopReq{
		Method: "core.stop",
		Token:  c.rpc.Token(),
	}

	var res *CoreStopRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreThreadKillReq struct {
	Method   string `json:"method"`
	Token    string `json:"token"`
	ThreadID string `json:"thread_id"`
}

type CoreThreadKillRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// ThreadKill kills a framework thread
func (c *core) ThreadKill(threadID string) (*CoreThreadKillRes, error) {
	req := &CoreThreadKillReq{
		Method:   "core.thread_kill",
		Token:    c.rpc.Token(),
		ThreadID: threadID,
	}

	var res *CoreThreadKillRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreThreadListReq struct {
	Method string
	Token  string
}

type CoreThreadListRes map[int]struct {
	Status   string `msgpack:"status" json:"status"`
	Critical bool   `msgpack:"critical" json:"critical"`
	Name     string `msgpack:"name" json:"name"`
	Started  string `msgpack:"started" json:"started"`
}

// ThreadList returns a list of framework threads
func (c *core) ThreadList() (*CoreThreadListRes, error) {
	req := &CoreThreadListReq{
		Method: "core.thread_list",
		Token:  c.rpc.Token(),
	}

	var res *CoreThreadListRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreUnsetgReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	OptionName string `json:"option_name"`
}

type CoreUnsetgRes struct {
	Result Result `msgpack:"result"`
}

// Unsetg unsets a global datastore option
func (c *core) Unsetg(optionName string) (*CoreUnsetgRes, error) {
	req := &CoreUnsetgReq{
		Method:     "core.unsetg",
		Token:      c.rpc.Token(),
		OptionName: optionName,
	}

	var res *CoreUnsetgRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type CoreVersionReq struct {
	Method string
	Token  string
}

type CoreVersionRes struct {
	Version string `msgpack:"version" json:"version"` // Framework version
	Ruby    string `msgpack:"ruby" json:"ruby"`       // Ruby version
	API     string `msgpack:"api" json:"api"`         // API version
}

func (version *CoreVersionRes) String() string {
	return fmt.Sprintf("version:%v ruby:%v api:%v", version.Version, version.Ruby, version.API)
}

// Version returns the RPC service versions
func (c *core) Version() (*CoreVersionRes, error) {
	req := &CoreVersionReq{
		Method: "core.version",
		Token:  c.rpc.Token(),
	}

	var res *CoreVersionRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
