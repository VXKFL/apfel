const registerForm = document.getElementById("registerForm");
registerForm.addEventListener("submit", (e)=>{
	e.preventDefault();

	// Disable submit button for 2 seconds
	const submitButton = document.getElementById("submitButton");
	submitButton.disabled = true;
	setTimeout(() => submitButton.disabled = false, 2000);

	let json = buildJson(registerForm);

	const validationResult = validateJson(json)

	if (validationResult != "") {
		window.alert(validationResult)
		return
	}

	delete json.Password2

	axios.post("api/register", json)
		.then( (response) => {
			console.log(response.data)
		})
});

function buildJson(form) {
	const json = {};
	const formData = new FormData(form);

	for(const pair of formData) {
		json[pair[0]] = pair[1];
	}

	return json;
}

function validateJson(json) {

	// Check if any field is empty
	if (!json.Password || !json.Password2 || !json.Name || !json.Email) {
		return "Bitte fülle alle Felder aus.";
	}

	// Check if passwords match
	if (json.Password != json.Password2) {
		return "Deine Passwörter stimmen nicht überein.";
	}

	// Check if password is to short
	if (json.Password.length < 6 ) {
		return "Dein Passwort muss mindestens sechs Zeichen lang sein."
	}

	// Regex email pattern, refer to https://www.regular-expressions.info/email.html
	// notice conversion to EcmaScript
	const emailPattern = /^(?=[a-z0-9@.!#$%&'*+\/=?^_‘{|}~-]{6,254}$)(?=[a-z0-9.!#$%&'*+\/=?^_‘{|}~-]{1,64}@)[a-z0-9!#$%&'*+\/=?^_‘{|}~-]+(?:\.[a-z0-9!#$%&'*+\/=?^_‘{|}~-]+)*@(?:(?=[a-z0-9-]{1,63}\.)[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+(?=[a-z0-9-]{1,63}$)[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/
	if (!emailPattern.test(json.Email)) {
		return "Deine E-Mail-Adresse hat ein ungültiges Format.";
	}

	return ""
}
