package twilio

import (
	"context"
	"net/url"
)

const trunkPathPart = "Trunks"

type TrunkService struct {
	client *Client
}

type Trunk struct {
	Sid                    string `json:"sid"`
	FriendlyName           string `json:"friendly_name"`
	DomainName             string `json:"domain_name"`
	DisasterRecoveryUrl    string `json:"disaster_recovery_url"`
	DisasterRecoveryMethod string `json:"disaster_recovery_method"`
	TransferMode           string `json:"transfer_mode"`
	Secure                 bool   `json:"secure"`
	CnamLookupEnabled      bool   `json:"cname_lookup_enabled"`
}

type TrunkPage struct {
	Meta   Meta     `json:"meta"`
	Trunks []*Trunk `json:"trunks"`
}

type TrunkPageIterator struct {
	p *PageIterator
}

// returns an elastic trunk
func (ets *TrunkService) Get(ctx context.Context, sid string) (*Trunk, error) {
	trunk := new(Trunk)
	err := ets.client.GetResource(ctx, trunkPathPart, sid, trunk)
	return trunk, err
}

// Create a trunk with the given url.Values. For more information on valid values,
// see https://www.twilio.com/docs/sip-trunking/api/trunk-resource#create-a-trunk-resource or use the
func (ets *TrunkService) Create(ctx context.Context, data url.Values) (*Trunk, error) {
	trunk := new(Trunk)
	err := ets.client.CreateResource(ctx, trunkPathPart, data, trunk)
	return trunk, err
}

func (ets *TrunkService) Delete(ctx context.Context, sid string) error {
	return ets.client.DeleteResource(ctx, trunkPathPart+sid, sid)
}

func (ets *TrunkService) Update(ctx context.Context, sid string, data url.Values) (*Trunk, error) {
	trunk := new(Trunk)
	err := ets.client.UpdateResource(ctx, trunkPathPart+sid, sid, data, trunk)
	return trunk, err
}

// Returns a list of trunks. For more information on valid values,
// seehttps://www.twilio.com/docs/sip-trunking/api/trunk-resource#read-multiple-trunk-resources
func (ets *TrunkService) GetPage(ctx context.Context, data url.Values) (*TrunkPage, error) {
	return ets.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (ets *TrunkService) GetPageIterator(data url.Values) *TrunkPageIterator {
	iter := NewPageIterator(ets.client, data, trunkPathPart)
	return &TrunkPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (ets *TrunkPageIterator) Next(ctx context.Context) (*TrunkPage, error) {
	rp := new(TrunkPage)
	err := ets.p.Next(ctx, rp)
	if err != nil {
		return nil, err
	}
	ets.p.SetNextPageURI(rp.Meta.NextPageURL)
	return rp, nil
}
