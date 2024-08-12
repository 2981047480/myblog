import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: 'login',
      component: () => import("@/views/login/LoginView.vue")
    },
    {
      path: "/backend",
      name: 'backend',
      component: () => import('@/views/backend/BackendView.vue'),
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
