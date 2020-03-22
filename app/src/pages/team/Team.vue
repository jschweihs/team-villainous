<template>
    <div class="content">
        <h1>Team</h1>
        <p class="page-description">
            If you are interested in joining Team Villainous then we would love to
            hear from you. Head over to the Contact page and send us your information.
            If we are interested then we will get back to you quickly!
        </p>
        <div v-if="groupedUsers && roles">
            <team-section
                v-for="(group, roleName) in groupedUsers"
                :key="roleName"
                :roleName="roleName"
                :group="group"
                @remove="removeUser"
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
        ...mapGetters(["groupedUsers", "roles"])
    },
    methods: {
        removeUser(user_id) {
            this.$store.dispatch("showModal", true);
            this.$store
                .dispatch("removeUser", user_id)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                })
                .catch(e => {
                    this.$store.dispatch("showModal", false);
                });
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
