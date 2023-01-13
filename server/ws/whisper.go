package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	whisperClient  = make(map[interface{}]*websocket.Conn)  // 消息通道
	whisperChannel = make(map[interface{}]chan interface{}) // websocket客户端链接池
	whisperMux     sync.Mutex                               // 互斥锁
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
	whisperMux.Lock()
	whisperClient[id] = conn
	whisperMux.Unlock()
}

// 获取消息管道
func getMsgMessageChannel(id interface{}) (m chan interface{}, exist bool) {
	whisperMux.Lock()
	m, exist = whisperChannel[id]
	whisperMux.Unlock()
	return
}

// 添加消息管道
func addMsgMessageChannel(id interface{}, m chan interface{}) {
	whisperMux.Lock()
	whisperChannel[id] = m
	whisperMux.Unlock()
}

// 移除客户端和管道
func deleteMsgClient(id, groupId interface{}) {
	whisperMux.Lock()
	delete(whisperClient, id)
	delete(whisperChannel, id)
	whisperMux.Unlock()
}

// 设置消息
func setMsgMessage(id, content interface{}) {
	whisperMux.Lock()
	if m, exist := whisperChannel[id]; exist {
		go func() {
			m <- content
		}()
	}
	whisperMux.Unlock()
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
