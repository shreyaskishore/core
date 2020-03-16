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
	err := service.validateResume(&resume)
	if err != nil {
		return "", fmt.Errorf("failed to create resume: %w", err)
	}

	err = service.addResume(&resume)
	if err != nil {
		return "", fmt.Errorf("failed to create resume: %w", err)
	}

	signedUri, err := service.getSignedUri(resume.BlobKey, "POST")
	if err != nil {
		return "", fmt.Errorf("failed to create signed uri: %w", err)
	}

	return signedUri, nil
}

func (service *resumeImpl) GetResumes() ([]model.Resume, error) {
	resumes, err := service.getResumes()
	if err != nil {
		return nil, fmt.Errorf("failed to get resumes: %w", err)
	}

	return resumes, nil
}

func (service *resumeImpl) validateResume(resume *model.Resume) error {
	// TODO: Enforce validation on resumes
	return nil
}

func (service *resumeImpl) addResume(resume *model.Resume) error {
	_, err := service.db.NamedExec("INSERT INTO resumes (username, first_name, last_name, email, graduation_month, graduation_year, major, degree, seeking, blob_key, approved) VALUES (:username, :first_name, :last_name, :email, :graduation_month, :graduation_year, :major, :degree, :seeking, :blob_key, :approved) ON DUPLICATE KEY UPDATE username = :username, first_name = :first_name, last_name = :last_name, email = :email, graduation_month = :graduation_month, graduation_year = :graduation_year, major = :major, degree = :degree, seeking = :seeking, blob_key = :blob_key, approved = :approved", resume)
	if err != nil {
		fmt.Errorf("failed to add resume to database: %w", err)
	}

	return nil
}

func (service *resumeImpl) getResumes() ([]model.Resume, error) {
	rows, err := service.db.NamedQuery("SELECT * FROM resumes", struct{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to query database for resumes: %w", err)
	}

	results := []model.Resume{}
	for rows.Next() {
		result := model.Resume{}
		err := rows.StructScan(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode row from database: %w", err)
		}
		results = append(results, result)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed reading rows from database: %w", err)
	}

	return results, nil
}

func (service *resumeImpl) getSignedUri(blob_key string, method string) (string, error) {
	// TODO: Implement signed uri
	return "https://not.implemented", nil
}
