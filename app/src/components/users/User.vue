<template>
	<section class="user">
		<h3>{{ user.username }}</h3>
		<p>Email: {{ user.email }}</p>
		<p>Name: {{ name }}</p>
		<p>Title: {{ user.title }}</p>
		<p>Status: {{ user.status }}</p>
		<button @click="editUser">Edit</button>
		<button @click="onRemove">Remove</button>
	</section>
</template>

<script>
	export default {
		props: {
			user: {
				type: Object,
				required: true
			}
		},
		computed: {
			name() {
				return this.user.m_name 
					? this.user.f_name + ' ' + this.user.m_name + ' ' + this.user.l_name 
					: this.user.f_name + ' ' + this.user.l_name;
			}
		},
		methods: {
			editUser() {
				this.$router.push('/admin/users/' + this.user.id);
			},
			onRemove() {
				this.$emit('remove', this.user.id);
			}
		}
	}
</script>

<style scoped>
	.user {
		background-color: var(--color-grey-medium);
		padding: 10px;
		cursor: pointer;
	}
</style>