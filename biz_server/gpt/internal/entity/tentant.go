package entity

type TenantEntity struct {
	ID            int32
	ParentId      int32
	Name          string
	IsApplication bool
	Model         string
	AccessKey     string
	SecretKey     string
	Desc          string
	Dashboard     string
}
