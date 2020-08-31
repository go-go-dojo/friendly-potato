// +build integration

package integrations_test

import (
	. "friendly-potato/integrations"
	"net"
	"os"
	"reflect"
	"testing"
)

func TestHttpBin_GetExtIP(t *testing.T) {

	t.Logf("setup httpbin")
	f, err := NewHttpBin()

	if err != nil {
		t.Logf("NewHttpBin() failed: %v",err)
		t.FailNow()
	}

	InitHttpBin()
	if err != nil {
		t.Logf("InitHttpBin() failed: %v",err)
		t.FailNow()
	}

	tests := []struct {
		name    string
		fields  *HttpBin
		wantIp  net.IP
		wantErr bool
	}{
		{
			name:    "test get external ip",
			fields:  f,
			wantIp:  f.Resolved,
			wantErr: false,
		},
	}

	expectedKnownIP, ok := os.LookupEnv("EXTERNAL_IP")
	if !ok {
		t.Log("EXTERNAL_IP env not set skipping")
	} else {
		t.Logf("got ext ip %v",expectedKnownIP)
		extTest := struct {
			name    string
			fields  *HttpBin
			wantIp  net.IP
			wantErr bool
		}{
			name:    "test known address",
			fields:  f,
			wantIp:  net.ParseIP(expectedKnownIP),
			wantErr: false,
		}
		tests = append(tests, extTest)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h,err := NewHttpBin()
			if err != nil {
				t.Logf("NewHttpBin() failed: %v",err)
				t.FailNow()
			}
			gotIp, err := h.GetExtIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExtIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIp, tt.wantIp) {
				t.Errorf("GetExtIP() gotIp = %v, want %v", gotIp, tt.wantIp)
			}
		})
	}
}
