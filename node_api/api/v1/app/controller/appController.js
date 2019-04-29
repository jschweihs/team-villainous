'use strict';

var User = require('../model/appModel');

exports.get_users = (req, res) => {
  // User.getUsers((err, users) => {
  //   console.log('controller')
  //   if (err) {
  //     res.send(err);
  //   }
  //   console.log('res', users);
  //   res.send(users);
  // });
};

exports.create_user = (req, res) => {
  // var new_user = new User(req.body);
  // // handles null error - check all required fields
  // if(!new_user.username) { 
  //   res.status(400).send({ error:true, message: 'Please provide username' });
  // } else {
  //   User.createUser(new_user, (err, user) => { 
  //   if (err) {
  //     res.send(err);
  //   }  
  //   res.json(user);
  //   });
  // }
};

exports.get_user_by_id = (req, res) => {
  // User.getUserById(req.params.user_id,(err, user) => {
  //   if (err) {
  //     res.send(err);
  //   }
  //   res.json(user);
  // });
};

exports.update_user = (req, res) => {
  // User.updateUser(req.params.user_id, new User(req.body), (err, user) => {
  //   if (err) {
  //     res.send(err);
  //   }
  //   res.json(user);
  // });
};

exports.delete_user = (req, res) => {
  // User.remove( req.params.user_id, (err, user) => {
  //   if (err) {
  //     res.send(err);
  //   }
  //   res.json({ message: 'User successfully deleted' });
  // });
};