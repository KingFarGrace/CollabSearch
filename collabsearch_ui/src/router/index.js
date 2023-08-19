import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('@/views/Index.vue'),
      redirect: '/user'
    },
    {
      path: '/login',
      component: () => import('@/views/account/LoginPage.vue')
    },
    {
      path: '/reg',
      component: () => import('@/views/account/RegisterPage.vue')
    },
    {
      path: '/user',
      component: () => import('@/views/account/AccountPage.vue')
    },
    {
      path: '/workspace',
      component: () => import('@/views/workspace/WorkspaceList.vue')
    },
    {
      path: '/workspace/handler',
      component: () => import('@/views/workspace/WorkspaceHandlerList.vue')
    },
    {
      path: '/workspace/create',
      component: () => import('@/views/workspace/WorkspaceCreator.vue')
    },
    {
      path: '/workspace/search',
      component: () => import('@/views/workspace/WorkspacePage.vue')
    },
    {
      path: '/workspace/manage',
      component: () => import('@/views/workspace/WorkspaceManager.vue')
    },
    {
      path: '/workspace/detail',
      component: () => import('@/views/workspace/WorkspaceDetailPage.vue')
    }
  ]
})

export default router
