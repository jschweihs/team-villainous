<template>
	<div class="content">
		<h1>Add New Blog Entry</h1>
		<div class="blog-form">
			<blog-form 
				@save="addEntry"
			/>
		</div>
	</div>
</template>

<script>

	import BlogForm from './../../../blog/BlogForm.vue';

	export default {
		components: {
			BlogForm
		},
		beforeCreate() {
			if(!this.$store.getters.jwt) {
				this.$router.push('/admin');
			}
		},
		methods: {
			addEntry(entry) {

				console.log('entry', entry);
				return this.$store.dispatch('addEntry', entry)
				.then(() => {
					this.$router.push('/admin/blog');
				})
				.catch(e => console.log(e));
			}
		},
	}

</script>

<style scoped>
	.content {
		background-color: #333;
		border-radius: 20px;
		padding: 20px;
	}

	h1 {
		border-bottom: 1px solid #444;
		margin: 0;
		margin-bottom: 20px;
	}
</style>