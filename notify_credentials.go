package twilio

import (
	"context"
	"net/url"
)

type CredentialType string

const credentialsPathPart = "Credentials"

const (
	TypeGCM = CredentialType("gcm")
	TypeFCM = CredentialType("fcm")
	TypeAPN = CredentialType("apn")
)

type NotifyCredentialsService struct {
	client *Client
}

type NotifyCredential struct {
	Sid          string         `json:"sid"`
	FriendlyName string         `json:"friendly_name"`
	AccountSid   string         `json:"account_sid"`
	Type         CredentialType `json:"type"`
	DateCreated  TwilioTime     `json:"date_created"`
	DateUpdated  TwilioTime     `json:"date_updated"`
	URL          string         `json:"url"`
}

type NotifyCredentialPage struct {
	Page
	Credentials []*NotifyCredential `json:"credentials"`
}

type NotifyCredentialPageIterator struct {
	p *PageIterator
}

func (n *NotifyCredentialsService) Create(ctx context.Context, data url.Values) (*NotifyCredential, error) {
	credential := new(NotifyCredential)
	err := n.client.CreateResource(ctx, credentialsPathPart, data, credential)
	return credential, err
}

func (n *NotifyCredentialsService) GetPage(ctx context.Context, data url.Values) (*NotifyCredentialPage, error) {
	iter := n.GetPageIterator(data)
	return iter.Next(ctx)
}

func (n *NotifyCredentialsService) Get(ctx context.Context, sid string) (*NotifyCredential, error) {
	credential := new(NotifyCredential)
	err := n.client.GetResource(ctx, credentialsPathPart, sid, credential)
	return credential, err
}

func (n *NotifyCredentialsService) Update(ctx context.Context, sid string, data url.Values) (*NotifyCredential, error) {
	credential := new(NotifyCredential)
	err := n.client.UpdateResource(ctx, credentialsPathPart, sid, data, credential)
	return credential, err
}

func (n *NotifyCredentialsService) Delete(ctx context.Context, sid string) error {
	return n.client.DeleteResource(ctx, credentialsPathPart, sid)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (n *NotifyCredentialsService) GetPageIterator(data url.Values) *NotifyCredentialPageIterator {
	iter := NewPageIterator(n.client, data, credentialsPathPart)
	return &NotifyCredentialPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (r *NotifyCredentialPageIterator) Next(ctx context.Context) (*NotifyCredentialPage, error) {
	rp := new(NotifyCredentialPage)
	err := r.p.Next(ctx, rp)
	if err != nil {
		return nil, err
	}
	r.p.SetNextPageURI(rp.NextPageURI)
	return rp, nil
}
