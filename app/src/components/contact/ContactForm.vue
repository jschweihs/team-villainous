<template>
    <form method="post" @submit.prevent="$emit('send', contact)">
        <label for="name">Name</label>
        <input
            type="text"
            v-model="contact.name"
            id="name"
            :class="{'form-error': !validName && showError}"
            placeholder="Kratos"
            @keyup="checkForErrors"
            required
        />

        <label for="email">Email</label>
        <input
            type="text"
            v-model="contact.email"
            id="email"
            :class="{'form-error': !validEmail && showError}"
            placeholder="godofwar@sparta.com"
            @keyup="checkForErrors"
            required
        />

        <label for="category">Category</label>
        <select
            v-model="contact.category"
            id="category"
            :class="{'form-error': !validCategory && showError}"
            @change="checkForErrors"
            required
        >
            <option value disabled selected>Select option...</option>
            <option value="Team Request">Team Request</option>
            <option value="Tournament Entry">Tournament Entry</option>
            <option value="Other">Other</option>
        </select>

        <label for="message">Message</label>
        <textarea
            id="message"
            :class="{'form-error': !validMessage && showError}"
            v-model="contact.message"
            placeholder="Enter message here..."
            @keyup="checkForErrors"
            requried
        ></textarea>

        <div class="submit-wrapper">
            <input type="submit" value="Send Message" :disabled="!validForm" />
            <div class="submit-overlay" @mouseover="showError = true; showSuccess=false"></div>
        </div>

        <p v-if="showError" class="error text-center mt1 mb1">{{ error }}</p>
        <p v-if="showSuccess" class="success mt1 mb1 text-center">Message sent successfully!</p>
    </form>
</template>

<script>
import Validator from "@utils/Validator";

export default {
    data() {
        return {
            // Contents of contact form
            contact: {
                name: "",
                email: "",
                category: "",
                message: ""
            },
            // Error handling
            error: "Name cannot be empty",
            showError: false,
            showSuccess: false
        };
    },
    computed: {
        // Determines if the name is valid
        validName() {
            return this.contact.name != "";
        },
        // Determines if the email is valid
        validEmail() {
            return Validator.validEmail(this.contact.email);
        },
        // Determines if the category is valid
        validCategory() {
            return this.contact.category != "";
        },
        // Determines if the message is valid
        validMessage() {
            return this.contact.message != "";
        },
        // Determines if the form is valid
        validForm() {
            return (
                this.validName &&
                this.validEmail &&
                this.validCategory &&
                this.validMessage
            );
        }
    },
    methods: {
        // Checks keystokes and updates form status
        checkForErrors() {
            if (!this.validName) {
                this.error = "Name cannot be empty";
            } else if (!this.validEmail) {
                this.error = "Email must be valid email address";
            } else if (!this.validCategory) {
                this.error = "Please select a category";
            } else if (!this.validMessage) {
                this.error = "Message cannot be empty";
            } else {
                this.error = "";
            }
        },
        // Clears out the contact form
        clearForm() {
            // Clear data
            this.contact.name = "";
            this.contact.email = "";
            this.contact.category = "";
            this.contact.message = "";
            // Clear error handling
            this.error = "Name cannot be empty";
            this.showError = false;
            this.showSuccess = true;
        }
    }
};
</script>

<style scoped>
form {
    width: 100%;
    margin: 0 auto;
}

label {
    color: white;
    display: block;
    font-size: 18px;
    margin-bottom: 4px;
    padding-left: 8px;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
}

input,
select,
textarea {
    display: block;
    width: 100%;
    font-size: 20px;
    font-family: inherit;
    padding: 12px 24px;
    margin-bottom: 10px;
    border-radius: 8px;
    border: 2px solid transparent;
    -webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    -moz-box-sizing: border-box; /* Firefox, other Gecko */
    box-sizing: border-box; /* Opera/IE 8+ */
    transition: all 0.25s;
}

input:focus,
select:focus,
textarea:focus {
    border: 2px solid var(--color-secondary);
    outline: none;
}

textarea {
    height: 200px;
}

input[type="submit"] {
    color: white;
    margin-top: 20px;
    background-color: var(--color-secondary);
    border: 0;
    height: 60px;
    cursor: pointer;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    border-radius: 8px;
    transform: translateY(0);
    transition: all 0.25s;
}

input[type="submit"]:hover {
    transform: translateY(-2px);
    background-color: var(--color-secondary-light);
}

input[type="submit"]:disabled {
    background-color: #bebebe;
    transform: translateY(0);
    pointer-events: auto;
}

select:required:invalid {
    color: grey;
}

select option:not(:disabled) {
    color: black;
}
</style>