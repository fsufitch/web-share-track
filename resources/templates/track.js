(function() {
    // https://stackoverflow.com/a/3855394
    var QTRACK = "_WST_TID";
    var qs = (function(a) {
	if (a == "") return {};
	var b = {};
	for (var i = 0; i < a.length; ++i)
	{
	    var p=a[i].split('=', 2);
	    if (p.length == 1)
		b[p[0]] = "";
	    else
		b[p[0]] = decodeURIComponent(p[1].replace(/\+/g, " "));
	}
	return b;
    })(window.location.search.substr(1).split('&'));

    // https://stackoverflow.com/a/11654596
    var updateQueryString = function(key, value, url) {
	if (!url) url = window.location.href;
	var re = new RegExp("([?&])" + key + "=.*?(&|#|$)(.*)", "gi"),
        hash;

	if (re.test(url)) {
            if (typeof value !== 'undefined' && value !== null)
		return url.replace(re, '$1' + key + "=" + value + '$2$3');
            else {
		hash = url.split('#');
		url = hash[0].replace(re, '$1$3').replace(/(&|\?)$/, '');
		if (typeof hash[1] !== 'undefined' && hash[1] !== null) 
                    url += '#' + hash[1];
		return url;
            }
	}
	else {
            if (typeof value !== 'undefined' && value !== null) {
		var separator = url.indexOf('?') !== -1 ? '&' : '?';
		hash = url.split('#');
		url = hash[0] + separator + key + '=' + value;
		if (typeof hash[1] !== 'undefined' && hash[1] !== null) 
                    url += '#' + hash[1];
		return url;
            }
            else
		return url;
	}
    }
    
    var updateUrlTrackId = function(responseData){
	var data = JSON.parse(responseData);
	var currTrackId = data['CurrTrackId'];
	var newUrl = updateQueryString(QTRACK, currTrackId, window.location.href);
	window.history.pushState({}, "", newUrl);
    };


    (function() {
	var sharerTrackId = qs[QTRACK] || "";

	var callback = "{{js .Callback}}" + sharerTrackId;
	
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
	    if (xhttp.readyState == 4 && xhttp.status == 200) {
		updateUrlTrackId(xhttp.responseText);
	    } else {
		//alert(xhttp.responseText);
	    }
	}
	xhttp.open("POST", callback, true);
	xhttp.send();
    })();
})()
