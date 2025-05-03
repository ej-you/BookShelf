const minWidth = 750
let iconIsSet = false

// +---------------+
// + Error message +
// +---------------+

document.addEventListener("DOMContentLoaded", function () {
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);

	let errorMessage = searchParams.get("message")
	let errorStatusCode = searchParams.get("statusCode")

	// if error message query-params is presented
	if (errorMessage !== null && errorStatusCode !== null) {
		searchParams.delete("message")
		searchParams.delete("statusCode")

		// if other query-params is presented
		if (searchParams.toString() !== "") {
			window.history.replaceState({}, document.title, `${location.pathname}?${searchParams.toString()}`);
		} else {
			window.history.replaceState({}, document.title, location.pathname);
		}

		const elemErrorMessage = document.getElementById("error-message");
		const elemErrorCode = document.getElementById("error-message-code");
		const elemErrorText = document.getElementById("error-message-text");

		elemErrorCode.textContent = errorStatusCode
		elemErrorText.textContent = errorMessage
		elemErrorMessage.style.display = "block"

		setTimeout(() => {
			console.log("Hide error message")
			elemErrorMessage.style.display = "none"
		}, 3000);
	}
});

// +-----------+
// + For forms +
// +-----------+

// add page GET query-params to POST request
function addQueryParamsToForm(formElem) {
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);
	formElem.action = formElem.action + "?" + searchParams.toString();
}

// add current page path with query-params like hidden input to form with POST request
function addHiddenNextAttributeToForm(formElem) {
	let currentPage = window.location.pathname + window.location.search;
	// create hidden input with current page path
	let input = document.createElement("input");
	input.type = "hidden";
	input.name = "next";
	input.value = currentPage;
	// add input to form
	formElem.appendChild(input);
}

// +-------------------+
// + Book back button +
// +-------------------+

document.addEventListener("DOMContentLoaded", function () {
	let currentPath = window.location.pathname;
	let currentQuery = window.location.search;

	// skip for book page
	if (currentPath.startsWith("/library/book")) {
		return
	}

	localStorage.setItem("book-back-url", currentPath + currentQuery);
});
