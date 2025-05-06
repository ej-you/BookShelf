const searchParams = new URLSearchParams(window.location.search);

// +----------------------------+
// + set and select order links +
// +----------------------------+

function getSortParams() {
	let sortField = searchParams.get("sortField")
	// if sortField is not presented
	if (sortField == null) {
		sortField = "title"
	}

	let sortOrder = searchParams.get("sortOrder")
	// if sortOrder is not presented
	if (sortOrder === null) {
		sortOrder = "asc"
	}

	return [sortField, sortOrder]
}

// update sort links href with selected sort field and order
function updateSortLinksHref(sortField, sortOrder) {
	document.querySelectorAll(".sort-field-link").forEach((sortFieldLink) => {
		sortFieldLink.href = sortFieldLink.href + `&sortOrder=${sortOrder}`
	});
	document.querySelectorAll(".sort-order-link").forEach((sortOrderLink) => {
		sortOrderLink.href = sortOrderLink.href + `&sortField=${sortField}`
	});
}

// add background to selected sort links and setup sort links href
function setupSortLinks() {
	const [sortField, sortOrder] = getSortParams()
	// sort fields
	document.querySelectorAll(".sort-field-link").forEach((sortFieldLink) => {
		if (sortFieldLink.dataset.linkValue === sortField) {
			sortFieldLink.classList.add("sort-link-selected")
		} else {
			sortFieldLink.classList.remove("sort-link-selected")
		}
	});
	// sort orders
	document.querySelectorAll(".sort-order-link").forEach((sortOrderLink) => {
		if (sortOrderLink.dataset.linkValue === sortOrder) {
			sortOrderLink.classList.add("sort-link-selected")
		} else {
			sortOrderLink.classList.remove("sort-link-selected")
		}
	});
	updateSortLinksHref(sortField, sortOrder)
}

document.addEventListener('DOMContentLoaded', setupSortLinks);

// +----------------------------------+
// + checkbox input for genres filter +
// +----------------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let genreList = searchParams.getAll("genres")

	document.querySelectorAll("#filter-genre-list .checkbox-btn input").forEach((genreInput) => {
		if (genreList.includes(genreInput.value)) {
			genreInput.checked = true;
		}
	})
});

// +-------------------------------+
// + number input for year filters +
// +-------------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let yearFrom = parseInt(searchParams.get("yearFrom"))
	if (isNaN(yearFrom) || yearFrom < 0) {
		yearFrom = 0
	}
	let yearTo = parseInt(searchParams.get("yearTo"))
	if (isNaN(yearTo) || yearTo > 2100) {
		yearTo = 2100
	}
	document.getElementById("filter-year-from").value = yearFrom;
	document.getElementById("filter-year-to").value = yearTo;
});

// +-----------------------------+
// + radio input for type filter +
// +-----------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let type = searchParams.get("type")
	if (type === null) {
		type = "all"
	}
	document.querySelectorAll("#filter-type .radio-button input").forEach((typeInput) => {
		if (typeInput.value === type) {
			typeInput.checked = true;
		}
	})
});

// +--------------+
// + query params +
// +--------------+

// add sort query-params values to filter form
document.addEventListener('DOMContentLoaded', () => {
	const [sortField, sortOrder] = getSortParams()
	document.getElementById("sort-field-in-filter").value = sortField
	document.getElementById("sort-order-in-filter").value = sortOrder
});

// add filter query params (from page URL) to sort links (then click to follow link)
document.addEventListener('DOMContentLoaded', () => {
	document.querySelectorAll(".sort-field-link, .sort-order-link").forEach(sortLink => {
		sortLink.addEventListener("click", function (event) {
			// abort link following
			event.preventDefault();
			// get sort link query params
			const linkSearchParams = new URLSearchParams(new URL(sortLink.href).search);
			// update page query params with sort link query params
			searchParams.set("sortField", linkSearchParams.get("sortField"));
			searchParams.set("sortOrder", linkSearchParams.get("sortOrder"));
			// follow the new link
			window.location.href = `${window.location.pathname}?${searchParams.toString()}`;
		});
	});
});
