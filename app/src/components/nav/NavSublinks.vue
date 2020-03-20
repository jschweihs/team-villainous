<template>
  <div id="nav-md-sublinks">

    <ul>

      <li>
        <a
          class="noselect"
          href="http://www.twitch.tv/team_villainous"
          target="_blank"
        >
          <i class="fab fa-twitch"></i>
          <span class="sublink-text"> Twitch.tv</span>
        </a>
      </li>

      <li>
        <a 
          class="noselect" 
          href="https://www.youtube.com/channel/UCP_WLhGqtjn2ocKtuytMS3Q" 
          target="_blank"
        >
          <i class="fab fa-youtube-square"></i>
          <span class="sublink-text"> Youtube</span>
        </a>
      </li>

      <li>
        <a 
          class="noselect" 
          href="https://www.facebook.com/TeamVillainous/" 
          target="_blank"
        >
          <i class="fab fa-facebook-square"></i>
          <span class="sublink-text"> Facebook</span>
        </a>
      </li>

      <li>
        <a 
          class="noselect" 
          href="https://twitter.com/Villainous_Team" 
          target="_blank"
        >
          <i class="fab fa-twitter-square"></i>
          <span class="sublink-text"> Twitter</span>
        </a>
      </li
      >
    </ul>

  </div>
</template>

<style scoped>
  #nav-md-sublinks {
    margin-top: 78px;
    width: 100%;
  }

  #nav-md-sublinks ul {
    list-style: none;
    margin: 0;
    padding: 0;
    display: table;
    width: 100%;
  }

  #nav-md-sublinks ul li {
    text-align: center;
    display: table-cell;
    padding: 0 10px;
  }

  #nav-md-sublinks ul li a {
    font-size: 18px;
    color: #EEE;
    display: block;
    font-weight: 900;
    transform: translateY(0);
    transition: all .25s;
  }

  #nav-md-sublinks ul li a:hover {
    color: #ffc200;
    transform: translateY(-2px);
  }

  @media screen and (max-width: 960px) {
    .sublink-text {
      display: none;
    }
  }
</style>
