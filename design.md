# API

The API provides the minimal core set of functionality needed to ACM@UIUC public facing infrastructure. There are three components to the core functionality: authentication, user management, group managment, and resume managment. Authentication is provided via oauth by a provider that can verify the user owns the email address of `<netid>@illinois.edu`. User management provides the ability to create an account, update account info, retrieve users, and mark users. Group management provides a read only interface to the truth store which is internally managed with git. Resume management allows students to upload their resumes and recruiters to the retrive all resumes.

## Routes
```
GET   /api/auth/oauth/:provider
GET   /api/auth/oauth/:provider/redirect
POST  /api/auth/oauth/:provider

GET   /api/user
POST  /api/user
GET   /api/user/filter
POST  /api/user/mark

GET   /api/group
GET   /api/group/verify

POST  /api/resume/upload
GET   /api/resume/filter
POST  /api/resume/approve
```

## Database Models
```
Token {
	Username string
	Token string
	Expiration int64
}

User {
	Username string
	FirstName string
	LastName string
	Mark string
}

Resume {
	Username string
	FirstName string
	LastName string
	Email string
	GraduationMonth int
	GraduationYear int
	Major string
	Degree string
	Seeking string
	BlobKey string
	Approved bool
}
```

## Gitstore Models
```
Group {
	Name string
	Description string
	Chairs string
	Members []GroupMember
	MeetingTime string
	MeetingLocation string
	Website string
	Email string
}

GroupMember {
	Role string
	Username string
	DisplayName string
	Email string
}
```

# Site

## Routes
```
GET  /
GET  /about
GET  /about/history
GET  /sigs
GET  /reflectionsprojections
GET  /hackathon
GET  /sponsors
GET  /join
GET  /login
GET  /logout
GET  /resumebook
GET  /resumeupload
GET  /intranet
GET  /intranet/resumemanager
GET  /intranet/usermanager
GET  /intranet/recruitermanager
GET  /intranet/recruitercreator
```

# Codebase

## Layout
```
/
	model/
		...
	service/
		auth/
			...
		user/
			...
		group/
			...
		resume/
			...
		...
	controller/
		auth/
			...
		user/
			...
		group/
			...
		...
	middleware/
		....
	server/
		...
	template/
		...
	database/
		migrations/
			...
		...
	test/
		service/
			auth/
				...
			user/
				...
			group/
				...
			resume/
				...
	static/
		images/
			...
		less/
			...
```
