<template>
    <div class="content login">
        <h1>Admin area</h1>

        <form @submit.prevent="login">
            <label for="email">Email</label>
            <input
                type="text"
                id="email"
                :class="{ error: showError && invalidEmail }"
                v-model="credentials.email"
                placeholder="johndoe@gmail.com"
            />

            <label for="password">Password</label>
            <input
                type="password"
                id="password"
                :class="{ error: showError && invalidPassword }"
                v-model="credentials.password"
                placeholder="p455w0rd"
            />

            <div class="submit-wrapper">
                <input type="submit" value="Login" :disabled="invalidForm" />
                <div class="submit-overlay" @mouseover="showError = true"></div>
            </div>

            <p class="text-error text-center" v-if="showError">{{ error }}</p>
        </form>
    </div>
</template>

<script>
import axios from "axios";

import Cookie from "@utils/Cookie";
import Validator from "@utils/Validator";

export default {
    data() {
        return {
            // Credentials used to login user
            credentials: {
                email: "",
                password: ""
            },
            // Error message to display
            error: "Email and/or password are not valid",
            // Defines if error is displayed
            showError: false
        };
    },
    computed: {
        // Returns if the email is valid
        invalidEmail() {
            return !Validator.validEmail(this.credentials.email);
        },
        // Returns if the password is valid
        invalidPassword() {
            return this.credentials.password == "";
        },
        // Returns if the form is valid
        invalidForm() {
            return this.invalidEmail || this.invalidPassword;
        }
    },
    watch: {
        // Watches form validity and updates error message accordingly
        invalidForm(v) {
            this.error = v ? this.error : "";
        }
    },
    methods: {
        // Attempt to login the user
        login() {
            // Display loader and hide any existing error messages
            this.showError = false;
            this.$store.dispatch("showModal", true);

            // Attempt to login the user
            this.$store
                .dispatch("login", this.credentials)
                .then(res => {
                    this.$store.dispatch("showModal", false);

                    if (
                        res.response &&
                        res.response.data &&
                        res.response.data.errors &&
                        res.response.data.errors[0] &&
                        res.response.data.errors[0].detail
                    ) {
                        // Handle error if returned
                        this.error = res.response.data.errors[0].detail;
                        this.showError = true;
                    } else if (
                        res.data &&
                        res.data.data &&
                        res.data.data.length > 0
                    ) {
                        // Successful login so redirect user
                        this.$router.push("/admin");
                    }
                })
                .catch(e => {
                    // Handle error if there was a problem getting response
                    this.$store.dispatch("showModal", false);
                    this.showError = true;
                    if (
                        e.response &&
                        e.response.data &&
                        e.response.data.error
                    ) {
                        this.error = e.response.data.error;
                        this.showError = true;
                    }
                });
        }
    }
};
</script>

<style scoped>
.login {
    width: 50%;
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

textarea {
    height: 200px;
}

input:focus,
select:focus,
textarea:focus {
    border: 2px solid var(--color-secondary);
    outline: none;
}

.text-center {
    text-align: center;
}

.text-error {
    color: var(--color-error);
}
.error {
    border: 2px solid var(--color-error);
}

input[type="submit"] {
    color: white;
    margin-top: 20px;
    background-color: var(--color-secondary);
    border: 0;
    height: 60px;
    cursor: pointer;
    transform: translateY(0);
    transition: all 0.25s;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
}

input[type="submit"]:hover {
    background-color: var(--color-secondary-light);
    transform: translateY(-2px);
}

input[type="submit"]:disabled {
    background-color: #bebebe;
    transform: translateY(0);
    pointer-events: auto;
}

@media only screen and (max-width: 600px) {
    .login {
        width: 100%;
    }
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
}
</style>
