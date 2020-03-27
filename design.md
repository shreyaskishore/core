# Core

Core provides the minimal core set of functionality needed to run ACM@UIUC public facing infrastructure. There are three components to the core functionality: authentication, user management, group managment, and resume managment. Core also generates and servers the html pages for the ACM@UIUC website. All additional functionality will be implemented via Extensions.

## API Service Overview

### Authentication
All authentication is handled via external oauth providers that enforce verification of email ownership. Authentication for students and facultu is provided via oauth by a provider that can verify the user owns the email address of `<netid>@illinois.edu`. Authentication for recruiters and sponsors is provided via oauth by a provider which can verify ownership of the recruiter's email address. At this the planned oauth providers for students and faculty are google and microsoft. The planned oauth providers for recruiters and sponsors are linkedin. Extensions can leverage Core inorder to verify the user's identity.

### User Managment
User management provides the ability to create an account, update account info, retrieve users, and mark users. Marks are used in order to indicate if the role of the user. Currently there are marks for basic users, paid users, and recruiters. Extensions can leverage Core in order to verify the mark of a user.

### Group Managment
Group management provides a read only interface to the truth store which is internally managed with git. In addition to directly serving group data, a user's membership to a specified group can be verified. Extensions can leverage Core inorder to verify group membership.

### Resume Managment
Resume management allows students to upload their resumes to the ACM@UIUC resume book. Recruiters are able to view the ACM@UIUC resume book, filter student resumes, and download specific resumes as pdf files.

## API Exposed Routes
Core exposes a set of routes that while primarly for internal use can be used by extensions in order to interact with data stored in Core. Permissions are enforced on these endpoint based on the sensitivity of the information which is being accessed or modified.

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
Core stores all of it's persitant user generated information in a MySQL database using the following models. Each model is stored in it's own table, where model fields corespond to columns.

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
Core stores all of it's static infrequently modified data in git and retreives it based on a predefined TTL policy over HTTPS. This information is stored as yaml which maps to the following models.

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

## Site Overview

## Routes
Core supports a number of routes which return HTML pages that are designed to be rendered in browser. These routes form the ACM@UIUC website. Permissions are enforced on some routes which expose sensitive data.

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
