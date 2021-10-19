const code = window.localStorage.getItem("Code");
document.getElementById("Code").innerText = code;

new QRCode(document.getElementById("qrcode"), code);
