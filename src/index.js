import express from "express";
import cors from "cors";
import { executeQuery } from "./config/database.js";


const app = express();

//Middleware JSON
app.use(express.json());

//Middleware CORS
app.use(cors());

//Rotas
app.get("/produtos", function (req, res) {


	executeQuery("SELECT * FROM testgrupo", [], function (err, result) {


		if (err) {
			res.status(500).json(err);
		} else {
			res.status(200).json(result);

		}

	});

});

app.listen(3000, function () {
	console.log("Servidor no ar")
});