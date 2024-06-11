import client from './db.config'; 

let isActive = false; 
const getDB = async () => { 
	console.log("Connecting to Database ..."); 
	if (isActive) { 
		console.log("Database already connected"); 
		return client; 
	} 
	try { 
		await client.connect(); 
		console.log("Database connected"); 
		isActive = true; 
		return client; 
	} catch (err: any) { 
		return client; 
	} 
}  

export default getDB;