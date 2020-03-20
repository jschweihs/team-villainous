<template>
    <div class="content">
        <div class="roles">
            <h1>Roles</h1>
            <roles-list :roles="roles" @add="addRole" @save="updateRole" @remove="removeRole" />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import RolesList from "@components/roles/RolesList.vue";

export default {
    components: {
        RolesList
    },
    computed: {
        ...mapGetters(["roles"])
    },
    methods: {
        addRole(name) {
            this.$store.dispatch("addRole", name);
        },
        updateRole(role) {
            this.$store.dispatch("updateRole", role);
        },
        removeRole(id) {
            this.$store.dispatch("removeRole", id);
        }
    },
    created() {
        this.$store.dispatch("showModal", true);
        this.$store
            .dispatch("getRoles")
            .then(res => {
                this.$store.dispatch("showModal", false);
            })
            .catch(e => console.log(e));
    }
};
</script>

<style scoped>
h1 {
    margin: 0;
    border-bottom: 1px solid var(--color-grey-light);
}

.roles {
    background-color: var(--color-grey-dark);
    padding: 10px;
    padding-top: 20px;
    border-radius: 20px;
}
</style>
