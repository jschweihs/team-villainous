<template>
    <div class="upcoming-event">
        <img
            :src="'/images/events/' + event.id + '.png'"
            width="992"
            onerror="this.onerror=null;this.src='/images/events/placeholder.jpg';"
        />
        <div class="upcoming-event-details">
            <h2 class="mb1">{{ event.name }}</h2>
            <p>{{ event.location }}</p>

            <p v-if="event.show_start_time">Begins: {{event.start_datetime | isoToDateTime }}</p>
            <p v-else>Begins: {{event.start_datetime | date }}</p>
            <p v-if="event.show_end_time" class="mb1">Ends: {{event.end_datetime | isoToDateTime }}</p>
            <p v-else class="mb1">Ends: {{event.end_datetime | date }}</p>

            <p v-if="event.referral_url" class="mb1">
                <a :href="event.referral_url">{{ event.referral_url }}</a>
            </p>
            <p class="mb1" v-if="event.description" v-html="event.description"></p>
            <p v-if="event.users && event.users.length > 0" v-html="usersContent"></p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        event: {
            type: Object,
            required: true
        },
        users: {
            type: Array,
            required: false
        }
    },
    computed: {
        usersContent() {
            if (
                this.users &&
                this.users.length > 0 &&
                this.event.users &&
                this.event.users.length > 0
            ) {
                // Initial content
                let content = "Villainous members attending: ";

                // Collect set of usernames
                this.event.users.forEach(user_id => {
                    const user = this.users.find(u => u.id == user_id);
                    content += user.username + ", ";
                });

                // Remove trailing commea
                content = content.substring(0, content.length - 2);

                return content;
            }
        }
    }
};
</script>

<style scoped>
.upcoming-event {
    background-color: #333;
    margin-bottom: 20px;
}

.upcoming-event h2 {
    color: white;
    margin-top: 0;
}

.upcoming-event img {
    display: block;
    max-width: 100%;
    border-top-right-radius: 20px;
    border-top-left-radius: 20px;
}

.upcoming-event p:last-child {
    margin-bottom: 0;
}

.upcoming-event-details {
    border: 1px solid var(--color-grey-light);
    border-top: 0;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    padding: 20px;
    background-color: var(--color-grey-medium);
}
</style>