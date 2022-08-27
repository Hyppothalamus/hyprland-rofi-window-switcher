package icons

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetIconName(class string) string {

    // first look /usr/share/pixmaps
    res := look_file("/usr/share/pixmaps", class)

    // if we get here icon not found
    // look in /usr/share/icons/hicolor/48x48/apps/
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
    return ""
}
