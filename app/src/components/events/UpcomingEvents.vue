<template>
    <div>
        <upcoming-event v-for="event in events" :key="event.id" :event="event" :users="users"></upcoming-event>
    </div>
</template>

<script>
import UpcomingEvent from "./UpcomingEvent.vue";

export default {
    components: {
        UpcomingEvent
    },
    props: {
        events: {
            type: Array,
            required: true
        },
        users: {
            type: Array,
            required: false
        }
    }
};
</script>