import { queryDatabase } from "routes/databaseFunctions";

const authorize = function () {
    // auth bearer token
    const getBasicAuthCredentials = (req: any) => {
        const authHeader = req.headers.authorization;
        if (!authHeader || !authHeader.startsWith('Basic ')) {
            throw new Error('Missing or invalid Authorization header');
        }
    
        const base64Credentials = authHeader.split(' ')[1];
        const credentials = Buffer.from(base64Credentials, 'base64').toString('ascii');
        const [username, password] = credentials.split(':');
    
        return { username, password };
    };

    return async (req: any, res: any, next: any) => {
        try {
            const { username: name, password: key } = getBasicAuthCredentials(req);
            const query = {
                text: "SELECT * FROM apikey WHERE name = $1 AND apikey = $2",
                values: [name.toLowerCase(), key]
            }

            if (!name || !key) {
                return res.status(400).json({ message: "Please enter your name and key in auth!" });
            }

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