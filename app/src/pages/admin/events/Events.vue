<template>
    <div class="content">
        <h1>Manage Events</h1>
        <router-link to="/admin/events/new" tag="button">Add New Event</router-link>
        <div v-if="events">
            <event
                v-for="(event, index) in activeEvents"
                :key="index"
                :event="event"
                @edit="editEvent"
                @remove="removeEvent"
                admin
            />

            <section>
                <p class="text-action text-right" @click="toggleShowInactiveEvents">
                    <span v-if="showInactiveEvents">Hide</span>
                    <span v-else>Show</span>
                    inactive events
                </p>
                <div class="mt1" v-if="showInactiveEvents">
                    <event
                        v-for="(event, index) in inactiveEvents"
                        :key="index"
                        :event="event"
                        @edit="editEvent"
                        @remove="removeEvent"
                        admin
                    />
                </div>
            </section>
        </div>
        <div v-else>No events found</div>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

import Event from "@components/events/Event.vue";

export default {
    components: {
        Event
    },
    data() {
        return {
            showInactiveEvents: false
        };
    },
    computed: {
        ...mapGetters(["events"]),
        activeEvents() {
            return this.events.filter(event => {
                return event.status == 1;
            });
        },
        inactiveEvents() {
            return this.events.filter(event => {
                return event.status == 2;
            });
        }
    },
    methods: {
        editEvent(id) {
            this.$router.push("/admin/events/" + id);
        },
        removeEvent(id) {
            this.$store.dispatch("showModal", true);
            this.$store.dispatch("removeEvent", id).then(res => {
                this.$store.dispatch("showModal", false);
            });
        },
        toggleShowInactiveEvents() {
            this.showInactiveEvents = !this.showInactiveEvents;
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

.text-action {
    padding-left: 8px;
    color: var(--color-secondary);
    cursor: pointer;
    transition: all 0.25s;
    font-size: 14px;
}

.text-action:hover {
    color: var(--color-secondary-light);
}
</style>
