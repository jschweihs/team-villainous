<template>
    <div class="content">
        <h1>Add New Event</h1>
        <div class="event-form">
            <event-form @save="addEvent" :users="users" />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import EventForm from "@components/events/EventForm.vue";

export default {
    components: {
        EventForm
    },
    computed: {
        ...mapGetters(["users"])
    },
    methods: {
        addEvent(payload) {
            // Update date params to fit RC3339
            if (payload.event.start_datetime != "") {
                payload.event.start_datetime += ":00Z";
            }
            if (payload.event.end_datetime != "") {
                payload.event.end_datetime += ":00Z";
            }

            // Convert data to correct type
            payload.event.type = parseInt(payload.event.type);
            payload.event.game_id = parseInt(payload.event.game_id);
            payload.event.status = parseInt(payload.event.status);

            for (let i = 0; i < payload.event.placements.length; i++) {
                payload.event.placements[i] = parseInt(
                    payload.event.placements[i]
                );
            }

            this.$store.dispatch("showModal", true);
            return this.$store
                .dispatch("addEvent", payload)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                    this.$router.push("/admin/events");
                })
                .catch(e => {
                    this.$store.dispatch("showModal", false);
                    console.log(e);
                });
        }
    },
    created() {
        this.$store.dispatch("showModal", true);

        this.$store
            .dispatch("getUsers")
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
