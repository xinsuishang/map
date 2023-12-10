package entity

type DomainMappingEntity struct {
	ID         int32
	TenantID   int32
	RegionID   string
	Domain     string
	BucketName string
	Desc       string
}
