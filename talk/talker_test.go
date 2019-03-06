package talk

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/makeitplay/arena"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

var upgrader = websocket.Upgrader{}

var serverTestConnections = map[string]*websocket.Conn{}

func echo(connectionName string) (hand http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		serverTestConnections[connectionName] = c
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			err = c.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}

func TestTalker_Connection(t *testing.T) {
	// Create test server with the echo handler.
	hanlder := echo("connctect_a")
	s := httptest.NewServer(hanlder)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	wsUrl, _ := url.Parse(u)

	logger := logrus.New()

	msgTeste := "a-nice-msg"
	msgReceived := ""
	ctxWaitALittle, ack := context.WithTimeout(context.Background(), 200*time.Millisecond)
	onMsg := func(bytes []byte) {
		msgReceived = string(bytes)
		ack()
	}

	onClose := func() {

	}

	myTalker := NewTalker(logger.WithField("test", "a"), onMsg, onClose)

	_, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	myTalker.Send([]byte(msgTeste))

	select {
	case <-ctxWaitALittle.Done():

	}
	assert.Equal(t, msgTeste, msgReceived)
}

func TestTalker_ClosingConnection(t *testing.T) {
	connectionName := "closing-test"
	// Create test server with the echo handler.
	hanlder := echo(connectionName)
	s := httptest.NewServer(hanlder)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	wsUrl, _ := url.Parse(u)

	logger := logrus.New()

	onMsg := func(bytes []byte) {}

	onClose := func() {
		assert.Fail(t, "should not be called when the connection is closed by the player")
	}

	myTalker := NewTalker(logger.WithField("test", "a"), onMsg, onClose)

	connectionCtx, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	myTalker.Close()
	select {
	case <-connectionCtx.Done():
		assert.Equal(t, context.Canceled, connectionCtx.Err(), "should receive a closing message")

	}
}
func TestTalker_UnnexpectedConnectionClosed(t *testing.T) {
	connectionName := "unnexpected-clossed-1-test"
	// Create test server with the echo handler.
	hanlder := echo(connectionName)
	s := httptest.NewServer(hanlder)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	wsUrl, _ := url.Parse(u)

	logger := logrus.New()

	onMsg := func(bytes []byte) {}

	onCloseWasCalled := false
	onClose := func() {
		onCloseWasCalled = true
	}

	myTalker := NewTalker(logger.WithField("test", "a"), onMsg, onClose)

	connectionCtx, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	go func() {
		serverTestConnections[connectionName].Close()
	}()
	select {
	case <-connectionCtx.Done():
		assert.Equal(t, context.Canceled, connectionCtx.Err(), "should receive a closing message")
	}

	assert.True(t, onCloseWasCalled)
}
