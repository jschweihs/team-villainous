<template>
  <div class="home">
    <div class="main-banner"/>
    <div class="home-container">
      <div class="blog" v-if="blog">
        <blog-entry
          v-for="entry in blog"
          :key="entry.id"
          :entry="entry"/>
      </div>
      <div class="sidebar">
        <div class="sidebar-content">
          <twitter>
            <a
              class="twitter-timeline"
              data-width="300"
              data-theme="dark"
              data-tweet-limit="3"
              data-link-color="#FAB81E"
              href="https://twitter.com/villainous_team?ref_src=twsrc%5Etfw"
            >
              Tweets by Team Villainous
          </a>
          </twitter>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

  import {mapGetters} from 'vuex';

  import BlogEntry from './../../components/blog/BlogEntry.vue';

  export default {
    components: {
      BlogEntry
    },
    data() {
      return {
        w_width: 300
      };
    },
    computed: {
      ...mapGetters([
        'blog'
      ])
    },
    created() {
      this.$store.dispatch('showModal', true);
      this.$store.dispatch('getBlog')
      .then(res => {
        this.$store.dispatch('showModal', false);
      })
      .catch(e => {
        this.$store.dispatch('showModal', true);
        console.log(e);
      });
    }
  }
</script>

<style>
.home {
  /* background-color: #262626; */
}
.main-banner {
  background-image:  url("/images/recruiting.png");
  background-size: contain;
  padding-top: 28.64%;
  width: 100%;
  box-shadow: 0 10px 40px #111;
}

.home-container {
  max-width: 992px;
  margin: 0 auto;
}
.blog {
  max-width: 627px;
  display: inline-block;
  margin: 20px 0;
  margin-top: 20px;
  padding: 20px;
  background-color: var(--color-grey-dark);
  vertical-align:top;
  border-radius: 30px;
  box-shadow: 0 -2px 10px #1C1C1C;
}

.blog h1 {
  color: white;
  text-align: center;
  margin: 0;
  font-weight: 100;
}

.blog h3 {
  color: white;
  font-size: 16px;
  text-align: center;
  font-weight: 100;
  margin: 0 0 16px 0;
}

.blog p {
  text-indent: 16px;
  letter-spacing: .25px;
  margin-bottom: 1rem;
}

.blog-promo {
  width: 100%;
  height: auto;
  padding: 0;
  box-shadow: 0 0 4px #222;
}
.blog hr {
  border-color: #666;
  margin-bottom: 16px;
}

.sidebar {
display: inline-block;
}

.sidebar-bg {
  min-height: 100%;
  width: 33.3%;
  min-width: 180px;
  background-color: #1a1a1a;
  z-index: -1;
  overflow: auto;
  display: inline-block;
  float: right;
}

.sidebar-content {
  padding: 20px;
  padding-right: 0;
}

@media screen and (max-width: 628px) {
  .main-banner {
    /* margin-top: 106px; */
  }
  .blog h1 {
    font-size: 24px;
  }
}

@media screen and (max-width: 684px) {
  .blog {
    border-radius: 0;
    margin-top: 10px;
  }
}

@media screen and (max-width: 1011px) {
  .blog {
    margin: 10px auto;
    padding-top: 10px;
    display: block;
  }

  .sidebar {
    display: none;
  }
}

</style>
