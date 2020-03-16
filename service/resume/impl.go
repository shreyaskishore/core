package resume

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/acm-uiuc/core/model"
)

type resumeImpl struct {
	db *sqlx.DB
}

func (service *resumeImpl) UploadResume(resume model.Resume) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (service *resumeImpl) GetResumes() ([]model.Resume, error) {
	return nil, fmt.Errorf("not implemented")
}
