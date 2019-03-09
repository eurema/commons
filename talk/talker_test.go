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

	myTalker := NewTalker(logger.WithField("test", "a"))

	_, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	myTalker.Send([]byte(msgTeste))

	ctxWaitALittle, ack := context.WithTimeout(context.Background(), 200*time.Millisecond)
	select {
	case newMsg := <-myTalker.Listen():
		msgReceived = string(newMsg)
		ack()
	case err := <-myTalker.ListenInterruption():
		assert.Fail(t, err.Text)
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

	myTalker := NewTalker(logger.WithField("test", "a"))

	connectionCtx, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	myTalker.Close()

	select {
	case <-myTalker.Listen():
		assert.Fail(t, "should not be called when the connection is closed by the player")
	case <-myTalker.ListenInterruption():
		assert.Fail(t, "should not be called when the connection is closed by the player")
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

	myTalker := NewTalker(logger.WithField("test", "a"))

	connectionCtx, err := myTalker.Connect(*wsUrl, arena.PlayerSpecifications{})
	assert.Nil(t, err)
	go func() {
		serverTestConnections[connectionName].Close()
	}()

	onCloseWasCalled := false
	select {
	case <-myTalker.Listen():
		assert.Fail(t, "should not be called when the connection is closed by the player")
	case <-myTalker.ListenInterruption():
		onCloseWasCalled = true
	case <-connectionCtx.Done():
		assert.Equal(t, context.Canceled, connectionCtx.Err(), "should receive a closing message")
	}

	assert.True(t, onCloseWasCalled)
}
