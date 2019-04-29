<template>
	<div>
		<h2>{{ role }}</h2>
	    <div class="team">
	    	<team-member 
	    		v-if="group" 
	    		v-for="user in group.users"
	    		:user="user"
	    		@remove="$emit('remove', user.id)"
	    		:admin="admin"
	    	/>
	    </div>
    </div>
</template>

<script>

	import TeamMember from './TeamMember.vue';

	export default {
		components: {
			TeamMember
		},
		props: {
			group: {
				type: Array,
				required: true
			},
			roles: {
				type: Array,
				required: true
			},
			admin: {
				type: Boolean,
				required: false
			}
		},
		computed: {
			role() {
				for(let i=0; i<this.roles.length; i++) {
					const r = this.roles[i];
					if (r.id == this.group.name) {
						return r.name;
						throw 0;
					}
				}
				return "Others"; // All users should have a role but this is default
			}
		}
	}
</script>

<style scoped>

h2 {
	color: white;
	text-align: center;
}

.team {
	width: 100%;
	display: -webkit-box;
	display: -moz-box;
	display: -ms-flexbox;
	display: -moz-flex;
	display: -webkit-flex;
	display: flex;
	-webkit-flex-direction: row;
	flex-direction: row;
	flex-wrap: wrap;
	-webkit-justify-content: space-around;
	justify-content: space-around;
	-webkit-align-content: space-around;
	align-content: space-around;
	margin-bottom: 40px;
}

</style>