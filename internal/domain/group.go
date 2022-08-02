package domain

type Group struct {
	ID        int      `json:"group_id,omitempty"`
	GroupName string   `json:"groupname"`
	Members   []string `json:"group_members,omitempty"`
	Subgroup  []string `json:"subgroup,omitempty"`
}
