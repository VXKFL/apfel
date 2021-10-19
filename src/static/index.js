if (!window.localStorage.getItem("Code")) {
	window.alert("Du bist noch nicht registriert.")
	window.location.replace("/register")
}

const code = window.localStorage.getItem("Code");
document.getElementById("Code").innerText = code;

new QRCode(document.getElementById("qrcode"), code);
