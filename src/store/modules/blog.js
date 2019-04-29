import axios from 'axios';

const state = {
	blog: 		null,
};

const mutations = {
	'ADD_ENTRY' (state, entry) {
		state.blog.unshift(entry);
	},
	'SET_BLOG' (state, blog) {
		state.blog = blog;
	},
	'UPDATE_ENTRY' (state, entry) {
		if(state.blog) {
			const entry_index = state.blog.findIndex(e => entry.id == e.id);
			state.blog[entry_index] = entry;
		}
	},
	'REMOVE_ENTRY' (state, entry_id) {
		state.blog = state.blog.filter(e => {
			return entry_id !== e.id;
		});
	}
};

const actions = {
	addEntry: ({commit}, payload) => {
		const entry = payload.entry;
		axios.post('http://teamvillainous.com/api/v1/blog/create.php', entry)
		.then(response => {
			console.log(response);
			commit('ADD_ENTRY', {
				...entry, 
				created: (new Date()).toISOString(), 
				id: response.data.id
			});
			// Upload new image if the path was altered
			if(payload.image_path != '') {
				// Need to append the id because it hasn't existed till right now
				payload.image.append('name', response.data.id);
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
	getBlog: ({commit}) => {
		// Only hit api if blog does not exist yet
		if(!this.a.state.blog) {
			return axios.post('http://teamvillainous.com/api/v1/blog/get_all.php')
			.then(response => {
				commit('SET_BLOG', response.data.entries);
			})
			.catch(error => {
				console.log(error);
			});
		} else {
			console.log('blog already in store, we good homie!');
		}	
	},
	getEntry: ({commit}) => {

	},
	updateEntry: ({commit}, payload) => {
		const entry = payload.entry;
		axios.put('http://teamvillainous.com/api/v1/blog/update.php', entry)
		.then(response => {
			commit('UPDATE_ENTRY', entry);

			// Upload new image if the path was altered
			if(payload.image_path != '') {
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
		})
		.catch(e => console.log(e));
	},
	removeEntry: ({commit}, entry_id) => {
		axios.post('http://teamvillainous.com/api/v1/blog/delete.php', {id: entry_id})
		.then(response => {
			console.log('Entry ' + entry_id + ' removed');
			commit('REMOVE_ENTRY', entry_id);
		})
		.catch(error => {
			console.log(error);
		})
	}
}

const getters = {
	blog: state => state.blog
}

export default {
	state,
	mutations,
	actions,
	getters
}

