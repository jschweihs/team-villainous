<template>
        <footer>
            <nav class="footer-links">
                <p>Copyright &copy; 2020</p>
                <p 
                    v-if="isAdmin"
                    class="pointer"
                    @click="logout"
                >
                    Logout
                </p>
            </nav>
        </footer>
</template>

<script>

    import { mapGetters } from 'vuex';

    export default {
        computed: {
            ...mapGetters([
                'currentAdmin'
            ]),
            isAdmin() {
                return this.$route.path.includes('admin');
            }
        },
        methods: {
            logout() {
                this.$store.dispatch('logoutUser');
                this.$router.push('/');
            }
        }
    }
</script>

<style scoped>

    footer {
        bottom: 0;
        padding: 1rem 2rem;
        background: linear-gradient(var(--color-primary), var(--color-primary-dark));   
        position: relative;     
    }

    footer::before {
        content: '';
        height: .4rem;
        width: 100%;
        position: absolute;
        z-index: 1;
        left: 0;
        top: -.4rem;
        background-image: linear-gradient(var(--color-secondary-very-light), var(--color-secondary), var(--color-secondary-very-dark));
    }

    .footer-links {
        margin: 0 auto;
        display: flex;
        justify-content: space-between;
    }

</style>