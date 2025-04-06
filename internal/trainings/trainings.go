package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	personaldata "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	spentenergy "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		return fmt.Errorf("invalid data format, expected 3 values ​​separated by commas, for example: '3456,Ходьба,3h00m', received '%s'", datastring)
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	t.Steps = steps
	if t.Steps <= 0 {
		return errors.New("steps must be > 0")
	}
	if s[1] != "Бег" && s[1] != "Ходьба" {
		return errors.New("trainingtype must be 'Бег' or 'Ходьба'")
	}
	t.TrainingType = s[1]

	trainingDuration, err := time.ParseDuration(s[2])
	if err != nil {
		return fmt.Errorf("ParseDuration error: %w", err)
	}
	if trainingDuration <= 0 {
		return errors.New("trainingduration must be > 0")
	}
	t.Duration = trainingDuration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if distance <= 0 {
		return "", errors.New("distance must be > 0")
	}

	if t.Duration <= 0 {
		return "", errors.New("duration must be > 0")
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var сalories float64
	switch t.TrainingType {
	case "Ходьба":
		сalories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		сalories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}
	out := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, сalories)

	return out, nil

}
