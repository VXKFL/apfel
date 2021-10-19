const registerForm = document.getElementById("registerForm");
registerForm.addEventListener("submit", (e)=>{
	e.preventDefault();

	// Disable submit button for 2 seconds
	const submitButton = document.getElementById("submitButton");
	submitButton.disabled = true;
	setTimeout(() => submitButton.disabled = false, 2000);

	const json = buildJson(registerForm);

	axios.post("api/register", json)
		.then( (response) => {
			console.log(response.data)
		})
});

function buildJson(form) {
	const json = {};
	const formData = new FormData(form);
	console.log(formData)
	for(const pair of formData) {
		json[pair[0]] = pair[1];
	}
	console.log(json)
	return json;
}
