package model

type ExamplePerson struct {
	ID         uint   `json:"id" gorm:"primaryKey;column:id"`
	FullName   string `form:"full_name" json:"full_name" gorm:"column:full_name"`
	Avatar     string `form:"avatar" json:"avatar" gorm:"column:avatar"`
	Post       string `form:"post" json:"post" gorm:"column:post"`
	Email      string `form:"email" json:"email" gorm:"column:email"`
	City       string `form:"city" json:"city" gorm:"column:city"`
	StartDate  string `form:"start_date" json:"start_date" gorm:"column:start_date"`
	Salary     string `form:"salary" json:"salary" gorm:"column:salary"`
	Age        string `form:"age" json:"age" gorm:"column:age"`
	Experience string `form:"experience" json:"experience" gorm:"column:experience"`
	Status     int    `form:"status" json:"status" gorm:"column:status;index"`
}
