import { createWebHashHistory, createRouter, RouteRecordRaw } from 'vue-router'
const publicRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/views/main/main.vue'),
    children: [
      {
        path: '/switches',
        name: 'switches',
        component: () => import('@/views/switches/switches.vue')
      },
      {
        path: '/arplist',
        name: 'arplist',
        component: () => import('@/views/arp/arp.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: publicRoutes
})
export default router
