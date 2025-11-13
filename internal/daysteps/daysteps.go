package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               //встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

// Parse парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	inData := strings.Split(datastring, ",")
	if len(inData) != 2 {
		return errors.New("not enough data")
	}
	countStep, err := strconv.Atoi(inData[0])
	if err != nil {
		return fmt.Errorf("incorrect conversion countStep: %w", err)
	}
	if countStep <= 0 {
		return errors.New("count steps must be greater than zero")
	}
	ds.Steps = countStep

	duration, err := time.ParseDuration(inData[1])

	if err != nil {
		return fmt.Errorf("incorrect conversion duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration  must be greater than zero")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	if err != nil {
		return "", err
	}
	outString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return outString, nil
}
