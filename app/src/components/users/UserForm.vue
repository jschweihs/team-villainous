<template>
	<div>
		<form @submit.prevent="save">
			<div class="flex">
				<div class="profile-pic">
					<div class="picture-preview">
						<label class="img-label">Profile Picture</label>
						<img 
							width="260" 
							height="260" 
							:src="user.profile_picture" 
							onclick="document.getElementById('profile-picture').click()"
						/>
						<label for="profile-picture" class="file-label">Select</label>
						<input 
							type="file" 
							id="profile-picture" 
							accept="image/jpeg"
							ref="profilepicture"
							v-model="image_path"
							@change="uploadImage"
						/>
					</div>
				</div>
				<div class="main-info">
					<label>Username*</label>
					<input type='text' v-model="user.username"/>
					<label>Email*</label>
					<input type='text' v-model="user.email"/>
					<div v-if="!edituser">
						<label>Password*</label>
						<input type='password' v-model="user.password"/>
						<label>Password Confirm*</label>
						<input type='password' v-model="user.password_confirm"/>
					</div>
					<label>Role*</label>
					<select v-model="user.role">
						<option 
							v-for="role in roles"
							:key="role.id"
							:value="role.id"
						>
							{{ role.name }}
						</option>
					</select>
					<label>Title</label>
					<input type='text' v-model="user.title"/>
				</div>
			</div>

			<hr/>

			<div class="flex">
				<div class="two-column mr-10">
					<label>First Name*</label>
					<input type='text' v-model="user.f_name"/>
					<label>Middle Name</label>
					<input type='text' v-model="user.m_name"/>
					<label>Last Name</label>
					<input type='text' v-model="user.l_name"/>
					<label>Description</label>
					<textarea v-model="user.description">{{ description }}</textarea>
				</div>
				<div class="two-column ml-10">
					<label>Date of Birth</label>
					<input 
						type='text'
						placeholder="mm/dd/yyyy"
						v-model="user.birth_date" 
						v-mask="'##/##/####'"
					/>
					<label>Address</label>
					<input type='text' v-model="user.address"/>
					<label>City</label>
					<input type='text' v-model="user.city"/>
					<div v-if="user.country == 'United States'">
						<label>State</label>
						<select v-model="user.province">
							<option 
								v-for="state in states"
								:key="state"
								:value="state"
							>
								{{ state }}
							</option>
						</select>
					</div>
					<div v-else>
						<label>Province</label>
						<input type="text" v-model="user.province"/>
					</div>
					
					<label>Zip code</label>
					<input type='text' v-model="user.zip"/>
					<label>Country</label>
					<select 
						v-model="user.country"
						@change="user.province = ''"
					>
						<option 
							v-for="country in countries"
							:key="country"
							:value="country"
						>
							{{ country }}
						</option>
					</select>
				</div>
			</div>

			<hr/>	

			<label>Facebook</label>
			<input type='text' v-model="user.facebook_url"/>
			<label>Twitter</label>
			<input type='text' v-model="user.twitter_url"/>
			<label>Instagram</label>
			<input type='text' v-model="user.instagram_url"/>
			<label>Twitch</label>
			<input type='text' v-model="user.twitch_url"/>
			<label>Youtube</label>
			<input type='text' v-model="user.youtube_url"/>
			<label>Other</label>
			<input type='text' v-model="user.other_url"/>

			<hr/>

			<label>PS4 Gamertag</label>
			<input type='text' v-model="user.ps4_gamertag"/>
			<label>XBox Gamertag</label>
			<input type='text' v-model="user.xbox_gamertag"/>
			<label>Steam Gamertag</label>
			<input type='text' v-model="user.steam_gamertag"/>

			<hr/>

			<input type="submit" value="Save User"/>

		</form>
	</div>
</template>

<script>

	import {states, countries} from './../../data/geo';

	import {mask} from 'vue-the-mask'

	export default {
		props: {
			edituser: {
				type: Object,
				required: false
			},
			roles: {
				type: Array,
				required: true
			}
		},
		data() {
			return {
				image_path: '',
				image_contents: '',
				states,
				countries,
				user: this.edituser
					? { 
						...this.edituser, 
						profile_picture: '/images/users/ViL_' + this.edituser.username + '.jpg'
					} 
					: {
						username: 			'',
						email: 				'',
						password: 			'',
						profile_picture: 	'/images/users/placeholder.jpg',
						f_name: 			'',
						m_name: 			'',
						l_name: 			'',
						title: 				'',
						address: 			'',
						city: 				'',
						province: 			'',
						zip: 				'',
						country: 			'United States',
						birth_date: 		'',
						description: 		'',
						role: 				'',
						status: 			1,
						facebook_url: 		'',
						twitter_url: 		'',
						instagram_url: 		'',
						twitch_url: 		'',
						youtube_url: 		'',
						other_url: 			'',
						ps4_gamertag: 		'',
						xbox_gamertag: 		'',
						steam_gamertag: 	'',
						created: 			'',
						updated: 			'',
					}
			};
		},
		computed: {
			dob() {
				if(this.user.birth_date) {
					const date_array = this.user.birth_date.split('/');
					return  date_array[2] + '-' + date_array[0] + '-' + date_array[1]
				}
			}
		},
		methods: {
			uploadImage(e) {
				this.image_contents = this.$refs.profilepicture.files[0];
				this.user.profile_picture = URL.createObjectURL(e.target.files[0]);
			},
			save() {
				this.user = { ...this.user, birth_date: this.dob };
				
				// Prep file
				let image = new FormData();
				image.append('image', this.image_contents);
				image.append('folder', 'users');
				image.append('name', 'ViL_' + this.user.username);

				this.$emit('save', { user: this.user, image, image_path: this.image_path });
			}
		},
		created() {
			if(this.user.birth_date) {
				const date_array = this.user.birth_date.split('-');
				this.user.birth_date = date_array[1] + '/' + date_array[2] + '/' + date_array[0]
			}	
		},
		directives: {mask}
	}
</script>

<style scoped>

	.flex {
		display: flex;
	}
	.two-column {
		flex: 50%;
	}

	.ml-10 {
		margin-left: 10px;
	}

	.mr-10 {
		margin-right: 10px;
	}

	hr {
		margin: 20px 0;
		border: 1px solid var(--color-grey-light);
	}

	img {
		width: 260px;
		height: 260px;
		display: block;
		margin: 0 auto;
		border-top-left-radius: 20px;
		border-top-right-radius: 20px;
		cursor: pointer;
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
		padding: 10px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 0;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	font-family: Nixie One,cursive;
	}

	textarea {
		height: 215px;
		resize: none;
	}

	input[type=submit] {
		color: white;
		margin-top: 20px;
		background-color: var(--color-secondary);
		border: 0;
		height: 60px;
		cursor: pointer;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
		border-radius: 8px;
		transform: translateY(0);
		transition: all .25s;
	}

	input[type=submit]:hover {
		transform: translateY(-2px);
		background-color: var(--color-secondary-light);
	}

	.profile-pic {
		width: 340px;
	}

	.img-label {
		margin-left: 40px;
	}

	.main-info {
		width: calc(100% - 340px);
	}

	.file-label {
		display: block;
		cursor: pointer;
		color: white;
		margin: 0 auto;
		background-color: var(--color-secondary);
		border: 0;
		height: 35px;
		cursor: pointer;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
		line-height: 35px;
	    width: 260px;
	    text-align: center;
	    box-sizing: border-box;
	    font-size: 20px;
	    border-bottom-left-radius: 20px;
		border-bottom-right-radius: 20px;
		transition: all .25s;
	}

	.file-label:hover {
		background-color: var(--color-secondary-light);
	}

	input[type=file] {
		opacity: 0;
	   	position: absolute;
	   	z-index: -1;
	}

</style>