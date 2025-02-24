package ui

import (
	"Habit-tracker/internal/models"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

func NewWindowAdd(MainWindow fyne.Window, DB *gorm.DB, icons *Icons, NoHabitText *canvas.Text) *fyne.Container {

	// Текст заголовка
	CreateTextHabits := canvas.NewText("Внесите информацию о вашей привычке", color.White)
	CreateTextHabits.TextSize = 25
	CreateTextHabits.Move(fyne.NewPos(0, 10))

	// Линия
	Line := canvas.NewLine(color.White)
	Line.StrokeWidth = 2.5
	Line.Position1 = fyne.NewPos(0, 65)
	Line.Position2 = fyne.NewPos(640, 65)

	// Поля ввода
	HabitNameEntry := widget.NewEntry()
	HabitNameEntry.SetPlaceHolder("Название привычки...")
	HabitNameEntry.Resize(fyne.NewSize(640, 50))
	HabitNameEntry.Move(fyne.NewPos(0, 80))

	DifficultiLevelEntry := widget.NewEntry()
	DifficultiLevelEntry.SetPlaceHolder("Степень важности (1-10)...")
	DifficultiLevelEntry.Resize(fyne.NewSize(640, 50))
	DifficultiLevelEntry.Move(fyne.NewPos(0, 135))

	MoreInfoEntry := widget.NewMultiLineEntry()
	MoreInfoEntry.SetPlaceHolder("Дополнительная информация...")
	MoreInfoEntry.Resize(fyne.NewSize(640, 145))
	MoreInfoEntry.Move(fyne.NewPos(0, 190))

	// Кнопка возврата
	ButtonBackOnMenu := widget.NewButton("", func() {
		var habits []models.Habits
		DB.Find(&habits)
		MainWindow.SetContent(NewMainWindow(MainWindow, habits, DB, icons))

		// Обновление полей ввода
		HabitNameEntry.Text = ""
		DifficultiLevelEntry.Text = ""
		MoreInfoEntry.Text = ""
	})
	ButtonBackOnMenu.SetIcon(icons.BackIcon)
	ButtonBackOnMenu.Resize(fyne.NewSize(50, 50))
	ButtonBackOnMenu.Move(fyne.NewPos(580, 5))

	// Кнопка сохранения
	SaveHabitsButton := widget.NewButton("Сохранить привычку", func() {

		// Обработка написания названия привычки
		errorLabelName := widget.NewLabel("Название привычки не должно быть пустым!")
		errorLabelName.TextStyle = fyne.TextStyle{Bold: true}
		if HabitNameEntry.Text == "" {
			dialog.NewCustom("Ошибка названия", "Закрыть", errorLabelName, MainWindow).Show()
			return
		}

		// Обработка написания уровня важности
		difficultiLevel, err := strconv.Atoi(DifficultiLevelEntry.Text)
		errorLabelDifficult := widget.NewLabel("Степень важности должна быть числом от 1 до 10!")
		errorLabelDifficult.TextStyle = fyne.TextStyle{Bold: true}
		if err != nil && difficultiLevel <= 1 || difficultiLevel > 11 || difficultiLevel < 1 {
			// Обработка ошибки, если введено не число
			dialog.NewCustom("Ошибка уровня важности", "Закрыть", errorLabelDifficult, MainWindow).Show()
			return
		}

		habit := models.Habits{
			HabitName:       HabitNameEntry.Text,
			DifficultiLevel: int(difficultiLevel),
			MoreInfo:        MoreInfoEntry.Text,
		}

		// Очищаем поля
		HabitNameEntry.SetText("")
		DifficultiLevelEntry.SetText("")
		MoreInfoEntry.SetText("")

		// Создаем новую привычку в базе данных
		DB.Create(&habit)

		// Обновляем список привычек из базы
		var habits []models.Habits
		DB.Find(&habits)

		// Обновляем текст "Нет привычек" в зависимости от наличия привычек
		if len(habits) == 0 {
			NoHabitText.Show()
		} else {
			NoHabitText.Hide()
		}

		// Обновляем окно с новыми привычками
		MainWindow.SetContent(NewMainWindow(MainWindow, habits, DB, icons))
	})

	SaveHabitsButton.SetIcon(icons.SaveIcon)
	SaveHabitsButton.Resize(fyne.NewSize(640, 50))
	SaveHabitsButton.Move(fyne.NewPos(0, 340))

	// Контейнер окна добавления привычки
	WindowAdd := container.NewWithoutLayout(
		HabitNameEntry,
		DifficultiLevelEntry,
		MoreInfoEntry,
		CreateTextHabits,
		Line,
		ButtonBackOnMenu,
		SaveHabitsButton,
	)

	return WindowAdd
}
