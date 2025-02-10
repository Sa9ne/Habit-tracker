package ui

import (
	"Habit-tracker/internal/models"
	"image/color"
	"sort"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

func NewMainWindow(MainWindow fyne.Window, habits []models.Habits, DB *gorm.DB, icons *Icons) *fyne.Container {
	var taskshabits *widget.List
	var NoHabitText *canvas.Text

	// Сортировка привычек по уровню сложности (от максимального к минимальному)
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].DifficultiLevel > habits[j].DifficultiLevel
	})

	// Кнопка добавления привычек
	ButtonHabitAdd := widget.NewButton("", func() {
		MainWindow.SetContent(NewWindowAdd(MainWindow, DB, icons, NoHabitText))
	})
	ButtonHabitAdd.SetIcon(icons.AddIcon)
	ButtonHabitAdd.Resize(fyne.NewSize(50, 50))
	ButtonHabitAdd.Move(fyne.NewPos(580, 5))

	// Текст "Ваши привычки"
	TextHabits := canvas.NewText("Ваши привычки:", color.White)
	TextHabits.TextSize = 40

	// Настройки линии
	Line := canvas.NewLine(color.White)
	Line.StrokeWidth = 2.5
	Line.Position1 = fyne.NewPos(0, 65)
	Line.Position2 = fyne.NewPos(640, 65)

	// Текст, если привычек нет
	NoHabitText = canvas.NewText("Вы еще не внесли привычки", color.White)
	NoHabitText.TextSize = 30
	NoHabitText.Move(fyne.NewPos(0, 70))
	if len(habits) != 0 {
		NoHabitText.Hide()
	}

	// Вывод информации о привычках
	taskshabits = widget.NewList(
		func() int {
			return len(habits)
		},
		func() fyne.CanvasObject {

			// Текст названия привычки
			LabelHabit := canvas.NewText("", color.White)
			LabelHabit.TextStyle = fyne.TextStyle{Bold: true}
			LabelHabit.TextSize = 20
			LabelHabit.Move(fyne.NewPos(35, 2.5))

			// Настройка уровня сложности
			DifficultiLVLIMG := canvas.NewImageFromResource(nil)
			DifficultiLVLIMG.Move(fyne.NewPos(2.5, 2.5))
			DifficultiLVLIMG.Resize(fyne.NewSize(30, 30))

			// Счетчик выполненных задач
			CompletedTasksText := canvas.NewText("Выполнено: 0", color.White)
			CompletedTasksText.TextSize = 20
			CompletedTasksText.Move(fyne.NewPos(300, 2.5)) // Размещаем счетчик рядом с привычкой

			// Кнопка информации
			ButtonInfo := widget.NewButtonWithIcon("", icons.InfoIcon, nil)
			ButtonInfo.Move(fyne.NewPos(565, 2.5))
			ButtonInfo.Resize(fyne.NewSize(30, 30))

			// Кнопка удаления
			ButtonDelete := widget.NewButtonWithIcon("", icons.DeleteIcon, nil)
			ButtonDelete.Move(fyne.NewPos(600, 2.5))
			ButtonDelete.Resize(fyne.NewSize(30, 30))

			// Кнопка готовности
			ButtonReadyCheck := widget.NewButtonWithIcon("", icons.ReadyIcon, nil)
			ButtonReadyCheck.Move(fyne.NewPos(530, 2.5))
			ButtonReadyCheck.Resize(fyne.NewSize(30, 30))

			//Огонь (Горит)
			IMGFireStreak := canvas.NewImageFromResource(icons.FireStreak)
			IMGFireStreak.Move(fyne.NewPos(460, 0))
			IMGFireStreak.Resize(fyne.NewSize(40, 40))
			IMGFireStreak.Hide()

			//Огонь (Потух)
			IMGFireLose := canvas.NewImageFromResource(icons.FireLose)
			IMGFireLose.Move(fyne.NewPos(460, 0))
			IMGFireLose.Resize(fyne.NewSize(40, 40))

			return container.NewWithoutLayout(LabelHabit, DifficultiLVLIMG, ButtonInfo, ButtonDelete, ButtonReadyCheck, CompletedTasksText, IMGFireStreak, IMGFireLose, ButtonReadyCheck)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			cont := co.(*fyne.Container)
			labelWithHabit := cont.Objects[0].(*canvas.Text)
			DifficultiLVLIMG := cont.Objects[1].(*canvas.Image)
			CompletedTasksText := cont.Objects[5].(*canvas.Text)
			buttonInfo := cont.Objects[2].(*widget.Button)
			ButtonDelete := cont.Objects[3].(*widget.Button)
			ButtonReadyCheck := cont.Objects[4].(*widget.Button)
			IMGFireStreak := cont.Objects[6].(*canvas.Image)
			IMGFireLose := cont.Objects[7].(*canvas.Image)

			ButtonReadyCheck.OnTapped = func() {
				dialog.ShowConfirm("Подтверждение", "Вы выполнили задачку на сегодня?", func(confirmed bool) {
					if confirmed {
						// Увеличиваем счетчик в модели и обновляем текст
						habits[lii].CompletedTasks++
						habits[lii].LastCompletedDate = time.Now() // Сохраняем текущую дату
						habits[lii].IsFireStreakActive = true      // Устанавливаем огонь как активный
						DB.Save(&habits[lii])                      // Сохраняем изменения в базе данных

						CompletedTasksText.Text = strconv.Itoa(habits[lii].CompletedTasks)
						CompletedTasksText.Refresh() // Обновляем отображение

						IMGFireStreak.Show()
						IMGFireLose.Hide()
					}
				}, MainWindow)
			}

			// Восстановление состояния огня при старте приложения
			if habits[lii].IsFireStreakActive {
				// Проверяем, не прошел ли уже день
				if habits[lii].LastCompletedDate.Day() != time.Now().Day() {
					// Если день не совпадает, значит, нужно сбросить огонь
					habits[lii].IsFireStreakActive = false
					DB.Save(&habits[lii]) // Сохраняем изменения в базе данных

					// Обновляем иконку огня
					IMGFireStreak.Hide()
					IMGFireLose.Show()
				} else {
					// Если день совпадает, оставляем огонь активным
					IMGFireStreak.Show()
					IMGFireLose.Hide()
				}
			} else {
				IMGFireStreak.Hide()
				IMGFireLose.Show()
			}

			// Название самой привычки
			labelWithHabit.Text = habits[lii].HabitName
			labelWithHabit.Refresh()

			// Условие по отбору фото для уровня сложности
			lvlIcons := []fyne.Resource{
				icons.FirstLVLIcon, icons.SecondLVLIcon, icons.ThirdLVLICon,
				icons.FourthLVLICon, icons.FithLVLICon, icons.SixthLVLICon,
				icons.SeventhLVLICon, icons.EighthLVLICon, icons.NinthLVLICon, icons.TenthLVLICon,
			}

			if habits[lii].DifficultiLevel >= 1 && habits[lii].DifficultiLevel <= 10 {
				DifficultiLVLIMG.Resource = lvlIcons[habits[lii].DifficultiLevel-1]
				DifficultiLVLIMG.Refresh()
			}

			// Отображаем текущий счетчик выполненных задач
			CompletedTasksText.Text = strconv.Itoa(habits[lii].CompletedTasks)
			CompletedTasksText.Refresh()
			CompletedTasksText.Move(fyne.NewPos(500, 2.5))

			// Нажатие на кнопку информации
			moreInfo := habits[lii].MoreInfo
			buttonInfo.OnTapped = func() {
				dialog.ShowInformation("Более подробная информация о привычке", moreInfo, MainWindow)
			}

			// Нажатие на кнопку удаления
			ButtonDelete.OnTapped = func() {
				dialog.ShowConfirm("Удалить привычку", "Вы уверены, что хотите удалить эту привычку?", func(confirmed bool) {
					if confirmed {
						DB.Delete(&habits[lii])
					}

					var updatedHabits []models.Habits
					DB.Find(&updatedHabits)
					MainWindow.SetContent(NewMainWindow(MainWindow, updatedHabits, DB, icons))

					if len(updatedHabits) == 0 {
						NoHabitText.Show()
					} else {
						NoHabitText.Hide()
					}
				}, MainWindow)
			}
		},
	)

	// Визуальные настройки панели с привычками
	TaskScroll := container.NewScroll(taskshabits)
	TaskScroll.Resize(fyne.NewSize(640, 300))
	TaskScroll.Move(fyne.NewPos(0, 70))

	MainCont := container.NewWithoutLayout(
		TextHabits,
		ButtonHabitAdd,
		Line,
		NoHabitText,
		TaskScroll,
	)

	return MainCont
}
