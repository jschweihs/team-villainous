<template>
	<div class="content">
		<h1>Blog</h1>
		<router-link tag="button" to="/admin/blog/new">New Entry</router-link>
		<div class="blog-list">
			<blog-preview
				v-for="(entry, index) in blog"
				:key="entry.id"
				:entry="entry"
				@edit="editEntry"
				@remove="removeEntry"
			/>
		</div>
	</div>
</template>

<script>

	import BlogPreview from './../../../blog/BlogPreview.vue';

	export default {
		components: {
			BlogPreview
		},
		computed: {
			blog() {
				return this.$store.getters.blog;
			}
		},
		methods: {
			editEntry(entry_id) {
				this.$router.push('/admin/blog/' + entry_id);
			},
			removeEntry(entry_id) {
				this.$store.dispatch('removeEntry', entry_id);
			}
		},
		created() {
			console.log('get blog');
			this.$store.dispatch("getBlog");
		}
	}
</script>

<style scoped>
	h1 {
		border-bottom: 1px solid #444;
		margin: 0;
		margin-bottom: 20px;
	}

	button {
		display: block;
		width: 100%;
		height: 60px;
		font-size: 24px;
		padding: 10px;
		margin: 20px 0;
		border-radius: 8px;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	color: white;
		background-color: #ffc200;
		border: 0;
		cursor: pointer;
	}
	.content {
		background-color: #333;
		border-radius: 20px;
		padding: 20px;
	}

	.blog-list {
		width: 627px;
		margin: 0 auto;
	}

</style>