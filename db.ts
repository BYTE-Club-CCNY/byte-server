import client from './db.config';

const getDB = async () => {
  console.log("Connecting to Database ...");

	try {
		await client.connect();
		console.log("Database connected");
		return client;
	} catch (err: any) {
		return client;
	} 
}

export default getDB;