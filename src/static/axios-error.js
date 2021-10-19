function axiosErrorHandler(action) {
	return (error) => {
		if("response" in error) { // if the error is axios-generated
			alert("Fehler beim " + action + "!\n"+axiosErrorString(error.response));
		} else {
			alert("Fehler beim " + action + "!\n"+error.message);
		}
		console.error(error);
	}
}

function axiosErrorString(response) {
	if (!response) {
		return "Keine Antwort erhalten";
	}

	let errorstr = "";
	if (response.status) {
		errorstr = "Status: " + response.status;
	}
	if (response.data) {
		errorstr += "\nAntwort: " + response.data;
	}
	return errorstr;
}
