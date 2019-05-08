<template>
	<div class="content">
		<h1>Manage Users</h1>
		<router-link to="/admin/users/add" tag="button">Add New User</router-link>
		<team-section 
			v-if="user_groups && roles" 
			v-for="group in user_groups"
			:group="group"
			:roles="roles"
			@remove="removeUser"
			admin
		/>
	</div>
</template>

<script>

	import TeamSection from './../../../users/TeamSection.vue';

	export default {
		components: {
			TeamSection
		},
		computed: {
			user_groups() {
				if(this.$store.getters.users) {
					return this.$store.getters.user_groups;
				}
			},
			roles() {
				return this.$store.getters.roles;
			},
			current_user() {
				return this.$store.getters.current_user;
			}
		},
		methods: {
			removeUser(user_id) {
				this.$store.dispatch('removeUser', user_id);
			}
		},
		beforeCreate() {
			if(!this.$store.getters.jwt) {
				this.$router.push('/admin');
			}
		},
		created() {
			this.$store.dispatch('getUsers');
			this.$store.dispatch('getRoles');
		}
	}
	
</script>

<style scoped>
	
	.content {
		background-color: #333;
		border-radius: 20px;
		padding: 20px;
	}

	h1 {
		border-bottom: 1px solid #444;
		margin: 0;
		margin-bottom: 20px;
	}

	button {
		display: block;
		width: 100%;
		height: 60px;
		font-size: 24px;
		padding: 10px;
		margin: 20px 0;
		border-radius: 8px;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	color: white;
		background-color: #ffc200;
		border: 0;
		cursor: pointer;
	}

</style>