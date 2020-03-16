<template>
	<div class="event-form">
		<form @submit.prevent="save">
			<label>Name *</label>
			<input type="text" v-model="event.title"/>
			<label class="img-label">Promo Picture</label>
			<img 
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
			<label>Location</label>
			<input type="text" v-model="event.location" placeholder="Allstate Arena in Chicago, IL"/>

            {{ event.showStartTime }}
            <label>Starts At</label>
			<input type="datetime-local" v-model="event.start_datetime" />

            <input type="checkbox" class="mb1" id="show-start-time" v-model="event.showStartTime">
            <label class="checkbox-label" for="show-start-time">Show time</label>

            <label>Ends At</label>
			<input type="datetime-local" v-model="event.end_datetime" />

            <input type="checkbox" class="mb1" id="show-end-time" v-model="event.showEndTime">
            <label class="checkbox-label" for="show-end-time">Show time</label>

            <label>Type</label>
            <select v-model="event.type">
                <option value="" selected disabled>Select event type...</option>
                <option value="1">Tournament</option>
                <option value="2">Other</option>
            </select>

            <label>Game</label>
            <select v-model="event.game_id">
                <option value="0" disabled selected="selected">Choose game (if applicable)</option>
                <option value="1">Fortnite</option>
                <option value="2">Call of Duty</option>
                <option value="3">Apex Legends</option>
            </select>

            <label>Description</label>
			<textarea v-model="event.description"></textarea>

            <label>Referral Url</label>
			<input type="text" v-model="event.referral_url" placeholder="http://google.com"/>

            <label>Status</label>
            <select v-model="event.status">
                <option value="1">Active</option>
                <option value="2">Inactive</option>
            </select>

            <label>Users</label>
            <div 
                class="multi-select" 
                v-for="index in numUsers" 
                :key="index"
            >
                <select 
                    type="text" 
                    class="multi-input"
                    v-model="event.users[index-1]" 
                >
                    <option value='' disabled selected>Select user...</option>
                    <option
                        v-for="user in users"
                        :key="user.id"
                        :value="user.id"
                    >
                        {{ user.username }}
                    </option>
                </select>
                <i class="fas fa-trash-alt select-icon" @click="removeUserField(index-1)"></i>
            </div>
            <p class="form-action mb1" @click="addUserField">Add Another User</p>

            <label>Placements</label>
            <div 
                class="multi-select" 
                v-for="count in numPlacements" 
                :key="count" 
            >
                <input 
                    type="text" 
                    class="multi-input"
                    v-model="event.placements[count-1]" 
                    placeholder="Enter placement..."
                />
                <i class="fas fa-trash-alt select-icon" @click="removePlacementField"></i>
            </div>
            <p class="form-action" @click="addPlacementField">Add Another Placement</p>

			<input type="submit" value="Save Event"/>
		</form>
	</div>
</template>

<script>
	export default {
		props: {
			editEvent: {
				type: Object,
				required: false
            },
            users: {
                required: true
            }
        },
		computed: {
			current_user() {
				return this.$store.getters.current_user;
			},
			image_name() {
				return this.event.title.toLowerCase().replace(' ', '_');
            },
            user_id() {
                return this.$store
            }
		},
		data() {
			return {
				image_path:     '',
                image_contents: '',
                numUsers:       1,
                numPlacements:  1,
				event: this.editEvent
					? { 
						...this.editEvent, 
						promo_picture: '/images/events/' + this.editEvent.id + '.jpg'
					} 
					: {
						title: 				'',
						user_id: 			this.current_user ? this.current_user.id : "1",
						username: 			this.current_user ? this.current_user.username : "Anon",
                        location:           '',
                        start_datetime:     '',
                        showStartTime:      false,
                        end_datetime:       '',
                        promo_picture: 		'/images/events/placeholder.jpg',
                        type:               '',
                        game_id:            0,
                        description:        '',
                        referral_url:       '',
                        status:             1,
                        users:              [''],
                        placements:         []
					}
			};
		},
		methods: {
			uploadImage(e) {
				this.image_contents = this.$refs.promopicture.files[0];
				this.event.promo_picture = URL.createObjectURL(e.target.files[0]);
			},
			save() {
                console.log("event:", this.event);

				// Prep file
				// let image = new FormData();
				// image.append('image', this.image_contents);
				// image.append('folder', 'events');
				// image.append('name', this.event.id);

				// this.$emit('save', { event: this.event, image, image_path: this.image_path });
            },
            addPlacementField() {
                this.numPlacements++;
                this.event.placements.push('');
                
            },
            removePlacementField(i) {
                this.numPlacements--;
                this.event.placements = this.event.placements.slice(0, i).concat(this.event.placements.slice(i + 1, this.event.placements.length));
            },
            addUserField() {
                this.numUsers++;
                this.event.users.push('');
            },
            removeUserField(i) {
                this.numUsers--;
                this.event.users = this.event.users.slice(0, i).concat(this.event.users.slice(i + 1, this.event.users.length));
            }
        }
	}
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

	img {
		width: 627px;
		display: block;
		margin: 0 auto;
		border-top-left-radius: 20px;
		border-top-right-radius: 20px;
		cursor: pointer;
	}

	label {
		color: #DEDEDE;
		display: block;
		font-size: 18px;
		margin-bottom: 4px;
		padding-left: 8px;
	}

	input, select, textarea {
		display: block;
		width: 100%;
		font-size: 24px;
		padding: 10px;
		margin-bottom: 10px;
		border-radius: 8px;
		border: 0;
		-webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
    	-moz-box-sizing: border-box;    /* Firefox, other Gecko */
    	box-sizing: border-box;         /* Opera/IE 8+ */
    	font-family: Nixie One,cursive;
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

	input[type=submit] {
		color: white;
		margin-top: 20px;
		background-color: var(--color-secondary);
		border: 0;
		height: 60px;
		cursor: pointer;
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
		border-radius: 8px;
		transform: translateY(0);
		transition: all .25s;
	}

	input[type=submit]:hover {
		transform: translateY(-2px);
		background-color: var(--color-secondary-light);
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
		font-family: 'Geizer', cursive;
		letter-spacing: 2px;
		line-height: 35px;
	    text-align: center;
	    box-sizing: border-box;
	    font-size: 20px;
	    border-bottom-left-radius: 20px;
	    border-bottom-right-radius: 20px;
		margin-bottom: 10px;
		transition: all .25s;
	}

	.file-label:hover {
		background-color: var(--color-secondary-light);
	}

	input[type=file] {
		opacity: 0;
	   	position: absolute;
	   	z-index: -1;
	}

    .form-action {
        padding-left: 8px;
        color: var(--color-secondary);
        cursor: pointer;
        transition: all .25s;
        font-size: 14px;
    }

    .form-action:hover {
        color: var(--color-secondary-light);
    }

    .multi-select{
        display: flex;
        flex-direction: row;
    }

    .select-icon {
        font-size: 3rem;
        margin-left: 20px;
        cursor: pointer;
        color: var(--color-secondary);
        transition: all .25s;
    }

    .select-icon:hover {
        color: var(--color-secondary-light);
    }

    input[type=checkbox] {
        width: auto;
        display: inline;
        margin-left: 8px;
    }

    .checkbox-label {
        display: inline;
        font-size: 14px;
        padding: 0;
    }

</style>