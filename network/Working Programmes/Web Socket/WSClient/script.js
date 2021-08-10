var socket = new WebSocket("ws://localhost:8080/chat");

socket.onopen = function(e) {
  socket.send("I've connected!");
};
socket.onmessage = function(event) {
  let messBox = document.getElementById("messages");
  messBox.innerHTML += event.data;
};

socket.onclose = function(event) {
  if (event.wasClean) {
    alert(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
  } else {
    // например, сервер убил процесс или сеть недоступна
    // обычно в этом случае event.code 1006
    alert('[close] Соединение прервано');
  }
};

function send() {
	let mes = document.getElementById("messInp").value;
	socket.send(mes);
}