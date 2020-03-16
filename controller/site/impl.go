package site

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"
)

type SiteController struct {
	svc *service.Service
}

func New(svc *service.Service) *SiteController {
	return &SiteController{
		svc: svc,
	}
}

func (controller *SiteController) Home(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "home", params)
}

func (controller *SiteController) About(ctx *context.Context) error {
	groups, err := controller.svc.Group.GetGroups()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting Groups")
	}

	committees, ok := groups[model.GroupCommittees]
	if !ok {
		return ctx.String(http.StatusBadRequest, "Failed Getting Committees")
	}

	params := struct {
		Authenticated bool
		Committees    []model.Group
	}{
		Authenticated: ctx.LoggedIn,
		Committees:    committees,
	}

	return ctx.Render(http.StatusOK, "about", params)
}

func (controller *SiteController) History(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "history", params)
}

func (controller *SiteController) ReflectionsProjections(ctx *context.Context) error {
	params := struct {
		Authenticated bool
		Editions      []struct {
			Year int
			Uri  string
		}
	}{
		Authenticated: ctx.LoggedIn,
		Editions: []struct {
			Year int
			Uri  string
		}{
			{
				Year: 2020,
				Uri:  "https://2020.reflectionsprojections.org",
			},
			{
				Year: 2019,
				Uri:  "https://2019.reflectionsprojections.org",
			},
			{
				Year: 2018,
				Uri:  "https://reflectionsprojections.github.io/rp2018",
			},
			{
				Year: 2017,
				Uri:  "https://reflectionsprojections.github.io/rp2017",
			},
			{
				Year: 2016,
				Uri:  "https://reflectionsprojections.github.io/rp2016",
			},
			{
				Year: 2015,
				Uri:  "https://reflectionsprojections.github.io/rp2015",
			},
			{
				Year: 2014,
				Uri:  "https://reflectionsprojections.github.io/rp2014",
			},
			{
				Year: 2013,
				Uri:  "https://reflectionsprojections.github.io/rp2013",
			},
			{
				Year: 2012,
				Uri:  "https://reflectionsprojections.github.io/rp2012",
			},
			{
				Year: 2011,
				Uri:  "https://reflectionsprojections.github.io/rp2011",
			},
			{
				Year: 2010,
				Uri:  "https://reflectionsprojections.github.io/rp2010",
			},
			{
				Year: 2009,
				Uri:  "https://reflectionsprojections.github.io/rp2009",
			},
			{
				Year: 2008,
				Uri:  "https://reflectionsprojections.github.io/rp2008",
			},
			{
				Year: 2007,
				Uri:  "https://reflectionsprojections.github.io/rp2007",
			},
			{
				Year: 2006,
				Uri:  "https://reflectionsprojections.github.io/rp2006",
			},
			{
				Year: 2005,
				Uri:  "https://reflectionsprojections.github.io/rp2005",
			},
			{
				Year: 2004,
				Uri:  "https://reflectionsprojections.github.io/rp2004",
			},
			{
				Year: 2003,
				Uri:  "https://reflectionsprojections.github.io/rp2003",
			},
			{
				Year: 2002,
				Uri:  "https://reflectionsprojections.github.io/rp2002",
			},
			{
				Year: 2001,
				Uri:  "https://reflectionsprojections.github.io/rp2001",
			},
			{
				Year: 2000,
				Uri:  "https://reflectionsprojections.github.io/rp2000",
			},
			{
				Year: 1999,
				Uri:  "https://reflectionsprojections.github.io/rp1999",
			},
			{
				Year: 1998,
				Uri:  "https://reflectionsprojections.github.io/rp1998",
			},
			{
				Year: 1997,
				Uri:  "https://reflectionsprojections.github.io/rp1997",
			},
			{
				Year: 1996,
				Uri:  "https://reflectionsprojections.github.io/rp1996",
			},
			{
				Year: 1995,
				Uri:  "https://reflectionsprojections.github.io/rp1995",
			},
		},
	}

	return ctx.Render(http.StatusOK, "reflectionsprojections", params)
}

func (controller *SiteController) HackIllinois(ctx *context.Context) error {
	params := struct {
		Authenticated bool
		Editions      []struct {
			Year int
			Uri  string
		}
	}{
		Authenticated: ctx.LoggedIn,
		Editions: []struct {
			Year int
			Uri  string
		}{
			{
				Year: 2020,
				Uri:  "https://2020.hackillinois.org",
			},
			{
				Year: 2019,
				Uri:  "https://2019.hackillinois.org",
			},
			{
				Year: 2018,
				Uri:  "https://2018.hackillinois.org",
			},
			{
				Year: 2017,
				Uri:  "https://2017.hackillinois.org",
			},
			{
				Year: 2016,
				Uri:  "https://2016.hackillinois.org",
			},
			{
				Year: 2015,
				Uri:  "https://2015.hackillinois.org",
			},
			{
				Year: 2014,
				Uri:  "https://2014.hackillinois.org",
			},
		},
	}

	return ctx.Render(http.StatusOK, "hackillinois", params)
}

func (controller *SiteController) Sponsors(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "sponsors", params)
}

func (controller *SiteController) Sigs(ctx *context.Context) error {
	groups, err := controller.svc.Group.GetGroups()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting Groups")
	}

	sigs, ok := groups[model.GroupSIGs]
	if !ok {
		return ctx.String(http.StatusBadRequest, "Failed Getting Sigs")
	}

	sigsColLeft := sigs[:len(sigs)/2]
	sigsColRight := sigs[len(sigs)/2:]

	params := struct {
		Authenticated bool
		SigsColLeft   []model.Group
		SigsColRight  []model.Group
	}{
		Authenticated: ctx.LoggedIn,
		SigsColLeft:   sigsColLeft,
		SigsColRight:  sigsColRight,
	}

	return ctx.Render(http.StatusOK, "sigs", params)
}

func (controller *SiteController) Login(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "login", params)
}

func (controller *SiteController) Logout(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})

	return ctx.Render(http.StatusOK, "logout", params)
}

func (controller *SiteController) Join(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "join", params)
}

func (controller *SiteController) ResumeUpload(ctx *context.Context) error {
	params := struct {
		Authenticated    bool
		GraduationMonths []int
		GraduationYears  []int
		Degrees          []string
		Seekings         []string
		Majors           []string
	}{
		Authenticated:    ctx.LoggedIn,
		GraduationMonths: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		GraduationYears:  []int{2020, 2021, 2022, 2023, 2024, 2025},
		Degrees:          []string{"Bachlors", "Masters", "PhD"},
		Seekings:         []string{"Internship", "Co Op", "Full Time"},
		Majors:           []string{"Computer Science", "Computer Engineering", "Electrical Enginering", "Mathematics", "Other Engineering", "Other Sciences", "Other"},
	}

	return ctx.Render(http.StatusOK, "resumeupload", params)
}

func (controller *SiteController) UserManager(ctx *context.Context) error {
	users, err := controller.svc.User.GetUsers()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting Users")
	}

	params := struct {
		Authenticated bool
		Users         []model.User
	}{
		Authenticated: ctx.LoggedIn,
		Users:         users,
	}

	return ctx.Render(http.StatusOK, "usermanager", params)
}

func (controller *SiteController) RecruiterCreator(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusOK, "recruitercreator", params)
}

func (controller *SiteController) RecruiterManager(ctx *context.Context) error {
	users, err := controller.svc.User.GetUsers()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting Users")
	}

	// TODO: Use filtering on GetUsers() instead once implemented
	recruiters := []model.User{}
	for _, user := range users {
		if user.Mark == model.UserMarkRecruiter {
			recruiters = append(recruiters, user)
		}
	}

	params := struct {
		Authenticated bool
		Users         []model.User
	}{
		Authenticated: ctx.LoggedIn,
		Users:         recruiters,
	}

	return ctx.Render(http.StatusOK, "recruitermanager", params)
}

func (controller *SiteController) Intranet(ctx *context.Context) error {
	roles := []string{}

	marksToRole := map[string]string{
		model.UserMarkBasic:     "Basic Member",
		model.UserMarkPaid:      "Paid Member",
		model.UserMarkRecruiter: "Recruiter",
	}

	user, err := controller.svc.User.GetUser(ctx.Username)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting User")
	}

	markRole, ok := marksToRole[user.Mark]
	if !ok {
		return ctx.String(http.StatusBadRequest, "Invalid User Mark")
	}
	roles = append(roles, markRole)

	isTop4, err := controller.svc.Group.VerifyMembership(ctx.Username, model.GroupCommittees, model.GroupTop4)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Membership Verification")
	}
	if isTop4 {
		roles = append(roles, "Top4")
	}

	cards := []struct {
		Title       string
		Description string
		Uri         string
	}{}

	if isTop4 {
		cards = append(cards, struct {
			Title       string
			Description string
			Uri         string
		}{
			Title:       "User Manager",
			Description: "Manage ACM@UIUC's users",
			Uri:         "/intranet/usermanager",
		})
	}

	if isTop4 {
		cards = append(cards, struct {
			Title       string
			Description string
			Uri         string
		}{
			Title:       "Recruiter Manager",
			Description: "Manage ACM@UIUC's recruiters",
			Uri:         "/intranet/recruitermanager",
		})
	}

	params := struct {
		Authenticated bool
		Username      string
		Roles         []string
		Cards         []struct {
			Title       string
			Description string
			Uri         string
		}
	}{
		Authenticated: ctx.LoggedIn,
		Username:      ctx.Username,
		Roles:         roles,
		Cards:         cards,
	}

	return ctx.Render(http.StatusOK, "intranet", params)
}

func (controller *SiteController) ResumeBook(ctx *context.Context) error {
	resumes, err := controller.svc.Resume.GetResumes()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Getting Resumes")
	}

	params := struct {
		Authenticated bool
		Resumes       []model.Resume
	}{
		Authenticated: ctx.LoggedIn,
		Resumes:       resumes,
	}

	return ctx.Render(http.StatusOK, "resumebook", params)
}

func (controller *SiteController) NotFound(ctx *context.Context) error {
	params := struct {
		Authenticated bool
	}{
		Authenticated: ctx.LoggedIn,
	}

	return ctx.Render(http.StatusNotFound, "notfound", params)
}
