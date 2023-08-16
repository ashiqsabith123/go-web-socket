const socket = new WebSocket("ws://localhost:3000/message");


    socket.addEventListener("open", (event) => {
        console.log("WebSocket connection opened:", event);
    
        // Send data to the server
        socket.send("Hello, server!");
    });


    socket.onmessage = function(event) {
        var messageDiv = document.getElementById("message");
        messageDiv.innerHTML = "Message from server: " + event.data;
    };
