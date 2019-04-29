<template>
  <div id="nav">
    <tv-nav-links/>
    <div id="lg-fill"/><div id="nav-wrapper"><div id="lg"><div id="nav-lg-logo"><router-link class="noselect" to="/" tag="img" src="/images/ui/nav_logo.png" @click.native="$store.dispatch('setNavPopup', false);"></router-link></div></div><div id="lg-md"/><div id="md"><tv-nav-sublinks/></div><div id="md-sm"/><div id="sm"/></div><div id="sm-fill"/>
  </div>
</template>

<script>
  import NavLinks       from './NavLinks.vue';
  import NavSublinks    from './NavSublinks.vue';

  export default {
    components: {
      tvNavLinks:       NavLinks,
      tvNavSublinks:    NavSublinks
    }
  }
</script>
<style scoped>

  #nav {
    margin: 0;
    padding: 0;
    width: 100%;
    height: 142px;
    position: absolute;
    top: 0;
    left: 0;
  }

  #nav div {
    display: inline-block;
  }

  #nav-wrapper {
    max-width: 992px;

  }

  #lg-fill {
    height: 142px;
    width: calc((100% - 992px)/2);
    background-image: url("/images/ui/nav_bg_lg.png");
  }

  #lg {
    height: 142px;
    width: 160px;
    background-image: url("/images/ui/nav_bg_lg.png");
    vertical-align: top;
    background-repeat: repeat-x;
  }

  #nav-lg-logo {
    padding-left: 20px;
    padding-top: 2px;
    cursor: pointer;
    width: 100%;
    height: auto;
    position: relative;
    z-index: 999999;
  }

  #lg-md {
    height: 142px;
    width: 103px;
    background-image: url("/images/ui/nav_bg_lg_md.png");
  }

  #md {
    height: 120px;

    min-width: 180px;
    width: 560px;
    background-image: url("/images/ui/nav_bg_md.png");
    vertical-align: top;
  }

  #md-sm {
    height: 120px;
    width: 69px;
    background-image: url("/images/ui/nav_bg_md_sm.png");
    vertical-align: top;
  }

  #sm {
    height: 80px;
    width: 100px;
    background-image: url("/images/ui/nav_bg_sm.png");
    vertical-align: top;
  }

  #sm-fill {
    height: 78px;
    width: calc((100% - 992px)/2);
    background-image: url("/images/ui/nav_bg_sm.png");
    vertical-align: top;
  }



  @media screen and (max-width: 628px) {
    #nav {
      zoom: .75;
    }
    #lg-fill, #lg, #lg-md {
      margin-top: 1px;
    }
    #md {
      min-width: 140px;
    }
    #sm {
      display: none !important;
    }
    #sm-fill {
      width: calc(100% - ((100vw - 449px) + 332px));
    }
  }
  @media screen and (max-width: 588px) {
    #sm-fill {
      width: calc(100% - 472px);
    }
  }
  @media screen and (max-width: 1008px) {
    #md {
      width: calc(100vw - 449px);
    }
  }
</style>
