package twilio

// ElasticTrunkService lets you interact with Elastic Trunk Resources.
type ElasticTrunkService struct {
	Trunks          *TrunkService
	OriginationUrls *OriginationUrlService
}
