package main

import (
	"fmt"
	"os"
)

// _____________________________________________________________
//
//  Variables.
// _____________________________________________________________
var lessCompiler bool
var answer string
var wait string
var easyLessConfig string
var lesspath string
var projectname string

// _____________________________________________________________
//
//  Main Function.
//  This function is auto-executed by Go.
// _____________________________________________________________
func main() {

	// _____________________________________________________________
	//
	//  Project Name
	//  The software ask you for project name
	// _____________________________________________________________
	fmt.Println("What is the project name?")
	fmt.Scanln(&projectname)
	if projectname == "" {
		projectname = "new-project"
	}
	fmt.Println("")

	// _____________________________________________________________
	//
	//  Eassy Less settings to Visual Studio Code
	//  Select "Y" if you want to use it, or select "N" if you don't.
	// _____________________________________________________________
	easyLessConfig = "// main:styles.less, out: ../css/styles.min.css, compress: true, sourceMap: false"

	fmt.Println("Would you like to use the Eassy LESS plugin to compile less on Visual Studio Code? (Y/N)")
	fmt.Scanln(&answer)
	if answer == "Y" || answer == "y" {
		lessCompiler = true
	} else {
		lessCompiler = false
	}

	// _____________________________________________________________
	//
	//  Create folders.
	//  Create folders as public/css, public/images, etc..
	// _____________________________________________________________
	createFolder(projectname + "/public/less")
	createFolder(projectname + "/public/css")
	createFolder(projectname + "/public/images")
	createFolder(projectname + "/public/js")

	// _____________________________________________________________
	//
	//  Create Less Files
	//  If you don't want less files, comment the next 5 lines.
	// _____________________________________________________________
	lesspath = projectname + "/public/less/"
	createVariablesDocument()
	createResetDocument()
	createLessDocuments()
	createStylesDocument()

	// _____________________________________________________________
	//
	//  Successfuly Message.
	// _____________________________________________________________
	fmt.Printf("\nProject created successfuly. Press ENTER to exit")

	// _____________________________________________________________
	//
	//  Press ENTER to exit
	// _____________________________________________________________
	fmt.Scanln(&wait)
}

// _____________________________________________________________
//
//  createFolder() function
//  It's a generic function to create sub-folders
// _____________________________________________________________
func createFolder(pathname string) {
	os.MkdirAll(pathname, os.ModePerm)
}

// _____________________________________________________________
//
//  createStylesDocument() function
//  This fuction create the styles.less with @import less files.
// _____________________________________________________________
func createStylesDocument() {
	file, err := os.Create(lesspath + "styles.less")
	if err != nil {
		fmt.Println("Hubo un problema con la creación del archivo styles.Less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	fmt.Fprintln(file, "@import url('reset.less');")
	fmt.Fprintln(file, "@import url('variables.less');")
	fmt.Fprintln(file, "@import url('functions.less');")
	fmt.Fprintln(file, "@import url('main.less');")
	fmt.Fprintln(file, "@import url('header.less');")
	fmt.Fprintln(file, "@import url('aside.less');")
	fmt.Fprintln(file, "@import url('navbar.less');")
	fmt.Fprintln(file, "@import url('articles.less');")
	fmt.Fprintln(file, "@import url('footer.less');")
	fmt.Fprintln(file, "@import url('forms.less');")
	fmt.Fprintln(file, "@import url('buttons.less');")
	fmt.Fprintln(file, "@import url('tables.less');")
	fmt.Fprintln(file, "@import url('helpers.less');")
	fmt.Fprintln(file, "@import url('responsive.less');")
	fmt.Fprintln(file, "@import url('animations.less');")
	defer file.Close()
}

// _____________________________________________________________
//
//  createVariablesDocument() function
//  This function create the variables.less sheet
// _____________________________________________________________
func createVariablesDocument() {

	selectors := [9]string{"@imports", "background", "border", "colors", "font", "height", "margin", "padding", "width"}

	file, err := os.Create(lesspath + "variables.less")
	if err != nil {
		fmt.Println("Hubo un problema con la creación del archivo variables.less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	for i := 0; i < 9; i++ {
		fmt.Fprintln(file, "/* =============================================================================")
		fmt.Fprintln(file, "* \t", selectors[i])
		fmt.Fprintln(file, "*  ========================================================================== */")
		fmt.Fprintln(file, "")
	}
	defer file.Close()
}

// _____________________________________________________________
//
//  createResetDocument() function
//  This function create the reset.less sheet
// _____________________________________________________________
func createResetDocument() {
	file, err := os.Create(lesspath + "reset.less")
	if err != nil {
		fmt.Println("Hubo un problema con la creación del archivo reset.Less")
		return
	}
	if lessCompiler == true {
		fmt.Fprintln(file, easyLessConfig)
	}
	fmt.Fprintln(file, "* {")
	fmt.Fprintln(file, "\t margin: 0;")
	fmt.Fprintln(file, "\t border: 0;")
	fmt.Fprintln(file, "\t padding: 0;")
	fmt.Fprintln(file, "\t outline: none;")
	fmt.Fprintln(file, "}")
	defer file.Close()
}

// _____________________________________________________________
//
//  createLessDocuments() function
//  This function create each less file used in this project
// _____________________________________________________________
func createLessDocuments() {

	files := [13]string{"functions.less", "header.less", "navbar.less", "main.less", "aside.less", "articles.less", "footer.less", "forms.less", "buttons.less", "tables.less", "helpers.less", "responsive.less", "animations.less"}

	for i := 0; i < len(files); i++ {
		file, err := os.Create(lesspath + files[i])
		if err != nil {
			fmt.Printf("Hubo un problema con la creación del archivo %v", files[i])
			panic(err)
		}
		if lessCompiler == true {
			fmt.Fprintln(file, easyLessConfig)
		}
		defer file.Close()
	}

}
