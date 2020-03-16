<template>
  <div id="nav-links">
    <div id="nav-links-bar">
      <ul>
        <li>
          <router-link class="noselect" to="/team" tag="a">Team</router-link>
        </li>
        <li>
          <router-link class="noselect" to="/partners" tag="a">Partners</router-link>
        </li>
        <li>
          <router-link class="noselect" to="/events" tag="a">Events</router-link>
        </li>
        <li>
          <router-link class="noselect" to="/shop" tag="a">Shop</router-link>
        </li>
        <li>
          <router-link class="noselect" to="/contact" tag="a">Contact</router-link>
        </li>
      </ul>
    </div>
    <div id="nav-links-btn" @click="toggleNavPopup">
      <div class="nav-links-btn-cmpt1"/>
      <div class="nav-links-btn-cmpt2"/>
      <div class="nav-links-btn-cmpt2"/>
    </div>
    <transition name="fade" mode="out-in">
      <div id="nav-links-popup" v-if="display_nav_popup">
        <ul>
          <li>
            <router-link
              class="noselect"
              to="/"
              tag="a"
            >
              <span class="invisible" @click="toggleNavPopup">Home</span>
            </router-link>
          </li>
          <li>
            <router-link
              class="noselect"
              to="/team"
              tag="a"
            >
              <span @click="toggleNavPopup">Team</span>
            </router-link>
          </li>
          <li>
            <router-link class="noselect" to="/partners" tag="a">
              <span @click="toggleNavPopup">Partners</span>
            </router-link>
          </li>
          <li>
            <router-link class="noselect" to="/events" tag="a">
              <span @click="toggleNavPopup">Events</span>
            </router-link>
          </li>
          <li>
            <router-link class="noselect" to="/shop" tag="a">
              <span @click="toggleNavPopup">Shop</span>
            </router-link>
          </li>
          <li>
            <router-link class="noselect" to="/contact" tag="a">
              <span @click="toggleNavPopup">Contact</span>
            </router-link>
          </li>
        </ul>
        <div class="close" @click="toggleNavPopup">
          <div class="circle"/>
          <div class="x left"/>
          <div class="x right"/>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>

  import {mapGetters} from 'vuex';

  export default {
    data() {
      return {
        window_width: 0
      };
    },
    computed: {
      ...mapGetters([
        'display_nav_popup'
      ])
    },
    watch: {
      window_width(width) {
        if(width > 712) {
          this.$store.dispatch('setNavPopup', false);
        }
      }
    },
    methods: {
      toggleNavPopup: function() {
        this.$store.dispatch('toggleNavPopup');
      }
    },
    mounted() {
      this.$nextTick(() => {
        window.addEventListener('resize', () => {
          this.window_width = window.innerWidth;
        })
      });
    }
  }
</script>
<style scoped>
  a {
    font-family: "Geizer";
  }
  #nav-links {
    position: absolute;
    z-index: 100;
    margin: 0;
    padding: 0;
    width: 100%;
  }

  #nav-links-bar ul {
    list-style: none;
    padding: 0;
    margin-left: calc(((100% - 992px) / 2) + 280px);
    max-width: 992px;
    width: 50%;
    margin-top: 20px;
    display: table;
  }

  #nav-links-bar ul li {
    text-align: center;
    display: table-cell;
    padding: 0 20px;
  }

  #nav-links-bar ul li a {
    font-size: 42px;
    text-shadow: 0 -4px 8px black;
    color: white;
    display: block;
    transform: translateY(0);
    transition: all .25s;
  }

   #nav-links-bar ul li a:hover {
    transform: translateY(-2px);
    text-shadow: 0 -10px 20px black;
    color: #ffc200;
  }

  #nav-links-btn {
    position: absolute;
    top: 16px;
    right: 30px;
    width: 40px;
    height: 40px;
    z-index: 10000;
    cursor: pointer;
    display:  none;
  }

  .nav-links-btn-cmpt1 {
    background-color: white;
    width: 40px;
    height: 6px;
    border-radius: 4px;
  }

  .nav-links-btn-cmpt2 {
    background-color: white;
    width: 40px;
    height: 6px;
    margin-top: 10px;
    border-radius: 4px;
  }

  #nav-links-popup {
    background-color: #333;
    border-radius: 4px;
    position: absolute;
    right: 20px;
    left: 20px;
    top: 64px;
    box-shadow: 0 0 20px black;
    z-index: 200;

  }
  #nav-links-popup ul {
    list-style: none;
    margin: 0;
    padding: 10px 20px;
  }

  #nav-links-popup ul li {
    padding: 10px 0;

  }

  #nav-links-popup ul li:not(:last-child) {
    border-bottom: 1px solid #222;
  }

  #nav-links-popup ul li a {
    color: white;
    font-size: 42px;
  }

  .close {
    position: absolute;
    top: 11px;
    right: 42px;
    cursor: pointer;
  }
  .circle {
    width: 20px;
    height: 20px;
    border: 4px solid white;
    border-radius: 50px/50px;
    position: absolute;
  }

  .x {
    width: 4px;
    height: 22px;
    left: 12px;
    top: 3px;
    background-color: white;
    position: absolute;
  }

  .left {
    transform: rotate(45deg);
  }

  .right {
    transform: rotate(-45deg);
  }

  @media screen and (max-width: 810px) {
    #nav-links-bar ul {
      display: none;
    }
    #nav-links-btn {
      display: block;
    }
  }

  @media screen and (max-width: 960px) {
    #nav-links-bar ul {
      margin-left: calc(((100% - 744px) / 2) + 180px);
      margin-top: 30px;
    }
    #nav-links-bar ul li a {
      font-size: 32px;
      text-shadow: 0 -4px 8px black;
      color: white;
      display: block;
    }
  }
</style>
