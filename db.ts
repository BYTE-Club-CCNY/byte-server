import client from './db.config'; 
import checkDB from 'dbChecker';

const getDB = async () => { 
	console.log("Connecting to Database ..."); 
	
	if(client.connected){
		try { 
			await client.connect(); 
			console.log("Database connected"); 
			return client; 
		} catch (err: any) { 
			throw Error("Client Not Found"); 
		} 	
	}
	return client;
}  

export default getDB;