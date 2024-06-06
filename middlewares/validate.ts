//so the error that is occurring is due to the fact that initially it would return a field is a invalid field but hten it would give me the error saying that we try to write after the response ahs ended
function validating(keys: string[], values: any[], requiredFields: string[], res: any) {
    for (let index = 0; index < keys.length; index++) {
        //check if the field is a required field, if not return a 400 status code
        if (!requiredFields.includes(keys[index])) {
            return res.status(400).json({ message: `${keys[index]} is a invalid field` });
        };

        //at this rate field(s) provided are required field(s) therefore now checking type of its value based on index
        switch (keys[index]) {
            case "name":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "short-desc":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "long-desc":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "team":
                if (!Array.isArray(values[index])) {
                    return res.status(400).json({ message: `${keys[index]} field must be a array of strings` });
                }
                break;
            case "link":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "image":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "tech-stack":
                if (!Array.isArray(values[index])) {
                    return res.status(400).json({ message: `${keys[index]} field must be a array of strings` });
                }
                break;
            case "cohort":
                if (typeof values[index] !== "string") {
                    return res.status(400).json({ message: `${keys[index]} field must be a string` });
                }
                break;
            case "topic":
                if (!Array.isArray(values[index])) {
                    return res.status(400).json({ message: `${keys[index]} field must be a array of strings` });
                }
                break;
        };
    };
    //at this point of code all fields are required and have correct typings, meaning no JSON was sent. if no response is sent, return false at the end of the function
    return false;
};

const validate = (req: any, res: any, next: any) => {

    const requiredFields = ["name", "short-desc", "long-desc", "team", "link", "image", "tech-stack", "cohort", "topic"];
    const values = Object.values(req.body);
    const keys = Object.keys(req.body).toString().toLowerCase().split(",");

    if (req.method === "POST") {
        if (Object.keys(req.body).length === 0) {
            return res.status(400).json({ message: "Please insert a object with all required fields!" });
        }
        //there should be 9 fields so if one is missing return a 400 status code indicating missing fields
        if (keys.length !== 9) {
            return res.status(400).json({ message: "Please insert all required fields, you are missing some fields!" });
        }
        else {
            /**check if all fields are required fields and have the correct typings, 
             * if validating sends a response, return to stop the execution of validate
             **/
            if (validating(keys, values, requiredFields, res)) {
                return;
            }
        }
    }
    // the request method is PUT
    else {
        if (Object.keys(req.body).length === 0) {
            return res.status(400).json({ message: "Please insert a object to update" });
        }
        /**check if all fields are required fields and have the correct typings, 
         * if validating sends a response, return to stop the execution of validate
         **/
        if (validating(keys, values, requiredFields, res)) {
            return;
        }
    }

    //at this rate all fields are filled with the correct typings so pass control to the next middleware to insert into DB
    next();
};

export default validate;