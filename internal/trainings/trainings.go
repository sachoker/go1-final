package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Создайте экспортируемую структуру Training, которая содержит следующие поля:

// Steps int — количество шагов, проделанных за тренировку.
// TrainingType string — тип тренировки(бег или ходьба).
// Duration time.Duration — длительность тренировки.
// personaldata.Personal — встроенная структура Personal из пакета personaldata, у которой есть метод Print().

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, " ")

	if len(parts) != 3 {
		return errors.New("invalid string format: expected 3 elements (steps, type, duration)")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error parsing steps: %w", err)
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpead := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calcErr error
	var calories float64
	switch t.TrainingType {
	case "Бег":
		calories, calcErr = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, calcErr = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("wrong training type")
	}
	if calcErr != nil {
		return "", fmt.Errorf("calories calc error: %w", calcErr)
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpead,
		calories,
	)
	return result, nil
}
