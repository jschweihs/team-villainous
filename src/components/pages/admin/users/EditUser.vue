<template>
	<div class="content">
		<h1>Edit User</h1>
		<div class="new-user-form">
			<user-form 
				v-if="user" 
				:edituser="user" 
				:roles="roles"
				@save="updateUser"
			/>
		</div>
	</div>
</template>

<script>

	import axios from 'axios';

	import UserForm from './../../../users/UserForm.vue';

	export default {
		components: {
			UserForm
		},
		data() {
			return {
				user: null
			};
		},
		computed: {
			roles() {
				return this.$store.getters.roles;
			}
		},
		methods: {
			updateUser(user) {
				this.$store.dispatch('updateUser', user)
				.then(() => {
					this.$router.push('/admin/users');
				})
				.catch(e => console.log(e));
			}
		},
		// beforeCreate() {
		// 	if(!this.$store.getters.jwt) {
		// 		this.$router.push('/admin');
		// 	}
		// },
		created() {
			this.$store.dispatch('getRoles');

			axios.get('http://teamvillainous.com/api/v1/user/get.php?id=' + this.$route.params.id)
			.then(response => {
				this.user = { ...response.data };
			})
			.catch(e => {
				console.log(e);
			})
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
</style>