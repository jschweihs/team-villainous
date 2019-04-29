<template>
  <div class="content half">
    <h1>Partners</h1>
    <ul class="partners">
      <li class="partner">
        <a href="#">
          <img width="180" height="180" src="/images/partners/erivals.jpg"/>
        </a>
        <p>eRivals is a safe online platform for gamers to compete in paid wagers or free tournaments.</p>
      </li>
      <li class="partner">
        <a href="#">
          <img width="180" height="180" src="/images/partners/reloadfocus.jpg"/>
        </a>
        <p>Description here...</p>
      </li>
      <li class="partner">
        <a href="#">
          <img width="180" height="180" src="/images/partners/aporiacustoms.jpg"/>
        </a>
        <p>Description here...</p>
      </li>
    </ul>
  </div>
</template>

<style scoped>
  .partners {
    margin: 0 auto;
    padding: 0;
    list-style: none;
    max-width: 600px;
  }

  .partner {
    clear: both;
    margin-bottom: 10px;
    height: 180px;
    background-color: #333;
    border-top-left-radius: 20px;
    border-bottom-left-radius: 20px;
  }

  .partner img {
    float: left;
    max-height: 180px;
    margin-right: 20px;
    border-top-left-radius: 20px;
    border-bottom-left-radius: 20px;
  }

  .partner p {
    padding: 20px;
    height: 138px;
    margin-left: 180px;
    border-top: 1px solid #444;
    border-right: 1px solid #444;
    border-bottom: 1px solid #444;
    border-top-right-radius: 20px;
    border-bottom-right-radius: 20px;
  }

  .partner:last-child {
    margin-bottom: 0;
  }

</style>
