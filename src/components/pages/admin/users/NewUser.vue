<template>
	<div class="content">
		<h1>Add New User</h1>
		<div class="new-user-form">
			<user-form 
				:roles="roles" 
				@save="addUser"
			/>
		</div>
	</div>
</template>

<script>

	import UserForm from './../../../users/UserForm.vue';

	export default {
		components: {
			UserForm
		},
		// beforeCreate() {
		// 	if(!this.$store.getters.jwt) {
		// 		this.$router.push('/admin');
		// 	}
		// },
		computed: {
			roles() {
				return this.$store.getters.roles;
			}
		},
		methods: {
			addUser(user) {
				console.log('user', user);
				return this.$store.dispatch('addUser', user)
				.then(() => {
					this.$router.push('/admin/users');
				})
				.catch(e => console.log(e));
			}
		},
		created() {
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
</style>