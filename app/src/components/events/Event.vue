<template>

    <div class="upcoming-event">
        <img 
            :src="'/images/events/ishoni_2019.jpg'" 
            width="992" 
            alt="No image provided" 
        />
        <div class="upcoming-event-details">
            <h2>{{ event.name }}</h2>
            <p>{{ event.location }}</p>
            <p>{{ event.start_datetime | isoToDateTime }} - {{ event.end_datetime | isoToDateTime }}</p>
            <p>{{ event.description }}</p>
            <div class="buttons">
              <button @click="$emit('edit', entry.id);">
                <i class="far fa-edit"></i>
              </button>
              <button @click="$emit('remove', entry.id);">
                <i class="far fa-trash-alt"></i>
              </button>
            </div>
        </div>      
    </div>

</template>

<script>

    export default {
        props: {
            event: {
                type: Object,
                required: true
            }
        }
    }
    
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

  .past-events {
    background-color: var(--color-grey-dark);
    list-style: none;
    margin: 0;
    margin-bottom: 20px;
    padding: 0;
    font-size: 0;
  }

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

  .gold {
    color: #cead00;
  }

  .silver {
    color: #c0c0c0;
  }

  .bronze {
    color: #cd7f32;
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

  .buttons {
		text-align: right;
  }
  
  button {
		color: var(--color-secondary);
		display: inline;
		background: none;
		border: none;
		cursor: pointer;
		-webkit-touch-callout: none; /* iOS Safari */
   		-webkit-user-select: none; /* Safari */
     	-khtml-user-select: none; /* Konqueror HTML */
       	-moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
        user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
    transform: translateY(0);
    transition: all .25s;
	}

	button:hover {
		transform: translateY(-2px);
		color: var(--color-secondary-light);
	}

</style>
