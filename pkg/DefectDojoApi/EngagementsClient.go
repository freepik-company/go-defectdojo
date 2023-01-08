package DefectDojoApi

import (
	"encoding/json"
	"fmt"
	"github.com/freepik-company/go-defectdojo/pkg/Client"
)

type EngagementsClient struct {
	client Client.Client
}

func NewEngagementsClient(client Client.Client) *EngagementsClient {
	return &EngagementsClient{client: client}
}

func (this EngagementsClient) All() *EngagementResult {
	response := this.client.Get("/api/v2/engagements/")
	result := new(EngagementResult)
	_ = json.Unmarshal([]byte(response), &result)

	return result
}

func (this EngagementsClient) Create(engagement Engagement) *Engagement {
	response := this.client.Post("/api/v2/engagements/", engagement.String())
	return newEngagementFromJson(response)
}

func (this EngagementsClient) Get(id int) *Engagement {
	response := this.client.Get(fmt.Sprintf("/api/v2/engagements/%d/", id))
	return newEngagementFromJson(response)
}

func (this EngagementsClient) Update(engagement Engagement) *Engagement {
	response := this.client.Put(
		fmt.Sprintf("/api/v2/engagements/%d/", engagement.Id),
		engagement.String(),
	)
	return newEngagementFromJson(response)
}

func (this EngagementsClient) Delete(engagement Engagement) {
	_ = this.client.Delete(fmt.Sprintf("/api/v2/engagements/%d/", engagement.Id))
}
