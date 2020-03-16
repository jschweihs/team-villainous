<template>
    <transition name="fade">
        <div class="modal" v-if="showModal">

            <!-- Modal content -->
            <div class="modal-content">

                <div class="loader-wrapper">
                    <div class="loader"></div>
                </div>
                
            </div>

        </div>
    </transition>
</template>

<script>

    import { mapGetters } from 'vuex';

    export default {
        computed: {
            ...mapGetters([
                'showModal'
            ])
        }
        
    }

</script>

<style scoped>

/* Modal */
.modal {
  position: fixed;
  z-index: 9999999;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  margin: 0 auto;
}


/* Spinner */

.loader-wrapper {
	position: relative;
	width: 100%;
	width: 120px;
	height: 120px;
	margin: 0 auto;
	margin-top: calc(50vh - 60px);

}

.loader {
  border: 16px solid #f3f3f3; /* Light grey */
  border-top: 16px solid var(--color-secondary);; /* Blue */
  border-radius: 50%;
  width: 120px;
  height: 120px;
  animation: spin 1s linear infinite;
  position: absolute;
  z-index: 99999999;
  margin: 0 auto;

}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

</style>
