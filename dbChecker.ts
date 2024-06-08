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
const INTERVAL = secondsToMs(60 * 60);
const TIMEOUT = secondsToMs(30);

// threading should happen at top level of server "setInterval"
function checkDB() {
    const database = spawn("bun", ["dbCheck.ts", TIMEOUT.toString()]);
    database.stdout.on("data", (data) => {
        console.log(data.toString());
    });

    database.on("exit", (code) => {
        if (code === 1) {
            console.log("DB is down");
        } else {
            console.log("DB is all good");
        }
    });
}

setInterval(checkDB, INTERVAL);
