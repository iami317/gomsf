package rpc

type plugin struct {
	rpc *RPC
}

type PluginLoadReq struct {
	Method     string            `json:"method"`
	Token      string            `json:"token"`
	PluginName string            `json:"plugin_name"`
	Options    map[string]string `json:"options"`
}

type PluginLoadRes struct {
	Result Result `msgpack:"result"`
}

// Load loads a plugin
func (p *plugin) Load(pluginName string, pluginOptions map[string]string) (*PluginLoadRes, error) {
	req := &PluginLoadReq{
		Method:     "plugin.load",
		Token:      p.rpc.Token(),
		PluginName: pluginName,
		Options:    pluginOptions,
	}

	var res *PluginLoadRes
	if err := p.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type PluginLoadedReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type PluginLoadedRes struct {
	Plugins []string `msgpack:"plugins" json:"plugins"`
}

// Loaded returns a list of loaded plugins
func (p *plugin) Loaded() (*PluginLoadedRes, error) {
	req := &PluginLoadedReq{
		Method: "plugin.loaded",
		Token:  p.rpc.Token(),
	}

	var res *PluginLoadedRes
	if err := p.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type PluginUnLoadReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	PluginName string `json:"plugin_name"`
}

type PluginUnLoadRes struct {
	Result Result `msgpack:"result"`
}

// Unload unloads a plugin
func (p *plugin) UnLoad(pluginName string) (*PluginUnLoadRes, error) {
	req := &PluginUnLoadReq{
		Method:     "plugin.unload",
		Token:      p.rpc.Token(),
		PluginName: pluginName,
	}

	var res *PluginUnLoadRes
	if err := p.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
