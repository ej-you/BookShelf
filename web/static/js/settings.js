// remove all query-params if exists (for removing passwdChangedOK query-param)
document.addEventListener("DOMContentLoaded", function () {
	const url = new URL(window.location.href);
	url.searchParams.delete("passwdChangedOK");
	window.history.replaceState({}, document.title, url.pathname + url.search);
});
