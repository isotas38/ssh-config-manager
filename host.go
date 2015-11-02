package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Hosts []*Host

type Host struct {
	Host                             []string
	Match                            []string
	AddressFamily                    []string
	BatchMode                        []string
	BindAddress                      []string
	CanonicalDomains                 []string
	CanonicalizefallbackLocal        []string
	Canonicalizehostname             []string
	CanonicalizemaxDots              []string
	CanonicalizepermittedCNAMEs      []string
	ChallengeResponseAuthentication  []string
	CheckHostIP                      []string
	Cipher                           []string
	Ciphers                          []string
	ClearAllForwardings              []string
	Compression                      []string
	CompressionLevel                 []string
	ConnectionAttempts               []string
	ConnectTimeout                   []string
	ControlMaster                    []string
	ControlPath                      []string
	ControlPersist                   []string
	DynamicForward                   []string
	EnableSSHKeysign                 []string
	EscapeChar                       []string
	ExitOnForwardFailure             []string
	FingerprintHash                  []string
	ForwardAgent                     []string
	ForwardX11                       []string
	ForwardX11Timeout                []string
	ForwardX11Trusted                []string
	GatewayPorts                     []string
	GlobalKnownHostsFile             []string
	GSSAPIAuthentication             []string
	GSSAPIDelegateCredentials        []string
	HashKnownHosts                   []string
	HostbasedAuthentication          []string
	HostbasedKeyTypes                []string
	HostKeyAlgorithms                []string
	HostKeyAlias                     []string
	HostName                         []string
	IdentitiesOnly                   []string
	IdentityFile                     []string
	IgnoreUnknown                    []string
	IPQoS                            []string
	KbdInteractiveAuthentication     []string
	KbdInteractiveDevices            []string
	KexAlgorithms                    []string
	LocalCommand                     []string
	LocalForward                     []string
	LogLevel                         []string
	MACs                             []string
	NoHostAuthenticationForLocalhost []string
	NumberOfPasswordPrompts          []string
	PasswordAuthentication           []string
	PermitLocalCommand               []string
	PKCS11Provider                   []string
	Port                             []string
	PreferredAuthentications         []string
	Protocol                         []string
	ProxyCommand                     []string
	ProxyUseFdpass                   []string
	PubkeyAuthentication             []string
	RekeyLimit                       []string
	RemoteForward                    []string
	RequestTTY                       []string
	RevokedHostKeys                  []string
	RhostsRSAAuthentication          []string
	RSAAuthentication                []string
	SendEnv                          []string
	ServerAliveCountMax              []string
	ServerAliveInterval              []string
	StreamLocalBindMask              []string
	StreamLocalBindUnlink            []string
	StrictHostKeyChecking            []string
	TCPKeepAlive                     []string
	Tunnel                           []string
	TunnelDevice                     []string
	UpdateHostKeys                   []string
	UsePrivilegedPort                []string
	User                             []string
	UserKnownHostsFile               []string
	VerifyHostKeyDNS                 []string
	VisualHostKey                    []string
	XAuthLocation                    []string
}

func (hosts Hosts) String() string {
	buf := &bytes.Buffer{}
	for _, v := range hosts {
		fmt.Fprintln(buf, v)
	}
	return buf.String()
}

func (host *Host) String() string {
	m := make(map[string][]string)
	tmp, _ := json.Marshal(*host)
	json.Unmarshal(tmp, &m)

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s %s\n", "Host", strings.Join(host.Host, " "))
	for k := range m {
		if k != "Host" && len(m[k]) > 0 {
			fmt.Fprintf(buf, "  %s %s\n", k, strings.Join(m[k], " "))
		}
	}
	return buf.String()
}

func (hosts Hosts) GetHost(name string) (int, *Host) {
	for i1, v1 := range hosts {
		for _, v2 := range v1.Host {
			if v2 == name {
				return i1, v1
			}
		}
	}
	return -1, nil
}

func (hosts Hosts) addHost(name, ip string) Hosts {
	host := &Host{
		Host: []string{name},
	}
	s := []string{ip}
	fVal := reflect.ValueOf(host).Elem()
	sVal := reflect.ValueOf(s)
	fVal.FieldByName("HostName").Set(sVal)
	hosts = append(hosts, host)
	return hosts
}

func (hosts Hosts) updateHost(name, ip string) Hosts {
	for i1, v1 := range hosts {
		for _, v2 := range v1.Host {
			if v2 == name {
				s := []string{ip}
				sVal := reflect.ValueOf(s)
				reflect.ValueOf(hosts[i1]).Elem().FieldByName("HostName").Set(sVal)
				return hosts
			}
		}
	}
	return nil
}

func (hosts Hosts) listHost() {
	for _, v1 := range hosts {
		for _, v2 := range v1.Host {
			fmt.Println(v2)
		}
	}
}

func (hosts Hosts) removeHost(name string) Hosts {
	for i1, v1 := range hosts {
		for _, v2 := range v1.Host {
			if v2 == name {
				copy(hosts[i1:], hosts[i1+1:])
				hosts[len(hosts)-1] = nil
				hosts = hosts[:len(hosts)-1]
				return hosts
			}
		}
	}
	return nil
}
