package armstrong

// Gitlab -
type GitlabAddGroupReq struct {
	ID int `json:"id,omitempty"`
	Type string `json:"type"`
	ProjectName string `json:"project_name"`
	Owner string `json:"owner"`
	UserId string `json:"userid"`
	DcDomain string `json:"dcdomain"`
	AdGroup string `json:"adgroup"`
}

type GitlabAddGroupResp struct {
	ID int `json:"id"`
	Message string `json:"message"`
	Status int `json:"status"`
}

type GitlabDeleteGroupResp struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

type GitlabGroupInfoResp struct {
	ID int `json:"id,omitempty"`
	ServiceName string `json:"service_name"`
	Type string `json:"project_type"`
	ProjectName string `json:"project_name"`
	Status string `json:"status"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	Owner string `json:"project_owner"`
	AdDomain string `json:"adDomain"`
	AdGroup string `json:"adgroup"`
	UserId string `json:"userid"`
}