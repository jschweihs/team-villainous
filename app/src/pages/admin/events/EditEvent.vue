<template>
    <div class="content">
        <h1>Edit Event</h1>
        <div class="edit-event-form">
            <event-form v-if="event" :editEvent="event" :users="users" @save="updateEvent" />
        </div>
    </div>
</template>

<script>
import axios from "axios";
import { mapGetters } from "vuex";

import EventForm from "@components/events/EventForm.vue";

export default {
    components: {
        EventForm
    },
    data() {
        return {
            event: null
        };
    },
    computed: {
        ...mapGetters(["users"]),
        id() {
            return this.$route.params.id;
        }
    },
    methods: {
        // Update an event
        updateEvent(payload) {
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

            this.$store
                .dispatch("updateEvent", payload)
                .then(res => {
                    if (res.data && res.data.data && res.data.data.name != "") {
                        // Updating event was successful so take them back to
                        // the events page
                        this.$router.push("/admin/events");
                    } else {
                        // There was an unknown problem with updating
                        // the event
                        alert("There was a problem saving this event");
                    }
                })
                .catch(e => console.log(e));
        }
    },
    created() {
        // Verify we have a valid id
        if (parseInt(this.id) != this.id) {
            this.$router.push("/admin/events");
        } else {
            // Gather event information and user list
            this.$store.dispatch("showModal", true);
            Promise.all([
                this.$store.dispatch("getEvents", this.id),
                this.$store.dispatch("getUsers")
            ])
                .then(res => {
                    this.event = this.$store.getters.event(this.id);
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
