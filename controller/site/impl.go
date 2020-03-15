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
