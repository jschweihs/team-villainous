<template>
    <div class="content half">
        <h1>Contact</h1>
        <form method="post" @submit.prevent="sendEmail">
            <label for="name">Name</label>
            <input type="text" v-model="email.name" id="name" placeholder="Kratos" required />

            <label for="email">Email</label>
            <input
                type="text"
                v-model="email.address"
                id="email"
                placeholder="godofwar@sparta.com"
                required
            />

            <label for="category">Category</label>
            <select v-modoel="email.category" required>
                <option value disabled selected>Select option..</option>
                <option>Team Request</option>
                <option>Tournament Entry</option>
            </select>
            <label for="message">Message</label>
            <textarea
                id="message"
                v-model="email.message"
                placeholder="Enter message here..."
                requried
            ></textarea>
            <input type="submit" />
        </form>
        <p>Thank you for reaching out to Team Villainous. Please allow our team to respond to your inquiry within 24 hours. You can also contact us on any of our social media.</p>
    </div>
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            email: {
                name: "",
                address: "",
                category: "",
                message: ""
            }
        };
    },
    methods: {
        sendEmail() {
            axios
                .post("http://teamvillainous.com/api/v1/file/send-email", {
                    ...this.email
                })
                .then(res => {
                    console.log("Email sent successfully");
                })
                .catch(e => console.log(e));
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
