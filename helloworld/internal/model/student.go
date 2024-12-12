package model

type Student struct {
	Id    int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name  string `gorm:"column:name;type:varchar(255)" json:"name"`
	Age   int64  `gorm:"column:age;type:bigint(20)" json:"age"`
	Total int64  `gorm:"column:total;type:bigint(20)" json:"total"`
}

func (m *Student) TableName() string {
	return "student"
}
