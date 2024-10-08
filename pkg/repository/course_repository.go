package repository

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetAll() ([]dto.CourseResponse, error) {
	query := "SELECT id, course_name, description FROM course"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []dto.CourseResponse
	for rows.Next() {
		var course dto.CourseResponse
		err := rows.Scan(&course.Id, &course.Course_name, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) GetById(courseId int) (dto.CourseResponse, error) {
	var course dto.CourseResponse
	query := "SELECT id, course_name, description FROM course WHERE id = $1"

	err := r.db.QueryRow(query, courseId).Scan(&course.Id, &course.Course_name, &course.Description)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *CourseRepository) GetByName(courseName string) (dto.CourseResponse, error) {
	var course dto.CourseResponse
	query := "SELECT id, course_name, description FROM course WHERE course_name = $1"

	err := r.db.QueryRow(query, courseName).Scan(&course.Id, &course.Course_name, &course.Description)
	if err != nil {
		return course, err
	}

	return course, nil
}
