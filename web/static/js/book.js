// change form action last path route (e.g. "smth" in action "/path/to/smth")
function changeFormAction(formElem, newLastRoute) {
	let formAction = formElem.action.split("/")
	formAction[formAction.length - 1] = newLastRoute
	formElem.action = formAction.join("/")
}

// +-------------------+
// + Movie back button +
// +-------------------+

document.addEventListener("DOMContentLoaded", function () {
	let backButtonURL = localStorage.getItem("book-back-url");

	if (backButtonURL == null) {
		backButtonURL = "/"
	}

	let backButton = document.getElementById("back-button")
	backButton.setAttribute("href", backButtonURL)
});
