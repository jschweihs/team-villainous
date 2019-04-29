<template>
	<div class="content">
		<h1>Team</h1>
		<p class="page-description">
			Our team is very diverse and we are proud to announce that we have a shit ton of people.  It's pretty cool and I think we will get even more if we continue to kick the ass we already kick.
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