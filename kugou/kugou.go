package kugou

import (
    "fmt"
    "crypto/md5"
    "io/ioutil"
    "encoding/hex"
    "net/http"
    re "regexp"
    json "github.com/bitly/go-simplejson"
)

type download interface {
    getName()
    getDownloadUrl()
    download()
}

type kugou struct {
    hash string
    songUrl string
    songName string
    downloadUrl string
}

func (kg *kugou)getName() {
    var choice rune
    pat := re.MustCompile("hash=([a-zA-Z0-9]+)")
    (*kg).hash = pat.FindStringSubmatch((*kg).songUrl)[1]
    fmt.Printf("要手动输入歌曲名称吗(建议选是)？(y/n)")
    fmt.Scanf("%v", &choice)
    if choice != 'y'{
        pat = re.MustCompile("filename=(.*?)&")
        (*kg).songName = pat.FindStringSubmatch((*kg).songUrl)[1]
    } else if choice == 'y' {
        fmt.Scanf("%v", &choice)
        fmt.Printf("请手动输入音乐的名字:")
        fmt.Scanf("%s", &(*kg).songName)
    }
}

func (kg *kugou)getDownloadUrl() {
    var cookie string
    
    key := (*kg).hash + "kgcloud"
    m := md5.New()
    m.Write([]byte(key))
    url := "http://trackercdn.kugou.com/i/?key=" + hex.EncodeToString(m.Sum(nil)) + "&cmd=4&acceptMp3=1&pid=1&hash=" + (*kg).hash
    client := &http.Client{}
    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("Cookie", cookie)
    request.Header.Set("Referer", url)
    response, _ := client.Do(request)
    body, _ := ioutil.ReadAll(response.Body)
    js, _ := json.NewJson(body)
    (*kg).downloadUrl, _ = js.Get("url").String()
}

func (kg *kugou)download() {
    
    songStream, _ := http.Get((*kg).downloadUrl)
    defer songStream.Body.Close()
    body, _ := ioutil.ReadAll(songStream.Body)
    ioutil.WriteFile((*kg).songName + ".mp3", body, 0644)
}

func do(k download) {
    k.getName()
    k.getDownloadUrl()
    k.download()
}

func Anction(url string) {
    song := &kugou{songUrl:url}
    do(song)
}
