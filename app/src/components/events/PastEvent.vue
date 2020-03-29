<template>
    <li class="past-event" :class="{'past-event-alt': alt}">
        <div class="past-event-detail past-event-name">{{ event.name }}</div>
        <div class="past-event-detail">{{ event.location }}</div>
        <div
            class="past-event-detail"
        >{{ event.start_datetime | shortDate}}-{{ event.end_datetime | shortDate}}</div>
        <div class="past-event-detail" v-html="placements"></div>
    </li>
</template>

<script>
export default {
    props: {
        event: {
            type: Object,
            required: true
        },
        alt: {
            type: Boolean,
            required: false
        }
    },
    computed: {
        placements() {
            if (this.event.placements && this.event.placements.length > 0) {
                // Sort placements
                const placements = this.event.placements.sort();

                // New content
                let content = "";

                // Add each placement
                placements.forEach(placement => {
                    // Start placement
                    content += "<span";

                    // Set class for 1st, 2nd, or 3rd place
                    switch (placement) {
                        case 1:
                            content += ' class="gold"';
                            break;
                        case 2:
                            content += ' class="silver"';
                            break;
                        case 3:
                            content += ' class="bronze"';
                            break;
                    }
                    content += ">";

                    // Set superscript
                    let superscript = "th";
                    switch (
                        placement
                            .toString()
                            .charAt(placement.toString().length - 1)
                    ) {
                        case "1":
                            superscript = "st";
                            break;
                        case "2":
                            superscript = "nd";
                            break;
                        case "3":
                            superscript = "rd";
                            break;
                    }

                    // 1st, 2nd, and 3rd place get a trophy icon
                    if (placement <= 3) {
                        content += '<i class="fas fa-trophy"></i> ';
                    }

                    // Add placement
                    content += placement + superscript;

                    content += "</span>";

                    // Add comma
                    content += ", ";
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
.past-event {
    color: white;
    padding: 10px;
}

.past-event-alt {
    background-color: var(--color-grey-medium);
}

.past-event-name {
    font-weight: bold !important;
    text-align: left !important;
}

.past-event-detail {
    font-size: 16px;
    display: inline-block;
    text-align: center;
    margin: 0;
    padding: 0;
    width: 25%;
}

@media screen and (max-width: 850px) {
    .past-event-name {
        display: block;
        width: 100%;
    }

    .past-event-detail {
        width: 33.33%;
    }
}

@media screen and (max-width: 628px) {
    .past-event-detail {
        display: block;
        width: 100%;
        margin: 10px 0;
    }
}
</style>
