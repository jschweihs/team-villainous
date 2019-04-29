import axios from 'axios';

const state = {
	users: 			null,
	current_user: 	null,
	jwt: 			'',
};

const mutations = {
	'ADD_USER' (state, user) {
		console.log('mutation');
		console.log(user);
		state.users.push(user);
	},
	'GET_USERS' (state, users) {
		state.users = users;
	},
	'UPDATE_USER' (state, user) {
		if(state.users) {
			const user_index = state.users.findIndex(u => user.id == u.id);
			state.users[user_index] = user;
		}
	},
	'REMOVE_USER' (state, user_id) {
		state.users = state.users.filter(user => {
			return user.id !== user_id;
		})
	},
	'SET_JWT' (state, jwt) {
		state.jwt = jwt;
	},
	'SET_CURRENT_USER' (state, user) {
		state.current_user = user;
	}
};

const actions = {
	addUser: ({commit}, payload) => {
		const user = payload.user;
		console.log('adding action store user', user);
		axios.post('http://teamvillainous.com/api/v1/user/create.php', { 
			...user,
			created: new Date(),
			updated: new Date()
		})
		.then(response => {
			console.log("User added successfully");
			console.log(response);
			console.log(user);
			commit('ADD_USER', {...user, id: response.data.id});
			// Upload new image if the path was altered
			if(payload.image_path != '') {
				// Save image to server
				axios.post(
					'http://teamvillainous.com/api/v1/file/upload_image.php',
					payload.image,
					{
				    	headers: {
				        	'Content-Type': 'multipart/form-data'
				    	}
				  	}
				)
				.then(response => {
					console.log('image set successfully');
				})
				.catch(e => console.log(e));
			}
			else {
				console.log('no image change!');
			}
		})
		.catch(e => console.log(e));
	},
	getUsers: ({commit}) => {
		console.log('get users in store');
		// Only hit api if users does not exist yet
		if(!this.a.state.users) {
			console.log('hitting api');
			return axios.post('http://teamvillainous.com/api/v1/user/get_all.php')
			.then(response => {
				console.log('Users loaded!');
				commit('GET_USERS', response.data.records);
			})
			.catch(error => {
				console.log(error);
			});
		} else {
			console.log('users already in store, we good homie!');
		}	
	},
	updateUser: ({commit}, payload) => {
		const user = payload.user;
		console.log('store user', user);
		console.log('store image', payload.image);
		console.log('store imagepath', payload.image_path);
		axios.put('http://teamvillainous.com/api/v1/user/update.php', {
			...user,
			updated: new Date()
		})
		.then(response => {
			console.log('User updated successfully');
			console.log(response);
			commit('UPDATE_USER', user);

			console.log(payload.image_path);
			console.log(payload.image);
			// Upload new image if the path was altered
			if(payload.image_path != '') {
				console.log('image changed');
				axios.post(
					'http://teamvillainous.com/api/v1/file/upload_image.php',
					payload.image,
					{
				    	headers: {
				        	'Content-Type': 'multipart/form-data'
				    	}
				  	}
				)
				.then(response => {
					console.log('image set successfully');
				})
				.catch(e => console.log(e));
			}
			else {
				console.log('no image change!');
			}
			
		})
		.catch(e => console.log(e));
	},
	removeUser: ({commit}, user_id) => {
		console.log(user_id);
		axios.post('http://teamvillainous.com/api/v1/user/delete.php', {id: user_id})
		.then(response => {
			console.log('User ' + user_id + ' removed');
			commit('REMOVE_USER', user_id);
		})
		.catch(error => {
			console.log(error);
		})
	},
	setAuth: ({commit}, data) => {
		commit('SET_JWT', data.jwt);
		commit('SET_CURRENT_USER', data.user);
	}
}

const getters = {
	users: state => state.users,
	user_groups: state => {
		let groups = [];
			state.users.forEach(user => {
				let added = false;
				if(groups.length == 0) {
					groups[0] = {
						name: user.role,
						users: [user]
					};
				} else {
					let group_exists = false;
					let i = 0;
					groups.forEach(group => {
						const role = user.role;
						const group_name = group.name;
						if(role == group_name) {
							group_exists = true;
						}
						if(group_exists && !added) {
							groups[i].users.push(user);
							added = true;
						}	
						i++;
					})
					if(!group_exists) {
						groups.push({
							name: user.role,
							users: [user]
						});
						added = true;
					}
				}
			});

		// Hack to make owners appear first
		let i=0;
		groups.forEach(group => {
			if(group.name == 9) {
				const temp_group = groups[0];
				groups[0] = group;
				groups[i] = temp_group;
			}
			i++;
		});
		return groups;
	},
	jwt: state =>  state.jwt,
	current_user: state => state.current_user
}

export default {
	state,
	mutations,
	actions,
	getters
}

