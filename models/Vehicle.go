package models

type Vehicle struct {
	myType string
	
}

func (v *Vehicle) Type () string {
	return v.myType
}