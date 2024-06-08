// import { Client } from "pg";
// import activateDb from "./db";
import { spawn } from "child_process";

// IDEA:
// have a thread running parrallel to the server that will check every 1hr
// if the database is still contected, changing the server to use local instead untill the
// new connection is established
//
// fn will try to establish a connection to the db in 30s intervals

function secondsToMs(d: number) {
    return d * 1000;
}

function checkDB() {
    const database = spawn("bun", ["dbCheck.ts"]);
    database.on("exit", (code) => {
        if (code === 1) {
            console.log("DB is down");
        } else {
            console.log("DB is all good");
        }
    });
}
const INTERVAL = secondsToMs(1);

setInterval(checkDB, INTERVAL);
