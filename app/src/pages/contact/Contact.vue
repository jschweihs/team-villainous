<template>
    <div class="content half">
        <h1>Contact</h1>
        <contact-form @send="sendEmail" ref="contactForm"></contact-form>
        <p>Thank you for reaching out to Team Villainous. Please allow our team to respond to your inquiry within 24 hours. You can also contact us on any of our social media.</p>
    </div>
</template>

<script>
import axios from "axios";

import ContactForm from "@components/contact/ContactForm.vue";

export default {
    components: {
        ContactForm
    },
    data() {
        return {
            // Display success message
            showSuccess: false
        };
    },
    methods: {
        sendEmail(contact) {
            this.$store.dispatch("showModal", true);
            this.$store
                .dispatch("sendEmail", contact)
                .then(res => {
                    this.$store.dispatch("showModal", false);
                    if (res.data && res.data.data) {
                        this.$refs.contactForm.clearForm();
                    }
                })
                .catch(e => {
                    this.$store.dispatch("showModal", false);
                    console.log(e);
                });
        }
    }
};
</script>