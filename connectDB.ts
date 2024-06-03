import getDB from "./db";
import logger from "./utils/logger";

const connectDB = async (req: any, res: any, next: any) => {
    try {
        req.client = await getDB();
        next(); 
    } catch (err: any) {
        logger.info(err.message);
        next("route");
    }
}

export default connectDB;
