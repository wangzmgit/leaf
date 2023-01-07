package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	messageClient  = make(map[interface{}]*websocket.Conn)  // 消息通道
	messageChannel = make(map[interface{}]chan interface{}) // websocket客户端链接池
	messageMux     sync.Mutex                               // 互斥锁
)

func MsgWsHandler(w http.ResponseWriter, r *http.Request, id uint) {
	conn, err := CreateWsConn(w, r)
	if err != nil {
		zap.L().Error("升级websocket失败，原因 " + err.Error())
		return
	}

	// 把与客户端的链接添加到客户端链接池中
	addMsgClient(id, conn)

	// 获取该客户端的消息通道
	m, exist := getMsgMessageChannel(id)
	if !exist {
		m = make(chan interface{})
		addMsgMessageChannel(id, m)
	}

	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteMsgClient(id, nil)
		return nil
	})

	WsHandler(conn, id, nil, m, deleteMsgClient)
}

func addMsgClient(id interface{}, conn *websocket.Conn) {
	messageMux.Lock()
	messageClient[id] = conn
	messageMux.Unlock()
}

// 获取消息管道
func getMsgMessageChannel(id interface{}) (m chan interface{}, exist bool) {
	messageMux.Lock()
	m, exist = messageChannel[id]
	messageMux.Unlock()
	return
}

// 添加消息管道
func addMsgMessageChannel(id interface{}, m chan interface{}) {
	messageMux.Lock()
	messageChannel[id] = m
	messageMux.Unlock()
}

// 移除客户端和管道
func deleteMsgClient(id, groupId interface{}) {
	messageMux.Lock()
	delete(messageClient, id)
	delete(messageChannel, id)
	messageMux.Unlock()
}

// 设置消息
func setMsgMessage(id, content interface{}) {
	messageMux.Lock()
	if m, exist := messageChannel[id]; exist {
		go func() {
			m <- content
		}()
	}
	messageMux.Unlock()
}

// 向用户发送消息
func SendMsg(id, content interface{}) {
	if id != 0 {
		_, exist := getMsgMessageChannel(id)
		if !exist {
			return
		}
	}
	setMsgMessage(id, content)
}
