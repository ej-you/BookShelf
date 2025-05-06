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

	// skip for book pages
	if (currentPath.startsWith("/book")) {
		return
	}

	localStorage.setItem("book-back-url", currentPath + currentQuery);
});

// +----------------------+
// + Apply theme and font +
// +----------------------+

// Returns list of theme css root vars
function getCssThemeVars() {
	// получаем все переменные, установленные в :root
	const rootStyle = document.documentElement.style;

	let styleList = document.styleSheets[1].cssRules[0].style

	const cssRootVars = [];
	for (let i = 0; i < styleList.length; i++) {
		let styleName = styleList[i];
		if (styleName.startsWith('--theme')) {
			cssRootVars.push(styleList[i].replace(/^--theme-/, ""));
		}
	}

	console.log(cssRootVars);
	return cssRootVars
}

// returns current font name
function getFont() {
	let font = localStorage.getItem("font");

	if (font == null) {
		font = "Sans-Serif"
	}
	return font
}

// returns current theme name
function getTheme() {
	let theme = localStorage.getItem("theme");

	if (theme == null) {
		theme = "Dark"
	}
	return theme
}

// apply font for app
function applyFont(font) {
	document.body.style.fontFamily = `var(--font-${font})`
}

// apply theme for app
function applyTheme(theme) {
	if (theme == "Dark") {
		theme = "dark"
	} else {
		theme = "light"
	}

	let cssRootVars = getCssThemeVars()

	let themeKey;
	let themeValue;
	cssRootVars.forEach(cssVar => {
		themeKey = "--theme-" + cssVar
		themeValue = `var(--${cssVar}-${theme})`
		document.documentElement.style.setProperty(themeKey, themeValue)
	})
}

document.addEventListener("DOMContentLoaded", function () {
	let font = getFont();
	let theme = getTheme();

	applyFont(font)
	applyTheme(theme)
});
