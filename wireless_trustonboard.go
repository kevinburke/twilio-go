package twilio

import (
	"context"
	"net/url"

	"golang.org/x/sync/errgroup"
)

const tobPathPart = "trust-onboard"
const simCertsPathPart = tobPathPart + "/sim-certificates"

type TrustOnBoardService struct {
	client *Client
}

// SimCertificate represents a Sim's Certificate resource.
type SimCertificate struct {
	Sid             string          `json:"sid"`
	SimSid          string          `json:"sim_sid"`
	ICCID           string          `json:"iccid"`
	ContentsPEM     string          `json:"contents_pem"`
	CertificateType CertificateType `json:"certificate_type"`
}

// GetSimCertificates gets the certificates for a collection of SIMs by SID.
// Uses threads to retrieve all certs simultaneously; if retrieving any cert
// fails, we return an error for the entire request.
//
// Requests are made in batches of 10
func (t *TrustOnBoardService) GetSimCertificates(ctx context.Context, sids ...string) ([]*SimCertificate, error) {
	if len(sids) == 0 {
		certs := make([]*SimCertificate, 0)
		return certs, nil
	}

	// Split the slice into batches of 20 items.
	certColls := make([][]*SimCertificate, (len(sids)/10)+1)
	g, errctx := errgroup.WithContext(ctx)
	batch := 10
	collIdx := 0
	for i := 0; i < len(sids); i += batch {
		j := i + batch
		if j > len(sids) {
			j = len(sids)
		}

		// Add this batch to query params
		queryParams := make(url.Values)
		for _, sid := range sids[i:j] {
			queryParams.Add("SimSid", sid)
		}

		curCollIdx := collIdx
		collIdx++

		g.Go(func() error {
			var batchCerts []*SimCertificate
			err := t.client.MakeRequest(errctx, "GET", simCertsPathPart, queryParams, &batchCerts)
			if err != nil {
				return err
			}

			certColls[curCollIdx] = batchCerts
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Put all the collections back together
	var certs []*SimCertificate
	for _, coll := range certColls {
		certs = append(certs, coll...)
	}

	return certs, nil
}
