<template>
    <div class="content">
        <!-- Upcoming events -->
        <h1>Upcoming Events</h1>
        <upcoming-events
            v-if="upcomingEvents && upcomingEvents.length > 0"
            :events="upcomingEvents"
            :users="users"
        ></upcoming-events>
        <p v-else>No upcoming events. Please stay tuned!</p>

        <!-- Past events -->
        <h1 v-if="pastEvents && pastEvents.length > 0">Past Events</h1>
        <past-events v-if="pastEvents && pastEvents.length > 0" :events="pastEvents"></past-events>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import UpcomingEvents from "@components/events/UpcomingEvents.vue";
import PastEvents from "@components/events/PastEvents.vue";

export default {
    components: {
        UpcomingEvents,
        PastEvents
    },
    computed: {
        ...mapGetters(["upcomingEvents", "pastEvents", "users"])
    },
    created() {
        // Fetch event data
        this.$store.dispatch("showModal", true);
        Promise.all([
            this.$store.dispatch("getEvents"),
            this.$store.dispatch("getUsers")
        ]).then(res => {
            this.$store.dispatch("showModal", false);
        });
    }
};
</script>

<style scoped>
</style>
