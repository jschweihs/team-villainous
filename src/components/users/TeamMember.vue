<template>
	<div class="team-m">
        <div class="team-member">
        	<h3 class="team-member-title">ViL_{{ user.username }}</h3>
        	<img 
        		width="260" 
        		height="260" 
        		:src="'/images/users/ViL_' + user.username + '.jpg'"
        		:alt="'ViL_' + user.username"
        	/>
        	<h3 class="team-member-name">{{ name }}</h3>
        	<h4 class="team-member-loc">{{ location }}</h4>
        	<p class="team-member-description">{{ user.description }}</p>
        	<div class="member-links">
        		<a
        			:href="'https://www.facebook.com/' + user.facebook_url"
        			target="_blank"
        			v-if="user.facebook_url != ''"
        		>
        			<i class="fab fa-facebook"></i>
        		</a>
        		<a
        			:href="'https://www.twitter.com/' + user.twitter_url"
        			target="_blank"
        			v-if="user.twitter_url != ''"
        		>
        			<i class="fab fa-twitter"></i>
        		</a>
        		<a
        			:href="'https://www.instagram.com/' + user.instagram_url"
        			target="_blank"
        			v-if="user.instagram_url != ''"
        		>
        			<i class="fab fa-instagram"></i>
        		</a>
        		<a
        			:href="'http://twitch.tv/' + user.twitch_url"
        			target="_blank"
        			v-if="user.twitch_url != ''"
        		>
        			<i class="fab fa-twitch"></i>
        		</a>
        		<a
        			:href="user.youtube_url"
        			target="_blank"
        			v-if="user.youtube_url != ''"
        		>
        			<i class="fab fa-youtube"></i>
        		</a>
        		<a
        			:href="user.other_url"
        			target="_blank"
        			v-if="user.other_url != ''"
        		>
        			<i class="fas fa-home"></i>
        		</a>
        	</div>
        	<div class="admin-links" v-if="admin">
    			<router-link
    				tag="a"
    				:to="'/admin/users/' + user.id"
    			>
    				<i class="far fa-edit"></i>
    			</router-link>
    			<button
    				@click="remove"
    			>
    				<i class="far fa-trash-alt"></i>
    			</router-link>
    		</div>
        </div>
    </div>
</template>

<script>
	export default {
		props: {
			user: {
				type: Object,
				required: true
			},
			admin: {
				type: Boolean,
				required: false
			}
		},
		computed: {
			name() {
				return this.user.m_name 
					? this.user.f_name + ' ' + this.user.m_name + ' ' + this.user.l_name 
					: this.user.f_name + ' ' + this.user.l_name;
			},
			location() {
				console.log(this.user);
				let loc = '';
				if(this.user.city) {
					loc += this.user.city + ', ';
				}
				if(this.user.province) {
					loc += this.user.province + ', ';
				}
				if(this.user.country) {
					loc += this.user.country;
				}
				return loc;
			}
		},
		methods: {
			remove() {
				if(confirm('Are you user you want to remove this user?')) {
					this.$emit('remove', this.user.id)
				}		
			}
		}
	}
</script>

<style scoped> 

	.team-m {
    	flex-grow: 1;
    	flex-shrink: 1;
    	flex-basis: auto;
    	margin-bottom: 20px;
  	}

	.team-member {
		width: 260px;
    	margin: 0 auto;
    	padding: 20px;
    	border: 1px solid #444;
    	border-radius: 20px;
    	background-color: #393939;
	}

	.team-member-title {
	    border-bottom: 1px solid #444;
	    margin: 0 -20px;
	    margin-bottom: 20px;
	    padding-bottom: 5px;
	}

	.team-member img {
		width: 260px;
		border-radius: 20px;
	}

	.team-member h3, h4 {
		text-align: center;
		color: white;

	}

	.team-member-loc {
		font-size: 14px;
		margin: 0;
	}

	.team-member-description {
		height: 150px;
		margin-top: 10px;
	}

	.team-member-name {
		margin-bottom: 10px;
	}

	.member-links {
		display: inline-block;
		padding-left: 10px;
		height: 18px;
	}

	.member-links a {
		padding-right: 10px;
	}

	.admin-links {
		display: inline-block;
		float: right;
		height: 18px;
		padding-right: 10px;
	}

	.admin-links a {
		padding-left: 10px;
	}

	.admin-links button {
		padding-left: 10px;
		background: none;
	    border: 0;
	    color: #ffc200;
	    font-size: 16px;
	    cursor: pointer;
	}

	.admin-links button:focus {
		outline: none;
	}
</style>