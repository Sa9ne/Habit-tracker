package main

import (
	"Habit-tracker/internal/models"
	"Habit-tracker/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	apps := app.New()
	apps.Settings().SetTheme(theme.DarkTheme())

	MainWindow := apps.NewWindow("Трекер привычек")
	MainWindow.Resize(fyne.NewSize(650, 400))
	MainWindow.SetFixedSize(true)
	MainWindow.CenterOnScreen()

	// Загрузка иконок
	icons := ui.LoadIcons()
	MainWindow.SetIcon(icons.MainIcon)

	// Работа с базой данных
	var habits []models.Habits
	DB, _ := gorm.Open(sqlite.Open("Habits.db"), &gorm.Config{})
	DB.AutoMigrate(&models.Habits{})
	DB.Find(&habits)

	// Запуск главного экрана
	MainCont := ui.NewMainWindow(MainWindow, habits, DB, icons)
	MainWindow.SetContent(MainCont)

	MainWindow.Show()
	apps.Run()
}
