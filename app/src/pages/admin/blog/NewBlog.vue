<template>
    <div class="content" v-if="currentUser">
        <h1>Add New Blog Entry</h1>
        <div class="blog-form">
            <blog-form @save="addEntry" :userID="currentUser.id" />
        </div>
    </div>
</template>

<script>
import BlogForm from "@components/blog/BlogForm.vue";

export default {
    components: {
        BlogForm
    },
    computed: {
        currentUser() {
            return this.$store.getters.currentUser;
        }
    },
    methods: {
        addEntry(entry) {
            this.$store.dispatch("showModal", true);
            return this.$store
                .dispatch("addEntry", entry)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                    if (
                        res.data &&
                        res.data.data &&
                        res.data.data.title != ""
                    ) {
                        // Adding blog entry was successful so take them back to
                        // the blog page
                        this.$router.push("/admin/blog");
                    } else {
                        // There was an unknown problem with adding
                        // a new blog entry
                        alert("There was a problem saving this blog entry");
                    }
                })
                .catch(e => console.log(e));
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
