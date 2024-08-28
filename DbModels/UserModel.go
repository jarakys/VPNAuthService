package DbModels

type UserModel struct {
	Id        uint `gorm:"primaryKey"`
	IsPremium bool
	LastVisit int64
}

func (UserModel) TableName() string {
	return "users"
}
