<template>
    <div class="content">
        <h1>Manage Users</h1>
        <router-link to="/admin/users/new" tag="button">Add New User</router-link>
        <div v-if="groupedUsers && roles">
            <team-section
                v-for="(group, roleName) in groupedUsers"
                :key="roleName"
                :group="group"
                :roleName="roleName"
                @remove="removeUser"
                isAdmin
            />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import TeamSection from "@components/users/TeamSection.vue";

export default {
    components: {
        TeamSection
    },
    computed: {
        ...mapGetters(["current_user", "roles", "groupedUsers"])
    },
    methods: {
        removeUser(user_id) {
            this.$store.dispatch("removeUser", user_id);
        }
    },
    created() {
        this.$store.dispatch("showModal", true);
        Promise.all([
            this.$store.dispatch("getUsers", { status: 1 }),
            this.$store.dispatch("getRoles")
        ])
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

button {
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
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
    border: 0;
    cursor: pointer;
    transform: translateY(0);
    transition: all 0.25s;
}

button:hover {
    background-color: var(--color-secondary-light);
    transform: translateY(-2px);
}
</style>
