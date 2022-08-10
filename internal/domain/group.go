package domain

type Group struct {
	ID          int    `json:"group_id,omitempty"`
	GroupName   string `json:"groupname"`
	Members     int    `json:"group_members,omitempty"`
	Subgroup    bool   `json:"subgroup,omitempty"`
	MotherGroup string `json:"mother_group,omitempty"`
}
