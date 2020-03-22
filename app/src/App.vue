<template>
    <div id="app">
        <vil-nav></vil-nav>
        <div class="body-wrapper">
            <transition name="slide" mode="out-in">
                <router-view />
            </transition>
        </div>
        <foot></foot>

        <!-- Loading modal -->
        <loading-modal></loading-modal>
    </div>
</template>

<script>
import Nav from "./components/nav/Nav.vue";
import Foot from "./components/footer/Foot.vue";
import LoadingModal from "./components/utils/LoadingModal.vue";

import Cookie from "./utils/Cookie";

export default {
    name: "app",
    components: {
        LoadingModal,
        vilNav: Nav,
        Foot
    },
    beforeCreate() {
        // Attempt to grab current user
        this.$store.dispatch("getCurrentUser");
    }
};
</script>

<style>
*,
*::before,
*::after {
    margin: 0;
    padding: 0;
}

html,
body {
    height: 100%;
}

body {
    background-color: var(--color-primary-dark);
}

#app {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    min-height: 100%;
}

.body-wrapper {
    background-color: var(--color-grey-very-dark);
    /* padding: 3rem 6rem; */
    flex: 1;
    padding-top: 2rem;
    padding-bottom: 2rem;
}

.slide-enter-active {
    animation: slide-in 200ms easy-out forwards;
}
.slide-leave-active {
    animation: slide-out 200ms ease-out forwards;
}

@keyframes slide-in {
    from {
        transform: translateY(-30px);
        opacity: 0;
    }
    to {
        transform: translateY(0);
        opacity: 1;
    }
}

@keyframes slide-out {
    from {
        transform: translateY(0);
        opacity: 1;
    }
    to {
        transform: translateY(-30px);
        opacity: 0;
    }
}
</style>
