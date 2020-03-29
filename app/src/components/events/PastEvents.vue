<template>
    <ul class="past-events">
        <past-event
            v-for="(event, index) in events"
            :key="event.id"
            :event="event"
            :alt="index%2==0"
        ></past-event>
    </ul>
</template>

<script>
import PastEvent from "@components/events/PastEvent.vue";

export default {
    components: {
        PastEvent
    },
    props: {
        events: {
            type: Array,
            requried: true
        }
    }
};
</script>

<style scoped>
.past-events {
    background-color: var(--color-grey-dark);
    list-style: none;
    margin: 0;
    margin-bottom: 20px;
    padding: 0;
    font-size: 0;
}
</style>