package twilio

// SIPService groups the SIP configuration resources exposed under the
// /SIP path of the core Twilio API: Credential Lists, Domains, and IP Access
// Control Lists, along with their sub-resources and Domain auth mappings.
//
// Access it via Client.SIP. For example, to register a SIP endpoint you would
// create a CredentialList and Credential, create a Domain, and then map the
// CredentialList to the Domain for registration:
//
//	cl, _ := client.SIP.CredentialLists.Create(ctx, clData)
//	client.SIP.CredentialLists.Credentials(cl.Sid).Create(ctx, credData)
//	d, _ := client.SIP.Domains.Create(ctx, domainData)
//	mapping := url.Values{"CredentialListSid": {cl.Sid}}
//	client.SIP.Domains.AuthRegistrationsCredentialListMappings(d.Sid).Create(ctx, mapping)
type SIPService struct {
	CredentialLists      *SIPCredentialListService
	Domains              *SIPDomainService
	IPAccessControlLists *SIPIPAccessControlListService
}
