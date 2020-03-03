auth
	login(username, password) -> Token, error
	logout(token) -> error
	verify(token) -> bool, error
	createLocalAccount(username, password) -> error
users
	getInfo(netid) -> UserData, error
	create(data) -> UserData, error
	markPaid(netid) -> error
	getInfos(filter) -> UserData[], error
groups
	getMemberships(netid) -> UserMemberships
	getGroups(groupType) -> Group[]
	verifyMembership(netid, group) -> bool, error
web
	index() -> HTML
	about() -> HTML
	sigs() -> HTML
	conference() -> HTML
	hackillinois() -> HTMl
	sponsors() -> HTML
	join() -> HTML
	intranet() -> HTML
	resumebook() -> HTML
	resumeupload() -> HTML


Token {
	netid
	token
}
LocalAccount {
	email
	hashedPassword
}
UserData {
	netid
	name
	graduationYear
	major
	resume
	memberType oneof[basic, paid, recruiter]
}
UserMemberships {
	netid
	memberships[]
}
Group {
	name
	chairs
	meetingTime
	meetingRoom
}


The auth module will act as an interface to the universities ldap server allowing us to verify a students login information. The auth module also allows for local account to be created that are not backed by ldap and are instead back by basic login info stored in a table of the database. When login fails to sucessfully complete an ldap login it will fall back to logging in with a local account. Local accounts are primarly for recruiters.

The user module will store basic information about the user which can be gathered from the university's ldap servers or our sign up form. When a student first signs up they will have a basic account. They can later be marked as paid which will change their membership type and allow them more access. The user module also exposes the ability to get all users which conform to a filer. This will be used by the recruiter portal in order to choose which resumes to display.

The groups module will be an interface to the truth store of group data. This truth store will be stored directly on github and we automatically pulled into the service. As a result the service will not need to be redeployed in order to update group information.

The web module is the frontend for the website and will provide users an interface for interacting with our infrastructure.
