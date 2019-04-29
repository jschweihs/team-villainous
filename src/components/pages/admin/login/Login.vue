<template>
	<div class="content login">
		<h1>Admin area</h1>
		<form @submit.prevent="login">
			<label>Email</label>
			<input type="text" v-model="email"/>
			<label>Password</label>
			<input type="password" v-model="password"/>
			<input type="submit" value="Login"/>
		</form>
	</div>
</template>

<script>

	import axios from 'axios';

	export default {
		data() {
			return {
				email: 			'',
				password: 		'',
				error: 			'',
				display_error: 	false
			}
		},
		methods: {
			login() {
				const params = {
					email: 		this.email,
					password: 	this.password
				}
				console.log(params);
				axios.post('http://teamvillainous.com/api/v1/user/login.php', params)
				.then(response => {
					console.log("success");
					console.log(response);
					this.$store.dispatch('setAuth', response.data);
					this.$router.push('/admin/home');
				})
				.catch(error => {
					console.log(error);
				})
			}
		}
	}
</script>

<style scoped>

	.login {
		width: 50%;
	}

	label {
		color: #DEDEDE;
		display: block;
		font-size: 18px;
		margin-bottom: 4px;
		padding-left: 8px;
	}

	input, select, textarea {
		display: block;
		width: 100%;
		font-size: 24px;
		font-family: Nixie One,cursive;
		padding: 10px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 0;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
	}

	textarea {
		height: 200px;
	}
	
	input[type=submit] {
		color: white;
		margin-top: 20px;
		background-color: #ffc200;
		border: 0;
		height: 60px;
		cursor: pointer;
	}
</style>