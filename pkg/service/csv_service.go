package service

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
)

type CSVService struct {
	repo repository.Predictions
}

func NewCSVService(repo repository.Predictions) *CSVService {
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
	tempFile, err := os.CreateTemp("", "upload-*.csv")
	if err != nil {
		return fmt.Errorf("не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Записываем содержимое файла во временный файл
	if _, err := io.Copy(tempFile, file); err != nil {
		return fmt.Errorf("не удалось записать файл: %v", err)
	}

	// Возвращаемся к началу файла для дальнейшего чтения
	tempFile.Seek(0, 0)

	// Отправляем POST запрос
	url := "http://localhost:5000/upload_csv" // Замените на ваш URL
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", tempFile.Name())
	if err != nil {
		return fmt.Errorf("не удалось создать часть формы: %v", err)
	}

	if _, err := io.Copy(part, tempFile); err != nil {
		return fmt.Errorf("не удалось скопировать файл в часть формы: %v", err)
	}

	// Закрываем writer, чтобы завершить формирование тела запроса
	writer.Close()

	// Выполняем POST запрос
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("не удалось создать запрос: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("получен неожиданный статус ответа: %s", resp.Status)
	}

	err = s.repo.SavePrediction(studentId, resp.Proto) // сохраняем данные в бдшку

	if err == nil {
		return fmt.Errorf("не удалось сохранить Предикшн студента: %v", err)
	}
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
