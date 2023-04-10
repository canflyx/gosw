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
      },
      {
        path: '/culist',
        name: 'culist',
        component: () => import('@/views/culist/culist.vue')
      },
      {
        path: '/cuscan',
        name: 'cuscan',
        component: () => import('@/views/cuscan/cuscan.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: publicRoutes
})
export default router
