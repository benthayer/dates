package main


import (
    "math/rand"
    "time"
    "net/http"
    "html/template"
)

func getRandomDate() time.Time {
    min := time.Date(1200, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    max := time.Date(2300, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
    delta := max - min

    sec := rand.Int63n(delta) + min

    return time.Unix(sec, 0)
}

func serveRandomDate(w http.ResponseWriter, r *http.Request) {
    const layoutUS = "January 2, 2006"
    date := getRandomDate()
    weekday := date.Weekday()

    type DayInfo struct {
        Date string
        Weekday string
    }

    t, _ := template.ParseFiles("index.html")

    t.Execute(w, &DayInfo{Date: date.Format(layoutUS), Weekday: weekday.String()})
}

func main() {
    http.HandleFunc("/", serveRandomDate)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

