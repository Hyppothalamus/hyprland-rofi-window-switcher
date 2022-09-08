package main

import (
	"Hyppothalamus/wayland-rofi-windows/commands"
	"Hyppothalamus/wayland-rofi-windows/icons"
	"encoding/json"
	"fmt"
	"strings"
)

type workspace struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type window struct {
    Class string `json:"class"`
    Title string `json:"title"`
    Workspace workspace `json:"workspace"`
}

func main() {

    // TODO change output to spit out json
    // better interprete data
    windows := parseOutput(commands.Command("hyprctl clients -j"))

    window := commands.Command(fmt.Sprintf("echo -en '%s' | rofi -dmenu -p windows", genTitles(&windows)))

    class := getClassFromTitle(strings.TrimSpace(strings.TrimSuffix(window, "\n")), &windows)

    commands.Command(fmt.Sprintf("hyprctl dispatch focuswindow %s", class))

    return
}

// TODO change to parse JSON 
// this will be faster and better optimized
func parseOutput(out string) []window {
   var windows []window
   json.Unmarshal([]byte(out), &windows)
   return windows
}

func genTitles(windows *[]window) string {
    var result string

    for _, v := range(*windows) {
        result += " " + fmt.Sprint(v.Workspace.Id) + "\t" + strings.ToLower(v.Class) + " - " + v.Title + "\\0icon\\x1f" + icons.GetIconName(v.Class) + "\n"
    }

    return result
}

func getClassFromTitle(title string, windows *[]window) string {

    for _, v := range(*windows) {
        if v.Title == strings.Split(title, " - ")[1] {
            return v.Class
        }
    }
    return ""
}
