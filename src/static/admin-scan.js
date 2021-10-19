QrScanner.WORKER_PATH = "../static/qr-scanner-worker.min.js";

const camList = document.getElementById('camList');
const camDisplay = document.getElementById("camDisplay")
const camMessage = document.getElementById("camMessage")

const scanner = new QrScanner(document.createElement("video"), result => console.log(result), error => console.log("Error!"));

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
		});

		camList.addEventListener('change', (event) => {
			scanner.setCamera(event.target.value);
			window.localStorage.setItem("LastSelectedCam", event.target.value)
		});

		camDisplay.append(scanner.$canvas);

		scanner.start().then( () => {
			scanner.setCamera(camList.value)
		})

	} else {
		camMessage.innerText = "Keine Kamera gefunden."
	}
})


