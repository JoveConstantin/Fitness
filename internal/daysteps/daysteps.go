package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	personaldata "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	spentenergy "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	row := strings.Split(datastring, ",")
	if len(row) != 2 {
		return fmt.Errorf("invalid data format, expected 2 values ​​separated by comma, for example: '678,0h50m', received '%s'", datastring)
	}
	steps, err := strconv.Atoi(row[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be > 0")
	}
	ds.Steps = steps
	walkLenth, err := time.ParseDuration(row[1])
	if err != nil {
		return fmt.Errorf("parseDuration error: %w", err)
	}
	if walkLenth <= 0 {
		return errors.New("walk <= 0 seconds")
	}
	ds.Duration = walkLenth
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("duration must be > 0")
	}
	distance := spentenergy.Distance(ds.Steps)
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	out := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.",
		ds.Steps, distance, calories,
	)
	return out, nil
}
