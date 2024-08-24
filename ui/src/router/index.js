import { createRouter, createWebHistory } from 'vue-router'
import app from '../stores/app'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: 'login',
      component: () => import("@/views/login/LoginView.vue")
    },
    {
      path: "/",
      name: 'frontend',
      component: () => import("@/views/frontend/FrontendView.vue"),
      children: [
        {
          path: "blogs/list",
          name: "FrontendBlogList",
          component: () => import('../views/frontend/Blog/BlogListView.vue')
        },
        {
          path: "blogs/info",
          name: "FrontendBlogEdit",
          component: () => import('../views/frontend/Blog/BlogInfoView.vue')
        }
      ]
    },
    {
      path: "/backend",
      name: 'backend',
      component: () => import('@/views/backend/BackendView.vue'),
      redirect: { name: 'BackendBlogList' },
      beforeEnter: () => {
        if (!app.value.token) {
          return { name: 'login' }
        }
      },
      children: [
        {
          path: "blogs/list",
          name: "BackendBlogList",
          component: () => import('../views/backend/Blog/BlogListView.vue')
        },
        {
          path: "blogs/edit",
          name: "BackendBlogEdit",
          component: () => import('../views/backend/Blog/BlogEditView.vue')
        }
      ]
    }
  ]
})

export default router
