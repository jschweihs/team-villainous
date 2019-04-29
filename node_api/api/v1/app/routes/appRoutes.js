'use strict';
module.exports = function(app) {
  var userList = require('../controller/appController');

  // todoList Routes
  app.route('/users')
  // .get(userList.get_users)
  // .post(userList.create_user);

  // app.route('/users/:userId')
  // .get(userList.get_user)
  // .put(userList.update_user)
  // .delete(userList.delete_user);
};