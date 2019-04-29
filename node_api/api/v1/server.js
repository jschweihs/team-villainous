const express = require('express'),
	app = express(),
	bodyParser = require('body-parser');

const port = process.env.PORT || 3000;

const mysql = require('mysql');

// Connection configurations
const con = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '',
    database: 'mydb'
});
 
// Connect to database
con.connect();

app.listen(port);

console.log('API server started on: ' + port);

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());

var routes = require('./app/routes/appRoutes'); // Importing route
routes(app); // Register the route