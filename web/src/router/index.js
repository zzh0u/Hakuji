import { createRouter, createWebHistory } from 'vue-router'
import CallbackView from '../views/CallbackView.vue'
// import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
    },
    {
      path: '/book/:id',
      name: 'book-detail',
      component: () => import('../views/BookDetail.vue'),
    },
    {
      path: '/callback',
      name: 'callback',
      component: CallbackView,
    },
    // {
    //   path: '/login',
    //   name: 'login',
    //   component: () => import('../views/LoginPage.vue'),
    // },
    // {
    //   path: '/profile/edit',
    //   name: 'profile-edit',
    //   component: () => import('../views/ProfileEditView.vue'),
    //   meta: { requiresAuth: true },
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
  ],
})

export default router
