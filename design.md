# API

The API provides the minimal core set of functionality needed to ACM@UIUC public facing infrastructure. There are three components to the core functionality: authentication, user management, and group managment. Authentication is provided via oauth by a provider that can verify the user owns the email address of `<netid>@illinois.edu`. User management provides the ability to create an account, update account info, retrieve users, and mark users. Group management provides a read only interface to the truth store which is internally managed with git.

## Routes
```
GET   /api/auth/oauth/:provider
GET   /api/auth/oauth/:provider/redirect
POST  /api/auth/oauth/:provider

GET   /api/user
POST  /api/user
GET   /api/user/filter
POST  /api/user/mark

GET   /api/group/filter
GET   /api/group/verify
```

## Database Models
```
Token {
	Netid string
	Token string
	Expiration int64
}

User {
	Netid string
	Name string
	GraduationYear int
	Major string
	Resume string
	Mark string
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
GET  /sigs
GET  /conference
GET  /hackathon
GET  /sponsors
GET  /join
GET  /intranet
GET  /resumebook
GET  /resumeupload
```

# Codebase

## Layout
```
/
	api/
		model/
			...
		service/
			auth/
				...
			user/
				...
			group/
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
		database/
			migrations/
				...
			...
	site/
		partial/
			...
		...
```
