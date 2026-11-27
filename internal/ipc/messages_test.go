package ipc

import (
	"net"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestBasicMarshalUnmarshal(t *testing.T) {
	interval, _ := time.ParseDuration("10s")
	ttl, _ := time.ParseDuration("1m")
	table := &Table{
		Name:     "mytable",
		Id:       12345,
		Port:     8080,
		Disabled: false,
		Interval: durationpb.New(interval),
		Timeout:  durationpb.New(ttl),
		Hosts: []*Host{
			&Host{
				Id:   33333,
				Kind: AddrKind_IPV4,
				Addr: net.ParseIP("10.12.12.1"),
			},
		},
	}

	bytes, err := proto.Marshal(table)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	var table2 Table
	err = proto.Unmarshal(bytes, &table2)
	if err != nil {
		t.Errorf("could not unmarshal message: %v", err)
	}

	if !proto.Equal(table, &table2) {
		t.Errorf("expected {%v}, got {%v}", table, &table2)
	}
}
