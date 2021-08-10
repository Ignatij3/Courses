async function request(url, data) {
	let res = await fetch(url, {
		method: met, // *GET, POST, PUT, DELETE, etc.
		mode: 'cors', // no-cors, *cors, same-origin
		cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
		credentials: 'same-origin', // include, *same-origin, omit
		headers: {
		  'Content-Type': 'application/json'
		  // 'Content-Type': 'application/x-www-form-urlencoded',
		},
		redirect: 'follow', // manual, *follow, error
		referrerPolicy: 'no-referrer', // no-referrer, *client
		body: JSON.stringify(data) // body data type must match "Content-Type" header
	  });
	let inf = await res;
	console.log(inf);
}
met = "GET";
function methodSelect() {
	console.log(met);
	met = document.getElementById("method").value;
	console.log(met);
	console.log(met === "POST");
	document.getElementById("Post").hidden = met !== "POST";
	document.getElementById("Put").hidden = met !== "PUT";
	document.getElementById("Delete").hidden = met !== "DELETE";
}

function pop() {
	
	request(document.getElementById("url").value, document.getElementById("body").value);
}
