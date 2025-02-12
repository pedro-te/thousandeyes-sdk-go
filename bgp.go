package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// BGP - BGP trace test
type BGP struct {
	// Common test fields
	AlertsEnabled      *bool                `json:"alertsEnabled,omitempty" te:"int-bool"`
	AlertRules         *[]AlertRule         `json:"alertRules,omitempty"`
	APILinks           *[]APILink           `json:"apiLinks,omitempty"`
	CreatedBy          *string              `json:"createdBy,omitempty"`
	CreatedDate        *string              `json:"createdDate,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Enabled            *bool                `json:"enabled,omitempty" te:"int-bool"`
	Groups             *[]GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         *string              `json:"modifiedBy,omitempty"`
	ModifiedDate       *string              `json:"modifiedDate,omitempty"`
	SavedEvent         *bool                `json:"savedEvent,omitempty" te:"int-bool"`
	SharedWithAccounts *[]SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             *int64               `json:"testId,omitempty"`
	TestName           *string              `json:"testName,omitempty"`
	Type               *string              `json:"type,omitempty"`
	LiveShare          *bool                `json:"liveShare,omitempty" te:"int-bool"`

	// Fields unique to this test
	BGPMonitors            *[]BGPMonitor `json:"bgpMonitors,omitempty"`
	IncludeCoveredPrefixes *bool         `json:"includeCoveredPrefixes,omitempty" te:"int-bool"`
	Prefix                 *string       `json:"prefix,omitempty"`
	UsePublicBGP           *bool         `json:"usePublicBgp,omitempty" te:"int-bool"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t BGP) MarshalJSON() ([]byte, error) {
	type aliasTest BGP

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *BGP) UnmarshalJSON(data []byte) error {
	type aliasTest BGP
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAlertRule - Adds an alert to agent test
func (t *BGP) AddAlertRule(id int64) {
	alertRule := AlertRule{RuleID: Int64(id)}
	*t.AlertRules = append(*t.AlertRules, alertRule)
}

// GetBGP  - get bgp test
func (c *Client) GetBGP(id int64) (*BGP, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &BGP{}, err
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateBGP - Create bgp test
func (c Client) CreateBGP(t BGP) (*BGP, error) {
	resp, err := c.post("/tests/bgp/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create test, response code %d", resp.StatusCode)
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteBGP - delete bgp test
func (c *Client) DeleteBGP(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/bgp/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete bgp test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateBGP - - Update bgp trace test
func (c *Client) UpdateBGP(id int64, t BGP) (*BGP, error) {
	resp, err := c.post(fmt.Sprintf("/tests/bgp/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]BGP
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
