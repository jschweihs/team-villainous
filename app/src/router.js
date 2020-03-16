// Main
import Vue from "vue";
import Router from "vue-router";
import store from "./store/store";

import Home from "./pages/home/Home.vue";
import Team from "./pages/team/Team.vue";
import Partners from "./pages/partners/Partners.vue";
import Events from "./pages/events/Events.vue";
import Shop from "./pages/shop/Shop.vue";
import Contact from "./pages/contact/Contact.vue";

// Admin
import AdminLogin from "./pages/login/Login.vue";
import AdminHome from "./pages/admin/home/Home.vue";
import AdminRoles from "./pages/admin/roles/Roles.vue";
import AdminUsers from "./pages/admin/users/Users.vue";
import AdminNewUser from "./pages/admin/users/NewUser.vue";
import AdminEditUser from "./pages/admin/users/EditUser.vue";
import AdminBlog from "./pages/admin/blog/Blog.vue";
import AdminNewBlog from "./pages/admin/blog/NewBlog.vue";
import AdminEditBlog from "./pages/admin/blog/EditBlog.vue";
import AdminEvents from "./pages/admin/events/Events.vue";
import AdminNewEvent from "./pages/admin/events/NewEvent.vue";
import AdminEditEvent from "./pages/admin/events/EditEvent.vue";

Vue.use(Router);

let router = new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      component: Home
    },
    {
      path: "/team",
      component: Team
    },
    {
      path: "/partners",
      component: Partners
    },
    {
      path: "/events",
      component: Events
    },
    {
      path: "/shop",
      component: Shop
    },
    {
      path: "/contact",
      component: Contact
    },
    // Admin
    {
      path: "/login",
      component: AdminLogin,
      meta: {
        skipWithLogin: true
      }
    },
    {
      path: "/admin",
      component: AdminHome,
      meta: {
        requiresLogin: true
      }
    },
    // Admin Roles
    {
      path: "/admin/roles",
      component: AdminRoles,
      meta: {
        requiresLogin: true
      }
    },
    // Admin Users
    {
      path: "/admin/users",
      component: AdminUsers,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/users/new",
      component: AdminNewUser,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/users/:id",
      component: AdminEditUser,
      meta: {
        requiresLogin: true
      }
    },
    // Admin Blog
    {
      path: "/admin/blog",
      component: AdminBlog,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/blog/new",
      component: AdminNewBlog,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/blog/:id",
      component: AdminEditBlog,
      meta: {
        requiresLogin: true
      }
    },
    // Admin Events
    {
      path: "/admin/events",
      component: AdminEvents,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/events/new",
      component: AdminNewEvent,
      meta: {
        requiresLogin: true
      }
    },
    {
      path: "/admin/events/:id",
      component: AdminEditEvent,
      meta: {
        requiresLogin: true
      }
    },
    // Default
    {
      path: "/*",
      component: Home
    }
  ]
});

// Check for token.  Seems messy that this would happen here

// const token = Cookie.getCookie('token');

// store.dispatch("setCurrentUser", token);

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresLogin)) {
    // Requires login
    const token = store.getters.token;
    if (store.getters.token) {
      next();
      return;
    }
    next("/login");
  } else if (to.matched.some(record => record.meta.skipWithLogin)) {
    if (store.getters.token) {
      next("/admin");
      return;
    }
    next();
  } else {
    next();
  }
});

export default router;
