package icons

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetIconName(class string) string {

    res := look_file("/usr/share/pixmaps", class)

    if res == "" {
        res = look_file("/usr/share/icons/hicolor/48x48/apps/", class)
    }

    return res
} 

func look_file(location string, class string) string {
    files, err := ioutil.ReadDir(location)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        file_name := file.Name()
        if strings.Contains(file_name, class) {
            var extension = filepath.Ext(file_name)
            var name = file_name[0:len(file_name)-len(extension)]
            return name
        }
    }
    for _, file := range files {
        file_name := file.Name()
        if strings.Contains(file_name, strings.ToLower(class)) {
            var extension = filepath.Ext(file_name)
            var name = file_name[0:len(file_name)-len(extension)]
            return name
        }
    }
    return ""
}
