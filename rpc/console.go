package rpc

type console struct {
	rpc *RPC
}

type ConsoleCreateReq struct {
	Method string
	Token  string
}

type ConsoleCreateRes struct {
	ID     string `msgpack:"id" json:"id"`
	Prompt string `msgpack:"prompt" json:"prompt"`
	Busy   bool   `msgpack:"busy" json:"busy"`
}

// Create creates a new framework console instance
func (c *console) Create() (*ConsoleCreateRes, error) {
	req := &ConsoleCreateReq{
		Method: "console.create",
		Token:  c.rpc.Token(),
	}

	var res *ConsoleCreateRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleDestroyReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
}

type ConsoleDestroyRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// Destroy deletes a framework console instance
func (c *console) Destroy(consoleID string) (*ConsoleDestroyRes, error) {
	req := &ConsoleDestroyReq{
		Method:    "console.destroy",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
	}

	var res *ConsoleDestroyRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleListReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ConsoleListRes map[string][]struct {
	ID     string `msgpack:"id" json:"id"`
	Prompt string `msgpack:"prompt" json:"prompt"`
	Busy   bool   `msgpack:"busy" json:"busy"`
}

// List returns a list of framework consoles
func (c *console) List() (*ConsoleListRes, error) {
	req := &ConsoleListReq{
		Method: "console.list",
		Token:  c.rpc.Token(),
	}

	var res *ConsoleListRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleReadReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
}

type ConsoleReadRes struct {
	Data   string `msgpack:"data" json:"data"`
	Prompt string `msgpack:"prompt" json:"prompt"`
	Busy   bool   `msgpack:"busy" json:"busy"`
}

// Read returns the framework console output in raw form
func (c *console) Read(consoleID string) (*ConsoleReadRes, error) {
	req := &ConsoleReadReq{
		Method:    "console.read",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
	}

	var res *ConsoleReadRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleSessionDetachReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
}

type ConsoleSessionDetachRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// SessionDetach detaches a framework session
func (c *console) SessionDetach(consoleID string) (*ConsoleSessionDetachRes, error) {
	req := &ConsoleSessionDetachReq{
		Method:    "console.session_detach",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
	}

	var res *ConsoleSessionDetachRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleSessionKillReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
}

type ConsoleSessionKillRes struct {
	Result Result `msgpack:"result" json:"result"`
}

// SessionKill kills a framework session
func (c *console) SessionKill(consoleID string) (*ConsoleSessionKillRes, error) {
	req := &ConsoleSessionKillReq{
		Method:    "console.session_kill",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
	}

	var res *ConsoleSessionKillRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleTabsReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
	InputLine string `json:"input_line"`
}

type ConsoleTabsRes struct {
	Tabs []string `msgpack:"tabs"`
}

// Tabs returns the tab-completed version of your input (such as a module path)
func (c *console) Tabs(consoleID, line string) (*ConsoleTabsRes, error) {
	req := &ConsoleTabsReq{
		Method:    "console.tabs",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
		InputLine: line,
	}

	var res *ConsoleTabsRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ConsoleWriteReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	ConsoleID string `json:"console_id"`
	Command   string `json:"command"`
}

type ConsoleWriteRes struct {
	Wrote uint32 `msgpack:"wrote" json:"wrote"`
}

// Write sends an input (such as a command) to the framework console
func (c *console) Write(consoleID, command string) (*ConsoleWriteRes, error) {
	req := &ConsoleWriteReq{
		Method:    "console.write",
		Token:     c.rpc.Token(),
		ConsoleID: consoleID,
		Command:   command,
	}

	var res *ConsoleWriteRes
	if err := c.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
