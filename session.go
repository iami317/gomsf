package gomsf

import (
	"fmt"
	"github.com/iami317/gomsf/rpc"
	"strings"
)

type SessionManager struct {
	rpc *rpc.RPC
}

func (s *SessionManager) List() (rpc.SessionListRes, error) {
	return s.rpc.Session.List()
}

func (s *SessionManager) Read(sid int) (string, error) {
	r, err := s.rpc.Session.MeterpreterRead(sid)
	if err != nil {
		return "", err
	}

	return r.Data, nil
}

func (s *SessionManager) Write(sid int, command string) error {
	if !strings.HasSuffix(command, "\n") {
		command = fmt.Sprintf("%s\n", command)
	}

	r, err := s.rpc.Session.MeterpreterWrite(sid, command)
	if err != nil {

		return err
	}
	if r.Result == rpc.FAILURE {
		return fmt.Errorf("cannot write command %s to session %d", command, sid)
	}
	return nil
}

func (s *SessionManager) ShellWrite(sid int, command string) error {
	return s.rpc.Session.ShellWrite(sid, command)
}

func (s *SessionManager) ShellRead(session int, readPointer uint32) (string, error) {
	return s.rpc.Session.ShellRead(session, readPointer)
}

func (s *SessionManager) RunSingle(session int, cmd string) (rpc.SessionMeterpreterRunSingleRes, error) {
	return s.rpc.Session.MeterpreterRunSingle(session, cmd)
}

func (s *SessionManager) SessionDetach(sid int) (rpc.SessionMeterpreterDetachRes, error) {
	return s.rpc.Session.MeterpreterSessionDetach(sid)
}

func (s *SessionManager) SessionKill(sid int) (rpc.SessionMeterpreterKillRes, error) {
	return s.rpc.Session.MeterpreterSessionKill(sid)
}

func (s *SessionManager) Modules(sid int) ([]string, error) {
	r, err := s.rpc.Session.CompatibleModules(sid)
	if err != nil {
		return nil, err
	}

	return r.Modules, nil
}

func (s *SessionManager) Stop(sid int) (string, error) {
	r, err := s.rpc.Session.Stop(sid)
	if err != nil {
		return "", err
	}
	return r, nil
}
