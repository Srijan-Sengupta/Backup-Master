package ui

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/Srijan-Sengupta/Backup-Master/backend"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("Backup master")
	fileOutputLocation := &widget.Entry{}
	fileInputLocation := &widget.Entry{}
	//log := &widget.Label{}
	var infinite *dialog.ProgressInfiniteDialog
	progress := &widget.ProgressBar{
		Min:   -1,
		Max:   100,
		Value: -1,
	}
	/*scroll := &widget.ScrollContainer{
		Content: log,
	}
	scroll.SetMinSize(fyne.Size{
		Height: 100,
		Width:  100,
	})*/
	w.SetContent(&widget.Box{
		Horizontal: false,
		Children: []fyne.CanvasObject{
			&widget.Box{
				Horizontal: true,
				Children: []fyne.CanvasObject{
					&widget.Label{Text: "Enter the file location : "},
					fileInputLocation,
					&widget.Button{Icon: theme.FolderIcon(), OnTapped: func() {
						browser := dialog.NewFolderOpen(func(f fyne.ListableURI, e error) {
							if e != nil {
								dialog.ShowError(e, w)
							}
							if f != nil {
								fileInputLocation.SetText(strings.Replace(f.String(), "file://", "", 1))
							}

						}, w)
						browser.Show()
					},
					},
				},
			},
			&widget.Box{
				Horizontal: true,
				Children: []fyne.CanvasObject{
					&widget.Label{Text: "Enter the output file location : "},
					fileOutputLocation,
					&widget.Button{Icon: theme.FileIcon(), OnTapped: func() {
						browser := dialog.NewFileSave(func(f fyne.URIWriteCloser, e error) {
							if e != nil {
								dialog.ShowError(e, w)
							}
							if f != nil {
								fileOutputLocation.SetText(strings.Replace(f.URI().String(), "file://", "", 1))

							}
						}, w)
						browser.Show()
					},
					},
				},
			},
			&widget.Button{Text: "Start to take the backup", OnTapped: func() {
				backend.Outputf = fileOutputLocation.Text
				backend.Inputf = fileInputLocation.Text
				backend.StartTakingBackup(func(msg string) {
					//log.SetText(log.Text + "\n" + msg)
					//scroll.ScrollToBottom()
				}, func(p float64) {

					if p == -1 {
						infinite = dialog.NewProgressInfinite("Reading from the directory", "Just a minute...", w)
						infinite.Show()
					} else if (p >= 0) && (infinite != nil) {
						//fmt.Println("Hello")
						infinite.Hide()
					}
					if p == 100 {
						finished(w)
					}
					progress.SetValue(p)
				})
			},
			},
			/*scroll,*/ progress,
		}})

	w.SetIcon(resourceIconPng)
	w.Resize(fyne.Size{
		Width:  250,
		Height: 250,
	})
	w.ShowAndRun()
}
func finished(w fyne.Window) {
	finished := dialog.NewInformation("Completed !!!", "I have compressed the given file....", w)
	finished.Show()
	finished.SetDismissText("OK")
}
