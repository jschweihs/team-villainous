<template>
	<div class="content">
		<h1>Team</h1>
		<p class="page-description">
			If you are interested in joining Team Villainous then we would love to hear from you.  Head over to the Contact page and send us your information.  If we are interested then we will get back to you quickly!
		</p>
		<team-section 
			v-if="user_groups && roles" 
			v-for="group in user_groups"
			:group="group"
			:roles="roles"
			@remove="removeUser"
		/>
	</div>
</template>

<script>

	import TeamSection from './../../users/TeamSection.vue';

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
			}
		},
		methods: {
			removeUser(user_id) {
				this.$store.dispatch('removeUser', user_id);
			}
		},
		created() {
			this.$store.dispatch('getUsers');
			this.$store.dispatch('getRoles');
			// .then(() => {
			// 	let ug = this.$store.getters.user_groups;
			// 	console.log('created');
			// 	console.log(ug);

			// 	ug.forEach(group => {
			// 		console.log('group');
			// 		this.user_groups.push(group);
			// 	})
			// });
		}
	}

</script>