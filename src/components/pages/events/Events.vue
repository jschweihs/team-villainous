<template>
  <div class="content">
    <h1>Upcoming Events</h1>
    <div class="upcoming-event">
      <img src="/images/events/dreamhack_dallas_2019.png" width="992"/>
      <div class="upcoming-event-details">
        <h2>Dreamhack Dallas 2019</h2>
        <p>Halo</p>
        <p>Brianna will be attending</p>
        <p>Dallas, TX</p>
        <p>May 31st - June 2nd 2019</p>
        <p>Business Cards Available</p>
      </div>
    </div>
    <h1>Past Events</h1>
    <ul class="past-events">
      <li class="past-event past-event-alt">
        <div class="past-event-detail past-event-name">MES 2019</div>
        <div class="past-event-detail">Detroit, MI</div>
        <div class="past-event-detail">05/05/2019-05/05/2019</div>
        <div class="past-event-detail">
          <span class="place bronze">
            <i class="fas fa-trophy"></i>
            3rd place
          </span>, 5th place
        </div>
        <!-- <div class="past-event-detail">
          5th place
        </div> -->
      </li>
      <li class="past-event">
        <div class="past-event-detail past-event-name">MES 2018</div>
        <div class="past-event-detail">Detroit, MI</div>
        <div class="past-event-detail">10/26/2018-10/28/2018</div>
        <div class="past-event-detail">
          <span class="gold">
            <i class="fas fa-trophy"></i>
            1st place
          </span>, 
          <span class="silver">
            <i class="fas fa-trophy"></i>
            2nd place
          </span>
        </div>
        <!-- <div class="past-event-detail silver">
          <i class="fas fa-trophy"></i>
          2nd place
        </div> -->
      </li>
    </ul>
  </div>
</template>

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
    border: 1px solid #444;
    border-top: 0;
    border-bottom-left-radius: 20px;
    border-bottom-right-radius: 20px;
    padding: 20px;
    background-color: #393939;
  }

  .past-events {
    background-color: #333;
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
    background-color: #393939;
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

</style>
