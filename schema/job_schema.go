package schema

import "advanced_programming/models"

type CampaignJobRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Company string `json:"company"`
	Base    string `json:"base"`
	Content string `json:"content"`
}

type CampaignJobResponse struct {
	Job *models.Job `json:"job"`
	Msg string      `json:"msg"`
}

type GetJobListRequest struct {
	Limit int
}
type GetJobListResponse struct {
	Msg       string        `json:"msg"`
	Jobs      []*models.Job `json:"jobs"`
	UserNames []string      `json:"user_names"`
}

type GetJobRequest struct {
	JobID int `json:"job_id"`
}

type GetJobResponse struct {
	Msg      string      `json:"msg"`
	UserName string      `json:"user_name"`
	Job      *models.Job `json:"job"`
}

type UpdateJobRequest struct {
	UserID  int    `json:"user_id"`
	JobID   int    `json:"job_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateJobResponse struct {
	Msg string `json:"msg"`
}

type DeleteJobRequest struct {
	JobID int `json:"job_id"`
}

type DeleteJobResponse struct {
	Msg string `json:"msg"`
}

type ApplyRequest struct {
	UserID int `json:"user_id"`
	JobID  int `json:"job_id"`
}

type ApplyResponse struct {
	Msg string `json:"msg"`
}
