package model

type SubjectStar struct {
	SubjectId uint64 `json:"subject_id"`
	StarId    uint64 `json:"star_id"`
	TimestampModel
}
