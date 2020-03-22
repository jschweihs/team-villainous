<template>
    <div class="content">
        <h1>Update Blog Entry</h1>
        <div class="blog-form">
            <blog-form v-if="entry" :editEntry="entry" @save="updateEntry" />
        </div>
    </div>
</template>

<script>
import BlogForm from "@components/blog/BlogForm.vue";

export default {
    components: {
        BlogForm
    },
    data() {
        return {
            entry: null
        };
    },
    computed: {
        id() {
            return this.$route.params.id;
        }
    },
    methods: {
        updateEntry(entry) {
            this.$store
                .dispatch("updateEntry", entry)
                .then(res => {
                    if (
                        res.data &&
                        res.data.data &&
                        res.data.data.title != ""
                    ) {
                        // Updating blog entry was successful so take them back to
                        // the blog page
                        this.$router.push("/admin/blog");
                    } else {
                        // There was an unknown problem with updating
                        // the blog entry
                        alert("There was a problem saving this blog entry");
                    }
                })
                .catch(e => console.log(e));
        }
    },
    created() {
        // Verify we have a valid id
        if (parseInt(this.id) != this.id) {
            this.$router.push("/admin/blog");
        } else {
            // Gather blog entry information
            this.$store.dispatch("showModal", true);
            this.$store
                .dispatch("getEntry", this.id)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                    this.entry = this.$store.getters.entry(this.id);
                })
                .catch(e => {
                    this.$store.dispatch("showModal", false);
                    console.log(e);
                });
        }
    }
};
</script>

<style scoped>
.content {
    background-color: var(--color-grey-dark);
    border-radius: 20px;
    padding: 20px;
}

h1 {
    border-bottom: 1px solid var(--color-grey-light);
    margin: 0;
    margin-bottom: 20px;
}
</style>
