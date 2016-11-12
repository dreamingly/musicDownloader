package main

import (
    "fmt"
    re "regexp"
    "musicDownloader/neteaseMusic"
    "musicDownloader/kugou"
)

func main() {
    var url string
    print("输入url:")
    fmt.Scanf("%s", &url)
    isNetease, _ := re.MatchString("163", url)
    if isNetease {
        neteaseMusic.Action(url)
    } else {
        kugou.Anction(url)
    }
}
