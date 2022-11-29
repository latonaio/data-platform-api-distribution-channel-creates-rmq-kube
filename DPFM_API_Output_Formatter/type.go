package dpfm_api_output_formatter

type DistributionChannel struct {
	DistributionChannel string `json:"DistributionChannel"`
	Text                Text   `json:"Text"`
}

type Text struct {
	DistributionChannel     string `json:"DistributionChannel"`
	Language                string `json:"Language"`
	DistributionChannelName string `json:"DistributionChannelName"`
}
