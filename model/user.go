package model

type User struct {
	Username       string `db:"username" json:"username"`
	FirstName      string `db:"first_name" json:"first_name"`
	LastName       string `db:"last_name" json:"last_name"`
	GraduationYear int32  `db:"graduation_year" json:"graduation_year"`
	Major          string `db:"major" json:"major"`
	Mark           string `db:"mark" json:"mark"`
}

const (
	UserMarkBasic     = "BASIC"
	UserMarkPaid      = "PAID"
	UserMarkRecruiter = "RECRUITER"
)

var (
	UserValidMarks = []string{UserMarkBasic, UserMarkPaid, UserMarkRecruiter}
)
