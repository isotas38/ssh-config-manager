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
	var hosts Hosts
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string)

	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		fields := strings.Fields(line)
		if fields[0] == "Host" {
			if _, ok := m["Host"]; ok {
				host := new(Host)
				tmp, _ := json.Marshal(m)
				json.Unmarshal(tmp, &host)
				hosts = append(hosts, host)
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
