# Go Slack Bot

The go-slack-bot application will be sending requests to a Go backend via Websocket, something called Socket-Mode in the slack world.

We will be using go-slack which is a library that supports the regular REST API, WebSockets, RTM, and Events.

The slack package will need to create a new Client for the socket mode, so we will have two clients. One that uses the regular API and one for the websocket events.

## Pre-requisites

To create a bot for your slack workspace or channel, you need to first create a slack app from https://api.slack.com/apps as mentioned [here](https://slack.com/intl/en-in/help/articles/115005265703-Create-a-bot-for-your-workspace)

- setup OAuth & permissions 
- select appropriate scopes.
- enable events from event subscriptions and subscribe to appropriate bot events

## Connection process

Websocket client is created by calling socketmode.New and given the regular client as input. I’ve also added a OptionAppLevelToken to the creation of the regular client since that is now needed to connect to the Socket.

The output should be connected and there will be a ping hello sent as shown below:-

```
➜  go-slack-bot git:(main) ✗ go run main.go 
socketmode: 2022/10/16 21:46:51 socket_mode_managed_conn.go:258: Starting SocketMode
slack-go/slack2022/10/16 21:46:57 socket_mode.go:30: Using URL: wss://wss-primary.slack.com/link/?ticket=<redacted>&app_id=<redacted>
socketmode: 2022/10/16 21:46:57 socket_mode_managed_conn.go:266: Dialing to websocket on url wss://wss-primary.slack.com/link/?ticket=<redacted>&app_id=<redacted>
socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:90: WebSocket connection succeeded on try 0
socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:439: Starting to receive message
socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:481: Incoming WebSocket message: {
  "type": "hello",
  "num_connections": 1,
  "debug_info": {
    "host": "applink-1",
    "build_number": 14,
    "approximate_connection_time": 18060
  },
  "connection_info": {
    "app_id": "A0469A3MM9D"
  }
}

socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:493: Finished to receive message
socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:439: Starting to receive message
socketmode: 2022/10/16 21:46:59 socket_mode_managed_conn.go:336: Received WebSocket message: {"type":"hello","num_connections":1,"debug_info":{"host":"applink-1","build_number":14,"approximate_connection_time":18060},"connection_info":{"app_id":"A0469A3MM9D"}}
socketmode: 2022/10/16 21:47:02 socket_mode_managed_conn.go:561: WebSocket ping message received: Ping from applink-1
socketmode: 2022/10/16 21:47:12 socket_mode_managed_conn.go:561: WebSocket ping message received: Ping from applink-1
socketmode: 2022/10/16 21:47:22 socket_mode_managed_conn.go:561: WebSocket ping message received: Ping from applink-1
socketmode: 2022/10/16 21:47:32 socket_mode_managed_conn.go:561: WebSocket ping message receive
```
## Listen for events

It’s time to start selecting all events to listen for. At the end of the program, we call socketClient.Run() which will be blocking and ingesting new Websocket messages on a channel at socketClient.Events. So we can use a for loop to continuously wait for new events, also the slack-go library comes with predefined Event types, so we can use a type switch to handle different types of Events easily. All events can be found [here](https://api.slack.com/events?filter=Events)

Since socketClient.Run() is blocking, we will spawn a goroutine that handles incoming messages in the background.

We will begin by simply logging in whenever an Event on the EventAPI is triggered in Slack. Since we first need to type switch the message on the websocket if it’s an EventsAPI type, then switch again based on the actual Event that occurred we will break the Event handling out into a separate function to avoid deeply nested switches.

[For reference](https://towardsdatascience.com/develop-a-slack-bot-using-golang-1025b3e606bc)

## How to test?

Make sure you have the environment in .env

```
go run main.go
```

Enter slack and mention bot by using @your_bot_name