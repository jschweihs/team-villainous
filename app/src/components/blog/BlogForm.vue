<template>
    <div class="blog-form">
        <form @submit.prevent="save">
            <label for="title">Title*</label>
            <input type="text" id="title" v-model="entry.title" />
            <label class="img-label" for="promo-display">Promo Picture</label>
            <img
                id="promo-display"
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
            <label for="content">Content</label>
            <textarea id="content" class="blog-content" v-model="entry.content"></textarea>
            <label for="preview">Preview</label>
            <textarea id="preview" class="preview" v-model="entry.preview"></textarea>
            <input type="submit" value="Save Blog Entry" />
        </form>
    </div>
</template>

<script>
export default {
    props: {
        editEntry: {
            type: Object,
            required: false
        },
        userID: {
            type: Number,
            required: true
        }
    },
    data() {
        return {
            image_path: "",
            image_contents: "",
            entry: this.editEntry
                ? {
                      ...this.editEntry,
                      promo_picture:
                          "/images/blog/" +
                          this.editEntry.id.toString() +
                          ".jpg"
                  }
                : {
                      title: "",
                      user_id: this.userID,
                      preview: "",
                      content: "",
                      promo_picture: "/images/blog/placeholder.jpg"
                  }
        };
    },
    methods: {
        // Prepares image to be uploaded and displays chosen image
        uploadImage(e) {
            // Get image
            this.image_contents = this.$refs.promopicture.files[0];

            // Set image source url to new image for preview
            this.entry.promo_picture = URL.createObjectURL(e.target.files[0]);
        },
        save() {
            // Create payload
            let payload = {
                entry: this.entry
            };

            // Prep file
            let image = new FormData();
            image.append("image", this.image_contents);
            image.append("folder", "blog");
            image.append("name", this.entry.id);

            // Append image to payload
            payload.image = image;
            payload.image_path = this.image_path;

            // Save entry
            this.$emit("save", payload);
        }
    }
};
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
    color: white;
    display: block;
    font-size: 18px;
    margin-bottom: 4px;
    padding-left: 8px;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
}

input,
select,
textarea {
    display: block;
    width: 100%;
    font-size: 20px;
    font-family: inherit;
    padding: 12px 24px;
    margin-bottom: 10px;
    border-radius: 8px;
    border: 2px solid transparent;
    -webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    -moz-box-sizing: border-box; /* Firefox, other Gecko */
    box-sizing: border-box; /* Opera/IE 8+ */
    transition: all 0.25s;
}

textarea {
    height: 215px;
    resize: none;
}

input:focus,
select:focus,
textarea:focus {
    border: 2px solid var(--color-secondary);
    outline: none;
}

input[type="submit"] {
    color: white;
    margin-top: 20px;
    background-color: var(--color-secondary);
    border: 0;
    height: 60px;
    cursor: pointer;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    border-radius: 8px;
    transform: translateY(0);
    transition: all 0.25s;
}

input[type="submit"]:hover {
    transform: translateY(-2px);
    background-color: var(--color-secondary-light);
}

input[type="submit"]:disabled {
    background-color: #bebebe;
    transform: translateY(0);
    pointer-events: auto;
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

/* input[type="submit"] {
    color: white;
    margin-top: 20px;
    background-color: var(--color-secondary);
    border: 0;
    height: 60px;
    cursor: pointer;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    border-radius: 8px;
    transform: translateY(0);
    transition: all 0.25s;
}

input[type="submit"]:hover {
    transform: translateY(-2px);
    background-color: var(--color-secondary-light);
} */

.file-label {
    display: block;
    cursor: pointer;
    color: white;
    margin: 0 auto;
    background-color: var(--color-secondary);
    border: 0;
    height: 35px;
    cursor: pointer;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    line-height: 35px;
    text-align: center;
    box-sizing: border-box;
    font-size: 20px;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    margin-bottom: 10px;
    transition: all 0.25s;
}

.file-label:hover {
    background-color: var(--color-secondary-light);
}

input[type="file"] {
    opacity: 0;
    position: absolute;
    z-index: -1;
}
</style>