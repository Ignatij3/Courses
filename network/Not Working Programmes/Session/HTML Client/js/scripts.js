async function request(url, data, sMethod) {
	var sBody;
    let res;
    if (sMethod != "GET") {
        res = await fetch(url, {
        method: sMethod, // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
          'Content-Type': 'application/json'
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *client
        body: JSON.stringify(data)// body data type must match "Content-Type" header
      });
    } else {
        res = await fetch(url);
    };
	
	let inf = await res;
	console.log(inf);
}

function pop() {
	request(document.getElementById("url").value, document.getElementById("body").value, document.getElementById("methods").value);
}

