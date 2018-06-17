package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
    "math/rand"
	"time"
    "gopkg.in/yaml.v2"
)

func main() {

    start := time.Now()
    c := make(chan TestCall)
    ticker := time.NewTicker(1 * time.Second)
    testSuite := newTestSuite()
    var aliveUsers = make(map[int]bool)
    
    urlsFile, _ := ioutil.ReadFile("urls.yml")
    urls := []string{}
    yaml.Unmarshal(urlsFile, &urls)
    maxUsers := 70
    seconds := 30
    inCresciendoSeconds := int(seconds*5/7)
    inDecresciendoSeconds := int(seconds/7)

    for i:=0; i<maxUsers; i++ {
        secondsToStart := int((float64(i)/float64(maxUsers))*float64(inCresciendoSeconds))
        secondsToEnd := int((float64(i)/float64(maxUsers))*float64(inDecresciendoSeconds))
        aliveUsers[i] = true

        go doAction(c, urls, start, User{
            Id: i,
            ActiveFrom: secondsToStart,
            ActiveTo: seconds-secondsToEnd,
        }, &aliveUsers)
    }

    go func() {
        for {
            <- ticker.C
            if (len(aliveUsers) == 0) {
                close(c)
                return
            }
            lastSecond := int(time.Now().UnixNano()/int64(time.Second)-1)
            testSuite.TestCalls[lastSecond].buildReport().paint()
            testSuite.addTestCall(generateTestCallByUrlAndUser(
                "/health?token=0e4d75ba-c640-44c1-a745-06ee51db4e93",
                User{},
            ))
        }
    }()

    for tc := range c {
        if (tc.From > 0) {
            testSuite.addTestCall(tc)
        }
        go doAction(c, urls, start, tc.User, &aliveUsers)
    }
}

func doAction (c chan TestCall, urls []string, start time.Time, user User, aliveUsers *map[int]bool) {
    time.Sleep(100 * time.Millisecond)
    currentSecond := int((time.Now().UnixNano()-start.UnixNano())/int64(1000000000))
    if (currentSecond >= user.ActiveFrom) {
        if (currentSecond < user.ActiveTo) {
            go visitUrl(c, urls, user)
        } else {
            delete(*aliveUsers, user.Id)
        }
    } else {
        go wait(c, user)
    }
}

func visitUrl(c chan TestCall, urls []string, user User) {
    position := rand.Intn(len(urls))
    currentUrl := urls[position]

    c <- generateTestCallByUrlAndUser(currentUrl, user)
}

func generateTestCallByUrlAndUser(url string, user User) TestCall {
    from := time.Now()
    resp, _ := http.Get("http://0.0.0.0:8100" + url)
    body, _ := ioutil.ReadAll(resp.Body)
    to := time.Now()
    defer resp.Body.Close()
    chr := CheckHealthResponse{}
    json.Unmarshal(body, &chr)

    return TestCall{
        User: user,
        Resp: chr,
        From: from.UnixNano(),
        To: to.UnixNano(),
    }
}

func wait(c chan TestCall, user User) {
    c <- TestCall{
        User: user,
    }
}
