// в процессе

package domain

type Person struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	YearOfBirth int    `json:"year_of_birth"`
	GroupName   string `json:"groupname"`
}
