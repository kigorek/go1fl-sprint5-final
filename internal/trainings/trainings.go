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

type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку.
	TrainingType          string        // тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	inData := strings.Split(datastring, ",")
	if len(inData) != 3 {
		return errors.New("not enough data")
	}
	countStep, err := strconv.Atoi(inData[0])
	if err != nil {
		return fmt.Errorf("incorrect conversion countStep: %w", err)
	}
	if countStep <= 0 {
		return errors.New("count staeps must be greater than zero")
	}
	t.Steps = countStep
	t.TrainingType = inData[1]

	duration, err := time.ParseDuration(inData[2])

	if err != nil {
		return fmt.Errorf("incorrect conversion duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration  must be greater than zero")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
// если тип тренировки не подходит, то ошибка
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("unknown type of training")
	}

	if err != nil {
		return "", err
	}

	outString := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return outString, nil

}
