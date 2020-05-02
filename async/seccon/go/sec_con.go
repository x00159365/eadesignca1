package seccon

import (
	"html/template"
	"net/http"
	"time"
)

var counts map[string]string = make(map[string]string)
var channelNames []string

func ProcessMessage(channel_name string, msg_content string) {
	// the function just counts the messages so message_content goes unused
	//counts[channel_name]++
	// MG changing to use the content
	counts[channel_name] = msg_content
}

func SetUpSecConPubSub(redis_url string, interval time.Duration, wait time.Duration, channel_names []string) {

	channelNames = channel_names
	SetUpPubSub("security-console", redis_url, interval, wait, channel_names, ProcessMessage)
}

func GetTableInnerHTML() string {
	var innerHTML string = "<tr>"
	for _, name := range channelNames {
		innerHTML += "<th>" + name + "</th>"
	}
	innerHTML += "</tr><tr>"
	for _, name := range channelNames {
		// MG
		//innerHTML += "<th>" + strconv.Itoa(counts[name]) + "</th>"
		innerHTML += "<th>" + counts[name] + "</th>"
	}
	innerHTML += "</tr>"
	return innerHTML
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	inserts := struct {
		TableInnerHTML template.HTML
	}{
		template.HTML(GetTableInnerHTML()),
	}
	t, _ := template.ParseFiles("console.html")

	t.Execute(w, inserts)
}
