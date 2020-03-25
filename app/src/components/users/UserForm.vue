<template>
    <div>
        <form @submit.prevent="save" enctype="multipart/form-data" id="user-form">
            <p v-if="showError" class="error text-center">{{ error }}</p>
            <fieldset>
                <legend>User Information</legend>

                <div class="flex">
                    <div class="profile-pic">
                        <div class="picture-preview">
                            <label for="promo-display" class="img-label">Profile Picture</label>
                            <img
                                id="promo-display"
                                width="260"
                                height="260"
                                :src="user.profile_picture"
                                onclick="document.getElementById('profile-picture').click()"
                            />

                            <label for="profile-picture" class="file-label">Select</label>
                            <input
                                type="file"
                                id="profile-picture"
                                accept="image/jpeg"
                                ref="profilepicture"
                                v-model="image_path"
                                @change="uploadImage"
                            />
                        </div>
                    </div>

                    <div class="main-info">
                        <label for="username">
                            Username
                            <span aria-hidden="true">*</span>
                        </label>
                        <input
                            type="text"
                            v-model="user.username"
                            id="username"
                            :class="{'form-error': !validUsername && showError}"
                            @keyup="checkForErrors"
                            aria-describedby="error"
                            placeholder="Vegeta"
                            requried
                        />

                        <label for="email">
                            Email
                            <span aria-hidden="true">*</span>
                        </label>
                        <input
                            type="text"
                            v-model="user.email"
                            id="email"
                            :class="{'form-error': !validEmail && showError}"
                            @keyup="checkForErrors"
                            placeholder="vegeta@planetvegeta.com"
                        />

                        <div v-if="!edituser">
                            <label for="password">
                                Password
                                <span aria-hidden="true">*</span>
                            </label>
                            <input
                                type="password"
                                v-model="user.password"
                                id="password"
                                :class="{'form-error': (!validPassword || !validPasswordMatch) && showError}"
                                @keyup="checkForErrors"
                            />
                            <label for="password-match">
                                Password Confirm
                                <span aria-hidden="true">*</span>
                            </label>
                            <input
                                type="password"
                                v-model="passwordMatch"
                                id="password-match"
                                :class="{'form-error': (!validPassword || !validPasswordMatch) && showError}"
                                @keyup="checkForErrors"
                            />
                        </div>

                        <label for="role">
                            Role
                            <span aria-hidden="true">*</span>
                        </label>
                        <select
                            v-model="user.role"
                            id="role"
                            :class="{'form-error': !validRole && showError}"
                            @keyup="checkForErrors"
                        >
                            <option value="0" disabled>Select role...</option>
                            <option
                                v-for="role in roles"
                                :key="role.id"
                                :value="role.id"
                            >{{ role.name }}</option>
                        </select>

                        <label for="title">Title</label>
                        <input
                            type="text"
                            v-model="user.title"
                            id="title"
                            placeholder="Price of All Saiyans"
                        />
                    </div>
                </div>
            </fieldset>

            <fieldset>
                <legend>Personal Information</legend>

                <div class="flex">
                    <div class="two-column mr-10">
                        <label for="f-name">
                            First Name
                            <span aria-hidden="true">*</span>
                        </label>
                        <input
                            type="text"
                            v-model="user.f_name"
                            id="f-name"
                            :class="{'form-error': !validFName && showError}"
                            @keyup="checkForErrors"
                            placeholder="Vegeta"
                        />
                        <label for="m-name">Middle Name</label>
                        <input type="text" v-model="user.m_name" id="m-name" />
                        <label for="l-name">Last Name</label>
                        <input type="text" v-model="user.l_name" id="l-name" />
                        <label for="description">Description</label>
                        <textarea v-model="user.description" id="description"></textarea>
                    </div>
                    <div class="two-column ml-10">
                        <label for="birth-date">Date of Birth</label>
                        <input
                            type="text"
                            :class="{'form-error': !validDOB && showError}"
                            id="birth-date"
                            placeholder="mm/dd/yyyy"
                            v-model="user.birth_date"
                            v-mask="'##/##/####'"
                            @keyup="checkForErrors"
                        />
                        <label for="address">Address</label>
                        <input type="text" v-model="user.address" id="address" />
                        <label for="city">City</label>
                        <input type="text" v-model="user.city" id="city" />
                        <div v-if="user.country == 'United States'">
                            <label for="state">State</label>
                            <select v-model="user.province" id="state">
                                <option value disabled>Select state/province...</option>
                                <option
                                    v-for="state in states"
                                    :key="state"
                                    :value="state"
                                >{{ state }}</option>
                            </select>
                        </div>
                        <div v-else>
                            <label for="province">Province</label>
                            <input type="text" v-model="user.province" id="province" />
                        </div>

                        <label for="zip">Zip code</label>
                        <input type="text" v-model="user.zip" id="zip" />
                        <label for="country">Country</label>
                        <select v-model="user.country" @change="user.province = ''" id="country">
                            <option value disabled>Select country...</option>
                            <option
                                v-for="country in countries"
                                :key="country"
                                :value="country"
                            >{{ country }}</option>
                        </select>
                    </div>
                </div>
            </fieldset>

            <fieldset>
                <legend>Social Media</legend>
                <label for="facebook">Facebook</label>
                <input
                    type="text"
                    v-model="user.facebook_url"
                    id="facebook"
                    placeholder="Facebook handle"
                />
                <label for="twitter">Twitter</label>
                <input
                    type="text"
                    v-model="user.twitter_url"
                    id="twitter"
                    placeholder="Twitter handle"
                />
                <label for="instagram">Instagram</label>
                <input
                    type="text"
                    v-model="user.instagram_url"
                    id="instagram"
                    placeholder="Instagram handle"
                />
                <label for="twitch">Twitch</label>
                <input
                    type="text"
                    v-model="user.twitch_url"
                    id="twitch"
                    placeholder="Twitch handle"
                />
                <label for="youtube">Youtube</label>
                <input
                    type="text"
                    v-model="user.youtube_url"
                    id="youtube"
                    placeholder="http://youtube.com/myChannel"
                />
                <label for="other">Other</label>
                <input
                    type="text"
                    v-model="user.other_url"
                    id="other"
                    placeholder="http://my-website.com"
                />
            </fieldset>

            <fieldset>
                <legend>Gamertags</legend>
                <label for="ps4-gamertag">PS4 Gamertag</label>
                <input type="text" v-model="user.ps4_gamertag" id="ps4-gamertag" />
                <label for="xbox-gamertag">XBox Gamertag</label>
                <input type="text" v-model="user.xbox_gamertag" id="xbox-gamertag" />
                <label for="steam-gamertag">Steam Gamertag</label>
                <input type="text" v-model="user.steam_gamertag" id="steam-gamertag" />
            </fieldset>

            <fieldset>
                <div class="submit-wrapper">
                    <input type="submit" value="Save User" :disabled="!validForm" />
                    <div class="submit-overlay" @mouseover="showError = true" @click="scrollToTop"></div>
                </div>
            </fieldset>

            <p v-if="showError" class="error text-center">{{ error }}</p>
        </form>
    </div>
</template>

<script>
import { states, countries } from "@data/geo";

import { mask } from "vue-the-mask";

import Validator from "@utils/Validator";

export default {
    props: {
        edituser: {
            type: Object,
            required: false
        },
        roles: {
            type: Array,
            required: true
        }
    },
    data() {
        return {
            image_path: "",
            image_contents: "",
            states,
            countries,
            user: this.edituser
                ? {
                      ...this.edituser,
                      password: "",
                      profile_picture:
                          "/images/users/ViL_" +
                          this.edituser.username.split(" ").join("") +
                          ".jpg"
                  }
                : {
                      username: "",
                      email: "",
                      password: "",
                      profile_picture: "/images/users/placeholder.jpg",
                      f_name: "",
                      m_name: "",
                      l_name: "",
                      title: "",
                      address: "",
                      city: "",
                      province: "",
                      zip: "",
                      country: "United States",
                      birth_date: "",
                      description: "",
                      role: 0,
                      privilege_id: 1,
                      status: 1,
                      facebook_url: "",
                      twitter_url: "",
                      instagram_url: "",
                      twitch_url: "",
                      youtube_url: "",
                      other_url: "",
                      ps4_gamertag: "",
                      xbox_gamertag: "",
                      steam_gamertag: ""
                  },
            passwordMatch: "",
            // UI control
            showError: false,
            error: "Username is missing"
        };
    },
    computed: {
        dob() {
            if (this.user.birth_date) {
                const dates = this.user.birth_date.split("/");
                return dates[2] + "-" + dates[0] + "-" + dates[1];
            }
        },
        //
        // Validity checks
        //
        // validUsername checks if a username is valid
        validUsername() {
            return this.user.username != "";
        },
        // validUsername checks if an email is valid
        validEmail() {
            return Validator.validEmail(this.user.email);
        },
        // validUsername checks if a password is valid
        validPassword() {
            if (!this.edituser) {
                return Validator.validPassword(this.user.password, 3);
            }
            return true;
        },
        // validUsername checks if both passwords match
        validPasswordMatch() {
            if (!this.edituser) {
                return this.user.password == this.passwordMatch;
            }
            return true;
        },
        // validUsername checks if a role is valid
        validRole() {
            return this.user.role > 0;
        },
        // validUsername checks if a first name is valid
        validFName() {
            return this.user.f_name != "";
        },
        validDOB() {
            if (this.user.birth_date != "") {
                console.log("birth date", this.user.birth_date);
                return Validator.validDate(this.user.birth_date);
            }
            return true;
        },
        // validUsername checks if the form is valid
        validForm() {
            return (
                this.validUsername &&
                this.validEmail &&
                this.validPassword &&
                this.validPasswordMatch &&
                this.validRole &&
                this.validFName &&
                this.validDOB
            );
        }
    },
    methods: {
        // Checks keystokes and updates form status
        checkForErrors() {
            if (!this.validUsername) {
                this.error = "Username cannot be empty";
            } else if (!this.validEmail) {
                this.error = "Email must be valid email address";
            } else if (!this.validPassword) {
                this.error =
                    "Password must be characters long and contain 1 lowercase letter, 1 uppercase letter, 1 number, and 1 special character";
            } else if (!this.validPasswordMatch) {
                this.error = "Passwords do not match";
            } else if (!this.validRole) {
                this.error = "Role cannot be empty";
            } else if (!this.validFName) {
                this.error = "First name cannot be empty";
            } else if (!this.validDOB) {
                this.error = "Date of birth is not formatted correctly";
            } else {
                this.error = "";
            }
        },
        // Prepares image to be uploaded and displays chosen image
        uploadImage(e) {
            // Get image
            this.image_contents = this.$refs.profilepicture.files[0];

            // Set image source url to new image for preview
            this.user.profile_picture = URL.createObjectURL(e.target.files[0]);
        },
        save() {
            // Create a new user with proper date format for dob
            const user = { ...this.user, birth_date: this.dob };

            // Create payload
            let payload = {
                user
            };

            // Handle image
            // Create new form data with image information
            let image = new FormData();
            image.append("image", this.image_contents);
            image.append("folder", "users");
            image.append("name", "ViL_" + user.username);

            // Append image to payload
            payload.image = image;
            payload.image_path = this.image_path;

            // Save user
            this.$emit("save", payload);
        },
        // Moves the view to the beginning of the form
        scrollToTop() {
            document.getElementById("user-form").scrollIntoView({
                behavior: "smooth"
            });
        }
    },
    created() {
        // Format date for input field
        if (this.user.birth_date) {
            const date_array = this.user.birth_date.split("-");
            this.user.birth_date =
                date_array[1] + "/" + date_array[2] + "/" + date_array[0];
        }
    },
    directives: { mask }
};
</script>

<style scoped>
.p1 {
    padding: 1rem;
}

.flex {
    display: flex;
}
.two-column {
    flex: 50%;
}

.ml-10 {
    margin-left: 10px;
}

.mr-10 {
    margin-right: 10px;
}

hr {
    margin: 20px 0;
    border: 1px solid var(--color-grey-light);
}

fieldset {
    padding: 0.5rem 0;
    border: 0;
    border-top: 1px solid var(--color-grey-light);
}

legend {
    color: white;
    font-family: inherit;
    padding: 0 0.5rem;
    text-align: center;
}

img {
    width: 260px;
    height: 260px;
    display: block;
    margin: 0 auto;
    border-top-left-radius: 20px;
    border-top-right-radius: 20px;
    cursor: pointer;
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
    height: 215px;
    resize: none;
}

input:focus,
select:focus,
textarea:focus {
    border: 2px solid var(--color-secondary);
    outline: none;
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

.profile-pic {
    width: 340px;
}

.img-label {
    margin-left: 40px;
}

.main-info {
    width: calc(100% - 340px);
}

.file-label {
    display: block;
    cursor: pointer;
    color: white;
    margin: 0 auto;
    background-color: var(--color-secondary);
    border: 0;
    height: 35px;
    cursor: pointer;
    font-family: "Geizer", cursive;
    letter-spacing: 2px;
    line-height: 35px;
    width: 260px;
    text-align: center;
    box-sizing: border-box;
    font-size: 20px;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    transition: all 0.25s;
}

.file-label:hover {
    background-color: var(--color-secondary-light);
}

input[type="file"] {
    opacity: 0;
    position: absolute;
    z-index: -1;
}
</style>
