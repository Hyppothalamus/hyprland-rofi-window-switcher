package main

import (
	"Hyppothalamus/wayland-rofi-windows/commands"
	"Hyppothalamus/wayland-rofi-windows/icons"
	"fmt"
	"strconv"
	"strings"
)

type window struct {
    class string
    title string
    workspace int
}

func main() {

    // TODO change output to spit out json
    // better interprete data
    windows := parseOutput(commands.Command("hyprctl clients"))

    window := commands.Command(fmt.Sprintf("echo -en '%s' | rofi -dmenu -p windows", genTitles(&windows)))

    class := getClassFromTitle(strings.TrimSpace(strings.TrimSuffix(window, "\n")), &windows)

    commands.Command(fmt.Sprintf("hyprctl dispatch focuswindow %s", class))

    return
}

// TODO change to parse JSON 
// this will be faster and better optimized
func parseOutput(out string) []window {
    strings_value := strings.Split(out, "Window")[1:]

    var windows []window

    for _, v := range(strings_value) {
        window := window{}
        for i, item := range(strings.Split(v, "\n")) {
            if item == "" || i == 0 {continue}
            key := strings.TrimSpace(strings.Split(item, ":")[0])
            value := strings.TrimSpace(strings.Split(item, ":")[1])
            switch key {
            case "class":
                window.class = value
                break
            case "title":
                window.title = value
                break
            case  "workspace":
                // TODO check lengt of workspace
                workspace, _ := strconv.Atoi(value[0:1])
                window.workspace = workspace
                break
            }
        }
        windows = append(windows, window)
    }
    return windows
}

func genTitles(windows *[]window) string {
    var result string

    for _, v := range(*windows) {
        result += " " + fmt.Sprint(v.workspace) + "\t" + v.class + " - " + v.title + "\\0icon\\x1f" + icons.GetIconName(v.class) + "\n"
    }

    return result
}

func getClassFromTitle(title string, windows *[]window) string {

    for _, v := range(*windows) {
        if v.title == strings.Split(title, " - ")[1] {
            return v.class
        }
    }
    return ""
}
