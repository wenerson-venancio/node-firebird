import express from "express";
import cors from "cors";
import { dbOptions } from "./config/database.js";
import firebird from "node-firebird";



const app = express();

//Middleware JSON
app.use(express.json());

//Middleware CORS
app.use(cors());

//Rotas
app.get("/produtos", function (req, res) {

	firebird.attach(dbOptions, function (err, db) {

		if (err)
			return res.status(500).json(err);

		//db Database
		db.query('SELECT * FROM TESTGRUPO', function (err, result) {

			//IMPORTANT: Close conecction
			db.detach();

			if (err) {
				return res.status(200).json(err);
			} else {
				return res.status(200).json(result);
			}

		})
	})

});

app.listen(3000, function () {
	console.log("Servidor no ar")
}) 