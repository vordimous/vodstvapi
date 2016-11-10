package forms

//ModelForm ...
type ModelForm interface {
	ToModel(m *interface{}) error
}
