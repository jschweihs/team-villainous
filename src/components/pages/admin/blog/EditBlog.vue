<template>
	<div class="content">
		<h1>Update Blog Entry</h1>
		<div class="blog-form">
			<blog-form
				v-if="entry" 
				:editentry="entry"
				@save="updateEntry"
			/>
		</div>
	</div>
</template>

<script>

	import BlogForm from './../../../blog/BlogForm.vue';

	import axios from 'axios';

	export default {
		components: {
			BlogForm
		},
		data() {
			return {
				entry: null
			};
		},
		// beforeCreate() {
		// 	if(!this.$store.getters.jwt) {
		// 		this.$router.push('/admin');
		// 	}
		// },
		computed: {

		},
		methods: {
			updateEntry(entry) {
				console.log('entry', entry);
				return this.$store.dispatch('updateEntry', entry)
				.then(() => {
					this.$router.push('/admin/blog');
				})
				.catch(e => console.log(e));
			}
		},
		created() {
			axios.get('http://teamvillainous.com/api/v1/blog/get.php?id=' + this.$route.params.id)
			.then(response => {
				this.entry = { ...response.data };
			})
			.catch(e => {
				console.log(e);
			})
		}
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