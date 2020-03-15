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
				Year: 2019,
				Uri:  "https://2019.reflectionsprojections.org",
			},
		},
	}

	return ctx.Render(http.StatusOK, "reflectionsprojections", params)
}
