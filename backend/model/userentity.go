package model

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserRegis struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type UserList struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Role      string `json:"role" `
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type ContentList struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type MentorList struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Skill string `json:"skill"`
}

type MentorSkill struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Skill string `json:"skill"`
	Image string `json:"image"`
}

type PayloadUser struct {
	Username string `json:"username" `
	Password string `json:"password"`
}

type Authorize struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
}

type MentorRegis struct {
	Skill   string `json:"skill"`
	Bio     string `json:"bio"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Image   string `json:"image"`
}

type MentorDetail struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Skill string `json:"skill"`
	Bio   string `json:"bio"`
	Image string `json:"image"`
}

type UserDetail struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Token struct {
	Token string `json:"token"`
	Role  string `json:"RLPP"`
	ID    int    `json:"id"`
}

type MentorKontak struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
