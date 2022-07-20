package domain

type Country struct {
	model
	Name string `sql:"size:1024" json:"name"`
}
