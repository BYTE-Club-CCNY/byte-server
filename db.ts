import client from './db.config';

const activateDb = async () => {
    console.log("Connecting to Database ...");

    try {
      await client.connect();
      console.log("Database connected");
      return client;
    } catch (err: any) {
      throw new Error(`Database connection error\n ${err.message}`);
    } 
}

export default activateDb;