package main

import (
	"Hyppothalamus/wayland-rofi-windows/icons"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type window struct {
    class string
    title string
    workspace int
}

func main() {

    icons_name := icons.GetIconName("discord")
    fmt.Printf("found icon: %s\n", icons_name)

    // TODO change output to spit out json
    // better interprete data
    out, err := exec.Command("hyprctl", "clients").Output()

    if err != nil {
        log.Fatal(err)
    }

    windows := parseOutput(string(out))
    
    cmd := exec.Command("rofi", "-dmenu", "-p", "windows")
    cmd.Stdin = strings.NewReader(genTitles(&windows))

    var window bytes.Buffer
    cmd.Stdout = &window

    err = cmd.Run()

    if err != nil {
        log.Fatal(err)
    }

    // TODO generate icons
    class := getClassFromTitle(strings.TrimSuffix(window.String(), "\n"), &windows)
    fmt.Printf("rofi: %q\n", string(window.Bytes()))
    fmt.Printf("selected window %s\n", class)

    _, err2 := exec.Command("hyprctl", "dispatch", "focuswindow", class).Output()

    if err2 != nil {
        log.Fatal(err)
    }
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
                window.workspace = int(value[0])
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
        result += v.title + "\n"
    }

    return result
}

func getClassFromTitle(title string, windows *[]window) string {

    for _, v := range(*windows) {
        if v.title == title {
            return v.class
        }
    }
    return ""
}
