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
          <span class="sublink-text">Twitch.tv</span>
        </a>
      </li>
      <li>
        <a class="noselect" href="#" target="_blank">
          <i class="fab fa-youtube-square"></i>
          <span class="sublink-text">Youtube</span>
        </a>
      </li>
      <li>
        <a class="noselect" href="#" target="_blank">
          <i class="fab fa-facebook-square"></i>
          <span class="sublink-text">Facebook</span>
        </a>
      </li>
      <li>
        <a class="noselect" href="#" target="_blank">
          <i class="fab fa-twitter-square"></i>
          <span class="sublink-text">Twitter</span>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        current_route: this.$route.path
      };
    }
  }
</script>

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
    color: #CCC;
    display: block;
    font-weight: 900;
  }

  @media screen and (max-width: 960px) {
    .sublink-text {
      display: none;
    }
  }
</style>
