if (typeof console === 'undefined') {
  console = {error:function(){},log:function(){}};
}

var s; // our socket
var currentRoom; // name of the room we are currently viewing
var roomsUL = document.querySelector('#rooms ul.rooms');
var newRoom = document.querySelector('#rooms form');
var newRoomInputField = newRoom.querySelector('input');
var roomMessagesUL = document.querySelector('#room ul.messages');
var composer = document.querySelector('#room .composer');
var composerInputField = composer.querySelector('.body');
var composingMessages = {}; // room -> string
var accessKeys = {};
var unseenMessageCount = {}; // room -> int


function makeRoomLI(room) {
  var name = room.name;

  var accessKeyOffs = 0, accessKey;
  for (; accessKeyOffs !== name.length; ++accessKeyOffs ) {
    accessKey = name[accessKeyOffs].toUpperCase();
    if (!accessKeys[accessKey]) {
      accessKeys[accessKey] = function() { viewRoom(name); }
      break;
    }
  }

  var li = document.createElement('li');
  li.id = 'roomli-' + name;

  var s, a = document.createElement('a');
  li.onclick = a.onclick = function(ev) { ev.stopPropagation(); viewRoom(name); }
  if (accessKey) {
    li.title = 'Ctrl-' + accessKey;
    if (accessKeyOffs > 0) {
      a.appendChild(document.createTextNode(name.substr(0,accessKeyOffs)));
    }
    s = document.createElement('span');
    s.className = 'accessKey';
    s.innerText = name.substr(accessKeyOffs,1);
    a.appendChild(s);
    a.appendChild(document.createTextNode(name.substr(accessKeyOffs+1)));
  } else {
    a.innerText = name;
  }
  a.href = '#'+name;
  li.appendChild(a);

  s = document.createElement('span');
  updateUnseenMessageCount(unseenMessageCount[name], s);
  li.appendChild(s);

  return li;
}


function updateUnseenMessageCount(room, unseenSpan) {
  var count = unseenMessageCount[room];
  if (!unseenSpan) {
    unseenSpan = roomsUL.querySelector('#roomli-'+room+' span.unseenCount');
  }
  if (!count) {
    if (unseenSpan.innerText === '' || !unseenSpan.innerText) {
      unseenSpan.innerText = '0';
    }
    unseenSpan.className = 'unseenCount';
    unseenSpan.title = '';
  } else {
    unseenSpan.innerText = String(count);
    unseenSpan.className = 'unseenCount show';
    if (count === 1) {
      unseenSpan.title = 'Unseen message';
    } else {
      unseenSpan.title = String(count) + ' unseen messages';
    }
  }
}


function makeMessageView(message) {
  var li = document.createElement('li');
  var author = document.createElement('span');
  author.innerText = message.author;
  author.className = 'author';
  li.appendChild(author);
  var body = document.createElement('span');
  body.innerText = message.body;
  body.className = 'body';
  li.appendChild(body);
  return li;
}


function showListItems(list, makefun, replaceExisting, items) {
  list.style.display = 'none';
  if (replaceExisting) { list.innerText = ''; }
  if (Array.isArray(items)) {
    items.forEach(function (item) {
      list.appendChild(makefun(item));
    });
  } else if (items) {
    Object.keys(items).forEach(function (k) {
      list.appendChild(makefun(items[k]));
    });
  }
  list.style.display = null; // avoids multiple draw operations
}

function showRooms(replaceExisting, rooms) {
  if (replaceExisting) {
    accessKeys = {};
  }
  var roomsList = [];
  Object.keys(rooms).forEach(function (k) {
    roomsList.push(rooms[k]);
  });
  roomsList.sort(function (a, b) {
    return b.name < a.name;
  })
  showListItems(roomsUL, makeRoomLI, replaceExisting, roomsList);
}

function showMessages(replaceExisting, messages) {
  showListItems(roomMessagesUL, makeMessageView, replaceExisting, messages);
}


// Navigate to a room
function viewRoom(roomName) {
  if (currentRoom === roomName) return;

  // store any message that hasn't been sent in the current room:
  var message = getComposerMessage();
  composingMessages[currentRoom] = message;

  currentRoom = roomName;
  document.location.hash = '#'+roomName;
  document.querySelector('#room > h2').innerText = roomName;

  // Restore any previously "being composed" message for this room
  var message = composingMessages[roomName];
  composerInputField.value = message || '';
  composerInputField.focus();
  if (message) {
    composerInputField.setSelectionRange(message.length, message.length);
  }

  if (unseenMessageCount[roomName]) {
    unseenMessageCount[roomName] = 0;
    updateUnseenMessageCount(roomName);
  }

  roomMessagesUL.innerText = '';
  fetchMessagesInRoom(roomName);
}


function fetchMessagesInRoom(roomName) {
  s.request('list-messages', roomName, function (err, messages) {
    if (err) return console.error(err.stack);
    showMessages(/*replaceExisting=*/false, messages);
  });
}


// Handle "newmsg" notifications broadcast by the server
gotalk.handleNotification('newmsg', function (m) {
  if (currentRoom === m.room) {
    showMessages(/*replaceExisting=*/false, [m.message]);
    updateUnseenMessageCount(m.room);
  } else {
    if (unseenMessageCount[m.room]) {
      unseenMessageCount[m.room]++;
    } else {
      unseenMessageCount[m.room] = 1;
    }
    updateUnseenMessageCount(m.room);
  }
});


// Handle "rooms" notifications broadcast by the server
var onrooms;
gotalk.handleNotification('rooms', function (rooms) {
  showRooms(/*replaceExisting=*/true, rooms);
  if (onrooms) {
    onrooms(rooms);
    onrooms = null;
  }
});


// We get assigned a username
gotalk.handleNotification('username', function (username) {
  Array.prototype.forEach.call(document.querySelectorAll('.my-username'), function (e) {
    e.innerText = username;
  });
});


function getComposerMessage() {
  return composerInputField.value.replace(/^[ \s\t\r\n]+|[ \s\t\r\n]+$/g, '');
}

function getNewRoomName() {
  return newRoomInputField.value.replace(/^[ \s\t\r\n]+|[ \s\t\r\n]+$/g, '');
}


// Takes any message entered in the composer and sends it as a new message
// to the current room.
composer.onsubmit = function () {
  var body = getComposerMessage();
  composerInputField.value = body;
  if (body.length) {
    composerInputField.enabled = false;
    s.request('send-message', {room:currentRoom, message:{body:body}}, function (err, res) {
      composerInputField.enabled = true;
      if (err) return console.error(err.stack);
      composerInputField.value = '';
    });
  }
  return false;
};


// Enable creation of new rooms
newRoom.onsubmit = function () {
  var roomName = getNewRoomName();
  newRoomInputField.value = roomName;
  if (roomName.length) {
    newRoomInputField.enabled = false;
    s.request('create-room', roomName, function (err, r) {
      newRoomInputField.enabled = true;
      if (err) return console.error(err.stack);
      newRoomInputField.value = '';
      viewRoom(r.name);
    });
  }
  return false;
};


// Update list of rooms when connection opens
function onConnect(s) {
  console.log("connection opened")
  if (currentRoom) {
    fetchMessagesInRoom(currentRoom);
  }
  onrooms = function(rooms) {
    // View a room, unless we're already viewing an existing room
    if (!rooms[currentRoom]) {
      var goToRoomName, h = document.location.hash;
      if (h && h !== '#' && rooms[h.substr(1)]) {
        goToRoomName = h.substr(1);
      } else {
        for (goToRoomName in rooms) { break; }
      }
      viewRoom(goToRoomName);
    }
  };
}


// Connect to server
var s = gotalk.connection()
  .on('open', onConnect)
  .on("close", err => {
    console.log("connection closed" + (err ? " with error: " + err : ""))
  })


// Intercept ctrl+accesskey
document.addEventListener('keydown', function (ev) {
  if (ev.ctrlKey) {
    var f = accessKeys[String.fromCharCode(ev.keyCode).toUpperCase()];
    if (f) {
      ev.preventDefault();
      ev.stopPropagation();
      f(ev);
    }
  }
}, true);
