package entity

type TenantEntity struct {
	ID        int32
	Name      string
	Region    bool
	Type      string
	AccessKey string
	SecretKey string
}
