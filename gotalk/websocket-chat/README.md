This is a multi-room chat app implemented in gotalk.

    go build && ./websocket-chat

Then open http://localhost:1235/ in a web browser.
Use several browser tabs/windows to simulate multiple people chatting.

Features:

- Uses Gotalk for both Go and JavaScript
- Implements a "broadcast" function for notifying all connected sockets
- Maintains state in-memory (restarting the server loses any rooms or messages)
