package rpc

import "encoding/json"

type session struct {
	rpc *RPC
}

type SessionCompatibleModulesReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
}

type SessionCompatibleModulesRes struct {
	Modules []string `msgpack:"modules" json:"modules"`
}

func (s *session) CompatibleModules(session int) (SessionCompatibleModulesRes, error) {
	req := &SessionCompatibleModulesReq{
		Method:    "session.compatible_modules",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionCompatibleModulesRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionCompatibleModulesRes{}, err
	}

	return res, nil
}

type SessionMeterpreterWriteReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
	Command   string `json:"command"`
}

type SessionMeterpreterWriteRes struct {
	Result Result `msgpack:"result"`
}

type SessionMeterpreterReadReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
}

type SessionMeterpreterReadRes struct {
	Data string `msgpack:"data" json:"data"`
}

type SessionMeterpreterRunSingleReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
	Command   string `json:"command"`
}

type SessionMeterpreterRunSingleRes SessionMeterpreterWriteRes

type SessionMeterpreterDetachReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
}

type SessionMeterpreterDetachRes SessionMeterpreterWriteRes

type SessionMeterpreterKillReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
}

type SessionMeterpreterKillRes SessionMeterpreterWriteRes

type SessionShellUpgradeReq struct {
	Method     string `json:"method"`
	Token      string `json:"token"`
	SessionID  int    `json:"session_id"`
	IPAddress  string `json:"ip_address"`
	PortNumber uint32 `json:"port_number"`
}

type SessionShellUpgradeRes SessionMeterpreterWriteRes

type SessionRingClearReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
}

type SessionRingClearRes SessionMeterpreterWriteRes

type SessionRingPutReq struct {
	Method    string `json:"method"`
	Token     string `json:"token"`
	SessionID int    `json:"session_id"`
	Command   string `json:"command"`
}

type SessionRingPutRes struct {
	WriteCount uint32 `msgpack:"write_count" json:"write_count"`
}

type SessionListReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type SessionListRes map[uint32]struct {
	Type        string `msgpack:"type" json:"type"`
	TunnelLocal string `msgpack:"tunnel_local" json:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer" json:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit" json:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload" json:"via_payload"`
	Description string `msgpack:"desc" json:"desc"`
	Info        string `msgpack:"info" json:"info"`
	Workspace   string `msgpack:"workspace" json:"workspace"`
	SessionHost string `msgpack:"session_host" json:"session_host"`
	SessionPort int    `msgpack:"session_port" json:"session_port"`
	Username    string `msgpack:"username" json:"username"`
	UUID        string `msgpack:"uuid" json:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid" json:"exploit_uuid"`
}

func (slr SessionListRes) String() string {
	bb, _ := json.Marshal(slr)
	return string(bb)
}

func (s *session) List() (SessionListRes, error) {
	req := &SessionListReq{
		Method: "session.list",
		Token:  s.rpc.Token(),
	}

	var res SessionListRes
	if err := s.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type SessionStopRes struct {
	Result Result `msgpack:"result"`
}

type SessionStopReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
	Sid    int    `json:"sid"`
}

func (s *session) Stop(sid int) (string, error) {
	req := &SessionStopReq{
		Method: "session.stop",
		Token:  s.rpc.Token(),
		Sid:    sid,
	}

	var res SessionStopRes
	if err := s.rpc.Call(req, &res); err != nil {
		return "", err
	}

	return string(res.Result), nil
}

type SessionShellWriteReq struct {
	Method    string
	Token     string
	SessionID int
	Command   string
}

type SessionShellWriteRes struct {
	WriteCount string `msgpack:"write_count"`
}

func (s *session) ShellWrite(session int, command string) error {
	req := &SessionShellWriteReq{
		Method:    "session.shell_write",
		Token:     s.rpc.Token(),
		SessionID: session,
		Command:   command,
	}

	var res SessionShellWriteRes
	if err := s.rpc.Call(req, &res); err != nil {
		return err
	}

	return nil
}

type SessionShellReadReq struct {
	Method      string
	Token       string
	SessionID   int
	ReadPointer uint32
}

type SessionShellReadRes struct {
	Seq  uint32 `msgpack:"seq"`
	Data string `msgpack:"data"`
}

func (s *session) ShellRead(session int, readPointer uint32) (string, error) {
	req := &SessionShellReadReq{
		Method:      "session.shell_read",
		Token:       s.rpc.Token(),
		SessionID:   session,
		ReadPointer: readPointer,
	}

	var res SessionShellReadRes
	if err := s.rpc.Call(req, &res); err != nil {
		return "", err
	}

	return res.Data, nil
}

func (s *session) MeterpreterWrite(session int, command string) (SessionMeterpreterWriteRes, error) {
	req := &SessionMeterpreterWriteReq{
		Method:    "session.meterpreter_write",
		Token:     s.rpc.Token(),
		SessionID: session,
		Command:   command,
	}

	var res SessionMeterpreterWriteRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterWriteRes{}, err
	}

	return res, nil
}

func (s *session) MeterpreterRead(session int) (SessionMeterpreterReadRes, error) {
	req := &SessionMeterpreterReadReq{
		Method:    "session.meterpreter_read",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionMeterpreterReadRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterReadRes{}, err
	}

	return res, nil
}

func (s *session) MeterpreterRunSingle(session int, command string) (SessionMeterpreterRunSingleRes, error) {
	req := &SessionMeterpreterRunSingleReq{
		Method:    "session.meterpreter_run_single",
		Token:     s.rpc.Token(),
		SessionID: session,
		Command:   command,
	}

	var res SessionMeterpreterRunSingleRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterRunSingleRes{}, err
	}

	return res, nil
}

func (s *session) MeterpreterSessionDetach(session int) (SessionMeterpreterDetachRes, error) {
	req := &SessionMeterpreterDetachReq{
		Method:    "session.meterpreter_session_detach",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionMeterpreterDetachRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterDetachRes{}, err
	}

	return res, nil
}

func (s *session) MeterpreterSessionKill(session int) (SessionMeterpreterKillRes, error) {
	req := &SessionMeterpreterKillReq{
		Method:    "session.meterpreter_session_kill",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionMeterpreterKillRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterKillRes{}, err
	}

	return res, nil
}

type SessionMeterpreterTabsReq struct {
	Method    string
	Token     string
	SessionID int
	InputLine string
}

type SessionMeterpreterTabsRes struct {
	Tabs []string `msgpack:"tabs"`
}

func (s *session) MeterpreterTabs(sessionID int, inputLine string) (SessionMeterpreterTabsRes, error) {
	req := &SessionMeterpreterTabsReq{
		Method:    "session.meterpreter_tabs",
		Token:     s.rpc.Token(),
		SessionID: sessionID,
		InputLine: inputLine,
	}

	var res SessionMeterpreterTabsRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionMeterpreterTabsRes{}, err
	}

	return res, nil
}

func (s *session) ShellUpgrade(session int, lhostAddress string, lportNumber uint32) (SessionShellUpgradeRes, error) {
	req := &SessionShellUpgradeReq{
		Method:     "session.shell_upgrade",
		Token:      s.rpc.Token(),
		SessionID:  session,
		IPAddress:  lhostAddress,
		PortNumber: lportNumber,
	}

	var res SessionShellUpgradeRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionShellUpgradeRes{}, err
	}

	return res, nil
}

func (s *session) RingClear(session int) (SessionRingClearRes, error) {
	req := &SessionRingClearReq{
		Method:    "session.ring_clear",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionRingClearRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionRingClearRes{}, err
	}

	return res, nil
}

type SessionRingLastReq struct {
	Method    string
	Token     string
	SessionID int
}

type SessionRingLastRes struct {
	Seq uint32 `msgpack:"seq"`
}

func (s *session) RingLast(session int) (SessionRingLastRes, error) {
	req := &SessionRingLastReq{
		Method:    "session.ring_last",
		Token:     s.rpc.Token(),
		SessionID: session,
	}

	var res SessionRingLastRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionRingLastRes{}, err
	}

	return res, nil
}

func (s *session) RingPut(session int, command string) (SessionRingPutRes, error) {
	req := &SessionRingPutReq{
		Method:    "session.ring_put",
		Token:     s.rpc.Token(),
		SessionID: session,
		Command:   command,
	}

	var res SessionRingPutRes
	if err := s.rpc.Call(req, &res); err != nil {
		return SessionRingPutRes{}, err
	}

	return res, nil
}
