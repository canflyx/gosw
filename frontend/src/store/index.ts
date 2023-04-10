import { defineStore } from 'pinia'
import {
  getPageListData,
  createPageDate,
  updatePageDate,
  deletePageDate,
  scanPageData
} from '@/service/main'

interface SwLister {
  id: number
  label: string
}

export const mainStore = defineStore('main', {
  state: () => {
    return {
      arpList: [],
      arpCount: 0,
      switchesList: [],
      switchesCount: 0,
      logList: [],
      logCount: 0,
      swList: [] as SwLister[],
      queryInfo: {},
      selectItem: [] as number[],
      tags: [] as any[],
      readCmd: []
    }
  },
  actions: {
    async getInitial() {
      const swData = await getPageListData('/switches/list', {
        page_size: 1000,
        page_number: 0
      })

      for (const sw of swData.data.items) {
        this.swList.push({ id: sw.ID, label: sw.ip })
      }
    },
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
      console.log(payload)
      if (payload.queryInfo) {
        queryInfo.value = this.queryInfo
      }
      let pageResult: any
      if (pageName === 'culog') {
        pageResult = await getPageListData(`maclist/log`, queryInfo)
      } else {
        pageResult = await getPageListData(`${pageName}/list`, queryInfo)
      }
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
          break
        case 'culog':
          this.logList = items
          this.logCount = total
      }
    },
    async deletePageDataAction(payload: any) {
      const { pageName, id } = payload
      const pageUrl = `/${pageName}/${id}`
      await deletePageDate(pageUrl)
      this.getPageListAction({ pageName: pageName })
    },
    scanPageDataAction(payload: any) {
      if (this.selectItem.length > 0) {
        const pageUrl = '/maclist/scan'
        scanPageData(pageUrl, {
          list: this.selectItem,
          read_cmd: this.readCmd,
          flag: payload.flag
        })
      }
    },
    scanSwDataAction() {
      if (
        this.tags.length < 1 ||
        !Array.isArray(this.readCmd) ||
        this.readCmd.length < 1
      ) {
        return
      }
      this.selectItem = []
      for (const tag of this.tags) {
        this.selectItem.push(tag.id)
      }
      this.scanPageDataAction({ flag: 2 })
    }
  }
})
