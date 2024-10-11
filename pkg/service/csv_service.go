package service

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/spf13/viper"

	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type CSVService struct {
	repo  repository.Predictions
	repo2 repository.Course
	repo3 repository.StudentCourse
	repo4 repository.Direction
}

func NewCSVService(repo repository.Predictions, repo2 repository.Course, repo3 repository.StudentCourse, repo4 repository.Direction) *CSVService {
	return &CSVService{
		repo:  repo,
		repo2: repo2,
		repo3: repo3,
		repo4: repo4,
	}
}
func parseInt(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func (s *CSVService) ValidateCSV(file io.Reader) error {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return errors.New("invalid CSV structure")
	}

	expectedHeaders := []string{
		"Operating System", "Analysis of Algorithm", "Programming Concept", "Software Engineering",
		"Computer Network", "Applied Mathematics", "Computer Security", "Hackathons attended",
		"Topmost Certification", "Personality", "Management or technical", "Leadership", "Team", "Self Ability"}

	if len(records) == 0 || !s.equalHeaders(records[0], expectedHeaders) {
		return errors.New("invalid CSV structure, expected columns: check the required format")
	}

	for i, row := range records[1:] {
		if len(row) != len(expectedHeaders) {
			return fmt.Errorf("invalid number of columns at row %d", i+2)
		}

		for j := 0; j <= 6; j++ {
			gradeStr := row[j]
			if gradeStr == "" {
				return fmt.Errorf("missing grade for subject at row %d, column %d", i+2, j+1)
			}

			grade, err := strconv.Atoi(gradeStr)
			if err != nil || grade < 0 || grade > 100 {
				return fmt.Errorf("invalid grade value at row %d, column %d, must be between 0 and 100", i+2, j+1)
			}
		}

		hackathonsStr := row[7]
		if hackathonsStr == "" {
			return fmt.Errorf("missing value for 'hackathons' at row %d", i+2)
		}
		hackathons, err := strconv.Atoi(hackathonsStr)
		if err != nil || hackathons < 0 {
			return fmt.Errorf("invalid 'Hackathons attended' value at row %d, must be a non-negative integer", i+2)
		}

		for j := 8; j < 14; j++ {
			if row[j] == "" {
				return fmt.Errorf("missing value for '%s' at row %d", expectedHeaders[j], i+2)
			}
		}
	}

	return nil
}

func (s *CSVService) PredictCSV(studentId int, file io.Reader) (*dto.PredictionResponseDto, error) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("invalid CSV structure")
	}

	row := records[1]

	courses := map[string]int{
		"Operating System":      parseInt(row[0]),
		"Analysis of Algorithm": parseInt(row[1]),
		"Programming Concept":   parseInt(row[2]),
		"Software Engineering":  parseInt(row[3]),
		"Computer Network":      parseInt(row[4]),
		"Applied Mathematics":   parseInt(row[5]),
		"Computer Security":     parseInt(row[6]),
	}

	for courseName, grade := range courses {
		courseID, err := s.repo2.FindCourseIDByName(courseName)
		if err != nil {
			return nil, fmt.Errorf("failed to find course ID for %s: %v", courseName, err)
		}

		err = s.repo3.AddStudentCourse(studentId, courseID, grade)
		if err != nil {
			return nil, fmt.Errorf("failed to save course data for student: %v", err)
		}
	}

	predictionRequest := dto.PredictionDataOfCSVRequest{
		OperatingSystem:      parseInt(row[0]),
		AnalysisOfAlgorithm:  parseInt(row[1]),
		ProgrammingConcept:   parseInt(row[2]),
		SoftwareEngineering:  parseInt(row[3]),
		ComputerNetwork:      parseInt(row[4]),
		AppliedMathematics:   parseInt(row[5]),
		ComputerSecurity:     parseInt(row[6]),
		HackathonsAttended:   parseInt(row[7]),
		TopmostCertification: row[8],
		Personality:          row[9],
		ManagementTechnical:  row[10],
		Leadership:           row[11],
		Team:                 row[12],
		SelfAbility:          row[13],
	}

	jsonData, err := json.Marshal(predictionRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(viper.GetString("url"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ML service status -->: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.New("failed to decode ML service response")
	}

	prediction, ok := result["predicted_track"].(string)
	if !ok {
		return nil, errors.New("invalid prediction format")
	}

<<<<<<< HEAD
	//******//

	err = s.repo.SavePrediction(studentId, prediction)
=======
	directionID, err := s.repo4.FindDirectionIDByName(prediction)
>>>>>>> e1e18e5e99ee210f33fd65ee2b1bb3d695728391
	if err != nil {
		return nil, fmt.Errorf("failed to find direction ID for predicted track: %v", err)
	}

<<<<<<< HEAD
	//******//

	return prediction, nil
=======
	err = s.repo.SavePrediction(studentId, directionID)
	if err != nil {
		return nil, errors.New("failed to save prediction")
	}

	response := &dto.PredictionResponseDto{
		PredictedTrack: prediction,
		StudentId:      studentId,
	}

	return response, nil
>>>>>>> e1e18e5e99ee210f33fd65ee2b1bb3d695728391
}

func (s *CSVService) equalHeaders(headers, expectedHeaders []string) bool {
	if len(headers) != len(expectedHeaders) {
		return false
	}

	for i, header := range headers {
		if strings.TrimSpace(strings.ToLower(header)) != strings.TrimSpace(strings.ToLower(expectedHeaders[i])) {
			return false
		}
	}
	return true
}
