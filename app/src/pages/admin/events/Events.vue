<template>
    <div class="content">
        <h1>Manage Events</h1>
        <router-link to="/admin/events/new" tag="button">Add New Event</router-link>
        <div v-if="events">
            <event
                v-for="(event, index) in events"
                :key="index"
                :event="event"
                @edit="editEvent"
                @remove="removeEvent"
                admin
            />
        </div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import Event from "@components/events/Event.vue";

export default {
    components: {
        Event
    },
    computed: {
        ...mapGetters(["events"])
    },
    methods: {
        editEvent(event_id) {
            this.$router.push("/admin/events/" + entry_id);
        },
        removeEvent(event_id) {
            // this.$store.dispatch('removeEvent', user_id);
        }
    },
    created() {
        this.$store.dispatch("showModal", true);
        this.$store
            .dispatch("getEvents")
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
</style>
