<template>
  <div class="content">
    <div class="admin-nav">
      <router-link class="admin-link" to="/admin/users" tag="button"
        >Manage Users</router-link
      >
      <router-link class="admin-link" to="/admin/blog" tag="button"
        >Manage Blog</router-link
      >
      <router-link class="admin-link" to="/admin/events" tag="button"
        >Manage Events</router-link
      >
      <router-link class="admin-link" to="/admin/roles" tag="button"
        >Manage Roles</router-link
      >
    </div>

    <section class="dashboard">
      <h1>Dashboard coming soon!</h1>
    </section>
  </div>
</template>

<script></script>

<style scoped>
.admin-nav {
  display: flex;
  flex-direction: row;
  margin-bottom: 2rem;
}

.admin-link {
  flex: 1;
  width: 100%;
  font-size: 24px;
  padding: 10px;
  margin: 10px;
  border-radius: 8px;
  border: 0;
  -webkit-box-sizing: border-box; /* Safari/Chrome, other WebKit */
  -moz-box-sizing: border-box; /* Firefox, other Gecko */
  box-sizing: border-box; /* Opera/IE 8+ */
  color: white;
  height: 60px;
  cursor: pointer;
  border-radius: 8px;
  font-family: "Geizer", cursive;
  letter-spacing: 2px;
  background-color: var(--color-secondary);
  transform: translateY(0);
  transition: all 0.25s;
}

.admin-link:hover {
  background-color: var(--color-secondary-light);
  transform: translateY(-2px);
}

@media screen and (max-width: 800px) {
  .two-column {
    flex: 100%;
  }
}
</style>
