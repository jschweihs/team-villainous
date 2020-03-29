<template>
    <div class="event-form">
        <form @submit.prevent="save" enctype="multipart/form-data" id="event-form">
            <p v-if="showError" class="error text-center">{{ error }}</p>

            <fieldset>
                <legend>Event Info</legend>
                <label for="name">Name *</label>
                <input
                    type="text"
                    id="name"
                    v-model="event.name"
                    :class="{'form-error': !validName && showError}"
                    @keyup="checkForErrors"
                />

                <label for="promo-display" class="img-label">Promo Picture</label>
                <img
                    id="promo-display"
                    :src="event.promo_picture"
                    onclick="document.getElementById('promo-picture').click()"
                />

                <label for="promo-picture" class="file-label">Select</label>
                <input
                    type="file"
                    id="promo-picture"
                    accept="image/jpeg"
                    ref="promopicture"
                    v-model="image_path"
                    @change="uploadImage"
                />

                <label for="location">Location</label>
                <input
                    type="text"
                    id="location"
                    v-model="event.location"
                    placeholder="Allstate Arena in Chicago, IL"
                />

                <label for="start-datetime">Starts At</label>
                <input
                    type="datetime-local"
                    id="start-datetime"
                    :class="{'form-error': (!validStartDatetime || !validTimeRange)&& showError}"
                    v-model="event.start_datetime"
                    @keyup="checkForErrors"
                />

                <input
                    type="checkbox"
                    id="show-start-time"
                    class="mb1"
                    v-model="event.show_start_time"
                />
                <label class="checkbox-label" for="show-start-time">Show time</label>

                <label for="end-datetime">Ends At</label>
                <input
                    type="datetime-local"
                    id="end-datetime"
                    :class="{'form-error': (!validEndDatetime || !validTimeRange) && showError}"
                    v-model="event.end_datetime"
                    @keyup="checkForErrors"
                />

                <input type="checkbox" class="mb1" id="show-end-time" v-model="event.show_end_time" />
                <label class="checkbox-label" for="show-end-time">Show time</label>

                <label for="type">Type</label>
                <select v-model="event.type" id="type">
                    <option value="0" selected disabled>Select event type...</option>
                    <option value="1">Tournament</option>
                    <option value="2">Other</option>
                </select>

                <label for="game-id">Game</label>
                <select v-model="event.game_id" id="game-id">
                    <option value="0" selected disabled>Choose game (if applicable)</option>
                    <option value="1">Fortnite</option>
                    <option value="2">Call of Duty</option>
                    <option value="3">Apex Legends</option>
                </select>

                <label for="description">Description</label>
                <textarea v-model="event.description" id="description"></textarea>

                <label for="referral-url">Referral Url</label>
                <input
                    type="text"
                    id="referral-url"
                    v-model="event.referral_url"
                    placeholder="http://google.com"
                />

                <label for="status">Status</label>
                <select v-model="event.status" id="status">
                    <option value="1">Active</option>
                    <option value="2">Inactive</option>
                </select>

                <label>Users</label>
                <div class="multi-select" v-for="index in numUsers" :key="index">
                    <select
                        type="text"
                        class="multi-input"
                        :class="{'form-error': !validUsers && showError}"
                        v-model="event.users[index - 1]"
                    >
                        <option value disabled>Select user...</option>
                        <option
                            v-for="user in users"
                            :key="user.id"
                            :value="user.id"
                        >{{ user.username }}</option>
                    </select>
                    <i class="fas fa-trash-alt select-icon" @click="removeUserField(index - 1)"></i>
                </div>
                <p class="form-action mb1" @click="addUserField">Add Another User</p>

                <label>Placements</label>
                <p>{{event.placemnts}}</p>
                <div class="multi-select" v-for="index in numPlacements" :key="index">
                    <input
                        type="number"
                        class="multi-input"
                        :class="{'form-error': !validPlacements && showError}"
                        v-model="event.placements[index - 1]"
                        placeholder="Enter placement..."
                        @keyup="checkForErrors"
                    />
                    <i
                        class="fas fa-trash-alt select-icon"
                        @click="removePlacementField(index - 1)"
                    ></i>
                </div>
                <p class="form-action" @click="addPlacementField">Add Another Placement</p>
            </fieldset>
            <fieldset>
                <div class="submit-wrapper">
                    <input type="submit" value="Save Event" :disabled="!validForm" />
                    <div class="submit-overlay" @mouseover="showError = true" @click="scrollToTop"></div>
                </div>
            </fieldset>

            <p v-if="showError" class="error text-center">{{ error }}</p>
        </form>
    </div>
</template>

<script>
import Validator from "@utils/Validator";

export default {
    props: {
        // Defines the existing event information
        editEvent: {
            type: Object,
            required: false
        },
        // Defines the list of users
        users: {
            required: true
        }
    },
    data() {
        return {
            // Form data
            event: this.editEvent
                ? {
                      ...this.editEvent,
                      start_datetime: this.editEvent.start_datetime.substring(
                          0,
                          16
                      ),
                      end_datetime: this.editEvent.end_datetime.substring(
                          0,
                          16
                      ),
                      promo_picture:
                          "/images/events/" + this.editEvent.id + ".jpg"
                  }
                : {
                      name: "",
                      location: "",
                      start_datetime: "",
                      show_start_time: false,
                      end_datetime: "",
                      show_end_time: false,
                      promo_picture: "/images/events/placeholder.jpg",
                      type: 0,
                      game_id: 0,
                      description: "",
                      referral_url: "",
                      status: 1,
                      users: [],
                      placements: []
                  },
            // Form control
            image_path: "",
            image_contents: "",
            numUsers: this.editEvent ? this.editEvent.users.length : 1,
            numPlacements: this.editEvent
                ? this.editEvent.placements.length
                : 1,
            // Error Management
            error: "Event name cannot be empty",
            showError: false
        };
    },
    computed: {
        // Determines if the event name is valid
        validName() {
            return this.event.name != "";
        },
        // Determines if the event start date/time is valid
        validStartDatetime() {
            return Validator.validDatetime(
                this.event.start_datetime,
                "-",
                "T",
                "YMD"
            );
        },
        // Determines if the event end date/time is valid
        validEndDatetime() {
            return Validator.validDatetime(
                this.event.end_datetime,
                "-",
                "T",
                "YMD"
            );
        },
        // Determines if the event time range is valid
        validTimeRange() {
            return (
                this.validStartDatetime &&
                this.validEndDatetime &&
                new Date(this.event.start_datetime) <
                    new Date(this.event.end_datetime)
            );
        },
        // Determines if the event placements is valid
        validPlacements() {
            if (this.event.placements.length > 1) {
                return (
                    new Set(this.event.placements).size ===
                    this.event.placements.length
                );
            }
            return true;
        },
        // Determines if the event end date/time is valid
        validUsers() {
            if (this.event.users.length > 1) {
                return (
                    new Set(this.event.users).size === this.event.users.length
                );
            }
            return true;
        },
        // Determines if the event form is valid
        validForm() {
            return (
                this.validName &&
                this.validStartDatetime &&
                this.validEndDatetime &&
                this.validTimeRange &&
                this.validPlacements &&
                this.validUsers
            );
        }
    },
    methods: {
        // Checks keystokes and updates form status
        checkForErrors() {
            if (!this.validName) {
                this.error = "Event name cannot be empty";
            } else if (!this.validPlacements) {
                this.error = "Placements cannot have duplicates";
            } else if (!this.validStartDatetime) {
                this.error = "Start date must be a valid date/time";
            } else if (!this.validEndDatetime) {
                this.error = "End date must be a valid date/time";
            } else if (!this.validTimeRange) {
                this.error =
                    "End date/time cannot be before the start date/time";
            } else if (!this.validPlacements) {
                this.error = "Placements cannot have duplicates";
            } else if (!this.validUsers) {
                this.error = "Users must not contain duplicates";
            } else {
                this.error = "";
            }
        },
        uploadImage(e) {
            // Get image
            this.image_contents = this.$refs.promopicture.files[0];

            // Set image source url to new image for preview
            this.event.promo_picture = URL.createObjectURL(e.target.files[0]);
        },
        save() {
            // Create payload
            let payload = {
                event: this.event
            };

            // Handle image
            // Create new form data with image information
            let image = new FormData();
            image.append("image", this.image_contents);
            image.append("folder", "events");
            image.append("name", "ViL_" + event.name);

            // Append image to payload
            payload.image = image;
            payload.image_path = this.image_path;

            // Save user
            this.$emit("save", payload);
        },
        // Add a placement field to the form
        addPlacementField() {
            this.numPlacements++;
            this.event.placements.push("");
        },
        // Remove a placement field from the form
        removePlacementField(i) {
            this.numPlacements--;
            this.event.placements = this.event.placements
                .slice(0, i)
                .concat(
                    this.event.placements.slice(
                        i + 1,
                        this.event.placements.length
                    )
                );
        },
        // Add a new user field to the form
        addUserField() {
            this.numUsers++;
            this.event.users.push("");
        },
        // Remove a user field from the form
        removeUserField(i) {
            this.numUsers--;
            this.event.users = this.event.users
                .slice(0, i)
                .concat(this.event.users.slice(i + 1, this.event.users.length));
        },
        // Moves the view to the beginning of the form
        scrollToTop() {
            document.getElementById("event-form").scrollIntoView({
                behavior: "smooth"
            });
        }
    }
};
</script>

<style scoped>
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

.event-form {
    width: 627px;
    margin: 0 auto;
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
    width: 627px;
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

.preview {
    height: 110px;
    resize: none;
}

.event-content {
    height: 370px;
    resize: none;
}

.promo-picture {
    margin-bottom: 20px;
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
    text-align: center;
    box-sizing: border-box;
    font-size: 20px;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    margin-bottom: 10px;
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

.form-action {
    padding-left: 8px;
    color: var(--color-secondary);
    cursor: pointer;
    transition: all 0.25s;
    font-size: 14px;
}

.form-action:hover {
    color: var(--color-secondary-light);
}

.multi-select {
    display: flex;
    flex-direction: row;
}

.select-icon {
    font-size: 3rem;
    margin-left: 20px;
    cursor: pointer;
    color: var(--color-secondary);
    transition: all 0.25s;
}

.select-icon:hover {
    color: var(--color-secondary-light);
}

input[type="checkbox"] {
    width: auto;
    display: inline;
    margin-left: 8px;
}

.checkbox-label {
    display: inline;
    font-size: 14px;
    padding: 0;
    font-family: inherit;
}
</style>
