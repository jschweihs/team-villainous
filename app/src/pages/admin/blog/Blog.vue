<template>
    <div class="content">
        <h1>Blog</h1>
        <router-link tag="button" to="/admin/blog/new">New Entry</router-link>
        <div class="blog-list" v-if="blog">
            <blog-preview
                v-for="entry in blog"
                :user="user(entry.user_id)"
                :key="entry.id"
                :entry="entry"
                @get-user="getUser"
                @edit="editEntry"
                @remove="removeEntry"
            />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import BlogPreview from "@components/blog/BlogPreview.vue";

export default {
    components: {
        BlogPreview
    },
    computed: {
        ...mapGetters(["blog"])
    },
    methods: {
        getUser(id) {
            this.$store.dispatch("showModal", true);
            this.$store
                .dispatch("getUser", id)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                })
                .catch(e => {
                    this.$store.dispatch("showModal", true);
                    console.log(e);
                });
        },
        user(id) {
            return this.$store.getters.user(id);
        },
        editEntry(entry_id) {
            this.$router.push("/admin/blog/" + entry_id);
        },
        removeEntry(entry_id) {
            this.$store.dispatch("showModal", true);
            this.$store.dispatch("removeEntry", entry_id).then(res => {
                this.$store.dispatch("showModal", false);
            });
        }
    },
    created() {
        this.$store.dispatch("showModal", true);
        Promise.all([
            this.$store.dispatch("getBlog"),
            this.$store.dispatch("getUsers")
        ])
            .then(res => {
                this.$store.dispatch("showModal", false);
            })
            .catch(e => {
                this.$store.dispatch("showModal", true);
                console.log(e);
            });
    }
};
</script>

<style scoped>
h1 {
    border-bottom: 1px solid var(--color-grey-light);
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
    -moz-box-sizing: border-box; /* Firefox, other Gecko */
    box-sizing: border-box; /* Opera/IE 8+ */
    color: white;
    background-color: var(--color-secondary);
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    border: 0;
    cursor: pointer;
    transform: translateY(0);
    transition: all 0.25s;
}

button:hover {
    transform: translateY(-2px);
    background-color: var(--color-secondary-light);
}
.content {
    background-color: var(--color-grey-dark);
    border-radius: 20px;
    padding: 20px;
}

.blog-list {
    width: 627px;
    margin: 0 auto;
}
</style>
