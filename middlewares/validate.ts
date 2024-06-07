//validate all fields and their types
function validating(keys: string[], values: any[], requiredFields: any, res: any) {
    for (let i = 0; i < values.length; i++) {
        //initial check if the field is a required field
        if (!(keys[i] in requiredFields)) {
            return res.status(400).json({ message: `Please insert a valid field name, ${keys[i]} is invalid`});
        }
        //check for the correct typing
        if (typeof values[i] !== requiredFields[keys[i]]) {
            return res.status(400).json({ message: `Please insert the correct typing for ${keys[i]}, it should be a ${requiredFields[keys[i]] === "object" ? "array of strings" : requiredFields[keys[i]]}!`});
        }
    }
    // if no response is sent meaning all validations passed, return false at the end of the function
    return false;
}
 
const validate = (req: any, res: any, next: any) => {

    const requiredFields = {name: "string", "short-desc": "string", "long-desc": "string", team: "object", link: "string", image: "string", "tech-stack": "object", cohort: "string", topic: "object"};
    const values = Object.values(req.body);
    const keys = Object.keys(req.body).toString().toLowerCase().split(",");

    if (req.method === "POST") {
        //initial check for empty request body
        if (Object.keys(req.body).length === 0) {
            return res.status(400).json({ message: "Please insert a object with all required fields!" });
        }
        if (keys.length !== 9) {
            return res.status(400).json({ message: "Please insert all required fields, you are missing some fields!" });
        }
        else {
            if (validating(keys, values, requiredFields, res)) {
                return;
            }
        }
    }
    else {
        //initial check for empty request body
        if (Object.keys(req.body).length === 0) {
            return res.status(400).json({ message: "Please insert a object to update!" });
        }
        else {
            if (validating(keys, values, requiredFields, res)) {
                return;
            }
        }
    }
    next();
};

export default validate;