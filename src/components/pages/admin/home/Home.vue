<template>
	<div class="content home">
		<div class="admin-nav">
			<router-link class="admin-link" to="/admin/users" tag="button">Manage Users</router-link>
			<router-link class="admin-link" to="/admin/blog" tag="button">Manage Blog</router-link>
		</div>
		<div class="roles">
			<h1>Roles</h1>
			<roles-list 
				:roles="roles"
				@add="addRole"
				@save="updateRole"
				@remove="removeRole"
			/>
		</div>
	</div>
</template>

<script>

	import RolesList from './../../../roles/RolesList.vue';

	export default {
		components: {
			RolesList
		},
		computed: {
			roles() {
				return this.$store.getters.roles;
			}
		},
		methods: {
			addRole(name) {
				this.$store.dispatch('addRole', name);
			},
			updateRole(role) {
				this.$store.dispatch('updateRole', role);
			},
			removeRole(id) {
				this.$store.dispatch('removeRole', id);
			}
		},
		beforeCreate() {
			if(!this.$store.getters.jwt) {
				this.$router.push('/admin');
			}
		},
		created() {
			this.$store.dispatch('getRoles');
		}
	}

</script>

<style scoped>

	.home {
		width: 50%;
	}
	
	h1 {
		margin: 0;
		border-bottom: 1px solid #444;
	}

	.admin-link {
		display: block;
		width: 100%;
		font-size: 24px;
		padding: 10px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 0;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	font-family: Nixie One,cursive;
		color: white;
		background-color: #ffc200;
		border: 0;
		height: 60px;
		cursor: pointer;
		border-radius: 8px;
	}

	.roles {
		background-color: #333;
		padding: 10px;
		padding-top: 20px;
		border-radius: 20px;
	}
	.flex {
		display: flex;
	}
	.two-column {
		flex: 50%;
	}

	@media screen and (max-width: 800px) {
		.two-column {
			flex: 100%;
		}
	}
	
</style>