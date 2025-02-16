package main

import (
	"Habit-tracker/internal/models"
	"Habit-tracker/internal/ui"
	"fmt"
	"os"
	"path/filepath"

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

	// Получаем путь к папке "Документы" пользователя
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка получения домашней директории:", err)
		return
	}

	dbPath := filepath.Join(homeDir, "Documents", "Habit-tracker")
	if err := os.MkdirAll(dbPath, os.ModePerm); err != nil {
		fmt.Println("Ошибка создания директории:", err)
		return
	}

	dbFile := filepath.Join(dbPath, "Habits.db")

	// Работа с базой данных
	var habits []models.Habits
	DB, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}

	DB.AutoMigrate(&models.Habits{})
	DB.Find(&habits)

	// Запуск главного экрана
	MainCont := ui.NewMainWindow(MainWindow, habits, DB, icons)
	MainWindow.SetContent(MainCont)

	MainWindow.Show()
	apps.Run()
}
