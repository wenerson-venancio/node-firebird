import firebird from "node-firebird";


const dbOptions = {
	host: 'localhost',
	port: 3050,
	database: '//Users//wenersonvenancio//ECO.ECO',
	user: 'SYSDBA',
	password: 'masterkey',
	lowercase_keys: false,
	role: null,
	pageSize: 4096
}

function executeQuery(ssql, params, callback) {


	firebird.attach(dbOptions, function (err, db) {

		if (err)
			return callback(err, []);

		//db Database
		db.query(ssql, params, function (err, result) {


			db.detach(); //IMPORTANT: Close conecction


			if (err) {
				return callback(err, []);
			} else {

				return callback(result);
			}

		});
	});

}

export { executeQuery };