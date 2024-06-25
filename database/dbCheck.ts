import getDB from "./db";
import client from "./db.config";

function runWithTimeout(fn: () => void, timeout: number, interval: number = 100) {
  return new Promise<void>((resolve, reject) => {
    const startTime = Date.now();

    fn();

    const handle = setInterval(() => {
      const elapsedTime = Date.now() - startTime;
      
      if (elapsedTime >= timeout) {
        clearInterval(handle);
        resolve();
      // } else {
      //   try {
      //     fn();
      //   } catch (error) {
      //     clearInterval(handle);
      //     reject(error);
      //   }
      }
    }, interval);
  });
}


const args: string[] = process.argv;
const timeout: number = parseInt(args[2]);

await runWithTimeout(getDB, timeout);