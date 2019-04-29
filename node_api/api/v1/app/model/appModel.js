'user strict';
var sql = require('./db.js');

// User object constructor
var User = function(user) {
    this.id             = user.id;
    this.username       = user.username;
    this.password       = user.password;
    this.email          = user.email;
    this.f_name         = user.f_name;
    this.m_name         = user.m_name;
    this.l_name         = user.l_name;
    this.title          = user.title;
    this.address        = user.address;
    this.city           = user.city;
    this.province       = user.province;
    this.zip            = user.zip;
    this.country        = user.country;
    this.birth_date     = user.birth_date;
    this.description    = user.description;
    this.status         = user.status;
    this.created        = new Date();
    this.update         = new Date();
};

User.createUser = (new_user, result) => {    
    sql.query("INSERT INTO users SET ?", newUser, (err, res) => { 
        if(err) {
            console.log("error: ", err);
            result(err, null);
        } else {
            console.log(res.insertId);
            result(null, res.insertId);
        }
    });           
};
User.getUserById = (user_id, result) => {
    sql.query("SELECT * FROM users WHERE id = ? ", user_id, (err, res) => {             
        if(err) {
            console.log("error: ", err);
            result(err, null);
        } else {
            result(null, res);
        }
    });   
};
User.getUsers = (result) => {
    sql.query("SELECT * FROM users", (err, res) => {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        } else {
            console.log('tasks : ', res);  
            result(null, res);
        }
    });   
};
User.updateUser = (id, user, result) => {
    sql.query("UPDATE users SET username = ? WHERE id = ?", [user.username, id], (err, res) => {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        } else {   
            result(null, res);
        }
    }); 
};
User.removeUser = (id, result) => {
    sql.query("DELETE FROM users WHERE id = ?", [id], (err, res) => {
        if(err) {
            console.log("error: ", err);
            result(null, err);
        } else {
            result(null, res);
        }
    }); 
};

module.exports = User;