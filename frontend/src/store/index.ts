import { defineStore } from 'pinia'
import {
  getPageListData,
  createPageDate,
  updatePageDate,
  deletePageDate,
  scanPageData
} from '@/service/main'

interface IQueryInfo {
  pageSize: number
  currentPage: number
  kws: object
}

export const mainStore = defineStore('main', {
  state: () => {
    return {
      arpList: [],
      arpCount: 0,
      switchesList: [],
      switchesCount: 0,
      queryInfo: {},
      selectItem: [] as number[]
    }
  },
  actions: {
    async updatePageDataAction(payload: any) {
      const { pageName, id, data } = payload
      const pageUrl = `/${pageName}/${id}`
      await updatePageDate(pageUrl, data)
      this.getPageListAction({ pageName: pageName })
    },
    async createPageDataAction(payload: any) {
      const pageName = payload.pageName
      await createPageDate(`/${pageName}`, payload.data)
      this.getPageListAction({ pageName: pageName })
    },
    async getPageListAction(payload: any) {
      const pageName = payload.pageName
      const queryInfo = payload.queryInfo
      if (payload.queryInfo) {
        queryInfo.value = this.queryInfo
      }
      const pageResult: any = await getPageListData(
        `${pageName}/list`,
        queryInfo
      )
      this.queryInfo = payload.queryInfo
      const { total, items } = pageResult.data
      switch (pageName) {
        case 'maclist':
          this.arpList = items
          this.arpCount = total
          break
        case 'switches':
          this.switchesList = items
          this.switchesCount = total
      }
    },
    async deletePageDataAction(payload: any) {
      const { pageName, id } = payload
      const pageUrl = `/${pageName}/${id}`
      await deletePageDate(pageUrl)
      this.getPageListAction({ pageName: pageName })
    },
    scanPageDataAction() {
      if (this.selectItem.length > 0) {
        const pageUrl = '/maclist/scan'
        scanPageData(pageUrl, { list: this.selectItem })
      }
    }
  }
})
