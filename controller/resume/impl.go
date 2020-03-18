package resume

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service"
)

type ResumeController struct {
	svc *service.Service
}

func New(svc *service.Service) *ResumeController {
	return &ResumeController{
		svc: svc,
	}
}

func (controller *ResumeController) UploadResume(ctx *context.Context) error {
	req := model.Resume{}

	err := ctx.Bind(&req)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed Bind")
		return err
	}

	req.BlobKey = req.Username
	req.Approved = false

	uri, err := controller.svc.Resume.UploadResume(req)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed Resume Upload")
		return err
	}

	return ctx.JSON(http.StatusOK, &struct {
		UploadUri string `json:"upload_uri"`
	}{
		UploadUri: uri,
	})
}

func (controller *ResumeController) GetResumes(ctx *context.Context) error {
	resumes, err := controller.svc.Resume.GetResumes()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed Resumes Lookup")
		return err
	}

	return ctx.JSON(http.StatusOK, resumes)
}

func (controller *ResumeController) ApproveResume(ctx *context.Context) error {
	req := model.Resume{}

	err := ctx.Bind(&req)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed Bind")
		return err
	}

	err = controller.svc.Resume.ApproveResume(req.Username)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed Resume Approval")
		return err
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}
