package main

import (
	"fmt"

	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	buffer := femto.NewBufferFromString("", "")
	buffer.Settings["filetype"] = "markdown"
	buffer.Settings["keepautoindent"] = true
	buffer.Settings["statusline"] = false
	buffer.Settings["softwrap"] = true
	buffer.Settings["scrollbar"] = true

	// var colorscheme femto.Colorscheme
	// if monokai := runtime.Files.FindFile(femto.RTColorscheme, "monokai"); monokai != nil {
	// 	if data, err := monokai.Data(); err == nil {
	// 		colorscheme = femto.ParseColorscheme(string(data))
	// 	}
	// }
	editor := femto.NewView(buffer)
	// editor.SetColorscheme(colorscheme)
	editor.SetRuntimeFiles(runtime.Files)
	editor.Readonly = false

	list := tview.NewList()
	// AddItem("List item 1", "Some explanatory text", 'a', nil).
	// AddItem("List item 2", "Some explanatory text", 'b', nil).
	// AddItem("List item 3", "Some explanatory text", 'c', nil).
	// AddItem("List item 4", "Some explanatory text", 'd', nil)

	runTimeFileList := runtime.Files.ListRuntimeFiles("monokai")
	for _, file := range runTimeFileList {
		list.AddItem(file.Name(), "", rune(file.Name()[0]), nil)
	}

	flex := tview.NewFlex().AddItem(list, 0, 1, false).AddItem(editor, 0, 2, true)
	app.SetFocus(editor)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", runTimeFileList)
}
