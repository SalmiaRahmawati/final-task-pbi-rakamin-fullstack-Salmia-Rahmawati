package app

type User struct {
	GormApp
	Username string  `gorm:"not null" json:"username" valid:"required~Username is required"`
	Email    string  `gorm:"unique" json:"email" valid:"required~Email is required,email"`
	Password string  `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
}

type UserSignUp struct {
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"required~Email is required,email"`
	Password string `json:"password" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
}

type UserSignUpOutput struct {
	GormApp
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"required~Email is required,email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	u.Password = helpers.HashPass(u.Password)
// 	err = nil
// 	return
// }

// type UserSignUp struct {
// 	Username string `json:"username" valid:"required~username is required"`
// 	Email    string `json:"email" valid:"required~email is required"`
// 	Password string `json:"password" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
// }

// type UserSignUpCreate struct {
// 	ID        uint64         `gorm:"primaryKey" json:"id"`
// 	Username  string         `json:"username"`
// 	Email     string         `json:"email"`
// 	CreatedAt *time.Time     `json:"created_at,omitempty"`
// 	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
// 	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
// }

// type UserSignIn struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type SignInOutput struct {
// 	Token string `json:"token"`
// }
