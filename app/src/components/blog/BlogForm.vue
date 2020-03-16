<template>
	<div class="blog-form">
		<form @submit.prevent="save">
			<label>Title*</label>
			<input type="text" v-model="entry.title"/>
			<label class="img-label">Promo Picture</label>
			<img 
				:src="entry.promo_picture" 
				onclick="document.getElementById('promo-picture').click()"
			/>
			<label for="promo-picture" class="file-label">Select</label>
			<input 
				type="file" 
				id="promo-picture" 
				accept="image/jpeg"
				ref="promopicture"
				v-model="image_path"
				@change="uploadImage"
			/>
			<label>Content</label>
			<textarea class="blog-content" v-model="entry.content"></textarea>
			<label>Preview</label>
			<textarea class="preview" v-model="entry.preview"></textarea>
			<input type="submit" value="Save Blog Entry"/>
		</form>
	</div>
</template>

<script>
	export default {
		props: {
			editentry: {
				type: Object,
				required: false
			}
		},
		computed: {
			current_user() {
				return this.$store.getters.current_user;
			},
			image_name() {
				return this.entry.title.toLowerCase().replace(' ', '_');
			}
		},
		data() {
			return {
				image_path: '',
				image_contents: '',
				entry: this.editentry
					? { 
						...this.editentry, 
						promo_picture: '/images/blog/' + this.editentry.id + '.jpg'
					} 
					: {
						title: 				'',
						user_id: 			this.current_user ? this.current_user.id : "1",
						username: 			this.current_user ? this.current_user.username : "Anon",
						preview: 			'',
						content: 			'',
						promo_picture: 		'/images/blog/placeholder.jpg',
					}
			};
		},
		methods: {
			uploadImage(e) {
				this.image_contents = this.$refs.promopicture.files[0];
				this.entry.promo_picture = URL.createObjectURL(e.target.files[0]);
			},
			save() {
				// Prep file
				let image = new FormData();
				image.append('image', this.image_contents);
				image.append('folder', 'blog');
				image.append('name', this.entry.id);

				this.$emit('save', { entry: this.entry, image, image_path: this.image_path });
			}
		}
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

	.blog-form {
		width: 627px;
		margin: 0 auto;
	}

	hr {
		margin: 20px 0;
		border: 1px solid var(--color-grey-light);
	}

	img {
		width: 627px;
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

	.preview {
		height: 110px;
		resize: none;
	}

	.blog-content {
		height: 370px;
		resize: none;
	}

	.promot-picture {
		margin-bottom: 20px;
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
	    text-align: center;
	    box-sizing: border-box;
	    font-size: 20px;
	    border-bottom-left-radius: 20px;
	    border-bottom-right-radius: 20px;
		margin-bottom: 10px;
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