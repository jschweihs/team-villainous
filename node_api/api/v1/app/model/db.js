'user strict';

var mysql = require('mysql');

//local mysql db connection
var connection = mysql.createConnection({
    host     : 'localhost',
    user     : 'teammkig_admin',
    password : 'TeamVil2018',
    database : 'teammkig_teamvillainous'
});

connection.connect(function(err) {
    if (err) throw err;
});

module.exports = connection;