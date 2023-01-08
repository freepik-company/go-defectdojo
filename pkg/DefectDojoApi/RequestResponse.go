package DefectDojoApi

type ReqResp struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type RequestResponse struct {
	ReqResp []ReqResp `json:"req_resp"`
}
