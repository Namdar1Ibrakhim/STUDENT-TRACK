package repository

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type StudentCourseRepository struct {
	db *sqlx.DB
}

func NewStudentCoursePostgres(db *sqlx.DB) *StudentCourseRepository {
	return &StudentCourseRepository{db: db}
}

func (r *StudentCourseRepository) AddStudentCourse(studentId int, courseId int, grades int) error {
    query := "INSERT INTO student_courses (student_id, course_id, grades) VALUES ($1, $2, $3)"
    _, err := r.db.Exec(query, studentId, courseId, grades)
    return err
}

func (r *StudentCourseRepository) GetAllStudentCourse() ([]dto.StudentCourseResponse, error) {
	query := "SELECT id, student_id, course_id, grades FROM student_course"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourses []dto.StudentCourseResponse
	for rows.Next() {
		var studentCourse dto.StudentCourseResponse
		err := rows.Scan(&studentCourse.Id, &studentCourse.Student_id, &studentCourse.Course_id, &studentCourse.Grades)
		if err != nil {
			return nil, err
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return studentCourses, nil
}

func (r *StudentCourseRepository) GetStudentCourseById(studentCourseId int) (dto.StudentCourseResponse, error) {
	var studentCourse dto.StudentCourseResponse
	query := "SELECT id, student_id, course_id, grades FROM student_course WHERE id = $1"

	err := r.db.QueryRow(query, studentCourseId).Scan(&studentCourse.Id, &studentCourse.Student_id, &studentCourse.Course_id, &studentCourse.Grades)
	if err != nil {
		return studentCourse, err
	}

	return studentCourse, nil
}

func (r *StudentCourseRepository) GetStudentCourseByStudentId(studentId int) (dto.StudentCourseResponse, error) {
	var studentCourse dto.StudentCourseResponse
	query := "SELECT id, student_id, course_id, grades FROM student_course WHERE student_id = $1"

	err := r.db.QueryRow(query, studentId).Scan(&studentCourse.Id, &studentCourse.Student_id, &studentCourse.Course_id, &studentCourse.Grades)
	if err != nil {
		return studentCourse, err
	}

	return studentCourse, nil
}

func (r *StudentCourseRepository) GetStudentCourseByCourseId(courseId int) (dto.StudentCourseResponse, error) {
	var studentCourse dto.StudentCourseResponse
	query := "SELECT id, student_id, course_id, grades FROM student_course WHERE course_id = $1"

	err := r.db.QueryRow(query, courseId).Scan(&studentCourse.Id, &studentCourse.Student_id, &studentCourse.Course_id, &studentCourse.Grades)
	if err != nil {
		return studentCourse, err
	}

	return studentCourse, nil
}

func (r *StudentCourseRepository) GetAllStudentCourseByFilter(limit, page int, sortByGrades *string) ([]dto.StudentCourseResponse, error) {
	// Вычисляем смещение (offset) для пагинации
	offset := (page - 1) * limit

	// Базовый SQL-запрос без сортировки
	query := "SELECT id, student_id, course_id, grades FROM student_course"

	// Добавляем сортировку, если она указана
	if sortByGrades != nil && (*sortByGrades == "asc" || *sortByGrades == "desc") {
		query += " ORDER BY grades " + *sortByGrades
	}

	// Добавляем лимит и смещение для пагинации
	query += " LIMIT $1 OFFSET $2"

	// Выполняем запрос
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourses []dto.StudentCourseResponse
	for rows.Next() {
		var studentCourse dto.StudentCourseResponse
		err := rows.Scan(&studentCourse.Id, &studentCourse.Student_id, &studentCourse.Course_id, &studentCourse.Grades)
		if err != nil {
			return nil, err
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return studentCourses, nil
}

