import client from './db.config'; 

// attempt to establish connection
// if exception of connection already established exists, return client
// if no connection can be made, return null
const getDB = async () => { 
	console.log("Connecting to Database ..."); 
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