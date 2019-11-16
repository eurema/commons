package talk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lugobots/arena"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"sync"
)

type Talker interface {
	Connect(mainCtx context.Context, url url.URL, playerSpec arena.PlayerSpecifications) (ctx context.Context, err error)
	Send(data []byte) error
	Listen() <-chan []byte
	ListenInterruption() <-chan *websocket.CloseError
	Close()
}

// channel is meant to make the websocket connection and communication easier.
type channel struct {
	ws                *websocket.Conn
	playerSpec        arena.PlayerSpecifications
	urlConnection     url.URL
	connectionCtx     context.Context
	connectionCloser  context.CancelFunc
	ReaderChan        chan []byte
	InterruptChan     chan *websocket.CloseError
	writingMitx       sync.Mutex
	logger            *logrus.Entry
	connectionOpenned bool
}

//	NewTalker creates a new talker that knows how to talk to the game server
func NewTalker(logger *logrus.Entry) Talker {
	return &channel{
		logger:        logger,
		ReaderChan:    make(chan []byte, 1),
		InterruptChan: make(chan *websocket.CloseError, 1),
	}
}

// Connect tries to open a new web socket connection with the game server
func (c *channel) Connect(mainCtx context.Context, url url.URL, playerSpec arena.PlayerSpecifications) (ctx context.Context, err error) {
	c.playerSpec = playerSpec
	c.urlConnection = url
	if err := c.dial(); err != nil {
		return nil, err
	}
	c.connectionCtx, c.connectionCloser = context.WithCancel(mainCtx)
	c.connectionOpenned = true
	go c.keepListenning()

	go func() {
		select {
		case <-mainCtx.Done():
			c.Close()
		}
	}()
	return c.connectionCtx, nil
}

// Listen send a new message when the game server send one
func (c *channel) Listen() <-chan []byte {
	return c.ReaderChan
}

func (c *channel) ListenInterruption() <-chan *websocket.CloseError {
	return c.InterruptChan
}

// Send allow the player to send a ws message to the game server
func (c *channel) Send(data []byte) error {
	c.writingMitx.Lock()
	defer c.writingMitx.Unlock()
	return c.ws.WriteMessage(websocket.TextMessage, data)
}

func (c *channel) Close() {
	c.connectionOpenned = false
	c.ws.WriteMessage(websocket.CloseNormalClosure, []byte("bye"))
	c.ws.Close()
}

func (c *channel) dial() error {
	connectHeader := http.Header{}
	specJson, err := json.Marshal(c.playerSpec)
	if err != nil {
		return fmt.Errorf("fail on bulding the player spec header: %s", err.Error())
	}
	connectHeader.Add("X-Player-Specs", string(specJson))

	c.ws, _, err = websocket.DefaultDialer.Dial(c.urlConnection.String(), connectHeader)
	if err != nil {
		return fmt.Errorf("fail on dialing to ws server: %s", err.Error())
	}
	return nil
}

func (c *channel) keepListenning() {
	for {
		msgType, message, err := c.ws.ReadMessage()
		if e, ok := err.(*websocket.CloseError); ok {
			if e.Code == websocket.CloseGoingAway || e.Code == websocket.CloseAbnormalClosure {
				c.InterruptChan <- e
			} else if e.Code == websocket.CloseNormalClosure && c.connectionOpenned {
				c.InterruptChan <- e
			} else {
				c.logger.Infof("Connection closed by the player (%d): %s", msgType, e)
			}
			if c.connectionOpenned { //something close not asked by us
				c.connectionCloser() //unexpected
			}
			return
		} else if e, ok := err.(net.Error); ok && c.connectionOpenned {
			c.connectionCloser() //unexpected
			c.logger.Infof("unnexpected connection closed (%d): %s", msgType, e)
			return
		} else if err != nil {
			c.connectionCloser() //unexpected
			return
		} else {
			c.ReaderChan <- message
		}
	}
}
