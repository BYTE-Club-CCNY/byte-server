import client from './db.config'; 

const getDB = async () => { 
	console.log("Connecting to Database ..."); 
	if(client._connected || client._connecting) return client;
	try{
		await client.connect()
		console.log("Connected to Database");
		return client;
	} catch(exception: any) {
		if(exception.toString() === "Error: Client has already been connected. You cannot reuse a client.") {
			console.log("Client Already Connected, returning original client");
			return client;
		}
		return null;
	}
}

export default getDB;