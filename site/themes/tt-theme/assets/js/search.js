// the length of the excerpts
const contextDive = 40;
let timerUserInput = false;

document.addEventListener("DOMContentLoaded", function () {
    const searchInputs = document.getElementsByClassName("search-query");
    Array.from(searchInputs).forEach(searchInput => {
        searchInput.addEventListener("keyup", function () {
            // don't start searching every time a key is pressed,
            // wait a bit till users stops typing
            if (timerUserInput) { clearTimeout(timerUserInput); }
            timerUserInput = setTimeout(
                function () {
                    search(searchInput);
                },
                500
            );
        });
    });
});

function search(el) {
    // clear previous search results
    let searchResults = el.parentNode.querySelector(".search-results");
    let searchQuery = el.value.trim();
    while (searchResults.firstChild) {
        searchResults.removeChild(searchResults.firstChild);
    }

    // ignore empty and short search queries
    if (searchQuery.length === 0 || searchQuery.length < 2) {
        searchResults.style.display = "none";
        return;
    }

    searchResults.style.display = "block";

    // load your index file
    getJSON("/index.json", function (contents) {
        var results = [];
        let regex = new RegExp(searchQuery, "i");
        // iterate through posts and collect the ones with matches
        contents.forEach(function (post) {
            // here you can also search in tags, categories
            // or whatever you put into the index.json layout
            if (post.title.match(regex) || post.content.match(regex)) {
                results.push(post);
            }
        });

        if (results.length > 0) {
            searchResults.appendChild(
                htmlToElement("<div class='search-summary'>Found: ".concat(results.length, " &middot; <a class='clear-search' href='#' onclick='clearSearchInput()'>clear</a></div>"))
            );

            // populate search results block with excerpts around the matched search query
            results.forEach(function (value, key) {
                let firstIndexOf = value.content.toLowerCase().indexOf(searchQuery.toLowerCase());
                let lastIndexOf = firstIndexOf + searchQuery.length;
                let spaceIndex = firstIndexOf - contextDive;
                if (spaceIndex > 0) {
                    spaceIndex = value.content.indexOf(" ", spaceIndex) + 1;
                    if (spaceIndex < firstIndexOf) { firstIndexOf = spaceIndex; }
                    else { firstIndexOf = firstIndexOf - contextDive / 2; }
                }
                else {
                    firstIndexOf = 0;
                }
                let lastSpaceIndex = lastIndexOf + contextDive;
                if (lastSpaceIndex < value.content.length) {
                    lastSpaceIndex = value.content.indexOf(" ", lastSpaceIndex);
                    if (lastSpaceIndex !== -1) { lastIndexOf = lastSpaceIndex; }
                    else { lastIndexOf = lastIndexOf + contextDive / 2; }
                }
                else {
                    lastIndexOf = value.content.length - 1;
                }

                let summary = value.content.substring(firstIndexOf, lastIndexOf);
                if (firstIndexOf !== 0) { summary = "...".concat(summary); }
                if (lastIndexOf !== value.content.length - 1) { summary = summary.concat("..."); }

                let div = "".concat("<div class='search-result' id='search-summary-", key, "'>")
                    .concat("<a href='", value.permalink, "'>")
                    .concat("<div class='search-result-title'>", value.title, "</div>")
                    .concat("<div class='search-result-summary'>", summary, "</div>")
                    .concat("</a></div>");
                searchResults.appendChild(htmlToElement(div));

                // optionaly highlight the search query in excerpts using mark.js
                new Mark(document.getElementById("search-summary-" + key))
                    .mark(searchQuery, { "separateWordSearch": false });
            });
        }
        else {
            searchResults.appendChild(
                htmlToElement("<div class='search-summary'>Nothing found. <a class='clear-search' href='#' onclick='clearSearchInput()'>clear</a></div>")
            );
        }
    });
}

function getJSON(url, fn) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", url);
    xhr.onload = function () {
        if (xhr.status === 200) {
            fn(JSON.parse(xhr.responseText));
        }
        else {
            console.error(
                "Error processing ".concat(url, ": ", xhr.status)
            );
        }
    };
    xhr.onerror = function () {
        console.error("Connection error: ".concat(xhr.status));
    };
    xhr.send();
}

// it is faster (more convenient)
// to generate an element from the raw HTML code
function htmlToElement(html) {
    let template = document.createElement("template");
    html = html.trim();
    template.innerHTML = html;
    return template.content.firstChild;
}

function clearSearchInput() {
    const searchInputs = document.getElementsByClassName("search-query");
    Array.from(searchInputs).forEach(input => { input.value = ""; search(input) });
};
