<template>
    <div class="content">
        <h1>Edit User</h1>
        <div class="new-user-form">
            <user-form v-if="user" :edituser="user" :roles="roles" @save="updateUser" />
        </div>
    </div>
</template>

<script>
import axios from "axios";
import { mapGetters } from "vuex";

import UserForm from "@components/users/UserForm.vue";

export default {
    components: {
        UserForm
    },
    data() {
        return {
            user: null
        };
    },
    computed: {
        ...mapGetters(["roles"]),
        id() {
            return this.$route.params.id;
        }
    },
    methods: {
        updateUser(user) {
            this.$store
                .dispatch("updateUser", user)
                .then(res => {
                    if (res.data.data && res.data.data.username != "") {
                        this.$router.push("/admin/users");
                    }
                })
                .catch(e => console.log(e));
        }
    },
    created() {
        // Verify we have a valid id
        if (parseInt(this.id) != this.id) {
            this.$router.push("/admin/users");
        } else {
            // Gather user information and role list
            this.$store.dispatch("showModal", true);
            Promise.all([
                this.$store.dispatch("getUser", this.id),
                this.$store.dispatch("getRoles")
            ])
                .then(res => {
                    this.user = this.$store.getters.user(this.id);
                    this.$store.dispatch("showModal", false);
                })
                .catch(e => {
                    console.log(e);
                    this.$store.dispatch("showModal", false);
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
