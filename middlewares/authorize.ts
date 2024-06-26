import { queryDatabase } from "routes/databaseFunctions";

const authorize = function () {
    return async (req: any, res: any, next: any) => {
        const name = req.headers.name.toLowerCase();
        const key = req.headers.authorization?.split(" ")[1];
        const query = {
            text: "SELECT * FROM apikey WHERE name = $1 AND apikey = $2",
            values: [name, key]
        }

        if (!name || !key) {
            return res.status(400).json({ message: "Please enter your name and key before accessing the database!" });
        }

        try {
            const result = await queryDatabase(req.client, query.text, query.values);
            if (result.rows.length === 0) {
                return res.status(401).json({ message: "Invalid name or key!" });
            }
            next();
        }
        catch (e: any) {
            return res.status(500).json({ message: e.message });
        }
    }
}

export default authorize;