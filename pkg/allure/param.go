package allure

type Parametrize struct {
	Name string
}

func (param *Parametrize) GetName() string {
	return param.Name
}
