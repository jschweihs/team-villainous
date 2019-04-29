// Main
import Home           from './components/pages/home/Home.vue';
import Team           from './components/pages/team/Team.vue';
import Partners       from './components/pages/partners/Partners.vue';
import Events         from './components/pages/events/Events.vue';
import Shop           from './components/pages/shop/Shop.vue';
import Contact        from './components/pages/contact/Contact.vue';

// Admin
import AdminLogin     from './components/pages/admin/login/Login.vue';
import AdminHome      from './components/pages/admin/home/Home.vue';
import AdminUsers     from './components/pages/admin/users/Users.vue';
import AdminNewUser   from './components/pages/admin/users/NewUser.vue';
import AdminEditUser  from './components/pages/admin/users/EditUser.vue';
import AdminBlog      from './components/pages/admin/blog/Blog.vue';
import AdminNewBlog   from './components/pages/admin/blog/NewBlog.vue';
import AdminEditBlog  from './components/pages/admin/blog/EditBlog.vue';

export const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/team',
    component: Team
  },
  {
    path: '/partners',
    component: Partners
  },
  {
    path: '/events',
    component: Events
  },
  {
    path: '/shop',
    component: Shop
  },
  {
    path: '/contact',
    component: Contact
  },
  // Admin
  {
    path: '/admin',
    component: AdminLogin
  },
  {
    path: '/admin/home',
    component: AdminHome
  },
  // Users
  {
    path: '/admin/users',
    component: AdminUsers
  },
  {
    path: '/admin/users/add',
    component: AdminNewUser
  },
  {
    path: '/admin/users/:id',
    component: AdminEditUser
  },
  // Blog
  {
    path: '/admin/blog',
    component: AdminBlog
  },
  {
    path: '/admin/blog/new',
    component: AdminNewBlog
  },
  {
    path: '/admin/blog/:id',
    component: AdminEditBlog
  },
  // Default
  {
    path: '/*',
    component: Home
  },
];
