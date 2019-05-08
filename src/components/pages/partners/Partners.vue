<template>
  <div class="content half">
    <h1>Partners</h1>
    <ul class="partners">
      <li class="partner">
        <a href="http://erivals.com" target="_blank">
          <img width="180" height="180" src="/images/partners/erivals.jpg"/>
        </a>
        <p>Team Villainous has partnered with Erivals to bring top tier talent to their competitive gaming platform. Erivals host tournaments that are both Free to Enter and & Pay to Play. We are proud to have Erivals as a valued partner. Go visit their website and sign up by clicking the logo.</p>
      </li>
      <li class="partner">
        <a href="https://aporiacustoms.com/" target="_blank">
          <img width="180" height="180" src="/images/partners/aporiacustoms.jpg"/>
        </a>
        <p>Aporia Customs is the sole producer for all Villainous merchanidse. Aporia Customs uses the highest quality of material offered and takes great pride in their hand crafted work. Do not forget to use code "ViL5" for 5% off at checkout! You can locate our store at the "Shop" tab in the navigation bar.</p>
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
    background-color: #393939;
  }

  .partner:last-child {
    margin-bottom: 0;
  }

</style>
