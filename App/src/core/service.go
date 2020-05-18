package [[ with .ModuleToParse ]][[.Name]][[ end ]]

//Service interface
type Service interface { [[ with .ModuleToParse.Model ]]
	Create[[.Name | ToCamel]](item Create[[.Name | ToCamel]]Req) (string, error)
	Get[[.Name | ToCamel]]ByID(id string) ([[.Name | ToCamel]], error)
	GetAll[[.Name | ToCamel | ToPlural]]() ([][[.Name | ToCamel]], error)
	Update[[.Name | ToCamel]](id string, update Update[[.Name | ToCamel]]Req) ([[.Name | ToCamel]], error)
	Delete[[.Name | ToCamel]](id string, deletedby string) ([[.Name | ToCamel]], error)[[ end ]]
}
