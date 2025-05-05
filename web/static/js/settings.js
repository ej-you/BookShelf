// +----------------+
// + Theme and font +
// +----------------+

document.addEventListener("DOMContentLoaded", function () {
	let theme = getTheme();

	document.querySelectorAll(".theme-setting").forEach((themeSettingButton) => {
		if (themeSettingButton.textContent === theme) {
			themeSettingButton.classList.add("selected")
			themeSettingButton.classList.remove("default")
		} else {
			themeSettingButton.classList.add("default")
			themeSettingButton.classList.remove("selected")
		}
	});

	let font = getFont();

	document.querySelectorAll(".font-setting").forEach((fontSettingButton) => {
		if (fontSettingButton.textContent === font) {
			fontSettingButton.classList.add("selected")
			fontSettingButton.classList.remove("default")
		} else {
			fontSettingButton.classList.add("default")
			fontSettingButton.classList.remove("selected")
		}
	});
});


function themeChange(buttonElem) {
	console.log(buttonElem)

	document.querySelectorAll(".theme-setting").forEach((themeSettingButton) => {
		if (themeSettingButton.textContent === buttonElem.textContent) {
			themeSettingButton.classList.add("selected")
			themeSettingButton.classList.remove("default")
		} else {
			themeSettingButton.classList.add("default")
			themeSettingButton.classList.remove("selected")
		}
	});

	localStorage.setItem("theme", buttonElem.textContent);
}

function fontChange(buttonElem) {
	console.log(buttonElem)

	document.querySelectorAll(".font-setting").forEach((fontSettingButton) => {
		if (fontSettingButton.textContent === buttonElem.textContent) {
			fontSettingButton.classList.add("selected")
			fontSettingButton.classList.remove("default")
		} else {
			fontSettingButton.classList.add("default")
			fontSettingButton.classList.remove("selected")
		}
	});

	localStorage.setItem("font", buttonElem.textContent);
	applyFont(buttonElem.textContent)
}