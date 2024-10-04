package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"io"
	"strconv"
	"strings"
)

type CSVService struct {
	repo repository.CSV
}

func NewCSVService(repo repository.CSV) *CSVService {
	return &CSVService{repo: repo}
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

func (s *CSVService) ProcessCSV(studentId int, file io.Reader) error {
	//////....
	return nil
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
