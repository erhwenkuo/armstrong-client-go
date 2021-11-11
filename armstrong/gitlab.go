package armstrong

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateGitlabService - add group
func (c *Client) GitlabAddGroup(gitlabAddGroupReq GitlabAddGroupReq) (*GitlabAddGroupResp, error) {
	rb, err := json.Marshal(gitlabAddGroupReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gitlab_add_group", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	gitlabAddGroupResp := GitlabAddGroupResp{}
	err = json.Unmarshal(body, &gitlabAddGroupResp)
	if err != nil {
		return nil, err
	}

	return &gitlabAddGroupResp, nil
}

// DeleteGitlabService - remove group
func (c *Client) GitlabDeleteGroup(gitlabServiceID string) (*GitlabDeleteGroupResp, error) {
	req, err := http.NewRequest("DELETE",fmt.Sprintf("%s/gitlab_delete_group/%s", c.HostURL, gitlabServiceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	gitlabDeleteGroupResp := GitlabDeleteGroupResp{}
	err = json.Unmarshal(body, &gitlabDeleteGroupResp)
	if err != nil {
		return nil, err
	}

	return &gitlabDeleteGroupResp, nil
}

// GetGitlabServiceInfo - retrieve info
func (c *Client) GetGitlabGroupInfo(gitlabServiceID string) (*GitlabGroupInfoResp, error) {
	req, err := http.NewRequest("GET",fmt.Sprintf("%s/gitlab_get/%s", c.HostURL, gitlabServiceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	gitlabGroupInfoResp := GitlabGroupInfoResp{}
	err = json.Unmarshal(body, &gitlabGroupInfoResp)
	if err != nil {
		return nil, err
	}
	return &gitlabGroupInfoResp, nil
}