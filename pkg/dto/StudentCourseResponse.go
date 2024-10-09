package dto

type StudentCourseResponse struct {
    Id            int     `json:"-" db:"id"`                   
	Student_id    int     `json:"student_id" binding:"required"`
	Course_id     int     `json:"course_id"  binding:"required"`
	Grades        int     `json:"grades"     binding:"required"`
}
