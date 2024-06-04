import getDB from "./db";
import logger from "./utils/logger";

let isActive = false;

const connectDB = async (req: any, res: any, next: any) => {
    try {
        if (isActive) {
            logger.info("Database already connected");
            return next();
        }
        req.client = await getDB();
        isActive = true;
        next(); 
    } catch (err: any) {
        logger.info(err.message);
        next("route");
    }
}

export default connectDB;
