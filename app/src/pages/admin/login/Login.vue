<template>
	<div class="content login">
		<h1>Admin area</h1>
		<form @submit.prevent="loginUser">

			<label for="email">Email</label>
			<input 
				type="text"
				id="email"
				:class="{error: showError && invalidEmail}"
				v-model="login.email" 
				placeholder="johndoe@gmail.com"
			/>

			<label for="password">Password</label>
			<input 
				type="password" 
				id="password"
				:class="{error: showError && invalidPassword}"
				v-model="login.password" 
				placeholder="p455w0rd"
			/>

			<div class="submit-wrapper">
				<input 
					type="submit"
					id="login-btn"
					value="Login"
					:disabled="invalidEmail || invalidPassword"
				/>
				<div class="submit-overlay" @mouseover="showError = true"></div>
			</div>
			
			<p class="text-error text-center" v-if="showError">{{ error }}</p>

		</form>

	</div>
</template>

<script>

	import axios from 'axios';

	import Cookie from './../../../utils/Cookie';

	export default {
		data() {
			return {
				login: {
					email: 			'',
					password: 		'',
				},
				error: 				'Email and/or password are missing or invalid',
				// UI Control
				showError: 			false
			}
		},
		computed: {
			showModal() {
				return this.$store.getters.showModal;
			},
			invalidEmail() {
				return !/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(this.login.email);
			},
			invalidPassword() {
				return this.login.password == '';
			},
			invalidForm() {
				return this.invalidEmail || this.invalidPassword;
			}
		},
		watch: {
			invalidForm(v) {
				this.error = v ? this.error : '';
			}
		},
		methods: {
			loginUser() {
				this.showError = false;
				this.$store.dispatch('showModal', true);
				this.$store.dispatch('loginUser', this.login)
				.then(res => {
						this.$store.dispatch('showModal', false);
						this.$router.push('/admin');
				})
				.catch(error => {
					this.$store.dispatch('showModal', false);
					this.showError = true;
					if(error.response && error.response.data && error.response.data.error) {
						this.error = error.response.data.error;
						this.showError = true;
					}
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
		color: white;
		display: block;
		font-size: 18px;
		margin-bottom: 4px;
		padding-left: 8px;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
	}

	input, select, textarea {
		display: block;
		width: 100%;
		font-size: 20px;
		font-family: inherit;
		padding: 12px 24px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 2px solid transparent;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
		box-sizing: border-box;         /* Opera/IE 8+ */
		transition: all .25s;
	}

	textarea {
		height: 200px;
	}

	input:focus, select:focus, textarea:focus {
		border: 2px solid var(--color-secondary);
		outline:none;
	}

	.submit-wrapper {
		position: relative;
	}

	.submit-overlay {
		position: absolute;
		width: 100%;
		height: 100%;
		z-index: 1;
		left: 0;
		top: 0;
	}

	.text-center {
		text-align: center;
	}

	.text-error {
		color: var(--color-error);
	}
	.error {
		border: 2px solid var(--color-error);
	}
	
	input[type=submit] {
		color: white;
		margin-top: 20px;
		background-color: #ffc200;
		border: 0;
		height: 60px;
		cursor: pointer;
		transform: translateY(0);
		transition: all .25s;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
	}

	input[type=submit]:hover {
		background-color: var(--color-secondary-light);
		transform: translateY(-2px);
	}

	input[type=submit]:disabled {
		background-color: #BEBEBE;
		transform: translateY(0);
		pointer-events: auto;
	}

	input[type=submit]:not(:disabled) + .submit-overlay{
		display: none;
	}

	input[type=submit]:disabled + .submit-overlay{
		display: block;
	}

	@media only screen and (max-width: 600px) {
		.login {
			width: 100%;
		}
	}


.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
	
</style>