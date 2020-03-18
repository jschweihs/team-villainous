<template>
    <div class="content">
        <h1>Add New User</h1>
        <div class="new-user-form">
            <user-form :roles="roles" @save="addUser" />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import UserForm from "@components/users/UserForm.vue";

export default {
    components: {
        UserForm
    },
    computed: {
        ...mapGetters(["roles"])
    },
    methods: {
        // Add a new user
        addUser(user) {
            return this.$store
                .dispatch("addUser", user)
                .then(res => {
                    if (
                        res.data &&
                        res.data.data &&
                        res.data.data.username &&
                        res.data.data.username != ""
                    ) {
                        // Adding user was successful so take them back to
                        // the users page
                        this.$router.push("/admin/users");
                    } else {
                        // There was an unknown problem with adding
                        // a new user
                        alert("There was a problem saving this user");
                    }
                })
                .catch(e => console.log(e));
        }
    },
    created() {
        // Gather roles upon loading
        this.$store.dispatch("showModal", true);
        this.$store
            .dispatch("getRoles")
            .then(res => {
                this.$store.dispatch("showModal", false);
            })
            .catch(e => {
                console.log(e);
                this.$store.dispatch("showModal", false);
            });
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
