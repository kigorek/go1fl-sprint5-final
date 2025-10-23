package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories производит расчет количества калорий, потраченных при ходьбе:
// возвращает два значения:
// - количество калорий, потраченных при ходьбе.
// - ошибку, если входные параметры некорректны
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, errors.New("steps must be greater than or zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}

// RunningSpentCalories производит расчет количества калорий, потраченных при беге:
// возвращает два значения:
// - количество калорий, потраченных при беге.
// - ошибку, если входные параметры некорректны
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, errors.New("steps must be greater than or zero")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than zero")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

// MeanSpeed принимает количество шагов steps, рост пользователя height и продолжительность активности duration
//
//	и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	meanSpeed := Distance(steps, height) / duration.Hours()
	return meanSpeed
}

// Distance принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	distance := (height * stepLengthCoefficient * float64(steps)) / mInKm
	return distance
}
