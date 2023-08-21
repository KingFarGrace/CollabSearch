import { createRouter, createWebHistory } from 'vue-router'
import { useAccountStore } from '@/stores/account'
import { storeToRefs } from 'pinia'

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
router.beforeEach((to, from, next) => {
  const { isLogin } = storeToRefs(useAccountStore())

  const isPublicRoute = to.path === '/login' || to.path === '/reg'

  if (!isPublicRoute && !isLogin.value) {
    return next('/login')
  }

  next()
})
export default router
