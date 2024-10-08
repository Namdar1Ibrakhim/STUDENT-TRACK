package dto

type CourseResponse struct {
    Id            int     `json:"-" db:"id"`                   
	Course_name   string  `json:"coursename" binding:"required"`
	Description   string  `json:"description" binding:"required"`

}