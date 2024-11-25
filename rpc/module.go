package rpc

import (
	"encoding/json"
	"fmt"
)

type module struct {
	rpc *RPC
}
type ModuleArchitecturesReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleArchitecturesRes []string

// Architectures returns a list of architecture names
func (m *module) Architectures() (*ModuleArchitecturesRes, error) {
	req := &ModuleArchitecturesReq{
		Method: "module.architectures",
		Token:  m.rpc.Token(),
	}

	var res *ModuleArchitecturesRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleAuxiliaryReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleAuxiliaryRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

// Auxiliary 返回一个辅助模块名称列表
func (m *module) Auxiliary() (*ModuleAuxiliaryRes, error) {
	req := &ModuleAuxiliaryReq{
		Method: "module.auxiliary",
		Token:  m.rpc.Token(),
	}

	var res *ModuleAuxiliaryRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleCheckReq struct {
	Method     string            `json:"method"`
	Token      string            `json:"token"`
	ModuleType string            `json:"module_type"`
	ModuleName string            `json:"module_name"`
	Options    map[string]string `json:"options"`
}

func (m *module) Check(moduleType, moduleName string, options map[string]string) error {
	req := &ModuleCheckReq{
		Method:     "module.execute",
		Token:      m.rpc.Token(),
		ModuleType: moduleType,
		ModuleName: moduleName,
		Options:    options,
	}

	if err := m.rpc.Call(req, nil); err != nil {
		return err
	}

	return nil
}

type ModuleExploitsReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleExploitsRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

func (m *module) Exploits() (*ModuleExploitsRes, error) {
	req := &ModuleExploitsReq{
		Method: "module.exploits",
		Token:  m.rpc.Token(),
	}

	var res *ModuleExploitsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleCompatibleEvasionPayloadsReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleName string `json:"module_name"`
	Target     int    `json:"target"`
}

type ModuleCompatibleEvasionPayloadsRes struct {
	Payloads []string `msgpack:"payloads" json:"payloads"`
}

// CompatibleEvasionPayloads returns the compatible target-specific payloads for an evasion module
func (m *module) CompatibleEvasionPayloads(moduleName string, target int) (*ModuleCompatibleEvasionPayloadsRes, error) {
	req := &ModuleCompatibleEvasionPayloadsReq{
		Method:     "module.compatible_evasion_payloads",
		Token:      m.rpc.Token(),
		ModuleName: moduleName,
		Target:     target,
	}

	var res *ModuleCompatibleEvasionPayloadsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleCompatiblePayloadsReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleName string `json:"module_name"`
	Target     int    `json:"target"`
}

type ModuleCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads" json:"payloads"`
}

// CompatiblePayloads returns the compatible payloads for a specific exploit
func (m *module) CompatiblePayloads(moduleName string, target int) (*ModuleCompatiblePayloadsRes, error) {
	req := &ModuleCompatiblePayloadsReq{
		Method:     "module.compatible_payloads",
		Token:      m.rpc.Token(),
		ModuleName: moduleName,
		Target:     target,
	}

	var res *ModuleCompatiblePayloadsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleCompatibleSessionsReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleName string `json:"module_name"`
}

type ModuleCompatibleSessionsRes struct {
	Sessions []string `msgpack:"sessions" json:"sessions"`
}

// ModuleCompatibleSessions returns the compatible sessions for a specific post module
func (m *module) CompatibleSessions(moduleName string) (*ModuleCompatibleSessionsRes, error) {
	req := &ModuleCompatibleSessionsReq{
		Method:     "module.compatible_sessions",
		Token:      m.rpc.Token(),
		ModuleName: moduleName,
	}

	var res *ModuleCompatibleSessionsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleEncodeReq struct {
	Method        string            `json:"method"`
	Token         string            `json:"token"`
	Data          string            `json:"data"`
	EncoderModule string            `json:"encoder_module"`
	Options       map[string]string `json:"options"`
}

type ModuleEncodeRes struct {
	Encoded []byte `msgpack:"encoded" json:"encoded"`
}

// Encode encodes data with an encoder
func (m *module) Encode(data, encoderModule string, options map[string]string) (*ModuleEncodeRes, error) {
	req := &ModuleEncodeReq{
		Method:        "module.encode",
		Token:         m.rpc.Token(),
		Data:          data,
		EncoderModule: encoderModule,
		Options:       options,
	}

	var res *ModuleEncodeRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModulePostReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModulePostRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

func (m *module) Post() (*ModulePostRes, error) {
	req := &ModulePostReq{
		Method: "module.post",
		Token:  m.rpc.Token(),
	}

	var res *ModulePostRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModulePayloadsReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModulePayloadsRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

func (m *module) Payloads() (*ModulePayloadsRes, error) {
	req := &ModulePayloadsReq{
		Method: "module.payloads",
		Token:  m.rpc.Token(),
	}

	var res *ModulePayloadsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModulePlatformsReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModulePlatformsRes []string

// Platforms 返回一个平台名称列表
func (m *module) Platforms() (*ModulePlatformsRes, error) {
	req := &ModulePlatformsReq{
		Method: "module.platforms",
		Token:  m.rpc.Token(),
	}

	var res *ModulePlatformsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleEncodersReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleEncodersRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

func (m *module) Encoders() (*ModuleEncodersRes, error) {
	req := &ModuleEncodersReq{
		Method: "module.encoders",
		Token:  m.rpc.Token(),
	}

	var res *ModuleEncodersRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleEncryptionFormatsReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleEncryptionFormatsRes []string

func (m *module) EncryptionFormats() (*ModuleEncryptionFormatsRes, error) {
	req := &ModuleEncryptionFormatsReq{
		Method: "module.encryption_formats",
		Token:  m.rpc.Token(),
	}

	var res *ModuleEncryptionFormatsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleEvasionReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleEvasionRes struct {
	Modules []string `msgpack:"modules"`
}

// Evasion 反杀毒软件的木马
func (m *module) Evasion() (*ModuleEvasionRes, error) {
	req := &ModuleEvasionReq{
		Method: "module.evasion",
		Token:  m.rpc.Token(),
	}

	var res *ModuleEvasionRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleNopsReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type ModuleNopsRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

// Nops 空指令模块
func (m *module) Nops() (*ModuleNopsRes, error) {
	req := &ModuleNopsReq{
		Method: "module.nops",
		Token:  m.rpc.Token(),
	}

	var res *ModuleNopsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleInfoReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleType string `json:"module_type"`
	ModuleName string `json:"module_name"`
}

type ModuleInfoRes struct {
	Name        string     `msgpack:"name" json:"name"`
	Description string     `msgpack:"description" json:"description"`
	License     string     `msgpack:"license" json:"license"`
	FilePath    string     `msgpack:"filepath" json:"filePath"`
	Version     string     `msgpack:"version" json:"version"`
	Rank        string     `msgpack:"rank" json:"rank"`
	References  [][]string `msgpack:"references" json:"references"`
	Authors     []string   `msgpack:"authors" json:"authors"`
	//Privileged    bool       `msgpack:"privileged" json:"privileged"`
	//DefaultTarget int        `msgpack:"default_target" json:"default_target"`
	//Arch           string     `msgpack:"arch" json:"arch"`
	DisclosureDate string   `msgpack:"disclosure_date" json:"disclosure_date"`
	Platform       []string `msgpack:"platform" json:"platform"`
}

func (mir ModuleInfoRes) String() string {
	bb, _ := json.Marshal(mir)
	return string(bb)
}

// Info 返回模块的元数据
func (m *module) Info(moduleType, moduleName string) (*ModuleInfoRes, error) {
	req := &ModuleInfoReq{
		Method:     "module.info",
		Token:      m.rpc.Token(),
		ModuleType: moduleType,
		ModuleName: moduleName,
	}
	var res *ModuleInfoRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleInfoHTMLReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleType string `json:"module_type"`
	ModuleName string `json:"module_name"`
}

type ModuleInfoHTMLRes string

// InfoHTML 返回关于HTML中模块的详细信息
func (m *module) InfoHTML(moduleType, moduleName string) (*ModuleInfoHTMLRes, error) {
	req := &ModuleInfoHTMLReq{
		Method:     "module.info_html",
		Token:      m.rpc.Token(),
		ModuleType: moduleType,
		ModuleName: moduleName,
	}

	var res *ModuleInfoHTMLRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleOptionsReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleType string `json:"module_type"`
	ModuleName string `json:"module_name"`
}

type ModuleOptionsRes map[string]struct {
	Type     string      `msgpack:"type" json:"type"`
	Required bool        `msgpack:"required" json:"required"`
	Advanced bool        `msgpack:"advanced" json:"advanced"`
	Evasion  bool        `msgpack:"evasion" json:"evasion"`
	Desc     string      `msgpack:"desc" json:"desc"`
	Default  interface{} `msgpack:"default" json:"default"`
	Enums    []string    `msgpack:"enums,omitempty" json:"enums,omitempty"`
}

func (m *module) Options(moduleType, moduleName string) (*ModuleOptionsRes, error) {
	req := &ModuleOptionsReq{
		Method:     "module.options",
		Token:      m.rpc.Token(),
		ModuleType: moduleType,
		ModuleName: moduleName,
	}

	var res *ModuleOptionsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleTargetCompatiblePayloadsReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	ModuleName string `json:"module_name"`
	ArchNumber uint32 `json:"arch_number"`
}

type ModuleTargetCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads" json:"payloads"`
}

func (m *module) TargetCompatiblePayloads(moduleName string, targetNumber uint32) (*ModuleTargetCompatiblePayloadsRes, error) {
	req := &ModuleTargetCompatiblePayloadsReq{
		Method:     "module.target_compatible_payloads",
		Token:      m.rpc.Token(),
		ModuleName: moduleName,
		ArchNumber: targetNumber,
	}

	var res *ModuleTargetCompatiblePayloadsRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ModuleExecuteReq struct {
	Method     string            `json:"method"`
	Token      string            `json:"token"`
	ModuleType string            `json:"module_type"`
	ModuleName string            `json:"module_name"`
	Options    map[string]string `json:"options"`
}

type ModuleExecuteRes struct {
	JobID uint32 `msgpack:"job_id" json:"job_id"`
	UUID  string `msgpack:"uuid" json:"uuid"`
}

func (mer ModuleExecuteRes) String() string {
	return fmt.Sprintf("job_id:%v uuid:%v", mer.JobID, mer.UUID)
}

/*
//moduleType exploit auxiliary post payload evasion
//moduleName windows/smb/ms08_067_netapi
*/
func (m *module) Execute(moduleType, moduleName string, options map[string]string) (*ModuleExecuteRes, error) {
	req := &ModuleExecuteReq{
		Method:     "module.execute",
		Token:      m.rpc.Token(),
		ModuleType: moduleType,
		ModuleName: moduleName,
		Options:    options,
	}

	var res *ModuleExecuteRes
	if err := m.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
