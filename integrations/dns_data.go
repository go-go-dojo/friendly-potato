package integrations

import (
	"errors"
	cf "github.com/cloudflare/cloudflare-go"
	"time"
)

type DNSResource struct{
	Zone string
	Name string
	cfRecordId string
	cfZoneId string
	cfProxied bool
	cfModifiedOn time.Time
}

func (f DNSResource)FQDN()string{
	return f.Name + "." + f.Zone
}

func (f *DNSResource)translateFromCloudflare(record cf.DNSRecord){
	if f==nil{
		*f = DNSResource{}
	}
	f.Zone=record.ZoneName
	f.Name=record.Name
	f.cfRecordId=record.ID
	f.cfZoneId=record.ZoneID
	f.cfProxied=record.Proxied
	f.cfModifiedOn=record.ModifiedOn
}


type Record struct {
	Resource  DNSResource
	Ttl       uint
	DnsType   string
	DnsData   string
	FetchTime time.Time
}

func (r *Record) translateFromCloudflare(record cf.DNSRecord)(err error){
	// check if we can cast to uint
	if record.TTL<0{
		err = errors.New("invalid ttl from cloudflare")
		return
	}else{
		r.Ttl = uint(record.TTL)
	}

	r.Resource.translateFromCloudflare(record)
	r.DnsType = record.Type
	r.DnsData = record.Content
	r.FetchTime = time.Now()
	return
}

func (r Record) translateToCloudflare()(record cf.DNSRecord){
	record = cf.DNSRecord{
		ID:        	r.Resource.cfRecordId,
		Type:       r.DnsType,
		Name:       r.Resource.Name,
		Content:    r.DnsData,
		Proxied:    r.Resource.cfProxied,
		TTL:        int(r.Ttl),
		ZoneID: 	r.Resource.cfZoneId,
		ZoneName:   r.Resource.Zone,
	}
	return
}
type DomainResource struct{
	Name string
	cfZoneId string
	cfModifiedOn time.Time
	cfStatus string
	cfNameSevers []string
}


func (d *DomainResource)translateFromCloudflare(zone cf.Zone){
	if d==nil{
		*d=DomainResource{}
	}
	d.Name = zone.Name
	d.cfZoneId = zone.ID
	d.cfModifiedOn = zone.ModifiedOn
	d.cfNameSevers = zone.NameServers
	d.cfStatus = zone.Status
	return
}

func (d DomainResource)translateToCloudflare()(zone cf.Zone){
	zone=cf.Zone{
		ID:                d.cfZoneId,
		Name:              d.Name,
	}
	return
}

type Zone struct {
	Resource DomainResource
	Records  []Record
}

func (z *Zone) appendRecords(record ...Record) {
	if len(record) == 0 {
		return
	}
	if z.Records == nil {
		z.Records = []Record{}
	}
	z.Records = append(z.Records, record...)
	return
}

func (z *Zone) translateFromCloudflare(zone cf.Zone) {
	 if z == nil{
	 	*z = Zone{}
	 }
	 dr := DomainResource{}
	 dr.translateFromCloudflare(zone)
	 *z = Zone{
		Resource: dr,
		Records:  []Record{},
	}
}

func (z Zone)translateToCloudflare()(zone cf.Zone){
	zone=z.Resource.translateToCloudflare()
	return
}

type Zones []Zone

func (z Zones) Names()(names []string){
	names = []string{}
	for _,zz := range z{
		names = append(names,zz.Resource.Name)
	}
	return
}

func (z *Zones) appendZone(zone ...Zone) {
	if z == nil {
		*z = []Zone{}
	}
	if len(zone) == 0 {
		return
	}
	za := []Zone(*z)
	za = append(za, zone...)
	*z = za
}

func (z *Zones)translateFromCloudflare(zones []cf.Zone){
	for _, zc := range zones {
		zz := Zone{}
		zz.translateFromCloudflare(zc)
		z.appendZone(zz)
	}
}