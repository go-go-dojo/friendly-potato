// +build integration

package integrations_test

import (
	. "friendly-potato/integrations"
	"net"
	"os"
	"reflect"
	"testing"
)

func TestGetExternalIP(t *testing.T) {
	type testS struct{
		name string
		wantErr error
		wantIP net.IP
	}
	type testsS []testS

	t.Log("setup httpbin")
	httpBin,err:= NewHttpBin()

	if err!=nil{
		t.Logf("failed to setup httpbin %v",err)
		t.FailNow()
	}

	t.Log("setup ipfy")
	ipFy,err:= NewIPFy()

	if err!=nil{
		t.Logf("failed to setup ipfy %v",err)
		t.FailNow()
	}
	tests:=testsS{
		{
			name:    "test known httpbin IP addr",
			wantErr: nil,
			wantIP:  httpBin.Resolved,
		},
		{
			name:    "test known ipfy IP addr",
			wantErr: nil,
			wantIP:  ipFy.Resolved,
		},
	}
	t.Log("check if we know the external ip")
	expectedIP, ok := os.LookupEnv("EXTERNAL_IP")
	if !ok {
		t.Logf("EXTERNAL_IP env not set skipping")
	}else{
		t.Logf("got ext ip %v",expectedIP)
		knownExt := testS{
			name:    "test known ext IP addr",
			wantErr: nil,
			wantIP:  net.ParseIP(expectedIP),
		}
		tests=append(tests,knownExt)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err:=InitHttpBin()
			if err!=nil{
				t.Logf("failed InitHttpBin(): %v",err)
				t.FailNow()
			}
			err=InitIPFy()
			if err!=nil{
				t.Logf("failed InitIPFy(): %v",err)
				t.FailNow()
			}
			gotErr, gotIP := GetExternalIP()
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetExternalIP() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotIP, tt.wantIP) {
				t.Errorf("GetExternalIP() gotIP = %v, want %v", gotIP, tt.wantIP)
			}
		})
	}
}
