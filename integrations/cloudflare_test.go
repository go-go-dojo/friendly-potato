//+build integration

package integrations_test

import (
	. "friendly-potato/integrations"
	"os"
	"reflect"
	"testing"
)

func TestInitAPI(t *testing.T) {
	t.Logf("lookup CF_TOKEN env")
	cfToken, ok := os.LookupEnv("CF_TOKEN")
	if !ok {
		t.Skipf("could not find cloudflare token as env SKIPPING...")
	}
	type args struct {
		apiToken string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test init cloudflare api",
			args:    args{cfToken},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitCloudFlareAPI(tt.args.apiToken); (err != nil) != tt.wantErr {
				t.Errorf("InitCloudFlareAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateZone(t *testing.T) {
	t.Logf("lookup CF_TOKEN env")
	cfToken, ok := os.LookupEnv("CF_TOKEN")
	if !ok {
		t.Skipf("could not find cloudflare token as env SKIPPING...")
	} else {
		err := InitCloudFlareAPI(cfToken)
		if err != nil {
			t.Logf("failed to initialize cloudflare api: %v", err)
		}
	}
	type args struct {
		zone Zone
	}
	tests := []struct {
		name            string
		args            args
		wantCreatedZone Zone
		wantErr         bool
	}{
		{
			name: "create zone blablabla.com",
			args: args{
				zone: Zone{
					Name: "blablabla.com",
				},
			},
			wantCreatedZone: Zone{
				Name: "blablabla.com",
			},
			wantErr: false,
		},
		{
			name: "create zone xptozone.com",
			args: args{
				zone: Zone{
					Name: "xptozone.com",
				},
			},
			wantCreatedZone: Zone{
				Name: "xptozone.com",
			},
			wantErr: false,
		},
		{
			name: "create duplicate zone xptozone.com",
			args: args{
				zone: Zone{
					Name: "xptozone.com",
				},
			},
			wantCreatedZone: Zone{
				Name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCreatedZone, err := CreateZone(tt.args.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateZone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCreatedZone.Name, tt.wantCreatedZone.Name) {
				t.Errorf("CreateZone() gotCreatedZone = %v, want %v", gotCreatedZone.Name, tt.wantCreatedZone.Name)
			}
		})
	}
}

func TestListZones(t *testing.T) {
	t.Logf("lookup CF_TOKEN env")
	cfToken, ok := os.LookupEnv("CF_TOKEN")
	if !ok {
		t.Skipf("could not find cloudflare token as env SKIPPING...")
	} else {
		err := InitCloudFlareAPI(cfToken)
		if err != nil {
			t.Logf("failed to initialize cloudflare api: %v", err)
		}
	}
	tests := []struct {
		name      string
		wantZones Zones
		wantErr   bool
	}{
		{
			name:      "list zone find zone xptozone.com, blablabla.com",
			wantZones: Zones{
				Zone{
					Name: "xptozone.com",
				},
				{
					Name: "blablabla.com",
				},
			},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotZones, err := ListZones()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListZones() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotZones.Names(), tt.wantZones.Names()) {
				t.Errorf("ListZones() gotZones = %v, want %v", gotZones.Names(), tt.wantZones.Names())
			}
		})
	}
}
