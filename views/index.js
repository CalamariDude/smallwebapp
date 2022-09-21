let addressbox = document.getElementById("address-box")

async function getAddresses() {
	let response = await fetch("http://localhost:3000/api/addresses");
	let data = await response.json()
	return data
}

getAddresses().then(data => {
	addresses = data.addresses
	for (let i = 0; i < addresses.length; i++) {
		addressElement = document.createElement("p")
		addressElement.append(document.createTextNode(addresses[i]))
		addressbox.append(addressElement)
	}
})

