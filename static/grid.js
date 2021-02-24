let drag = false
let isStart = false
let wall = []
let startId = 0
let endId = 0

function mouseDown() {
    drag = true
    console.log("DOWN: ", drag)
}

function mouseUp() {
    drag = false
    console.log("UP: ", drag)
}

function hover(element) {
    if (drag) {
        element.style.backgroundColor = "#332727";
        wall.push(parseInt(element.id))
    }
}

function randomizePoints() {
    startId = Math.floor(Math.random() * Math.floor(1499));
    endId = Math.floor(Math.random() * Math.floor(1499));
    clearGrid()
}

function sendGridState(algoId) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/", true);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.send(JSON.stringify({
        "length": wall.length,
        "wall": wall,
        "algo": parseInt(algoId),
        "start_id": startId,
        "end_id": endId
    }));

    xhr.onload = function() {
        console.log("Response: ")
        let resp = JSON.parse(this.responseText);
        console.log(resp);
        paintDiscovery(resp)
    }
}