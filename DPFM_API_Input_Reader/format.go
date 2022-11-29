package dpfm_api_input_reader

import (
	"data-platform-api-distribution-channel-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToDistributionChannel() *requests.DistributionChannel {
	data := sdc.DistributionChannel
	return &requests.DistributionChannel{
		DistributionChannel: data.DistributionChannel,
	}
}

func (sdc *SDC) ConvertToDistributionChannelText() *requests.Text {
	dataDistributionChannel := sdc.DistributionChannel
	data := sdc.DistributionChannel.Text
	return &requests.Text{
		DistributionChannel:     dataDistributionChannel.DistributionChannel,
		Language:                data.Language,
		DistributionChannelName: data.DistributionChannelName,
	}
}
