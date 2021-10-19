QrScanner.WORKER_PATH = "../static/qr-scanner-worker.min.js";

const camList = document.getElementById('camList');
const camDisplay = document.getElementById("camDisplay")
const camMessage = document.getElementById("camMessage")

const scanner = new QrScanner(document.createElement("video"), QRResultHandler, error => {});

QrScanner.hasCamera().then( (hasCamera) => {
	if (hasCamera) {
		camMessage.innerText = "Suche nach QR-Codes ..."

		QrScanner.listCameras(true).then( (cameras) => {
			let foundLastSelectedCam = false;
			cameras.forEach(camera => {
				const option = document.createElement('option');
				option.value = camera.id;
				option.text = camera.label;
				camList.add(option);

				if (option.value == window.localStorage.getItem("LastSelectedCam")) {
					foundLastSelectedCam = true;
				}
			});

			if (foundLastSelectedCam) {
				camList.value = window.localStorage.getItem("LastSelectedCam")
			} else {
				window.localStorage.setItem("LastSelectedCam", camList.value)
			}

			scanner.setCamera(camList.value).then( () => {
				scanner.start()
			})
		});

		camList.addEventListener('change', (event) => {
			scanner.setCamera(event.target.value);
			window.localStorage.setItem("LastSelectedCam", event.target.value)
		});

		camDisplay.append(scanner.$canvas);

	} else {
		camMessage.innerText = "Keine Kamera gefunden."
	}
})

var lastResult = null;
var counter = 0;
var scannerStopped = false;

function QRResultHandler(result) {
	if (result == null || scannerStopped) {
		return;
	}

	if (result == lastResult) {
		counter++;
		if(counter == 10) {
			counter = 0;
			lastResult = null;

			scannerStopped = true;
			scanner.stop()

			axios.post("../api/event/1", { Code: result })
				.then( (response) => {
					window.location.replace("/admin");
				})
				.catch((error) => {

					scanner.setCamera(camList.value).then( () => {
						scanner.start();
						scannerStopped = false;
					})
					axiosErrorHandler("Eintragen")(error)
				});
		}
	} else {
		lastResult = result;
		counter = 0;
	}

}

