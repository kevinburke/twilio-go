package twilio

import (
	"context"
	"net/url"
)

type OriginationUrlService struct {
	client *Client
}

type OriginationUrl struct {
	Sid          string `json:"sid"`
	FriendlyName string `json:"friendly_name"`
	Weight       string `json:"weight"`
	Enabled      bool   `json:"enabled"`
	DateCreated  string `json:"date_created"`
	DateUpdated  string `json:"date_updated"`
	Priority     int    `json:"priority"`
	SipUrl       string `json:"sip_url"`
	TrunkSid     string `json:"trunk_sid"`
	Url          string `json:"url"`
}

type OriginationUrlPage struct {
	Meta           Meta              `json:"meta"`
	OriginationUrl []*OriginationUrl `json:"origination_urls"`
}

type OriginationUrlIterator struct {
	p *PageIterator
}

// returns an elastic trunk
func (ous *OriginationUrlService) Get(ctx context.Context, sid string) (*Trunk, error) {
	trunk := new(Trunk)
	err := ous.client.GetResource(ctx, trunkPathPart+"/"+WorkflowPathPart, sid, trunk)
	return trunk, err
}

// see https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#create-an-originationurl-resource
func (ous *OriginationUrlService) Create(ctx context.Context, data url.Values) (*Trunk, error) {
	trunk := new(Trunk)
	err := ous.client.CreateResource(ctx, trunkPathPart, data, trunk)
	return trunk, err
}

func (ous *OriginationUrlService) Delete(ctx context.Context, sid string) error {
	return ous.client.DeleteResource(ctx, trunkPathPart+sid, sid)
}

func (ous *OriginationUrlService) Update(ctx context.Context, sid string, data url.Values) (*Trunk, error) {
	trunk := new(Trunk)
	err := ous.client.UpdateResource(ctx, trunkPathPart+sid, sid, data, trunk)
	return trunk, err
}

// see https://www.twilio.com/docs/sip-trunking/api/originationurl-resource#fetch-an-originationurl-resource
func (ous *OriginationUrlService) GetPage(ctx context.Context, data url.Values) (*OriginationUrlPage, error) {
	return ous.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (ous *OriginationUrlService) GetPageIterator(data url.Values) *OriginationUrlIterator {
	iter := NewPageIterator(ous.client, data, trunkPathPart)
	return &OriginationUrlIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (ous *OriginationUrlIterator) Next(ctx context.Context) (*OriginationUrlPage, error) {
	rp := new(OriginationUrlPage)
	err := ous.p.Next(ctx, rp)
	if err != nil {
		return nil, err
	}
	ous.p.SetNextPageURI(rp.Meta.NextPageURL)
	return rp, nil
}
