package main

import (
	"bufio"
	"container/list"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

var connid int
var conns *list.List

func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()

	//產生使用者名稱
	connid++
	id := connid

	fmt.Printf("connection id: %d\n", id)

	item := conns.PushBack(ws)
	defer conns.Remove(item)

	name := fmt.Sprintf("user%d", id)

	SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))

	r := bufio.NewReader(ws)

	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Printf("disconnected id: %d\n", id)
			SendMessage(item, fmt.Sprintf("%s offline\n", name))
			break
		}

		fmt.Printf("%s: %s", name, data)

		SendMessage(item, fmt.Sprintf("%s\t> %s", name, data))
	}
}

func SendMessage(self *list.Element, data string) {
	// for _, item := range conns {
	for item := conns.Front(); item != nil; item = item.Next() {
		ws, ok := item.Value.(*websocket.Conn)
		if !ok {
			panic("item not *websocket.Conn")
		}

		if item == self {
			continue
		}

		io.WriteString(ws, data)
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>golang websocket chatroom</title>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    <script>
        var ws = new WebSocket("ws://127.0.0.1:5000/chatroom");
        ws.onopen = function(e){
            console.log("onopen");
            console.dir(e);
        };
        ws.onmessage = function(e){
            console.log("onmessage");
            console.dir(e);
            $('#log').append('<p>'+e.data+'<p>');
            $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
        };
        ws.onclose = function(e){
            console.log("onclose");
            console.dir(e);
        };
        ws.onerror = function(e){
            console.log("onerror");
            console.dir(e);
        };
        $(function(){
            $('#msgform').submit(function(){
                ws.send($('#msg').val()+"\n");
                $('#log').append('<p style="color:red;">您 > '+$('#msg').val()+'<p>');
                $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
                $('#msg').val('');
                return false;
            });
        });
    </script>
</head>
<body>
    <div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
    </div>
    <div>
        <form id="msgform">
            <input type="text" id="msg" size="60" />
        </form>
    </div>
</body>
</html>`
	io.WriteString(w, html)
}

func main() {
	fmt.Printf("Welcome chatroom server!")

	connid = 0
	conns = list.New()

	http.Handle("/chatroom", websocket.Handler(ChatroomServer))
	http.HandleFunc("/", Client)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
