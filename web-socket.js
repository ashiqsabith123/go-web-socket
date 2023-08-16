var socket = new WebSocket("ws://localhost:3000/message");
        
        socket.onmessage = function(event) {
            var messageDiv = document.getElementById("message");
            messageDiv.innerHTML += "<p>Message from server: " + event.data + "</p>";
        };
        
        document.getElementById("messageForm").addEventListener("submit", function(event) {
            event.preventDefault();
            var messageInput = document.getElementById("messageInput");
            var message = messageInput.value;
            console.log(message);
            socket.send(message);
            messageInput.value = "";
        });
