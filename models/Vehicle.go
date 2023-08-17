package models

type Vehicle struct {
	MyType string
}

func (v *Vehicle) Type() string {
	return v.MyType
}
