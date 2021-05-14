// States
let cols = 0
let rows = 0
let isMouseDown = false
let wall = []
let grid = []
let startId = 246
let endId = 262
const wallColor = "#332727"
const cellColor = "#FFF"
const API_CALC_PATH = "https://dhnp1uigod.execute-api.us-east-1.amazonaws.com/test/calculate"
// END

function newGrid() {
    wall = []
    isMouseDown = false
    grid = []
    cols = 30
    rows = 20

    generateGrid(rows, cols, "grid")
    paintEndpoints()
}

function generateGrid(rows, cols, gridId) {
    let gridContainer = document.getElementById(gridId);

    for (let i = 0; i < rows; i++) {
        let gridRow = document.createElement("div")
        gridRow.setAttribute("class", "grid-row")

        for (let j = 0; j < cols; j++) {
            const cellId = (i*cols) + j
            const cell = newGridCell(cellId)
            gridRow.appendChild(cell)
        }
        gridContainer.appendChild(gridRow)
    }
}

function newGridCell(id) {
    let cell = document.createElement("div")
    cell.setAttribute("id", id)
    cell.setAttribute("class", "cell")
    cell.addEventListener("mousedown", mouseDownHandler, false)
    cell.addEventListener("mouseup", mouseUpHandler, false)
    cell.addEventListener("mouseover", function() {paintCellHandler(cell)}, false)
    cell.addEventListener("click", function() {paintSingleCellHandler(cell)}, false)

    return cell
}

function clearWalls() {
    wall.forEach(id => document.getElementById(id).style.backgroundColor = cellColor);
    wall = []
}

function clearGrid() {
    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            const cellId = (i*cols) + j
            document.getElementById(cellId).style.backgroundColor = cellColor
        }
    }
}

function clearAll() {
    //clearWalls()
    clearGrid()
    paintEndpoints()
}

function paintEndpoints() {
    document.getElementById(startId).style.backgroundColor = "#59d059"
    document.getElementById(endId).style.backgroundColor = "#f55e5e"
}

function randomizePoints() {
    document.getElementById(startId).style.backgroundColor = cellColor
    document.getElementById(endId).style.backgroundColor = cellColor
    startId = Math.floor(Math.random() * Math.floor(rows*cols));
    endId = Math.floor(Math.random() * Math.floor(rows*cols));
    clearAll()
}


// API
function calculatePath(algoId) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", API_CALC_PATH, true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    //xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');

    xhr.send(JSON.stringify({
        "wall": wall,
        "algo": parseInt(algoId),
        "start_id": startId,
        "end_id": endId,
        "rows": rows,
        "cols": cols,
    }));

    xhr.onload = function() {
        console.log("Response: ")
        let resp = JSON.parse(this.responseText);
        console.log(resp);
        paintDiscovery(resp)
    }
}
// END

// Path Paint

function start() {
    clearAll()
    let algoId = document.getElementById("algo").value
    calculatePath(algoId)
}

function paintDiscovery(status) {
    if (status.Data.length === 0) {
        for (let i = 0; i < status.Path.length; i++) {
            setTimeout(() => {
                paintPath(status.Path, status.EndId)
            }, 10 * i);
        }
    } else {
        for (let i = 1; i <= status.Data.length; i++) {
            if (i === status.Data.length) {
                setTimeout(() => {
                    paintPath(status.Path, status.EndId)
                }, 10 * i);
            }
            let node = status.Data[i]
            setTimeout(function () {
                document.getElementById(node.Id).style.backgroundColor = "#47acfc"
            }, 10 * i);
        }
    }
}

function paintPath(path, endId) {

    for (let i = 1; i < path.length; i++) {
        let node = path[i]
        setTimeout(function () {
            document.getElementById(node.Id).style.backgroundColor = "yellow"
            document.getElementById(endId).style.backgroundColor = "#f55e5e"
        }, 10 * i);
    }
}

// END

// Cell grid handlers

function mouseDownHandler() {
    isMouseDown = true
    console.log("DOWN: ", isMouseDown)
}

function mouseUpHandler() {
    isMouseDown = false
    console.log("UP: ", isMouseDown)
}

function paintCellHandler(element) {
    if (isMouseDown) {
        element.style.backgroundColor = wallColor;
        wall.push(parseInt(element.id))
    }
}

function paintSingleCellHandler(element) {
    element.style.backgroundColor = wallColor;
    wall.push(parseInt(element.id))
}

// END