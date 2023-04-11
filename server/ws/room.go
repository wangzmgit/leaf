package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/vo"
)

var (
	roomClient  = make(map[interface{}]map[interface{}]*websocket.Conn)  //房间
	roomChannel = make(map[interface{}]map[interface{}]chan interface{}) // 房间消息通道
	roomMux     sync.Mutex                                               // 互斥锁
)

// 处理ws请求
func RoomWsHandler(w http.ResponseWriter, r *http.Request, roomId uint, clientId string) {
	conn, err := CreateWsConn(w, r)
	if err != nil {
		zap.L().Error("升级websocket失败，原因 " + err.Error())
		return
	}

	// 把与客户端的链接添加到客户端链接池中
	addRoomClient(clientId, roomId, conn)

	// 获取该客户端的消息通道
	m, exist := getRoomMessageChannel(clientId, roomId)
	if !exist {
		m = make(chan interface{})
		addRoomMessageChannel(clientId, roomId, m)
	}

	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteRoomClient(clientId, roomId)
		return nil
	})

	// 广播房间人数
	BroadcastNumber(roomId)

	WsHandler(conn, clientId, roomId, m, deleteRoomClient)
}

func addRoomClient(id, groupId interface{}, conn *websocket.Conn) {
	roomMux.Lock()
	if roomClient[groupId] == nil {
		roomClient[groupId] = make(map[interface{}]*websocket.Conn)
	}
	roomClient[groupId][id] = conn
	roomMux.Unlock()
}

// 获取消息管道
func getRoomMessageChannel(id, groupId interface{}) (m chan interface{}, exist bool) {
	roomMux.Lock()
	m, exist = roomChannel[groupId][id]
	roomMux.Unlock()
	return
}

// 添加消息管道
func addRoomMessageChannel(id, groupId interface{}, m chan interface{}) {
	roomMux.Lock()
	if roomChannel[groupId] == nil {
		roomChannel[groupId] = make(map[interface{}]chan interface{})
	}
	roomChannel[groupId][id] = m
	roomMux.Unlock()
}

// 移除客户端和管道
func deleteRoomClient(id, groupId interface{}) {
	roomMux.Lock()
	delete(roomClient[groupId], id)
	delete(roomChannel[groupId], id)
	roomMux.Unlock()
	BroadcastNumber(groupId) //广播房间人数
}

// 设置消息到房间内所有客户端
func setMessageAllClient(groupId, content interface{}) {
	roomMux.Lock()
	all := roomChannel[groupId]
	roomMux.Unlock()
	go func() {
		for _, m := range all {
			m <- content
		}
	}()
}

// 广播房间人数
func BroadcastNumber(groupId interface{}) {
	setMessageAllClient(groupId, &vo.RoomVO{
		Number: len(roomClient[groupId]),
	})
}
