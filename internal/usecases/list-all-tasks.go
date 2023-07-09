package usecases

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/mkafonso/time-slice/internal/helpers"
	"github.com/mkafonso/time-slice/internal/models"
)

type ListAllTasksManager struct {
	Tasks *models.Tasks
}

func (tm *ListAllTasksManager) ListAllTasks(tasks *models.Tasks) {
	clearScreen()
	fmt.Print("\n")

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created at"},
			{Align: simpletable.AlignCenter, Text: "Completed  at"},
		},
	}

	var cells [][]*simpletable.Cell

	for index, task := range *tasks {
		index++

		item := helpers.PendingColor(task.Name)
		done := helpers.PendingColor("No")
		createdAd := helpers.PendingColor(task.CreatedAt.Format(time.RFC822))
		completedAt := helpers.PendingColor(task.CreatedAt.Format(time.RFC822))
		if task.Done {
			item = helpers.DoneColor(fmt.Sprintf("🎉 %s", task.Name))
			done = helpers.DoneColor("Yes")
			createdAd = helpers.DoneColor(createdAd)
			completedAt = helpers.DoneColor(completedAt)
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: item},
			{Text: done},
			{Text: createdAd},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf(" 🚀 You have %d pending tasks", countPendingTasks(tasks))},
	}}

	table.Println()
	fmt.Print("\n\n\n")
}

func countPendingTasks(tasks interface{}) int {
	tasksSlice := tasks.(*models.Tasks)

	total := 0

	for _, task := range *tasksSlice {
		if !task.Done {
			total++
		}
	}

	return total
}

func clearScreen() {
	var clearCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
