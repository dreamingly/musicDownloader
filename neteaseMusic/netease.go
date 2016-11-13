package neteaseMusic

import (
    "io/ioutil"
    //"log"
    json "github.com/bitly/go-simplejson"
    re "regexp"
    "net/http"
)

type Download interface
{
    getSongId()
    getSongUrl()
    getArtistsDetail()
    download()
}

type Netease struct {
    songId string
    songUrl string
    DownloadUrl string
    songName string
    artistsName string
}

func (net *Netease) getSongId() {
    songRe := re.MustCompile(`[\d]+`)
    net.songId = songRe.FindAllString((*net).songUrl, -1)[1]
}

func (net *Netease) getSongUrl() {
    net.songUrl = "http://music.163.com/api/song/detail/?id=" + net.songId + "&ids=%5B" + net.songId +"%5D"
    page, _ := http.Get((*net).songUrl)
    page1, _ := ioutil.ReadAll(page.Body)
    page.Body.Close()
    mjson, _ := json.NewJson(page1)
    net.DownloadUrl = mjson.Get("songs").GetIndex(0).Get("mp3Url").MustString()
}

func (net *Netease) getArtistsDetail() {
    page, _ := http.Get(net.songUrl)
    page1, _ := ioutil.ReadAll(page.Body)
    page.Body.Close()
    mjson, _ := json.NewJson(page1)
    net.songName = mjson.Get("songs").GetIndex(0).Get("name").MustString()
    net.artistsName = mjson.Get("songs").GetIndex(0).Get("artists").GetIndex(0).Get("name").MustString()
}

func (net *Netease) download() {
    songStream, _ := http.Get(net.DownloadUrl)
    defer songStream.Body.Close()
    body, _ := ioutil.ReadAll(songStream.Body)
    _ = ioutil.WriteFile((*net).artistsName + "-" + net.songName + ".mp3", body, 0644)
}

/*
func err(errCode error) {
    if errCode != nil {
        log.Fatal(err)
    }
}
*/

func do(nts Download) {
    nts.getSongId()
    nts.getSongUrl()
    nts.getArtistsDetail()
    nts.download()
}

func Action(url string) {
    song := &Netease{songUrl: url}
    do(song)
}
