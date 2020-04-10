package model

// Village represent village model
type Village struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// District represent district model
type District struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Villages []Village `json:"villages"`
}

// Regency represent regency model
type Regency struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Districts []District `json:"districts"`
}

// Province represent province model
type Province struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Regencies []Regency `bson:"regencies" json:"regencies,omitempty"`
}
