import server from "../server";
import * as chai from 'chai';
import chaiHttp from "chai-http";
process.env.NODE_ENV = "test";
chai.use(chaiHttp);
import { request, expect } from 'chai';

describe("Basic Server Test", () => {
	it("should return 200", (done) => {
		chai.request(server)
		// request(server)
		.get("/")
		.end((err, res) => {
			expect(res).to.have.status(200);
			expect(res.text).to.equal("BYTE @ CCNY");
			done();
		});
	});
});