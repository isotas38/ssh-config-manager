package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func Parse(r io.Reader) (Hosts, error) {
	m := make(map[string][]string)
	var hosts Hosts
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if fields[0] == "Host" {
			if _, ok := m["Host"]; ok {
				host := new(Host)
				tmp, _ := json.Marshal(m)
				json.Unmarshal(tmp, &host)
				hosts = append(hosts, host)
				m = make(map[string][]string)
			}
		}
		m[fields[0]] = fields[1:]
	}
	if _, ok := m["Host"]; ok {
		host := new(Host)
		tmp, _ := json.Marshal(m)
		json.Unmarshal(tmp, &host)
		hosts = append(hosts, host)
	}
	return hosts, nil
}
