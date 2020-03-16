<template>
  <div class="content">
    <h1>Add New Event</h1>
    <div class="event-form">
      <event-form @save="addEvent" :users="users" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

import EventForm from "@components/events/EventForm.vue";

export default {
  components: {
    EventForm
  },
  computed: {
    ...mapGetters(["users"])
  },
  methods: {
    addEvent(event) {
      this.$store.dispatch("showModal", true);
      return this.$store
        .dispatch("addEvent", event)
        .then(() => {
          this.$store.dispatch("showModal", false);
          this.$router.push("/admin/events");
        })
        .catch(e => {
          this.$store.dispatch("showModal", false);
          console.log(e);
        });
    }
  },
  created() {
    this.$store.dispatch("showModal", true);

    this.$store
      .dispatch("getUsers")
      .then(res => {
        this.$store.dispatch("showModal", false);
      })
      .catch(e => {
        console.log(e);
        this.$store.dispatch("showModal", false);
      });
  }
};
</script>

<style scoped>
.content {
  background-color: var(--color-grey-dark);
  border-radius: 20px;
  padding: 20px;
}

h1 {
  border-bottom: 1px solid var(--color-grey-light);
  margin: 0;
  margin-bottom: 20px;
}
</style>
