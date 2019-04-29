<template>
  <div class="content">
    <h1>Upcoming Events</h1>
    <div class="upcoming-event">
      <img src="/images/events/mes_apr5.png"/>
      <div class="upcoming-event-details">
        <h2>MES 2019</h2>
        <p>COD</p>
        <p>Schwee, Tiggie, Deprive</p>
        <p>Chicago, IL</p>
        <p>Sep. 29th, 2009 - Sep. 30th, 2009</p>
        <p>Merch availible</p>
      </div>
    </div>
    <h1>Past Events</h1>
    <ul class="past-events">
      <li class="past-event">
        <div class="past-event-detail past-event-name">MES 2018</div>
        <div class="past-event-detail">Detroit, MI</div>
        <div class="past-event-detail">09/31/2018-10/02/2018</div>
        <div class="past-event-detail silver">
          <i class="fas fa-trophy"></i>
          2nd place
        </div>
      </li>
      <li class="past-event past-event-alt">
        <div class="past-event-detail past-event-name">MES 2017</div>
        <div class="past-event-detail">Orlando, FL</div>
        <div class="past-event-detail">04/05/2017-04/07/2017</div>
        <div class="past-event-detail bronze">
          <i class="fas fa-trophy"></i>
          3rd place
        </div>
      </li>
      <li class="past-event">
        <div class="past-event-detail past-event-name">MES 2016</div>
        <div class="past-event-detail">Columbus, OH</div>
        <div class="past-event-detail">04/03/2016-04/04/2016</div>
        <div class="past-event-detail gold">
          <i class="fas fa-trophy"></i>
          1st place
        </div>
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
    min-width: 100px;
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
