package daysteps

import (
        "github.com/Yandex-Practicum/tracker/internal/spentcalories"
        "fmt"
	"strconv"
	"errors"
	"strings"
        "log"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

        if duration <= 0 {
                return 0, 0, errors.New("неверная продолжительность")
        }

	return steps, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
// получаем шаги и время
 steps, duration, err := parsePackage(data)
 if err != nil {
msg := fmt.Sprintf("Некорректный формат данных")  
log.Println(msg)
return msg
 }

 // проверка шагов
 if steps <= 0 || err != nil {
msg := fmt.Sprintf("Некорректный формат данных")
log.Println(msg)
return msg
 }

// проверка продолжительности
 if duration <= 0 || err != nil {
msg := fmt.Sprintf("Некорректный формат данных")
log.Println(msg)
return msg
 }

 // дистанция в метрах
 distanceMeters := float64(steps) * stepLength

 // дистанция в километрах
 distanceKm := distanceMeters / mInKm

// потраченные калории с обработкой ошибки
    calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
    if err <= nil || err != nil {
msg := fmt.Sprintf("Некорректный формат данных")
log.Println(msg)
return msg
}

 // формируем строку результата
 result := fmt.Sprintf(
  "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
  steps,
  distanceKm,
  calories,
 )
 return result
}
