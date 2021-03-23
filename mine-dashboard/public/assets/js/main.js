const STATUS_STARTING = "starting"
const STATUS_STOPPING = "stopping"
const STATUS_STOPPED = "stopped"
const STATUS_RUNNING = "running"

const STATUS_UPDATE_INTERVAL = 6 * 1000;

window.addEventListener('load', () => {
    updateServerStatus();
    setInterval(updateServerStatus, STATUS_UPDATE_INTERVAL);

    const startBtn = document.getElementById("start-button")
    const stopBtn = document.getElementById("stop-button")

    startBtn.onclick = () => confirm("Do you realy want to play ?") && doStart()
    stopBtn.onclick = () => confirm("Do you want to stop server ?") && doStop()
});


async function updateServerStatus() {
    const statusEl = document.getElementById("status-fld")
    const buttons = [
        {name: 'start', el: document.getElementById("start-button")},
        {name: 'stop', el: document.getElementById("stop-button")},
        {name: 'wait', el: document.getElementById("wait-button")},
    ]

    const statusRes = await fetch("/status")
    const statusJson = await statusRes.json()
    const status = statusJson.status

    statusEl.textContent = status
    statusEl.className = "status-"+status

    buttons.forEach((o) => {
        if(status == STATUS_STOPPED && o.name == "start") {
            o.el.classList.remove("hidden")
            return
        }

        if(status == STATUS_RUNNING && o.name == "stop") {
            o.el.classList.remove("hidden")
            return
        }

        if(status == STATUS_STARTING || status == STATUS_STOPPING){
            if (o.name == "wait") {
                o.el.classList.remove("hidden")
                return
            }
        }

        o.el.classList.add("hidden")
    })
}

async function doStop() {
    console.log("Stopping")
    await fetch("/stop")
}

async function doStart() {
    console.log("Starting")
    await fetch("/start")
}
