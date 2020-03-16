<template>
	<div>
		<form @submit.prevent="add">
			<input type="text" v-model="name" placeholder="Add new role..."/>
			<input type="submit" value="Add"/></form>
		<ul class="role-list">
			<li 
				v-for="(role, index) in roles"
				:key="role.id"
			>
				<role-row 
					:role="role"
					:index="index"
					@remove="$emit('remove', role.id)"
					@save="save"
				/>
			</li>
		</ul>
	</div>
</template>

<script>
	import RoleRow from './RoleRow.vue';

	export default {
		components: {
			RoleRow
		},
		props: {
			roles: {
				type: Array,
				required: true
			}
		},
		data() {
			return {
				name: ''
			}
		},
		methods: {
			add() {
				this.$emit('add', this.name);
				this.name = ''
			},
			save(role) {
				this.$emit('save', role);
			}
		}
	}
</script>

<style scoped>
	.role-list {
		list-style-type: none;
		padding: 0;
		margin-bottom: 0;
	}

	input, select, textarea {
		display: inline-block;
		margin-left: 10px;
		width: calc(75% - 15px);
		font-size: 24px;
		padding: 10px 20px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 0;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	font-family: Nixie One,cursive;
	}

	input[type=submit] {
		margin: 0 10px;
		width: calc(25% - 20px);
		color: white;
		margin-top: 20px;
		border: 0;
		height: 48px;
		cursor: pointer;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
		background-color: var(--color-secondary);
		transform: translateY(0);
		transition: all .25s;
	}

	input[type=submit]:hover {
		background-color: var(--color-secondary-light);
		transform: translateY(-2px);
	}

</style>