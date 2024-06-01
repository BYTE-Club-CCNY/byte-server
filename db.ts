import client from './db.config';

let connectedClient: any;

const activateDb = async () => {
    console.log("Connecting to Database ...");

    try {
      await client.connect();
      console.log("Database connected");
      connectedClient = client;
    } catch (err: any) {
      throw new Error(`Database connection error\n ${err.message}`);
    } 
}

export default activateDb;
export { connectedClient };